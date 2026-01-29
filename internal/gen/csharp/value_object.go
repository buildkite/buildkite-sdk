package csharp

import (
	"fmt"
	"strings"
)

// ObjectProperty represents a property in an object
type ObjectProperty struct {
	Name        string
	Type        string
	JsonName    string
	Description string
	IsNullable  bool
	IsRequired  bool
}

// ObjectValue generates C# code for an object/class type
type ObjectValue struct {
	Name                 string
	Description          string
	Properties           []ObjectProperty
	Implements           []string
	AdditionalProperties string // Type for dictionary-style objects
}

// CSharp generates C# code for a class
func (o ObjectValue) CSharp() (string, error) {
	if o.Name == "" {
		return "", nil
	}

	// If this is a dictionary type (additionalProperties only, no fixed properties)
	if o.AdditionalProperties != "" && len(o.Properties) == 0 {
		return o.generateDictionaryType(), nil
	}

	return o.generateClass(), nil
}

func (o ObjectValue) generateClass() string {
	var sb strings.Builder

	// XML documentation
	if o.Description != "" {
		sb.WriteString("/// <summary>\n")
		for _, line := range strings.Split(o.Description, "\n") {
			sb.WriteString(fmt.Sprintf("/// %s\n", strings.TrimSpace(line)))
		}
		sb.WriteString("/// </summary>\n")
	}

	// Class declaration
	className := ToTitleCase(o.Name)
	sb.WriteString(fmt.Sprintf("public class %s", className))

	if len(o.Implements) > 0 {
		sb.WriteString(" : ")
		sb.WriteString(strings.Join(o.Implements, ", "))
	}

	sb.WriteString("\n{\n")

	// Properties
	for _, prop := range o.Properties {
		// Property documentation
		if prop.Description != "" {
			sb.WriteString("    /// <summary>\n")
			desc := strings.ReplaceAll(prop.Description, "\n", " ")
			sb.WriteString(fmt.Sprintf("    /// %s\n", desc))
			sb.WriteString("    /// </summary>\n")
		}

		// JSON property name attribute if different
		propName := ToTitleCase(prop.Name)
		if prop.JsonName != "" && prop.JsonName != propName && prop.JsonName != ToSnakeCase(propName) {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", prop.JsonName))
		}

		// Determine type with nullability
		typeName := prop.Type
		if prop.IsNullable && !strings.HasSuffix(typeName, "?") && !isReferenceType(typeName) {
			typeName += "?"
		}

		// Escape reserved words
		escapedName := EscapeReservedWord(propName)

		sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n\n", typeName, escapedName))
	}

	sb.WriteString("}\n")

	return sb.String()
}

func (o ObjectValue) generateDictionaryType() string {
	var sb strings.Builder

	if o.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", o.Description))
		sb.WriteString("/// </summary>\n")
	}

	className := ToTitleCase(o.Name)
	valueType := o.AdditionalProperties
	if valueType == "" {
		valueType = "object"
	}

	sb.WriteString(fmt.Sprintf("public class %s : Dictionary<string, %s>\n{\n", className, valueType))
	sb.WriteString(fmt.Sprintf("    public %s() : base() { }\n", className))
	sb.WriteString(fmt.Sprintf("    public %s(IDictionary<string, %s> dictionary) : base(dictionary) { }\n", className, valueType))
	sb.WriteString("}\n")

	return sb.String()
}

// isReferenceType checks if a type is a reference type (already nullable)
func isReferenceType(typeName string) bool {
	referenceTypes := map[string]bool{
		"string": true,
		"object": true,
	}

	// Lists, arrays, and custom classes are reference types
	if strings.HasPrefix(typeName, "List<") ||
		strings.HasPrefix(typeName, "Dictionary<") ||
		strings.HasSuffix(typeName, "[]") {
		return true
	}

	// Check if it's a known reference type
	if referenceTypes[typeName] {
		return true
	}

	// Custom classes (PascalCase names) are reference types
	if len(typeName) > 0 && typeName[0] >= 'A' && typeName[0] <= 'Z' {
		return true
	}

	return false
}

// CSharpType returns the C# type name
func (o ObjectValue) CSharpType() string {
	return ToTitleCase(o.Name)
}

// CSharpInlineType returns the inline type for use in properties
func (o ObjectValue) CSharpInlineType() string {
	return o.CSharpType()
}
