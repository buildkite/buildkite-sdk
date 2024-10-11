package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var allowDependencyFailure = schema_types.NewField().
	Name("allow_dependency_failure").
	Description("Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.").
	Boolean()

var branches = schema_types.NewField().
	Name("branches").
	Description("The branch pattern defining which branches will include this block step in their builds.").
	String()

var dependsOn = schema_types.NewField().
	Name("depends_on").
	Description("A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.").
	StringArray()

var ifField = schema_types.NewField().
	Name("if").
	Description("A boolean expression that omits the step when false. See Using conditionals for supported expressions.").
	String()

var key = schema_types.NewField().
	Name("key").
	Description("A unique string to identify the block step.").
	String()

var label = schema_types.NewField().
	Name("label").
	Description("The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.").
	String()

var prompt = schema_types.NewField().
	Name("prompt").
	Description("The instructional message displayed in the dialog box when the unblock step is activated.").
	String()

var skip = schema_types.NewField().
	Name("skip").
	Description("Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.").
	Boolean()
