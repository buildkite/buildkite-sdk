package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

type PipelineSchemaGenerator struct {
	Definitions map[string]schema.PropertyDefinition
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
				Name: propertyName,
				Type: union,
			}, nil
		}
		if property.Items.OneOf != nil {
			union, err := p.UnionDefinitionToUnionValue(propertyName, property.Description, property.Items.OneOf)
			if err != nil {
				return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			return Array{
				Name: propertyName,
				Type: union,
			}, nil
		}

		if property.Items.Ref != "" {
			if property.Items.Ref.IsNested() {
				return nil, fmt.Errorf("nested refs are not supported")
			}

			refName := property.Items.Ref.Name()
			property := p.Definitions[refName]
			arrayType, err := p.PropertyDefinitionToValue(refName, property)
			if err != nil {
				return nil, fmt.Errorf("finding ref def: %v", err)
			}

			return Array{
				Name:      propertyName,
				Type:      arrayType,
				Reference: true,
			}, nil
		}

		switch property.Items.Type {
		case "string":
			return Array{
				Name: propertyName,
				Type: String{},
			}, nil
		case "integer":
			return Array{
				Name: propertyName,
				Type: Number{},
			}, nil
		default:
			panic(fmt.Sprintf("unsupported array type [%s]", propertyName.Value))
		}
	}

	// Object
	if property.Type == "object" {
		properties := make(map[string]Value, len(property.Properties))
		for name, prop := range property.Properties {
			if prop.Ref != "" {
				refName := prop.Ref.Name()
				refProp := p.Definitions[refName]

				if prop.Ref.IsNested() {
					refProp = property.Properties[refName]
					if refProp.Ref != "" {
						refProp = p.Definitions[refProp.Ref.Name()]
					}

					fmt.Println(refName)
					fmt.Println(refProp.Ref)
				}

				refVal, err := p.PropertyDefinitionToValue(refName, refProp)
				if err != nil {
					return nil, fmt.Errorf("converting reference for [%s]: %v", propertyName.Value, err)
				}
				properties[name] = PropertyReference{
					Name: refName,
					Ref:  string(prop.Ref),
					Type: refVal,
				}
				continue
			}

			objProp, err := p.PropertyDefinitionToValue(
				fmt.Sprintf("%s%s", propertyName.Value, utils.DashCaseToTitleCase(name)),
				prop,
			)
			if err != nil {
				return nil, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
			}
			properties[name] = objProp
		}

		return Object{
			Name:       propertyName,
			Properties: properties,
		}, nil
	}

	// String
	if property.Type == "string" {
		return String{
			Name: propertyName,
		}, nil
	}

	// Number
	if property.Type == "integer" {
		return Number{
			Name: propertyName,
		}, nil
	}

	// Boolean
	if property.Type == "boolean" {
		return Boolean{
			Name: propertyName,
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
			typeIdentifiers = append(typeIdentifiers, PropertyReference{
				Name: refName,
				Ref:  string(item.Ref),
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
			properties := make(map[string]Value, len(item.Properties))
			for name, prop := range item.Properties {
				propMap := make(map[string]schema.PropertyDefinition, 1)
				propMap[name] = prop
				objProp, err := p.PropertyDefinitionToValue(
					name,
					propMap[name],
				)
				if err != nil {
					return Union{}, fmt.Errorf("converting object property to value [%s]: %v", propertyName.Value, err)
				}
				properties[name] = objProp
			}

			typeIdentifiers = append(typeIdentifiers, Object{
				Name:       propertyName,
				Properties: properties,
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
				property := p.Definitions[refName]
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
				Name: propertyName,
			})
			continue
		}

		if item.Type == "integer" {
			typeIdentifiers = append(typeIdentifiers, Number{
				Name: propertyName,
			})
			continue
		}

		if item.Type == "boolean" {
			typeIdentifiers = append(typeIdentifiers, Boolean{
				Name: propertyName,
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
	Go() (string, error)
	GoStructType() string
	GoStructKey(isUnion bool) string

	IsReference() bool
}
