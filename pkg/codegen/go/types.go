package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

func newTypesFile(fields []schema_types.Field, steps []schema.Step) string {
	file := newFile()
	file.AddImport("encoding/json", "encoding/json")

	for _, field := range fields {
		def := field.GetDefinition()
		file.AppendCode(
			fmt.Sprintf("%s\n", def.Typ.GoType()),
		)
	}

	for _, step := range steps {
		def := step.ToObjectField().GetDefinition()
		file.AppendCode(
			fmt.Sprintf("%s\n", def.Typ.GoType()),
		)
	}

	return file.Render()
}
