package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

type Object struct {
	Name       PropertyName
	Properties map[string]Value
}

func (o Object) IsReference() bool {
	return false
}

func (o Object) GoStructType() string {
	return o.Name.ToTitleCase()
}

func (o Object) GoStructKey(isUnion bool) string {
	return o.Name.ToTitleCase()
}

func (o Object) Go() (string, error) {
	// TODO: support other map types
	if len(o.Properties) == 0 {
		return fmt.Sprintf("type %s = map[string]string", o.Name.ToTitleCase()), nil
	}

	codeBlock := utils.NewCodeBlock()

	objectStruct := utils.NewGoStruct(o.Name.ToTitleCase(), nil)
	for name, val := range o.Properties {
		structKey := val.GoStructKey(false)
		structType := val.GoStructType()
		isPointer := false

		// Object
		if obj, ok := val.(Object); ok {
			nestedObjName := NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), obj.Name.ToTitleCase()))
			nestedObj := Object{
				Name:       nestedObjName,
				Properties: obj.Properties,
			}

			objLines, err := nestedObj.Go()
			if err != nil {
				return "", fmt.Errorf("generating nested object for [%s]: %v", o.Name.Value, err)
			}

			structType = nestedObjName.ToTitleCase()
			codeBlock.AddLines(objLines)
		}

		// Enum
		if enum, ok := val.(Enum); ok {
			nestedEnum := Enum{
				Name:        NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), enum.Name.ToTitleCase())),
				Description: enum.Description,
				Values:      enum.Values,
				Default:     enum.Default,
			}

			enumLines, err := nestedEnum.Go()
			if err != nil {
				return "", fmt.Errorf("generating enum lines for struct [%s]: %v", o.Name.Value, err)
			}

			structType = nestedEnum.GoStructType()
			codeBlock.AddLines(enumLines)
		}

		// Union
		if union, ok := val.(Union); ok {
			nestedUnion := Union{
				Name:            NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), union.Name.ToTitleCase())),
				Description:     union.Description,
				TypeIdentifiers: union.TypeIdentifiers,
			}

			unionLines, err := nestedUnion.Go()
			if err != nil {
				return "", fmt.Errorf("generating union lines for struct [%s]: %v", o.Name.Value, err)
			}

			structType = nestedUnion.GoStructType()
			codeBlock.AddLines(unionLines)
		}

		objectStruct.AddItem(structKey, structType, name, isPointer)
	}

	structLines, err := objectStruct.Write()
	if err != nil {
		return "", fmt.Errorf("writing out object struct [%s]: %v", objectStruct.Name, err)
	}

	codeBlock.AddLines(
		structLines,
	)

	return codeBlock.String(), nil
}
