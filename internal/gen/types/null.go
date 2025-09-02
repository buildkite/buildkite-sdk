package types

type Null struct {
	Reference bool
}

func (n Null) IsReference() bool {
	return n.Reference
}

func (Null) IsPrimative() bool {
	return false
}

func (Null) GoStructKey(isUnion bool) string {
	return "Null"
}

func (Null) GoStructType() string {
	return "Null"
}

func (Null) Go() (string, error) {
	return `type Null struct{}

func (Null) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}`, nil
}
