package schema

import (
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
)

type environmentVariableDefinition struct {
	Name        string
	Description string
	Typ         schema_types.SchemaType
	Metadata    map[string]string
	Dynamic     bool
}

type EnvironmentVariable struct {
	name        string
	description string
	typ         schema_types.SchemaType
	metadata    map[string]string
	dynamic     bool
}

func (e EnvironmentVariable) GetDefinition() environmentVariableDefinition {
	return environmentVariableDefinition{
		Name:        e.name,
		Description: e.description,
		Typ:         e.typ,
		Metadata:    e.metadata,
		Dynamic:     e.dynamic,
	}
}

func (e *EnvironmentVariable) Description(desc string) *EnvironmentVariable {
	e.description = desc
	return e
}

func (e *EnvironmentVariable) Dynamic() *EnvironmentVariable {
	e.dynamic = true
	return e
}

func (e *EnvironmentVariable) String() EnvironmentVariable {
	e.typ = schema_types.Simple.String()
	return *e
}

func (e *EnvironmentVariable) StringArray(delimiter string) EnvironmentVariable {
	e.typ = schema_types.Array.String()
	e.metadata["delimiter"] = delimiter
	return *e
}

func (e *EnvironmentVariable) Number() EnvironmentVariable {
	e.typ = schema_types.Simple.Number()
	return *e
}

func (e *EnvironmentVariable) Boolean() EnvironmentVariable {
	e.typ = schema_types.Simple.Boolean()
	return *e
}

func NewEnvVar(name string) *EnvironmentVariable {
	return &EnvironmentVariable{
		name:     name,
		metadata: make(map[string]string),
	}
}

