package types

import (
	"fmt"
	"slices"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
)

type Object struct {
	Name                 PropertyName
	Description          string
	Properties           *orderedmap.OrderedMap
	AdditionalProperties *Value
	Required             []string
}

func (o Object) GetDescription() string {
	return o.Description
}

func (o Object) IsReference() bool {
	return false
}

func (Object) IsPrimative() bool {
	return false
}

// Go
func (o Object) GoStructType() string {
	return o.Name.ToTitleCase()
}

func (o Object) GoStructKey(isUnion bool) string {
	return o.Name.ToTitleCase()
}

func (o Object) Go() (string, error) {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		block := utils.NewCodeBlock()

		if o.Description != "" {
			block.AddLines(fmt.Sprintf("// %s", o.Description))
		}

		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			block.AddLines(fmt.Sprintf("type %s = map[string]%s", o.Name.ToTitleCase(), prop.GoStructType()))
			return block.String(), nil
		}

		block.AddLines(fmt.Sprintf("type %s = map[string]interface{}", o.Name.ToTitleCase()))
		return block.String(), nil
	}

	codeBlock := utils.NewCodeBlock()

	objectStruct := utils.NewGoStruct(o.Name.ToTitleCase(), o.Description, nil)
	for _, name := range keys {
		prop, _ := o.Properties.Get(name)
		val := prop.(Value)

		structKey := val.GoStructKey(false)
		structType := val.GoStructType()
		description := val.GetDescription()
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
				Description:          obj.Description,
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

		objectStruct.AddItem(structKey, structType, name, description, isPointer)
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

// TypeScript
func (o Object) TypeScript() (string, error) {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		block := utils.NewCodeBlock()

		if o.Description != "" {
			block.AddLines(utils.NewTypeDocComment(o.Description))
		}

		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			block.AddLines(fmt.Sprintf("export type %s = Record<string, %s>", o.Name.ToTitleCase(), prop.TypeScriptInterfaceType()))
			return block.String(), nil
		}

		block.AddLines(fmt.Sprintf("export type %s = Record<string, any>", o.Name.ToTitleCase()))
		return block.String(), nil
	}

	tsInterface := utils.NewTypeScriptInterface(o.Name.ToTitleCase(), o.Description)
	for _, name := range keys {
		prop, _ := o.Properties.Get(name)
		val := prop.(Value)

		structType := val.TypeScriptInterfaceType()
		required := slices.Contains(o.Required, name)

		// Property Reference
		if ref, ok := val.(PropertyReference); ok {
			switch ref.Type.(type) {
			case String:
				tsInterface.AddItem(name, "string", val.GetDescription(), required)
				continue
			case Number:
				tsInterface.AddItem(name, "number", val.GetDescription(), required)
				continue
			case Boolean:
				tsInterface.AddItem(name, "boolean", val.GetDescription(), required)
				continue
			default:
				tsInterface.AddItem(name, utils.CamelCaseToTitleCase(ref.Ref.Name()), val.GetDescription(), required)
				continue
			}
		}

		// Nested Object
		if obj, ok := val.(Object); ok {
			keys := obj.Properties.Keys()
			if len(keys) == 0 {
				tsInterface.AddItem(name, "Record<string,any>", obj.Description, required)
				continue
			}

			tsObject := utils.NewTypeScriptInterface("", obj.Description)
			for _, propName := range keys {
				nestedProp, _ := obj.Properties.Get(propName)
				nestedVal := nestedProp.(Value)
				nestedDescription := nestedVal.GetDescription()
				nestedRequired := slices.Contains(obj.Required, propName)

				if ref, ok := nestedVal.(PropertyReference); ok {
					switch ref.Type.(type) {
					case String:
						tsObject.AddItem(propName, "string", nestedDescription, required)
						continue
					case Number:
						tsObject.AddItem(propName, "number", nestedDescription, required)
						continue
					case Boolean:
						tsObject.AddItem(propName, "boolean", nestedDescription, required)
						continue
					default:
						tsObject.AddItem(propName, utils.CamelCaseToTitleCase(ref.Ref.Name()), nestedDescription, required)
						continue
					}
				}
				tsObject.AddItem(propName, nestedVal.TypeScriptInterfaceType(), nestedDescription, nestedRequired)
			}

			objString, err := tsObject.WriteUnionObject()
			if err != nil {
				return "", fmt.Errorf("generating nested object: %v", err)
			}

			tsInterface.AddItem(name, objString, obj.Description, required)
			continue
		}

		tsInterface.AddItem(name, structType, val.GetDescription(), required)
	}

	tsInterfaceString, err := tsInterface.Write()
	if err != nil {
		return "", fmt.Errorf("writing ts interface: %v", err)
	}

	block := utils.NewCodeBlock(
		tsInterfaceString,
	)
	return block.String(), nil
}

