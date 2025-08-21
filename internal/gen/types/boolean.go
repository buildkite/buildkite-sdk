package types

import "fmt"

type Boolean struct {
	Name PropertyName
}

func (b Boolean) IsReference() bool {
	return false
}

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
