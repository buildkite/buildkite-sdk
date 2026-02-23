using System.Text.Json;
using System.Text.Json.Serialization;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;
using Buildkite.Sdk.Schema;

namespace Buildkite.Sdk;

public class Pipeline
{
    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower,
        DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
        WriteIndented = true
    };

    private static readonly ISerializer YamlSerializer = new SerializerBuilder()
        .WithNamingConvention(UnderscoredNamingConvention.Instance)
        .ConfigureDefaultValuesHandling(DefaultValuesHandling.OmitNull)
        .Build();

    public object? Agents { get; set; }
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
        if (Agents is AgentsObject agentsObj)
        {
            agentsObj[key] = value;
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
        return Agents switch
        {
            AgentsObject obj => obj.Count > 0,
            AgentsList list => list.Count > 0,
            _ => Agents != null
        };
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
