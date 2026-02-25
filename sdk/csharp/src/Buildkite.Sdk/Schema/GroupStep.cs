using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// A group step organizes steps into a collapsible group.
/// </summary>
public class GroupStep : IStep
{
    /// <summary>The label for the group.</summary>
    public string? Group { get; set; }

    /// <summary>Alias for Group.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Label { get => Group; set => Group = value; }

    /// <summary>Alias for Group.</summary>
    [JsonIgnore]
    [YamlIgnore]
    public string? Name { get => Group; set => Group = value; }

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

    /// <summary>The steps within this group.</summary>
    public List<IGroupStep> Steps { get; set; } = new();

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

    /// <summary>Notifications for this group.</summary>
    public List<INotification>? Notify { get; set; }

    /// <summary>Add a step to this group.</summary>
    public GroupStep AddStep(IGroupStep step)
    {
        Steps.Add(step);
        return this;
    }
}
