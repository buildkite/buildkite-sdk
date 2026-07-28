import { Pipeline } from "@buildkite/buildkite-sdk"
import * as fs from "fs"
import { execSync } from "child_process"

const pipeline = new Pipeline()

const plugins = [
  { "docker#v5.11.0": {
     image: "buildkite-sdk-tools:latest",
     "propagate-environment": true,
     environment: [
       "GITHUB_TOKEN",
       "NPM_TOKEN",
       "PYPI_TOKEN",
       "GEM_HOST_API_KEY",
       "NUGET_API_KEY"
     ]
  }},
  { "rubygems-oidc#v0.2.0": { role: "rg_oidc_akr_emf87k6zphtb7x7adyrk" } },
  { "aws-assume-role-with-web-identity#v1.4.0": {
    "role-arn": "arn:aws:iam::597088016345:role/pipeline-buildkite-buildkite-sdk",
    "session-tags": ["organization_slug", "organization_id", "pipeline_slug"],
  }},
  { "aws-ssm#v1.0.0": {
    parameters: {
      NPM_TOKEN: "/prod/buildkite-sdk/npm-token",
      PYPI_TOKEN: "/prod/buildkite-sdk/pypi-token",
      GITHUB_TOKEN: "/prod/buildkite-sdk/github-token",
      NUGET_API_KEY: "/prod/buildkite-sdk/nuget-api-key"
    }
  }}
]

pipeline.addStep({
    key: "install",
    label: ":test_tube: Install",
    plugins: [
        ...plugins,
        { "artifacts#v1.9.2": {
            upload: ["node_modules"],
            compressed: "node_modules.tgz"
        }}
    ],
    commands: [
        "mise trust",
        "npm install --ignore-scripts"
    ]
})

const languagePlugins = [
    ...plugins,
    { "artifacts#v1.9.2": {
        download: ["node_modules"],
        compressed: "node_modules.tgz"
    }}
]

const languageTargets = [
  {
    icon: ":typescript:",
    label: "Typescript",
    key: "typescript",
    sdkLabel: "sdk-typescript",
    appLabel: "app-typescript"
  },
  {
    icon: ":python:",
    label: "Python",
    key: "python",
    sdkLabel: "sdk-python",
    appLabel: "app-python"
  },
  {
    icon: ":go:",
    label: "Go",
    key: "go",
    sdkLabel: "sdk-go",
    appLabel: "app-go"
  },
  {
    icon: ":ruby:",
    label: "Ruby",
    key: "ruby",
    sdkLabel: "sdk-ruby",
    appLabel: "app-ruby"
  },
  {
    icon: ":csharp:",
    label: "C#",
    key: "csharp",
    sdkLabel: "sdk-csharp",
    appLabel: "app-csharp"
  }
]

// Skip publishing SDKs unchanged since the previous release tag. Lockstep bumps
// touch every manifest, so those don't count.

function git(args: string): string {
    return execSync(`git ${args}`, { encoding: "utf-8" }).trim();
}

// Files the release rewrites to carry the new version. Deliberately excludes
// lockfiles.
const VERSION_MANIFESTS = new Set([
    "sdk/go/project.json",
    "sdk/python/pyproject.toml",
    "sdk/typescript/package.json",
    "sdk/ruby/lib/buildkite/version.rb",
    "sdk/ruby/project.json",
    "sdk/csharp/src/Buildkite.Sdk/Buildkite.Sdk.csproj",
]);

function releasedVersions(): string[] {
    return git(
        "tag --list 'v[0-9]*.[0-9]*.[0-9]*' --sort=-v:refname --merged HEAD"
    )
        .split("\n")
        .filter(Boolean)
        .map((tag) => tag.slice(1));
}

function resolveBaseTag(versions: string[]): string | null {
    if (versions.length === 0) {
        return null;
    }

    // Tagged HEAD is the release itself, so diff from the one before.
    const headTags = git("tag --points-at HEAD").split("\n").filter(Boolean);
    const base = headTags.includes(`v${versions[0]}`)
        ? versions[1]
        : versions[0];

    return base ? `v${base}` : null;
}

function isVersionBumpOnly(
    base: string,
    file: string,
    versions: string[]
): boolean {
    if (!VERSION_MANIFESTS.has(file)) {
        return false;
    }

    const diff = git(`diff --unified=0 ${base}..HEAD -- "${file}"`);
    const removed: string[] = [];
    const added: string[] = [];

    for (const line of diff.split("\n")) {
        if (line.startsWith("---") || line.startsWith("+++")) continue;
        if (line.startsWith("-")) removed.push(line.slice(1));
        else if (line.startsWith("+")) added.push(line.slice(1));
    }

    if (removed.length === 0 || removed.length !== added.length) {
        return false;
    }

    // Blank out released versions before comparing.
    const strip = (line: string) =>
        versions.reduce((acc, version) => acc.split(version).join("\0"), line);

    return added.every((line, i) => strip(line) === strip(removed[i]));
}

function changedFiles(
    base: string,
    sdkPath: string,
    versions: string[]
): string[] {
    return git(`diff --name-only ${base}..HEAD -- ${sdkPath}`)
        .split("\n")
        .filter(Boolean)
        .filter((file) => !isVersionBumpOnly(base, file, versions));
}

const releasedVersionList = releasedVersions();
const baseTag = resolveBaseTag(releasedVersionList);

// Longest first so `0.12.0` is stripped before `0.1.0` matches inside it.
const strippableVersions = [...releasedVersionList].sort(
    (a, b) => b.length - a.length
);

console.log(
    baseTag
        ? `Gating publish steps on changes since ${baseTag}.`
        : "No previous release tag found. Publishing every SDK."
);

languageTargets.forEach((target) => {
    const sdkPath = `sdk/${target.key}`;
    const skipPublish =
        baseTag &&
        changedFiles(baseTag, sdkPath, strippableVersions).length === 0
            ? `No changes in ${sdkPath} since ${baseTag}`
            : undefined;

    console.log(
        skipPublish
            ? `  ${target.label}: skipping publish (${skipPublish})`
            : `  ${target.label}: publishing`
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
                `nx test ${target.sdkLabel}`
            ],
        },
        {
            key: `${target.key}-publish`,
            label: ":rocket: Publish",
            depends_on: [`${target.key}-test`],
            ...(skipPublish ? { skip: skipPublish } : {}),
            plugins: languagePlugins,
            commands: [
                "mise trust",
                `nx install ${target.sdkLabel}`,
                `nx build ${target.sdkLabel}`,
                `nx run ${target.sdkLabel}:publish`
            ],
        },
        ]
    })
})

fs.writeFileSync(".buildkite/pipeline.json", pipeline.toJSON())
