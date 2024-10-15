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

type array struct{}

func (array) String() SchemaArray {
	return SchemaArray{
		Items: SchemaString{},
	}
}

func (array) Number() SchemaArray {
	return SchemaArray{
		Items: SchemaNumber{},
	}
}

func (array) StringMap() SchemaArray {
	return SchemaArray{
		Items: Map.String(),
	}
}

func (array) NumberMap() SchemaArray {
	return SchemaArray{
		Items: Map.Number(),
	}
}

func (array) AnyMap() SchemaArray {
	return SchemaArray{
		Items: Map.Any(),
	}
}

func (array) Union(name string, fields []Field) SchemaArray {
	return SchemaArray{
		Items: Union.New(name, fields),
	}
}

var Array = array{}
