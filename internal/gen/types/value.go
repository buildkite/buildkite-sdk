package types

import (
	"fmt"
	"sort"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
)

type PipelineSchemaGenerator struct {
	Definitions *orderedmap.OrderedMap
	Properties  *orderedmap.OrderedMap
}

var pipelineFunctions = `func (p Pipeline) ToJSON() (string, error) {
    rawJSON, err := json.Marshal(p)
	if err != nil {
	    return "", err
	}
	return string(rawJSON), nil
}

func (p *Pipeline) AddStep(step PipelineStepsUnion) {
	steps := p.Steps
	if steps == nil {
		steps = &[]PipelineStepsUnion{}
	}

	newSteps := append(*steps, step)
	p.Steps = &newSteps
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
	env := *p.Env
	if p.Env == nil {
		env = map[string]interface{}{}
	}

	env[key] = value
	p.Env = &env
}

func (p *Pipeline) AddNotify(notify BuildNotifyUnion) {
	foo := []BuildNotifyUnion{notify}
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
}`

func (p PipelineSchemaGenerator) GeneratePipelineSchema() (string, error) {
	goStruct := utils.NewGoStruct("Pipeline", "", nil)

	for _, name := range p.Properties.Keys() {
		val, _ := p.Properties.Get(name)
		prop := val.(schema.SchemaProperty)

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
	val, _ := p.Definitions.Get(firstKey)
	currentDef := val.(schema.PropertyDefinition)

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

func (p PipelineSchemaGenerator) PropertyDefinitionToValue(name string, property schema.PropertyDefinition) (Value, error) {
	propertyName := NewPropertyName(name)

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
		}, nil
	}

	// Array
	if property.Type == "array" {
		if property.Items.AnyOf != nil {
			union, err := p.UnionDefinitionToUnionValue(propertyName, property.Description, property.Items.AnyOf)
			if err != nil {
				return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        union,
			}, nil
		}
		if property.Items.OneOf != nil {
			union, err := p.UnionDefinitionToUnionValue(propertyName, property.Description, property.Items.OneOf)
			if err != nil {
				return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        union,
			}, nil
		}

		if property.Items.Ref != "" {
			if property.Items.Ref.IsNested() {
				return nil, fmt.Errorf("nested refs are not supported")
			}

			refName := property.Items.Ref.Name()
			property := p.ResolveReference(property.Items.Ref)
			arrayType, err := p.PropertyDefinitionToValue(refName, property)
			if err != nil {
				return nil, fmt.Errorf("finding ref def: %v", err)
			}

			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        arrayType,
				Reference:   true,
			}, nil
		}

		switch property.Items.Type {
		case "string":
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        String{},
			}, nil
		case "integer":
			return Array{
				Name:        propertyName,
				Description: property.Description,
				Type:        Number{},
			}, nil
		default:
			panic(fmt.Sprintf("unsupported array type [%s]", propertyName.Value))
		}
	}

	// Object
	if property.Type == "object" {
		properties := orderedmap.New()
		for name, prop := range property.Properties {
			if prop.Ref != "" {
				refProp := p.ResolveReference(prop.Ref)
				refVal, err := p.PropertyDefinitionToValue(name, refProp)
				if err != nil {
					return nil, fmt.Errorf("converting reference for [%s]: %v", propertyName.Value, err)
				}

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

			objProp, err := p.PropertyDefinitionToValue(
				propName,
				prop,
			)
			if err != nil {
				return nil, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
			}
			properties.Set(name, objProp)
		}
		properties.SortKeys(sort.Strings)

		if property.AdditionalProperties.Type != "" {
			propDef := schema.PropertyDefinition{
				Type:        schema.PropertyType(property.AdditionalProperties.Type),
				Description: property.AdditionalProperties.Description,
				Items:       property.AdditionalProperties.Items,
			}

			additionalProperties, err := p.PropertyDefinitionToValue("", propDef)
			if err != nil {
				return nil, fmt.Errorf("determing type of additional properties: %v", err)
			}

			return Object{
				Name:                 propertyName,
				Description:          propDef.Description,
				Properties:           properties,
				AdditionalProperties: &additionalProperties,
				Required:             property.Required,
			}, nil
		}

		return Object{
			Name:        propertyName,
			Description: property.Description,
			Properties:  properties,
			Required:    property.Required,
		}, nil
	}

	// String
	if property.Type == "string" {
		return String{
			Name:        propertyName,
			Description: property.Description,
		}, nil
	}

	// Number
	if property.Type == "integer" {
		return Number{
			Name:        propertyName,
			Description: property.Description,
		}, nil
	}

	// Boolean
	if property.Type == "boolean" {
		return Boolean{
			Name:        propertyName,
			Description: property.Description,
		}, nil
	}

	return nil, fmt.Errorf("type for [%s] has not been implemented", name)
}

