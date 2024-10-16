package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

var printFn = `func (s *stepBuilder) Print() error {
	jsonBytes, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("pipeline.json", jsonBytes, os.ModePerm)
}`

var stepBuilderStruct = fmt.Sprintf(`type stepBuilder struct {
	Steps []interface{} %s
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{}
}`, "`json:\"steps\"`")

func newStepBuilderMethod(name schema_types.AttributeName) string {
	return utils.CodeGen.NewCodeBlock(
		fmt.Sprintf("func (s *stepBuilder) Add%s(step *%s) *stepBuilder {", name.TitleCase(), name.TitleCase()),
		"    s.Steps = append(s.Steps, step)",
		"    return s",
		"}",
	).Display()
}

func newStepBuilderFile(steps []schema.Step) string {
	file := newFile()
	file.AddImport("encoding/json", "encoding/json")
	file.AddImport("os", "os")
	file.AppendCode(stepBuilderStruct)

	for _, step := range steps {
		file.AppendCode(newStepBuilderMethod(schema_types.AttributeName(step.Name)))
	}

	file.AppendCode(printFn)

	return file.Render()
}
