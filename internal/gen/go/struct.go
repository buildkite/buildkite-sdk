package gogen

import (
	"fmt"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type GoStructItem struct {
	Name        string
	Description string
	Value       string
	Pointer     bool
	TagName     string
}

type GoStruct struct {
	Name        string
	Description string
	Items       *utils.OrderedMap[GoStructItem]
}

func (g *GoStruct) AddItem(key, val, tagName, description string, isPointer bool) {
	g.Items.Set(key, GoStructItem{
		Name:        key,
		Description: description,
		Value:       val,
		Pointer:     isPointer,
		TagName:     tagName,
	})
}

func (g GoStruct) WriteConstraintInterface() string {
	contents := utils.NewCodeBlock()

	keys := g.Items.Keys()
	items := make([]string, len(keys))
	for i, key := range keys {
		item, _ := g.Items.Get(key)
		items[i] = item.Value
	}

	contents.AddLines(
		fmt.Sprintf("type %sValues interface {", g.Name),
		fmt.Sprintf("    %s", strings.Join(items, " | ")),
		"}",
	)

	return contents.String()
}

func (g GoStruct) WriteMarshalFunction() string {
	contents := utils.NewCodeBlock(
		fmt.Sprintf("func (e %s) MarshalJSON() ([]byte, error) {", g.Name),
	)

	g.Items.SortKeys()
	for _, key := range g.Items.Keys() {
		item, _ := g.Items.Get(key)
		contents.AddLinesWithIndent(4, fmt.Sprintf("if e.%s != nil {", item.Name))
		contents.AddLinesWithIndent(8, fmt.Sprintf("return json.Marshal(e.%s)", item.Name))
		contents.AddLinesWithIndent(4, "}")
	}

	contents.AddLinesWithIndent(4, "return json.Marshal(nil)")
	contents.AddLines("}")
	return contents.String()
}

func (g GoStruct) Write() string {
	contents := utils.NewCodeBlock()

	if g.Description != "" {
		contents.AddLines(NewGoComment(g.Description))
	}

	contents.AddLines(fmt.Sprintf("type %s struct {", g.Name))

	g.Items.SortKeys()
	for _, key := range g.Items.Keys() {
		item, _ := g.Items.Get(key)

		if item.Description != "" {
			contents.AddLines(fmt.Sprintf("    %s", NewGoComment(item.Description)))
		}

		pointer := ""
		if item.Pointer {
			pointer = "*"
		}

		tag := ""
		if item.TagName != "" {
			tag = fmt.Sprintf(" `json:\"%s,omitempty\"`", item.TagName)
		}

		contents.AddLines(fmt.Sprintf("    %s %s%s%s", item.Name, pointer, item.Value, tag))
	}

	contents.AddLines("}")
	return contents.String()
}

func NewGoStruct(name, description string, items []GoStructItem) *GoStruct {
	return &GoStruct{
		Name:        name,
		Description: description,
		Items:       utils.NewOrderedMap[GoStructItem](nil),
	}
}
