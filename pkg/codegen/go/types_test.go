package go_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestNewTypesFile(t *testing.T) {
	t.Run("should create a types file", func(t *testing.T) {
		result := newTypesFile(schema.PipelinesSchema.Types, schema.PipelinesSchema.Steps)
		assert.Greater(t, len(result), 0)
	})
}
