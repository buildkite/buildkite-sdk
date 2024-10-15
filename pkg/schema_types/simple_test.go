package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaString(t *testing.T) {
	testValue := SchemaString{}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "string",
		GoType:         "string",
	})
}

func TestString(t *testing.T) {
	t.Run("should create a string", func(t *testing.T) {
		result := Simple.String()
		assert.Equal(t, "string", result.GoType())
	})
}

func TestSchemaNumber(t *testing.T) {
	testValue := SchemaNumber{}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "number",
		GoType:         "int",
	})
}

func TestNumber(t *testing.T) {
	t.Run("should create a number", func(t *testing.T) {
		result := Simple.Number()
		assert.Equal(t, "int", result.GoType())
	})
}

func TestSchemaBoolean(t *testing.T) {
	testValue := SchemaBoolean{}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "boolean",
		GoType:         "bool",
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should create a boolean", func(t *testing.T) {
		result := Simple.Boolean()
		assert.Equal(t, "bool", result.GoType())
	})
}

func TestSchemaAny(t *testing.T) {
	testValue := SchemaAny{}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "any",
		GoType:         "interface{}",
	})
}

func TestAny(t *testing.T) {
	t.Run("should create a boolean", func(t *testing.T) {
		result := Simple.Any()
		assert.Equal(t, "interface{}", result.GoType())
	})
}
