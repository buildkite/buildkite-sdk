package schema_types

import (
	"fmt"
	"strings"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type SchemaObject struct {
	Name       AttributeName
	Properties []Field
}

func (s SchemaObject) TypeScriptType() string {
	tsInterface := utils.CodeBlock{
		fmt.Sprintf("export interface %s {", s.Name.TitleCase()),
	}

	var properties []string
	for _, p := range s.Properties {
		optionalMarker := ""
		if !p.required {
			optionalMarker = "?"
		}

		if p.fieldref != nil {
			properties = append(properties, utils.CodeBlock{
				utils.NewCodeComment(p.description, 0),
				fmt.Sprintf("%s%s: %s;", p.name.CamelCase(), optionalMarker, p.fieldref.name.TitleCase()),
			}.DisplayIndent(4))
			continue
		}

		propType := p.typ.TypeScriptType()

		properties = append(properties, utils.CodeBlock{
			utils.NewCodeComment(p.description, 0),
			fmt.Sprintf("%s%s: %s;", p.name.CamelCase(), optionalMarker, propType),
		}.DisplayIndent(4))
	}

	tsInterface = append(tsInterface, properties...)
	tsInterface = append(tsInterface, "}")

	return tsInterface.Display()
}

func (s SchemaObject) GoType() string {
	var properties []string
	for _, prop := range s.Properties {
		var propType string
		if prop.fieldref != nil {
			propType = prop.fieldref.typ.GoType()
		} else {
			propType = prop.typ.GoType()
		}

		properties = append(properties, fmt.Sprintf("%s\n%s\n",
			utils.NewCodeComment(prop.description, 4),
			fmt.Sprintf("    %s %s `json:\"%s,omitempty\"`", s.Name.TitleCase(), propType, prop.name),
		))
	}

	return utils.CodeBlock{
		fmt.Sprintf("type %s struct {", s.Name.TitleCase()),
		strings.Join(properties, "\n"),
		"}",
	}.Display()
}
