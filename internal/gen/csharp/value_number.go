package csharp

// NumberValue generates C# code for a number/integer type
type NumberValue struct {
	Name        string
	Description string
	IsInteger   bool
}

// CSharp generates C# code for a number type
func (n NumberValue) CSharp() (string, error) {
	// Numbers are primitives, no wrapper needed
	return "", nil
}

// CSharpType returns the C# type name
func (n NumberValue) CSharpType() string {
	if n.IsInteger {
		return "int"
	}
	return "double"
}

// CSharpInlineType returns the inline type for use in properties
func (n NumberValue) CSharpInlineType() string {
	return n.CSharpType()
}
