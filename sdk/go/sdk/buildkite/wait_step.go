package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

type WaitStep struct {
	AllowDependencyFailure *bool
	Branches               []string
	ContinueOnFailure      *bool
	DependsOn              DependsOn
	ID                     *string
	Identifier             *string
	If                     *string
	Key                    *string
	Label                  *string
	Name                   *string
	Wait                   *string
}

func (step WaitStep) toPipelineStep() *PipelineStep {
	waitStep := &PipelineStep{
		ID:         step.ID,
		Identifier: step.Identifier,
		If:         step.If,
		Key:        step.Key,
		Label:      step.Label,
		Name:       step.Name,
		Wait: &schema.Label{
			String: step.Wait,
		},
	}

	if step.AllowDependencyFailure != nil {
		waitStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.Branches != nil {
		waitStep.Branches = &schema.Branches{
			StringArray: step.Branches,
		}
	}

	if step.ContinueOnFailure != nil {
		waitStep.ContinueOnFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.ContinueOnFailure,
		}
	}

	if step.DependsOn != nil {
		waitStep.DependsOn = step.DependsOn.toSchema()
	}

	return waitStep
}
