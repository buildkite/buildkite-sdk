package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var wait = schema_types.NewField().
	Name("wait").
	Description("When providing options for the wait step, you will need to set this value to \"~\".").
	Required().
	String()

var continueOnFailure = schema_types.NewField().
	Name("continue_on_failure").
	Description("Run the next step, even if the previous step has failed.").
	Boolean()

var waitStep = Step{
	Name:        "wait",
	Description: "A wait step waits for all previous steps to have successfully completed before allowing following jobs to continue.",
	Fields: []schema_types.Field{
		allowDependencyFailure,
		continueOnFailure,
		dependsOn,
		ifField,
		wait,
	},
}
