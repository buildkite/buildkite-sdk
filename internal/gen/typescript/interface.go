package typescript

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type TypeScriptInterfaceItem struct {
	Name        string
	Description string
	Value       string
	Required    bool
}

type TypeScriptInterface struct {
	Name        string
	Description string
	Inline      bool
	Items       []TypeScriptInterfaceItem
}

func (t *TypeScriptInterface) AddItem(name, value, description string, required bool) {
	t.Items = append(t.Items, TypeScriptInterfaceItem{
		Name:        name,
		Description: description,
		Value:       value,
		Required:    required,
	})
}

func (t *TypeScriptInterface) Write() string {
	// Create a new code block for the interface.
	codeBlock := utils.NewCodeBlock()

	// If the interface has a description add a
	// TypeDoc comment.
	comment := NewTypeDocComment(t.Description)
	if !t.Inline && comment != "" {
		codeBlock.AddLines(comment)
	}

	// Add the first line of the interface.
	firstLine := fmt.Sprintf("export interface %s {", t.Name)
	if t.Inline {
		firstLine = "{"
	}
	codeBlock.AddLines(firstLine)

	// Add each item to the interface.
	for _, item := range t.Items {
		// If the item has a description add a
		// TypeDoc comment.
		if item.Description != "" {
			codeBlock.AddLines(
				NewTypeDocCommentWithIndent(item.Description, 4),
			)
		}

		// Add the item to the interface code block.
		optional := "?"
		if item.Required {
			optional = ""
		}

		codeBlock.AddLinesWithIndent(
			4,
			fmt.Sprintf("%s%s: %v", item.Name, optional, item.Value),
		)
	}

	// Add a closing bracket
	codeBlock.AddLines("}")

	// Return the string
	return codeBlock.String()
}

func NewTypeScriptInterface(name, description string, inline bool) *TypeScriptInterface {
	return &TypeScriptInterface{
		Name:        name,
		Description: description,
		Inline:      inline,
	}
}
