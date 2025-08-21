package types

import "fmt"

type String struct {
	Name PropertyName
}

func (s String) IsReference() bool {
	return false
}

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
