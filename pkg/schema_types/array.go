package schema_types

import "fmt"

type SchemaArray struct {
	Items SchemaType
}

func (SchemaArray) IsUnion() bool {
	return false
}

func (s SchemaArray) TypeScriptType() string {
	return fmt.Sprintf("%s[]", s.Items.TypeScriptType())
}

func (s SchemaArray) GoType() string {
	if s.Items.IsUnion() {
		return s.Items.GoType()
	}

	return fmt.Sprintf("[]%s", s.Items.GoType())
}