func (o Object) TypeScriptInterfaceKey() string {
	return o.Name.Value
}

func (o Object) TypeScriptInterfaceType() string {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		return "Record<string, any>"
	}

	return o.Name.ToTitleCase()
}

func (o Object) TypeScriptImports() string {
	var imports []string
	for _, key := range o.Properties.Keys() {
		prop, _ := o.Properties.Get(key)
		val := prop.(Value)

		if union, ok := val.(Union); ok {
			imports = append(imports, union.TypeScriptImports())
		}

		if ref, ok := val.(PropertyReference); ok {
			imports = append(imports,
				fmt.Sprintf("import {%s} from \"./%s.ts\"", ref.TypeScriptInterfaceType(), ref.Name),
			)
		}
	}
	return strings.Join(imports, "\n")
}

// Python
func (o Object) Python() (string, error) {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		block := utils.NewCodeBlock()

		if o.Description != "" {
			block.AddLines(fmt.Sprintf("# %s", o.Description))
		}

		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			block.AddLines(fmt.Sprintf("%s = Dict[str, %s]", o.Name.ToTitleCase(), prop.PythonClassType()))
			return block.String(), nil
		}

		block.AddLines(fmt.Sprintf("%s = Dict[str, Any]", o.Name.ToTitleCase()))
		return block.String(), nil
	}

	codeBlock := utils.NewCodeBlock()
	pyClass := utils.NewPythonClass(o.Name.ToTitleCase(), o.Description)
	pyTypedDict := utils.NewPythonClass(fmt.Sprintf("%sArgs", o.Name.ToTitleCase()), o.Description)

	for _, name := range keys {
		prop, _ := o.Properties.Get(name)
		val := prop.(Value)

		// Reserved words
		if name == "async" {
			pyTypedDict.AddItem("async", "Literal[True, False, 'true', 'false']", "", "async", val.GetDescription(), false, false)
			pyClass.AddItem("step_async", "Literal[True, False, 'true', 'false']", "", "async", val.GetDescription(), false, false)
			continue
		}

		if name == "if" {
			pyTypedDict.AddItem("if", "If", "", "if", val.GetDescription(), false, false)
			pyClass.AddItem("step_if", "If", "", "if", val.GetDescription(), false, false)
			continue
		}

		if name == "with" {
			pyTypedDict.AddItem("with", "MatrixElementList | MatrixAdjustmentsWithObject", "", "with", val.GetDescription(), true, false)
			pyClass.AddItem("matrix_with", "MatrixElementList | MatrixAdjustmentsWithObject", "", "with", val.GetDescription(), true, false)
			continue
		}

		structType := val.PythonClassType()
		dictStructType := structType
		constructorName := ""
		description := val.GetDescription()
		isObjectArray := false
		required := slices.Contains(o.Required, name)

		// Object
		if obj, ok := val.(Object); ok {
			keys := obj.Properties.Keys()
			if len(keys) == 0 {
				pyTypedDict.AddItem(name, "Dict[str, Any]", "", "", description, required, isObjectArray)
				pyClass.AddItem(name, "Dict[str, Any]", "", "", description, required, isObjectArray)
				continue
			}

			dictStructType = fmt.Sprintf("%sArgs", dictStructType)
			constructorName = structType
			nestedPyClass := utils.NewPythonClass(structType, description)
			nestedPyTypeDict := utils.NewPythonClass(fmt.Sprintf("%sArgs", structType), description)
			for _, propName := range keys {
				nestedProp, _ := obj.Properties.Get(propName)
				nestedVal := nestedProp.(Value)
				nestedType := nestedVal.PythonClassType()
				nestedRequired := slices.Contains(obj.Required, propName)

				nestedDictType := nestedType
				if obj, ok := nestedVal.(Object); ok {
					if len(obj.Properties.Keys()) > 0 {
						nestedDictType = fmt.Sprintf("%sArgs", nestedType)
					}
				}

				nestedPyTypeDict.AddItem(propName, nestedDictType, "", "", nestedVal.GetDescription(), nestedRequired, false)
				nestedPyClass.AddItem(propName, nestedType, "", "", nestedVal.GetDescription(), nestedRequired, false)
			}

			nestedObjectClass, err := nestedPyClass.Write()
			if err != nil {
				return "", fmt.Errorf("writing nested class [%s]: %v", o.Name.Value, err)
			}
			codeBlock.AddLines(nestedObjectClass)

			nestedObjectTypedDict, err := nestedPyTypeDict.WriteTypedDict()
			if err != nil {
				return "", fmt.Errorf("writing nested class [%s]: %v", o.Name.Value, err)
			}
			codeBlock.AddLines(nestedObjectTypedDict)
		}

		// PropertyReference
		if ref, ok := val.(PropertyReference); ok {
			if obj, ok := ref.Type.(Object); ok {
				if len(obj.Properties.Keys()) > 0 {
					dictStructType = fmt.Sprintf("%sArgs", structType)
					constructorName = structType
				}
			}
		}

		// Array
		if array, ok := val.(Array); ok {
			if obj, ok := array.Type.(Object); ok {
				if len(obj.Properties.Keys()) > 0 {
					dictStructType = fmt.Sprintf("List[%sArgs]", obj.PythonClassType())
					constructorName = obj.PythonClassType()
					isObjectArray = true
				}
			}
		}

		// Union
		if union, ok := val.(Union); ok {
			typeIdentifiers := make([]Value, len(union.TypeIdentifiers))
			for i, typ := range union.TypeIdentifiers {
				if obj, ok := typ.(Object); ok {
					typeIdentifiers[i] = Object{
						Name:                 NewPropertyName(fmt.Sprintf("%sArgs", obj.Name.Value)),
						Properties:           obj.Properties,
						AdditionalProperties: obj.AdditionalProperties,
						Required:             obj.Required,
					}
					continue
				}

				if ref, ok := typ.(PropertyReference); ok {
					typeIdentifiers[i] = PropertyReference{
						Name: fmt.Sprintf("%sArgs", ref.Name),
						Ref:  schema.PropertyReferenceString(fmt.Sprintf("%sArgs", ref.Name)),
						Type: ref.Type,
					}
					continue
				}

				typeIdentifiers[i] = typ
			}

			newUnion := Union{
				Name:            union.Name,
				Description:     union.Description,
				TypeIdentifiers: typeIdentifiers,
			}

			dictStructType = newUnion.PythonClassType()
		}

		pyTypedDict.AddItem(name, dictStructType, "", "", description, required, isObjectArray)
		pyClass.AddItem(name, structType, constructorName, "", description, required, isObjectArray)
	}

	pyTypedDictString, err := pyTypedDict.WriteTypedDict()
	if err != nil {
		return "", fmt.Errorf("writing python typed dict: %v", err)
	}
	codeBlock.AddLines(pyTypedDictString, "")

	pyClassString, err := pyClass.Write()
	if err != nil {
		return "", fmt.Errorf("writing python class: %v", err)
	}
	codeBlock.AddLines(pyClassString)
	return codeBlock.String(), nil
}

func (o Object) PythonClassKey() string {
	return utils.CamelCaseToSnakeCase(o.Name.Value)
}

func (o Object) PythonClassType() string {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			return fmt.Sprintf("Dict[str, %s]", prop.PythonClassType())
		}

		return "Dict[str, Any]"
	}

	return o.Name.ToTitleCase()
}
