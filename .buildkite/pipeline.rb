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

# If the job has an associated tag that looks like a new version, add a publish step.
commands.push("npm run publish") if !tag.nil? && tag.start_with?("v")

pipeline.add_step(
  label: ":hammer_and_wrench: Install, test, build, publish",
  plugins: [{ "docker#v5.11.0": { image: "buildkite-sdk-tools:latest" } }],
  commands: commands
)

puts pipeline.to_json
