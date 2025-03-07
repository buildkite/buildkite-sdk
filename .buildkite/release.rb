require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new

commands = [
  "mise trust",
  "npm install",
  "npm test",
  "npm run build",
  "npm run docs",
  "npm run apps",
  "npm run publish"
]

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } },
  { "rubygems-oidc#v0.2.0": { role: "rg_oidc_akr_emf87k6zphtb7x7adyrk" } },
  { "aws-assume-role-with-web-identity#v1.0.0": {
    "role-arn": "arn:aws:iam::597088016345:role/marketing-website-production-pipeline-role"
  }},
  { "aws-ssm#v1.0.0": {
    parameters: {
      NPM_TOKEN: "prod/buildkite-sdk/npm-token",
      PYPI_TOKEN: "prod/buildkite-sdk/pypi-token",
      GITHUB_TOKEN: "prod/buildkite-sdk/github-token"
    }
  }}
]

pipeline.add_step(
  label: ":hammer_and_wrench: Install, test, build, publish",
  plugins: plugins,
  commands: commands
)

puts pipeline.to_json
