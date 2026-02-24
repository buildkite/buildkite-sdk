using System.Text.Json;
using System.Text.Json.Serialization;
using YamlDotNet.Core;
using YamlDotNet.Core.Events;
using YamlDotNet.Serialization;

namespace Buildkite.Sdk.Schema;

/// <summary>
/// An item in a depends_on list: either a step key string or a Dependency object.
/// </summary>
[JsonConverter(typeof(DependsOnItemJsonConverter))]
public class DependsOnItem
{
    private readonly object _value;

    private DependsOnItem(object value) => _value = value;

    public static DependsOnItem FromString(string key) => new(key);
    public static DependsOnItem FromDependency(Dependency dep) => new(dep);

    public static implicit operator DependsOnItem(string key) => FromString(key);
    public static implicit operator DependsOnItem(Dependency dep) => FromDependency(dep);

    public object Value => _value;
}

/// <summary>
/// Represents the depends_on configuration for a step.
/// Can be a single string, list of strings, list of dependency objects, or a mixed list.
/// </summary>
[JsonConverter(typeof(DependsOnJsonConverter))]
public class DependsOn
{
    private readonly object _value;

    private DependsOn(object value) => _value = value;

    public static DependsOn FromString(string key) => new(key);
    public static DependsOn FromStrings(params string[] keys) =>
        new(keys.Select(k => DependsOnItem.FromString(k)).ToList());
    public static DependsOn FromDependencies(params Dependency[] deps) =>
        new(deps.Select(d => DependsOnItem.FromDependency(d)).ToList());
    public static DependsOn FromItems(params DependsOnItem[] items) => new(items.ToList());

    public static implicit operator DependsOn(string key) => FromString(key);
    public static implicit operator DependsOn(string[] keys) => FromStrings(keys);
    public static implicit operator DependsOn(Dependency[] deps) => FromDependencies(deps);
    public static implicit operator DependsOn(DependsOnItem[] items) => FromItems(items);

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
/// Soft fail configuration. Can be true/false, "true"/"false", or a list of exit status conditions.
/// </summary>
[JsonConverter(typeof(SoftFailJsonConverter))]
public class SoftFail
{
    private readonly object _value;

    private SoftFail(object value) => _value = value;

    public static SoftFail FromBool(bool value) => new(value);
    public static SoftFail FromString(string value) => new(value);
    public static SoftFail FromConditions(params SoftFailCondition[] conditions) => new(conditions.ToList());

    public static implicit operator SoftFail(bool value) => FromBool(value);

    public object Value => _value;
}

/// <summary>
/// A soft fail condition based on exit status.
/// </summary>
public class SoftFailCondition
{
    /// <summary>The exit status that should be treated as a soft fail. Can be "*" or an integer.</summary>
    [JsonConverter(typeof(SoftFailExitStatusJsonConverter))]
    public SoftFailExitStatus? ExitStatus { get; set; }
}

/// <summary>
/// Represents an exit status for soft fail conditions: either "*" (any) or a specific integer.
/// </summary>
[JsonConverter(typeof(SoftFailExitStatusJsonConverter))]
public class SoftFailExitStatus
{
    private readonly object _value;

    private SoftFailExitStatus(object value) => _value = value;

    public static SoftFailExitStatus FromWildcard() => new("*");
    public static SoftFailExitStatus FromInt(int status) => new(status);

    public static implicit operator SoftFailExitStatus(int status) => FromInt(status);

    public object Value => _value;
}

/// <summary>
/// Skip configuration. Can be a boolean or a reason string.
/// </summary>
[JsonConverter(typeof(SkipJsonConverter))]
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

#region JSON Converters

internal class DependsOnItemJsonConverter : JsonConverter<DependsOnItem>
{
    public override DependsOnItem? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        => throw new NotSupportedException("Deserialization is not supported.");

    public override void Write(Utf8JsonWriter writer, DependsOnItem value, JsonSerializerOptions options)
    {
        switch (value.Value)
        {
            case string s:
                writer.WriteStringValue(s);
                break;
            case Dependency dep:
                JsonSerializer.Serialize(writer, dep, options);
                break;
        }
    }
}

internal class DependsOnJsonConverter : JsonConverter<DependsOn>
{
    public override DependsOn? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        => throw new NotSupportedException("Deserialization is not supported.");

    public override void Write(Utf8JsonWriter writer, DependsOn value, JsonSerializerOptions options)
    {
        switch (value.Value)
        {
            case string s:
                writer.WriteStringValue(s);
                break;
            case List<DependsOnItem> list:
                JsonSerializer.Serialize(writer, list, options);
                break;
        }
    }
}

internal class SoftFailJsonConverter : JsonConverter<SoftFail>
{
    public override SoftFail? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        => throw new NotSupportedException("Deserialization is not supported.");

