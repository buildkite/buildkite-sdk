package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

func NewTypesFile(fields []schema_types.Field, steps []schema.Step) (string, error) {
	file := NewFile()

	for _, field := range fields {
		def := field.GetDefinition()
		switch def.Typ.(type) {
		case schema_types.SchemaUnion:
			file.code = append(file.code,
				fmt.Sprintf("type %s = %s\n", def.Name.TitleCase(), def.Typ.TypeScriptType()),
			)
		default:
			file.code = append(file.code,
				fmt.Sprintf("%s\n", def.Typ.TypeScriptType()),
			)
		}
	}

	for _, step := range steps {
		def := step.ToObjectField().GetDefinition()
		switch def.Typ.(type) {
		case schema_types.SchemaUnion:
			file.code = append(file.code,
				fmt.Sprintf("type %s = %s\n", def.Name.TitleCase(), def.Typ.TypeScriptType()),
			)
		default:
			file.code = append(file.code,
				fmt.Sprintf("%s\n", def.Typ.TypeScriptType()),
			)
		}
	}

	return file.String(), nil
}
