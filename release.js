const [from, to] = [process.argv[2], process.argv[3]];

if (!from || !to) {
    console.error(`Invalid 'from' or 'to' version: '${from}', '${to}'`);
    process.exit(1);
}

(async () => {
    const { replaceInFileSync } = await import("replace-in-file");
    const { simpleGit } = await import("simple-git");
    const { execSync } = await import("child_process");
    const { Octokit } = await import("@octokit/rest");

    const paths = [
        "sdk/go/project.json",
        "sdk/python/pyproject.toml",
        "sdk/typescript/package.json",
        "sdk/ruby/lib/buildkite/version.rb",
        "sdk/ruby/project.json",
    ];

    const git = simpleGit();
    const branch = `release/v${to}`

    await git.checkoutLocalBranch(branch)

    // Bump versions.
    replaceInFileSync({
        files: paths,
        from,
        to,
    });

    // Build all SDKs.
    execSync("npm run build", { stdio: "inherit" });

    // Commit and tag.
    await git.add("sdk"); // Include everything here, as lockfiles will also have changed.
    await git.add("project.json"); // As this contains the new version.
    await git.commit(`Release v${to}`);
    await git.addTag(`v${to}`);
    await git.addTag(`sdk/go/v${to}`);

    // Push the commit and tags. This is what triggers publishing.
    await git.push("origin", branch, { "--tags": true });

    // Auth with GitHub.
    const octokit = new Octokit({
        auth: process.env.GITHUB_TOKEN,
    });

    // Create a GitHub release.
    const response = await octokit.rest.repos.createRelease({
        owner: "buildkite",
        repo: "buildkite-sdk",
        tag_name: `v${to}`,
        name: `Release v${to}`,
        body: [
            `* https://www.npmjs.com/package/@buildkite/buildkite-sdk/v/${to}`,
            `* https://pypi.org/project/buildkite-sdk/${to}/`,
            `* https://pkg.go.dev/github.com/buildkite/buildkite-sdk/sdk/go@v${to}`,
            `* https://rubygems.org/gems/buildkite-sdk/versions/${to}`,
        ].join("\n"),
        draft: false,
        prerelease: false,
    });

    console.log("Release created successfully:", response.data.html_url);
})();
