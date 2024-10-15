package schema_types

type SchemaType interface {
	TypeScriptType() string
	GoType() string
	IsUnion() bool
}
