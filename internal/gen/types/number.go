package types

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Number struct {
	Name        PropertyName
	Description string
}

func (n Number) GetDescription() string {
	return n.Description
}

func (Number) IsReference() bool {
	return false
}

func (Number) IsPrimative() bool {
	return true
}

// Go
func (n Number) Go() (string, error) {
	block := utils.NewCodeBlock()
	if n.Description != "" {
		block.AddLines(fmt.Sprintf("// %s", n.Description))
	}

	block.AddLines(fmt.Sprintf("type %s = int", n.Name.ToTitleCase()))
	return block.String(), nil
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
	typ := typescript.NewType(
		n.Name.ToTitleCase(),
		n.Description,
		"number",
	)
	return typ.String(), nil
}

func (n Number) TypeScriptInterfaceKey() string {
	return n.Name.Value
}

func (n Number) TypeScriptInterfaceType() string {
	return "number"
}

// Python
func (n Number) Python() (string, error) {
	block := utils.NewCodeBlock()
	if n.Description != "" {
		block.AddLines(fmt.Sprintf("# %s", n.Description))
	}

	block.AddLines(fmt.Sprintf("%s = int", n.Name.ToTitleCase()))
	return block.String(), nil
}

func (n Number) PythonClassKey() string {
	return utils.CamelCaseToTitleCase(n.Name.Value)
}

func (n Number) PythonClassType() string {
	return "int"
}
