using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class TriggerStepTests
{
    [Fact]
    public void TriggerStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Label = ":rocket: Trigger Deploy"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("trigger: deploy-pipeline", yaml);
        Assert.Contains(":rocket: Trigger Deploy", yaml);
    }

    [Fact]
    public void TriggerStep_WithBuild_GeneratesBuildConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Message = "Deploying to production"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("branch: main", yaml);
        Assert.Contains("message: Deploying to production", yaml);
    }

    [Fact]
    public void TriggerStep_WithAsync_GeneratesAsyncConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "another-pipeline",
            Async = true
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("async: true", yaml);
    }

    [Fact]
    public void TriggerStep_WithBuildEnv_GeneratesEnvConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Env = new Dictionary<string, object?> { ["DEPLOY_ENV"] = "production" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("DEPLOY_ENV: production", yaml);
    }
}
