package types

import (
	"fmt"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Array struct {
	Name      PropertyName
	Type      Value
	Reference bool
}

func (a Array) IsReference() bool {
	return a.Reference
}

func (Array) IsPrimative() bool {
	return false
}

// Go
func (a Array) GoStructType() string {
	if a.IsReference() {
		return fmt.Sprintf("[]%s", a.Type.GoStructType())
	}

	switch a.Type.(type) {
	case String:
		return "[]string"
	case Boolean:
		return "[]bool"
	case Number:
		return "[]int"
	case Union:
		return fmt.Sprintf("[]%sUnion", a.Name.ToTitleCase())
	default:
		return fmt.Sprintf("[]%s", a.Name.ToTitleCase())
	}
}

func (a Array) GoStructKey(isUnion bool) string {
	if isUnion {
		switch a.Type.(type) {
		case String:
			return "StringArray"
		case Boolean:
			return "BoolArray"
		case Number:
			return "IntArray"
		case Union:
			return fmt.Sprintf("%sUnion", a.Name.ToTitleCase())
		default:
			return a.Name.ToTitleCase()
		}
	}

	return a.Name.ToTitleCase()
}

func (a Array) Go() (string, error) {
	union, ok := a.Type.(Union)
	if !a.IsReference() && ok {
		item := Union{
			Name:            NewPropertyName(fmt.Sprintf("%sUnion", a.Name.Value)),
			Description:     union.Description,
			TypeIdentifiers: union.TypeIdentifiers,
		}

		lines := utils.NewCodeBlock()
		itemLines, err := item.Go()
		if err != nil {
			return "", fmt.Errorf("generating lines for union in array [%s]: %v", a.Name.Value, err)
		}

		lines.AddLines(
			itemLines,
			fmt.Sprintf("type %s = []%sUnion", a.Name.ToTitleCase(), a.Type.GoStructType()),
		)
		return lines.String(), nil
	}

	return fmt.Sprintf("type %s = []%s", a.Name.ToTitleCase(), a.Type.GoStructType()), nil
}

// TypeScript
func (a Array) TypeScript() (string, error) {
	if union, ok := a.Type.(Union); ok {
		codeBlock := utils.NewCodeBlock(
			fmt.Sprintf("export type %s = (%s)[]", a.Name.ToTitleCase(), union.TypeScriptInterfaceType()),
		)

		return codeBlock.String(), nil
	}

	return fmt.Sprintf("export type %s = %s[]", a.Name.ToTitleCase(), a.Type.TypeScriptInterfaceType()), nil
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
	return a.Name.ToTitleCase()
}

// Python
func (a Array) Python() (string, error) {
	codeBlock := utils.NewCodeBlock()
	listType := a.Type.PythonClassType()
	listLength := 1

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
					unionTypeParts = append(unionTypeParts, fmt.Sprintf("%sDict", unionType))
				}

				continue
			}

			if ref, ok := typ.(PropertyReference); ok {
				if _, ok := ref.Type.(Object); ok {
					unionTypeParts = append(unionTypeParts, fmt.Sprintf("%sDict", typ.PythonClassType()))
				}
			}

			unionTypeParts = append(unionTypeParts, typ.PythonClassType())
		}
		listLength = len(unionTypeParts)
		listType = strings.Join(unionTypeParts, ",")
	}

	if _, ok := a.Type.(Object); ok {
		listLength = 2
		listType = fmt.Sprintf("%s,%sDict", listType, listType)
	}

	pyType := fmt.Sprintf("type %s = List[Union[%s]]", a.Name.ToTitleCase(), listType)
	if listLength == 1 {
		pyType = fmt.Sprintf("type %s = List[%s]", a.Name.ToTitleCase(), listType)
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
