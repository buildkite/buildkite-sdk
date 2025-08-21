package schema

import "strings"

type PropertyReferenceString string

func (p PropertyReferenceString) Name() string {
	parts := strings.Split(string(p), "/")
	return parts[len(parts)-1]
}

func (p PropertyReferenceString) IsNested() bool {
	parts := strings.Split(string(p), "/")
	return len(parts) > 3
}

type PropertyDefinitionItems struct {
	Type  string                  `json:"type,omitempty"`
	AnyOf []PropertyDefinition    `json:"anyOf,omitempty"`
	OneOf []PropertyDefinition    `json:"oneOf,omitempty"`
	Ref   PropertyReferenceString `json:"$ref,omitempty"`
}

type PropertyDefinition struct {
	Ref         PropertyReferenceString       `json:"$ref,omitempty"`
	Enum        []any                         `json:"enum,omitempty"`
	Type        string                        `json:"type"`
	Description string                        `json:"description"`
	Default     any                           `json:"default,omitempty"`
	OneOf       []PropertyDefinition          `json:"oneOf,omitempty"`
	AnyOf       []PropertyDefinition          `json:"anyOf,omitempty"`
	Items       PropertyDefinitionItems       `json:"items,omitempty"`
	Properties  map[string]PropertyDefinition `json:"properties,omitempty"`
	Examples    []any                         `json:"examples,omitempty"`
}
