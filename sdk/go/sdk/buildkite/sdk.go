package buildkite

import (
	"encoding/json"

	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"
	"gopkg.in/yaml.v3"
)

func NewPipeline() *pipeline {
	return &pipeline{
		Schema: schema.Schema{},
	}
}

type pipeline struct {
	schema.Schema

	Agents *schema.Agents  `json:"agents,omitempty"`
	Steps  []*PipelineStep `json:"steps"`
}

type pipelineStep interface {
	toPipelineStep() *PipelineStep
}

func (p *pipeline) AddStep(step pipelineStep) {
	p.Steps = append(p.Steps, step.toPipelineStep())
}

func (p *pipeline) ToJSON() (string, error) {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *pipeline) ToYAML() (string, error) {
	data, err := yaml.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
