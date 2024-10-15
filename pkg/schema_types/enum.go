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

func (SchemaEnum) IsUnion() bool {
	return false
}

func (s SchemaEnum) TypeScriptType() string {
	enumProps := utils.CodeGen.NewCodeBlock()
	for _, val := range s.Values {
		enumProps = append(enumProps, fmt.Sprintf("    %s = \"%s\",", strings.ToUpper(val), val))
	}

	return utils.CodeGen.NewCodeBlock(
		fmt.Sprintf("export enum %s {", utils.String.SnakeCaseToTitleCase(s.Name)),
		enumProps.Display(),
		"}\n",
	).Display()
}

func (s SchemaEnum) GoType() string {
	name := utils.String.SnakeCaseToTitleCase(s.Name)
	consts := utils.CodeGen.NewCodeBlock()
	for _, val := range s.Values {
		consts = append(consts, fmt.Sprintf("    %s %s = \"%s\"", strings.ToUpper(val), name, val))
	}

	return utils.CodeGen.NewCodeBlock(
		fmt.Sprintf("type %s string", name),
		"const (",
		consts.Display(),
		")",
	).Display()
}

type enum struct{}

func (enum) String(name string, values []string) SchemaEnum {
	return SchemaEnum{
		Name:   name,
		Values: values,
	}
}

var Enum = enum{}
