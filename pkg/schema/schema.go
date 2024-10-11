package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var PipelinesSchema = PipelineSchema{
	Name: "Buildkite Pipeline Schema",
	Types: []schema_types.Field{
		build,
		blockedState,
		fields,
		retryOptions,
		selectInputOption,
		selectInput,
		textInput,
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
