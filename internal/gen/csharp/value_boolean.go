package csharp

// BooleanValue generates C# code for a boolean type
type BooleanValue struct {
	Name        string
	Description string
}

// CSharp generates C# code for a boolean type
func (b BooleanValue) CSharp() (string, error) {
	// Booleans are primitives, no wrapper needed
	return "", nil
}

// CSharpType returns the C# type name
func (b BooleanValue) CSharpType() string {
	return "bool"
}

// CSharpInlineType returns the inline type for use in properties
func (b BooleanValue) CSharpInlineType() string {
	return "bool"
}
