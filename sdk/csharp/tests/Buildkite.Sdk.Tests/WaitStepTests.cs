using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class WaitStepTests
{
    [Fact]
    public void WaitStep_Basic_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep());

        var yaml = pipeline.ToYaml();

        Assert.Contains("wait", yaml);
    }

    [Fact]
    public void WaitStep_WithLabel_GeneratesLabel()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep
        {
            Wait = "Wait for all tests"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("wait: Wait for all tests", yaml);
    }

    [Fact]
    public void WaitStep_WithContinueOnFailure_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep
        {
            ContinueOnFailure = true
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("continue_on_failure: true", yaml);
    }

    [Fact]
    public void WaitStep_WithConditional_GeneratesIfCondition()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep
        {
            If = "build.branch == 'main'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("if: build.branch == 'main'", yaml);
    }

    [Fact]
    public void WaitStep_WithDependsOn_GeneratesDependency()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep
        {
            DependsOn = "build-step"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("depends_on: build-step", yaml);
    }
}
