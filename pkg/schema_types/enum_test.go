package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaEnum(t *testing.T) {
	testValue := SchemaEnum{
		Name:   "enum",
		Values: []string{"one", "two"},
	}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "export enum Enum {\n    ONE = \"one\",\n    TWO = \"two\",\n}\n",
		GoType:         "type Enum string\nconst (\n    ONE Enum = \"one\"\n    TWO Enum = \"two\"\n)",
	})
}

func TestEnum(t *testing.T) {
	t.Run("should create a string enum", func(t *testing.T) {
		result := Enum.String("enum", []string{"one", "two"})
		assert.Equal(t, 2, len(result.Values))
	})
}
