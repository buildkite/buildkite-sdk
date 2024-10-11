// This file is auto-generated please do not edit

class Environment {
    // The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.
    public BUILDKITE_AGENT_ACCESS_TOKEN(): string {
        return process.env.BUILDKITE_AGENT_ACCESS_TOKEN;
    }
    // The value of the debug [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
    public BUILDKITE_AGENT_DEBUG(): boolean {
        return Boolean(process.env.BUILDKITE_AGENT_DEBUG);
    }
    // The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_AGENT_DISCONNECT_AFTER_JOB(): boolean {
        return Boolean(process.env.BUILDKITE_AGENT_DISCONNECT_AFTER_JOB);
    }
    // The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT(): number {
        return Number(process.env.BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT);
    }
    // The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.
    public BUILDKITE_AGENT_ENDPOINT(): string {
        return process.env.BUILDKITE_AGENT_ENDPOINT;
    }
    // A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).
    public BUILDKITE_AGENT_EXPERIMENT(): string[] {
        return process.env.BUILDKITE_AGENT_EXPERIMENT.split(",");
    }
    // The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_AGENT_HEALTH_CHECK_ADDR(): string {
        return process.env.BUILDKITE_AGENT_HEALTH_CHECK_ADDR;
    }
    // The UUID of the agent.
    public BUILDKITE_AGENT_ID(): string {
        return process.env.BUILDKITE_AGENT_ID;
    }
    // The value of each agent tag. The tag name is appended to the end of the variable name. They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file. The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.
    public BUILDKITE_AGENT_META_DATA(...strs: string[]): string {
        return process.env[strs.join("_").toUpperCase()];
    }
    // The name of the agent that ran the job.
    public BUILDKITE_AGENT_NAME(): string {
        return process.env.BUILDKITE_AGENT_NAME;
    }
    // The process ID of the agent.
    public BUILDKITE_AGENT_PID(): string {
        return process.env.BUILDKITE_AGENT_PID;
    }
    // The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_ARTIFACT_PATHS(): string[] {
        return process.env.BUILDKITE_ARTIFACT_PATHS.split(";");
    }
    // The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public BUILDKITE_ARTIFACT_UPLOAD_DESTINATION(): string {
        return process.env.BUILDKITE_ARTIFACT_UPLOAD_DESTINATION;
    }
    // The path to the directory containing the `buildkite-agent` binary.
    public BUILDKITE_BIN_PATH(): string {
        return process.env.BUILDKITE_BIN_PATH;
    }
    // The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.
    public BUILDKITE_BRANCH(): string {
        return process.env.BUILDKITE_BRANCH;
    }
    // The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_BUILD_CHECKOUT_PATH(): string {
        return process.env.BUILDKITE_BUILD_CHECKOUT_PATH;
    }
    // The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
    public BUILDKITE_BUILD_AUTHOR(): string {
        return process.env.BUILDKITE_BUILD_AUTHOR;
    }
    // The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
    public BUILDKITE_BUILD_AUTHOR_EMAIL(): string {
        return process.env.BUILDKITE_BUILD_AUTHOR_EMAIL;
    }
    // The name of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public BUILDKITE_BUILD_CREATOR(): string {
        return process.env.BUILDKITE_BUILD_CREATOR;
    }
    // The notification email of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public BUILDKITE_BUILD_CREATOR_EMAIL(): string {
        return process.env.BUILDKITE_BUILD_CREATOR_EMAIL;
    }
    // A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public BUILDKITE_BUILD_CREATOR_TEAMS(): string[] {
        return process.env.BUILDKITE_BUILD_CREATOR_TEAMS.split(":");
    }
    // The UUID of the build.
    public BUILDKITE_BUILD_ID(): string {
        return process.env.BUILDKITE_BUILD_ID;
    }
    // The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.
    public BUILDKITE_BUILD_NUMBER(): number {
        return Number(process.env.BUILDKITE_BUILD_NUMBER);
    }
    // The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_BUILD_PATH(): string {
        return process.env.BUILDKITE_BUILD_PATH;
    }
    // The url for this build on Buildkite.
    public BUILDKITE_BUILD_URL(): string {
        return process.env.BUILDKITE_BUILD_URL;
    }
    // The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.
    public BUILDKITE_CANCEL_GRACE_PERIOD(): number {
        return Number(process.env.BUILDKITE_CANCEL_GRACE_PERIOD);
    }
    // The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_CANCEL_SIGNAL(): string {
        return process.env.BUILDKITE_CANCEL_SIGNAL;
    }
    // Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.
    public BUILDKITE_CLEAN_CHECKOUT(): boolean {
        return Boolean(process.env.BUILDKITE_CLEAN_CHECKOUT);
    }
    // The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.
    public BUILDKITE_CLUSTER_ID(): string {
        return process.env.BUILDKITE_CLUSTER_ID;
    }
    // The command that will be run for the job.
    public BUILDKITE_COMMAND(): string {
        return process.env.BUILDKITE_COMMAND;
    }
    // The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_COMMAND_EVAL(): boolean {
        return Boolean(process.env.BUILDKITE_COMMAND_EVAL);
    }
    // The exit code from the last command run in the command hook.
    public BUILDKITE_COMMAND_EXIT_STATUS(): number {
        return Number(process.env.BUILDKITE_COMMAND_EXIT_STATUS);
    }
    // The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.
    public BUILDKITE_COMMIT(): string {
        return process.env.BUILDKITE_COMMIT;
    }
    // The path to the agent config file.
    public BUILDKITE_CONFIG_PATH(): string {
        return process.env.BUILDKITE_CONFIG_PATH;
    }
    // The path to the file containing the job's environment variables.
    public BUILDKITE_ENV_FILE(): string {
        return process.env.BUILDKITE_ENV_FILE;
    }
    // The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_GIT_CLEAN_FLAGS(): string {
        return process.env.BUILDKITE_GIT_CLEAN_FLAGS;
    }
    // The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_GIT_CLONE_FLAGS(): string {
        return process.env.BUILDKITE_GIT_CLONE_FLAGS;
    }
    // The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_GIT_SUBMODULES(): boolean {
        return Boolean(process.env.BUILDKITE_GIT_SUBMODULES);
    }
    // The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public BUILDKITE_GITHUB_DEPLOYMENT_ID(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_ID;
    }
    // The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT;
    }
    // The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public BUILDKITE_GITHUB_DEPLOYMENT_TASK(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_TASK;
    }
    // The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD;
    }
    // The UUID of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public BUILDKITE_GROUP_ID(): string {
        return process.env.BUILDKITE_GROUP_ID;
    }
    // The value of the `key` attribute of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public BUILDKITE_GROUP_KEY(): string {
        return process.env.BUILDKITE_GROUP_KEY;
    }
    // The label/name of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public BUILDKITE_GROUP_LABEL(): string {
        return process.env.BUILDKITE_GROUP_LABEL;
    }
    // The value of the `hooks-path` [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
    public BUILDKITE_HOOKS_PATH(): string {
        return process.env.BUILDKITE_HOOKS_PATH;
    }
    // A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.
    public BUILDKITE_IGNORED_ENV(): string[] {
        return process.env.BUILDKITE_IGNORED_ENV.split(",");
    }
    // The internal UUID Buildkite uses for this job.
    public BUILDKITE_JOB_ID(): string {
        return process.env.BUILDKITE_JOB_ID;
    }
    // The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_JOB_LOG_TMPFILE(): string {
        return process.env.BUILDKITE_JOB_LOG_TMPFILE;
    }
    // The label/name of the current job.
    public BUILDKITE_LABEL(): string {
        return process.env.BUILDKITE_LABEL;
    }
    // The exit code of the last hook that ran, used internally by the hooks.
    public BUILDKITE_LAST_HOOK_EXIT_STATUS(): number {
        return Number(process.env.BUILDKITE_LAST_HOOK_EXIT_STATUS);
    }
    // The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_LOCAL_HOOKS_ENABLED(): boolean {
        return Boolean(process.env.BUILDKITE_LOCAL_HOOKS_ENABLED);
    }
    // The message associated with the build, usually the commit message. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.
    public BUILDKITE_MESSAGE(): string {
        return process.env.BUILDKITE_MESSAGE;
    }
    // The organization name on Buildkite as used in URLs.
    public BUILDKITE_ORGANIZATION_SLUG(): string {
        return process.env.BUILDKITE_ORGANIZATION_SLUG;
    }
    // The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.
    public BUILDKITE_PARALLEL_JOB(): number {
        return Number(process.env.BUILDKITE_PARALLEL_JOB);
    }
    // The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.
    public BUILDKITE_PARALLEL_JOB_COUNT(): number {
        return Number(process.env.BUILDKITE_PARALLEL_JOB_COUNT);
    }
    // The default branch for this pipeline.
    public BUILDKITE_PIPELINE_DEFAULT_BRANCH(): string {
        return process.env.BUILDKITE_PIPELINE_DEFAULT_BRANCH;
    }
    // The displayed pipeline name on Buildkite.
    public BUILDKITE_PIPELINE_NAME(): string {
        return process.env.BUILDKITE_PIPELINE_NAME;
    }
    // The ID of the source code provider for the pipeline's repository.
    public BUILDKITE_PIPELINE_PROVIDER(): string {
        return process.env.BUILDKITE_PIPELINE_PROVIDER;
    }
    // The pipeline slug on Buildkite as used in URLs.
    public BUILDKITE_PIPELINE_SLUG(): string {
        return process.env.BUILDKITE_PIPELINE_SLUG;
    }
    // A colon separated list of the pipeline's non-private team slugs.
    public BUILDKITE_PIPELINE_TEAMS(): string[] {
        return process.env.BUILDKITE_PIPELINE_TEAMS.split(":");
    }
    // A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).
    public BUILDKITE_PLUGIN_CONFIGURATION(): string {
        return process.env.BUILDKITE_PLUGIN_CONFIGURATION;
    }
    // The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.
    public BUILDKITE_PLUGIN_NAME(): string {
        return process.env.BUILDKITE_PLUGIN_NAME;
    }
    // A JSON object containing a list plugins used in the step, and their configuration.
    public BUILDKITE_PLUGINS(): string {
        return process.env.BUILDKITE_PLUGINS;
    }
    // The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_PLUGINS_ENABLED(): boolean {
        return Boolean(process.env.BUILDKITE_PLUGINS_ENABLED);
    }
    // The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_PLUGINS_PATH(): string {
        return process.env.BUILDKITE_PLUGINS_PATH;
    }
    // Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_PLUGIN_VALIDATION(): boolean {
        return Boolean(process.env.BUILDKITE_PLUGIN_VALIDATION);
    }
    // The number of the pull request or `false` if not a pull request.
    public BUILDKITE_PULL_REQUEST(): number {
        return Number(process.env.BUILDKITE_PULL_REQUEST);
    }
    // The base branch that the pull request is targeting or `""` if not a pull request.
    public BUILDKITE_PULL_REQUEST_BASE_BRANCH(): string {
        return process.env.BUILDKITE_PULL_REQUEST_BASE_BRANCH;
    }
    // Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.
    public BUILDKITE_PULL_REQUEST_DRAFT(): boolean {
        return Boolean(process.env.BUILDKITE_PULL_REQUEST_DRAFT);
    }
    // The repository URL of the pull request or `""` if not a pull request.
    public BUILDKITE_PULL_REQUEST_REPO(): string {
        return process.env.BUILDKITE_PULL_REQUEST_REPO;
    }
    // The UUID of the original build this was rebuilt from or `""` if not a rebuild.
    public BUILDKITE_REBUILT_FROM_BUILD_ID(): string {
        return process.env.BUILDKITE_REBUILT_FROM_BUILD_ID;
    }
    // The UUID of the original build this was rebuilt from or `""` if not a rebuild.
    public BUILDKITE_REBUILT_FROM_BUILD_NUMBER(): string {
        return process.env.BUILDKITE_REBUILT_FROM_BUILD_NUMBER;
    }
    // A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_REFSPEC(): string {
        return process.env.BUILDKITE_REFSPEC;
    }
    // The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public BUILDKITE_REPO(): string {
        return process.env.BUILDKITE_REPO;
    }
    // The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).
    public BUILDKITE_REPO_MIRROR(): string {
        return process.env.BUILDKITE_REPO_MIRROR;
    }
    // How many times this job has been retried.
    public BUILDKITE_RETRY_COUNT(): number {
        return Number(process.env.BUILDKITE_RETRY_COUNT);
    }
    // The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public BUILDKITE_S3_ACCESS_KEY_ID(): string {
        return process.env.BUILDKITE_S3_ACCESS_KEY_ID;
    }
    // The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public BUILDKITE_S3_ACCESS_URL(): string {
        return process.env.BUILDKITE_S3_ACCESS_URL;
    }
    // The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the 'buildkite-agent artifact upload' command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the 'environment', 'pre-checkout' or 'pre-command' hooks.
