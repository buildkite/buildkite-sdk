package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type schemaTypeTestExpectations struct {
	TypeScriptType string
	GoType         string
}

func testSchemaType(t *testing.T, testValue SchemaType, expectations schemaTypeTestExpectations) {
	t.Run("should return false for IsUnion", func(t *testing.T) {
		assert.Equal(t, false, testValue.IsUnion())
	})

	t.Run("should return the typescript type", func(t *testing.T) {
		assert.Equal(t, expectations.TypeScriptType, testValue.TypeScriptType())
	})

	t.Run("should return the go type", func(t *testing.T) {
		assert.Equal(t, expectations.GoType, testValue.GoType())
	})
}
