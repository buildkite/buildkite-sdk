package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaObject(t *testing.T) {
	testValue := SchemaObject{
		Name: "test_object",
		Properties: []Field{
			{
				name: "name",
				typ:  SchemaString{},
			},
		},
	}

	testSchemaType(t, testValue, schemaTypeTestExpectations{
		TypeScriptType: "export interface TestObject {\n    \n    name?: string;\n}",
		GoType:         "type TestObject struct {\n\n    Name string `json:\"name,omitempty\"`\n\n}",
	})
}

func TestObject(t *testing.T) {
	t.Run("should create an object", func(t *testing.T) {
		result := Object.New("test_object", []Field{
			{
				name: "name",
				typ:  SchemaString{},
			},
		})
		assert.Equal(t, 1, len(result.Properties))
	})
}
