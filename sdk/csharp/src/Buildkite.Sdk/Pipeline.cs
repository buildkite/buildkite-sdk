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
        WriteIndented = true,
        Converters = { new UnionConverterFactory() }
    };

    private static readonly ISerializer YamlSerializer = new SerializerBuilder()
        .WithNamingConvention(UnderscoredNamingConvention.Instance)
        .ConfigureDefaultValuesHandling(DefaultValuesHandling.OmitNull)
        .WithTypeConverter(new UnionYamlTypeConverter())
        .Build();

    public OneOf<AgentsObject, AgentsList>? Agents { get; set; }
    public Dictionary<string, object?> Env { get; set; } = new();
    public List<INotification> Notify { get; set; } = new();
    public List<IStep> Steps { get; set; } = new();
    public List<string> Secrets { get; set; } = new();
    public string? Image { get; set; }

    public Pipeline AddStep(IStep step)
    {
        Steps.Add(step);
        return this;
    }

    public Pipeline AddAgent(string key, string value)
    {
        if (Agents?.AsT1 is AgentsObject agentsObj)
        {
            agentsObj[key] = value;
        }
        else if (Agents?.AsT2 is not null)
        {
            throw new InvalidOperationException(
                "Cannot use AddAgent when Agents is already set as a list. Use SetAgents to replace.");
        }
        else
        {
            Agents = new AgentsObject { [key] = value };
        }
        return this;
    }

    public Pipeline SetAgents(AgentsObject agents)
    {
        Agents = agents;
        return this;
    }

    public Pipeline SetAgents(AgentsList agents)
    {
        Agents = agents;
        return this;
    }

    public Pipeline SetImage(string image)
    {
        Image = image;
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
        if (pipeline.Image != null)
            Image = pipeline.Image;
        return this;
    }

    private bool HasAgents()
    {
        if (Agents == null) return false;
        if (Agents.Value.AsT1 is AgentsObject obj) return obj.Count > 0;
        if (Agents.Value.AsT2 is AgentsList list) return list.Count > 0;
        return false;
    }

    private BuildkitePipeline Build()
    {
        var pipeline = new BuildkitePipeline();

        if (HasAgents())
            pipeline.Agents = Agents;
        if (Env.Count > 0)
            pipeline.Env = Env;
        if (Notify.Count > 0)
            pipeline.Notify = Notify;
        if (Steps.Count > 0)
            pipeline.Steps = Steps;
        if (Secrets.Count > 0)
            pipeline.Secrets = Secrets;
        if (Image != null)
            pipeline.Image = Image;

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
