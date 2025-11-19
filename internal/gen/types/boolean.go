package types

import (
	"fmt"

	gogen "github.com/buildkite/buildkite-sdk/internal/gen/go"
	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
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

func (Boolean) IsPrimitive() bool {
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
	typ := gogen.NewType(
		b.Name.ToTitleCase(),
		b.Description,
		b.GoStructType(),
	)
	return typ.String(), nil
}

// TypeScript
func (b Boolean) TypeScript() string {
	typ := typescript.NewType(
		b.Name.ToTitleCase(),
		b.Description,
		"boolean",
	)
	return typ.String()
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

	block.AddLines(fmt.Sprintf("%s = bool", b.Name.ToTitleCase()))
	return block.String(), nil
}

func (b Boolean) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(b.Name.Value)
}

func (b Boolean) PythonClassType() string {
	return "bool"
}
