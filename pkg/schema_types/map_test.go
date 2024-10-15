package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaMap(t *testing.T) {
	testValue := SchemaMap{
		Items: SchemaString{},
	}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "Record<string, string>",
		GoType:         "map[string]string",
	})
}

func TestMap(t *testing.T) {
	t.Run("create a string map", func(t *testing.T) {
		result := Map.String()
		assert.Equal(t, "map[string]string", result.GoType())
	})

	t.Run("create a number map", func(t *testing.T) {
		result := Map.Number()
		assert.Equal(t, "map[string]int", result.GoType())
	})

	t.Run("create an any map", func(t *testing.T) {
		result := Map.Any()
		assert.Equal(t, "map[string]interface{}", result.GoType())
	})
}
