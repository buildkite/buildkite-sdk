using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class CommandStepTests
{
    [Fact]
    public void CommandStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = ":dotnet: Build",
            Key = "build",
            Command = "dotnet build --configuration Release"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":dotnet: Build", yaml);
        Assert.Contains("key: build", yaml);
        Assert.Contains("command: dotnet build --configuration Release", yaml);
    }

    [Fact]
    public void CommandStep_WithAgents_GeneratesAgentConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Agents = new AgentsObject { ["queue"] = "linux", ["os"] = "ubuntu" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("queue: linux", yaml);
        Assert.Contains("os: ubuntu", yaml);
    }

    [Fact]
    public void CommandStep_WithEnv_GeneratesEnvConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, object?> { ["NODE_ENV"] = "production" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("NODE_ENV: production", yaml);
    }

    [Fact]
    public void CommandStep_WithParallelism_GeneratesParallelConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "dotnet test",
            Parallelism = 5
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("parallelism: 5", yaml);
    }

    [Fact]
    public void CommandStep_WithTimeout_GeneratesTimeoutConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Long running task",
            Command = "./long-task.sh",
            TimeoutInMinutes = 60
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("timeout_in_minutes: 60", yaml);
    }

    [Fact]
    public void CommandStep_WithDependsOn_GeneratesDependencyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Key = "build",
            Label = "Build",
            Command = "make build"
        });
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = "build"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("depends_on: build", yaml);
    }

    [Fact]
    public void CommandStep_WithConditional_GeneratesIfConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            If = "build.branch == 'main'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("if: build.branch == 'main'", yaml);
    }
}
