package schema

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PropertyReferenceString string

func (p PropertyReferenceString) Name() string {
	parts := strings.Split(string(p), "/")
	return parts[len(parts)-1]
}

func (p PropertyReferenceString) Keys() []string {
	parts := strings.Split(string(p), "/")
	return parts[2:]
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

type PropertyAdditionalProperties struct {
	Type        string                  `json:"type,omitempty"`
	Description string                  `json:"description,omitempty"`
	Items       PropertyDefinitionItems `json:"items,omitempty"`
}

func (p *PropertyAdditionalProperties) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as a boolean
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*p = PropertyAdditionalProperties{}
		return nil
	}

	// If not a boolean, try unmarshaling as an object
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err == nil {
		var items PropertyDefinitionItems
		itemsString, err := json.Marshal(obj["items"])
		if err != nil {
			return fmt.Errorf("marshal additional properties items: %v", err)
		}

		err = json.Unmarshal(itemsString, &items)
		if err != nil {
			return fmt.Errorf("unmarhsaling additional properties: %v", err)
		}

		var typeStr, descStr string
		if t, ok := obj["type"].(string); ok {
			typeStr = t
		}
		if d, ok := obj["description"].(string); ok {
			descStr = d
		}

		*p = PropertyAdditionalProperties{
			Type:        typeStr,
			Description: descStr,
			Items:       items,
		}
		return nil
	}

	return fmt.Errorf("field is neither a boolean nor an object: %s", string(data))
}

type PropertyType string

func (p *PropertyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*p = PropertyType(s)
		return nil
	}

	// Attempt to unmarshal as a string array
	var sa []string
	if err := json.Unmarshal(data, &sa); err == nil {
		// TODO: make this not a hack
		*p = PropertyType("string")
		return nil
	}

	// If neither worked, return an error
	return fmt.Errorf("cannot unmarshal %s into PropertyType", data)
}

type PropertyDefinition struct {
	Ref                  PropertyReferenceString       `json:"$ref,omitempty"`
	Enum                 []any                         `json:"enum,omitempty"`
	Type                 PropertyType                  `json:"type"`
	Description          string                        `json:"description"`
	Default              any                           `json:"default,omitempty"`
	OneOf                []PropertyDefinition          `json:"oneOf,omitempty"`
	AnyOf                []PropertyDefinition          `json:"anyOf,omitempty"`
	Items                PropertyDefinitionItems       `json:"items,omitempty"`
	Properties           map[string]PropertyDefinition `json:"properties,omitempty"`
	AdditionalProperties PropertyAdditionalProperties  `json:"additionalProperties"`
	Examples             []any                         `json:"examples,omitempty"`
	Required             []string                      `json:"required"`
}
