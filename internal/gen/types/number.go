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
