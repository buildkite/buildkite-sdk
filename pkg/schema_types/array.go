package schema_types

import "fmt"

type SchemaArray struct {
	Items SchemaType
}

func (s SchemaArray) TypeScriptType() string {
	return fmt.Sprintf("%s[]", s.Items.TypeScriptType())
}

func (s SchemaArray) GoType() string {
	return fmt.Sprintf("[]%s", s.Items.GoType())
}
