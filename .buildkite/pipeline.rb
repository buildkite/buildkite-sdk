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
      upload: ["node_modules"],
      compressed: "node_modules.tgz"
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
    download: ["node_modules"],
    compressed: "node_modules.tgz"
  }}
]

language_targets = [
  {
    icon: ":typescript:",
    label: "Typescript",
    key: "typescript",
    sdk_label: "sdk-typescript",
    app_label: "app-typescript"
  },
  {
    icon: ":python:",
    label: "Python",
    key: "python",
    sdk_label: "sdk-python",
    app_label: "app-python"
  },
  {
    icon: ":go:",
    label: "Go",
    key: "go",
    sdk_label: "sdk-go",
    app_label: "app-go"
  },
  {
    icon: ":ruby:",
    label: "Ruby",
    key: "ruby",
    sdk_label: "sdk-ruby",
    app_label: "app-ruby"
  },
  {
    icon: ":csharp:",
    label: "C#",
    key: "csharp",
    sdk_label: "sdk-csharp",
    app_label: "app-csharp"
  }
]

language_targets.each do |target|
  pipeline.add_step(
    depends_on: "install",
    key: "#{target[:key]}",
    group: "#{target[:icon]} #{target[:label]}",
    steps: [
      {
        key: "#{target[:key]}-test",
        label: ":test_tube: Test",
        plugins: language_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:sdk_label]}",
          "nx test #{target[:sdk_label]}"
        ],
      },
      {
        key: "#{target[:key]}-build",
        label: ":package: Build",
        plugins: language_plugins,
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
        plugins: language_plugins,
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
        plugins: language_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:app_label]}",
          "nx run #{target[:app_label]}:run"
        ],
      },
    ]
  )
end

puts pipeline.to_json
