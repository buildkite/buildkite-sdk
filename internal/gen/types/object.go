package types

import (
	"fmt"
	"slices"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type Object struct {
	Name                 PropertyName
	Description          string
	Properties           *utils.OrderedMap[Value]
	AdditionalProperties *Value
	Required             []string

	// Is the object nested in another definition
	IsNested bool
}

func (o Object) GetDescription() string {
	return o.Description
}

func (o Object) IsReference() bool {
	return false
}

func (Object) IsPrimitive() bool {
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
		val, err := o.Properties.Get(name)
		if err != nil {
			return "", err
		}

		structKey := val.GoStructKey(false)
		structType := val.GoStructType()
		description := val.GetDescription()
		isPointer := true

		// Array
		if array, ok := val.(Array); ok {
			isPointer = false
			structKey = utils.DashCaseToTitleCase(name)

			// Array of inline objects (e.g. items defined directly in-place
			// rather than via a $ref) have no separate named struct written
			// out elsewhere, so write one here instead of referencing a name
			// that was never defined. $ref'd arrays already have their own
			// named type written elsewhere, so leave those alone.
			if obj, ok := array.Type.(Object); ok && !array.IsReference() && len(obj.Properties.Keys()) > 0 {
				nestedObjName := NewPropertyName(fmt.Sprintf("%s%s", o.Name.ToTitleCase(), structKey))
				nestedObj := Object{
					Name:                 nestedObjName,
					Description:          obj.Description,
					Properties:           obj.Properties,
					AdditionalProperties: obj.AdditionalProperties,
					Required:             obj.Required,
				}

				objLines, err := nestedObj.Go()
				if err != nil {
					return "", fmt.Errorf("generating nested array object for [%s]: %v", o.Name.Value, err)
				}

				structType = fmt.Sprintf("[]%s", nestedObjName.ToTitleCase())
				codeBlock.AddLines(objLines)
			}
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
func (o Object) TypeScript() string {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		recordValue := "any"
		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			recordValue = prop.TypeScriptInterfaceType()
		}

		typeValue := fmt.Sprintf("Record<string, %s>", recordValue)
		if o.IsNested {
			return typeValue
		}

		typ := typescript.NewType(
			o.Name.ToTitleCase(),
			o.Description,
			typeValue,
		)
		return typ.String()
	}

	tsInterface := typescript.NewTypeScriptInterface(o.Name.ToTitleCase(), o.Description, o.IsNested)
	for _, name := range keys {
		val, _ := o.Properties.Get(name)
		structType := val.TypeScriptInterfaceType()
		required := slices.Contains(o.Required, name)
		tsInterface.AddItem(name, structType, val.GetDescription(), required)
	}

	return tsInterface.Write()
}

func (o Object) TypeScriptInterfaceKey() string {
	return o.Name.Value
}

func (o Object) TypeScriptInterfaceType() string {
	if o.IsNested {
		return o.TypeScript()
	}

	return o.Name.ToTitleCase()
}

// writeNestedPythonClasses returns the Python class/TypedDict source for a
// nested object, recursing into any of its properties that are themselves
// objects (or arrays of objects) so that deeply nested structures get every
// class they reference actually written out, not just referenced by name.
func writeNestedPythonClasses(structType, description string, properties *utils.OrderedMap[Value], required []string) (string, error) {
	var sb strings.Builder

	nestedPyClass := utils.NewPythonClass(structType, description)
	nestedPyTypeDict := utils.NewPythonClass(fmt.Sprintf("%sArgs", structType), description)

	for _, propName := range properties.Keys() {
		nestedVal, _ := properties.Get(propName)
		nestedType := nestedVal.PythonClassType()
		nestedRequired := slices.Contains(required, propName)
		nestedDictType := nestedType

		if obj, ok := nestedVal.(Object); ok {
			if len(obj.Properties.Keys()) > 0 {
				nestedDictType = fmt.Sprintf("%sArgs", nestedType)
				nested, err := writeNestedPythonClasses(obj.PythonClassType(), obj.Description, obj.Properties, obj.Required)
				if err != nil {
					return "", err
				}
				sb.WriteString(nested)
			}
		}

		if array, ok := nestedVal.(Array); ok && !array.IsReference() {
			if obj, ok := array.Type.(Object); ok {
				if len(obj.Properties.Keys()) > 0 {
					nestedDictType = fmt.Sprintf("List[%sArgs]", obj.PythonClassType())
					nested, err := writeNestedPythonClasses(obj.PythonClassType(), obj.Description, obj.Properties, obj.Required)
					if err != nil {
						return "", err
					}
					sb.WriteString(nested)
				}
			}
		}

		nestedPyTypeDict.AddItem(propName, nestedDictType, "", "", nestedVal.GetDescription(), nestedRequired, false)
		nestedPyClass.AddItem(propName, nestedType, "", "", nestedVal.GetDescription(), nestedRequired, false)
	}

	nestedObjectClass, err := nestedPyClass.Write()
	if err != nil {
		return "", fmt.Errorf("writing nested class [%s]: %v", structType, err)
	}
	sb.WriteString(nestedObjectClass)
	sb.WriteString("\n")

	nestedObjectTypedDict, err := nestedPyTypeDict.WriteTypedDict()
	if err != nil {
		return "", fmt.Errorf("writing nested class [%s]: %v", structType, err)
	}
	sb.WriteString(nestedObjectTypedDict)
	sb.WriteString("\n")

	return sb.String(), nil
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
		val, _ := o.Properties.Get(name)

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
			nested, err := writeNestedPythonClasses(structType, description, obj.Properties, obj.Required)
			if err != nil {
				return "", fmt.Errorf("writing nested class [%s]: %v", o.Name.Value, err)
			}
			codeBlock.AddLines(nested)
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

		if name == "artifact_paths" {
			structType = "str | Path | List[str | Path]"
			dictStructType = structType
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

// CSharp
func (o Object) CSharp() (string, error) {
	keys := o.Properties.Keys()
	if len(keys) == 0 {
		if o.AdditionalProperties != nil {
			prop := *o.AdditionalProperties
			return fmt.Sprintf("public class %s : Dictionary<string, %s>\n{\n    public %s() : base() { }\n    public %s(IDictionary<string, %s> dictionary) : base(dictionary) { }\n}\n",
				o.Name.ToTitleCase(), prop.CSharpType(), o.Name.ToTitleCase(), o.Name.ToTitleCase(), prop.CSharpType()), nil
		}
		return fmt.Sprintf("public class %s : Dictionary<string, object>\n{\n    public %s() : base() { }\n    public %s(IDictionary<string, object> dictionary) : base(dictionary) { }\n}\n",
			o.Name.ToTitleCase(), o.Name.ToTitleCase(), o.Name.ToTitleCase()), nil
	}

	var sb strings.Builder
	className := o.Name.ToTitleCase()

	if o.Description != "" {
		sb.WriteString("/// <summary>\n")
		for _, line := range strings.Split(o.Description, "\n") {
			sb.WriteString(fmt.Sprintf("/// %s\n", strings.TrimSpace(line)))
		}
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString(fmt.Sprintf("public class %s\n{\n", className))

	for _, name := range keys {
		val, err := o.Properties.Get(name)
		if err != nil {
			return "", err
		}

		description := val.GetDescription()
		if description != "" {
			sb.WriteString("    /// <summary>\n")
			desc := strings.ReplaceAll(description, "\n", " ")
			sb.WriteString(fmt.Sprintf("    /// %s\n", desc))
			sb.WriteString("    /// </summary>\n")
		}

		propName := utils.DashCaseToTitleCase(name)
		if propName != name {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", name))
		}

		typeName := val.CSharpType()
		if !slices.Contains(o.Required, name) && !isReferenceType(typeName) {
			typeName += "?"
		}

		sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n\n", typeName, propName))
	}

	sb.WriteString("}\n")
	return sb.String(), nil
}

func isReferenceType(typeName string) bool {
	referenceTypes := map[string]bool{
		"string": true,
		"object": true,
	}
	if strings.HasPrefix(typeName, "List<") ||
		strings.HasPrefix(typeName, "Dictionary<") ||
		strings.HasSuffix(typeName, "[]") {
		return true
	}
	if referenceTypes[typeName] {
		return true
	}
	if len(typeName) > 0 && typeName[0] >= 'A' && typeName[0] <= 'Z' {
		return true
	}
	return false
}

func (o Object) CSharpType() string {
	return o.Name.ToTitleCase()
}
