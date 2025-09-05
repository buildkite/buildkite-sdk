package types

import "fmt"

type Boolean struct {
	Name PropertyName
}

func (b Boolean) IsReference() bool {
	return false
}

func (Boolean) IsPrimative() bool {
	return true
}

// Go
func (Boolean) GoStructType() string {
	return "bool"
}

func (b Boolean) GoStructKey(isUnion bool) string {
	if isUnion {
		return "Bool"
	}

	return b.Name.ToTitleCase()
}

func (b Boolean) Go() (string, error) {
	return fmt.Sprintf("type %s = string", b.Name.ToTitleCase()), nil
}

// TypeScript

func (b Boolean) TypeScript() (string, error) {
	return fmt.Sprintf("type %s = boolean", b.Name.ToTitleCase()), nil
}

func (b Boolean) TypeScriptInterfaceKey() string {
	return b.Name.Value
}

func (b Boolean) TypeScriptInterfaceType() string {
	return "boolean"
}
