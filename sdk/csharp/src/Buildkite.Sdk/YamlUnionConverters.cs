using YamlDotNet.Core;
using YamlDotNet.Core.Events;
using YamlDotNet.Serialization;

namespace Buildkite.Sdk;

/// <summary>
/// Generic YamlDotNet type converter that handles all union types.
/// Note: YAML deserialization is not supported - this SDK only produces YAML output.
/// </summary>
public sealed class UnionYamlTypeConverter : IYamlTypeConverter
{
    public bool Accepts(Type type)
    {
        if (!type.IsGenericType) return false;

        var genericDef = type.GetGenericTypeDefinition();
        return genericDef == typeof(OneOrMany<>) ||
               genericDef == typeof(BoolOr<>) ||
               genericDef == typeof(StringOr<>) ||
               genericDef == typeof(OneOf<,>) ||
               genericDef == typeof(OneOf<,,>) ||
               genericDef == typeof(Nullable<>);
    }

    public object? ReadYaml(IParser parser, Type type, ObjectDeserializer rootDeserializer)
    {
        throw new NotSupportedException(
            "Reading union types from YAML is not supported. This SDK only produces YAML output.");
    }

    public void WriteYaml(IEmitter emitter, object? value, Type type, ObjectSerializer serializer)
    {
        if (value == null)
        {
            emitter.Emit(new Scalar(null, "null"));
            return;
        }

        var actualType = value.GetType();
        if (!actualType.IsGenericType)
        {
            serializer(value, actualType);
            return;
        }

        var genericDef = actualType.GetGenericTypeDefinition();

        if (genericDef == typeof(OneOrMany<>))
        {
            WriteUnionValue(value, actualType, serializer);
        }
        else if (genericDef == typeof(BoolOr<>))
        {
            WriteBoolOr(emitter, value, actualType, serializer);
        }
        else if (genericDef == typeof(StringOr<>))
        {
            WriteStringOr(emitter, value, actualType, serializer);
        }
        else if (genericDef == typeof(OneOf<,>) || genericDef == typeof(OneOf<,,>))
        {
            WriteUnionValue(value, actualType, serializer);
        }
        else
        {
            serializer(value, actualType);
        }
    }

    private static void WriteUnionValue(object value, Type type, ObjectSerializer serializer)
    {
        var valueProp = type.GetProperty("Value")!;
        var innerValue = valueProp.GetValue(value);

        if (innerValue != null)
        {
            serializer(innerValue, innerValue.GetType());
        }
    }

    private static void WriteBoolOr(IEmitter emitter, object value, Type type, ObjectSerializer serializer)
    {
        var valueProp = type.GetProperty("Value")!;
        var isBoolProp = type.GetProperty("IsBool")!;

        var innerValue = valueProp.GetValue(value);
        var isBool = (bool)isBoolProp.GetValue(value)!;

        if (innerValue == null)
        {
            return;
        }

        if (isBool)
        {
            emitter.Emit(new Scalar(null, ((bool)innerValue) ? "true" : "false"));
        }
        else
        {
            serializer(innerValue, innerValue.GetType());
        }
    }

    private static void WriteStringOr(IEmitter emitter, object value, Type type, ObjectSerializer serializer)
    {
        var valueProp = type.GetProperty("Value")!;
        var isStringProp = type.GetProperty("IsString")!;

        var innerValue = valueProp.GetValue(value);
        var isString = (bool)isStringProp.GetValue(value)!;

        if (innerValue == null)
        {
            return;
        }

        if (isString)
        {
            emitter.Emit(new Scalar(null, (string)innerValue));
        }
        else
        {
            serializer(innerValue, innerValue.GetType());
        }
    }
}
