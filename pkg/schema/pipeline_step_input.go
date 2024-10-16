package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var input = schema_types.NewField().
	Name("input").
	Description("The label for this input step.").
	String()

var fieldsField = schema_types.NewField().
	Name("fields").
	Description("An input step is used to collect information from a user.").
	FieldRef(&fields)

var inputStep = Step{
	Name:        "input",
	Description: "An input step is used to collect information from a user.",
	Fields: []schema_types.Field{
		allowDependencyFailure,
		branches,
		dependsOn,
		fieldsField,
		ifField,
		input,
		key,
		prompt,
	},
}
