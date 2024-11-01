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

func (SchemaObject) IsUnion() bool {
	return false
}

func (s SchemaObject) Fields() []Field {
	return s.Properties
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
				utils.CodeGen.Comment.TypeScript(p.description, 0),
				fmt.Sprintf("%s%s: %s;", string(p.name), optionalMarker, p.fieldref.name.TitleCase()),
			}.DisplayIndent(4))
			continue
		}

		propType := p.typ.TypeScriptType()

		properties = append(properties, utils.CodeBlock{
			utils.CodeGen.Comment.TypeScript(p.description, 0),
			fmt.Sprintf("%s%s: %s;", string(p.name), optionalMarker, propType),
		}.DisplayIndent(4))
	}

	tsInterface = append(tsInterface, properties...)
	tsInterface = append(tsInterface, "}")

	return tsInterface.Display()
}

func (s SchemaObject) GoType() string {
	var properties []string
	for _, prop := range s.Properties {
		if prop.fieldref != nil {
			arrayType := ""
			typeName := fmt.Sprintf("*%s", prop.fieldref.name.TitleCase())
			if typ, ok := prop.fieldref.typ.(SchemaArray); ok {
				arrayType = "[]"

				if typ.IsUnion() || typ.Items.IsUnion() {
					typeName = prop.fieldref.name.TitleCase()
				}
			}

			omitEmpty := ",omitempty"
			if _, ok := prop.fieldref.typ.(SchemaNull); ok {
				omitEmpty = ""
			}

			properties = append(properties, utils.CodeBlock{
				utils.CodeGen.Comment.Go(prop.description, 0),
				fmt.Sprintf("%s %s%s `json:\"%s%s\"`", prop.name.TitleCase(), arrayType, typeName, prop.name, omitEmpty),
			}.DisplayIndent(4))
			continue
		}

		propType := prop.typ.GoType()
		switch prop.typ.(type) {
		case SchemaNull:
			properties = append(properties, utils.CodeBlock{
				utils.CodeGen.Comment.Go(prop.description, 0),
				fmt.Sprintf("%s %s `json:\"%s\"`", prop.name.TitleCase(), prop.name.TitleCase(), prop.name),
			}.DisplayIndent(4))
		case SchemaUnion:
			properties = append(properties, utils.CodeBlock{
				utils.CodeGen.Comment.Go(prop.description, 0),
				fmt.Sprintf("%s %s `json:\"%s,omitempty\"`", prop.name.TitleCase(), prop.name.TitleCase(), prop.name),
			}.DisplayIndent(4))
		default:
			properties = append(properties, fmt.Sprintf("%s\n%s\n",
				utils.CodeGen.Comment.Go(prop.description, 4),
				fmt.Sprintf("    %s %s `json:\"%s,omitempty\"`", prop.name.TitleCase(), propType, prop.name),
			))
		}
	}

	return utils.CodeBlock{
		fmt.Sprintf("type %s struct {", s.Name.TitleCase()),
		strings.Join(properties, "\n"),
		"}",
	}.Display()
}

type object struct{}

func (object) New(name string, props []Field) SchemaObject {
	return SchemaObject{
		Name:       AttributeName(name),
		Properties: props,
	}
}

var Object = object{}
