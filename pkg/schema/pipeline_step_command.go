package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var agents = schema_types.NewField().
	Name("agents").
	Description("A map of agent tag keys to values to target specific agents for this step.").
	StringMap()

var artifactPaths = schema_types.NewField().
	Name("artifact_paths").
	Description("The glob path or paths of artifacts to upload from this step.").
	StringArray()

var cancelOnBuildFailing = schema_types.NewField().
	Name("cancel_on_build_failing").
	Description("Setting this attribute to true cancels the job as soon as the build is marked as failing.").
	Boolean()

var commands = schema_types.NewField().
	Name("commands").
	Description("The shell command to run during this step.").
	StringArray()

var concurrency = schema_types.NewField().
	Name("concurrency").
	Description("The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.").
	Number()

var concurrencyGroup = schema_types.NewField().
	Name("concurrency_group").
	Description("A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.").
	String()

var environment = schema_types.NewField().
	Name("env").
	Description("A map of environment variables for this step.").
	StringMap()

// TODO: Fully implement this
var matrix = schema_types.NewField().
	Name("matrix").
	Description("An array of values to be used in the matrix expansion.").
	StringArray()

var parallelism = schema_types.NewField().
	Name("parallelism").
	Description("The number of parallel jobs that will be created based on this step.").
	Number()

// TODO: Look into how we could eventually type this out.
var plugins = schema_types.NewField().
	Name("plugins").
	Description("An array of plugins for this step.").
	AnyMapArray()

var priority = schema_types.NewField().
	Name("priority").
	Description("Adjust the priority for a specific job, as a positive or negative integer.").
	Number()

var retryOptionsField = schema_types.NewField().
	Name("retry").
	Description("The conditions for retrying this step.").
	FieldRef(&retryOptions)

var softFail = schema_types.NewField().
	Name("soft_fail").
	Description("Make all exit statuses soft-fail.").
	Boolean()

var timeoutInMinutes = schema_types.NewField().
	Name("timeout_in_minutes").
	Description("The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.").
	Number()

var commandStep = Step{
	Name:        "command",
	Description: "A command step runs one or more shell commands on one or more agents.",
	Fields: []schema_types.Field{
		agents,
		allowDependencyFailure,
		artifactPaths,
		branches,
		cancelOnBuildFailing,
		commands,
		concurrency,
		concurrencyGroup,
		dependsOn,
		environment,
		ifField,
		key,
		label,
		matrix,
		parallelism,
		plugins,
		priority,
		retryOptionsField,
		skip,
		softFail,
		timeoutInMinutes,
	},
}
