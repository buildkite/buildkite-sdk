package types

import (
	"fmt"
	"sort"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
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

func (e Enum) GetDescription() string {
	return e.Description
}

func (e Enum) IsReference() bool {
	return false
}

func (Enum) IsPrimative() bool {
	return false
}

// Go
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

	enumInterface := utils.NewGoConstraintInterface(fmt.Sprintf("%sValues", enumDefinitionName), e.Description)
	enumDefinition := utils.NewGoStruct(enumDefinitionName, e.Description, nil)

	enumMarshalFunction := utils.NewCodeBlock(
		fmt.Sprintf("func (e %s) MarshalJSON() ([]byte, error) {", displayName),
	)

	enumTypes := orderedmap.New()
	for _, val := range e.Values {
		parsed := parseEnumValue(val)
		typ := parsed.Type.GoStructType()
		enumTypes.Set(typ, typ)
	}
	enumTypes.SortKeys(sort.Strings)

	// If there is only one type in the values.
	if len(enumTypes.Keys()) == 1 {
		for _, typ := range enumTypes.Keys() {
			if typ != "string" {
				panic("type not supported in single enum")
			}

			lines.AddLines(
				fmt.Sprintf("type %s %s", enumDefinitionName, typ),
				fmt.Sprintf(fmt.Sprintf("// %s", e.Description)),
				fmt.Sprintf("var %sValues = map[string]%s{", enumDefinitionName, enumDefinitionName),
			)

			for _, val := range e.Values {
				lines.AddLines(
					fmt.Sprintf("   \"%s\": \"%v\",", val, val),
				)
			}

			lines.AddLines("}")
		}
		return lines.String(), nil
	}

	for _, typ := range enumTypes.Keys() {
		titleCaseType := utils.CamelCaseToTitleCase(typ)

		enumInterface.AddItem(typ)
		enumDefinition.AddItem(titleCaseType, typ, "", "", true)

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

// TypeScript
func (e Enum) TypeScript() (string, error) {
	parts := make([]string, len(e.Values))
	for i, val := range e.Values {
		if _, ok := val.(string); ok {
			parts[i] = fmt.Sprintf("'%v'", val)
			continue
		}

		parts[i] = fmt.Sprintf("%v", val)
	}
	values := strings.Join(parts, " | ")

	block := utils.NewCodeBlock()
	if e.Description != "" {
		block.AddLines(fmt.Sprintf("// %s", e.Description))
	}

	block.AddLines(fmt.Sprintf("export type %s = %s", e.Name.ToTitleCase(), values))
	return block.String(), nil
}

func (e Enum) TypeScriptInterfaceKey() string {
	return e.Name.Value
}

func (e Enum) TypeScriptInterfaceType() string {
	parts := make([]string, len(e.Values))
	for i, val := range e.Values {
		if _, ok := val.(string); ok {
			parts[i] = fmt.Sprintf("'%v'", val)
			continue
		}

		parts[i] = fmt.Sprintf("%v", val)
	}

	return strings.Join(parts, " | ")
}

// Python
func (e Enum) Python() (string, error) {
	block := utils.NewCodeBlock()
	if e.Description != "" {
		block.AddLines(fmt.Sprintf("# %s", e.Description))
	}

	block.AddLines(fmt.Sprintf("type %s = %s", e.Name.ToTitleCase(), e.PythonClassType()))
	return block.String(), nil
}

func (e Enum) PythonClassKey() string {
	return utils.CamelCaseToTitleCase(e.Name.Value)
}

func (e Enum) PythonClassType() string {
	parts := make([]string, len(e.Values))
	for i, val := range e.Values {
		if val == true {
			parts[i] = "True"
			continue
		}

		if val == false {
			parts[i] = "False"
			continue
		}

		if _, ok := val.(string); ok {
			parts[i] = fmt.Sprintf("'%v'", val)
			continue
		}

		parts[i] = fmt.Sprintf("%v", val)
	}

	return fmt.Sprintf("Literal[%s]", strings.Join(parts, ","))
}
