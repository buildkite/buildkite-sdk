package sdk_tests_go

import (
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func newInputStep() *bk.Input {
	return &bk.Input{
		Input:                  "My cool input step",
		AllowDependencyFailure: true,
		Branches:               "branch",
		DependsOn:              []string{"build"},
		If:                     "branch == \"branch\"",
		Key:                    "input-thing",
		Prompt:                 "Tell me about yourself",
		Fields: []bk.Fields{
			&bk.TextInput{
				Key:  "name",
				Text: "What is your name?",
			},
		},
	}
}

func TestInputStep(t *testing.T) {
	validator := newPipelineSchemaValidator(t)

	t.Run("should render an input step", func(t *testing.T) {
		step := newInputStep()
		pipeline := bk.NewStepBuilder().AddInput(step)
		validator.ValidatePipeline(t, pipeline)
	})

	t.Run("should render a text and select input", func(t *testing.T) {
		pipeline := bk.NewStepBuilder().AddInput(&bk.Input{
			Input: "My cool input step",
			Fields: []bk.Fields{
				&bk.TextInput{
					Key:  "name",
					Text: "What is your name?",
				},
				&bk.SelectInputAttribute{
					Key: "type",
					Options: []bk.SelectInputOption{
						{
							Label: "Customer",
							Value: "customer",
						},
						{
							Label: "Employee",
							Value: "employee",
						},
					},
				},
			},
		})

		validator.ValidatePipeline(t, pipeline)
	})
}
