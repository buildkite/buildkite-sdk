namespace Buildkite.Sdk.Schema;

/// <summary>
/// A wait step waits for all previous steps to complete before continuing.
/// </summary>
public class WaitStep : IStep, IGroupStep
{
    /// <summary>Optional label for the wait step.</summary>
    public string? Wait { get; set; } = "";

    /// <summary>Alias for Wait.</summary>
    public string? Label { get; set; }

    /// <summary>Alias for Wait.</summary>
    public string? Name { get; set; }

    /// <summary>A unique identifier for this step.</summary>
    public string? Key { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Id { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Identifier { get; set; }

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
