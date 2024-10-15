package schema_types

// String
type SchemaString struct{}

func (SchemaString) IsUnion() bool {
	return false
}

func (SchemaString) TypeScriptType() string {
	return "string"
}

func (SchemaString) GoType() string {
	return "string"
}

// Number
type SchemaNumber struct{}

func (SchemaNumber) IsUnion() bool {
	return false
}

func (SchemaNumber) TypeScriptType() string {
	return "number"
}

func (SchemaNumber) GoType() string {
	return "int"
}

// Boolean
type SchemaBoolean struct{}

func (SchemaBoolean) IsUnion() bool {
	return false
}

func (SchemaBoolean) TypeScriptType() string {
	return "boolean"
}

func (SchemaBoolean) GoType() string {
	return "bool"
}

// Any
type SchemaAny struct{}

func (SchemaAny) IsUnion() bool {
	return false
}

func (SchemaAny) TypeScriptType() string {
	return "any"
}

func (SchemaAny) GoType() string {
	return "interface{}"
}

// Simple
type simple struct{}

func (simple) String() SchemaString {
	return SchemaString{}
}

func (simple) Number() SchemaNumber {
	return SchemaNumber{}
}

func (simple) Boolean() SchemaBoolean {
	return SchemaBoolean{}
}

func (simple) Any() SchemaAny {
	return SchemaAny{}
}

var Simple = simple{}
