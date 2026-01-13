namespace Buildkite.Sdk.Schema;

/// <summary>
/// An input step collects information from the user before continuing.
/// </summary>
public class InputStep : IStep, IGroupStep
{
    /// <summary>The label for the input step.</summary>
    public string? Input { get; set; }

    /// <summary>Alias for Input.</summary>
    public string? Label { get; set; }

    /// <summary>Alias for Input.</summary>
    public string? Name { get; set; }

    /// <summary>A unique identifier for this step.</summary>
    public string? Key { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Id { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Identifier { get; set; }

    /// <summary>The message displayed in the input dialog.</summary>
    public string? Prompt { get; set; }

    /// <summary>Input fields to display.</summary>
    public List<Field>? Fields { get; set; }

    /// <summary>Branch filter pattern.</summary>
    public OneOrMany<string>? Branches { get; set; }

    /// <summary>A boolean expression to conditionally run this step.</summary>
    public string? If { get; set; }

    /// <summary>Step keys this step depends on.</summary>
    public OneOf<string, List<string>, List<Dependency>>? DependsOn { get; set; }

    /// <summary>Whether to proceed if a dependency fails.</summary>
    public bool? AllowDependencyFailure { get; set; }

    /// <summary>Team slugs or IDs permitted to provide input.</summary>
    public OneOrMany<string>? AllowedTeams { get; set; }

    /// <summary>The build state when waiting for input: 'passed', 'failed', or 'running'.</summary>
    public string? BlockedState { get; set; }

    /// <summary>The step type.</summary>
    public string Type => "input";
}
