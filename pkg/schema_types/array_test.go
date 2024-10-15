package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaArray(t *testing.T) {
	testValue := SchemaArray{
		Items: SchemaString{},
	}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "string[]",
		GoType:         "[]string",
	})

	t.Run("should return a struct and interface for unions", func(t *testing.T) {
		expectation := "type testFieldDefinition struct {\n\n}\ntype TestField interface {\n    ToTestField() testFieldDefinition\n}"
		testUnionValue := SchemaArray{
			Items: Union.New("test_field", []Field{}),
		}
		assert.Equal(t, expectation, testUnionValue.GoType())
	})
}

func TestArray(t *testing.T) {
	t.Run("should create a string array", func(t *testing.T) {
		result := Array.String()
		assert.Equal(t, "[]string", result.GoType())
	})

	t.Run("should create a number array", func(t *testing.T) {
		result := Array.Number()
		assert.Equal(t, "[]int", result.GoType())
	})

	t.Run("should create a string map array", func(t *testing.T) {
		result := Array.StringMap()
		assert.Equal(t, "[]map[string]string", result.GoType())
	})

	t.Run("should create a number map array", func(t *testing.T) {
		result := Array.NumberMap()
		assert.Equal(t, "[]map[string]int", result.GoType())
	})

	t.Run("should create an any map array", func(t *testing.T) {
		result := Array.AnyMap()
		assert.Equal(t, "[]map[string]interface{}", result.GoType())
	})

	t.Run("should create a union array", func(t *testing.T) {
		result := Array.Union("test_union", []Field{
			{
				name: "one",
			},
			{
				name: "two",
			},
		})
		assert.Equal(t, "(One | Two)[]", result.TypeScriptType())
	})
}
