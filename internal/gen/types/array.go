package types

import (
	"fmt"
	"strings"

	gogen "github.com/buildkite/buildkite-sdk/internal/gen/go"
	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Array struct {
	Name        PropertyName
	Description string
	Type        Value
	Reference   bool
}

func (a Array) GetDescription() string {
	return a.Description
}

func (a Array) IsReference() bool {
	return a.Reference
}

func (Array) IsPrimitive() bool {
	return false
}

// Go
func (a Array) GoStructType() string {
	if a.IsReference() || isPrimitiveValue(a.Type) {
		return fmt.Sprintf("[]%s", a.Type.GoStructType())
	}

	typeValue := a.Name.ToTitleCase()
	if isUnionValue(a.Type) {
		typeValue = fmt.Sprintf("%sItem", typeValue)
	}
	return fmt.Sprintf("[]%s", typeValue)
}

func (a Array) GoStructKey(isUnion bool) string {
	if isUnion {
		if isPrimitiveValue(a.Type) {
			return fmt.Sprintf("%sArray", utils.CamelCaseToTitleCase(a.Type.GoStructType()))
		}

		if isUnionValue(a.Type) {
			return fmt.Sprintf("%sItem", a.Name.ToTitleCase())
		}
	}

	return a.Name.ToTitleCase()
}

func (a Array) Go() (string, error) {
	contents := utils.NewCodeBlock()
	typeValue := a.GoStructType()

	if isUnionValue(a.Type) && !a.IsReference() {
		union := a.Type.(Union)
		item := Union{
			Name:            NewPropertyName(fmt.Sprintf("%sItem", a.Name.Value)),
			Description:     union.Description,
			TypeIdentifiers: union.TypeIdentifiers,
		}

		itemLines, err := item.Go()
		if err != nil {
			return "", fmt.Errorf("generating lines for union in array [%s]: %v", a.Name.Value, err)
		}

		contents.AddLines(itemLines)
	}

	typ := gogen.NewType(
		a.Name.ToTitleCase(),
		a.Description,
		typeValue,
	)
	contents.AddLines(typ.String())
	return contents.String(), nil
}

// TypeScript
func (a Array) TypeScript() string {
	arrayType := a.Type.TypeScriptInterfaceType()
	if _, ok := a.Type.(Union); ok {
		arrayType = fmt.Sprintf("(%s)", arrayType)
	}

	typ := typescript.NewType(
		a.Name.ToTitleCase(),
		a.Description,
		fmt.Sprintf("%s[]", arrayType),
	)
	return typ.String()
}

func (a Array) TypeScriptInterfaceType() string {
	if a.IsReference() {
		return fmt.Sprintf("%s[]", a.Type.GoStructType())
	}

	switch a.Type.(type) {
	case String:
		return "string[]"
	case Boolean:
		return "boolean[]"
	case Number:
		return "number[]"
	default:
		return fmt.Sprintf("%s[]", a.Name.ToTitleCase())
	}
}

func (a Array) TypeScriptInterfaceKey() string {
	return a.Name.Value
}

// Python
func (a Array) Python() (string, error) {
	codeBlock := utils.NewCodeBlock()
	listType := a.Type.PythonClassType()

	if union, ok := a.Type.(Union); ok {
		var unionTypeParts []string
		for _, typ := range union.TypeIdentifiers {
			if obj, ok := typ.(Object); ok {
				nestedObj := Object{
					Name:                 NewPropertyName(fmt.Sprintf("%sObject", obj.Name.Value)),
					Properties:           obj.Properties,
					AdditionalProperties: obj.AdditionalProperties,
				}

				objLines, err := nestedObj.Python()
				if err != nil {
					return "", fmt.Errorf("generating object lines for union [%s]: %v", a.Name.Value, err)
				}

				codeBlock.AddLines(objLines)
				unionType := nestedObj.PythonClassType()
				unionTypeParts = append(unionTypeParts, unionType)
				if len(obj.Properties.Keys()) > 0 {
					unionTypeParts = append(unionTypeParts, fmt.Sprintf("%sArgs", unionType))
				}

				continue
			}

			if ref, ok := typ.(PropertyReference); ok {
				if _, ok := ref.Type.(Object); ok {
					unionTypeParts = append(unionTypeParts, fmt.Sprintf("%sArgs", typ.PythonClassType()))
				}
			}

			unionTypeParts = append(unionTypeParts, typ.PythonClassType())
		}
		listType = strings.Join(unionTypeParts, " | ")
	}

	if _, ok := a.Type.(Object); ok {
		listType = fmt.Sprintf("%s | %sArgs", listType, listType)
	}

	pyType := fmt.Sprintf("%s = List[%s]", a.Name.ToTitleCase(), listType)

	if a.Description != "" {
		codeBlock.AddLines(fmt.Sprintf("# %s", a.Description))
	}

	codeBlock.AddLines(pyType)
	return codeBlock.String(), nil
}

func (a Array) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(a.Name.Value)
}

func (a Array) PythonClassType() string {
	return fmt.Sprintf("List[%s]", a.Type.PythonClassType())
}
