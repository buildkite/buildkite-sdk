package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

type mockPlugin struct{}

func newCommandStep() *bk.Command {
	return &bk.Command{
		Agents: map[string]string{
			"queue": "mac",
		},
		AllowDependencyFailure: false,
		ArtifactPaths: []string{
			"artifact.json",
		},
		Branches:             "branch",
		CancelOnBuildFailing: true,
		Commands:             []string{"run.sh"},
		Concurrency:          1,
		ConcurrencyGroup:     "build",
		DependsOn:            []string{"test"},
		Env: map[string]string{
			"FOO": "BAR",
		},
		If:          "branch == \"main\"",
		Key:         "build",
		Label:       "Build my thing",
		Matrix:      []string{"mac", "windows"},
		Parallelism: 1,
		Plugins: []map[string]interface{}{
			{
				"plugin": mockPlugin{},
			},
		},
		Priority: 1,
		Retry: &bk.Retry{
			Automatic: true,
		},
		SoftFail:         true,
		TimeoutInMinutes: 10,
	}
}

func TestCommandStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render a command step", func(t *testing.T) {
		step := newCommandStep()
		pipeline := bk.NewStepBuilder().AddCommand(step)
		validator.ValidatePipeline(t, pipeline)
	})
}
