package buildkite

import (
	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"
)

type PipelineStep struct {
	AllowDependencyFailure *schema.AllowDependencyFailureUnion `json:"allow_dependency_failure,omitempty"`
	Block                  *schema.Block                       `json:"block,omitempty"`
	BlockedState           *schema.BlockedState                `json:"blocked_state,omitempty"`
	Branches               *schema.Branches                    `json:"branches,omitempty"`
	DependsOn              *schema.DependsOn                   `json:"depends_on,omitempty"`
	Fields                 []schema.Field                      `json:"fields,omitempty"`
	ID                     *string                             `json:"id,omitempty"`
	Identifier             *string                             `json:"identifier,omitempty"`
	If                     *string                             `json:"if,omitempty"`
	Key                    *string                             `json:"key,omitempty"`
	Label                  *string                             `json:"label,omitempty"`
	Name                   *string                             `json:"name,omitempty"`
	Prompt                 *string                             `json:"prompt,omitempty"`
	Type                   *schema.BlockStepType               `json:"type,omitempty"`
	Input                  *schema.Input                       `json:"input,omitempty"`
	Agents                 *schema.Agents                      `json:"agents,omitempty"`
	ArtifactPaths          []string                            `json:"artifact_paths,omitempty"`
	Cache                  *schema.Cache                       `json:"cache,omitempty"`
	CancelOnBuildFailing   *schema.AllowDependencyFailureUnion `json:"cancel_on_build_failing,omitempty"`
	Command                *schema.CommandUnion                `json:"command,omitempty"`
	Commands               *schema.CommandUnion                `json:"commands,omitempty"`
	Concurrency            *int64                              `json:"concurrency,omitempty"`
	ConcurrencyGroup       *string                             `json:"concurrency_group,omitempty"`
	ConcurrencyMethod      *schema.ConcurrencyMethod           `json:"concurrency_method,omitempty"`
	Env                    map[string]interface{}              `json:"env,omitempty"`
	Matrix                 *schema.MatrixUnion                 `json:"matrix,omitempty"`
	Notify                 []schema.BlockStepNotify            `json:"notify,omitempty"`
	Parallelism            *int64                              `json:"parallelism,omitempty"`
	Plugins                *schema.Plugins                     `json:"plugins,omitempty"`
	Priority               *int64                              `json:"priority,omitempty"`
	Retry                  *schema.Retry                       `json:"retry,omitempty"`
	Signature              *schema.Signature                   `json:"signature,omitempty"`
	Skip                   *schema.Skip                        `json:"skip,omitempty"`
	SoftFail               *schema.SoftFail                    `json:"soft_fail,omitempty"`
	TimeoutInMinutes       *int64                              `json:"timeout_in_minutes,omitempty"`
	Script                 *schema.CommandStep                 `json:"script,omitempty"`
	ContinueOnFailure      *schema.AllowDependencyFailureUnion `json:"continue_on_failure,omitempty"`
	Wait                   *schema.Label                       `json:"wait,omitempty"`
	Waiter                 *schema.WaitStep                    `json:"waiter,omitempty"`
	Async                  *schema.AllowDependencyFailureUnion `json:"async,omitempty"`
	Build                  *schema.Build                       `json:"build,omitempty"`
	Trigger                *schema.Trigger                     `json:"trigger,omitempty"`
	Group                  *string                             `json:"group,omitempty"`
	Steps                  []PipelineStep                      `json:"steps,omitempty"`
}

func commandStepToGroupStep(step schema.CommandStep) *PipelineStep {
	var notify []schema.BlockStepNotify
	for _, val := range step.Notify {
		notify = append(notify, schema.BlockStepNotify{
			FluffyBuildNotify: &schema.FluffyBuildNotify{
				BasecampCampfire: val.NotifyClass.BasecampCampfire,
				If:               val.NotifyClass.If,
				Slack: &schema.IndecentSlack{
					TentacledSlack: (*schema.TentacledSlack)(val.NotifyClass.Slack.FluffySlack),
				},
				GithubCheck:        val.NotifyClass.GithubCheck,
				GithubCommitStatus: (*schema.TentacledGithubCommitStatus)(val.NotifyClass.GithubCommitStatus),
			},
		})
	}

	return &PipelineStep{
		Agents:                 step.Agents,
		AllowDependencyFailure: step.AllowDependencyFailure,
		ArtifactPaths:          step.ArtifactPaths,
		Branches:               step.Branches,
		Cache:                  step.Cache,
		CancelOnBuildFailing:   step.CancelOnBuildFailing,
		Command:                step.Command,
		Commands:               step.Commands,
		Concurrency:            step.Concurrency,
		ConcurrencyGroup:       step.ConcurrencyGroup,
		ConcurrencyMethod:      step.ConcurrencyMethod,
		DependsOn:              step.DependsOn,
		Env:                    step.Env,
		ID:                     step.ID,
		Identifier:             step.Identifier,
		If:                     step.If,
		Key:                    step.Key,
		Label:                  step.Label,
		Matrix:                 step.Matrix,
		Name:                   step.Name,
		Notify:                 notify,
		Parallelism:            step.Parallelism,
		Plugins:                step.Plugins,
		Priority:               step.Priority,
		Retry:                  step.Retry,
		Signature:              step.Signature,
		Skip:                   step.Skip,
		SoftFail:               step.SoftFail,
		TimeoutInMinutes:       step.TimeoutInMinutes,
		Type:                   (*schema.BlockStepType)(step.Type),
	}
}

