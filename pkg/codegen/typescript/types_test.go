package typescript_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/stretchr/testify/assert"
)

func TestRenderType(t *testing.T) {
	t.Run("should render a union", func(t *testing.T) {
		obj1 := schema_types.NewField().Name("obj1").Object([]schema_types.Field{
			schema_types.NewField().Name("name").String(),
		})

		obj2 := schema_types.NewField().Name("obj2").Object([]schema_types.Field{
			schema_types.NewField().Name("name").String(),
		})

		union := schema_types.NewField().Name("union").Union("union", obj1, obj2)
		result := renderType(union)
		assert.Equal(t, "type Union = (Obj1 | Obj2)\n", result)
	})

	t.Run("should render a type", func(t *testing.T) {
		field := schema_types.NewField().Name("field").String()
		result := renderType(field)
		assert.Equal(t, "string\n", result)
	})
}

func TestNewTypesFile(t *testing.T) {
	t.Run("shoudl render a types file", func(t *testing.T) {
		file := newTypesFile(schema.PipelinesSchema.Types, schema.PipelinesSchema.Steps)
		assert.Greater(t, len(file), 0)
	})
}
