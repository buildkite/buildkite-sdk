using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// A wait step waits for all previous steps to complete before continuing.
/// </summary>
public class WaitStep : IStep, IGroupStep
{
    /// <summary>Optional label for the wait step.</summary>
    public string? Wait { get; set; } = "";

    /// <summary>Alias for Wait.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Label { get => Wait; set => Wait = value; }

    /// <summary>Alias for Wait.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Name { get => Wait; set => Wait = value; }

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

    /// <summary>A boolean expression to conditionally run this step.</summary>
    public string? If { get; set; }

    /// <summary>Branch filter pattern.</summary>
    public object? Branches { get; set; }

    /// <summary>Step keys this step depends on.</summary>
    public DependsOn? DependsOn { get; set; }

    /// <summary>Whether to proceed if a dependency fails.</summary>
    public bool? AllowDependencyFailure { get; set; }

    /// <summary>Continue even if previous steps failed.</summary>
    public bool? ContinueOnFailure { get; set; }
}
