package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var block = schema_types.NewField().
	Name("block").
	Description("The label for this block step.").
	String()

var blockedState = schema_types.NewField().
	Name("blocked_state").
	Description("The state that the build is set to when the build is blocked by this block step. The default is passed. When the blocked_state of a block step is set to failed, the step that triggered it will be stuck in the running state until it is manually unblocked. Default: passed Values: passed, failed, running").
	Enum("passed", "failed", "running")

var blockedStateField = schema_types.NewField().
	Name("blocked_state").
	Description("The state that the build is set to when the build is blocked by this block step. The default is passed. When the blocked_state of a block step is set to failed, the step that triggered it will be stuck in the running state until it is manually unblocked. Default: passed Values: passed, failed, running").
	FieldRef(&blockedState)

var blockStep = Step{
	Name:        "block",
	Description: "A block step is used to pause the execution of a build and wait on a team member to unblock it using the web or the API.",
	Fields: []schema_types.Field{
		allowDependencyFailure,
		block,
		blockedStateField,
		branches,
		dependsOn,
		fieldsField,
		ifField,
		key,
		prompt,
	},
}
