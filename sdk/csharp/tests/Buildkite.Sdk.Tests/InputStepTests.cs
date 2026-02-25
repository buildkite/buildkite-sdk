using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class InputStepTests
{
    [Fact]
    public void InputStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = ":memo: Provide Release Details",
            Prompt = "Please fill in the release information"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("input: ':memo: Provide Release Details'", yaml);
        Assert.Contains("prompt: Please fill in the release information", yaml);
    }

    [Fact]
    public void InputStep_WithKey_GeneratesKeyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Release Info",
            Key = "release-info"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("input: Release Info", yaml);
        Assert.Contains("key: release-info", yaml);
    }

    [Fact]
    public void InputStep_WithTextField_GeneratesFieldConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Release",
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

        Assert.Contains("text: Release Version", yaml);
        Assert.Contains("key: release-version", yaml);
        Assert.Contains("required: true", yaml);
    }

    [Fact]
    public void InputStep_WithSelectField_GeneratesFieldConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Choose Target",
            Fields = new List<Field>
            {
                new SelectField
                {
                    Select = "Region",
                    Key = "region",
                    Options = new List<SelectOption>
                    {
                        new SelectOption { Label = "US East", Value = "us-east-1" },
                        new SelectOption { Label = "EU West", Value = "eu-west-1" }
                    }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("select: Region", yaml);
        Assert.Contains("key: region", yaml);
        Assert.Contains("label: US East", yaml);
        Assert.Contains("value: us-east-1", yaml);
        Assert.Contains("label: EU West", yaml);
        Assert.Contains("value: eu-west-1", yaml);
    }

    [Fact]
    public void InputStep_WithBlockedState_GeneratesStateConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Gather Info",
            BlockedState = "running"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("blocked_state: running", yaml);
    }

    [Fact]
    public void InputStep_WithConditional_GeneratesIfConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Production Details",
            If = "build.branch == 'main'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("if: build.branch == 'main'", yaml);
    }

    [Fact]
    public void InputStep_WithDependsOn_GeneratesDependencyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Deploy Details",
            DependsOn = "build-step"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("depends_on: build-step", yaml);
    }
}
