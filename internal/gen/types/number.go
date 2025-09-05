package types

import "fmt"

type Number struct {
	Name PropertyName
}

func (Number) IsReference() bool {
	return false
}

func (Number) IsPrimative() bool {
	return true
}

// Go
func (n Number) Go() (string, error) {
	return fmt.Sprintf("type %s = int", n.Name.ToTitleCase()), nil
}

func (Number) GoStructType() string {
	return "int"
}

func (n Number) GoStructKey(isUnion bool) string {
	if isUnion {
		return "Int"
	}

	return n.Name.ToTitleCase()
}

// TypeScript
func (n Number) TypeScript() (string, error) {
	return fmt.Sprintf("type %s = number", n.Name.ToTitleCase()), nil
}

func (n Number) TypeScriptInterfaceKey() string {
	return n.Name.Value
}

func (n Number) TypeScriptInterfaceType() string {
	return "number"
}
