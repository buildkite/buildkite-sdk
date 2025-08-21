package types

import (
	"strings"

	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

type PropertyReference struct {
	Name string
	Ref  string
	Type Value
}

func (p PropertyReference) IsReference() bool {
	return true
}

func (p PropertyReference) IsNested() bool {
	parts := strings.Split(p.Ref, "/")
	return len(parts) > 3
}

func (p PropertyReference) Go() (string, error) {
	return utils.CamelCaseToTitleCase(p.Name), nil
}

func (p PropertyReference) GoStructType() string {
	return utils.CamelCaseToTitleCase(p.Name)
}

func (p PropertyReference) GoStructKey(isUnion bool) string {
	return utils.CamelCaseToTitleCase(p.Name)
}
