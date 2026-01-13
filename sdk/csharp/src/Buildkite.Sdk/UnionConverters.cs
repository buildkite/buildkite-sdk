using System.Text.Json;
using System.Text.Json.Serialization;

namespace Buildkite.Sdk;

/// <summary>
/// JSON converter for OneOrMany&lt;T&gt; that serializes single values directly
/// and lists as JSON arrays.
/// </summary>
public sealed class OneOrManyConverter<T> : JsonConverter<OneOrMany<T>>
{
    public override OneOrMany<T> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        if (reader.TokenType == JsonTokenType.StartArray)
        {
            var list = JsonSerializer.Deserialize<List<T>>(ref reader, options);
            return list != null ? OneOrMany<T>.FromMany(list) : default;
        }

        var single = JsonSerializer.Deserialize<T>(ref reader, options);
        return single != null ? OneOrMany<T>.FromOne(single) : default;
    }

    public override void Write(Utf8JsonWriter writer, OneOrMany<T> value, JsonSerializerOptions options)
    {
        if (value.Value == null)
        {
            writer.WriteNullValue();
            return;
        }

        if (value.IsList)
        {
            JsonSerializer.Serialize(writer, value.List, options);
        }
        else
        {
            JsonSerializer.Serialize(writer, value.Single, options);
        }
    }
}

/// <summary>
/// JSON converter for BoolOr&lt;T&gt; that handles boolean or typed value serialization.
/// </summary>
public sealed class BoolOrConverter<T> : JsonConverter<BoolOr<T>>
{
    public override BoolOr<T> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        if (reader.TokenType == JsonTokenType.True || reader.TokenType == JsonTokenType.False)
        {
            return BoolOr<T>.FromBool(reader.GetBoolean());
        }

        var value = JsonSerializer.Deserialize<T>(ref reader, options);
        return value != null ? BoolOr<T>.FromValue(value) : default;
    }

    public override void Write(Utf8JsonWriter writer, BoolOr<T> value, JsonSerializerOptions options)
    {
        if (value.Value == null)
        {
            writer.WriteNullValue();
            return;
        }

        if (value.IsBool)
        {
            writer.WriteBooleanValue(value.Bool!.Value);
        }
        else
        {
            JsonSerializer.Serialize(writer, value.TypedValue, options);
        }
    }
}

/// <summary>
/// JSON converter for StringOr&lt;T&gt; that handles string or typed value serialization.
/// </summary>
public sealed class StringOrConverter<T> : JsonConverter<StringOr<T>>
{
    public override StringOr<T> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        if (reader.TokenType == JsonTokenType.String)
        {
            return StringOr<T>.FromString(reader.GetString()!);
        }

        var value = JsonSerializer.Deserialize<T>(ref reader, options);
        return value != null ? StringOr<T>.FromValue(value) : default;
    }

    public override void Write(Utf8JsonWriter writer, StringOr<T> value, JsonSerializerOptions options)
    {
        if (value.Value == null)
        {
            writer.WriteNullValue();
            return;
        }

        if (value.IsString)
        {
            writer.WriteStringValue(value.String);
        }
        else
        {
            JsonSerializer.Serialize(writer, value.TypedValue, options);
        }
    }
}

/// <summary>
/// JSON converter for OneOf&lt;T1, T2&gt;.
/// Attempts to deserialize as T1 first, then T2.
/// Throws JsonException if neither type can be deserialized.
/// </summary>
public sealed class OneOfConverter<T1, T2> : JsonConverter<OneOf<T1, T2>>
{
    public override OneOf<T1, T2> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        if (reader.TokenType == JsonTokenType.Null)
        {
            return default;
        }

        using var doc = JsonDocument.ParseValue(ref reader);
        var json = doc.RootElement.GetRawText();

        try
        {
            var t1 = JsonSerializer.Deserialize<T1>(json, options);
            if (t1 != null) return OneOf<T1, T2>.FromT1(t1);
        }
        catch { }

        try
        {
            var t2 = JsonSerializer.Deserialize<T2>(json, options);
            if (t2 != null) return OneOf<T1, T2>.FromT2(t2);
        }
        catch { }

