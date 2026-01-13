namespace Buildkite.Sdk;

/// <summary>
/// Represents a value that can be either a single item of type T or a list of T.
/// Commonly used for fields like commands, branches, artifact_paths, etc.
/// A default instance represents "unset" (no value).
/// </summary>
public readonly struct OneOrMany<T>
{
    private readonly object? _value;

    private OneOrMany(object? value) => _value = value;

    public static OneOrMany<T> FromOne(T value) => new(value);
    public static OneOrMany<T> FromMany(List<T> values) => new(values);

    public static implicit operator OneOrMany<T>(T value) => FromOne(value);
    public static implicit operator OneOrMany<T>(T[] values) => FromMany(values.ToList());
    public static implicit operator OneOrMany<T>(List<T> values) => FromMany(values);

    /// <summary>Whether this union has a value (is not default/unset).</summary>
    public bool HasValue => _value is not null;
    public bool IsSingle => _value is T;
    public bool IsList => _value is List<T>;
    /// <summary>Returns the single value, or default(T) if not holding a single value.</summary>
    public T? Single => _value is T t ? t : default;
    /// <summary>Returns the list value, or null if not holding a list.</summary>
    public List<T>? List => _value as List<T>;
    public object? Value => _value;
}

/// <summary>
/// Represents a value that can be either a boolean or a value of type T.
/// Commonly used for fields like soft_fail, skip, etc.
/// A default instance represents "unset" (no value).
/// </summary>
public readonly struct BoolOr<T>
{
    private readonly object? _value;

    private BoolOr(object? value) => _value = value;

    public static BoolOr<T> FromBool(bool value) => new(value);
    public static BoolOr<T> FromValue(T value) => new(value);

    public static implicit operator BoolOr<T>(bool value) => FromBool(value);
    public static implicit operator BoolOr<T>(T value) => FromValue(value);

    /// <summary>Whether this union has a value (is not default/unset).</summary>
    public bool HasValue => _value is not null;
    public bool IsBool => _value is bool;
    public bool IsValue => _value is T;
    public bool? Bool => _value is bool b ? b : null;
    public T? TypedValue => _value is T t ? t : default;
    public object? Value => _value;
}

/// <summary>
/// Represents a value that can be either a string or a value of type T.
/// Commonly used for fields like skip (bool/string) where we need string-specific handling.
/// A default instance represents "unset" (no value).
/// </summary>
public readonly struct StringOr<T>
{
    private readonly object? _value;

    private StringOr(object? value) => _value = value;

    public static StringOr<T> FromString(string value) => new(value);
    public static StringOr<T> FromValue(T value) => new(value);

    public static implicit operator StringOr<T>(string value) => FromString(value);
    public static implicit operator StringOr<T>(T value) => FromValue(value);

    /// <summary>Whether this union has a value (is not default/unset).</summary>
    public bool HasValue => _value is not null;
    public bool IsString => _value is string;
    public bool IsValue => _value is T;
    public string? String => _value as string;
    public T? TypedValue => _value is T t ? t : default;
    public object? Value => _value;
}

/// <summary>
/// Represents a value that can be one of two types.
/// A default instance represents "unset" (no value).
/// </summary>
public readonly struct OneOf<T1, T2>
{
    private readonly object? _value;
    private readonly int _index;

    private OneOf(object? value, int index)
    {
        _value = value;
        _index = index;
    }

    public static OneOf<T1, T2> FromT1(T1 value) => new(value, 1);
    public static OneOf<T1, T2> FromT2(T2 value) => new(value, 2);

    public static implicit operator OneOf<T1, T2>(T1 value) => FromT1(value);
    public static implicit operator OneOf<T1, T2>(T2 value) => FromT2(value);

    /// <summary>Whether this union has a value (is not default/unset).</summary>
    public bool HasValue => _index != 0;
    public bool IsT1 => _index == 1;
    public bool IsT2 => _index == 2;
    public T1? AsT1 => _index == 1 && _value is T1 t ? t : default;
    public T2? AsT2 => _index == 2 && _value is T2 t ? t : default;
    public object? Value => _value;
}

/// <summary>
/// Represents a value that can be one of three types.
/// A default instance represents "unset" (no value).
/// </summary>
public readonly struct OneOf<T1, T2, T3>
{
    private readonly object? _value;
    private readonly int _index;

    private OneOf(object? value, int index)
    {
        _value = value;
        _index = index;
    }

    public static OneOf<T1, T2, T3> FromT1(T1 value) => new(value, 1);
    public static OneOf<T1, T2, T3> FromT2(T2 value) => new(value, 2);
    public static OneOf<T1, T2, T3> FromT3(T3 value) => new(value, 3);

    public static implicit operator OneOf<T1, T2, T3>(T1 value) => FromT1(value);
    public static implicit operator OneOf<T1, T2, T3>(T2 value) => FromT2(value);
    public static implicit operator OneOf<T1, T2, T3>(T3 value) => FromT3(value);

    /// <summary>Whether this union has a value (is not default/unset).</summary>
    public bool HasValue => _index != 0;
    public bool IsT1 => _index == 1;
    public bool IsT2 => _index == 2;
    public bool IsT3 => _index == 3;
    public T1? AsT1 => _index == 1 && _value is T1 t ? t : default;
    public T2? AsT2 => _index == 2 && _value is T2 t ? t : default;
    public T3? AsT3 => _index == 3 && _value is T3 t ? t : default;
    public object? Value => _value;
}
