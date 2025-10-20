package buildkite

type env string

var Strings = struct {
	// Always `true`
	BUILDKITE env
	// The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.
	BUILDKITE_AGENT_ACCESS_TOKEN env
	// The value of the `debug` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_AGENT_DEBUG env
	// The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_AGENT_DISCONNECT_AFTER_JOB env
	// The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT env
	// The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.
	BUILDKITE_AGENT_ENDPOINT env
	// A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).
	BUILDKITE_AGENT_EXPERIMENT env
	// The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_AGENT_HEALTH_CHECK_ADDR env
	// The UUID of the agent.
	BUILDKITE_AGENT_ID env
	// The value of each [agent tag](/docs/agent/v3/cli-start#setting-tags). The tag name is appended to the end of the variable name. They can be set using the `--tags` flag on the `buildkite-agent start` command, or in the [agent configuration file](/docs/agent/v3/configuration). The [Queue tag](/docs/agent/v3/queues) is specifically used for isolating jobs and agents, and appears as the `BUILDKITE_AGENT_META_DATA_QUEUE` environment variable.
	BUILDKITE_AGENT_META_DATA_ env
	// The name of the agent that ran the job.
	BUILDKITE_AGENT_NAME env
	// The process ID of the agent.
	BUILDKITE_AGENT_PID env
	// The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_ARTIFACT_PATHS env
	// The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	BUILDKITE_ARTIFACT_UPLOAD_DESTINATION env
	// The path to the directory containing the `buildkite-agent` binary.
	BUILDKITE_BIN_PATH env
	// The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.
	BUILDKITE_BRANCH env
	// The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_BUILD_CHECKOUT_PATH env
	// The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
	BUILDKITE_BUILD_AUTHOR env
	// The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
	BUILDKITE_BUILD_AUTHOR_EMAIL env
	// The name of the user who created the build. The value differs depending on how the build was created:
	//
	// - **Buildkite dashboard:** Set based on who manually created the build.
	// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
	// - **Webhook:** Set based on which user is attached to the API Key used.
	BUILDKITE_BUILD_CREATOR env
	// The notification email of the user who created the build. The value differs depending on how the build was created:
	//
	// - **Buildkite dashboard:** Set based on who manually created the build.
	// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
	// - **Webhook:** Set based on which user is attached to the API Key used.
	BUILDKITE_BUILD_CREATOR_EMAIL env
	// A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:
	//
	// - **Buildkite dashboard:** Set based on who manually created the build.
	// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
	// - **Webhook:** Set based on which user is attached to the API Key used.
	BUILDKITE_BUILD_CREATOR_TEAMS env
	// The UUID of the build.
	BUILDKITE_BUILD_ID env
	// The build number. This number increases with every build, and is guaranteed to be unique within each pipeline.
	BUILDKITE_BUILD_NUMBER env
	// The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_BUILD_PATH env
	// The url for this build on Buildkite.
	BUILDKITE_BUILD_URL env
	// The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.
	BUILDKITE_CANCEL_GRACE_PERIOD env
	// The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_CANCEL_SIGNAL env
	// Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.
	BUILDKITE_CLEAN_CHECKOUT env
	// The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.
	BUILDKITE_CLUSTER_ID env
	// The command that will be run for the job.
	BUILDKITE_COMMAND env
	// The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_COMMAND_EVAL env
	// The exit code from the last command run in the command hook.
	BUILDKITE_COMMAND_EXIT_STATUS env
	// The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.
	BUILDKITE_COMMIT env
	// The path to the agent config file.
	BUILDKITE_CONFIG_PATH env
	// The path to the file containing the job's environment variables.
	BUILDKITE_ENV_FILE env
	// The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_GIT_CLEAN_FLAGS env
	// The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_GIT_CLONE_FLAGS env
	// The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_GIT_SUBMODULES env
	// The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
	BUILDKITE_GITHUB_DEPLOYMENT_ID env
	// The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
	BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT env
	// The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
	BUILDKITE_GITHUB_DEPLOYMENT_TASK env
	// The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
	BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD env
	// The UUID of the [group step](/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
	BUILDKITE_GROUP_ID env
	// The value of the `key` attribute of the [group step](/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
	BUILDKITE_GROUP_KEY env
	// The label/name of the [group step](/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
	BUILDKITE_GROUP_LABEL env
	// The value of the `hooks-path` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_HOOKS_PATH env
	// A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.
	BUILDKITE_IGNORED_ENV env
	// The internal UUID Buildkite uses for this job.
	BUILDKITE_JOB_ID env
	// The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_JOB_LOG_TMPFILE env
	// The label/name of the current job.
	BUILDKITE_LABEL env
	// The exit code of the last hook that ran, used internally by the hooks.
	BUILDKITE_LAST_HOOK_EXIT_STATUS env
	// The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_LOCAL_HOOKS_ENABLED env
	// The message associated with the build, usually the commit message or the message provided when the build is triggered. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.
	BUILDKITE_MESSAGE env
	// The UUID of the organization.
	BUILDKITE_ORGANIZATION_ID env
	// The organization name on Buildkite as used in URLs.
	BUILDKITE_ORGANIZATION_SLUG env
	// The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.
	BUILDKITE_PARALLEL_JOB env
	// The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.
	BUILDKITE_PARALLEL_JOB_COUNT env
	// The default branch for this pipeline.
	BUILDKITE_PIPELINE_DEFAULT_BRANCH env
	// The UUID of the pipeline.
	BUILDKITE_PIPELINE_ID env
	// The displayed pipeline name on Buildkite.
	BUILDKITE_PIPELINE_NAME env
	// The ID of the source code provider for the pipeline's repository.
	BUILDKITE_PIPELINE_PROVIDER env
	// The pipeline slug on Buildkite as used in URLs.
	BUILDKITE_PIPELINE_SLUG env
	// A colon separated list of the pipeline's non-private team slugs.
	BUILDKITE_PIPELINE_TEAMS env
	// A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).
	BUILDKITE_PLUGIN_CONFIGURATION env
	// The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.
	BUILDKITE_PLUGIN_NAME env
	// A JSON object containing a list plugins used in the step, and their configuration.
	BUILDKITE_PLUGINS env
	// The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_PLUGINS_ENABLED env
	// The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_PLUGINS_PATH env
	// Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_PLUGIN_VALIDATION env
	// The number of the pull request or `false` if not a pull request.
	BUILDKITE_PULL_REQUEST env
	// The base branch that the pull request is targeting or `""` if not a pull request.`
	BUILDKITE_PULL_REQUEST_BASE_BRANCH env
	// Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.
	BUILDKITE_PULL_REQUEST_DRAFT env
	// The repository URL of the pull request or `""` if not a pull request.
	BUILDKITE_PULL_REQUEST_REPO env
	// The UUID of the original build this was rebuilt from or `""` if not a rebuild.
	BUILDKITE_REBUILT_FROM_BUILD_ID env
	// The number of the original build this was rebuilt from or `""` if not a rebuild.
	BUILDKITE_REBUILT_FROM_BUILD_NUMBER env
	// A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_REFSPEC env
	// The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
	BUILDKITE_REPO env
	// The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).
	BUILDKITE_REPO_MIRROR env
	// How many times this job has been retried.
	BUILDKITE_RETRY_COUNT env
	// The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	BUILDKITE_S3_ACCESS_KEY_ID env
	// The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	BUILDKITE_S3_ACCESS_URL env
	// The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	//
	// Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).
	BUILDKITE_S3_ACL env
	// The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	BUILDKITE_S3_DEFAULT_REGION env
	// The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.
	BUILDKITE_S3_SECRET_ACCESS_KEY env
	// Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
	BUILDKITE_S3_SSE_ENABLED env
	// The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_SHELL env
	// The source of the event that created the build.
	BUILDKITE_SOURCE env
	// The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).
	BUILDKITE_SSH_KEYSCAN env
	// A unique string that identifies a step.
	BUILDKITE_STEP_ID env
	// The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.
	BUILDKITE_STEP_KEY env
	// The name of the tag being built, if this build was triggered from a tag.
	BUILDKITE_TAG env
	// The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as "passed".
	BUILDKITE_TIMEOUT env
	// Set to `"datadog"` to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using `localhost:8126`, or `DD_AGENT_HOST:DD_AGENT_APM_PORT`.
	//
	// Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)
	BUILDKITE_TRACING_BACKEND env
	// The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.
	BUILDKITE_TRIGGERED_FROM_BUILD_ID env
	// The number of the build that triggered this build or `""` if the build was not triggered from another build.
	BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER env
	// The slug of the pipeline that was used to trigger this build or `""` if the build was not triggered from another build.
	BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG env
	// The name of the user who unblocked the build.
	BUILDKITE_UNBLOCKER env
	// The notification email of the user who unblocked the build.
	BUILDKITE_UNBLOCKER_EMAIL env
	// The UUID of the user who unblocked the build.
	BUILDKITE_UNBLOCKER_ID env
	// A colon separated list of non-private team slugs that the user who unblocked the build belongs to.
	BUILDKITE_UNBLOCKER_TEAMS env
	// Always `true`.
	CI env
}{
	BUILDKITE:                                     "BUILDKITE",
	BUILDKITE_AGENT_ACCESS_TOKEN:                  "BUILDKITE_AGENT_ACCESS_TOKEN",
	BUILDKITE_AGENT_DEBUG:                         "BUILDKITE_AGENT_DEBUG",
	BUILDKITE_AGENT_DISCONNECT_AFTER_JOB:          "BUILDKITE_AGENT_DISCONNECT_AFTER_JOB",
	BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT: "BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT",
	BUILDKITE_AGENT_ENDPOINT:                      "BUILDKITE_AGENT_ENDPOINT",
	BUILDKITE_AGENT_EXPERIMENT:                    "BUILDKITE_AGENT_EXPERIMENT",
	BUILDKITE_AGENT_HEALTH_CHECK_ADDR:             "BUILDKITE_AGENT_HEALTH_CHECK_ADDR",
	BUILDKITE_AGENT_ID:                            "BUILDKITE_AGENT_ID",
	BUILDKITE_AGENT_META_DATA_:                    "BUILDKITE_AGENT_META_DATA_",
	BUILDKITE_AGENT_NAME:                          "BUILDKITE_AGENT_NAME",
	BUILDKITE_AGENT_PID:                           "BUILDKITE_AGENT_PID",
	BUILDKITE_ARTIFACT_PATHS:                      "BUILDKITE_ARTIFACT_PATHS",
	BUILDKITE_ARTIFACT_UPLOAD_DESTINATION:         "BUILDKITE_ARTIFACT_UPLOAD_DESTINATION",
	BUILDKITE_BIN_PATH:                            "BUILDKITE_BIN_PATH",
	BUILDKITE_BRANCH:                              "BUILDKITE_BRANCH",
	BUILDKITE_BUILD_CHECKOUT_PATH:                 "BUILDKITE_BUILD_CHECKOUT_PATH",
	BUILDKITE_BUILD_AUTHOR:                        "BUILDKITE_BUILD_AUTHOR",
	BUILDKITE_BUILD_AUTHOR_EMAIL:                  "BUILDKITE_BUILD_AUTHOR_EMAIL",
	BUILDKITE_BUILD_CREATOR:                       "BUILDKITE_BUILD_CREATOR",
	BUILDKITE_BUILD_CREATOR_EMAIL:                 "BUILDKITE_BUILD_CREATOR_EMAIL",
	BUILDKITE_BUILD_CREATOR_TEAMS:                 "BUILDKITE_BUILD_CREATOR_TEAMS",
	BUILDKITE_BUILD_ID:                            "BUILDKITE_BUILD_ID",
	BUILDKITE_BUILD_NUMBER:                        "BUILDKITE_BUILD_NUMBER",
	BUILDKITE_BUILD_PATH:                          "BUILDKITE_BUILD_PATH",
	BUILDKITE_BUILD_URL:                           "BUILDKITE_BUILD_URL",
	BUILDKITE_CANCEL_GRACE_PERIOD:                 "BUILDKITE_CANCEL_GRACE_PERIOD",
	BUILDKITE_CANCEL_SIGNAL:                       "BUILDKITE_CANCEL_SIGNAL",
	BUILDKITE_CLEAN_CHECKOUT:                      "BUILDKITE_CLEAN_CHECKOUT",
	BUILDKITE_CLUSTER_ID:                          "BUILDKITE_CLUSTER_ID",
	BUILDKITE_COMMAND:                             "BUILDKITE_COMMAND",
	BUILDKITE_COMMAND_EVAL:                        "BUILDKITE_COMMAND_EVAL",
	BUILDKITE_COMMAND_EXIT_STATUS:                 "BUILDKITE_COMMAND_EXIT_STATUS",
	BUILDKITE_COMMIT:                              "BUILDKITE_COMMIT",
	BUILDKITE_CONFIG_PATH:                         "BUILDKITE_CONFIG_PATH",
	BUILDKITE_ENV_FILE:                            "BUILDKITE_ENV_FILE",
	BUILDKITE_GIT_CLEAN_FLAGS:                     "BUILDKITE_GIT_CLEAN_FLAGS",
	BUILDKITE_GIT_CLONE_FLAGS:                     "BUILDKITE_GIT_CLONE_FLAGS",
	BUILDKITE_GIT_SUBMODULES:                      "BUILDKITE_GIT_SUBMODULES",
	BUILDKITE_GITHUB_DEPLOYMENT_ID:                "BUILDKITE_GITHUB_DEPLOYMENT_ID",
	BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT:       "BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT",
	BUILDKITE_GITHUB_DEPLOYMENT_TASK:              "BUILDKITE_GITHUB_DEPLOYMENT_TASK",
	BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD:           "BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD",
	BUILDKITE_GROUP_ID:                            "BUILDKITE_GROUP_ID",
	BUILDKITE_GROUP_KEY:                           "BUILDKITE_GROUP_KEY",
	BUILDKITE_GROUP_LABEL:                         "BUILDKITE_GROUP_LABEL",
	BUILDKITE_HOOKS_PATH:                          "BUILDKITE_HOOKS_PATH",
	BUILDKITE_IGNORED_ENV:                         "BUILDKITE_IGNORED_ENV",
	BUILDKITE_JOB_ID:                              "BUILDKITE_JOB_ID",
	BUILDKITE_JOB_LOG_TMPFILE:                     "BUILDKITE_JOB_LOG_TMPFILE",
	BUILDKITE_LABEL:                               "BUILDKITE_LABEL",
	BUILDKITE_LAST_HOOK_EXIT_STATUS:               "BUILDKITE_LAST_HOOK_EXIT_STATUS",
	BUILDKITE_LOCAL_HOOKS_ENABLED:                 "BUILDKITE_LOCAL_HOOKS_ENABLED",
	BUILDKITE_MESSAGE:                             "BUILDKITE_MESSAGE",
	BUILDKITE_ORGANIZATION_ID:                     "BUILDKITE_ORGANIZATION_ID",
	BUILDKITE_ORGANIZATION_SLUG:                   "BUILDKITE_ORGANIZATION_SLUG",
	BUILDKITE_PARALLEL_JOB:                        "BUILDKITE_PARALLEL_JOB",
	BUILDKITE_PARALLEL_JOB_COUNT:                  "BUILDKITE_PARALLEL_JOB_COUNT",
	BUILDKITE_PIPELINE_DEFAULT_BRANCH:             "BUILDKITE_PIPELINE_DEFAULT_BRANCH",
	BUILDKITE_PIPELINE_ID:                         "BUILDKITE_PIPELINE_ID",
	BUILDKITE_PIPELINE_NAME:                       "BUILDKITE_PIPELINE_NAME",
	BUILDKITE_PIPELINE_PROVIDER:                   "BUILDKITE_PIPELINE_PROVIDER",
	BUILDKITE_PIPELINE_SLUG:                       "BUILDKITE_PIPELINE_SLUG",
	BUILDKITE_PIPELINE_TEAMS:                      "BUILDKITE_PIPELINE_TEAMS",
	BUILDKITE_PLUGIN_CONFIGURATION:                "BUILDKITE_PLUGIN_CONFIGURATION",
	BUILDKITE_PLUGIN_NAME:                         "BUILDKITE_PLUGIN_NAME",
	BUILDKITE_PLUGINS:                             "BUILDKITE_PLUGINS",
	BUILDKITE_PLUGINS_ENABLED:                     "BUILDKITE_PLUGINS_ENABLED",
	BUILDKITE_PLUGINS_PATH:                        "BUILDKITE_PLUGINS_PATH",
	BUILDKITE_PLUGIN_VALIDATION:                   "BUILDKITE_PLUGIN_VALIDATION",
	BUILDKITE_PULL_REQUEST:                        "BUILDKITE_PULL_REQUEST",
	BUILDKITE_PULL_REQUEST_BASE_BRANCH:            "BUILDKITE_PULL_REQUEST_BASE_BRANCH",
	BUILDKITE_PULL_REQUEST_DRAFT:                  "BUILDKITE_PULL_REQUEST_DRAFT",
	BUILDKITE_PULL_REQUEST_REPO:                   "BUILDKITE_PULL_REQUEST_REPO",
	BUILDKITE_REBUILT_FROM_BUILD_ID:               "BUILDKITE_REBUILT_FROM_BUILD_ID",
	BUILDKITE_REBUILT_FROM_BUILD_NUMBER:           "BUILDKITE_REBUILT_FROM_BUILD_NUMBER",
	BUILDKITE_REFSPEC:                             "BUILDKITE_REFSPEC",
	BUILDKITE_REPO:                                "BUILDKITE_REPO",
	BUILDKITE_REPO_MIRROR:                         "BUILDKITE_REPO_MIRROR",
	BUILDKITE_RETRY_COUNT:                         "BUILDKITE_RETRY_COUNT",
	BUILDKITE_S3_ACCESS_KEY_ID:                    "BUILDKITE_S3_ACCESS_KEY_ID",
	BUILDKITE_S3_ACCESS_URL:                       "BUILDKITE_S3_ACCESS_URL",
	BUILDKITE_S3_ACL:                              "BUILDKITE_S3_ACL",
	BUILDKITE_S3_DEFAULT_REGION:                   "BUILDKITE_S3_DEFAULT_REGION",
	BUILDKITE_S3_SECRET_ACCESS_KEY:                "BUILDKITE_S3_SECRET_ACCESS_KEY",
	BUILDKITE_S3_SSE_ENABLED:                      "BUILDKITE_S3_SSE_ENABLED",
	BUILDKITE_SHELL:                               "BUILDKITE_SHELL",
	BUILDKITE_SOURCE:                              "BUILDKITE_SOURCE",
	BUILDKITE_SSH_KEYSCAN:                         "BUILDKITE_SSH_KEYSCAN",
	BUILDKITE_STEP_ID:                             "BUILDKITE_STEP_ID",
	BUILDKITE_STEP_KEY:                            "BUILDKITE_STEP_KEY",
	BUILDKITE_TAG:                                 "BUILDKITE_TAG",
	BUILDKITE_TIMEOUT:                             "BUILDKITE_TIMEOUT",
	BUILDKITE_TRACING_BACKEND:                     "BUILDKITE_TRACING_BACKEND",
	BUILDKITE_TRIGGERED_FROM_BUILD_ID:             "BUILDKITE_TRIGGERED_FROM_BUILD_ID",
	BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER:         "BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER",
	BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG:  "BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG",
	BUILDKITE_UNBLOCKER:                           "BUILDKITE_UNBLOCKER",
	BUILDKITE_UNBLOCKER_EMAIL:                     "BUILDKITE_UNBLOCKER_EMAIL",
	BUILDKITE_UNBLOCKER_ID:                        "BUILDKITE_UNBLOCKER_ID",
	BUILDKITE_UNBLOCKER_TEAMS:                     "BUILDKITE_UNBLOCKER_TEAMS",
	CI:                                            "CI",
}
