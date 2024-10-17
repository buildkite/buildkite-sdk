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
	definitions := utils.CodeGen.NewCodeBlock()
	definitionFields := make(map[string]bool)

	for _, val := range s.Values {
		var transformAssignments []string
		switch val.typ.(type) {
		case SchemaObject:
			fields := val.typ.(SchemaObject).Fields()
			for _, field := range fields {
				if _, ok := definitionFields[field.name.TitleCase()]; !ok {
					var propType string
					if field.fieldref != nil {
						switch field.fieldref.typ.(type) {
						case SchemaEnum:
						case SchemaArray:
						case SchemaObject:
							propType = field.fieldref.name.TitleCase()
						default:
							propType = field.fieldref.typ.GoType()
						}
					} else {
						propType = field.typ.GoType()
					}

					definitionFields[field.name.TitleCase()] = true
					definitions = append(definitions, fmt.Sprintf("%s\n    %s %s `json:\"%s,omitempty\"`",
						utils.CodeGen.Comment.Go(field.description, 4),
						field.name.TitleCase(),
						propType,
						string(field.name),
					))

					transformAssignments = append(transformAssignments, fmt.Sprintf("        %s: x.%s,",
						field.name.TitleCase(),
						field.name.TitleCase(),
					))
				}
			}
		default:
			panic("non object unions are currently not supported")
		}

		transformFuncs = append(transformFuncs, fmt.Sprintf("func (x %s) To%s() %sDefinition {\n    return %sDefinition{\n%s\n    }\n}",
			val.name.TitleCase(),
			s.Name.TitleCase(),
			s.Name.CamelCase(),
			s.Name.CamelCase(),
			strings.Join(transformAssignments, "\n"),
		))
	}

	block = append(block, fmt.Sprintf("type %sDefinition struct {\n%s\n}",
		s.Name.CamelCase(),
		strings.Join(definitions, "\n"),
	))
	block = append(block, fmt.Sprintf("type %s interface {\n    To%s() %sDefinition\n}",
		s.Name.TitleCase(),
		s.Name.TitleCase(),
		s.Name.CamelCase(),
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
