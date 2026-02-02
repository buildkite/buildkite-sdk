package csharp

import "fmt"

// ArrayValue generates C# code for an array type
type ArrayValue struct {
	Name        string
	Description string
	ItemType    string // The C# type of array items
}

// CSharp generates C# code for an array type alias
func (a ArrayValue) CSharp() (string, error) {
	if a.Name == "" {
		return "", nil
	}

	// Generate a type alias comment (C# uses List<T> directly)
	return GenerateTypeAlias(a.Name, fmt.Sprintf("List<%s>", a.ItemType), a.Description), nil
}

// CSharpType returns the C# type name
func (a ArrayValue) CSharpType() string {
	return fmt.Sprintf("List<%s>", a.ItemType)
}

// CSharpInlineType returns the inline type for use in properties
func (a ArrayValue) CSharpInlineType() string {
	return a.CSharpType()
}
