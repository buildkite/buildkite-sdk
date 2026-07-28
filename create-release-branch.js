const target = process.argv[2];

if (!/^\d+\.\d+\.\d+$/.test(target || "")) {
    console.error(
        `Usage: node create-release-branch.js <version>, got '${target}'`,
    );
    process.exit(1);
}

// Go is absent: it has no version file, its version is the sdk/go/vX.Y.Z tag.
const manifests = [
    {
        file: "sdk/typescript/package.json",
        pattern: /("version":\s*")(\d+\.\d+\.\d+)(")/,
    },
    {
        file: "sdk/python/pyproject.toml",
        pattern: /^(version = ")(\d+\.\d+\.\d+)(")/m,
    },
    {
        file: "sdk/ruby/lib/buildkite/version.rb",
        pattern: /(VERSION = ")(\d+\.\d+\.\d+)(")/,
    },
    {
        file: "sdk/csharp/src/Buildkite.Sdk/Buildkite.Sdk.csproj",
        pattern: /(<Version>)(\d+\.\d+\.\d+)(<\/Version>)/,
    },
];

(async () => {
    const fs = await import("fs");
    const { simpleGit } = await import("simple-git");
    const { execSync } = await import("child_process");

    const git = simpleGit();
    const branch = `release/v${target}`;

    await git.checkoutLocalBranch(branch);

    // Matches the version field, not the version string, so a dependency
    // pinned at the same number survives.
    for (const { file, pattern } of manifests) {
        const before = fs.readFileSync(file, "utf-8");
        const match = pattern.exec(before);

        if (!match) {
            console.error(`No version field found in ${file}`);
            process.exit(1);
        }

        fs.writeFileSync(file, before.replace(pattern, `$1${target}$3`));
        console.log(`${file}: ${match[2]} -> ${target}`);
    }

    // Build all SDKs.
    execSync("npm run build", { stdio: "inherit" });

    // Commit and push.
    await git.add("sdk"); // Include everything here, as lockfiles will also have changed.
    await git.commit(`Release v${target}`);
    await git.push("origin", branch);

    console.log("Release branch created");
})();
