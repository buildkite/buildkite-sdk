package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

var stepBuilderCode = `import * as fs from "fs";
%s

class StepBuilder {
	private steps: any[] = [];

	public write() {
		fs.writeFileSync("pipeline.json", JSON.stringify({ steps: this.steps }, null, 4));
	}`

func newStepBuilderFile(pipelineSchema schema.PipelineSchema) string {
	file := NewFile()
	file.code = append(file.code, stepBuilderCode)
	interfaces := utils.CodeBlock{}

	for _, step := range pipelineSchema.Steps {
		file.imports.AddImport("./types", "types")

		stepName := utils.String.Capitalize(step.Name)
		file.code = append(file.code, utils.CodeBlock{
			utils.CodeGen.Comment.TypeScript(step.Description, 0),
			fmt.Sprintf("public add%sStep(args: types.%s): this {", stepName, stepName),
			"    this.steps.push({ ...args });",
			"    return this;",
			"}",
		}.DisplayIndent(4))
	}

	file.code[0] = fmt.Sprintf(file.code[0], interfaces.Display())

	return fmt.Sprintf("%s\n}\nexport default StepBuilder;", file.String())
}
