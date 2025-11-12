RSpec.describe Buildkite do
  it "has a version number" do
    expect(Buildkite::VERSION).not_to be_nil
  end

  it "builds a simple pipeline" do
    pipeline = Buildkite::Pipeline.new

    pipeline.add_step(label: "some-label", command: "echo 'Hello, World!'")

    json = { steps: [{ label: "some-label", "command": "echo 'Hello, World!'" }] }
    expected = JSON.pretty_generate(json, indent: "    ")
    expect(pipeline.to_json).to eq(expected)
  end

  it "builds a pipeline" do
    pipeline = Buildkite::Pipeline.new

    pipeline.add_agent("queue", "hosted")
    pipeline.add_environment_variable("FOO", "bar")
    pipeline.add_step(label: "some-label", command: "echo 'Hello, World!'")
    pipeline.set_secrets(["MY_SECRET"])

    json = {
      steps: [{ label: "some-label", "command": "echo 'Hello, World!'" }],
      agents: { "queue": "hosted" },
      env: { "FOO": "bar" },
      secrets: ["MY_SECRET"]
    }
    expected = JSON.pretty_generate(json, indent: "    ")
    expect(pipeline.to_json).to eq(expected)
  end
end
