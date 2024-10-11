package schema_types

import "fmt"

type SchemaMap struct {
	Items SchemaType
}

func (s SchemaMap) TypeScriptType() string {
	return fmt.Sprintf("Record<string, %s>", s.Items.TypeScriptType())
}

func (s SchemaMap) GoType() string {
	return fmt.Sprintf("map[string]%s", s.Items.GoType())
}
