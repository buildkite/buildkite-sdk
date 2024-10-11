package schema_types

import (
	"fmt"
	"strings"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type SchemaEnum struct {
	Name   string
	Values []string
}

func (s SchemaEnum) TypeScriptType() string {
	enumProps := utils.CodeBlock{}
	for _, val := range s.Values {
		enumProps = append(enumProps, fmt.Sprintf("    %s = \"%s\",", strings.ToUpper(val), val))
	}

	return utils.CodeBlock{
		fmt.Sprintf("export enum %s {", utils.SnakeCaseToTitleCase(s.Name)),
		enumProps.Display(),
		"}\n",
	}.Display()
}

func (s SchemaEnum) GoType() string {
	consts := utils.CodeBlock{}
	for _, val := range s.Values {
		consts = append(consts, fmt.Sprintf("    %s %s = \"%s\"", strings.ToUpper(val), s.Name, val))
	}

	return utils.CodeBlock{
		fmt.Sprintf("type %s string", s.Name),
		"const (",
		consts.Display(),
		")",
	}.Display()
}
