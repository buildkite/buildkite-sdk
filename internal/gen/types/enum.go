package types

import (
	"fmt"
	"sort"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
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

func (Enum) IsPrimitive() bool {
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
				fmt.Sprintf("// %s", e.Description),
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
func (e Enum) TypeScript() string {
	typ := typescript.NewType(
		e.Name.ToTitleCase(),
		e.Description,
		e.TypeScriptInterfaceType(),
	)

	return typ.String()
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

	block.AddLines(fmt.Sprintf("%s = %s", e.Name.ToTitleCase(), e.PythonClassType()))
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

// CSharp
func (e Enum) canBeCSharpEnum() bool {
	if len(e.Values) == 0 {
		return false
	}
	for _, v := range e.Values {
		if s, ok := v.(string); ok {
			if s == "*" || strings.ContainsAny(s, " -./") {
				return false
			}
		}
	}
	return true
}

func (e Enum) CSharp() (string, error) {
	if len(e.Values) == 0 {
		return "", nil
	}

	if !e.canBeCSharpEnum() {
		return e.generateCSharpStringConstants(), nil
	}

	return e.generateCSharpEnum(), nil
}

func (e Enum) generateCSharpEnum() string {
	var sb strings.Builder

	if e.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", e.Description))
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString("[JsonConverter(typeof(JsonStringEnumConverter))]\n")
	sb.WriteString(fmt.Sprintf("public enum %s\n{\n", e.Name.ToTitleCase()))

	for i, val := range e.Values {
		strVal := fmt.Sprintf("%v", val)
		enumName := sanitizeEnumValue(strVal)
		if enumName != strVal {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", strVal))
		}
		sb.WriteString(fmt.Sprintf("    %s", enumName))
		if i < len(e.Values)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

func (e Enum) generateCSharpStringConstants() string {
	var sb strings.Builder

	if e.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", e.Description))
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString(fmt.Sprintf("public static class %sValues\n{\n", e.Name.ToTitleCase()))

	for _, val := range e.Values {
		strVal := fmt.Sprintf("%v", val)
		constName := sanitizeEnumValue(strVal)
		if constName == "" {
			constName = "Empty"
		}
		sb.WriteString(fmt.Sprintf("    public const string %s = \"%s\";\n", constName, strVal))
	}

	sb.WriteString("}\n")
	return sb.String()
}

func sanitizeEnumValue(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
		if i == 0 && r >= '0' && r <= '9' {
			result := "_" + sb.String()
			sb.Reset()
			sb.WriteString(result)
		}
	}
	result := sb.String()
	if result == "" {
		return result
	}
	runes := []rune(result)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 32
	}
	return string(runes)
}

func (e Enum) CSharpType() string {
	if !e.canBeCSharpEnum() {
		return "string"
	}
	return e.Name.ToTitleCase()
}
