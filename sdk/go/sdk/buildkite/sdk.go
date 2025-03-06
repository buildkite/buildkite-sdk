package buildkite

import (
	"encoding/json"

	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"
	"gopkg.in/yaml.v3"
)

func NewPipeline() pipeline {
	return pipeline{
		Schema: schema.Schema{},
	}
}

type pipeline struct {
	schema.Schema

	Agents *schema.Agents  `json:"agents,omitempty"`
	Steps  []*PipelineStep `json:"steps"`
}

func (p *pipeline) AddCommandStep(step schema.CommandStep) {
	p.Steps = append(p.Steps, commandStepToGroupStep(step))
}

func (p *pipeline) AddWaitStep(step schema.WaitStep) {
	p.Steps = append(p.Steps, waitStepToGroupStep(step))
}

func (p *pipeline) AddBlockStep(step schema.BlockStep) {
	p.Steps = append(p.Steps, blockStepToGroupStep(step))
}

func (p *pipeline) AddInputStep(step schema.InputStep) {
	p.Steps = append(p.Steps, inputStepToGroupStep(step))
}

func (p *pipeline) AddTriggerStep(step schema.TriggerStep) {
	p.Steps = append(p.Steps, triggerStepToGroupStep(step))
}

func (p *pipeline) AddGroupStep(step GroupStep) {
	p.Steps = append(p.Steps, groupStepToPipelineStep(step))
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
