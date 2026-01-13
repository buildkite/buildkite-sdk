package csharp

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// TypeMapping maps JSON schema types to C# types
var TypeMapping = map[string]string{
	"string":  "string",
	"integer": "int",
	"number":  "double",
	"boolean": "bool",
	"object":  "object",
	"array":   "List",
}

// ToTitleCase converts a string to TitleCase (PascalCase)
func ToTitleCase(s string) string {
	if s == "" {
		return s
	}

	// Handle snake_case
	if strings.Contains(s, "_") {
		parts := strings.Split(s, "_")
		for i, part := range parts {
			parts[i] = ToTitleCase(part)
		}
		return strings.Join(parts, "")
	}

	// Handle dash-case
	if strings.Contains(s, "-") {
		parts := strings.Split(s, "-")
		for i, part := range parts {
			parts[i] = ToTitleCase(part)
		}
		return strings.Join(parts, "")
	}

	// Simple capitalize first letter
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	title := ToTitleCase(s)
	if title == "" {
		return title
	}
	runes := []rune(title)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SanitizeEnumValue converts a string value to a valid C# enum identifier
func SanitizeEnumValue(s string) string {
	// Remove invalid characters
	reg := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	sanitized := reg.ReplaceAllString(s, "_")

	// Ensure it starts with a letter
	if len(sanitized) > 0 && unicode.IsDigit(rune(sanitized[0])) {
		sanitized = "_" + sanitized
	}

	// Convert to TitleCase
	return ToTitleCase(sanitized)
}

// CSharpReservedWords are C# reserved keywords that need escaping
var CSharpReservedWords = map[string]bool{
	"abstract": true, "as": true, "base": true, "bool": true,
	"break": true, "byte": true, "case": true, "catch": true,
	"char": true, "checked": true, "class": true, "const": true,
	"continue": true, "decimal": true, "default": true, "delegate": true,
	"do": true, "double": true, "else": true, "enum": true,
	"event": true, "explicit": true, "extern": true, "false": true,
	"finally": true, "fixed": true, "float": true, "for": true,
	"foreach": true, "goto": true, "if": true, "implicit": true,
	"in": true, "int": true, "interface": true, "internal": true,
	"is": true, "lock": true, "long": true, "namespace": true,
	"new": true, "null": true, "object": true, "operator": true,
	"out": true, "override": true, "params": true, "private": true,
	"protected": true, "public": true, "readonly": true, "ref": true,
	"return": true, "sbyte": true, "sealed": true, "short": true,
	"sizeof": true, "stackalloc": true, "static": true, "string": true,
	"struct": true, "switch": true, "this": true, "throw": true,
	"true": true, "try": true, "typeof": true, "uint": true,
	"ulong": true, "unchecked": true, "unsafe": true, "ushort": true,
	"using": true, "virtual": true, "void": true, "volatile": true,
	"while": true, "async": true, "await": true,
}

// EscapeReservedWord escapes a C# reserved word
func EscapeReservedWord(s string) string {
	lower := strings.ToLower(s)
	if CSharpReservedWords[lower] {
		return "@" + s
	}
	return s
}

// GetCSharpType returns the C# type for a JSON schema type
func GetCSharpType(jsonType string, isNullable bool) string {
	csType, ok := TypeMapping[jsonType]
	if !ok {
		csType = "object"
	}

	if isNullable {
		return csType + "?"
	}
	return csType
}

// GenerateTypeAlias generates a C# type alias using global using
func GenerateTypeAlias(name, targetType, description string) string {
	var sb strings.Builder

	if description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", description))
		sb.WriteString("/// </summary>\n")
	}

	// C# doesn't have direct type aliases, so we create a wrapper or use the type directly
	// For simple aliases, we'll generate a comment and the actual type reference
	sb.WriteString(fmt.Sprintf("// Type alias: %s = %s\n", name, targetType))

	return sb.String()
}

// GenerateUnionType generates a C# representation of a union type
func GenerateUnionType(name, description string, variants []string) string {
	var sb strings.Builder

	if description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", description))
		sb.WriteString("/// </summary>\n")
	}

	// Generate JsonDerivedType attributes for polymorphic serialization
	for _, variant := range variants {
		sb.WriteString(fmt.Sprintf("[JsonDerivedType(typeof(%s))]\n", variant))
	}

	sb.WriteString(fmt.Sprintf("public interface I%s { }\n", name))

	return sb.String()
}
