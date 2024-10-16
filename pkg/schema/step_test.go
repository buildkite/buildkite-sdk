package schema

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	t.Run("should transform a step into a field", func(t *testing.T) {
		step := Step{
			Name:        "test-step",
			Description: "a test step",
			Fields: []schema_types.Field{
				schema_types.NewField().String(),
			},
		}

		field := step.ToObjectField().GetDefinition()
		assert.Equal(t, step.Name, string(field.Name))
		assert.Equal(t, step.Description, field.Description)

		_, isObject := field.Typ.(schema_types.SchemaObject)
		assert.Equal(t, true, isObject)
	})
}
