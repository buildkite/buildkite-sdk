package csharp

import (
	"fmt"
	"strings"
)

// UnionVariant represents a variant in a union type
type UnionVariant struct {
	Name string
	Type string
}

// UnionValue generates C# code for a union type
type UnionValue struct {
	Name        string
	Description string
	Variants    []UnionVariant
}

// CSharp generates C# code for a union type
func (u UnionValue) CSharp() (string, error) {
	if u.Name == "" || len(u.Variants) == 0 {
		return "", nil
	}

	// For unions with only primitive types, use object
	allPrimitive := true
	for _, v := range u.Variants {
		if !isPrimitiveType(v.Type) {
			allPrimitive = false
			break
		}
	}

	if allPrimitive {
		// Simple union of primitives - just use object
		return "", nil
	}

	// For complex unions, generate an interface with JsonDerivedType attributes
	return u.generateInterface(), nil
}

func (u UnionValue) generateInterface() string {
	var sb strings.Builder

	if u.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", u.Description))
		sb.WriteString("/// </summary>\n")
	}

	// Add JsonDerivedType for each non-primitive variant
	for _, variant := range u.Variants {
		if !isPrimitiveType(variant.Type) && variant.Type != "object" {
			sb.WriteString(fmt.Sprintf("[JsonDerivedType(typeof(%s))]\n", variant.Type))
		}
	}

	interfaceName := "I" + ToTitleCase(u.Name)
	sb.WriteString(fmt.Sprintf("public interface %s { }\n", interfaceName))

	return sb.String()
}

func isPrimitiveType(t string) bool {
	primitives := map[string]bool{
		"string": true,
		"int":    true,
		"long":   true,
		"double": true,
		"float":  true,
		"bool":   true,
		"object": true,
	}
	return primitives[t]
}

// CSharpType returns the C# type name
func (u UnionValue) CSharpType() string {
	// If all variants are primitives, return object
	allPrimitive := true
	for _, v := range u.Variants {
		if !isPrimitiveType(v.Type) {
			allPrimitive = false
			break
		}
	}

	if allPrimitive {
		return "object"
	}

	return "I" + ToTitleCase(u.Name)
}

// CSharpInlineType returns the inline type for use in properties
func (u UnionValue) CSharpInlineType() string {
	return u.CSharpType()
}

// GenerateUnionHelpers generates helper methods for working with unions
func GenerateUnionHelpers(name string, variants []UnionVariant) string {
	var sb strings.Builder

	typeName := ToTitleCase(name)

	sb.WriteString(fmt.Sprintf("public static class %sExtensions\n{\n", typeName))

	for _, variant := range variants {
		if !isPrimitiveType(variant.Type) {
			variantName := ToTitleCase(variant.Name)
			sb.WriteString(fmt.Sprintf("    public static bool Is%s(this I%s value) => value is %s;\n",
				variantName, typeName, variant.Type))
			sb.WriteString(fmt.Sprintf("    public static %s? As%s(this I%s value) => value as %s;\n",
				variant.Type, variantName, typeName, variant.Type))
		}
	}

	sb.WriteString("}\n")

	return sb.String()
}
