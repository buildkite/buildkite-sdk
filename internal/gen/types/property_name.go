package types

import (
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type PropertyName struct {
	Value string
}

func (s PropertyName) ToTitleCase() string {
	if strings.Contains(s.Value, "_") {
		return utils.DashCaseToTitleCase(s.Value)
	}

	return utils.CamelCaseToTitleCase(s.Value)
}

func NewPropertyName(name string) PropertyName {
	return PropertyName{
		Value: name,
	}
}
