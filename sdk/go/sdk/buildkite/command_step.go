package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

type CommandStep struct {
	Agents                 map[string]interface{}
	AllowDependencyFailure *bool
	ArtifactPaths          []string
	Branches               []string
	Cache                  Cache
	CancelOnBuildFailing   *bool
	Commands               []string
	Concurrency            *int64
	ConcurrencyGroup       *string
	ConcurrencyMethod      *schema.ConcurrencyMethod
	DependsOn              DependsOn
	Env                    map[string]interface{}
	ID                     *string
	Identifier             *string
	If                     *string
	Key                    *string
	Label                  *string
	Matrix                 Matrix
	Name                   *string
	Notify                 []CommandNotify
	Parallelism            *int64
	Plugins                map[string]interface{}
	Priority               *int64
	Retry                  Retry
	Signature              *Signature
	Skip                   *bool
	SoftFail               SoftFail
	TimeoutInMinutes       *int64
}

func (step CommandStep) toPipelineStep() *PipelineStep {
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

	commandStep := &PipelineStep{
		Commands: &schema.CommandUnion{
			StringArray: step.Commands,
		},

		ArtifactPaths:     step.ArtifactPaths,
		Concurrency:       step.Concurrency,
		ConcurrencyGroup:  step.ConcurrencyGroup,
		ConcurrencyMethod: step.ConcurrencyMethod,
		Env:               step.Env,
		ID:                step.ID,
		Identifier:        step.Identifier,
		If:                step.If,
		Key:               step.Key,
		Label:             step.Label,
		Name:              step.Name,
		Notify:            notify,
		Parallelism:       step.Parallelism,
		Priority:          step.Priority,
		TimeoutInMinutes:  step.TimeoutInMinutes,
	}

	if step.Agents != nil {
		commandStep.Agents = &schema.Agents{
			AnythingMap: step.Agents,
		}
	}

	if step.AllowDependencyFailure != nil {
		commandStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.Branches != nil {
		commandStep.Branches = &schema.Branches{
			StringArray: step.Branches,
		}
	}

	if step.Cache != nil {
		commandStep.Cache = step.Cache.toSchemaCache()
	}

	if step.CancelOnBuildFailing != nil {
		commandStep.CancelOnBuildFailing = &schema.AllowDependencyFailureUnion{
			Bool: step.CancelOnBuildFailing,
		}
	}

	if step.DependsOn != nil {
		commandStep.DependsOn = step.DependsOn.toSchema()
	}

	if step.Matrix != nil {
		commandStep.Matrix = step.Matrix.toSchema()
	}

	if step.Plugins != nil {
		commandStep.Plugins = &schema.Plugins{
			AnythingMap: step.Plugins,
		}
	}

	if step.Retry != nil {
		commandStep.Retry = step.Retry.toSchema()
	}

	if step.Signature != nil {
		commandStep.Signature = step.Signature.toSchema()
	}

	if step.Skip != nil {
		commandStep.Skip = &schema.Skip{
			Bool: step.Skip,
		}
	}

	if step.SoftFail != nil {
		commandStep.SoftFail = step.SoftFail.toSchema()
	}

	return commandStep
}
