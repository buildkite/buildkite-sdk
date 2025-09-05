package types

import "fmt"

type String struct {
	Name PropertyName
}

func (s String) IsReference() bool {
	return false
}

func (String) IsPrimative() bool {
	return true
}

// Go
func (s String) Go() (string, error) {
	return fmt.Sprintf("type %s = string", s.Name.ToTitleCase()), nil
}

func (s String) GoStructType() string {
	return "string"
}

func (s String) GoStructKey(isUnion bool) string {
	if isUnion {
		return "String"
	}

	return s.Name.ToTitleCase()
}

// TypeScript
func (s String) TypeScript() (string, error) {
	return fmt.Sprintf("export type %s = string", s.Name.ToTitleCase()), nil
}

func (s String) TypeScriptInterfaceKey() string {
	return s.Name.Value
}

func (String) TypeScriptInterfaceType() string {
	return "string"
}
