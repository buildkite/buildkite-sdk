package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	t.Run("should create a new field", func(t *testing.T) {
		field := NewField()
		assert.Equal(t, AttributeName(""), field.name)
	})

	t.Run("should set the name of a field", func(t *testing.T) {
		field := NewField().Name("test_field")
		assert.Equal(t, AttributeName("test_field"), field.name)
	})

	t.Run("should set the description of a field", func(t *testing.T) {
		field := NewField().Description("test description")
		assert.Equal(t, "test description", field.description)
	})

	t.Run("should mark a field as required", func(t *testing.T) {
		field := NewField().Required()
		assert.Equal(t, true, field.required)
	})

	t.Run("should create a fieldref field", func(t *testing.T) {
		fieldRef := NewField().Name("test").Description("test desc").Object([]Field{
			{
				name:        "some_field",
				description: "some field",
				typ:         SchemaString{},
			},
		})

		field := NewField().FieldRef(&fieldRef)
		assert.Equal(t, &fieldRef, field.fieldref)
	})

	t.Run("should create a string field", func(t *testing.T) {
		field := NewField().String()
		assert.Equal(t, SchemaString{}, field.typ)
	})

	t.Run("should create a string array field", func(t *testing.T) {
		field := NewField().StringArray()
		assert.Equal(t, SchemaArray{Items: SchemaString{}}, field.typ)
	})

	t.Run("should create a string map field", func(t *testing.T) {
		field := NewField().StringMap()
		assert.Equal(t, SchemaMap{Items: SchemaString{}}, field.typ)
	})

	t.Run("should create a number field", func(t *testing.T) {
		field := NewField().Number()
		assert.Equal(t, SchemaNumber{}, field.typ)
	})

	t.Run("should create a number map field", func(t *testing.T) {
		field := NewField().NumberMap()
		assert.Equal(t, SchemaMap{Items: SchemaNumber{}}, field.typ)
	})

	t.Run("should create a boolean field", func(t *testing.T) {
		field := NewField().Boolean()
		assert.Equal(t, SchemaBoolean{}, field.typ)
	})

	t.Run("should create an any map array field", func(t *testing.T) {
		field := NewField().AnyMapArray()
		assert.Equal(t, SchemaArray{Items: Map.Any()}, field.typ)
	})

	t.Run("should create an object field", func(t *testing.T) {
		field := NewField().Object([]Field{})
		assert.Equal(t, SchemaObject{Name: "", Properties: []Field{}}, field.typ)
	})

	t.Run("should create a union field", func(t *testing.T) {
		field := NewField().Union("test_field", Field{name: "one"}, Field{name: "two"})
		assert.Equal(t, SchemaUnion{Name: "test_field", Values: []Field{{name: "one"}, {name: "two"}}}, field.typ)
	})

	t.Run("should create a union array field", func(t *testing.T) {
		field := NewField().UnionArray("test_field", Field{name: "one"}, Field{name: "two"})
		expectation := SchemaArray{
			Items: SchemaUnion{
				Name:   "test_field",
				Values: []Field{{name: "one"}, {name: "two"}},
			},
		}
		assert.Equal(t, expectation, field.typ)
	})

	t.Run("should create an enum field", func(t *testing.T) {
		field := NewField().Enum("one", "two")
		assert.Equal(t, SchemaEnum{Values: []string{"one", "two"}}, field.typ)
	})

	t.Run("should emit the typescript identifier", func(t *testing.T) {
		field := NewField().Name("test_field").String()
		assert.Equal(t, "TestField", field.TypeScriptIdentifier())
	})

	t.Run("should get the definition of a field", func(t *testing.T) {
		field := NewField().
			Name("test_field").
			Description("test description").
			Required().
			String()

		def := field.GetDefinition()
		assert.Equal(t, AttributeName("test_field"), def.Name)
		assert.Equal(t, "test description", def.Description)
		assert.Equal(t, true, def.Required)
		assert.Equal(t, SchemaString{}, def.Typ)
	})
}
