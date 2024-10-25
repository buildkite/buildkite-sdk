package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var group = schema_types.NewField().
	Name("group").
	Description("Name of the group in the UI. In YAML, if you don't want a label, pass a `~`. Can also be provided in the `label` attribute if `null` is provided to the `group` attribute.").
	String()

var notify = schema_types.NewField().
	Name("notify").
	Description("Allows you to trigger build notifications to different services. You can also choose to conditionally send notifications based on pipeline events.").
	StringArray()

var steps = schema_types.NewField().
	Name("steps").
	Description("A list of steps in the group; at least 1 step is required. Allowed step types: wait, trigger, command/commands, block, input.").
	UnionArray(
		"steps",
		blockStep.ToObjectField(),
		commandStep.ToObjectField(),
		inputStep.ToObjectField(),
		triggerStep.ToObjectField(),
		waitStep.ToObjectField(),
	)

var stepsField = schema_types.NewField().
	Name("steps").
	Description("A list of steps in the group; at least 1 step is required. Allowed step types: wait, trigger, command/commands, block, input.").
	FieldRef(&steps)

var groupStep = Step{
	Name:        "group",
	Description: "A group step can contain various sub-steps, and display them in a single logical group on the Build page.",
	Fields: []schema_types.Field{
		allowDependencyFailure,
		dependsOn,
		group,
		ifField,
		key,
		label,
		notify,
		skip,
		stepsField,
	},
}