    public override void Write(Utf8JsonWriter writer, SoftFail value, JsonSerializerOptions options)
    {
        switch (value.Value)
        {
            case bool b:
                writer.WriteBooleanValue(b);
                break;
            case string s:
                writer.WriteStringValue(s);
                break;
            case List<SoftFailCondition> conditions:
                JsonSerializer.Serialize(writer, conditions, options);
                break;
        }
    }
}

internal class SoftFailExitStatusJsonConverter : JsonConverter<SoftFailExitStatus>
{
    public override SoftFailExitStatus? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        => throw new NotSupportedException("Deserialization is not supported.");

    public override void Write(Utf8JsonWriter writer, SoftFailExitStatus value, JsonSerializerOptions options)
    {
        switch (value.Value)
        {
            case string s:
                writer.WriteStringValue(s);
                break;
            case int i:
                writer.WriteNumberValue(i);
                break;
        }
    }
}

internal class SkipJsonConverter : JsonConverter<Skip>
{
    public override Skip? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        => throw new NotSupportedException("Deserialization is not supported.");

    public override void Write(Utf8JsonWriter writer, Skip value, JsonSerializerOptions options)
    {
        switch (value.Value)
        {
            case bool b:
                writer.WriteBooleanValue(b);
                break;
            case string s:
                writer.WriteStringValue(s);
                break;
        }
    }
}

#endregion

#region YAML Type Converters

// YAML 1.1 boolean literals that must be quoted when emitting strings.
internal static class YamlQuoting
{
    private static readonly HashSet<string> YamlBooleans = new(StringComparer.OrdinalIgnoreCase)
    {
        "true", "false", "yes", "no", "on", "off", "y", "n"
    };

    public static Scalar SafeStringScalar(string value)
    {
        if (YamlBooleans.Contains(value))
            return new Scalar(null, null, value, ScalarStyle.DoubleQuoted, true, false);
        return new Scalar(value);
    }
}

internal class DependsOnItemYamlConverter : IYamlTypeConverter
{
    public bool Accepts(Type type) => type == typeof(DependsOnItem);

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
        => throw new NotSupportedException("Deserialization is not supported.");

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value is not DependsOnItem item) return;

        switch (item.Value)
        {
            case string s:
                emitter.Emit(new Scalar(s));
                break;
            case Dependency dep:
                serializer(dep, typeof(Dependency));
                break;
        }
    }
}

internal class DependsOnYamlConverter : IYamlTypeConverter
{
    public bool Accepts(Type type) => type == typeof(DependsOn);

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
        => throw new NotSupportedException("Deserialization is not supported.");

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value is not DependsOn dependsOn) return;

        switch (dependsOn.Value)
        {
            case string s:
                emitter.Emit(new Scalar(s));
                break;
            case List<DependsOnItem> list:
                serializer(list, typeof(List<DependsOnItem>));
                break;
        }
    }
}

internal class SoftFailYamlConverter : IYamlTypeConverter
{
    public bool Accepts(Type type) => type == typeof(SoftFail);

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
        => throw new NotSupportedException("Deserialization is not supported.");

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value is not SoftFail softFail) return;

        switch (softFail.Value)
        {
            case bool b:
                emitter.Emit(new Scalar(null, b ? "true" : "false"));
                break;
            case string s:
                emitter.Emit(new Scalar(null, null, s, ScalarStyle.DoubleQuoted, true, false));
                break;
            case List<SoftFailCondition> conditions:
                serializer(conditions, typeof(List<SoftFailCondition>));
                break;
        }
    }
}

internal class SoftFailExitStatusYamlConverter : IYamlTypeConverter
{
    public bool Accepts(Type type) => type == typeof(SoftFailExitStatus);

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
        => throw new NotSupportedException("Deserialization is not supported.");

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value is not SoftFailExitStatus status) return;

        switch (status.Value)
        {
            case string s:
                emitter.Emit(new Scalar(null, null, s, ScalarStyle.DoubleQuoted, true, false));
                break;
            case int i:
                emitter.Emit(new Scalar(i.ToString()));
                break;
        }
    }
}

internal class SkipYamlConverter : IYamlTypeConverter
{
    public bool Accepts(Type type) => type == typeof(Skip);

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
        => throw new NotSupportedException("Deserialization is not supported.");

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value is not Skip skip) return;

        switch (skip.Value)
        {
            case bool b:
                emitter.Emit(new Scalar(null, b ? "true" : "false"));
                break;
            case string s:
                emitter.Emit(YamlQuoting.SafeStringScalar(s));
                break;
        }
    }
}

#endregion
