package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func newWaitStep() *bk.Wait {
	return &bk.Wait{
		AllowDependencyFailure: true,
		ContinueOnFailure:      true,
		DependsOn:              []string{"test"},
		If:                     "foo == bar",
	}
}

func TestWaitStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render a wait step", func(t *testing.T) {
		step := newWaitStep()
		pipeline := bk.NewStepBuilder().AddWait(step)
		validator.ValidatePipeline(t, pipeline)
	})

	t.Run("should render a valid wait step even if user provides `wait` argument", func(t *testing.T) {
		val := bk.WaitLabel("alsk")
		pipeline := bk.NewStepBuilder().AddWait(&bk.Wait{
			Wait: &val,
		})

		validator.ValidatePipeline(t, pipeline)
	})

}
