package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func TestGroupStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render a valid group step", func(t *testing.T) {
		pipeline := bk.NewStepBuilder().AddGroup(&bk.Group{
			AllowDependencyFailure: false,
			DependsOn:              []string{"build"},
			Group:                  "My great group",
			If:                     "brach == \"main\"",
			Key:                    "group",
			Label:                  "cool",
			Notify:                 []string{"github_commit_status"},
			Skip:                   false,
			Steps: []bk.Steps{
				newBlockStep(),
				newCommandStep(),
				newInputStep(),
				newTriggerStep(),
				newWaitStep(),
			},
		})

		validator.ValidatePipeline(t, pipeline)
	})
}
