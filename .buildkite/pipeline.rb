require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } }
]

pipeline.add_step(
  key: "install",
  label: ":test_tube: Install",
  plugins: [
    *plugins,
    { "artifacts#v1.9.2": {
      upload: ["node_modules"]
      compressed: node_modules.tgz
    }}
  ],
  commands: [
    "mise trust",
    "npm install --ignore-scripts"
  ]
)

language_plugins = [
  *plugins,
  { "artifacts#v1.9.2": {
    download: ["node_modules"]
    compressed: node_modules.tgz
  }}
]

pipeline.add_step(
  key: "node",
  group: ":typescript: TypeScript",
  steps: [
    {
      key: "test",
      label: ":test_tube: Test",
      plugins: language_plugins,
      commands: [
        "mise trust",
        "nx install:typescript --verbose",
        "nx test:typescript"
      ],
    },
  ]
)

# pipeline.add_step(
#   key: "test",
#   label: ":test_tube: Test",
#   plugins: plugins,
#   commands: [
#     "mise trust",
#     "npm test"
#   ]
# )

# pipeline.add_step(
#   label: ":package: Build",
#   plugins: plugins,
#   commands: [
#     "mise trust",
#     "npm run build"
#   ]
# )

# pipeline.add_step(
#   label: ":books: Docs",
#   key: "docs",
#   depends_on: ["build","test"],
#   plugins: plugins,
#   commands: [
#     "mise trust",
#     "npm run docs"
#   ]
# )

# pipeline.add_step(
#   label: ":lab_coat: Apps",
#   key: "apps",
#   depends_on: ["build","test"],
#   plugins: plugins,
#   commands: [
#     "mise trust",
#     "npm run apps"
#   ]
# )

puts pipeline.to_json
