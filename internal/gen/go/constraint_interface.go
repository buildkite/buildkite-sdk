package gogen

import (
	"fmt"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type GoConstraintInterface struct {
	Name        string
	Description string
	Items       []string
}

func (g *GoConstraintInterface) AddItem(item string) {
	g.Items = append(g.Items, item)
}

func (g *GoConstraintInterface) Write() string {
	contents := utils.NewCodeBlock()

	if g.Description != "" {
		contents.AddLines(NewGoComment(g.Description))
	}

	contents.AddLines(
		fmt.Sprintf("type %s interface {", g.Name),
		fmt.Sprintf("    %s", strings.Join(g.Items, " | ")),
		"}",
	)

	return contents.String()
}

func NewGoConstraintInterface(name, description string) *GoConstraintInterface {
	return &GoConstraintInterface{
		Name:        name,
		Description: description,
	}
}
