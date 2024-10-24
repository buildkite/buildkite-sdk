package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func newTriggerStep() *bk.Trigger {
	return &bk.Trigger{
		AllowDependencyFailure: true,
		Async:                  false,
		Branches:               "branch",
		Build: &bk.Build{
			Message:  "message",
			Commit:   "abc123",
			Branch:   "branch",
			MetaData: map[string]string{"foo": "bar"},
			Env:      map[string]string{"FOO": "BAR"},
		},
		DependsOn: []string{"build"},
		If:        "branch == \"branch\"",
		Label:     "my trigger step",
		Skip:      false,
		SoftFail:  true,
		Trigger:   "some-pipeline",
	}
}

func TestTriggerStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render a trigger step", func(t *testing.T) {
		step := newTriggerStep()
		pipeline := bk.NewStepBuilder().AddTrigger(step)
		validator.ValidatePipeline(t, pipeline)
	})

	t.Run("should render a trigger step when `build` argument is missing", func(t *testing.T) {
		pipeline := bk.NewStepBuilder().AddTrigger(&bk.Trigger{
			AllowDependencyFailure: true,
			Async:                  false,
			Branches:               "branch",
			DependsOn:              []string{"build"},
			If:                     "branch == \"branch\"",
			Label:                  "my trigger step",
			Skip:                   false,
			SoftFail:               true,
			Trigger:                "some-pipeline",
		})

		validator.ValidatePipeline(t, pipeline)
	})
}
