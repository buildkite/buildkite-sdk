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
            Env = new Dictionary<string, string> { ["NODE_ENV"] = "production" }
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

    [Fact]
    public void CommandStep_WithStringPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[] { "docker#v5.0.0", "ecr#v2.0.0" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("ecr#v2.0.0", yaml);
    }

    [Fact]
    public void CommandStep_WithConfiguredPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[]
            {
                new Dictionary<string, object>
                {
                    ["docker#v5.0.0"] = new { image = "node:18" }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("image: node:18", yaml);
    }

    [Fact]
    public void CommandStep_WithMixedPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[]
            {
                "artifacts#v1.0.0",
                new Dictionary<string, object>
                {
                    ["docker#v5.0.0"] = new { image = "node:18", environment = new[] { "NODE_ENV=production" } }
                },
                new Dictionary<string, object>
                {
                    ["ecr#v2.0.0"] = new { login = true }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("artifacts#v1.0.0", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("image: node:18", yaml);
        Assert.Contains("ecr#v2.0.0", yaml);
        Assert.Contains("login: true", yaml);
    }

    [Fact]
    public void CommandStep_WithSecretsAsList_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            Secrets = new[] { "my-secret", "another-secret" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("secrets:", yaml);
        Assert.Contains("my-secret", yaml);
        Assert.Contains("another-secret", yaml);
    }

    [Fact]
    public void CommandStep_WithSecretsAsMap_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            Secrets = new Dictionary<string, string>
            {
                ["MY_SECRET"] = "org/secret-name",
                ["API_KEY"] = "org/api-key"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("secrets:", yaml);
        Assert.Contains("MY_SECRET: org/secret-name", yaml);
        Assert.Contains("API_KEY: org/api-key", yaml);
    }
}
