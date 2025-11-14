package gogen

import (
	"fmt"
	"sort"

	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
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
	Items       *orderedmap.OrderedMap
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

func (g GoStruct) Write() string {
	contents := utils.NewCodeBlock()

	if g.Description != "" {
		contents.AddLines(NewGoComment(g.Description))
	}

	contents.AddLines(fmt.Sprintf("type %s struct {", g.Name))

	g.Items.SortKeys(sort.Strings)
	keys := g.Items.Keys()
	for _, key := range keys {
		val, _ := g.Items.Get(key)
		item := val.(GoStructItem)

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
		Items:       orderedmap.New(),
	}
}
