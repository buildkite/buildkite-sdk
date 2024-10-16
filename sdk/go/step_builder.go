// This file is auto-generated. Do not edit.
package buildkite
import (
    "encoding/json"
    "os"
)
type stepBuilder struct {
	Steps []interface{} `json:"steps"`
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{}
}
func (s *stepBuilder) AddBlock(step *Block) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) AddCommand(step *Command) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) AddGroup(step *Group) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) AddInput(step *Input) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) AddTrigger(step *Trigger) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) AddWait(step *Wait) *stepBuilder {
    s.Steps = append(s.Steps, step)
    return s
}
func (s *stepBuilder) Print() error {
	jsonBytes, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("pipeline.json", jsonBytes, os.ModePerm)
}