namespace Buildkite.Sdk.Schema;

/// <summary>
/// Represents a complete Buildkite pipeline configuration.
/// </summary>
public class BuildkitePipeline
{
    /// <summary>Default agents for all steps in the pipeline.</summary>
    public AgentsObject? Agents { get; set; }

    /// <summary>Environment variables for all steps in the pipeline.</summary>
    public Dictionary<string, object?>? Env { get; set; }

    /// <summary>Notifications for the pipeline.</summary>
    public List<INotification>? Notify { get; set; }

    /// <summary>The steps in the pipeline.</summary>
    public List<IStep>? Steps { get; set; }

    /// <summary>Secrets available to the pipeline.</summary>
    public List<string>? Secrets { get; set; }

    /// <summary>Container image for Kubernetes stack.</summary>
    public string? Image { get; set; }
}
