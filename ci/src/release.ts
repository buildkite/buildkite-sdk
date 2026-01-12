import { Pipeline } from "@buildkite/buildkite-sdk"
import * as fs from "fs"

const pipeline = new Pipeline()

const plugins = [
  { "docker#v5.11.0": {
     image: "buildkite-sdk-tools:latest",
     "propagate-environment": true,
     environment: [
       "GITHUB_TOKEN",
       "NPM_TOKEN",
       "PYPI_TOKEN",
       "GEM_HOST_API_KEY"
     ]
  }},
  { "rubygems-oidc#v0.2.0": { role: "rg_oidc_akr_emf87k6zphtb7x7adyrk" } },
  { "aws-assume-role-with-web-identity#v1.0.0": {
    "role-arn": "arn:aws:iam::597088016345:role/pipeline-buildkite-buildkite-sdk",
    "session-tags": ["organization_slug", "organization_id", "pipeline_slug"],
  }},
  { "aws-ssm#v1.0.0": {
    parameters: {
      NPM_TOKEN: "/prod/buildkite-sdk/npm-token",
      PYPI_TOKEN: "/prod/buildkite-sdk/pypi-token",
      GITHUB_TOKEN: "/prod/buildkite-sdk/github-token"
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
  }
]

languageTargets.forEach((target) => {
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
