require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

def run_language_pipeline(target)
  language_pipeline = Buildkite::Pipeline.new
  language_pipeline_plugins = [
    { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } },
    { "artifacts#v1.9.2": {
      download: ["node_modules"],
      compressed: "node_modules.tgz"
    }},
  ]

  language_pipeline.add_step(
    key: "#{target[:key]}",
    group: "#{target[:icon]} #{target[:label]}",
    steps: [
      {
        key: "#{target[:key]}-test",
        label: ":test_tube: Test",
        plugins: language_pipeline_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:sdk_label]}",
          "nx test #{target[:sdk_label]}"
        ],
      },
      {
        key: "#{target[:key]}-build",
        label: ":package: Build",
        plugins: language_pipeline_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:sdk_label]}",
          "nx build #{target[:sdk_label]}"
        ],
      },
      {
        key: "#{target[:key]}-docs",
        label: ":books: Docs",
        depends_on: ["#{target[:key]}-test","#{target[:key]}-build"],
        plugins: language_pipeline_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:sdk_label]}",
          "nx run #{target[:sdk_label]}:docs:build"
        ],
      },
      {
        label: ":lab_coat: Apps",
        key: "#{target[:key]}-apps",
        depends_on: ["#{target[:key]}-test","#{target[:key]}-build"],
        plugins: language_pipeline_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:app_label]}",
          "nx run #{target[:app_label]}:run"
        ],
      },
    ]
  )

  puts language_pipeline.to_json
end
