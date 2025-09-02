package types

import (
	"strings"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

type PropertyReference struct {
	Name string
	Ref  schema.PropertyReferenceString
	Type Value
}

func (p PropertyReference) IsReference() bool {
	return true
}

func (PropertyReference) IsPrimative() bool {
	return false
}

func (p PropertyReference) IsNested() bool {
	parts := strings.Split(string(p.Ref), "/")
	return len(parts) > 3
}

func (p PropertyReference) Go() (string, error) {
	return utils.CamelCaseToTitleCase(p.Name), nil
}

func (p PropertyReference) GoStructType() string {
	switch p.Type.(type) {
	case String:
		return p.Type.GoStructType()
	case Number:
		return p.Type.GoStructType()
	case Boolean:
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
