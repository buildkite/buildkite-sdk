package csharp

// StringValue generates C# code for a string type
type StringValue struct {
	Name        string
	Description string
}

// CSharp generates C# code for a string type alias
func (s StringValue) CSharp() (string, error) {
	// In C#, simple string types don't need a wrapper - they're just `string`
	// We only generate a type if it has special meaning
	if s.Name == "" {
		return "", nil
	}

	return GenerateTypeAlias(s.Name, "string", s.Description), nil
}

// CSharpType returns the C# type name
func (s StringValue) CSharpType() string {
	if s.Name != "" {
		return ToTitleCase(s.Name)
	}
	return "string"
}

// CSharpInlineType returns the inline type for use in properties
func (s StringValue) CSharpInlineType() string {
	return "string"
}
