package types

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Boolean struct {
	Name        PropertyName
	Description string
}

func (b Boolean) GetDescription() string {
	return b.Description
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
	block := utils.NewCodeBlock()

	if b.Description != "" {
		block.AddLines(fmt.Sprintf("// %s", b.Description))
	}

	block.AddLines(fmt.Sprintf("type %s = boolean", b.Name.ToTitleCase()))
	return block.String(), nil
}

func (b Boolean) TypeScriptInterfaceKey() string {
	return b.Name.Value
}

func (b Boolean) TypeScriptInterfaceType() string {
	return "boolean"
}

// Python
func (b Boolean) Python() (string, error) {
	block := utils.NewCodeBlock()
	if b.Description != "" {
		block.AddLines(fmt.Sprintf("# %s", b.Description))
	}

	block.AddLines(fmt.Sprintf("type %s = bool", b.Name.ToTitleCase()))
	return block.String(), nil
}

func (b Boolean) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(b.Name.Value)
}

func (b Boolean) PythonClassType() string {
	return "bool"
}
