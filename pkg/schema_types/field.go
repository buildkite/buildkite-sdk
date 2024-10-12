package schema_types

type fieldDefinition struct {
	Name        AttributeName
	Description string
	Required    bool
	Typ         SchemaType
	Fieldref    Field
}

type Field struct {
	name        AttributeName
	description string
	required    bool
	typ         SchemaType
	fieldref    *Field
}

func (f Field) GetDefinition() fieldDefinition {
	def := fieldDefinition{
		Name:        f.name,
		Description: f.description,
		Required:    f.required,
		Typ:         f.typ,
	}

	if f.fieldref != nil {
		def.Fieldref = *f.fieldref
	}

	return def
}

func (f *Field) Name(name string) *Field {
	f.name = AttributeName(name)
	return f
}

func (f *Field) Description(desc string) *Field {
	f.description = desc
	return f
}

func (f *Field) Required() *Field {
	f.required = true
	return f
}

func (f *Field) FieldRef(ref *Field) Field {
	f.fieldref = ref
	return *f
}

func (f *Field) String() Field {
	f.typ = Simple.String()
	return *f
}

func (f *Field) StringArray() Field {
	f.typ = Array.String()
	return *f
}

func (f *Field) StringMap() Field {
	f.typ = Map.String()
	return *f
}

func (f *Field) Number() Field {
	f.typ = Simple.Number()
	return *f
}

func (f *Field) NumberMap() Field {
	f.typ = Map.Number()
	return *f
}

func (f *Field) Boolean() Field {
	f.typ = Simple.Boolean()
	return *f
}

func (f *Field) AnyMapArray() Field {
	f.typ = Array.AnyMap()
	return *f
}

func (f *Field) Object(props []Field) Field {
	f.typ = Object.New(string(f.name), props)
	return *f
}

func (f *Field) Union(name string, fields ...Field) Field {
	f.typ = Union.New(name, fields)
	return *f
}

func (f *Field) UnionArray(name string, fields ...Field) Field {
	f.typ = Array.Union(name, fields)
	return *f
}

func (f *Field) Enum(values ...string) Field {
	f.typ = Enum.String(string(f.name), values)
	return *f
}

func (f Field) TypeScriptIdentifier() string {
	return f.name.TitleCase()
}

func NewField() *Field {
	return &Field{}
}
