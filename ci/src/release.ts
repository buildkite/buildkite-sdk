import { Pipeline } from "@buildkite/buildkite-sdk"
import * as fs from "fs"
import { execFileSync } from "child_process"

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

function git(...args: string[]): string {
    return execFileSync("git", args, {
        encoding: "utf-8",
        maxBuffer: 64 * 1024 * 1024,
    });
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

const RELEASE_TAG = /^v(\d+)\.(\d+)\.(\d+)$/;

interface Release {
    tag: string;
    parts: number[];
}

function compareParts(a: number[], b: number[]): number {
    return a[0] - b[0] || a[1] - b[1] || a[2] - b[2];
}

function releases(): Release[] {
    return git("tag", "--list", "--merged", "HEAD")
        .split("\n")
        .map((line) => RELEASE_TAG.exec(line.trim()))
        .filter((match): match is RegExpExecArray => match !== null)
        .map((match) => ({
            tag: match[0],
            parts: [+match[1], +match[2], +match[3]],
        }))
        .sort((a, b) => compareParts(b.parts, a.parts));
}

function pendingVersion(): number[] | null {
    const manifest = JSON.parse(
        fs.readFileSync("sdk/typescript/package.json", "utf-8")
    );
    const match = RELEASE_TAG.exec(`v${manifest.version}`);
    return match ? [+match[1], +match[2], +match[3]] : null;
}

function resolveBaseTag(pending: number[] | null): string | null {
    if (!pending) {
        return null;
    }

    const previous = releases().find(
        (release) => compareParts(release.parts, pending) < 0
    );

    return previous ? previous.tag : null;
}

function fileAt(rev: string, file: string): string | null {
    try {
        return git("show", `${rev}:${file}`);
    } catch {
        return null;
    }
}

function isVersionBumpOnly(
    base: string,
    file: string,
    versions: string[]
): boolean {
    if (!VERSION_MANIFESTS.has(file)) {
        return false;
    }

    const before = fileAt(base, file);
    const after = fileAt("HEAD", file);

    if (before === null || after === null) {
        return false;
    }

    const strip = (text: string) =>
        versions.reduce((acc, version) => acc.split(version).join("\0"), text);

    return strip(before) === strip(after);
}

function changedFiles(
    base: string,
    sdkPath: string,
    versions: string[]
): string[] {
    return git("diff", "--name-only", `${base}..HEAD`, "--", sdkPath)
        .split("\n")
        .map((line) => line.trim())
        .filter(Boolean)
        .filter((file) => !isVersionBumpOnly(base, file, versions));
}

const pending = pendingVersion();
const baseTag = resolveBaseTag(pending);

const strippableVersions = [
    ...new Set([
        ...releases().map((release) => release.parts.join(".")),
        ...(pending ? [pending.join(".")] : []),
    ]),
].sort((a, b) => b.length - a.length);

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
