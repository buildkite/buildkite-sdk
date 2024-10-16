package go_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/stretchr/testify/assert"
)

func TestNewStepBuilderMethod(t *testing.T) {
	t.Run("should render a step builder method", func(t *testing.T) {
		stepName := schema_types.AttributeName("test")
		result := newStepBuilderMethod(stepName)
		assert.Equal(t, "func (s *stepBuilder) AddTest(step *Test) *stepBuilder {\n    s.Steps = append(s.Steps, step)\n    return s\n}", result)
	})
}

func TestNewStepBuilderFile(t *testing.T) {
	t.Run("should create a step builder file", func(t *testing.T) {
		step := schema.Step{
			Name:        "test",
			Description: "description",
			Fields:      []schema_types.Field{},
		}

		result := newStepBuilderFile([]schema.Step{step})
		assert.Greater(t, len(result), 0)
	})
}
