package types

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
)

type Value interface {
	Go() (string, error)
	GoStructType() string
	GoStructKey(isUnion bool) string

	IsReference() bool
}

func unionDefinitionToUnionValue(propertyName PropertyName, description string, items []schema.PropertyDefinition) (Union, error) {
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
				Name:        propertyName,
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
				objProp, err := PropertyDefinitionToValue(
					propMap,
					name,
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
				union, err := unionDefinitionToUnionValue(propertyName, item.Description, item.Items.AnyOf)
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
				union, err := unionDefinitionToUnionValue(propertyName, item.Description, item.Items.OneOf)
				if err != nil {
					return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
				}

				typeIdentifiers = append(typeIdentifiers, Array{
					Name: propertyName,
					Type: union,
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
				fmt.Println(item.AnyOf)
				panic("unsupported array type")
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

func PropertyDefinitionToValue(definitions map[string]schema.PropertyDefinition, name string) (Value, error) {
	property, ok := definitions[name]
	if !ok {
		return nil, fmt.Errorf("no type for [%s] found in definitions", name)
	}

	propertyName := NewPropertyName(name)

	// Union
	if property.OneOf != nil {
		return unionDefinitionToUnionValue(propertyName, property.Description, property.OneOf)
	}
	if property.AnyOf != nil {
		return unionDefinitionToUnionValue(propertyName, property.Description, property.AnyOf)
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
			union, err := unionDefinitionToUnionValue(propertyName, property.Description, property.Items.AnyOf)
			if err != nil {
				return Union{}, fmt.Errorf("converting array union defintion for union [%s]: %v", propertyName.Value, err)
			}

			return Array{
				Name: propertyName,
				Type: union,
			}, nil
		}
		if property.Items.OneOf != nil {
			union, err := unionDefinitionToUnionValue(propertyName, property.Description, property.Items.OneOf)
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
			fmt.Println(refName)
			arrayType, err := PropertyDefinitionToValue(definitions, refName)
			if err != nil {
				return nil, fmt.Errorf("finding ref def: %v", err)
			}

			return Array{
				Name: propertyName,
				Type: arrayType,
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
			panic("unsupported array type")
		}
	}

	// Object
	if property.Type == "object" {
		properties := make(map[string]Value, len(property.Properties))
		for name, prop := range property.Properties {
			if prop.Ref != "" {
				if prop.Ref.IsNested() {
					return nil, fmt.Errorf("nested references not implemented")
				}

				refProp, err := PropertyDefinitionToValue(definitions, prop.Ref.Name())
				if err != nil {
					return nil, fmt.Errorf("converting reference for [%s]: %v", propertyName.Value, err)
				}
				properties[name] = PropertyReference{
					Name: prop.Ref.Name(),
					Ref:  string(prop.Ref),
					Type: refProp,
				}
				continue
			}

			objProp, err := PropertyDefinitionToValue(property.Properties, name)
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
