package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

var async = schema_types.NewField().
	Name("async").
	Description(`If set to true the step will immediately continue, regardless of the success of the triggered build. If set to false the step will wait for the triggered build to complete and continue only if the triggered build passed.
Note that when async is set to true, as long as the triggered build starts, the original pipeline will show that as successful. The original pipeline does not get updated after subsequent steps or after the triggered build completes.`).
	Boolean()

var build = schema_types.NewField().
	Name("build").
	Description("An optional map of attributes for the triggered build. Available attributes: branch, commit, env, message, meta_data").
	Object([]schema_types.Field{
		schema_types.NewField().
			Name("message").
			Description("The message for the build. Supports emoji.").
			String(),
		schema_types.NewField().
			Name("commit").
			Description("The commit hash for the build.").
			String(),
		schema_types.NewField().
			Name("branch").
			Description("The branch for the build.").
			String(),
		schema_types.NewField().
			Name("meta_data").
			Description("A map of meta-data for the build.").
			StringMap(),
		schema_types.NewField().
			Name("env").
			Description("A map of environment variables for the build.").
			StringMap(),
	})

var buildField = schema_types.
	NewField().
	Name("build").
	Description("An optional map of attributes for the triggered build. Available attributes: branch, commit, env, message, meta_data").
	FieldRef(&build)

var trigger = schema_types.NewField().
	Name("trigger").
	Description("The slug of the pipeline to create a build. You can find it in the URL of your pipeline, and it corresponds to the name of the pipeline, converted to kebab-case.").
	Required().
	String()

var triggerSoftFail = schema_types.NewField().
	Name("soft_fail").
	Description("When true, failure of the triggered build will not cause the triggering build to fail.").
	Boolean()

var triggerStep = Step{
	Name:        "trigger",
	Description: "A trigger step creates a build on another pipeline.",
	Fields: []schema_types.Field{
		allowDependencyFailure,
		async,
		branches,
		buildField,
		dependsOn,
		ifField,
		label,
		skip,
		triggerSoftFail,
		trigger,
	},
}
