using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class BlockStepTests
{
    [Fact]
    public void BlockStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = ":rocket: Deploy to Production?",
            Prompt = "Are you sure you want to deploy?"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":rocket: Deploy to Production?", yaml);
        Assert.Contains("prompt: Are you sure you want to deploy?", yaml);
    }

    [Fact]
    public void BlockStep_WithKey_GeneratesKeyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Confirm",
            Key = "confirm-deploy"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("key: confirm-deploy", yaml);
    }

    [Fact]
    public void BlockStep_WithTextField_GeneratesFieldConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Release",
            Fields = new List<Field>
            {
                new TextField
                {
                    Text = "Release Version",
                    Key = "release-version",
                    Required = true
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("release-version", yaml);
    }

    [Fact]
    public void BlockStep_WithSelectField_GeneratesFieldConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Choose Environment",
            Fields = new List<Field>
            {
                new SelectField
                {
                    Select = "Environment",
                    Key = "environment",
                    Options = new List<SelectOption>
                    {
                        new SelectOption { Label = "Staging", Value = "staging" },
                        new SelectOption { Label = "Production", Value = "production" }
                    }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("environment", yaml);
        Assert.Contains("staging", yaml);
        Assert.Contains("production", yaml);
    }

    [Fact]
    public void BlockStep_WithFields_SerializesDerivedProperties()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Release",
            Fields = new List<Field>
            {
                new TextField
                {
                    Text = "Release Version",
                    Key = "release-version",
                    Required = true
                },
                new SelectField
                {
                    Select = "Environment",
                    Key = "environment",
                    Options = new List<SelectOption>
                    {
                        new SelectOption { Label = "Staging", Value = "staging" },
                        new SelectOption { Label = "Production", Value = "production" }
                    }
                }
            }
        });

        var json = pipeline.ToJson();

        Assert.Contains("\"text\":", json);
        Assert.Contains("\"select\":", json);
        Assert.Contains("\"options\":", json);
    }

    [Fact]
    public void BlockStep_WithBlockedState_GeneratesStateConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Manual Approval",
            BlockedState = "running"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("blocked_state: running", yaml);
    }
}
