package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
)

type Object struct {
	Name                 PropertyName
	Properties           *orderedmap.OrderedMap
	AdditionalProperties *Value
}

func (o Object) IsReference() bool {
	return false
}

func (Object) IsPrimative() bool {
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
	keys := o.Properties.Keys()

	if len(keys) == 0 {
		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			return fmt.Sprintf("type %s = map[string]%s", o.Name.ToTitleCase(), prop.GoStructType()), nil
		}

		return fmt.Sprintf("type %s = map[string]string", o.Name.ToTitleCase()), nil
	}

	codeBlock := utils.NewCodeBlock()

	objectStruct := utils.NewGoStruct(o.Name.ToTitleCase(), nil)
	for _, name := range keys {
		prop, _ := o.Properties.Get(name)
		val := prop.(Value)

		structKey := val.GoStructKey(false)
		structType := val.GoStructType()
		isPointer := true

		// Array
		if _, ok := val.(Array); ok {
			isPointer = false
			structKey = utils.DashCaseToTitleCase(name)
		}

		// Object
		if obj, ok := val.(Object); ok {
			structKey = utils.DashCaseToTitleCase(name)
			nestedObjName := NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), structKey))
			nestedObj := Object{
				Name:                 nestedObjName,
				Properties:           obj.Properties,
				AdditionalProperties: obj.AdditionalProperties,
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
			structKey = utils.DashCaseToTitleCase(name)
			nestedEnum := Enum{
				Name:        NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), structKey)),
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
				Name:            NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), utils.DashCaseToTitleCase(name))),
				Description:     union.Description,
				TypeIdentifiers: union.TypeIdentifiers,
			}

			unionLines, err := nestedUnion.Go()
			if err != nil {
				return "", fmt.Errorf("generating union lines for struct [%s]: %v", o.Name.Value, err)
			}

			structType = nestedUnion.GoStructType()
			structKey = utils.DashCaseToTitleCase(name)
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
