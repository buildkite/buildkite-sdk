package schema_types

import (
	"fmt"
	"strings"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type SchemaUnion struct {
	Name   AttributeName
	Values []Field
}

func (SchemaUnion) IsUnion() bool {
	return true
}

func (s SchemaUnion) TypeScriptType() string {
	unionValues := make([]string, len(s.Values))
	for i, val := range s.Values {
		unionValues[i] = val.TypeScriptIdentifier()
	}

	return fmt.Sprintf("(%s)", strings.Join(unionValues, " | "))
}

func (s SchemaUnion) GoType() string {
	block := utils.CodeBlock{}
	transformFuncs := utils.CodeGen.NewCodeBlock()

	for _, val := range s.Values {
		transformFuncs = append(transformFuncs, fmt.Sprintf("func (x *%s) MarshalJSON() ([]byte, error) {\n    return json.Marshal(*x)\n}",
			val.name.TitleCase(),
		))
	}

	block = append(block, fmt.Sprintf("type %s interface {\n    MarshalJSON() ([]byte, error)\n}",
		s.Name.TitleCase(),
	))
	block = append(block, transformFuncs...)

	return block.Display()
}

// Union
type union struct{}

func (union) New(name string, fields []Field) SchemaUnion {
	return SchemaUnion{
		Name:   AttributeName(name),
		Values: fields,
	}
}

var Union = union{}
