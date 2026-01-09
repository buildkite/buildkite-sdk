using System.Text.Json;
using System.Text.Json.Serialization;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;
using Buildkite.Sdk.Schema;

namespace Buildkite.Sdk;

internal sealed class SnakeCaseNamingPolicy : JsonNamingPolicy
{
    public override string ConvertName(string name)
    {
        if (string.IsNullOrEmpty(name))
            return name;

        var builder = new System.Text.StringBuilder();
        for (int i = 0; i < name.Length; i++)
        {
            char c = name[i];
            if (char.IsUpper(c))
            {
                if (i > 0)
                    builder.Append('_');
                builder.Append(char.ToLowerInvariant(c));
            }
            else
            {
                builder.Append(c);
            }
        }
        return builder.ToString();
    }
}

public class Pipeline
{
    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        PropertyNamingPolicy = new SnakeCaseNamingPolicy(),
        DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
        WriteIndented = true
    };

    private static readonly ISerializer YamlSerializer = new SerializerBuilder()
        .WithNamingConvention(UnderscoredNamingConvention.Instance)
        .ConfigureDefaultValuesHandling(DefaultValuesHandling.OmitNull)
        .Build();

    public AgentsObject Agents { get; set; } = new();
    public Dictionary<string, object?> Env { get; set; } = new();
    public List<INotification> Notify { get; set; } = new();
    public List<IStep> Steps { get; set; } = new();
    public List<string> Secrets { get; set; } = new();

    public Pipeline AddStep(IStep step)
    {
        Steps.Add(step);
        return this;
    }

    public Pipeline AddAgent(string key, string value)
    {
        Agents[key] = value;
        return this;
    }

    public Pipeline AddEnvironmentVariable(string key, object? value)
    {
        Env[key] = value;
        return this;
    }

    public Pipeline AddNotify(INotification notification)
    {
        Notify.Add(notification);
        return this;
    }

    public Pipeline SetSecrets(List<string> secrets)
    {
        Secrets = secrets;
        return this;
    }

    public Pipeline SetPipeline(BuildkitePipeline pipeline)
    {
        if (pipeline.Agents != null)
            Agents = pipeline.Agents;
        if (pipeline.Env != null)
            Env = pipeline.Env;
        if (pipeline.Notify != null)
            Notify = pipeline.Notify;
        if (pipeline.Steps != null)
            Steps = pipeline.Steps;
        if (pipeline.Secrets != null)
            Secrets = pipeline.Secrets;
        return this;
    }

    private BuildkitePipeline Build()
    {
        var pipeline = new BuildkitePipeline();

        if (Agents.Count > 0)
            pipeline.Agents = Agents;
        if (Env.Count > 0)
            pipeline.Env = Env;
        if (Notify.Count > 0)
            pipeline.Notify = Notify;
        if (Steps.Count > 0)
            pipeline.Steps = Steps;
        if (Secrets.Count > 0)
            pipeline.Secrets = Secrets;

        return pipeline;
    }

    public string ToJson()
    {
        return JsonSerializer.Serialize(Build(), JsonOptions);
    }

    public string ToYaml()
    {
        return YamlSerializer.Serialize(Build());
    }
}
