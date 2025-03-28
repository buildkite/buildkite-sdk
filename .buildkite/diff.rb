require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

language_targets = {
  typescript: {
    icon: ":typescript:",
    label: "Typescript",
    key: "typescript",
    sdk_label: "sdk-typescript",
    app_label: "app-typescript"
  },
  python: {
    icon: ":python:",
    label: "Python",
    key: "python",
    sdk_label: "sdk-python",
    app_label: "app-python"
  },
  go: {
    icon: ":go:",
    label: "Go",
    key: "go",
    sdk_label: "sdk-go",
    app_label: "app-go"
  },
  ruby: {
    icon: ":ruby:",
    label: "Ruby",
    key: "ruby",
    sdk_label: "sdk-ruby",
    app_label: "app-ruby"
  }
}

pipeline = Buildkite::Pipeline.new

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest", "mount-buildkite-agent": true } }
]

pipeline.add_step(
  key: "install",
  label: ":test_tube: Install",
  plugins: [
    *plugins,
    { "artifacts#v1.9.2": {
      upload: ["node_modules"],
      compressed: "node_modules.tgz"
    }}
  ],
  commands: [
    "mise trust",
    "npm install --ignore-scripts"
  ]
)

pipeline.add_step(
  depends_on: "install",
  key: "upload-language-pipelines",
  label: ":buildkite: Generate Language Pipelines",
  plugins: [
    *plugins,
    { "artifacts#v1.9.2": {
      download: ["node_modules"],
      compressed: "node_modules.tgz"
    }},
    { "monorepo-diff#v1.3.0": {
      diff: "git diff --name-only main...HEAD",
      watch: [
        {
          path: "sdk/typescript",
          config: {
            command: "npx nx gen:pipeline-typescript | buildkite-agent pipeline upload",
          },
        },
      ]
    }},
  ]
)

puts pipeline.to_json
