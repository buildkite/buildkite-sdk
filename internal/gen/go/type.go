package gogen

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type goType struct {
	Name        string
	Description string
	Value       string
}

func (g goType) String() string {
	block := utils.NewCodeBlock()

	if g.Description != "" {
		block.AddLines(NewGoComment(g.Description))
	}

	block.AddLines(
		fmt.Sprintf("type %s = %s", g.Name, g.Value),
	)

	return block.String()
}

func NewType(name, description, value string) goType {
	return goType{
		Name:        name,
		Description: description,
		Value:       value,
	}
}
