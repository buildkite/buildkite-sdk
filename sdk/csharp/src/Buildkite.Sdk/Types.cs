using System.Collections.Generic;

namespace Buildkite.Sdk
{
    /// <summary>
    /// The state that the build is set to when the build is blocked by this block step
    /// </summary>
    public enum BlockedState
    {
        Failed,
        Passed,
        Running
    }

    /// <summary>
    /// Control command order, allowed values are 'ordered' (default) and 'eager'.
    /// If you use this attribute, you must also define concurrency_group and concurrency.
    /// </summary>
    public enum ConcurrencyMethod
    {
        Eager,
        Ordered
    }

    public enum NotifyEnum
    {
        GithubCheck,
        GithubCommitStatus
    }

    /// <summary>
    /// The exit signal reason, if any, that may be retried
    /// </summary>
    public enum SignalReason
    {
        AgentRefused,
        AgentStop,
        Cancel,
        Empty,
        None,
        ProcessRunError,
        SignatureRejected
    }

    public class DependsOn
    {
        public bool? AllowFailure { get; set; }
        public string? Step { get; set; }
    }

    /// <summary>
    /// A list of input fields required to be filled out before unblocking the step
    /// </summary>
    public class Field
    {
        /// <summary>
        /// The value that is pre-filled in the text field or pre-selected in the dropdown
        /// </summary>
        public object? Default { get; set; }

        /// <summary>
        /// The format must be a regular expression implicitly anchored to the beginning and end of
        /// the input and is functionally equivalent to the HTML5 pattern attribute.
        /// </summary>
        public string? Format { get; set; }

        /// <summary>
        /// The explanatory text that is shown after the label
        /// </summary>
        public string? Hint { get; set; }

        /// <summary>
        /// The meta-data key that stores the field's input
        /// </summary>
        public string Key { get; set; } = string.Empty;

        /// <summary>
        /// Whether the field is required for form submission
        /// </summary>
        public bool? Required { get; set; }

        /// <summary>
        /// The text input name
        /// </summary>
        public string? Text { get; set; }

        /// <summary>
        /// Whether more than one option may be selected
        /// </summary>
        public bool? Multiple { get; set; }

        public List<Option>? Options { get; set; }

        /// <summary>
        /// The select input name
        /// </summary>
        public string? Select { get; set; }
    }

    public class Option
    {
        /// <summary>
        /// The text displayed directly under the select field's label
        /// </summary>
        public string? Hint { get; set; }

        /// <summary>
        /// The text displayed on the select list item
        /// </summary>
        public string Label { get; set; } = string.Empty;

        /// <summary>
        /// Whether the field is required for form submission
        /// </summary>
        public bool? Required { get; set; }

        /// <summary>
        /// The value to be stored as meta-data
        /// </summary>
        public string Value { get; set; } = string.Empty;
    }

    public class CacheObject
    {
        public string? Name { get; set; }
        public List<string> Paths { get; set; } = new List<string>();
        public string? Size { get; set; }
    }

    public class SoftFail
    {
        /// <summary>
        /// The exit status number that will cause this job to soft-fail
        /// </summary>
        public object? ExitStatus { get; set; }
    }

    /// <summary>
    /// An adjustment to a Build Matrix
    /// </summary>
    public class Adjustment
    {
        public object? Skip { get; set; }
        public object? SoftFail { get; set; }
        public object? With { get; set; }
    }

    /// <summary>
    /// Configuration for multi-dimension Build Matrix
    /// </summary>
    public class MatrixObject
    {
        /// <summary>
        /// List of Build Matrix adjustments
        /// </summary>
        public List<Adjustment>? Adjustments { get; set; }

        public object? Setup { get; set; }
    }

    public class GithubCommitStatus
    {
        /// <summary>
        /// GitHub commit status name
        /// </summary>
        public string? Context { get; set; }
    }

    public class NotifySlack
    {
        public List<string>? Channels { get; set; }
        public string? Message { get; set; }
    }

    public class Notify
    {
        public string? BasecampCampfire { get; set; }
        public string? If { get; set; }
        public object? Slack { get; set; }
        public GithubCommitStatus? GithubCommitStatus { get; set; }
        public Dictionary<string, object>? GithubCheck { get; set; }
    }

    public class PipelineNotify
    {
        public string? Email { get; set; }
        public string? If { get; set; }
        public string? BasecampCampfire { get; set; }
        public object? Slack { get; set; }
        public string? Webhook { get; set; }
        public string? PagerdutyChangeEvent { get; set; }
        public GithubCommitStatus? GithubCommitStatus { get; set; }
        public Dictionary<string, object>? GithubCheck { get; set; }
    }

    /// <summary>
    /// The conditions for retrying this step.
    /// </summary>
    public class Retry
    {
        /// <summary>
        /// Whether to allow a job to retry automatically. If set to true, the retry conditions are
        /// set to the default value.
        /// </summary>
        public object? Automatic { get; set; }

        /// <summary>
        /// Whether to allow a job to be retried manually
        /// </summary>
        public object? Manual { get; set; }
    }

    public class AutomaticRetry
    {
        /// <summary>
        /// The exit status number that will cause this job to retry
        /// </summary>
        public object? ExitStatus { get; set; }

        /// <summary>
        /// The number of times this job can be retried
        /// </summary>
        public int? Limit { get; set; }

        /// <summary>
        /// The exit signal, if any, that may be retried
        /// </summary>
        public string? Signal { get; set; }

        /// <summary>
        /// The exit signal reason, if any, that may be retried
        /// </summary>
        public SignalReason? SignalReason { get; set; }
    }

    public class ManualRetry
    {
        /// <summary>
        /// Whether or not this job can be retried manually
        /// </summary>
        public bool? Allowed { get; set; }

        /// <summary>
        /// Whether or not this job can be retried after it has passed
        /// </summary>
        public bool? PermitOnPassed { get; set; }

        /// <summary>
        /// A string that will be displayed in a tooltip on the Retry button in Buildkite. This will
        /// only be displayed if the allowed attribute is set to false.
        /// </summary>
        public string? Reason { get; set; }
    }

    /// <summary>
    /// The signature of the command step, generally injected by agents at pipeline upload
    /// </summary>
    public class Signature
    {
        /// <summary>
        /// The algorithm used to generate the signature
        /// </summary>
        public string? Algorithm { get; set; }

        /// <summary>
        /// The fields that were signed to form the signature value
        /// </summary>
        public List<string>? SignedFields { get; set; }

        /// <summary>
        /// The signature value, a JWS compact signature with a detached body
        /// </summary>
        public string? Value { get; set; }
    }

    /// <summary>
    /// Properties of the build that will be created when the step is triggered
    /// </summary>
    public class Build
    {
        /// <summary>
        /// The branch for the build
        /// </summary>
        public string? Branch { get; set; }

        /// <summary>
        /// The commit hash for the build
        /// </summary>
        public string? Commit { get; set; }

        public Dictionary<string, object>? Env { get; set; }

        /// <summary>
        /// The message for the build (supports emoji)
        /// </summary>
        public string? Message { get; set; }

        /// <summary>
        /// Meta-data for the build
        /// </summary>
        public Dictionary<string, object>? MetaData { get; set; }
    }
}
