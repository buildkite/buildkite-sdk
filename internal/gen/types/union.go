package types

import (
	"fmt"
	"slices"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Union struct {
	Name            PropertyName
	Description     string
	TypeIdentifiers []Value
}

func (u Union) GetDescription() string {
	return u.Description
}

func (u Union) IsReference() bool {
	return false
}

func (u Union) IsPrimative() bool {
	return false
}

// Go
func (u Union) GoStructKey(isUnion bool) string {
	return u.Name.ToTitleCase()
}

func (u Union) GoStructType() string {
	return u.Name.ToTitleCase()
}

func (u Union) Go() (string, error) {
	lines := utils.NewCodeBlock()

	displayName := u.Name.ToTitleCase()
	unionValuesName := fmt.Sprintf("%sValues", displayName)

	unionInterface := utils.NewGoConstraintInterface(unionValuesName, u.Description)
	unionDefinition := utils.NewGoStruct(displayName, u.Description, nil)

	unionMarshalFunction := utils.NewCodeBlock(
		fmt.Sprintf("func (e %s) MarshalJSON() ([]byte, error) {", displayName),
	)

	for _, typ := range u.TypeIdentifiers {
		titleCaseType := typ.GoStructType()
		structKey := typ.GoStructKey(true)
		isPointer := true

		// Enum
		if enum, ok := typ.(Enum); ok {
			enumLines, err := enum.Go()
			if err != nil {
				return "", fmt.Errorf("generating enum lines for union [%s]: %v", u.Name.Value, err)
			}

			lines.AddLines(enumLines)
		}

		// Object
		if obj, ok := typ.(Object); ok {
			nestedObj := Object{
				Name:                 NewPropertyName(fmt.Sprintf("%sObject", obj.Name.Value)),
				Properties:           obj.Properties,
				AdditionalProperties: obj.AdditionalProperties,
			}

			objLines, err := nestedObj.Go()
			if err != nil {
				return "", fmt.Errorf("generating object lines for union [%s]: %v", u.Name.Value, err)
			}

			titleCaseType = nestedObj.GoStructType()
			lines.AddLines(objLines)
		}

		// Array
		if array, ok := typ.(Array); ok {
			isPointer = false

			switch array.Type.(type) {
			case String:
			case Boolean:
			case Number:
			default:
				arrayLines, err := array.Go()
				if err != nil {
					return "", fmt.Errorf("generating array lines for union [%s]: %v", u.Name.Value, err)
				}

				lines.AddLines(arrayLines)
			}
		}

		unionInterface.AddItem(titleCaseType)
		unionDefinition.AddItem(structKey, titleCaseType, "", typ.GetDescription(), isPointer)
		unionMarshalFunction.AddLines(
			fmt.Sprintf("    if e.%s != nil {\n        return json.Marshal(e.%s)\n    }", structKey, structKey),
		)
	}

	unionMarshalFunction.AddLines("    return json.Marshal(nil)\n}")

	unionInterfaceLines, err := unionInterface.Write()
	if err != nil {
		return "", fmt.Errorf("generating union interface [%s]: %v", u.Name.Value, err)
	}

	unionDefinitionLines, err := unionDefinition.Write()
	if err != nil {
		return "", fmt.Errorf("generating union interface [%s]: %v", u.Name.Value, err)
	}

	lines.AddLines(
		unionInterfaceLines,
		unionDefinitionLines,
		unionMarshalFunction.String(),
	)

	return lines.String(), nil
}

// TypeScript
func (u Union) TypeScriptInterfaceKey() string {
	return u.Name.Value
}

func (u Union) TypeScriptInterfaceType() string {
	parts := make([]string, len(u.TypeIdentifiers))
	for i, typ := range u.TypeIdentifiers {
		if obj, ok := typ.(Object); ok {
			if obj.AdditionalProperties != nil {
				prop := *obj.AdditionalProperties
				parts[i] = fmt.Sprintf("Record<string, %s>", prop.TypeScriptInterfaceType())
				continue
			}

			if obj.Properties == nil {
				parts[i] = "Record<string, any>"
				continue
			}

			block := utils.NewTypeScriptInterface("", obj.Description)
			for _, name := range obj.Properties.Keys() {
				prop, _ := obj.Properties.Get(name)
				val := prop.(Value)
				required := slices.Contains(obj.Required, name)

				block.AddItem(name, val.TypeScriptInterfaceType(), val.GetDescription(), required)
			}

			res, _ := block.WriteUnionObject()
			parts[i] = res
			continue
		}

		parts[i] = typ.TypeScriptInterfaceType()
	}
	return strings.Join(parts, " | ")
}

func (u Union) TypeScript() (string, error) {
	block := utils.NewCodeBlock()
	if u.Description != "" {
		block.AddLines(utils.NewTypeDocComment(u.Description))
	}

	block.AddLines(
		fmt.Sprintf("export type %s = %s", u.Name.ToTitleCase(), u.TypeScriptInterfaceType()),
	)

	return block.String(), nil
}

func (u Union) TypeScriptImports() string {
	var imports []string
	for _, typ := range u.TypeIdentifiers {
		if ref, ok := typ.(PropertyReference); ok {
			imports = append(imports,
				fmt.Sprintf("import {%s} from \"./%s.ts\"", ref.TypeScriptInterfaceType(), ref.Name),
			)
		}
	}
	return strings.Join(imports, "\n")
}

// Python
func (u Union) Python() (string, error) {
	codeBlock := utils.NewCodeBlock()

	var parts []string
	for _, typ := range u.TypeIdentifiers {
		// Nested Object
		if obj, ok := typ.(Object); ok {
			nestedObj := Object{
				Name:                 NewPropertyName(fmt.Sprintf("%sObject", obj.Name.Value)),
				Description:          obj.Description,
				Properties:           obj.Properties,
				AdditionalProperties: obj.AdditionalProperties,
			}

			objLines, err := nestedObj.Python()
			if err != nil {
				return "", fmt.Errorf("generating object lines for union [%s]: %v", u.Name.Value, err)
			}

			codeBlock.AddLines(objLines)
			parts = append(parts, nestedObj.PythonClassType())
			continue
		}

		if ref, ok := typ.(PropertyReference); ok {
			if obj, ok := ref.Type.(Object); ok {
				if len(obj.Properties.Keys()) > 0 {
					parts = append(parts, fmt.Sprintf("%sArgs", typ.PythonClassType()))
				}
			}
		}

		parts = append(parts, typ.PythonClassType())
	}

	if u.Description != "" {
		codeBlock.AddLines(fmt.Sprintf("# %s", u.Description))
	}

	codeBlock.AddLines(fmt.Sprintf("type %s = %s", u.Name.ToTitleCase(), strings.Join(parts, " | ")))
	return codeBlock.String(), nil
}

func (u Union) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(u.Name.Value)
}

func (u Union) PythonClassType() string {
	codeBlock := utils.NewCodeBlock()
	parts := make([]string, len(u.TypeIdentifiers))
	for i, typ := range u.TypeIdentifiers {
		parts[i] = typ.PythonClassType()
	}
	codeBlock.AddLines(strings.Join(parts, " | "))
	return codeBlock.String()
}
