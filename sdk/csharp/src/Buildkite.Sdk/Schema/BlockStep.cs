namespace Buildkite.Sdk.Schema;

/// <summary>
/// A block step pauses the build and requires manual unblocking.
/// </summary>
public class BlockStep : IStep, IGroupStep
{
    /// <summary>The label for the block step.</summary>
    public string? Block { get; set; }

    /// <summary>Alias for Block.</summary>
    public string? Label { get; set; }

    /// <summary>Alias for Block.</summary>
    public string? Name { get; set; }

    /// <summary>A unique identifier for this step.</summary>
    public string? Key { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Id { get; set; }

    /// <summary>Alias for Key.</summary>
    public string? Identifier { get; set; }

    /// <summary>The message displayed in the unblock dialog.</summary>
    public string? Prompt { get; set; }

    /// <summary>Input fields to display in the unblock dialog.</summary>
    public List<Field>? Fields { get; set; }

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

    /// <summary>Team slugs or IDs permitted to unblock this step.</summary>
    public OneOrMany<string>? AllowedTeams { get; set; }

    /// <summary>The build state when blocked: 'passed', 'failed', or 'running'.</summary>
    public string? BlockedState { get; set; }

    /// <summary>The step type.</summary>
    public string Type => "block";
}

/// <summary>
/// A field in a block or input step.
/// </summary>
public abstract class Field
{
    /// <summary>The key to store the field value.</summary>
    public string? Key { get; set; }

    /// <summary>Whether this field is required.</summary>
    public bool? Required { get; set; }

    /// <summary>Hint text displayed below the field.</summary>
    public string? Hint { get; set; }
}

/// <summary>
/// A text input field.
/// </summary>
public class TextField : Field
{
    /// <summary>The label for the text field.</summary>
    public string? Text { get; set; }

    /// <summary>The default value.</summary>
    public string? Default { get; set; }
}

/// <summary>
/// A select (dropdown) field.
/// </summary>
public class SelectField : Field
{
    /// <summary>The label for the select field.</summary>
    public string? Select { get; set; }

    /// <summary>The options for the select field.</summary>
    public List<SelectOption>? Options { get; set; }

    /// <summary>The default value.</summary>
    public string? Default { get; set; }

    /// <summary>Whether multiple options can be selected.</summary>
    public bool? Multiple { get; set; }
}

/// <summary>
/// An option in a select field.
/// </summary>
public class SelectOption
{
    /// <summary>The display label for the option.</summary>
    public string? Label { get; set; }

    /// <summary>The value for the option.</summary>
    public string? Value { get; set; }
}
