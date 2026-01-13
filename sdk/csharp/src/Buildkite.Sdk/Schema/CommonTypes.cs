using System.Text.Json.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// Represents the depends_on configuration for a step.
/// Can be a single string, list of strings, or list of dependency objects.
/// </summary>
public class DependsOn
{
    private readonly object _value;

    private DependsOn(object value) => _value = value;

    public static DependsOn FromString(string key) => new(key);
    public static DependsOn FromStrings(params string[] keys) => new(keys.ToList());
    public static DependsOn FromDependencies(params Dependency[] dependencies) => new(dependencies.ToList());

    public static implicit operator DependsOn(string key) => FromString(key);
    public static implicit operator DependsOn(string[] keys) => FromStrings(keys);

    public object Value => _value;
}

/// <summary>
/// A dependency specification with optional allow_failure setting.
/// </summary>
public class Dependency
{
    /// <summary>The key of the step to depend on.</summary>
    public string? Step { get; set; }

    /// <summary>Whether to allow this dependency to fail.</summary>
    public bool? AllowFailure { get; set; }
}

/// <summary>
/// Retry configuration for command steps.
/// </summary>
public class Retry
{
    /// <summary>Automatic retry configuration.</summary>
    public AutomaticRetry? Automatic { get; set; }

    /// <summary>Manual retry configuration.</summary>
    public ManualRetry? Manual { get; set; }
}

/// <summary>
/// Automatic retry configuration.
/// </summary>
public class AutomaticRetry
{
    /// <summary>The exit status that triggers a retry. Use "*" for any exit status.</summary>
    public object? ExitStatus { get; set; }

    /// <summary>The maximum number of automatic retries.</summary>
    public int? Limit { get; set; }

    /// <summary>The exit signal that triggers a retry.</summary>
    public string? Signal { get; set; }

    /// <summary>The exit signal reason that triggers a retry.</summary>
    public string? SignalReason { get; set; }
}

/// <summary>
/// Manual retry configuration.
/// </summary>
public class ManualRetry
{
    /// <summary>Whether manual retries are allowed.</summary>
    public bool? Allowed { get; set; }

    /// <summary>Whether manual retry requires a permit.</summary>
    public bool? PermitOnPassed { get; set; }

    /// <summary>The reason for the manual retry being restricted.</summary>
    public string? Reason { get; set; }
}

/// <summary>
/// Soft fail configuration. Can be true/false or a list of exit statuses.
/// </summary>
public class SoftFail
{
    private readonly object _value;

    private SoftFail(object value) => _value = value;

    public static SoftFail FromBool(bool value) => new(value);
    public static SoftFail FromExitStatuses(params SoftFailCondition[] conditions) => new(conditions.ToList());

    public static implicit operator SoftFail(bool value) => FromBool(value);

    public object Value => _value;
}

/// <summary>
/// A soft fail condition based on exit status.
/// </summary>
public class SoftFailCondition
{
    /// <summary>The exit status that should be treated as a soft fail.</summary>
    public object? ExitStatus { get; set; }
}

/// <summary>
/// Skip configuration. Can be a boolean or a reason string.
/// </summary>
public class Skip
{
    private readonly object _value;

    private Skip(object value) => _value = value;

    public static Skip FromBool(bool value) => new(value);
    public static Skip FromReason(string reason) => new(reason);

    public static implicit operator Skip(bool value) => FromBool(value);
    public static implicit operator Skip(string reason) => FromReason(reason);

    public object Value => _value;
}

/// <summary>
/// Cache configuration for a step.
/// </summary>
public class Cache
{
    /// <summary>The name of the cache.</summary>
    public string? Name { get; set; }

    /// <summary>The paths to cache.</summary>
    public List<string> Paths { get; set; } = new();

    /// <summary>The maximum size of the cache.</summary>
    public string? Size { get; set; }
}

/// <summary>
/// Matrix configuration for generating multiple jobs from a single step.
/// </summary>
public class Matrix
{
    /// <summary>The setup for the matrix.</summary>
    public Dictionary<string, List<string>> Setup { get; set; } = new();

    /// <summary>Adjustments to the matrix.</summary>
    public List<MatrixAdjustment>? Adjustments { get; set; }
}

/// <summary>
/// An adjustment to a matrix configuration.
/// </summary>
public class MatrixAdjustment
{
    /// <summary>The values to match for this adjustment.</summary>
    public Dictionary<string, string>? With { get; set; }

    /// <summary>Whether to skip this combination.</summary>
    public Skip? Skip { get; set; }

    /// <summary>Soft fail configuration for this combination.</summary>
    public SoftFail? SoftFail { get; set; }
}

/// <summary>
/// Signature configuration for signed pipelines.
/// </summary>
public class Signature
{
    /// <summary>The algorithm used for signing.</summary>
    public string? Algorithm { get; set; }

    /// <summary>The fields that were signed.</summary>
    public List<string>? SignedFields { get; set; }

    /// <summary>The signature value.</summary>
    public string? Value { get; set; }
}


