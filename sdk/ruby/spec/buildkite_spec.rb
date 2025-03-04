RSpec.describe Buildkite do
  it "has a version number" do
    expect(Buildkite::VERSION).not_to be_nil
  end

  it "does something useful" do
    pipeline = Buildkite::Pipeline.new

    pipeline.add_step(label: "some-label", command: "echo 'Hello, World!'")

    json = { steps: [{ label: "some-label", "command": "echo 'Hello, World!'" }] }
    expected = JSON.pretty_generate(json, indent: "    ")
    expect(pipeline.to_json).to eq(expected)
  end
end
