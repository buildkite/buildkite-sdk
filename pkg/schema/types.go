package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var retryOptions = schema_types.NewField().
	Name("retry").
	Description("The conditions for retrying this step.").
	Object([]schema_types.Field{
		schema_types.NewField().
			Name("automatic").
			Description("Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.").
			Boolean(),
		schema_types.NewField().
			Name("manual").
			Description("Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.").
			Boolean(),
	})

var textInput = schema_types.NewField().
	Name("text_input").
	Description("A description of a text input for a pipeline step.").
	Object([]schema_types.Field{
		schema_types.NewField().
			Name("key").
			Description("The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.").
			String(),
		schema_types.NewField().
			Name("text").
			Description("The label for the text input.").
			String(),
		schema_types.NewField().
			Name("hint").
			Description("The explanatory text that is shown after the label.").
			String(),
		schema_types.NewField().
			Name("required").
			Description("A boolean value that defines whether the field is required for form submission.").
			Boolean(),
		schema_types.NewField().
			Name("default").
			Description("The value that is pre-filled in the text field.").
			String(),
	})

var selectInputOption = schema_types.NewField().
	Name("select_input_option").
	Description("The conditions for retrying this step.").
	Object([]schema_types.Field{
		schema_types.NewField().
			Name("label").
			Description("The text displayed for the option.").
			String(),
		schema_types.NewField().
			Name("value").
			Description("The value to be stored as meta-data (to be later retrieved using the buildkite-agent meta-data command).").
			String(),
	})

var selectInput = schema_types.NewField().
	Name("select_input_attribute").
	Description("A description of a select input for a pipeline step.").
	Object([]schema_types.Field{
		schema_types.NewField().
			Name("key").
			Description("The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.").
			String(),
		schema_types.NewField().
			Name("select").
			Description("The label for the select input.").
			String(),
		schema_types.NewField().
			Name("hint").
			Description("The explanatory text that is shown after the label.").
			String(),
		schema_types.NewField().
			Name("required").
			Description("A boolean value that defines whether the field is required for form submission.").
			Boolean(),
		schema_types.NewField().
			Name("default").
			Description("The value that is pre-filled in the text field.").
			String(),
		schema_types.NewField().
			Name("multiple").
			Description("A boolean value that defines whether multiple options may be selected. When multiple options are selected, they are delimited in the meta-data field by a line break").
			Boolean(),
		schema_types.NewField().
			Name("options").
			Description("The list of select field options. For 6 or less options they'll be displayed as radio buttons, otherwise they'll be displayed in a dropdown box. If selecting multiple options is permitted the options will be displayed as checkboxes.").
			FieldRef(&selectInputOption),
	})

var fields = schema_types.NewField().
	Name("fields").
	Description("An input step is used to collect information from a user.").
	Union(textInput, selectInput)
