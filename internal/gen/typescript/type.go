package typescript

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type typescriptType struct {
	Name        string
	Description string
	Value       string
}

func (t typescriptType) String() string {
	block := utils.NewCodeBlock()

	if t.Description != "" {
		block.AddLines(NewTypeDocComment(t.Description))
	}

	block.AddLines(
		fmt.Sprintf("export type %s = %s", t.Name, t.Value),
	)

	return block.String()
}

func NewType(name, description, value string) typescriptType {
	return typescriptType{
		Name:        name,
		Description: description,
		Value:       value,
	}
}
