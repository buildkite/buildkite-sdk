package buildkite

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type PipelineStep interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type Pipeline struct {
	Steps []PipelineStep `json:"steps" yaml:"steps"`
}

func (p *Pipeline) AddCommandStep(step CommandStep) {
	p.Steps = append(p.Steps, &CommandUnion{
		CommandStep: &step,
	})
}

func (p *Pipeline) AddWaitStep(step WaitStep) {
	p.Steps = append(p.Steps, &Label{
		WaitStep: &step,
	})
}

func (p *Pipeline) AddInputStep(step InputStep) {
	p.Steps = append(p.Steps, &Input{
		InputStep: &step,
	})
}

func (p *Pipeline) AddTriggerStep(step TriggerStep) {
	p.Steps = append(p.Steps, &Trigger{
		TriggerStep: &step,
	})
}

func (p *Pipeline) AddBlockStep(step BlockStep) {
	p.Steps = append(p.Steps, &Block{
		BlockStep: &step,
	})
}

func (p *Pipeline) AddGroupStep(step GroupStepClass) {
	p.Steps = append(p.Steps, &SchemaStep{
		GroupStepClass: &step,
	})
}

func (p *Pipeline) ToJSON() (string, error) {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *Pipeline) ToYAML() (string, error) {
	data, err := yaml.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
