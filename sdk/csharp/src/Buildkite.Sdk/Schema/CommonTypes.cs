namespace Buildkite.Sdk.Schema;

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
/// A soft fail condition based on exit status.
/// </summary>
public class SoftFailCondition
{
    /// <summary>The exit status that should be treated as a soft fail.</summary>
    public object? ExitStatus { get; set; }
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
    public StringOr<bool>? Skip { get; set; }

    /// <summary>Soft fail configuration for this combination.</summary>
    public BoolOr<List<SoftFailCondition>>? SoftFail { get; set; }
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

/// <summary>
/// Plugin configuration.
/// </summary>
public class PluginConfig
{
    /// <summary>The plugin name (e.g., "docker#v5.0.0").</summary>
    public string Name { get; }

    /// <summary>The plugin configuration.</summary>
    public object? Config { get; }

    public PluginConfig(string name, object? config = null)
    {
        Name = name;
        Config = config;
    }
}
