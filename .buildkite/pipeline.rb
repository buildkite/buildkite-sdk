require_relative("../sdk/ruby/lib/buildkite")
require_relative("../sdk/ruby/lib/environment")

pipeline = Buildkite::Pipeline.new
tag = ENV[Environment::BUILDKITE_TAG]

commands = [
  "mise trust",
  "npm install",
  "npm test",
  "npm run build",
  "npm run docs",
  "npm run apps"
]

plugins = [
  { "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } }
]

# If the job has an associated tag that looks like a new version, add a publish step.
if !tag.nil? && tag.start_with?("v")
  commands.push("npm run publish")
  plugins.push({ "rubygems-oidc#v0.2.0": { role: "rg_oidc_akr_emf87k6zphtb7x7adyrk" } })
  plugins.push({ "aws-assume-role-with-web-identity#v1.0.0": {
    "role-arn": "arn:aws:iam::597088016345:role/marketing-website-production-pipeline-role"
  }})
  plugins.push({ "aws-ssm#v1.0.0": {
    parameters: {
      NPM_TOKEN: "prod/buildkite-sdk/npm-token",
      PYPI_TOKEN: "prod/buildkite-sdk/pypi-token",
      GITHUB_TOKEN: "prod/buildkite-sdk/github-token"
    }
  }})
end

pipeline.add_step(
  label: ":hammer_and_wrench: Install, test, build, publish",
  plugins: plugins,
  commands: commands
)

puts pipeline.to_json
