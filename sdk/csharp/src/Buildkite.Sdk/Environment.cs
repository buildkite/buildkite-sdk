namespace Buildkite.Sdk;

/// <summary>
/// Provides strongly-typed access to Buildkite environment variables.
/// These environment variables are set by the Buildkite agent during build execution.
/// </summary>
public static class EnvironmentVariable
{
    // Build Information
    /// <summary>The UUID of the build</summary>
    public static string? BuildId => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_ID");

    /// <summary>The build number for this pipeline</summary>
    public static string? BuildNumber => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_NUMBER");

    /// <summary>The URL for the build</summary>
    public static string? BuildUrl => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_URL");

    /// <summary>The message associated with the build (commit message or custom)</summary>
    public static string? Message => System.Environment.GetEnvironmentVariable("BUILDKITE_MESSAGE");

    /// <summary>The creator of the build</summary>
    public static string? BuildCreator => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CREATOR");

    /// <summary>The email of the build creator</summary>
    public static string? BuildCreatorEmail => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CREATOR_EMAIL");

    /// <summary>The source of the build (e.g., webhook, ui, api, schedule)</summary>
    public static string? Source => System.Environment.GetEnvironmentVariable("BUILDKITE_SOURCE");

    // Repository Information
    /// <summary>The branch being built</summary>
    public static string? Branch => System.Environment.GetEnvironmentVariable("BUILDKITE_BRANCH");

    /// <summary>The commit SHA being built</summary>
    public static string? Commit => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMIT");

    /// <summary>The git tag being built, if any</summary>
    public static string? Tag => System.Environment.GetEnvironmentVariable("BUILDKITE_TAG");

    /// <summary>The repository URL</summary>
    public static string? Repo => System.Environment.GetEnvironmentVariable("BUILDKITE_REPO");

    // Pipeline Information
    /// <summary>The slug of the pipeline</summary>
    public static string? PipelineSlug => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_SLUG");

    /// <summary>The name of the pipeline</summary>
    public static string? PipelineName => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_NAME");

    /// <summary>The UUID of the pipeline</summary>
    public static string? PipelineId => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_ID");

    /// <summary>The default branch for this pipeline</summary>
    public static string? PipelineDefaultBranch => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_DEFAULT_BRANCH");

    /// <summary>The provider for the pipeline (e.g., github, bitbucket)</summary>
    public static string? PipelineProvider => System.Environment.GetEnvironmentVariable("BUILDKITE_PIPELINE_PROVIDER");

    // Organization Information
    /// <summary>The slug of the organization</summary>
    public static string? OrganizationSlug => System.Environment.GetEnvironmentVariable("BUILDKITE_ORGANIZATION_SLUG");

    // Agent Information
    /// <summary>The UUID of the agent running the job</summary>
    public static string? AgentId => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_ID");

    /// <summary>The name of the agent running the job</summary>
    public static string? AgentName => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_NAME");

    /// <summary>The agent's meta-data (tags)</summary>
    public static string? AgentMetaData => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_META_DATA");

    // Job Information
    /// <summary>The UUID of the current job</summary>
    public static string? JobId => System.Environment.GetEnvironmentVariable("BUILDKITE_JOB_ID");

    /// <summary>The label of the current step</summary>
    public static string? Label => System.Environment.GetEnvironmentVariable("BUILDKITE_LABEL");

    /// <summary>The key of the current step</summary>
    public static string? StepKey => System.Environment.GetEnvironmentVariable("BUILDKITE_STEP_KEY");

    /// <summary>The UUID of the current step</summary>
    public static string? StepId => System.Environment.GetEnvironmentVariable("BUILDKITE_STEP_ID");

    /// <summary>The command being run</summary>
    public static string? Command => System.Environment.GetEnvironmentVariable("BUILDKITE_COMMAND");

    /// <summary>The timeout for the command in minutes</summary>
    public static string? Timeout => System.Environment.GetEnvironmentVariable("BUILDKITE_TIMEOUT");

    // Parallelism
    /// <summary>The index of the current parallel job (0, 1, 2, ...)</summary>
    public static string? ParallelJob => System.Environment.GetEnvironmentVariable("BUILDKITE_PARALLEL_JOB");

