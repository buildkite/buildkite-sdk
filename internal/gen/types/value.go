package types

import (
	"fmt"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

type PipelineSchemaGenerator struct {
	Definitions *utils.OrderedMap[schema.PropertyDefinition]
	Properties  *utils.OrderedMap[schema.SchemaProperty]
}

func NewPipelineSchemaGenerator(pipelineSchema schema.PipelineSchema) PipelineSchemaGenerator {
	return PipelineSchemaGenerator{
		Definitions: utils.NewOrderedMap(pipelineSchema.Definitions),
		Properties:  utils.NewOrderedMap(pipelineSchema.Properties),
	}
}

var pipelineFunctions = `func (p Pipeline) ToJSON() (string, error) {
    rawJSON, err := json.Marshal(p)
	if err != nil {
	    return "", err
	}
	return string(rawJSON), nil
}

func (p *Pipeline) AddStep(step pipelineStep) {
	steps := p.Steps
	if steps == nil {
		steps = &[]PipelineStepsItem{}
	}

	newSteps := append(*steps, step.toStepUnion())
	p.Steps = &newSteps
}

func (p *Pipeline) SetSecrets(secrets *Secrets) {
	p.Secrets = secrets
}

func (p *Pipeline) AddAgent(key string, value any) {
	agents := map[string]interface{}{}
	if p.Agents != nil {
		agents = *p.Agents.AgentsObject
	}

	agents[key] = value
	p.Agents = &Agents{
		AgentsObject: &agents,
	}
}

func (p *Pipeline) AddEnvironmentVariable(key string, value any) {
	env := map[string]interface{}{}
	if p.Env != nil {
		env = *p.Env
	}

	env[key] = value
	p.Env = &env
}

func (p *Pipeline) AddNotify(notify BuildNotifyItem) {
	foo := []BuildNotifyItem{notify}
	p.Notify = &foo
}

func (p *Pipeline) ToYAML() (string, error) {
	data, err := p.ToJSON()
	if err != nil {
		return "", err
	}

	var output strings.Builder
	input := strings.NewReader(data)
	if err := json2yaml.Convert(&output, input); err != nil {
		return "", fmt.Errorf("converting JSON to YAML: %v", err)
	}

	return output.String(), nil
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

type pipelineStep interface {
	toStepUnion() PipelineStepsItem
}

func (s BlockStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		BlockStep: &s,
	}
}
func (s CommandStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		CommandStep: &s,
	}
}
func (s GroupStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		GroupStep: &s,
	}
}
func (s InputStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		InputStep: &s,
	}
}
func (s NestedBlockStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		NestedBlockStep: &s,
	}
}
func (s NestedCommandStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		NestedCommandStep: &s,
	}
}
func (s NestedInputStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		NestedInputStep: &s,
	}
}
func (s NestedTriggerStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		NestedTriggerStep: &s,
	}
}
func (s NestedWaitStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		NestedWaitStep: &s,
	}
}
func (s StringBlockStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		StringBlockStep: &s,
	}
}
func (s StringInputStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		StringInputStep: &s,
	}
}
func (s StringWaitStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		StringWaitStep: &s,
	}
}
func (s TriggerStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		TriggerStep: &s,
	}
}
func (s WaitStep) toStepUnion() PipelineStepsItem {
	return PipelineStepsItem{
		WaitStep: &s,
	}
}

func Value[T any](val T) *T {
	return &val
}
`

func (p PipelineSchemaGenerator) GenerateCSharpPipelineSchema() (string, error) {
	var sb strings.Builder

	sb.WriteString("/// <summary>\n")
	sb.WriteString("/// Represents a Buildkite pipeline configuration.\n")
	sb.WriteString("/// </summary>\n")
	sb.WriteString("public class BuildkitePipeline\n{\n")

	for _, name := range p.Properties.Keys() {
		prop, err := p.Properties.Get(name)
		if err != nil {
			return "", fmt.Errorf("generating pipeline schema: %v", err)
		}

		propName := utils.DashCaseToTitleCase(name)
		typeName := utils.ToTitleCase(prop.Ref.Name())

		if propName != name {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", name))
		}

		sb.WriteString(fmt.Sprintf("    public %s? %s { get; set; }\n\n", typeName, propName))
	}

	sb.WriteString("}\n")
	return sb.String(), nil
}

func (p PipelineSchemaGenerator) GeneratePipelineSchema() (string, error) {
	goStruct := utils.NewGoStruct("Pipeline", "", nil)

	for _, name := range p.Properties.Keys() {
		prop, err := p.Properties.Get(name)
		if err != nil {
			return "", fmt.Errorf("generating pipeline schema: %v", err)
		}

		structKey := utils.DashCaseToTitleCase(name)
		structType := utils.CamelCaseToTitleCase(prop.Ref.Name())
		goStruct.AddItem(structKey, structType, name, "", true)
	}

	structString, err := goStruct.Write()
	if err != nil {
		return "", fmt.Errorf("generating pipeline struct")
	}

	codeBlock := utils.NewCodeBlock(
		structString,
		"",
		pipelineFunctions,
	)

	return codeBlock.String(), nil
}

func (p PipelineSchemaGenerator) ResolveReference(ref schema.PropertyReferenceString) schema.PropertyDefinition {
	keys := ref.Keys()
	firstKey := keys[0]
	currentDef, err := p.Definitions.Get(firstKey)
	if err != nil {
		panic(fmt.Sprintf("reference not found for %s", firstKey))
	}

	if len(keys) == 1 {
		return currentDef
	}

	// Nested references contain 'properties' in they keys slice
	// so we skip it here to get the actual reference.
	for _, key := range keys[2:] {
		currentDef = currentDef.Properties[key]
	}

	if currentDef.Ref != "" {
		return p.ResolveReference(currentDef.Ref)
	}

	return currentDef
}

func (p PipelineSchemaGenerator) PropertyDefinitionToValue(name string, property schema.PropertyDefinition) (Value, []string, error) {
	propertyName := NewPropertyName(name)
	dependencies := []string{}

	// Union
	if property.OneOf != nil {
		return p.UnionDefinitionToUnionValue(propertyName, property.Description, property.OneOf)
	}
	if property.AnyOf != nil {
		return p.UnionDefinitionToUnionValue(propertyName, property.Description, property.AnyOf)
	}

	// Enum
	if property.Enum != nil {
		return Enum{
			Name:        propertyName,
			Description: property.Description,
			Values:      property.Enum,
			Default:     property.Default,
		}, dependencies, nil
	}

	// Array
	if property.Type == "array" {
		if property.Items.AnyOf != nil {
			union, unionDependencies, err := p.UnionDefinitionToUnionValue(propertyName, property.Description, property.Items.AnyOf)
			if err != nil {
				return Union{}, dependencies, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			dependencies = append(dependencies, unionDependencies...)
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        union,
			}, unionDependencies, nil
		}
		if property.Items.OneOf != nil {
			union, unionDependencies, err := p.UnionDefinitionToUnionValue(propertyName, property.Description, property.Items.OneOf)
			if err != nil {
				return Union{}, dependencies, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			dependencies = append(dependencies, unionDependencies...)
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        union,
			}, dependencies, nil
		}

		if property.Items.Ref != "" {
			if property.Items.Ref.IsNested() {
				return nil, dependencies, fmt.Errorf("nested refs are not supported")
			}

			refName := property.Items.Ref.Name()
			property := p.ResolveReference(property.Items.Ref)
			arrayType, _, err := p.PropertyDefinitionToValue(refName, property)
			if err != nil {
				return nil, dependencies, fmt.Errorf("finding ref def: %v", err)
			}

			dependencies = append(dependencies, refName)
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        arrayType,
				Reference:   true,
			}, dependencies, nil
		}

		switch property.Items.Type {
		case "string":
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        String{},
			}, dependencies, nil
		case "integer":
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        Number{},
			}, dependencies, nil
		default:
			panic(fmt.Sprintf("unsupported array type [%s]", propertyName.Value))
		}
	}

	// Object
	if property.Type == "object" {
		properties := utils.NewOrderedMap[Value](nil)
		for name, prop := range property.Properties {
			if prop.Ref != "" {
				refProp := p.ResolveReference(prop.Ref)
				refVal, objDependencies, err := p.PropertyDefinitionToValue(name, refProp)
				if err != nil {
					return nil, dependencies, fmt.Errorf("converting reference for [%s]: %v", propertyName.Value, err)
				}

				dependencies = append(dependencies, objDependencies...)
				properties.Set(name, PropertyReference{
					Name: name,
					Ref:  prop.Ref,
					Type: refVal,
				})
				continue
			}

			var propName string
			switch prop.Type {
			case "string":
				fallthrough
			case "integer":
				fallthrough
			case "boolean":
				propName = name
			default:
				propName = fmt.Sprintf("%s%s", propertyName.Value, utils.DashCaseToTitleCase(name))
			}

			objProp, objDependencies, err := p.PropertyDefinitionToValue(
				propName,
				prop,
			)
			if err != nil {
				return nil, dependencies, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
			}

			dependencies = append(dependencies, objDependencies...)

			// Nested object
			if nestedObject, ok := objProp.(Object); ok {
				objProp = Object{
					Name:                 nestedObject.Name,
					Description:          nestedObject.Description,
					Properties:           nestedObject.Properties,
					AdditionalProperties: nestedObject.AdditionalProperties,
					Required:             nestedObject.Required,
					IsNested:             true,
				}
			}

			properties.Set(name, objProp)
		}
		properties.SortKeys()

		if property.AdditionalProperties.Type != "" {
			propDef := schema.PropertyDefinition{
				Type:        schema.PropertyType(property.AdditionalProperties.Type),
				Description: property.AdditionalProperties.Description,
				Items:       property.AdditionalProperties.Items,
			}

			additionalProperties, _, err := p.PropertyDefinitionToValue("", propDef)
			if err != nil {
				return nil, dependencies, fmt.Errorf("determing type of additional properties: %v", err)
			}

			return Object{
				Name:                 propertyName,
				Description:          propDef.Description,
				Properties:           properties,
				AdditionalProperties: &additionalProperties,
				Required:             property.Required,
			}, dependencies, nil
		}

		return Object{
			Name:        propertyName,
			Description: property.Description,
			Properties:  properties,
			Required:    property.Required,
		}, dependencies, nil
	}

	// String
	if property.Type == "string" {
		return String{
			Name:        propertyName,
			Description: property.Description,
		}, dependencies, nil
	}

	// Number
	if property.Type == "integer" {
		return Number{
			Name:        propertyName,
			Description: property.Description,
		}, dependencies, nil
	}

	// Boolean
	if property.Type == "boolean" {
		return Boolean{
			Name:        propertyName,
			Description: property.Description,
		}, dependencies, nil
	}

	return nil, dependencies, fmt.Errorf("type for [%s] has not been implemented", name)
}

