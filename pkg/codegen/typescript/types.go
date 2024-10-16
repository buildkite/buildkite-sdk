package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

func renderType(field schema_types.Field) string {
	def := field.GetDefinition()
	switch def.Typ.(type) {
	case schema_types.SchemaObject, schema_types.SchemaEnum:
		return fmt.Sprintf("%s\n", def.Typ.TypeScriptType())
	default:
		return fmt.Sprintf("type %s = %s\n", def.Name.TitleCase(), def.Typ.TypeScriptType())
	}
}

func newTypesFile(fields []schema_types.Field, steps []schema.Step) string {
	file := newFile()

	for _, field := range fields {
		file.AppendCode(renderType(field))
	}

	for _, step := range steps {
		file.AppendCode(renderType(step.ToObjectField()))
	}

	return file.Render()
}
