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
                Env = new Dictionary<string, string> { ["DEPLOY_ENV"] = "production" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("DEPLOY_ENV: production", yaml);
    }

    [Fact]
    public void TriggerStep_WithBuildEnv_PlacesEnvUnderBuild()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Env = new Dictionary<string, string>
                {
                    ["DEPLOY_ENV"] = "production",
                    ["REGION"] = "us-east-1"
                }
            }
        });

        var yaml = pipeline.ToYaml();

        // env should be nested under build
        var lines = yaml.Split('\n');
        var buildLineIdx = Array.FindIndex(lines, l => l.TrimEnd() == "  build:");
        Assert.True(buildLineIdx >= 0, "Expected 'build:' under step");
        var envLineIdx = Array.FindIndex(lines, buildLineIdx, l => l.TrimEnd() == "    env:");
        Assert.True(envLineIdx > buildLineIdx, "Expected 'env:' nested under 'build:'");
    }

    [Fact]
    public void TriggerStep_WithEmptyBuildEnv_SerializesAsEmptyMapping()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Env = new Dictionary<string, string>()
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("env: {}", yaml);
    }

    [Fact]
    public void TriggerStep_WithNullBuildEnv_OmitsEnvFromOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Env = null
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.DoesNotContain("env:", yaml);
    }

    [Fact]
    public void TriggerStep_WithMultipleBuildEnvVars_GeneratesAllEnvVars()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Build = new TriggerBuild
            {
                Branch = "main",
                Env = new Dictionary<string, string>
                {
                    ["DEPLOY_ENV"] = "production",
                    ["REGION"] = "us-east-1",
                    ["DRY_RUN"] = "false"
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("DEPLOY_ENV: production", yaml);
        Assert.Contains("REGION: us-east-1", yaml);
        Assert.Contains("DRY_RUN: false", yaml);
    }
}
