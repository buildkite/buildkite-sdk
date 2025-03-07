require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } }
]

pipeline.add_step(
  key: "test",
  label: ":test_tube: Test",
  plugins: plugins,
  commands: [
    "npm test"
  ]
)

pipeline.add_step(
  label: ":test_tube: Build",
  plugins: plugins,
  commands: [
    "npm run build"
  ]
)

pipeline.add_step(
  label: ":test_tube: Docs",
  key: "docs",
  depends_on: ["build","test"],
  plugins: plugins,
  commands: [
    "npm run docs"
  ]
)

pipeline.add_step(
  label: ":test_tube: Apps",
  key: "apps",
  depends_on: ["build","test"],
  plugins: plugins,
  commands: [
    "npm run apps"
  ]
)

puts pipeline.to_json
