package buildkite

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func NewPipeline() *pipeline {
	return &pipeline{}
}

type pipeline struct {
	Agents map[string]any  `json:"agents,omitempty"`
	Env    map[string]any  `json:"env,omitempty"`
	Notify PipelineNotify  `json:"notify,omitempty"`
	Steps  []*PipelineStep `json:"steps"`
}

type pipelineStep interface {
	toPipelineStep() *PipelineStep
}

func (p *pipeline) AddStep(step pipelineStep) {
	p.Steps = append(p.Steps, step.toPipelineStep())
}

func (p *pipeline) AddAgent(key string, value any) {
	if p.Agents == nil {
		p.Agents = make(map[string]any)
	}

	p.Agents[key] = value
}

func (p *pipeline) AddEnvironmentVariable(key string, value any) {
	if p.Env == nil {
		p.Env = make(map[string]any)
	}

	p.Env[key] = value
}

func (p *pipeline) AddNotify(notify PipelineNotify) {
	p.Notify = notify
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
