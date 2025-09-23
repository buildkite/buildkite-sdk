package types

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

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

// Python
func (s String) Python() (string, error) {
	return fmt.Sprintf("type %s = str", s.Name.ToTitleCase()), nil
}

func (s String) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(s.Name.Value)
}

func (s String) PythonClassType() string {
	return "str"
}
