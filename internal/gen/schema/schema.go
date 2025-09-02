package schema

import (
	"encoding/json"
	"fmt"
	"os"
)

type PipelineSchema struct {
	Title       string                        `json:"title,omitempty"`
	Schema      string                        `json:"$schema,omitempty"`
	FileMatch   []string                      `json:"fileMatch,omitempty"`
	Type        string                        `json:"type,omitempty"`
	Required    []string                      `json:"required,omitempty"`
	Definitions map[string]PropertyDefinition `json:"definitions,omitempty"`
	Properties  map[string]SchemaProperty     `json:"properties,omitempty"`
}

type SchemaProperty struct {
	Ref PropertyReferenceString `json:"$ref,omitempty"`
}

func ReadSchema() (PipelineSchema, error) {
	file, err := os.ReadFile("internal/gen/schema.json")
	if err != nil {
		return PipelineSchema{}, fmt.Errorf("reading schema file: %v", err)
	}

	var schema PipelineSchema
	err = json.Unmarshal(file, &schema)
	if err != nil {
		return PipelineSchema{}, fmt.Errorf("parsing schema: %v", err)
	}

	return schema, nil
}
