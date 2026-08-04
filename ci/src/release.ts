import { Pipeline } from "@buildkite/buildkite-sdk";
import * as fs from "fs";
import { execFileSync } from "child_process";

const pipeline = new Pipeline();

const plugins = [
    {
        "docker#v5.11.0": {
            image: "buildkite-sdk-tools:latest",
            "propagate-environment": true,
            environment: [
                "GITHUB_TOKEN",
                "NPM_TOKEN",
                "PYPI_TOKEN",
                "GEM_HOST_API_KEY",
                "NUGET_API_KEY",
            ],
        },
    },
    { "rubygems-oidc#v0.2.0": { role: "rg_oidc_akr_emf87k6zphtb7x7adyrk" } },
    {
        "aws-assume-role-with-web-identity#v1.4.0": {
            "role-arn":
                "arn:aws:iam::597088016345:role/pipeline-buildkite-buildkite-sdk",
            "session-tags": [
                "organization_slug",
                "organization_id",
                "pipeline_slug",
            ],
        },
    },
    {
        "aws-ssm#v1.0.0": {
            parameters: {
                NPM_TOKEN: "/prod/buildkite-sdk/npm-token",
                PYPI_TOKEN: "/prod/buildkite-sdk/pypi-token",
                GITHUB_TOKEN: "/prod/buildkite-sdk/github-token",
                NUGET_API_KEY: "/prod/buildkite-sdk/nuget-api-key",
            },
        },
    },
];

pipeline.addStep({
    key: "install",
    label: ":test_tube: Install",
    plugins: [
        ...plugins,
        {
            "artifacts#v1.9.2": {
                upload: ["node_modules"],
                compressed: "node_modules.tgz",
            },
        },
    ],
    commands: ["mise trust", "npm install --ignore-scripts"],
});

const languagePlugins = [
    ...plugins,
    {
        "artifacts#v1.9.2": {
            download: ["node_modules"],
            compressed: "node_modules.tgz",
        },
    },
];

const languageTargets = [
    {
        icon: ":typescript:",
        label: "Typescript",
        key: "typescript",
        sdkLabel: "sdk-typescript",
        appLabel: "app-typescript",
    },
    {
        icon: ":python:",
        label: "Python",
        key: "python",
        sdkLabel: "sdk-python",
        appLabel: "app-python",
    },
    {
        icon: ":go:",
        label: "Go",
        key: "go",
        sdkLabel: "sdk-go",
        appLabel: "app-go",
    },
    {
        icon: ":ruby:",
        label: "Ruby",
        key: "ruby",
        sdkLabel: "sdk-ruby",
        appLabel: "app-ruby",
    },
    {
        icon: ":csharp:",
        label: "C#",
        key: "csharp",
        sdkLabel: "sdk-csharp",
        appLabel: "app-csharp",
    },
];

function git(...args: string[]): string {
    return execFileSync("git", args, {
        encoding: "utf-8",
        maxBuffer: 64 * 1024 * 1024,
    });
}

// Each returns a URL that is 200 when the version is published, 404 when not.
const REGISTRY_URL: Record<string, (version: string) => string> = {
    typescript: (v) =>
        `https://registry.npmjs.org/@buildkite%2Fbuildkite-sdk/${v}`,
    python: (v) => `https://pypi.org/pypi/buildkite-sdk/${v}/json`,
    ruby: (v) =>
        `https://rubygems.org/api/v2/rubygems/buildkite-sdk/versions/${v}.json`,
    csharp: (v) =>
        `https://api.nuget.org/v3-flatcontainer/buildkite.sdk/${v}/buildkite.sdk.${v}.nupkg`,
    go: (v) =>
        `https://proxy.golang.org/github.com/buildkite/buildkite-sdk/sdk/go/@v/v${v}.info`,
};

const RELEASE_TAG = /\/v(\d+\.\d+\.\d+)$/;

const MANIFEST: Record<string, { file: string; pattern: RegExp }> = {
    typescript: {
        file: "sdk/typescript/package.json",
        pattern: /"version":\s*"(\d+\.\d+\.\d+)"/,
    },
    python: {
        // Scoped to [project], so a version key in another table cannot win.
        file: "sdk/python/pyproject.toml",
        pattern: /\[project\][\s\S]*?^version\s*=\s*"(\d+\.\d+\.\d+)"/m,
    },
    ruby: {
        file: "sdk/ruby/lib/buildkite/version.rb",
        pattern: /VERSION\s*=\s*"(\d+\.\d+\.\d+)"/,
    },
    csharp: {
        file: "sdk/csharp/src/Buildkite.Sdk/Buildkite.Sdk.csproj",
        pattern: /<Version>(\d+\.\d+\.\d+)<\/Version>/,
    },
};

// All tags, not just those merged into HEAD: a squash or rebase merge leaves
// the release tag unreachable, and nx resolves the same way via
// checkAllBranchesWhen. An unpublished tag must point at HEAD before it ships.
function taggedVersion(key: string): string | null {
    const versions = git("tag", "--list", `sdk/${key}/v*`)
        .split("\n")
        .map((line) => RELEASE_TAG.exec(line.trim()))
        .filter((match): match is RegExpExecArray => match !== null)
        .map((match) => match[1].split(".").map(Number));

    if (!versions.length) {
        return null;
    }

    versions.sort((a, b) => b[0] - a[0] || b[1] - a[1] || b[2] - a[2]);
    return versions[0].join(".");
}