        throw new JsonException(
            $"Could not deserialize value as {typeof(T1).Name} or {typeof(T2).Name}: {json}");
    }

    public override void Write(Utf8JsonWriter writer, OneOf<T1, T2> value, JsonSerializerOptions options)
    {
        if (value.Value == null)
        {
            writer.WriteNullValue();
            return;
        }

        JsonSerializer.Serialize(writer, value.Value, value.Value.GetType(), options);
    }
}

/// <summary>
/// JSON converter for OneOf&lt;T1, T2, T3&gt;.
/// Attempts to deserialize as T1 first, then T2, then T3.
/// Throws JsonException if no type can be deserialized.
/// </summary>
public sealed class OneOfConverter<T1, T2, T3> : JsonConverter<OneOf<T1, T2, T3>>
{
    public override OneOf<T1, T2, T3> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    {
        if (reader.TokenType == JsonTokenType.Null)
        {
            return default;
        }

        using var doc = JsonDocument.ParseValue(ref reader);
        var json = doc.RootElement.GetRawText();

        try
        {
            var t1 = JsonSerializer.Deserialize<T1>(json, options);
            if (t1 != null) return OneOf<T1, T2, T3>.FromT1(t1);
        }
        catch { }

        try
        {
            var t2 = JsonSerializer.Deserialize<T2>(json, options);
            if (t2 != null) return OneOf<T1, T2, T3>.FromT2(t2);
        }
        catch { }

        try
        {
            var t3 = JsonSerializer.Deserialize<T3>(json, options);
            if (t3 != null) return OneOf<T1, T2, T3>.FromT3(t3);
        }
        catch { }

        throw new JsonException(
            $"Could not deserialize value as {typeof(T1).Name}, {typeof(T2).Name}, or {typeof(T3).Name}: {json}");
    }

    public override void Write(Utf8JsonWriter writer, OneOf<T1, T2, T3> value, JsonSerializerOptions options)
    {
        if (value.Value == null)
        {
            writer.WriteNullValue();
            return;
        }

        JsonSerializer.Serialize(writer, value.Value, value.Value.GetType(), options);
    }
}

/// <summary>
/// Factory for creating union converters dynamically.
/// </summary>
public class UnionConverterFactory : JsonConverterFactory
{
    public override bool CanConvert(Type typeToConvert)
    {
        if (!typeToConvert.IsGenericType) return false;

        var genericDef = typeToConvert.GetGenericTypeDefinition();
        return genericDef == typeof(OneOrMany<>) ||
               genericDef == typeof(BoolOr<>) ||
               genericDef == typeof(StringOr<>) ||
               genericDef == typeof(OneOf<,>) ||
               genericDef == typeof(OneOf<,,>);
    }

    public override JsonConverter CreateConverter(Type typeToConvert, JsonSerializerOptions options)
    {
        var genericDef = typeToConvert.GetGenericTypeDefinition();
        var typeArgs = typeToConvert.GetGenericArguments();

        if (genericDef == typeof(OneOrMany<>))
        {
            var converterType = typeof(OneOrManyConverter<>).MakeGenericType(typeArgs);
            return (JsonConverter)Activator.CreateInstance(converterType)!;
        }

        if (genericDef == typeof(BoolOr<>))
        {
            var converterType = typeof(BoolOrConverter<>).MakeGenericType(typeArgs);
            return (JsonConverter)Activator.CreateInstance(converterType)!;
        }

        if (genericDef == typeof(StringOr<>))
        {
            var converterType = typeof(StringOrConverter<>).MakeGenericType(typeArgs);
            return (JsonConverter)Activator.CreateInstance(converterType)!;
        }

        if (genericDef == typeof(OneOf<,>))
        {
            var converterType = typeof(OneOfConverter<,>).MakeGenericType(typeArgs);
            return (JsonConverter)Activator.CreateInstance(converterType)!;
        }

        if (genericDef == typeof(OneOf<,,>))
        {
            var converterType = typeof(OneOfConverter<,,>).MakeGenericType(typeArgs);
            return (JsonConverter)Activator.CreateInstance(converterType)!;
        }

        throw new NotSupportedException($"Cannot create converter for {typeToConvert}");
    }
}
