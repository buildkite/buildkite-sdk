package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	resp, err := http.Get("https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var schema PipelineSchema
	err = json.Unmarshal(body, &schema)
	if err != nil {
		return PipelineSchema{}, fmt.Errorf("parsing schema: %v", err)
	}

	return schema, nil
}
