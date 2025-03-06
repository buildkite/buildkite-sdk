package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

type GroupStep struct {
	DependsOn              DependsOn      `json:"depends_on,omitempty"`
	Group                  *string        `json:"group,omitempty"`
	If                     *string        `json:"if,omitempty"`
	Key                    *string        `json:"key,omitempty"`
	ID                     *string        `json:"id,omitempty"`
	Identifier             *string        `json:"identifier,omitempty"`
	Label                  *string        `json:"label,omitempty"`
	Name                   *string        `json:"name,omitempty"`
	AllowDependencyFailure *bool          `json:"allow_dependency_failure,omitempty"`
	Skip                   *bool          `json:"skip,omitempty"`
	Notify                 []StepNotify   `json:"notify,omitempty"`
	Steps                  []PipelineStep `json:"steps,omitempty"`
}

func (step GroupStep) toPipelineStep() *PipelineStep {
	groupName := Value("~")
	if step.Group != nil {
		groupName = step.Group
	}

	notify := make([]schema.BlockStepNotify, len(step.Notify))
	for i, val := range step.Notify {
		notify[i] = *val.toSchema()
	}

	groupStep := &PipelineStep{
		Group:      groupName,
		If:         step.If,
		ID:         step.ID,
		Identifier: step.Identifier,
		Key:        step.Key,
		Label:      step.Label,
		Name:       step.Name,
		Notify:     notify,
		Steps:      step.Steps,
	}

	if step.AllowDependencyFailure != nil {
		groupStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.DependsOn != nil {
		groupStep.DependsOn = step.DependsOn.toSchema()
	}

	if step.Skip != nil {
		groupStep.Skip = &schema.Skip{
			Bool: step.Skip,
		}
	}

	return groupStep
}