// Throws rather than returning null: every failure here means we cannot tell
// what would ship, and skipping on that is how a release goes missing quietly.
function releaseVersion(key: string): string {
    if (key === "go") {
        const version = taggedVersion(key);
        if (!version) {
            throw new Error("no sdk/go/v* tag found");
        }
        return version;
    }

    const manifest = MANIFEST[key];
    if (!manifest) {
        throw new Error(`no version source configured for "${key}"`);
    }

    let contents: string;
    try {
        contents = fs.readFileSync(manifest.file, "utf-8");
    } catch {
        throw new Error(`could not read ${manifest.file}`);
    }

    const match = manifest.pattern.exec(contents);
    if (!match) {
        throw new Error(
            `no X.Y.Z version in ${manifest.file}; prereleases are not supported`,
        );
    }

    return match[1];
}

// An unpublished version must be built from the exact commit its tag names.
// This also covers shared build inputs outside the SDK directory.
function reproducibilityProblem(key: string, version: string): string | null {
    const tag = `sdk/${key}/v${version}`;

    if (!git("tag", "--list", tag).trim()) {
        return `${tag} does not exist; push the release tag`;
    }

    const taggedCommit = git("rev-parse", `${tag}^{commit}`).trim();
    const head = git("rev-parse", "HEAD").trim();

    if (taggedCommit !== head) {
        return (
            `${tag} points to ${taggedCommit.slice(0, 12)}, not this build ` +
            `${head.slice(0, 12)}; trigger the pipeline at the tagged commit`
        );
    }

    return null;
}

function isPublished(key: string, version: string): boolean | null {
    try {
        const status = execFileSync(
            "curl",
            [
                "-s",
                "-o",
                "/dev/null",
                "-w",
                "%{http_code}",
                "--max-time",
                "20",
                "--retry",
                "2",
                "--retry-delay",
                "1",
                REGISTRY_URL[key](version),
            ],
            { encoding: "utf-8" },
        ).trim();

        if (status === "200") return true;
        if (status === "404") return false;
        return null;
    } catch {
        return null;
    }
}

// Buildkite can reuse a checkout whose tags are missing, stale, or deleted.
// The remote tag set is authoritative for release planning.
try {
    execFileSync(
        "git",
        ["fetch", "--force", "--prune", "--prune-tags", "--tags", "origin"],
        { stdio: "inherit" },
    );
} catch {
    console.error("Could not refresh release tags from origin.");
    process.exit(1);
}

if (!git("tag", "--list", "sdk/*/v*").trim()) {
    console.error(
        "No sdk/<language>/v* tags exist after fetching from origin.",
    );
    process.exit(1);
}

const plans = languageTargets.map((target) => {
    let version: string;
    try {
        version = releaseVersion(target.key);
    } catch (error) {
        console.error(`${target.label}: ${(error as Error).message}`);
        process.exit(1);
    }

    const published = isPublished(target.key, version);
    if (published === true) {
        return { key: target.key, skip: `${version} is already published` };
    }
    if (published === null) {
        console.warn(
            `  ${target.label}: registry lookup failed, publishing ${version} anyway`,
        );
    }

    // Only a version that is about to ship needs to match its tag. Checking
    // earlier would let one SDK's drift block every other SDK's release.
    const blocked = reproducibilityProblem(target.key, version);

    return { key: target.key, version, blocked };
});

const planFor = new Map(plans.map((plan) => [plan.key, plan]));

languageTargets.forEach((target) => {
    const plan = planFor.get(target.key);
    const skipPublish = plan && "skip" in plan ? plan.skip : undefined;
    const blocked = plan && "blocked" in plan ? plan.blocked : null;

    console.log(
        skipPublish
            ? `  ${target.label}: skipping publish (${skipPublish})`
            : blocked
              ? `  ${target.label}: BLOCKED - ${blocked}`
              : `  ${target.label}: publishing ${plan && "version" in plan ? plan.version : "?"}`,
    );

    pipeline.addStep({
        depends_on: "install",
        key: `${target.key}`,
        group: `${target.icon} ${target.label}`,
        steps: [
            {
                key: `${target.key}-test`,
                label: ":test_tube: Test",
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx test ${target.sdkLabel}`,
                ],
            },
            {
                key: `${target.key}-publish`,
                label: ":rocket: Publish",
                depends_on: [`${target.key}-test`],
                ...(skipPublish ? { skip: skipPublish } : {}),
                ...(!blocked && target.key === "go" && plan && "version" in plan
                    ? { env: { GO_RELEASE_VERSION: plan.version } }
                    : {}),
                plugins: languagePlugins,
                // A blocked SDK fails its own step rather than skipping, so it is
                // red rather than quietly absent, and the others still publish.
                commands: blocked
                    ? [`echo ${JSON.stringify(blocked)}`, "exit 1"]
                    : [
                          "mise trust",
                          `nx install ${target.sdkLabel}`,
                          `nx build ${target.sdkLabel}`,
                          `nx run ${target.sdkLabel}:publish`,
                      ],
            },
        ],
    });
});

fs.writeFileSync(".buildkite/pipeline.json", pipeline.toJSON());