// 
// Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).
    public BUILDKITE_S3_ACL(): string {
        return process.env.BUILDKITE_S3_ACL;
    }
    // The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public BUILDKITE_S3_DEFAULT_REGION(): string {
        return process.env.BUILDKITE_S3_DEFAULT_REGION;
    }
    // The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.
    public BUILDKITE_S3_SECRET_ACCESS_KEY(): string {
        return process.env.BUILDKITE_S3_SECRET_ACCESS_KEY;
    }
    // Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public BUILDKITE_S3_SSE_ENABLED(): boolean {
        return Boolean(process.env.BUILDKITE_S3_SSE_ENABLED);
    }
    // The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_SHELL(): string {
        return process.env.BUILDKITE_SHELL;
    }
    // The source of the event that created the build.
    public BUILDKITE_SOURCE(): string {
        return process.env.BUILDKITE_SOURCE;
    }
    // The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).
    public BUILDKITE_SSH_KEYSCAN(): boolean {
        return Boolean(process.env.BUILDKITE_SSH_KEYSCAN);
    }
    // A unique string that identifies a step.
    public BUILDKITE_STEP_ID(): string {
        return process.env.BUILDKITE_STEP_ID;
    }
    // The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.
    public BUILDKITE_STEP_KEY(): string {
        return process.env.BUILDKITE_STEP_KEY;
    }
    // The name of the tag being built, if this build was triggered from a tag.
    public BUILDKITE_TAG(): string {
        return process.env.BUILDKITE_TAG;
    }
    // The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as "passed".
    public BUILDKITE_TIMEOUT(): number {
        return Number(process.env.BUILDKITE_TIMEOUT);
    }
    // Set to "datadog" to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using 'localhost:8126', or 'DD_AGENT_HOST:DD_AGENT_APM_PORT'.
