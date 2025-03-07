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
    "mise trust",
    "npm test"
  ]
)

pipeline.add_step(
  label: ":package: Build",
  plugins: plugins,
  commands: [
    "mise trust",
    "npm run build"
  ]
)

pipeline.add_step(
  label: ":books: Docs",
  key: "docs",
  depends_on: ["build","test"],
  plugins: plugins,
  commands: [
    "mise trust",
    "npm run docs"
  ]
)

pipeline.add_step(
  label: ":lab_coat: Apps",
  key: "apps",
  depends_on: ["build","test"],
  plugins: plugins,
  commands: [
    "mise trust",
    "npm run apps"
  ]
)

puts pipeline.to_json
