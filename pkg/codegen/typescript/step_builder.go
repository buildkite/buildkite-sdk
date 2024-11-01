package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

var stepBuilderCode = `
class StepBuilder {
	public steps: any[] = [];

	public write() {
		fs.writeFileSync("pipeline.json", JSON.stringify({ steps: this.steps }, null, 4));
	}`

func renderStepFunction(step schema.Step) string {
	stepName := utils.String.Capitalize(step.Name)
	return utils.CodeGen.NewCodeBlock(
		utils.CodeGen.Comment.TypeScript(step.Description, 0),
		fmt.Sprintf("public add%sStep(args: types.%s): this {", stepName, stepName),
		"    this.steps.push({ ...args });",
		"    return this;",
		"}",
	).DisplayIndent(4)
}

func newStepBuilderFile(pipelineSchema schema.PipelineSchema) string {
	file := newFile()

	file.AddImport("fs", "fs")
	file.AddImport("types", "./types")
	file.AppendCode(stepBuilderCode)

	for _, step := range pipelineSchema.Steps {
		file.AppendCode(renderStepFunction(step))
	}

	file.AppendCode("\n}\nexport default StepBuilder;")

	return file.Render()
}
