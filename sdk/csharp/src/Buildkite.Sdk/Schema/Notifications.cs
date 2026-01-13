namespace Buildkite.Sdk.Schema;

/// <summary>
/// Email notification configuration.
/// </summary>
public class EmailNotification : INotification
{
    /// <summary>The email address to notify.</summary>
    public string? Email { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// Slack notification configuration.
/// </summary>
public class SlackNotification : INotification
{
    /// <summary>Slack channel or configuration.</summary>
    public object? Slack { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// Detailed Slack notification configuration.
/// </summary>
public class SlackConfig
{
    /// <summary>Slack channels to notify.</summary>
    public List<string>? Channels { get; set; }

    /// <summary>Message to send.</summary>
    public string? Message { get; set; }
}

/// <summary>
/// Webhook notification configuration.
/// </summary>
public class WebhookNotification : INotification
{
    /// <summary>The webhook URL to notify.</summary>
    public string? Webhook { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// PagerDuty notification configuration.
/// </summary>
public class PagerDutyNotification : INotification
{
    /// <summary>PagerDuty change event configuration.</summary>
    public string? PagerdutyChangeEvent { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// Basecamp notification configuration.
/// </summary>
public class BasecampNotification : INotification
{
    /// <summary>The Basecamp campfire URL.</summary>
    public string? BasecampCampfire { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// GitHub commit status notification.
/// </summary>
public class GitHubCommitStatusNotification : INotification
{
    /// <summary>GitHub commit status configuration.</summary>
    public GitHubCommitStatusConfig? GithubCommitStatus { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// Configuration for GitHub commit status.
/// </summary>
public class GitHubCommitStatusConfig
{
    /// <summary>The context for the commit status.</summary>
    public string? Context { get; set; }
}

/// <summary>
/// GitHub check notification.
/// </summary>
public class GitHubCheckNotification : INotification
{
    /// <summary>GitHub check configuration.</summary>
    public GitHubCheckConfig? GithubCheck { get; set; }

    /// <summary>A boolean expression to conditionally send this notification.</summary>
    public string? If { get; set; }
}

/// <summary>
/// Configuration for GitHub check.
/// </summary>
public class GitHubCheckConfig
{
    /// <summary>The context for the check.</summary>
    public string? Context { get; set; }
}