func blockStepToGroupStep(step schema.BlockStep) *PipelineStep {
	return &PipelineStep{
		AllowDependencyFailure: step.AllowDependencyFailure,
		Block: &schema.Block{
			String: step.Block,
		},
		BlockedState: step.BlockedState,
		Branches:     step.Branches,
		DependsOn:    step.DependsOn,
		Fields:       step.Fields,
		ID:           step.ID,
		Identifier:   step.Identifier,
		If:           step.If,
		Key:          step.Key,
		Label:        step.Label,
		Name:         step.Name,
		Prompt:       step.Prompt,
		Type:         (*schema.BlockStepType)(step.Type),
	}
}

func waitStepToGroupStep(step schema.WaitStep) *PipelineStep {
	return &PipelineStep{
		AllowDependencyFailure: step.AllowDependencyFailure,
		Branches:               step.Branches,
		ContinueOnFailure:      step.ContinueOnFailure,
		DependsOn:              step.DependsOn,
		ID:                     step.ID,
		Identifier:             step.Identifier,
		If:                     step.If,
		Key:                    step.Key,
		Label:                  step.Label,
		Name:                   step.Name,
		Type:                   (*schema.BlockStepType)(step.Type),
		Wait: &schema.Label{
			String: step.Wait,
		},
	}
}

func inputStepToGroupStep(step schema.InputStep) *PipelineStep {
	return &PipelineStep{
		AllowDependencyFailure: step.AllowDependencyFailure,
		Branches:               step.Branches,
		DependsOn:              step.DependsOn,
		Fields:                 step.Fields,
		ID:                     step.ID,
		Identifier:             step.Identifier,
		If:                     step.If,
		Input: &schema.Input{
			String: step.Input,
		},
		Key:    step.Key,
		Label:  step.Label,
		Name:   step.Name,
		Prompt: step.Prompt,
		Type:   (*schema.BlockStepType)(step.Type),
	}
}

func triggerStepToGroupStep(step schema.TriggerStep) *PipelineStep {
	return &PipelineStep{
		AllowDependencyFailure: step.AllowDependencyFailure,
		Async:                  step.Async,
		Branches:               step.Branches,
		Build:                  step.Build,
		DependsOn:              step.DependsOn,
		ID:                     step.ID,
		Identifier:             step.Identifier,
		If:                     step.If,
		Key:                    step.Key,
		Label:                  step.Label,
		Name:                   step.Name,
		Skip:                   step.Skip,
		SoftFail:               step.SoftFail,
		Trigger: &schema.Trigger{
			String: &step.Trigger,
		},
		Type: (*schema.BlockStepType)(step.Type),
	}
}

type GroupStep struct {
	DependsOn              *schema.DependsOn        `json:"depends_on,omitempty"`
	Group                  *string                  `json:"group,omitempty"`
	If                     *string                  `json:"if,omitempty"`
	Key                    *string                  `json:"key,omitempty"`
	ID                     *string                  `json:"id,omitempty"`
	Identifier             *string                  `json:"identifier,omitempty"`
	Label                  *string                  `json:"label,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	AllowDependencyFailure *bool                    `json:"allow_dependency_failure,omitempty"`
	Skip                   *bool                    `json:"skip,omitempty"`
	Notify                 []schema.BlockStepNotify `json:"notify,omitempty"`
	Steps                  []PipelineStep           `json:"steps,omitempty"`
}

func groupStepToPipelineStep(step GroupStep) *PipelineStep {
	return &PipelineStep{
		DependsOn:  step.DependsOn,
		Group:      step.Group,
		If:         step.If,
		ID:         step.ID,
		Identifier: step.Identifier,
		Key:        step.Key,
		Label:      step.Label,
		Name:       step.Name,
		AllowDependencyFailure: &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		},
		Skip: &schema.Skip{
			Bool: step.Skip,
		},
		Notify: step.Notify,
		Steps:  step.Steps,
	}
}
