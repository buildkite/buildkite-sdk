namespace Buildkite.Sdk;

/// <summary>
/// Provides strongly-typed access to Buildkite environment variables.
/// These environment variables are set by the Buildkite agent during build execution.
/// </summary>
public static class EnvironmentVariable
{
    // Build Information
    /// <summary>The UUID of the build.</summary>
    public static string? BuildId => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_ID");

    /// <summary>The build number. This number increases with every build, and is guaranteed to be unique within each pipeline.</summary>
    public static string? BuildNumber => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_NUMBER");

    /// <summary>The URL for this build on Buildkite.</summary>
    public static string? BuildUrl => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_URL");

    /// <summary>The message associated with the build, usually the commit message or the message provided when the build is triggered.</summary>
    public static string? Message => System.Environment.GetEnvironmentVariable("BUILDKITE_MESSAGE");

    /// <summary>The name of the user who authored the commit being built. May be unverified.</summary>
    public static string? BuildAuthor => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_AUTHOR");

    /// <summary>The notification email of the user who authored the commit being built. May be unverified.</summary>
    public static string? BuildAuthorEmail => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_AUTHOR_EMAIL");

    /// <summary>The name of the user who created the build.</summary>
    public static string? BuildCreator => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CREATOR");

    /// <summary>The notification email of the user who created the build.</summary>
    public static string? BuildCreatorEmail => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CREATOR_EMAIL");

    /// <summary>A colon-separated list of non-private team slugs that the build creator belongs to.</summary>
    public static string? BuildCreatorTeams => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CREATOR_TEAMS");

    /// <summary>The source of the event that created the build (e.g., webhook, ui, api, schedule).</summary>
    public static string? Source => System.Environment.GetEnvironmentVariable("BUILDKITE_SOURCE");

    /// <summary>The UUID value of the cluster, if the build has an associated cluster_queue.</summary>
    public static string? ClusterId => System.Environment.GetEnvironmentVariable("BUILDKITE_CLUSTER_ID");

    // Repository Information
    /// <summary>The branch being built.</summary>
    public static string? Branch => System.Environment.GetEnvironmentVariable("BUILDKITE_BRANCH");

    /// <summary>The git commit object of the build. Usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like HEAD.</summary>
    public static string? Commit => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMIT");

    /// <summary>The name of the tag being built, if this build was triggered from a tag.</summary>
    public static string? Tag => System.Environment.GetEnvironmentVariable("BUILDKITE_TAG");

    /// <summary>The repository URL of the pipeline.</summary>
    public static string? Repo => System.Environment.GetEnvironmentVariable("BUILDKITE_REPO");

    /// <summary>The path to the shared git mirror.</summary>
    public static string? RepoMirror => System.Environment.GetEnvironmentVariable("BUILDKITE_REPO_MIRROR");

    /// <summary>A custom refspec for the buildkite-agent bootstrap script to use when checking out code.</summary>
    public static string? Refspec => System.Environment.GetEnvironmentVariable("BUILDKITE_REFSPEC");

    /// <summary>Whether the build should perform a clean checkout.</summary>
    public static bool CleanCheckout => System.Environment.GetEnvironmentVariable("BUILDKITE_CLEAN_CHECKOUT") == "true";

    /// <summary>Whether git submodules should be cloned.</summary>
    public static bool GitSubmodules => System.Environment.GetEnvironmentVariable("BUILDKITE_GIT_SUBMODULES") == "true";

    /// <summary>The value of the git-clean-flags agent configuration option.</summary>
    public static string? GitCleanFlags => System.Environment.GetEnvironmentVariable("BUILDKITE_GIT_CLEAN_FLAGS");

    /// <summary>The value of the git-clone-flags agent configuration option.</summary>
    public static string? GitCloneFlags => System.Environment.GetEnvironmentVariable("BUILDKITE_GIT_CLONE_FLAGS");

    // Pipeline Information
    /// <summary>The pipeline slug on Buildkite as used in URLs.</summary>
    public static string? PipelineSlug => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_SLUG");

    /// <summary>The displayed pipeline name on Buildkite.</summary>
    public static string? PipelineName => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_NAME");

    /// <summary>The UUID of the pipeline.</summary>
    public static string? PipelineId => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_ID");

    /// <summary>The default branch for this pipeline.</summary>
    public static string? PipelineDefaultBranch => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_DEFAULT_BRANCH");

    /// <summary>The ID of the source code provider for the pipeline's repository.</summary>
    public static string? PipelineProvider => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_PROVIDER");

