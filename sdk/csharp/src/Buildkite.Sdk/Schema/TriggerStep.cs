using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// A trigger step creates a build on another pipeline.
/// </summary>
public class TriggerStep : IStep, IGroupStep
{
    /// <summary>The slug of the pipeline to trigger.</summary>
    public string? Trigger { get; set; }

    /// <summary>The label displayed in the Buildkite UI.</summary>
    public string? Label { get; set; }

    /// <summary>Alias for Label.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Name { get => Label; set => Label = value; }

    /// <summary>A unique identifier for this step.</summary>
    public string? Key { get; set; }

    /// <summary>Alias for Key.</summary>
    [JsonIgnore]
    [YamlIgnore]
    [Obsolete("Use Key instead.")]
    public string? Id { get => Key; set => Key = value; }

    /// <summary>Alias for Key.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Identifier { get => Key; set => Key = value; }

    /// <summary>Whether to run the triggered build asynchronously.</summary>
    public bool? Async { get; set; }

    /// <summary>Properties for the triggered build.</summary>
    public TriggerBuild? Build { get; set; }

    /// <summary>Branch filter pattern.</summary>
    public object? Branches { get; set; }

    /// <summary>A boolean expression to conditionally run this step.</summary>
    public string? If { get; set; }

    /// <summary>Glob pattern to run step only if matching files changed.</summary>
    public object? IfChanged { get; set; }

    /// <summary>Step keys this step depends on.</summary>
    public object? DependsOn { get; set; }

    /// <summary>Whether to proceed if a dependency fails.</summary>
    public bool? AllowDependencyFailure { get; set; }

    /// <summary>Whether to skip this step.</summary>
    public object? Skip { get; set; }

    /// <summary>Soft fail configuration.</summary>
    public object? SoftFail { get; set; }
}

/// <summary>
/// Properties for a triggered build.
/// </summary>
public class TriggerBuild
{
    /// <summary>The branch to build.</summary>
    public string? Branch { get; set; }

    /// <summary>The commit to build.</summary>
    public string? Commit { get; set; }

    /// <summary>The build message.</summary>
    public string? Message { get; set; }

    /// <summary>Environment variables for the triggered build.</summary>
    public Dictionary<string, string>? Env { get; set; }

    /// <summary>Metadata for the triggered build.</summary>
    public Dictionary<string, string>? MetaData { get; set; }
}
