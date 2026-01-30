package types

import (
	"fmt"
	"strings"

	gogen "github.com/buildkite/buildkite-sdk/internal/gen/go"
	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Union struct {
	Name            PropertyName
	Description     string
	TypeIdentifiers []Value
}

func (u Union) GetDescription() string {
	return u.Description
}

func (u Union) IsReference() bool {
	return false
}

func (u Union) IsPrimitive() bool {
	return false
}

// Go
func (u Union) GoStructKey(isUnion bool) string {
	return u.Name.ToTitleCase()
}

func (u Union) GoStructType() string {
	return u.Name.ToTitleCase()
}

func (u Union) Go() (string, error) {
	lines := utils.NewCodeBlock()
	unionDefinition := gogen.NewGoStruct(u.Name.ToTitleCase(), u.Description, nil)

	for _, typ := range u.TypeIdentifiers {
		isPointer := true
		structKey := typ.GoStructKey(true)
		structType := typ.GoStructType()
		structTag := ""

		var extraDefinition Value
		switch val := typ.(type) {
		case Enum:
			extraDefinition = typ
		case Object:
			extraDefinition = Object{
				Name:                 NewPropertyName(fmt.Sprintf("%sObject", val.Name.Value)),
				Properties:           val.Properties,
				AdditionalProperties: val.AdditionalProperties,
			}
			structType = extraDefinition.GoStructType()
		case Array:
			isPointer = false
			if !isPrimitiveValue(val.Type) {
				extraDefinition = val
			}
		}

		if extraDefinition != nil {
			extraLines, err := extraDefinition.Go()
			if err != nil {
				return "", err
			}
			lines.AddLines(extraLines)
		}

		unionDefinition.AddItem(structKey, structType, structTag, typ.GetDescription(), isPointer)
	}

	lines.AddLines(
		unionDefinition.WriteConstraintInterface(),
		unionDefinition.Write(),
		unionDefinition.WriteMarshalFunction(),
	)

	return lines.String(), nil
}

// TypeScript
func (u Union) TypeScriptInterfaceKey() string {
	return u.Name.Value
}

func (u Union) TypeScriptInterfaceType() string {
	parts := make([]string, len(u.TypeIdentifiers))
	for i, typ := range u.TypeIdentifiers {
		// Object
		if obj, ok := typ.(Object); ok {
			parts[i] = obj.TypeScript()
			continue
		}

		parts[i] = typ.TypeScriptInterfaceType()
	}
	return strings.Join(parts, " | ")
}

func (u Union) TypeScript() string {
	typ := typescript.NewType(
		u.Name.ToTitleCase(),
		u.Description,
		u.TypeScriptInterfaceType(),
	)
	return typ.String()
}

// Python
func (u Union) Python() (string, error) {
	codeBlock := utils.NewCodeBlock()

	var parts []string
	for _, typ := range u.TypeIdentifiers {
		// Nested Object
		if obj, ok := typ.(Object); ok {
			nestedObj := Object{
				Name:                 NewPropertyName(fmt.Sprintf("%sObject", obj.Name.Value)),
				Description:          obj.Description,
				Properties:           obj.Properties,
				AdditionalProperties: obj.AdditionalProperties,
			}

			objLines, err := nestedObj.Python()
			if err != nil {
				return "", fmt.Errorf("generating object lines for union [%s]: %v", u.Name.Value, err)
			}

			codeBlock.AddLines(objLines)
			pythonClassType := nestedObj.PythonClassType()
			parts = append(parts, pythonClassType)

			if !strings.HasPrefix(pythonClassType, "Dict") {
				parts = append(parts, fmt.Sprintf("%sArgs", nestedObj.PythonClassType()))
			}

			continue
		}

		if ref, ok := typ.(PropertyReference); ok {
			if obj, ok := ref.Type.(Object); ok {
				if len(obj.Properties.Keys()) > 0 {
					parts = append(parts, fmt.Sprintf("%sArgs", typ.PythonClassType()))
				}
			}
		}

		parts = append(parts, typ.PythonClassType())
	}

	if u.Description != "" {
		codeBlock.AddLines(fmt.Sprintf("# %s", u.Description))
	}

	codeBlock.AddLines(fmt.Sprintf("%s = %s", u.Name.ToTitleCase(), strings.Join(parts, " | ")))
	return codeBlock.String(), nil
}

func (u Union) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(u.Name.Value)
}

func (u Union) PythonClassType() string {
	codeBlock := utils.NewCodeBlock()
	parts := make([]string, len(u.TypeIdentifiers))
	for i, typ := range u.TypeIdentifiers {
		parts[i] = typ.PythonClassType()
	}
	codeBlock.AddLines(strings.Join(parts, " | "))
	return codeBlock.String()
}
