package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func newBlockStep() *bk.Block {
	failedState := bk.FAILED

	return &bk.Block{
		AllowDependencyFailure: false,
		Block:                  "my block step",
		BlockedState:           &failedState,
		Branches:               "branch",
		DependsOn:              []string{"build"},
		Fields: []bk.Fields{
			&bk.TextInput{
				Key:  "name",
				Text: "What is your name?",
			},
		},
		If:     "branch == \"main\"",
		Key:    "details",
		Prompt: "Fill out details",
	}
}

func TestBlockStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render a block step", func(t *testing.T) {
		step := newBlockStep()
		pipeline := bk.NewStepBuilder().AddBlock(step)
		validator.ValidatePipeline(t, pipeline)
	})
}
