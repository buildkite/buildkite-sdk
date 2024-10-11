package schema_types

// String
type SchemaString struct{}

func (SchemaString) TypeScriptType() string {
	return "string"
}

func (SchemaString) GoType() string {
	return "string"
}

// Number
type SchemaNumber struct{}

func (SchemaNumber) TypeScriptType() string {
	return "number"
}

func (SchemaNumber) GoType() string {
	return "int"
}

// Boolean
type SchemaBoolean struct{}

func (SchemaBoolean) TypeScriptType() string {
	return "boolean"
}

func (SchemaBoolean) GoType() string {
	return "bool"
}

// Any
type SchemaAny struct{}

func (SchemaAny) TypeScriptType() string {
	return "any"
}

func (SchemaAny) GoType() string {
	return "interface{}"
}
