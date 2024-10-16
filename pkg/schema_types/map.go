package schema_types

import "fmt"

type SchemaMap struct {
	Items SchemaType
}

func (SchemaMap) IsUnion() bool {
	return false
}

func (s SchemaMap) TypeScriptType() string {
	return fmt.Sprintf("Record<string, %s>", s.Items.TypeScriptType())
}

func (s SchemaMap) GoType() string {
	return fmt.Sprintf("map[string]%s", s.Items.GoType())
}

type smap struct{}

func (smap) String() SchemaMap {
	return SchemaMap{
		Items: SchemaString{},
	}
}

func (smap) Number() SchemaMap {
	return SchemaMap{
		Items: SchemaNumber{},
	}
}

func (smap) Any() SchemaMap {
	return SchemaMap{
		Items: SchemaAny{},
	}
}

var Map = smap{}
