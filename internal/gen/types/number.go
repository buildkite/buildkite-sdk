package types

import (
	"fmt"

	gogen "github.com/buildkite/buildkite-sdk/internal/gen/go"
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

func (Number) IsPrimitive() bool {
	return true
}

// Go
func (n Number) Go() (string, error) {
	typ := gogen.NewType(
		n.Name.ToTitleCase(),
		n.Description,
		"int",
	)
	return typ.String(), nil
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
func (n Number) TypeScript() string {
	typ := typescript.NewType(
		n.Name.ToTitleCase(),
		n.Description,
		"number",
	)
	return typ.String()
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
