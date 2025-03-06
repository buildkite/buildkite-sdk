require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } }
]

pipeline.add_step(
  label: ":test_tube: Test",
  plugins: plugins,
  commands: [
    "mise trust",
    "npm install",
    "npm test",
  ],
  artifact_paths: [
    "node_modules/**/*"
  ]
)

artifact_plugin = { "artifacts#v1.9.2": { download: [
  "node_modules/**/*"
] } }

pipeline.add_step(
  label: ":test_tube: Build",
  plugins: [
    *plugins,
    artifact_plugin
  ],
  commands: [
    "mise trust",
    "npm run build",
  ]
)

pipeline.add_step(
  label: ":test_tube: Docs",
  plugins: [
    *plugins,
    artifact_plugin
  ],
  commands: [
    "mise trust",
    "npm run docs",
  ]
)

pipeline.add_step(
  label: ":test_tube: Apps",
  plugins: [
    *plugins,
    artifact_plugin
  ],
  commands: [
    "mise trust",
    "npm run apps",
  ]
)

puts pipeline.to_json
