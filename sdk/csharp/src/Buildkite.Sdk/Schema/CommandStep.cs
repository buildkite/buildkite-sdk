namespace Buildkite.Sdk.Schema;

/// <summary>
/// A command step runs one or more shell commands on an agent.
/// </summary>
public class CommandStep : IStep, IGroupStep
{
    /// <summary>The label displayed in the Buildkite UI. Supports emoji.</summary>
    public string? Label { get; set; }

    /// <summary>Alias for Label.</summary>
    public string? Name { get; set; }

    /// <summary>A unique identifier for this step.</summary>
    public string? Key { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Id { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Identifier { get; set; }

    /// <summary>The shell command(s) to run.</summary>
    public OneOrMany<string>? Command { get; set; }

    /// <summary>Alias for Command.</summary>
    public OneOrMany<string>? Commands { get; set; }

    /// <summary>Agent query rules for targeting specific agents.</summary>
    public OneOf<AgentsObject, AgentsList>? Agents { get; set; }

    /// <summary>Environment variables for this step.</summary>
    public Dictionary<string, object?>? Env { get; set; }

    /// <summary>Branch filter pattern.</summary>
    public OneOrMany<string>? Branches { get; set; }

    /// <summary>A boolean expression to conditionally run this step.</summary>
    public string? If { get; set; }

    /// <summary>Glob pattern to run step only if matching files changed.</summary>
    public OneOrMany<string>? IfChanged { get; set; }

    /// <summary>Step keys this step depends on.</summary>
    public OneOf<string, List<string>, List<Dependency>>? DependsOn { get; set; }

    /// <summary>Whether to proceed if a dependency fails.</summary>
    public bool? AllowDependencyFailure { get; set; }

    /// <summary>Whether to skip this step. Can be bool or string reason.</summary>
    public StringOr<bool>? Skip { get; set; }

    /// <summary>Retry configuration.</summary>
    public Retry? Retry { get; set; }

    /// <summary>Soft fail configuration.</summary>
    public BoolOr<List<SoftFailCondition>>? SoftFail { get; set; }

    /// <summary>Maximum time in minutes for the job to run.</summary>
    public int? TimeoutInMinutes { get; set; }

    /// <summary>Glob patterns for artifacts to upload.</summary>
    public OneOrMany<string>? ArtifactPaths { get; set; }

    /// <summary>The number of parallel jobs to create.</summary>
    public int? Parallelism { get; set; }

    /// <summary>Maximum jobs from this step that can run concurrently.</summary>
    public int? Concurrency { get; set; }

    /// <summary>A unique name for the concurrency group.</summary>
    public string? ConcurrencyGroup { get; set; }

    /// <summary>Control command order: 'ordered' (default) or 'eager'.</summary>
    public string? ConcurrencyMethod { get; set; }

    /// <summary>Priority of the job (higher = more priority).</summary>
    public int? Priority { get; set; }

    /// <summary>Plugins to use with this step.</summary>
    public List<PluginConfig>? Plugins { get; set; }

    /// <summary>Matrix configuration for multiple job variations.</summary>
    public Matrix? Matrix { get; set; }

    /// <summary>Notifications for this step.</summary>
    public List<INotification>? Notify { get; set; }

    /// <summary>Cache configuration.</summary>
    public OneOf<Cache, List<Cache>>? Cache { get; set; }

    /// <summary>Cancel the job if the build is marked as failing.</summary>
    public bool? CancelOnBuildFailing { get; set; }

    /// <summary>Secrets to make available to the step.</summary>
    public OneOrMany<string>? Secrets { get; set; }

    /// <summary>Signature for signed pipelines.</summary>
    public Signature? Signature { get; set; }

    /// <summary>Container image for Kubernetes stack.</summary>
    public string? Image { get; set; }
}
