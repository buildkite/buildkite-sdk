const { execFileSync, spawnSync } = require("child_process");

const SDKS = ["typescript", "python", "go", "ruby", "csharp"];
const BUMPS = ["patch", "minor", "major"];
const RELEASE_TAG = /\/v(\d+\.\d+\.\d+)$/;
// The bump is a flag, not a positional: nx parses a bare positional as a
// project name and never forwards it here.
const USAGE =
    "Usage: nx release:create [--bump=patch|minor|major] [--dry-run] [--force]";

let bump = "patch";
let dryRun = false;
let force = false;

// Unknown arguments are rejected rather than ignored: silently dropping
// --dry-run would turn a preview into a real release.
for (const arg of process.argv.slice(2)) {
    if (!arg) continue;

    const bumpFlag = /^--bump=(.+)$/.exec(arg);

    if (arg === "--dry-run") {
        dryRun = true;
    } else if (arg === "--force") {
        force = true;
    } else if (bumpFlag && BUMPS.includes(bumpFlag[1])) {
        bump = bumpFlag[1];
    } else {
        console.error(`Unknown argument '${arg}'.\n${USAGE}`);
        process.exit(1);
    }
}

function git(...args) {
    return execFileSync("git", args, {
        encoding: "utf-8",
        maxBuffer: 64 * 1024 * 1024,
    });
}

// Required even for a dry run: detection compares <tag>..HEAD, so uncommitted
// work is invisible and the preview would report "unchanged" for an SDK you
// have just edited.
if (git("status", "--porcelain").trim()) {
    console.error("Working tree is not clean. Commit or stash first.");
    process.exit(1);
}

// Tag resolution deliberately considers unmerged and squashed branches, so a
// stale local tag set produces the wrong version. Fetch rather than suggest it.
// No --force: a local tag that disagrees with the remote is a conflict to look
// at, not something to silently overwrite with release tags in play.
// git writes its per-ref report to stderr, so both streams are read to surface
// which tag conflicts rather than only that the fetch failed.
const fetch = spawnSync("git", ["fetch", "origin", "--tags"], {
    encoding: "utf-8",
});
const report = `${fetch.stdout || ""}${fetch.stderr || ""}`.trim();

if (report.includes("[rejected]")) {
    console.error(
        "A local tag disagrees with the remote. Releasing from either would " +
            "produce the wrong version, so resolve it first.\n\n" +
            report,
    );
    process.exit(1);
}

if (fetch.error || fetch.status !== 0) {
    console.error(
        "Could not fetch tags, so the local tag set may be stale." +
            (report ? `\n\n${report}` : ""),
    );
    process.exit(1);
}

if (!dryRun) {
    const branch = git("branch", "--show-current").trim() || "(detached HEAD)";
    const head = git("rev-parse", "HEAD").trim();
    const remoteMain = git("rev-parse", "refs/remotes/origin/main").trim();

    if (branch !== "main" || head !== remoteMain) {
        console.error(
            "Create releases only from main at the current origin/main commit.\n" +
                `Current branch: ${branch}\n` +
                `HEAD: ${head}\n` +
                `origin/main: ${remoteMain}\n\n` +
                "Switch to main and update it before trying again.",
        );
        process.exit(1);
    }
}

// All tags, not just those merged into HEAD. A squash or rebase merge leaves
// the release tag unreachable, and nx resolves the same way via
// checkAllBranchesWhen, so both must agree on which tag is newest.
function latestTag(key) {
    const tags = git("tag", "--list", `sdk/${key}/v*`)
        .split("\n")
        .map((line) => ({
            tag: line.trim(),
            match: RELEASE_TAG.exec(line.trim()),
        }))
        .filter((entry) => entry.match)
        .map((entry) => ({
            tag: entry.tag,
            parts: entry.match[1].split(".").map(Number),
        }));

    if (!tags.length) {
        return null;
    }

    tags.sort(
        (a, b) =>
            b.parts[0] - a.parts[0] ||
            b.parts[1] - a.parts[1] ||
            b.parts[2] - a.parts[2],
    );

    return tags[0].tag;
}

// project.json is excluded from selection. It sets build and pack commands, so
// the release pipeline does count it when checking an artifact against its tag,
// but treating a build-config edit as reason to release produces exactly the
// no-op version bumps this tooling exists to avoid. --force includes it, for a
// deliberate packaging release.
function changedSince(key, tag) {
    const paths = [`sdk/${key}`];
    if (!force) {
        paths.push(`:(exclude)sdk/${key}/project.json`);
    }

    return git("diff", "--name-only", `${tag}..HEAD`, "--", ...paths)
        .split("\n")
        .filter(Boolean);
}

const changed = [];

for (const key of SDKS) {
    const tag = latestTag(key);

    // Every SDK has a baseline tag in this repo, so a missing one means the
    // tags were not fetched. Releasing everything on that assumption would
    // bump SDKs that have not changed.
    if (!tag) {
        console.error(
            `\nNo sdk/${key}/v* tag found. Fetch tags before releasing:\n` +
                "  git fetch --tags",
        );
        process.exit(1);
    }

    const files = changedSince(key, tag);
    if (files.length) {
        console.log(`  ${key}: ${files.length} file(s) changed since ${tag}`);
        changed.push(key);
    } else {
        console.log(`  ${key}: unchanged since ${tag}`);
    }
}

if (!changed.length) {
    console.log(
        force
            ? "\nNothing to release."
            : "\nNothing to release. Use --force to release build-config changes.",
    );
    process.exit(0);
}

const projects = changed.map((key) => `sdk-${key}`).join(",");
console.log(`\nReleasing ${bump}${dryRun ? " (dry run)" : ""}: ${projects}\n`);

execFileSync(
    "npx",
    [
        "nx",
        "release",
        bump,
        `--projects=${projects}`,
        "--skip-publish",
        ...(dryRun ? ["--dry-run"] : []),
    ],
    { stdio: "inherit" },
);

if (!dryRun) {
    console.log(
        "\nReview the release commit and tags, then push them:\n" +
            "  git show --stat\n" +
            "  git push --follow-tags",
    );
}
