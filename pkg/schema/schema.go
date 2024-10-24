package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var PipelinesSchema = PipelineSchema{
	Version: "0.0.1",
	Name:    "Buildkite Pipeline Schema",
	Types: []schema_types.Field{
		build,
		blockedState,
		fields,
		retryOptions,
		selectInputOption,
		selectInput,
		textInput,
		steps,
		wait,
	},
	Steps: []Step{
		blockStep,
		commandStep,
		groupStep,
		inputStep,
		triggerStep,
		waitStep,
	},
	Environment: environmentVariables,
}
