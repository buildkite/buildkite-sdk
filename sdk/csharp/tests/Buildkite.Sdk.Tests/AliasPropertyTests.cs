using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class AliasPropertyTests
{
    [Fact]
    public void CommandStep_NameDelegatesToLabel()
    {
        var step = new CommandStep { Name = "Build" };
        Assert.Equal("Build", step.Label);
        step.Label = "Test";
        Assert.Equal("Test", step.Name);
    }

    [Fact]
    public void CommandStep_IdDelegatesToKey()
    {
#pragma warning disable CS0618
        var step = new CommandStep { Id = "build" };
#pragma warning restore CS0618
        Assert.Equal("build", step.Key);
    }

    [Fact]
    public void CommandStep_IdentifierDelegatesToKey()
    {
        var step = new CommandStep { Identifier = "build" };
        Assert.Equal("build", step.Key);
        step.Key = "test";
        Assert.Equal("test", step.Identifier);
    }

    [Fact]
    public void CommandStep_CommandsDelegatesToCommand()
    {
        var cmds = new[] { "make build", "make test" };
        var step = new CommandStep { Commands = cmds };
        Assert.Same(cmds, step.Command);
        step.Command = "echo hello";
        Assert.Equal("echo hello", step.Commands);
    }

    [Fact]
    public void CommandStep_AliasesNotSerialized()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Name = "Build",
            Identifier = "build",
            Commands = "make",
        });

        var yaml = pipeline.ToYaml();
        Assert.Contains("label: Build", yaml);
        Assert.Contains("key: build", yaml);
        Assert.Contains("command: make", yaml);
        Assert.DoesNotContain("name:", yaml);
        Assert.DoesNotContain("identifier:", yaml);
        Assert.DoesNotContain("commands:", yaml);

        var json = pipeline.ToJson();
        Assert.Contains("\"label\"", json);
        Assert.Contains("\"key\"", json);
        Assert.Contains("\"command\"", json);
        Assert.DoesNotContain("\"name\"", json);
        Assert.DoesNotContain("\"identifier\"", json);
        Assert.DoesNotContain("\"commands\"", json);
    }

    [Fact]
    public void BlockStep_LabelAndNameDelegateToBlock()
    {
        var step = new BlockStep { Label = "Approve" };
        Assert.Equal("Approve", step.Block);
        step.Name = "Deploy Gate";
        Assert.Equal("Deploy Gate", step.Block);
        Assert.Equal("Deploy Gate", step.Label);
    }

    [Fact]
    public void BlockStep_AliasesNotSerialized()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Label = "Approve",
            Identifier = "approve",
        });

        var yaml = pipeline.ToYaml();
        Assert.Contains("block: Approve", yaml);
        Assert.Contains("key: approve", yaml);
        Assert.DoesNotContain("label:", yaml);
        Assert.DoesNotContain("name:", yaml);
        Assert.DoesNotContain("identifier:", yaml);
    }

    [Fact]
    public void WaitStep_LabelAndNameDelegateToWait()
    {
        var step = new WaitStep { Label = "Hold" };
        Assert.Equal("Hold", step.Wait);
        step.Name = "Pause";
        Assert.Equal("Pause", step.Wait);
    }

    [Fact]
    public void GroupStep_LabelAndNameDelegateToGroup()
    {
        var step = new GroupStep { Label = "Tests" };
        Assert.Equal("Tests", step.Group);
        step.Name = "Checks";
        Assert.Equal("Checks", step.Group);
    }

    [Fact]
    public void InputStep_LabelAndNameDelegateToInput()
    {
        var step = new InputStep { Label = "Config" };
        Assert.Equal("Config", step.Input);
        step.Name = "Settings";
        Assert.Equal("Settings", step.Input);
    }

    [Fact]
    public void TriggerStep_NameDelegatesToLabel()
    {
        var step = new TriggerStep { Name = "Deploy" };
        Assert.Equal("Deploy", step.Label);
        step.Label = "Ship";
        Assert.Equal("Ship", step.Name);
    }

    [Fact]
    public void TriggerStep_AliasesNotSerialized()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy-pipeline",
            Name = "Deploy",
            Identifier = "deploy",
        });

        var yaml = pipeline.ToYaml();
        Assert.Contains("label: Deploy", yaml);
        Assert.Contains("key: deploy", yaml);
        Assert.DoesNotContain("name:", yaml);
        Assert.DoesNotContain("identifier:", yaml);
    }
}