var environmentVariables = []EnvironmentVariable{
	NewEnvVar("BUILDKITE_AGENT_ACCESS_TOKEN").
		Description("The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.").
		String(),

	NewEnvVar("BUILDKITE_AGENT_DEBUG").
		Description("The value of the debug [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB").
		Description("The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT").
		Description("The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).").
		Number(),

	NewEnvVar("BUILDKITE_AGENT_ENDPOINT").
		Description("The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.").
		String(),

	NewEnvVar("BUILDKITE_AGENT_EXPERIMENT").
		Description("A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).").
		StringArray(","),

	NewEnvVar("BUILDKITE_AGENT_HEALTH_CHECK_ADDR").
		Description("The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_AGENT_ID").
		Description("The UUID of the agent.").
		String(),

	NewEnvVar("BUILDKITE_AGENT_META_DATA").
		Description("The value of each agent tag. The tag name is appended to the end of the variable name. They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file. The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.").
		Dynamic().
		String(),

	NewEnvVar("BUILDKITE_AGENT_NAME").
		Description("The name of the agent that ran the job.").
		String(),

	NewEnvVar("BUILDKITE_AGENT_PID").
		Description("The process ID of the agent.").
		String(),

	NewEnvVar("BUILDKITE_ARTIFACT_PATHS").
		Description("The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		StringArray(";"),

	NewEnvVar("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION").
		Description("The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.").
		String(),

	NewEnvVar("BUILDKITE_BIN_PATH").
		Description("The path to the directory containing the `buildkite-agent` binary.").
		String(),

	NewEnvVar("BUILDKITE_BRANCH").
		Description("The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.").
		String(),

	NewEnvVar("BUILDKITE_BUILD_CHECKOUT_PATH").
		Description("The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_BUILD_AUTHOR").
		Description("The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.").
		String(),

	NewEnvVar("BUILDKITE_BUILD_AUTHOR_EMAIL").
		Description("The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.").
		String(),

	NewEnvVar("BUILDKITE_BUILD_CREATOR").
		Description(`The name of the user who created the build. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`).
		String(),

	NewEnvVar("BUILDKITE_BUILD_CREATOR_EMAIL").
		Description(`The notification email of the user who created the build. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`).
		String(),

	NewEnvVar("BUILDKITE_BUILD_CREATOR_TEAMS").
		Description(`A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`).
		StringArray(":"),

	NewEnvVar("BUILDKITE_BUILD_ID").
		Description("The UUID of the build.").
		String(),

	NewEnvVar("BUILDKITE_BUILD_NUMBER").
		Description("The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.").
		Number(),

	NewEnvVar("BUILDKITE_BUILD_PATH").
		Description("The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_BUILD_URL").
		Description("The url for this build on Buildkite.").
		String(),

	NewEnvVar("BUILDKITE_CANCEL_GRACE_PERIOD").
		Description("The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.").
		Number(),

	NewEnvVar("BUILDKITE_CANCEL_SIGNAL").
		Description("The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_CLEAN_CHECKOUT").
		Description("Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.").
		Boolean(),

	NewEnvVar("BUILDKITE_CLUSTER_ID").
		Description("The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.").
		String(),

	NewEnvVar("BUILDKITE_COMMAND").
		Description("The command that will be run for the job.").
		String(),

	NewEnvVar("BUILDKITE_COMMAND_EVAL").
		Description("The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_COMMAND_EXIT_STATUS").
		Description("The exit code from the last command run in the command hook.").
		Number(),

	NewEnvVar("BUILDKITE_COMMIT").
		Description("The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.").
		String(),

	NewEnvVar("BUILDKITE_CONFIG_PATH").
		Description("The path to the agent config file.").
		String(),

	NewEnvVar("BUILDKITE_ENV_FILE").
		Description("The path to the file containing the job's environment variables.").
		String(),

	NewEnvVar("BUILDKITE_GIT_CLEAN_FLAGS").
		Description("The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_GIT_CLONE_FLAGS").
		Description("The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_GIT_SUBMODULES").
		Description("The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_GITHUB_DEPLOYMENT_ID").
		Description("The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).").
		String(),

	NewEnvVar("BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT").
		Description("The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).").
		String(),

	NewEnvVar("BUILDKITE_GITHUB_DEPLOYMENT_TASK").
		Description("The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).").
		String(),

	NewEnvVar("BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD").
		Description("The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).").
		String(),

	NewEnvVar("BUILDKITE_GROUP_ID").
		Description("The UUID of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.").
		String(),

	NewEnvVar("BUILDKITE_GROUP_KEY").
		Description("The value of the `key` attribute of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.").
		String(),

	NewEnvVar("BUILDKITE_GROUP_LABEL").
		Description("The label/name of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.").
		String(),

	NewEnvVar("BUILDKITE_HOOKS_PATH").
		Description("The value of the `hooks-path` [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_IGNORED_ENV").
		Description("A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.").
		StringArray(","),

	NewEnvVar("BUILDKITE_JOB_ID").
		Description("The internal UUID Buildkite uses for this job.").
		String(),

	NewEnvVar("BUILDKITE_JOB_LOG_TMPFILE").
		Description("The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_LABEL").
		Description("The label/name of the current job.").
		String(),

	NewEnvVar("BUILDKITE_LAST_HOOK_EXIT_STATUS").
		Description("The exit code of the last hook that ran, used internally by the hooks.").
		Number(),

	NewEnvVar("BUILDKITE_LOCAL_HOOKS_ENABLED").
		Description("The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_MESSAGE").
		Description("The message associated with the build, usually the commit message. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.").
		String(),

	NewEnvVar("BUILDKITE_ORGANIZATION_SLUG").
		Description("The organization name on Buildkite as used in URLs.").
		String(),

	NewEnvVar("BUILDKITE_PARALLEL_JOB").
		Description("The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.").
		Number(),

	NewEnvVar("BUILDKITE_PARALLEL_JOB_COUNT").
		Description("The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.").
		Number(),

	NewEnvVar("BUILDKITE_PIPELINE_DEFAULT_BRANCH").
		Description("The default branch for this pipeline.").
		String(),

	NewEnvVar("BUILDKITE_PIPELINE_NAME").
		Description("The displayed pipeline name on Buildkite.").
		String(),

	NewEnvVar("BUILDKITE_PIPELINE_PROVIDER").
		Description("The ID of the source code provider for the pipeline's repository.").
		String(),

	NewEnvVar("BUILDKITE_PIPELINE_SLUG").
		Description("The pipeline slug on Buildkite as used in URLs.").
		String(),

	NewEnvVar("BUILDKITE_PIPELINE_TEAMS").
		Description("A colon separated list of the pipeline's non-private team slugs.").
		StringArray(":"),

	NewEnvVar("BUILDKITE_PLUGIN_CONFIGURATION").
		Description("A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).").
		String(),

	NewEnvVar("BUILDKITE_PLUGIN_NAME").
		Description("The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.").
		String(),

	NewEnvVar("BUILDKITE_PLUGINS").
		Description("A JSON object containing a list plugins used in the step, and their configuration.").
		String(),

	NewEnvVar("BUILDKITE_PLUGINS_ENABLED").
		Description("The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_PLUGINS_PATH").
		Description("The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_PLUGIN_VALIDATION").
		Description("Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_PULL_REQUEST").
		Description("The number of the pull request or `false` if not a pull request.").
		Number(),

	NewEnvVar("BUILDKITE_PULL_REQUEST_BASE_BRANCH").
		Description("The base branch that the pull request is targeting or `\"\"` if not a pull request.").
		String(),

	NewEnvVar("BUILDKITE_PULL_REQUEST_DRAFT").
		Description("Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.").
		Boolean(),

	NewEnvVar("BUILDKITE_PULL_REQUEST_REPO").
		Description("The repository URL of the pull request or `\"\"` if not a pull request.").
		String(),

	NewEnvVar("BUILDKITE_REBUILT_FROM_BUILD_ID").
		Description("The UUID of the original build this was rebuilt from or `\"\"` if not a rebuild.").
		String(),

	NewEnvVar("BUILDKITE_REBUILT_FROM_BUILD_NUMBER").
		Description("The UUID of the original build this was rebuilt from or `\"\"` if not a rebuild.").
		String(),

	NewEnvVar("BUILDKITE_REFSPEC").
		Description("A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_REPO").
		Description("The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.").
		String(),

	NewEnvVar("BUILDKITE_REPO_MIRROR").
		Description("The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).").
		String(),

	NewEnvVar("BUILDKITE_RETRY_COUNT").
		Description("How many times this job has been retried.").
		Number(),

	NewEnvVar("BUILDKITE_S3_ACCESS_KEY_ID").
		Description("The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.").
		String(),

	NewEnvVar("BUILDKITE_S3_ACCESS_URL").
		Description("The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.").
		String(),

	NewEnvVar("BUILDKITE_S3_ACL").
		Description(`The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the 'buildkite-agent artifact upload' command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the 'environment', 'pre-checkout' or 'pre-command' hooks.

Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).`).
		String(),

	NewEnvVar("BUILDKITE_S3_DEFAULT_REGION").
		Description("The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.").
		String(),

	NewEnvVar("BUILDKITE_S3_SECRET_ACCESS_KEY").
		Description("The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.").
		String(),

	NewEnvVar("BUILDKITE_S3_SSE_ENABLED").
		Description("Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.").
		Boolean(),

	NewEnvVar("BUILDKITE_SHELL").
		Description("The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).").
		String(),

	NewEnvVar("BUILDKITE_SOURCE").
		Description("The source of the event that created the build.").
		String(),

	NewEnvVar("BUILDKITE_SSH_KEYSCAN").
		Description("The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).").
		Boolean(),

	NewEnvVar("BUILDKITE_STEP_ID").
		Description("A unique string that identifies a step.").
		String(),

	NewEnvVar("BUILDKITE_STEP_KEY").
		Description("The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.").
		String(),

	NewEnvVar("BUILDKITE_TAG").
		Description("The name of the tag being built, if this build was triggered from a tag.").
		String(),

	NewEnvVar("BUILDKITE_TIMEOUT").
		Description("The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as \"passed\".").
		Number(),

	NewEnvVar("BUILDKITE_TRACING_BACKEND").
		Description(`Set to "datadog" to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using 'localhost:8126', or 'DD_AGENT_HOST:DD_AGENT_APM_PORT'.

Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)`).
		String(),

	NewEnvVar("BUILDKITE_TRIGGERED_FROM_BUILD_ID").
		Description("The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.").
		String(),

	NewEnvVar("BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER").
		Description("The number of the build that triggered this build or `\"\"` if the build was not triggered from another build.").
		String(),

	NewEnvVar("BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG").
		Description("The slug of the pipeline that was used to trigger this build or `\"\"` if the build was not triggered from another build.").
		String(),

	NewEnvVar("BUILDKITE_UNBLOCKER").
		Description("The name of the user who unblocked the build.").
		String(),

	NewEnvVar("BUILDKITE_UNBLOCKER_EMAIL").
		Description("The notification email of the user who unblocked the build.").
		String(),

	NewEnvVar("BUILDKITE_UNBLOCKER_ID").
		Description("The UUID of the user who unblocked the build.").
		String(),

	NewEnvVar("BUILDKITE_UNBLOCKER_TEAMS").
		Description("A colon separated list of non-private team slugs that the user who unblocked the build belongs to.").
		StringArray(":"),
}
