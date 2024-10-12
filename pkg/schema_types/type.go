package schema_types

type SchemaType interface {
	TypeScriptType() string
	GoType() string
	IsUnion() bool
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

var Simple = simple{}

// Array
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

func (array) Custom(typ SchemaType) SchemaArray {
	return SchemaArray{
		Items: SchemaMap{
			Items: typ,
		},
	}
}

var Array = array{}

// Map
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

func (smap) Custom(typ SchemaType) SchemaMap {
	return SchemaMap{
		Items: typ,
	}
}

var Map = smap{}

// Object
type object struct{}

func (object) New(name string, props []Field) SchemaObject {
	return SchemaObject{
		Name:       AttributeName(name),
		Properties: props,
	}
}

var Object = object{}

// Union
type union struct{}

func (union) New(name string, fields []Field) SchemaUnion {
	return SchemaUnion{
		Name:   AttributeName(name),
		Values: fields,
	}
}

var Union = union{}

// Enum
type enum struct{}

func (enum) String(name string, values []string) SchemaEnum {
	return SchemaEnum{
		Name:   name,
		Values: values,
	}
}

var Enum = enum{}
