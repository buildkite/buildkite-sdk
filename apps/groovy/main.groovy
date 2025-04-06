import buildkite.Pipeline
import buildkite.StepTypes

def pipeline = new Pipeline()

pipeline.addStep([
    label: "some-label",
    command: "echo 'Hello, world!'",
])

def outputDir = new File("../../out/apps/groovy")
outputDir.mkdirs()

new File(outputDir, "pipeline.json").text = pipeline.toJSON()
new File(outputDir, "pipeline.yaml").text = pipeline.toYAML()
