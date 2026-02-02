using System.Text.Json.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// Marker interface for all pipeline step types.
/// </summary>
[JsonDerivedType(typeof(CommandStep))]
[JsonDerivedType(typeof(BlockStep))]
[JsonDerivedType(typeof(InputStep))]
[JsonDerivedType(typeof(WaitStep))]
[JsonDerivedType(typeof(TriggerStep))]
[JsonDerivedType(typeof(GroupStep))]
public interface IStep { }

/// <summary>
/// Marker interface for all notification types.
/// </summary>
[JsonDerivedType(typeof(EmailNotification))]
[JsonDerivedType(typeof(SlackNotification))]
[JsonDerivedType(typeof(WebhookNotification))]
[JsonDerivedType(typeof(PagerDutyNotification))]
[JsonDerivedType(typeof(BasecampNotification))]
[JsonDerivedType(typeof(GitHubCommitStatusNotification))]
[JsonDerivedType(typeof(GitHubCheckNotification))]
public interface INotification { }

/// <summary>
/// Marker interface for steps that can appear inside a group.
/// </summary>
[JsonDerivedType(typeof(CommandStep))]
[JsonDerivedType(typeof(BlockStep))]
[JsonDerivedType(typeof(InputStep))]
[JsonDerivedType(typeof(WaitStep))]
[JsonDerivedType(typeof(TriggerStep))]
public interface IGroupStep : IStep { }
