package types

import (
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type PropertyReference struct {
	Name string
	Ref  schema.PropertyReferenceString
	Type Value
}

func (p PropertyReference) GetDescription() string {
	return p.Type.GetDescription()
}

func (p PropertyReference) IsReference() bool {
	return true
}

func (PropertyReference) IsPrimitive() bool {
	return false
}

func (p PropertyReference) IsNested() bool {
	parts := strings.Split(string(p.Ref), "/")
	return len(parts) > 3
}

// isPrimitiveType checks if the referenced type is a primitive type
func (p PropertyReference) isPrimitiveType() bool {
	switch p.Type.(type) {
	case String, Number, Boolean:
		return true
	default:
		return false
	}
}

// Go
func (p PropertyReference) Go() (string, error) {
	return utils.CamelCaseToTitleCase(p.Name), nil
}

func (p PropertyReference) GoStructType() string {
	if p.isPrimitiveType() {
		return p.Type.GoStructType()
	}

	if p.Type != nil {
		return utils.CamelCaseToTitleCase(p.Ref.Name())
	}

	return utils.CamelCaseToTitleCase(p.Name)
}

func (p PropertyReference) GoStructKey(isUnion bool) string {
	if strings.Contains(p.Name, "_") {
		return utils.DashCaseToTitleCase(p.Name)
	}

	return utils.CamelCaseToTitleCase(p.Name)
}

// TypeScript
func (p PropertyReference) TypeScript() string {
	return p.Name
}

func (p PropertyReference) TypeScriptInterfaceKey() string {
	return p.Name
}

func (p PropertyReference) TypeScriptInterfaceType() string {
	if p.isPrimitiveType() {
		return p.Type.TypeScriptInterfaceType()
	}

	name := p.Ref.Name()
	if strings.Contains(name, "_") {
		return utils.DashCaseToTitleCase(name)
	}

	return utils.CamelCaseToTitleCase(name)
}

// Python
func (p PropertyReference) Python() (string, error) {
	return "", nil
}

func (p PropertyReference) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(p.Name)
}

func (p PropertyReference) PythonClassType() string {
	if p.isPrimitiveType() {
		return p.Type.PythonClassType()
	}

	return utils.CamelCaseToTitleCase(p.Ref.Name())
}

// CSharp
func (p PropertyReference) CSharp() (string, error) {
	return "", nil
}

func (p PropertyReference) CSharpType() string {
	if p.isPrimitiveType() {
		return p.Type.CSharpType()
	}

	name := p.Ref.Name()
	if strings.Contains(name, "_") {
		return utils.DashCaseToTitleCase(name)
	}

	return utils.CamelCaseToTitleCase(name)
}
