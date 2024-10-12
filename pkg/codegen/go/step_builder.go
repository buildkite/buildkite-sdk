package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

func newStepBuilderMethod(name schema_types.AttributeName) string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n",
		fmt.Sprintf("func (s *stepBuilder) Add%s(step *%s) *stepBuilder {", name.TitleCase(), name.TitleCase()),
		"    s.Steps = append(s.Steps, step)",
		"    return s",
		"}",
	)
}

func newStepBuilderFile(pipelineSchema schema.PipelineSchema) string {
	file := NewFile()
	file.imports.AddImports("encoding/json", "os")
	file.code = append(file.code, fmt.Sprintf(`type stepBuilder struct {
	Steps []interface{} %s
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{}
}`, "`json:\"steps\"`"))

	for _, step := range pipelineSchema.Steps {
		file.code = append(file.code, newStepBuilderMethod(schema_types.AttributeName(step.Name)))
	}

	file.code = append(file.code, fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"func (s *stepBuilder) Print() error {",
		`	jsonBytes, err := json.MarshalIndent(s, "", "\t")`,
		"	if err != nil {",
		"	    return err",
		"	}",
		"",
		"    return os.WriteFile(\"pipeline.json\", jsonBytes, os.ModePerm)",
		"}",
	))

	return file.String()
}
