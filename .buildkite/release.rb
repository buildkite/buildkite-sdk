require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new

plugins = [
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
    "role-arn": "arn:aws:iam::597088016345:role/marketing-website-production-pipeline-role"
  }},
  { "aws-ssm#v1.0.0": {
    parameters: {
      NPM_TOKEN: "/prod/buildkite-sdk/npm-token",
      PYPI_TOKEN: "/prod/buildkite-sdk/pypi-token",
      GITHUB_TOKEN: "/prod/buildkite-sdk/github-token"
    }
  }}
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
        key: "#{target[:key]}-publish",
        label: ":rocket: Publish",
        depends_on: ["#{target[:key]}-test","#{target[:key]}-build"],
        plugins: language_plugins,
        commands: [
          "mise trust",
          "nx install #{target[:sdk_label]}",
          "nx run #{target[:sdk_label]}:publish"
        ],
      },
    ]
  )
end

puts pipeline.to_json
