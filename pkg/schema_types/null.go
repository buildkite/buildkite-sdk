package schema_types

import "fmt"

type SchemaNull struct {
	Name AttributeName
}

func (SchemaNull) IsUnion() bool {
	return false
}

func (s SchemaNull) TypeScriptType() string {
	return ""
}

func (s SchemaNull) GoType() string {
	return fmt.Sprintf("type %s string\nfunc (%s) MarshalJSON() ([]byte,error) {\n return json.Marshal(nil) \n}",
		s.Name.TitleCase(),
		s.Name.TitleCase(),
	)
}

type null struct{}

func (null) New(name string) SchemaNull {
	return SchemaNull{
		Name: AttributeName(name),
	}
}

var Null = null{}