// 
// Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)
    public BUILDKITE_TRACING_BACKEND(): string {
        return process.env.BUILDKITE_TRACING_BACKEND;
    }
    // The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.
    public BUILDKITE_TRIGGERED_FROM_BUILD_ID(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_ID;
    }
    // The number of the build that triggered this build or `""` if the build was not triggered from another build.
    public BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER;
    }
    // The slug of the pipeline that was used to trigger this build or `""` if the build was not triggered from another build.
    public BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG;
    }
    // The name of the user who unblocked the build.
    public BUILDKITE_UNBLOCKER(): string {
        return process.env.BUILDKITE_UNBLOCKER;
    }
    // The notification email of the user who unblocked the build.
    public BUILDKITE_UNBLOCKER_EMAIL(): string {
        return process.env.BUILDKITE_UNBLOCKER_EMAIL;
    }
    // The UUID of the user who unblocked the build.
    public BUILDKITE_UNBLOCKER_ID(): string {
        return process.env.BUILDKITE_UNBLOCKER_ID;
    }
    // A colon separated list of non-private team slugs that the user who unblocked the build belongs to.
    public BUILDKITE_UNBLOCKER_TEAMS(): string[] {
        return process.env.BUILDKITE_UNBLOCKER_TEAMS.split(":");
    }
}


export default Environment;