package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

func newTypesFile(fields []schema_types.Field, steps []schema.Step) string {
	file := NewFile()

	for _, field := range fields {
		def := field.GetDefinition()
		switch def.Typ.(type) {
		default:
			file.code = append(file.code,
				fmt.Sprintf("%s\n", def.Typ.GoType()),
			)
		}
	}

	for _, step := range steps {
		def := step.ToObjectField().GetDefinition()
		switch def.Typ.(type) {
		default:
			file.code = append(file.code,
				fmt.Sprintf("%s\n", def.Typ.GoType()),
			)
		}
	}

	return file.String()
}
