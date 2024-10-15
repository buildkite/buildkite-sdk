package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaUnion(t *testing.T) {
	testValue := SchemaUnion{
		Name: "union",
		Values: []Field{
			{
				name: "field_one",
				typ: SchemaObject{
					Name:       "field_one",
					Properties: []Field{},
				},
			},
			{
				name: "field_two",
				typ: SchemaObject{
					Name:       "field_two",
					Properties: []Field{},
				},
			},
		},
	}

	t.Run("should return true for IsUnion", func(t *testing.T) {
		assert.Equal(t, true, testValue.IsUnion())
	})

	t.Run("should return the typescript type", func(t *testing.T) {
		assert.Equal(t, "(FieldOne | FieldTwo)", testValue.TypeScriptType())
	})

	t.Run("should return the go type", func(t *testing.T) {
		assert.Equal(t, "type unionDefinition struct {\n\n}\ntype Union interface {\n    ToUnion() unionDefinition\n}\nfunc (x FieldOne) ToUnion() unionDefinition {\n    return unionDefinition{\n\n    }\n}\nfunc (x FieldTwo) ToUnion() unionDefinition {\n    return unionDefinition{\n\n    }\n}", testValue.GoType())
	})
}

func TestUnion(t *testing.T) {
	t.Run("should create a union", func(t *testing.T) {
		result := Union.New("union", []Field{
			{
				name: "name",
				typ:  SchemaString{},
			},
			{
				name: "age",
				typ:  SchemaNumber{},
			},
		})
		assert.Equal(t, 2, len(result.Values))
	})
}
