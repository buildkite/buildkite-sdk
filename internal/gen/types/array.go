package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

type Array struct {
	Name      PropertyName
	Type      Value
	Reference bool
}

func (a Array) IsReference() bool {
	return a.Reference
}

func (Array) IsPrimative() bool {
	return false
}

func (a Array) GoStructType() string {
	if a.IsReference() {
		return fmt.Sprintf("[]%s", a.Type.GoStructType())
	}

	switch a.Type.(type) {
	case String:
		return "[]string"
	case Boolean:
		return "[]bool"
	case Number:
		return "[]int"
	case Union:
		return fmt.Sprintf("[]%sUnion", a.Name.ToTitleCase())
	default:
		return fmt.Sprintf("[]%s", a.Name.ToTitleCase())
	}
}

func (a Array) GoStructKey(isUnion bool) string {
	if isUnion {
		switch a.Type.(type) {
		case String:
			return "StringArray"
		case Boolean:
			return "BoolArray"
		case Number:
			return "IntArray"
		case Union:
			return fmt.Sprintf("%sUnion", a.Name.ToTitleCase())
		default:
			return a.Name.ToTitleCase()
		}
	}

	return a.Name.ToTitleCase()
}

func (a Array) Go() (string, error) {
	union, ok := a.Type.(Union)
	if !a.IsReference() && ok {
		item := Union{
			Name:            NewPropertyName(fmt.Sprintf("%sUnion", a.Name.Value)),
			Description:     union.Description,
			TypeIdentifiers: union.TypeIdentifiers,
		}

		lines := utils.NewCodeBlock()
		itemLines, err := item.Go()
		if err != nil {
			return "", fmt.Errorf("generating lines for union in array [%s]: %v", a.Name.Value, err)
		}

		lines.AddLines(
			itemLines,
			fmt.Sprintf("type %s = []%sUnion", a.Name.ToTitleCase(), a.Type.GoStructType()),
		)
		return lines.String(), nil
	}

	return fmt.Sprintf("type %s = []%s", a.Name.ToTitleCase(), a.Type.GoStructType()), nil
}
