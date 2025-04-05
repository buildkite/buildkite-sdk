import buildkite.Pipeline
import buildkite.StepTypes

def pipeline = new Pipeline()

pipeline.addStep([
    label: "some-label",
    command: "echo 'Hello, world!'"
])

print pipeline.toJSON()