    /// <summary>The total number of parallel jobs</summary>
    public static string? ParallelJobCount => System.Environment.GetEnvironmentVariable("BUILDKITE_PARALLEL_JOB_COUNT");

    // Retry Information
    /// <summary>The number of times this job has been retried</summary>
    public static string? RetryCount => System.Environment.GetEnvironmentVariable("BUILDKITE_RETRY_COUNT");

    /// <summary>Whether this build is a rebuild</summary>
    public static string? Rebuilt => System.Environment.GetEnvironmentVariable("BUILDKITE_REBUILT_FROM_BUILD_ID");

    // Pull Request Information
    /// <summary>The pull request number if this is a PR build</summary>
    public static string? PullRequest => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST");

    /// <summary>The base branch for the pull request</summary>
    public static string? PullRequestBaseBranch => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST_BASE_BRANCH");

    /// <summary>The repository URL of the pull request</summary>
    public static string? PullRequestRepo => System.Environment.GetEnvironmentVariable("BUILDKITE_PULL_REQUEST_REPO");

    // Paths
    /// <summary>The path to the build checkout directory</summary>
    public static string? BuildCheckoutPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_CHECKOUT_PATH");

    /// <summary>The path to the build directory</summary>
    public static string? BuildPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BUILD_PATH");

    /// <summary>The path to the agent's bin directory</summary>
    public static string? BinPath => System.Environment.GetEnvironmentVariable("BUILDKITE_BIN_PATH");

    /// <summary>The path to the agent's hooks directory</summary>
    public static string? HooksPath => System.Environment.GetEnvironmentVariable("BUILDKITE_HOOKS_PATH");

    /// <summary>The path to the agent's plugins directory</summary>
    public static string? PluginsPath => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGINS_PATH");

    // Artifact Paths
    /// <summary>The artifact paths to upload</summary>
    public static string? ArtifactPaths => System.Environment.GetEnvironmentVariable("BUILDKITE_ARTIFACT_PATHS");

    /// <summary>The S3 bucket for artifact uploads</summary>
    public static string? ArtifactUploadDestination => System.Environment.GetEnvironmentVariable("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION");

    // Feature Flags / Settings
    /// <summary>Whether the agent is in debug mode</summary>
    public static bool AgentDebug => System.Environment.GetEnvironmentVariable("BUILDKITE_AGENT_DEBUG") == "true";

    /// <summary>Whether git submodules should be cloned</summary>
    public static bool GitSubmodules => System.Environment.GetEnvironmentVariable("BUILDKITE_GIT_SUBMODULES") == "true";

    /// <summary>Whether the build is triggered</summary>
    public static bool TriggeredFromBuildId => !string.IsNullOrEmpty(System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_ID"));

    /// <summary>The build ID that triggered this build</summary>
    public static string? TriggeredFromBuildIdValue => System.Environment.GetEnvironmentVariable("BUILDKITE_TRIGGERED_FROM_BUILD_ID");

    /// <summary>Whether this is a Buildkite environment</summary>
    public static bool IsBuildkite => System.Environment.GetEnvironmentVariable("BUILDKITE") == "true";

    /// <summary>Whether running in CI</summary>
    public static bool IsCi => System.Environment.GetEnvironmentVariable("CI") == "true";

    // Group Information
    /// <summary>The UUID of the group step</summary>
    public static string? GroupId => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_ID");

    /// <summary>The key of the group step</summary>
    public static string? GroupKey => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_KEY");

    /// <summary>The label of the group step</summary>
    public static string? GroupLabel => System.Environment.GetEnvironmentVariable("BUILDKITE_GROUP_LABEL");

    // Unblocker Information
    /// <summary>The ID of the user who unblocked this step</summary>
    public static string? UnblockedById => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER_ID");

    /// <summary>The email of the user who unblocked this step</summary>
    public static string? UnblockerEmail => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER_EMAIL");

    /// <summary>The name of the user who unblocked this step</summary>
    public static string? UnblockerName => System.Environment.GetEnvironmentVariable("BUILDKITE_UNBLOCKER");

    // Plugin Information
    /// <summary>The JSON representation of plugins for the current step</summary>
    public static string? Plugins => System.Environment.GetEnvironmentVariable("BUILDKITE_PLUGINS");
}