func (p PipelineSchemaGenerator) UnionDefinitionToUnionValue(propertyName PropertyName, description string, items []schema.PropertyDefinition) (Union, []string, error) {
	var typeIdentifiers []Value
	dependencies := []string{}
	for _, item := range items {
		// Skip Null
		if item.Type == "null" {
			continue
		}

		// Reference
		if item.Ref != "" {
			refName := item.Ref.Name()
			refDef := p.ResolveReference(item.Ref)
			refTyp, _, err := p.PropertyDefinitionToValue(refName, refDef)
			if err != nil {
				return Union{}, dependencies, fmt.Errorf("looking up reference type: %v", err)
			}

			dependencies = append(dependencies, string(item.Ref.Name()))
			typeIdentifiers = append(typeIdentifiers, PropertyReference{
				Name: refName,
				Ref:  item.Ref,
				Type: refTyp,
			})
			continue
		}

		// Enum
		if item.Enum != nil {
			typeIdentifiers = append(typeIdentifiers, Enum{
				Name:        NewPropertyName(fmt.Sprintf("%sEnum", propertyName.Value)),
				Description: item.Description,
				Values:      item.Enum,
				Default:     item.Default,
			})
			continue
		}

		// Object
		if item.Type == "object" {
			properties := utils.NewOrderedMap[Value](nil)
			for name, prop := range item.Properties {
				objProp, _, err := p.PropertyDefinitionToValue(
					name,
					prop,
				)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
				}

				// Nested object
				if nestedObject, ok := objProp.(Object); ok {
					objProp = Object{
						Name:                 nestedObject.Name,
						Properties:           nestedObject.Properties,
						AdditionalProperties: nestedObject.AdditionalProperties,
						Required:             nestedObject.Required,
						IsNested:             true,
					}
				}

				properties.Set(name, objProp)
			}
			properties.SortKeys()

			if item.AdditionalProperties.Type != "" {
				propDef := schema.PropertyDefinition{
					Type:        schema.PropertyType(item.AdditionalProperties.Type),
					Description: item.AdditionalProperties.Description,
					Items:       item.AdditionalProperties.Items,
				}

				additionalProperties, _, err := p.PropertyDefinitionToValue("", propDef)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("determing type of additional properties: %v", err)
				}

				typeIdentifiers = append(typeIdentifiers, Object{
					Name:                 propertyName,
					Properties:           properties,
					AdditionalProperties: &additionalProperties,
					Required:             item.Required,
					IsNested:             true,
				})
				continue
			}

			typeIdentifiers = append(typeIdentifiers, Object{
				Name:       propertyName,
				Properties: properties,
				Required:   item.Required,
				IsNested:   true,
			})
			continue
		}

		// Array
		if item.Type == "array" {
			if item.Items.AnyOf != nil {
				union, unionDependencies, err := p.UnionDefinitionToUnionValue(propertyName, item.Description, item.Items.AnyOf)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
				}

				dependencies = append(dependencies, unionDependencies...)
				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: union,
				})
				continue
			}
			if item.Items.OneOf != nil {
				union, unionDependencies, err := p.UnionDefinitionToUnionValue(propertyName, item.Description, item.Items.OneOf)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
				}

				dependencies = append(dependencies, unionDependencies...)
				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: union,
				})
				continue
			}

			if item.Items.Ref != "" {
				if item.Items.Ref.IsNested() {
					return Union{}, dependencies, fmt.Errorf("nested refs are not supported")
				}

				refName := item.Items.Ref.Name()
				property, err := p.Definitions.Get(refName)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("getting definition for %s", refName)
				}

				arrayType, _, err := p.PropertyDefinitionToValue(refName, property)
				if err != nil {
					return Union{}, dependencies, fmt.Errorf("finding ref def: %v", err)
				}

				dependencies = append(dependencies, string(item.Items.Ref.Name()))
				typeIdentifiers = append(typeIdentifiers, Array{
					Name: NewPropertyName(refName),
					Type: arrayType,
				})
				continue
			}

			switch item.Items.Type {
			case "string":
				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: String{},
				})
				continue
			case "integer":
				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: Number{},
				})
				continue
			default:
				panic(fmt.Sprintf("unsupported array type [%s]", propertyName.Value))
			}
		}

		if item.Type == "string" {
			typeIdentifiers = append(typeIdentifiers, String{
				Name:        propertyName,
				Description: item.Description,
			})
			continue
		}

		if item.Type == "integer" {
			typeIdentifiers = append(typeIdentifiers, Number{
				Name:        propertyName,
				Description: item.Description,
			})
			continue
		}

		if item.Type == "boolean" {
			typeIdentifiers = append(typeIdentifiers, Boolean{
				Name:        propertyName,
				Description: item.Description,
			})
			continue
		}

		return Union{}, dependencies, fmt.Errorf("union type not implemented")
	}

	return Union{
		Name:            propertyName,
		Description:     description,
		TypeIdentifiers: typeIdentifiers,
	}, dependencies, nil
}

type Value interface {
	GetDescription() string

	// Go
	Go() (string, error)
	GoStructType() string
	GoStructKey(isUnion bool) string

	// TypeScript
	TypeScript() string
	TypeScriptInterfaceKey() string
	TypeScriptInterfaceType() string

	// Python
	Python() (string, error)
	PythonClassKey() string
	PythonClassType() string

	// CSharp
	CSharp() (string, error)
	CSharpType() string

	IsReference() bool
	IsPrimitive() bool
}