func (p PipelineSchemaGenerator) UnionDefinitionToUnionValue(propertyName PropertyName, description string, items []schema.PropertyDefinition) (Union, error) {
	var typeIdentifiers []Value
	for _, item := range items {
		// Skip Null
		if item.Type == "null" {
			continue
		}

		// Reference
		if item.Ref != "" {
			refName := item.Ref.Name()
			refDef := p.ResolveReference(item.Ref)
			refTyp, err := p.PropertyDefinitionToValue(refName, refDef)
			if err != nil {
				return Union{}, fmt.Errorf("looking up reference type: %v", err)
			}

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
			properties := orderedmap.New()
			for name, prop := range item.Properties {
				objProp, err := p.PropertyDefinitionToValue(
					name,
					prop,
				)
				if err != nil {
					return Union{}, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
				}
				properties.Set(name, objProp)
			}
			properties.SortKeys(sort.Strings)

			if item.AdditionalProperties.Type != "" {
				propDef := schema.PropertyDefinition{
					Type:        schema.PropertyType(item.AdditionalProperties.Type),
					Description: item.AdditionalProperties.Description,
					Items:       item.AdditionalProperties.Items,
				}

				additionalProperties, err := p.PropertyDefinitionToValue("", propDef)
				if err != nil {
					return Union{}, fmt.Errorf("determing type of additional properties: %v", err)
				}

				typeIdentifiers = append(typeIdentifiers, Object{
					Name:                 propertyName,
					Properties:           properties,
					AdditionalProperties: &additionalProperties,
					Required:             item.Required,
				})
				continue
			}

			typeIdentifiers = append(typeIdentifiers, Object{
				Name:       propertyName,
				Properties: properties,
				Required:   item.Required,
			})
			continue
		}

		// Array
		if item.Type == "array" {
			if item.Items.AnyOf != nil {
				union, err := p.UnionDefinitionToUnionValue(propertyName, item.Description, item.Items.AnyOf)
				if err != nil {
					return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
				}

				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: union,
				})
				continue
			}
			if item.Items.OneOf != nil {
				union, err := p.UnionDefinitionToUnionValue(propertyName, item.Description, item.Items.OneOf)
				if err != nil {
					return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
				}

				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: union,
				})
				continue
			}

			if item.Items.Ref != "" {
				if item.Items.Ref.IsNested() {
					return Union{}, fmt.Errorf("nested refs are not supported")
				}

				refName := item.Items.Ref.Name()
				val, _ := p.Definitions.Get(refName)
				property := val.(schema.PropertyDefinition)

				arrayType, err := p.PropertyDefinitionToValue(refName, property)
				if err != nil {
					return Union{}, fmt.Errorf("finding ref def: %v", err)
				}

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

		return Union{}, fmt.Errorf("union type not implemented")
	}

	return Union{
		Name:            propertyName,
		Description:     description,
		TypeIdentifiers: typeIdentifiers,
	}, nil
}

type Value interface {
	GetDescription() string

	// Go
	Go() (string, error)
	GoStructType() string
	GoStructKey(isUnion bool) string

	// TypeScript
	TypeScript() (string, error)
	TypeScriptInterfaceKey() string
	TypeScriptInterfaceType() string

	// Python
	Python() (string, error)
	PythonClassKey() string
	PythonClassType() string

	IsReference() bool
	IsPrimative() bool
}
