package csharp

import (
	"fmt"
	"strings"
)

// EnumValue generates C# code for an enum type
type EnumValue struct {
	Name        string
	Description string
	Values      []string
}

// CSharp generates C# code for an enum
func (e EnumValue) CSharp() (string, error) {
	if e.Name == "" || len(e.Values) == 0 {
		return "", nil
	}

	// Check if all values are strings that can be enum members
	// If they contain special chars or are complex, use string constants instead
	canBeEnum := true
	for _, v := range e.Values {
		if v == "*" || strings.ContainsAny(v, " -./") {
			canBeEnum = false
			break
		}
	}

	if !canBeEnum {
		// Generate a class with string constants
		return e.generateStringConstants(), nil
	}

	return e.generateEnum(), nil
}

func (e EnumValue) generateEnum() string {
	var sb strings.Builder

	if e.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", e.Description))
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString("[JsonConverter(typeof(JsonStringEnumConverter))]\n")
	sb.WriteString(fmt.Sprintf("public enum %s\n{\n", ToTitleCase(e.Name)))

	for i, val := range e.Values {
		enumName := SanitizeEnumValue(val)
		if enumName != val {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", val))
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

func (e EnumValue) generateStringConstants() string {
	var sb strings.Builder

	if e.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", e.Description))
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString(fmt.Sprintf("public static class %sValues\n{\n", ToTitleCase(e.Name)))

	for _, val := range e.Values {
		constName := SanitizeEnumValue(val)
		if constName == "" {
			constName = "Empty"
		}
		sb.WriteString(fmt.Sprintf("    public const string %s = \"%s\";\n", constName, val))
	}

	sb.WriteString("}\n")

	return sb.String()
}

// CSharpType returns the C# type name
func (e EnumValue) CSharpType() string {
	return ToTitleCase(e.Name)
}

// CSharpInlineType returns the inline type for use in properties
func (e EnumValue) CSharpInlineType() string {
	return e.CSharpType()
}
