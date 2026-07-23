package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
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
	Ref         PropertyReferenceString    `json:"$ref,omitempty"`
	Description string                     `json:"description,omitempty"`
	AllOf       []SchemaProperty           `json:"allOf,omitempty"`
	Properties  map[string]json.RawMessage `json:"properties,omitempty"`
}

// RemovedProperties lists the property keys a schema property narrows away
// with a literal `false` subschema, e.g. the pipeline-level checkout
// forbidding the step-only ssh_secret key.
func (s SchemaProperty) RemovedProperties() []string {
	removed := []string{}
	for key, raw := range s.Properties {
		if string(raw) == "false" {
			removed = append(removed, key)
		}
	}
	sort.Strings(removed)
	return removed
}

// Reference resolves the property's definition reference: either a direct
// $ref, or one wrapped in an allOf (used to narrow a definition, e.g. the
// pipeline-level checkout property, which forbids the step-only ssh_secret
// key). Narrowing keywords alongside the allOf are validation-only and do not
// change the generated type.
func (s SchemaProperty) Reference() PropertyReferenceString {
	if s.Ref != "" {
		return s.Ref
	}
	for _, member := range s.AllOf {
		if ref := member.Reference(); ref != "" {
			return ref
		}
	}
	return ""
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
