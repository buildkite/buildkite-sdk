import { Pipeline } from "@buildkite/buildkite-sdk"
import * as fs from "fs"

const pipeline = new Pipeline()
const plugins = [
    { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } },
]

pipeline.addStep({
    key: "install",
    label: ":test_tube: Install",
    plugins: [
        ...plugins,
        { "artifacts#v1.9.2": {
            upload: ["node_modules"],
            compressed: "node_modules.tgz"
        }},
    ],
    commands: [
        "mise trust",
        "npm install --ignore-scripts",
    ],
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
  }
]

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
                    `nx install ${target.sdkLabel}`,
                    "nx gen:build",
                    `nx gen:types-${target.key}`,
                    "exit $(git diff --exit-code)",
                ],
            },
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
                key: `${target.key}-build`,
                label: ":package: Build",
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx build ${target.sdkLabel}`
                ],
            },
            {
                key: `${target.key}-docs`,
                label: ":books: Docs",
                depends_on: [`${target.key}-test`,`${target.key}-build`],
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.sdkLabel}`,
                    `nx run ${target.sdkLabel}:docs:build`
                ],
            },
            {
                label: ":lab_coat: Apps",
                key: `${target.key}-apps`,
                depends_on: [`${target.key}-test`,`${target.key}-build`],
                plugins: languagePlugins,
                commands: [
                    "mise trust",
                    `nx install ${target.appLabel}`,
                    `nx run ${target.appLabel}:run`
                ],
            },
        ]
    })
})

fs.writeFileSync(".buildkite/pipeline.json", pipeline.toJSON())
