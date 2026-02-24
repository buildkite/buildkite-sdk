using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class PipelineTests
{
    [Fact]
    public void Pipeline_WithCommandStep_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = ":hammer: Build",
            Command = "dotnet build"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":hammer: Build", yaml);
        Assert.Contains("command: dotnet build", yaml);
    }

    [Fact]
    public void Pipeline_WithCommandStep_GeneratesCorrectJson()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = ":hammer: Build",
            Command = "dotnet build"
        });

        var json = pipeline.ToJson();

        Assert.Contains(":hammer: Build", json);
        Assert.Contains("dotnet build", json);
    }

    [Fact]
    public void Pipeline_WithMultipleSteps_GeneratesAllSteps()
    {
        var pipeline = new Pipeline();
        pipeline
            .AddStep(new CommandStep { Label = ":test_tube: Test", Command = "dotnet test" })
            .AddStep(new WaitStep())
            .AddStep(new CommandStep { Label = ":rocket: Deploy", Command = "./deploy.sh" });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":test_tube: Test", yaml);
        Assert.Contains(":rocket: Deploy", yaml);
        Assert.Contains("wait", yaml);
    }

    [Fact]
    public void Pipeline_WithAgents_IncludesAgentConfiguration()
    {
        var pipeline = new Pipeline();
        pipeline.AddAgent("queue", "default");
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });

        var yaml = pipeline.ToYaml();

        Assert.Contains("queue: default", yaml);
    }

    [Fact]
    public void Pipeline_WithEnvVars_IncludesEnvironmentVariables()
    {
        var pipeline = new Pipeline();
        pipeline.AddEnvironmentVariable("MY_VAR", "my_value");
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });

        var yaml = pipeline.ToYaml();

        Assert.Contains("MY_VAR: my_value", yaml);
    }

    [Fact]
    public void Pipeline_WithMultipleEnvVars_IncludesAllEnvironmentVariables()
    {
        var pipeline = new Pipeline();
        pipeline.AddEnvironmentVariable("MY_VAR", "my_value");
        pipeline.AddEnvironmentVariable("OTHER_VAR", "other_value");
        pipeline.AddEnvironmentVariable("THIRD_VAR", "third_value");
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });

        var yaml = pipeline.ToYaml();

        Assert.Contains("MY_VAR: my_value", yaml);
        Assert.Contains("OTHER_VAR: other_value", yaml);
        Assert.Contains("THIRD_VAR: third_value", yaml);
    }

    [Fact]
    public void Pipeline_SetPipeline_WithEnvVars_IncludesAllEnvironmentVariables()
    {
        var pipeline = new Pipeline();
        pipeline.SetPipeline(new BuildkitePipeline
        {
            Env = new Dictionary<string, string>
            {
                ["NODE_ENV"] = "production",
                ["CI"] = "true"
            },
            Steps = new List<IStep>
            {
                new CommandStep { Label = "Build", Command = "make" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("NODE_ENV: production", yaml);
        Assert.Contains("CI: true", yaml);
    }

    [Fact]
    public void Pipeline_Empty_GeneratesEmptyOutput()
    {
        var pipeline = new Pipeline();

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Equal("{}\n", yaml);
        Assert.Equal("{}", json);
    }
}