    /// <summary>A colon-separated list of the pipeline's non-private team slugs.</summary>
    public static string? PipelineTeams => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_TEAMS");

    // Organization Information
    /// <summary>The organization name on Buildkite as used in URLs.</summary>
    public static string? OrganizationSlug => System.Environment.GetEnvironmentVariable("BUILDKITE_ORGANIZATION_SLUG");

    /// <summary>The UUID of the organization.</summary>
    public static string? OrganizationId => System.Environment.GetEnvironmentVariable("BUILDKITE_ORGANIZATION_ID");

    // Agent Information
    /// <summary>The UUID of the agent.</summary>
    public static string? AgentId => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_ID");

    /// <summary>The name of the agent that ran the job.</summary>
    public static string? AgentName => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_NAME");

    /// <summary>The process ID of the agent.</summary>
    public static string? AgentPid => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_PID");

    /// <summary>The agent's meta-data (tags).</summary>
    public static string? AgentMetaData => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_META_DATA");

    /// <summary>The agent session token for the job. Used by the agent artifact and meta-data commands.</summary>
    public static string? AgentAccessToken => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_ACCESS_TOKEN");

    /// <summary>Whether the agent is in debug mode.</summary>
    public static bool AgentDebug => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_DEBUG") == "true";

    /// <summary>The value of the disconnect-after-job agent configuration option.</summary>
    public static string? AgentDisconnectAfterJob => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB");

    /// <summary>The value of the disconnect-after-idle-timeout agent configuration option.</summary>
    public static string? AgentDisconnectAfterIdleTimeout => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT");

    /// <summary>The value of the endpoint agent configuration option.</summary>
    public static string? AgentEndpoint => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_ENDPOINT");

    /// <summary>A list of the experimental agent features that are currently enabled.</summary>
    public static string? AgentExperiment => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_EXPERIMENT");

    /// <summary>The value of the health-check-addr agent configuration option.</summary>
    public static string? AgentHealthCheckAddr => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_HEALTH_CHECK_ADDR");

    // Job Information
    /// <summary>The internal UUID Buildkite uses for this job.</summary>
    public static string? JobId => System.Environment.GetEnvironmentVariable("BUILDKITE_JOB_ID");

    /// <summary>The label/name of the current job.</summary>
    public static string? Label => System.Environment.GetEnvironmentVariable("BUILDKITE_LABEL");

    /// <summary>The value of the key attribute, a unique string set by you to identify a step.</summary>
    public static string? StepKey => System.Environment.GetEnvironmentVariable("BUILDKITE_STEP_KEY");

    /// <summary>A unique string that identifies a step.</summary>
    public static string? StepId => System.Environment.GetEnvironmentVariable("BUILDKITE_STEP_ID");

    /// <summary>The command that will be run for the job.</summary>
    public static string? Command => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMAND");

    /// <summary>The exit code from the last command run in the command hook.</summary>
    public static string? CommandExitStatus => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMAND_EXIT_STATUS");

    /// <summary>Whether command evaluation is enabled (opposite of no-command-eval agent config).</summary>
    public static bool CommandEval => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMAND_EVAL") == "true";

    /// <summary>The number of minutes until Buildkite automatically cancels this job, or false if no timeout is set.</summary>
    public static string? Timeout => System.Environment.GetEnvironmentVariable("BUILDKITE_TIMEOUT");

    /// <summary>The path to a temporary file containing the logs for this job. Requires enabling the enable-job-log-tmpfile agent configuration option.</summary>
    public static string? JobLogTmpfile => System.Environment.GetEnvironmentVariable("BUILDKITE_JOB_LOG_TMPFILE");

    /// <summary>The exit code of the last hook that ran, used internally by the hooks.</summary>
    public static string? LastHookExitStatus => System.Environment.GetEnvironmentVariable("BUILDKITE_LAST_HOOK_EXIT_STATUS");

    // Parallelism
    /// <summary>The index of the current parallel job (0, 1, 2, ...).</summary>
    public static string? ParallelJob => System.Environment.GetEnvironmentVariable("BUILDKITE_PARALLEL_JOB");

    /// <summary>The total number of parallel jobs created from a parallel build step.</summary>
    public static string? ParallelJobCount => System.Environment.GetEnvironmentVariable("BUILDKITE_PARALLEL_JOB_COUNT");

