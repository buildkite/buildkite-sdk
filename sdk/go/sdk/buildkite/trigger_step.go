package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

// Build
type Build schema.Build

type TriggerStep struct {
	AllowDependencyFailure *bool
	Async                  *bool
	Branches               []string
	Build                  *Build
	DependsOn              DependsOn
	ID                     *string
	Identifier             *string
	If                     *string
	Key                    *string
	Label                  *string
	Name                   *string
	Skip                   *bool
	SoftFail               SoftFail
	Trigger                string
}

func (step TriggerStep) toPipelineStep() *PipelineStep {
	triggerStep := &PipelineStep{
		ID:         step.ID,
		Identifier: step.Identifier,
		If:         step.If,
		Key:        step.Key,
		Label:      step.Label,
		Name:       step.Name,
		Trigger: &schema.Trigger{
			String: &step.Trigger,
		},
	}

	if step.AllowDependencyFailure != nil {
		triggerStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.Async != nil {
		triggerStep.Async = &schema.AllowDependencyFailureUnion{
			Bool: step.Async,
		}
	}

	if step.Branches != nil {
		triggerStep.Branches = &schema.Branches{
			StringArray: step.Branches,
		}
	}

	if step.Build != nil {
		triggerStep.Build = &schema.Build{
			Branch:   step.Build.Branch,
			Commit:   step.Build.Commit,
			Env:      step.Build.Env,
			Message:  step.Build.Message,
			MetaData: step.Build.MetaData,
		}
	}

	if step.DependsOn != nil {
		triggerStep.DependsOn = step.DependsOn.toSchema()
	}

	if step.Skip != nil {
		triggerStep.Skip = &schema.Skip{
			Bool: step.Skip,
		}
	}

	if step.SoftFail != nil {
		triggerStep.SoftFail = step.SoftFail.toSchema()
	}

	return triggerStep
}
