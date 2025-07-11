package buildkite

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/json2yaml"
)

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

type Pipeline struct {
	Agents map[string]any  `json:"agents,omitempty"`
	Env    map[string]any  `json:"env,omitempty"`
	Notify *PipelineNotify `json:"notify,omitempty"`
	Steps  []*PipelineStep `json:"steps"`
}

type pipelineStep interface {
	ToPipelineStep() *PipelineStep
}

func (p *Pipeline) AddStep(step pipelineStep) {
	p.Steps = append(p.Steps, step.ToPipelineStep())
}

func (p *Pipeline) AddAgent(key string, value any) {
	if p.Agents == nil {
		p.Agents = make(map[string]any)
	}

	p.Agents[key] = value
}

func (p *Pipeline) AddEnvironmentVariable(key string, value any) {
	if p.Env == nil {
		p.Env = make(map[string]any)
	}

	p.Env[key] = value
}

func (p *Pipeline) AddNotify(notify *PipelineNotify) {
	p.Notify = notify
}

func (p *Pipeline) ToJSON() (string, error) {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *Pipeline) ToYAML() (string, error) {
	data, err := p.ToJSON()
	if err != nil {
		return "", err
	}

	var output strings.Builder
	input := strings.NewReader(data)
	if err := json2yaml.Convert(&output, input); err != nil {
		return "", fmt.Errorf("converting JSON to YAML: %v", err)
	}

	return output.String(), nil
}