    // Retry Information
    /// <summary>How many times this job has been retried.</summary>
    public static string? RetryCount => System.Environment.GetEnvironmentVariable("BUILDKITE_RETRY_COUNT");

    /// <summary>Whether this build is a rebuild.</summary>
    public static string? Rebuilt => System.Environment.GetEnvironmentVariable("BUILDKITE_REBUILT_FROM_BUILD_ID");

    /// <summary>The number of the original build this was rebuilt from, or empty if not a rebuild.</summary>
    public static string? RebuiltFromBuildNumber => System.Environment.GetEnvironmentVariable("BUILDKITE_REBUILT_FROM_BUILD_NUMBER");

    // Pull Request Information
    /// <summary>The pull request number, or "false" if not a pull request.</summary>
    public static string? PullRequest => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST");

    /// <summary>The base branch that the pull request is targeting, or empty if not a pull request.</summary>
    public static string? PullRequestBaseBranch => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST_BASE_BRANCH");

    /// <summary>The repository URL of the pull request, or empty if not a pull request.</summary>
    public static string? PullRequestRepo => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST_REPO");

    /// <summary>Whether the pull request is a draft.</summary>
    public static bool PullRequestDraft => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST_DRAFT") == "true";

    // Triggered Build Information
    /// <summary>Whether the build is triggered.</summary>
    public static bool TriggeredFromBuildId => !string.IsNullOrEmpty(System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_ID"));

    /// <summary>The build ID that triggered this build.</summary>
    public static string? TriggeredFromBuildIdValue => System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_ID");

    /// <summary>The number of the build that triggered this build, or empty if not triggered.</summary>
    public static string? TriggeredFromBuildNumber => System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER");

    /// <summary>The slug of the pipeline that was used to trigger this build, or empty if not triggered.</summary>
    public static string? TriggeredFromBuildPipelineSlug => System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG");

    // Paths
    /// <summary>The path where the agent has checked out your code for this build.</summary>
    public static string? BuildCheckoutPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CHECKOUT_PATH");

    /// <summary>The value of the build-path agent configuration option.</summary>
    public static string? BuildPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_PATH");

    /// <summary>The path to the directory containing the buildkite-agent binary.</summary>
    public static string? BinPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BIN_PATH");

    /// <summary>The value of the hooks-path agent configuration option.</summary>
    public static string? HooksPath => System.Environment.GetEnvironmentVariable("BUILDKITE_HOOKS_PATH");

    /// <summary>The value of the plugins-path agent configuration option.</summary>
    public static string? PluginsPath => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGINS_PATH");

    /// <summary>The path to the agent config file.</summary>
    public static string? ConfigPath => System.Environment.GetEnvironmentVariable("BUILDKITE_CONFIG_PATH");

    /// <summary>The path to the file containing the job's environment variables.</summary>
    public static string? EnvFile => System.Environment.GetEnvironmentVariable("BUILDKITE_ENV_FILE");

    // Artifacts
    /// <summary>The artifact paths to upload after the job, if any have been specified.</summary>
    public static string? ArtifactPaths => System.Environment.GetEnvironmentVariable("BUILDKITE_ARTIFACT_PATHS");

    /// <summary>The path where artifacts will be uploaded.</summary>
    public static string? ArtifactUploadDestination => System.Environment.GetEnvironmentVariable("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION");

    // S3 Configuration
    /// <summary>The access key ID for your S3 IAM user, for use with private S3 buckets.</summary>
    public static string? S3AccessKeyId => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_ACCESS_KEY_ID");

    /// <summary>The access URL for your private S3 bucket, if you are using a proxy.</summary>
    public static string? S3AccessUrl => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_ACCESS_URL");

    /// <summary>The Access Control List to be set on artifacts being uploaded to your private S3 bucket.</summary>
    public static string? S3Acl => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_ACL");

    /// <summary>The region of your private S3 bucket.</summary>
    public static string? S3DefaultRegion => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_DEFAULT_REGION");

    /// <summary>The secret access key for your S3 IAM user, for use with private S3 buckets.</summary>
    public static string? S3SecretAccessKey => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_SECRET_ACCESS_KEY");

    /// <summary>Whether to enable encryption for the artifacts in your private S3 bucket.</summary>
    public static bool S3SseEnabled => System.Environment.GetEnvironmentVariable("BUILDKITE_S3_SSE_ENABLED") == "true";

    // Cancel Configuration
    /// <summary>The value of the cancel-grace-period agent configuration option in seconds.</summary>
    public static string? CancelGracePeriod => System.Environment.GetEnvironmentVariable("BUILDKITE_CANCEL_GRACE_PERIOD");

    /// <summary>The value of the cancel-signal agent configuration option.</summary>
    public static string? CancelSignal => System.Environment.GetEnvironmentVariable("BUILDKITE_CANCEL_SIGNAL");

    // Plugin Information
    /// <summary>A JSON object containing a list of plugins used in the step, and their configuration.</summary>
    public static string? Plugins => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGINS");

    /// <summary>Whether plugins are enabled (opposite of no-plugins agent config).</summary>
    public static bool PluginsEnabled => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGINS_ENABLED") == "true";

    /// <summary>A JSON string holding the current plugin's configuration.</summary>
    public static string? PluginConfiguration => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGIN_CONFIGURATION");

    /// <summary>The current plugin's name, with all letters in uppercase and spaces replaced with underscores.</summary>
    public static string? PluginName => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGIN_NAME");

    /// <summary>Whether to validate plugin configuration and requirements.</summary>
    public static bool PluginValidation => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGIN_VALIDATION") == "true";

    // Group Information
    /// <summary>The UUID of the group step the job belongs to.</summary>
    public static string? GroupId => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_ID");

    /// <summary>The value of the key attribute of the group step the job belongs to.</summary>
    public static string? GroupKey => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_KEY");

    /// <summary>The label/name of the group step the job belongs to.</summary>
    public static string? GroupLabel => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_LABEL");

    // Unblocker Information
    /// <summary>The ID of the user who unblocked this step.</summary>
    public static string? UnblockedById => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER_ID");

    /// <summary>The notification email of the user who unblocked the build.</summary>
    public static string? UnblockerEmail => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER_EMAIL");

    /// <summary>The name of the user who unblocked the build.</summary>
    public static string? UnblockerName => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER");

    /// <summary>A colon-separated list of non-private team slugs that the user who unblocked the build belongs to.</summary>
    public static string? UnblockerTeams => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER_TEAMS");

    // GitHub Deployment Information
    /// <summary>The GitHub deployment ID. Only available on builds triggered by a GitHub Deployment.</summary>
    public static string? GithubDeploymentId => System.Environment.GetEnvironmentVariable("BUILDKITE_GITHUB_DEPLOYMENT_ID");

    /// <summary>The name of the GitHub deployment environment. Only available on builds triggered by a GitHub Deployment.</summary>
    public static string? GithubDeploymentEnvironment => System.Environment.GetEnvironmentVariable("BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT");

    /// <summary>The name of the GitHub deployment task. Only available on builds triggered by a GitHub Deployment.</summary>
    public static string? GithubDeploymentTask => System.Environment.GetEnvironmentVariable("BUILDKITE_GITHUB_DEPLOYMENT_TASK");

    /// <summary>The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a GitHub Deployment.</summary>
    public static string? GithubDeploymentPayload => System.Environment.GetEnvironmentVariable("BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD");

    // Shell and Hooks
    /// <summary>The value of the shell agent configuration option.</summary>
    public static string? Shell => System.Environment.GetEnvironmentVariable("BUILDKITE_SHELL");

    /// <summary>Whether local hooks are enabled (opposite of no-local-hooks agent config).</summary>
    public static bool LocalHooksEnabled => System.Environment.GetEnvironmentVariable("BUILDKITE_LOCAL_HOOKS_ENABLED") == "true";

    /// <summary>Whether SSH keyscan is enabled (opposite of no-ssh-keyscan agent config).</summary>
    public static bool SshKeyscan => System.Environment.GetEnvironmentVariable("BUILDKITE_SSH_KEYSCAN") == "true";

    // Tracing
    /// <summary>The tracing backend for the agent (e.g., "datadog").</summary>
    public static string? TracingBackend => System.Environment.GetEnvironmentVariable("BUILDKITE_TRACING_BACKEND");

    // Miscellaneous
    /// <summary>A list of environment variables that are protected and will be overridden, used internally by the agent.</summary>
    public static string? IgnoredEnv => System.Environment.GetEnvironmentVariable("BUILDKITE_IGNORED_ENV");

    /// <summary>Whether this is a Buildkite environment.</summary>
    public static bool IsBuildkite => System.Environment.GetEnvironmentVariable("BUILDKITE") == "true";

    /// <summary>Whether running in CI.</summary>
    public static bool IsCi => System.Environment.GetEnvironmentVariable("CI") == "true";
}
