// This file is auto-generated. Do not edit.
package buildkite
import (
    "os"
    "encoding/json"
)
type StepBuilder struct {
	Steps []interface{} `json:"steps"`
}

func NewStepBuilder() *StepBuilder {
	return &StepBuilder{}
}
func (s *StepBuilder) AddBlock(step *Block) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) AddCommand(step *Command) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) AddGroup(step *Group) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) AddInput(step *Input) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) AddTrigger(step *Trigger) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) AddWait(step *Wait) *StepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *StepBuilder) Print() error {
	jsonBytes, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("pipeline.json", jsonBytes, os.ModePerm)
}