package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

// Blocked State
type blockedState struct {
	Failed  schema.BlockedState
	Passed  schema.BlockedState
	Running schema.BlockedState
}

var BlockedState = blockedState{
	Failed:  schema.Failed,
	Passed:  schema.Passed,
	Running: schema.Running,
}

type BlockStep struct {
	AllowDependencyFailure *bool
	Block                  *string
	BlockedState           *schema.BlockedState
	Branches               []string
	DependsOn              DependsOn
	Fields                 []Field
	ID                     *string
	Identifier             *string
	If                     *string
	Key                    *string
	Label                  *string
	Name                   *string
	Prompt                 *string
}

func (step BlockStep) toPipelineStep() *PipelineStep {
	fields := make([]schema.Field, len(step.Fields))
	for i, field := range step.Fields {
		fields[i] = field.toSchema()
	}

	blockStep := &PipelineStep{
		Block: &schema.Block{
			String: step.Block,
		},
		BlockedState: step.BlockedState,
		Fields:       fields,
		ID:           step.ID,
		Identifier:   step.Identifier,
		If:           step.If,
		Key:          step.Key,
		Label:        step.Label,
		Name:         step.Name,
		Prompt:       step.Prompt,
	}

	if step.AllowDependencyFailure != nil {
		blockStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.Branches != nil {
		blockStep.Branches = &schema.Branches{
			StringArray: step.Branches,
		}
	}

	if step.DependsOn != nil {
		blockStep.DependsOn = step.DependsOn.toSchema()
	}

	return blockStep
}
