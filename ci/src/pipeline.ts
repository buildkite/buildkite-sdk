import { Pipeline } from "@buildkite/buildkite-sdk";
import * as fs from "fs";
import toml from "toml";

const pipeline = new Pipeline();
const plugins = [
    {
        "docker#v5.11.0": {
            image: "buildkite-sdk-tools:latest",
            "propagate-environment": true,
        },
    },
];

function getMiseConfig(): {
    node: string[];
    go: string[];
    python: string[];
    ruby: string[];
} {
    try {
        const data = fs.readFileSync("./mise.apps.toml");
        const config = toml.parse(data.toString());
        return config.tools;
    } catch (parseErr) {
        console.log("Error parsing TOML:", parseErr);
    }
}

const languageVersions = getMiseConfig();

// Install
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

// Check for version updates
const misePackages = ["node", "python", "go", "ruby"];
pipeline.addStep({
    key: "version-check",
    group: ":heavy_check_mark: Version Check",
    steps: misePackages.map((pkg) => ({
        plugins,
        key: `${pkg}-version-check`,
        label: `:${pkg}: ${pkg}`,
        soft_fail: true,
        command: [
            "mise trust",
            `mise upgrade ${pkg} --bump`,
            "exit $(git diff --exit-code)",
        ],
    })),
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

interface Target {
    icon: string;
    label: string;
    key: string;
    sdkLabel: string;
    appLabel: string;
    versions: string[];
}

const languageTargets: Target[] = [
    {
        icon: ":typescript:",
        label: "Typescript",
        key: "typescript",
        sdkLabel: "sdk-typescript",
        appLabel: "app-typescript",
        versions: languageVersions["node"],
    },
    {
        icon: ":python:",
        label: "Python",
        key: "python",
        sdkLabel: "sdk-python",
        appLabel: "app-python",
        versions: languageVersions["python"],
    },
    {
        icon: ":go:",
        label: "Go",
        key: "go",
        sdkLabel: "sdk-go",
        appLabel: "app-go",
        versions: languageVersions["go"],
    },
    {
        icon: ":ruby:",
        label: "Ruby",
        key: "ruby",
        sdkLabel: "sdk-ruby",
        appLabel: "app-ruby",
        versions: languageVersions["ruby"],
    },
];

function generateAppCommands(key: string, appLabel: string) {
    let language = key;
    if (key === "typescript") {
        language = "node";
    }

    let appInstallCommand = `mise exec ${language}@{{matrix}} -- nx install ${appLabel}`;
    if (language === "python") {
        appInstallCommand = `mise exec ${language}@{{matrix}} -- pip install --no-cache-dir uv black && nx install ${appLabel}`;
    }
    if (language === "node") {
        appInstallCommand = `mise exec ${language}@{{matrix}} -- npm install -g nx && npm install && nx install ${appLabel}`;
    }

    return [
        "mise trust mise.apps.toml",
        `mise install ${language}@{{matrix}}`,
        appInstallCommand,
        `mise exec ${language}@{{matrix}} -- nx run ${appLabel}:run`,
    ];
}

languageTargets.forEach((target) => {
    pipeline.addStep({
        depends_on: ["install"],
        key: `${target.key}`,
        group: `${target.icon} ${target.label}`,
        steps: [
            {
                key: `${target.key}-diff`,
                label: ":git: Diff",
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    "nx gen:build",
                    `nx gen:types-${target.key}`,
                    "export DIFF=$(git diff --exit-code)",
                    "exit $DIFF",
                ],
            },
            {
                key: `${target.key}-test`,
                label: ":test_tube: Test",
                plugins: languagePlugins,
                env: {
                    // Jest Issue: https://github.com/jestjs/jest/issues/15888
                    // Node Issue: https://github.com/nodejs/node/issues/60704
                    NODE_OPTIONS: "--localstorage-file=./jest-storage",
                },
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx test ${target.sdkLabel}`,
                ],
            },
            {
                key: `${target.key}-build`,
                label: ":package: Build",
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx build ${target.sdkLabel}`,
                ],
            },
            {
                key: `${target.key}-docs`,
                label: ":books: Docs",
                depends_on: [`${target.key}-test`, `${target.key}-build`],
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx run ${target.sdkLabel}:docs:build`,
                ],
            },
            {
                label: ":lab_coat: Apps",
                key: `${target.key}-apps`,
                depends_on: [`${target.key}-test`, `${target.key}-build`],
                plugins: languagePlugins,
                commands: generateAppCommands(target.key, target.appLabel),
                matrix: target.versions,
            },
        ],
    });
});

fs.writeFileSync(".buildkite/pipeline.json", pipeline.toJSON());
