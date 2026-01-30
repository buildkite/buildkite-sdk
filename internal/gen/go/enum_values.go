package gogen

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type GoEnumValues struct {
	EnumName    string
	Description string
	Values      []any
}

func (g GoEnumValues) Write() string {
	lines := utils.NewCodeBlock(
		fmt.Sprintf("type %s string", g.EnumName),
	)

	if g.Description != "" {
		lines.AddLines(NewGoComment(g.Description))
	}

	lines.AddLines(
		fmt.Sprintf("var %sValues = map[string]%s{", g.EnumName, g.EnumName),
	)

	for _, val := range g.Values {
		lines.AddLinesWithIndent(
			4,
			fmt.Sprintf("\"%s\": \"%v\",", val, val),
		)
	}

	lines.AddLines("}")
	return lines.String()
}

func NewGoEnumValues(name string, values []any) GoEnumValues {
	return GoEnumValues{
		EnumName: name,
		Values:   values,
	}
}
