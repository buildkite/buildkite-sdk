package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

func parseEnumValue(val any) EnumValue {
	if s, ok := val.(string); ok {
		return EnumValue{
			Value: s,
			Type:  String{},
		}
	}

	if b, ok := val.(bool); ok {
		return EnumValue{
			Value: b,
			Type:  Boolean{},
		}
	}

	if n, ok := val.(int); ok {
		return EnumValue{
			Value: n,
			Type: Number{
				Name: NewPropertyName("Int"),
			},
		}
	}

	panic("not implemented enum type")
}

type EnumValue struct {
	Type  Value
	Value any
}

type Enum struct {
	Name        PropertyName
	Description string
	Values      []any
	Default     any
}

func (e Enum) IsReference() bool {
	return false
}

func (Enum) IsPrimative() bool {
	return false
}

func (e Enum) GoStructType() string {
	return e.Name.ToTitleCase()
}

func (e Enum) GoStructKey(isUnion bool) string {
	return e.Name.ToTitleCase()
}

func (e Enum) Go() (string, error) {
	lines := utils.NewCodeBlock()

	displayName := e.Name.ToTitleCase()
	enumDefinitionName := e.GoStructType()

	enumInterface := utils.NewGoConstraintInterface(fmt.Sprintf("%sValues", enumDefinitionName))
	enumDefinition := utils.NewGoStruct(enumDefinitionName, nil)

	enumMarshalFunction := utils.NewCodeBlock(
		fmt.Sprintf("func (e %s) MarshalJSON() ([]byte, error) {", displayName),
	)

	enumTypes := map[string]string{}
	for _, val := range e.Values {
		parsed := parseEnumValue(val)
		typ := parsed.Type.GoStructType()
		enumTypes[typ] = typ
	}

	// If there is only one type in the values.
	if len(enumTypes) == 1 {
		for typ := range enumTypes {
			if typ != "string" {
				panic("type not supported in single enum")
			}

			lines.AddLines(
				fmt.Sprintf("type %s %s", enumDefinitionName, typ),
				fmt.Sprintf("var %sValues = map[string]%s{", enumDefinitionName, enumDefinitionName),
			)

			for _, val := range e.Values {
				lines.AddLines(
					fmt.Sprintf("   \"%s\": \"%v\",", val, val),
				)
			}

			lines.AddLines("}")
			return lines.String(), nil
		}
	}

	for typ := range enumTypes {
		titleCaseType := utils.CamelCaseToTitleCase(typ)

		enumInterface.AddItem(typ)
		enumDefinition.AddItem(titleCaseType, typ, "", true)

		enumMarshalFunction.AddLines(
			fmt.Sprintf("    if e.%s != nil {\n        return json.Marshal(e.%s)\n    }", titleCaseType, titleCaseType),
		)
	}

	enumMarshalFunction.AddLines("    return json.Marshal(nil)\n}")

	enumInterfaceString, err := enumInterface.Write()
	if err != nil {
		return "", fmt.Errorf("generating interface for [%s]: %v", e.Name.Value, err)
	}

	enumDefinitionString, err := enumDefinition.Write()
	if err != nil {
		return "", fmt.Errorf("generating struct for [%s]: %v", e.Name.Value, err)
	}

	lines.AddLines(
		enumInterfaceString,
		enumDefinitionString,
		enumMarshalFunction.String(),
	)

	return lines.String(), nil
}
