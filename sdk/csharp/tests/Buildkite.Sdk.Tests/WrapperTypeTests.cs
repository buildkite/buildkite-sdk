using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class DependsOnTests
{
    [Fact]
    public void DependsOn_FromString_SerializesAsString()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = "build"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("depends_on: build", yaml);
        Assert.Contains("\"depends_on\": \"build\"", json);
    }

    [Fact]
    public void DependsOn_FromStringArray_SerializesAsList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = new[] { "build", "lint" }
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- build", yaml);
        Assert.Contains("- lint", yaml);
        Assert.Contains("\"build\"", json);
        Assert.Contains("\"lint\"", json);
    }

    [Fact]
    public void DependsOn_FromDependencyArray_SerializesAsObjectList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = new[] { new Dependency { Step = "build", AllowFailure = true } }
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("step: build", yaml);
        Assert.Contains("allow_failure: true", yaml);
        Assert.Contains("\"step\": \"build\"", json);
        Assert.Contains("\"allow_failure\": true", json);
    }

    [Fact]
    public void DependsOn_FromDependencyWithoutAllowFailure_OmitsAllowFailure()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = new[] { new Dependency { Step = "build" } }
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("step: build", yaml);
        Assert.DoesNotContain("allow_failure", yaml);
        Assert.Contains("\"step\": \"build\"", json);
        Assert.DoesNotContain("allow_failure", json);
    }

    [Fact]
    public void DependsOn_MixedList_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = DependsOn.FromItems(
                "lint",
                new Dependency { Step = "build", AllowFailure = true }
            )
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- lint", yaml);
        Assert.Contains("step: build", yaml);
        Assert.Contains("\"lint\"", json);
        Assert.Contains("\"step\": \"build\"", json);
    }

    [Fact]
    public void DependsOn_WrapperInstance_NeverSerializesValueProperty()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = DependsOn.FromString("build")
        });

        var json = pipeline.ToJson();

        Assert.DoesNotContain("\"value\"", json);
        Assert.Contains("\"depends_on\": \"build\"", json);
    }

    [Fact]
    public void DependsOn_FromStrings_SerializesAsList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = DependsOn.FromStrings("build", "lint", "format")
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- build", yaml);
        Assert.Contains("- lint", yaml);
        Assert.Contains("- format", yaml);
        Assert.Contains("\"build\"", json);
        Assert.Contains("\"lint\"", json);
        Assert.Contains("\"format\"", json);
    }

    [Fact]
    public void DependsOn_FromDependencies_SerializesAsList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = DependsOn.FromDependencies(
                new Dependency { Step = "build" },
                new Dependency { Step = "lint", AllowFailure = true }
            )
        });

        var json = pipeline.ToJson();

        Assert.Contains("\"step\": \"build\"", json);
        Assert.Contains("\"step\": \"lint\"", json);
        Assert.Contains("\"allow_failure\": true", json);
    }
}

public class SoftFailTests
{
    [Theory]
    [InlineData("badvalue")]
    [InlineData("")]
    [InlineData("yes")]
    public void SoftFail_FromString_RejectsInvalidValues(string value)
    {
        Assert.Throws<ArgumentException>(() => SoftFail.FromString(value));
    }

    [Theory]
    [InlineData("true")]
    [InlineData("false")]
    public void SoftFail_FromString_AcceptsValidValues(string value)
    {
        var sf = SoftFail.FromString(value);
        Assert.Equal(value, sf.Value);
    }

    [Fact]
    public void SoftFail_True_SerializesAsBool()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = true
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("soft_fail: true", yaml);
        Assert.Contains("\"soft_fail\": true", json);
    }

    [Fact]
    public void SoftFail_False_SerializesAsBool()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = false
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("soft_fail: false", yaml);
        Assert.Contains("\"soft_fail\": false", json);
    }

    [Fact]
    public void SoftFail_FromStringEnum_SerializesAsQuotedString()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromString("true")
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        // YAML must quote "true" to distinguish from boolean true
        Assert.Contains("\"true\"", yaml);
        Assert.Contains("\"soft_fail\": \"true\"", json);
    }

    [Fact]
    public void SoftFail_FromConditions_WithIntExitStatus()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromConditions(
                new SoftFailCondition { ExitStatus = SoftFailExitStatus.FromInt(1) }
            )
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("exit_status: 1", yaml);
        Assert.Contains("\"exit_status\": 1", json);
    }

    [Fact]
    public void SoftFail_FromConditions_WithWildcardExitStatus()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromConditions(
                new SoftFailCondition { ExitStatus = SoftFailExitStatus.FromWildcard() }
            )
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("\"*\"", yaml);
        Assert.Contains("\"exit_status\": \"*\"", json);
    }

    [Fact]
    public void SoftFail_FromConditions_WithImplicitIntConversion()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromConditions(
                new SoftFailCondition { ExitStatus = 42 }
            )
        });

        var json = pipeline.ToJson();

        Assert.Contains("\"exit_status\": 42", json);
    }

    [Fact]
    public void SoftFail_FromConditions_MultipleConditions()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromConditions(
                new SoftFailCondition { ExitStatus = 1 },
                new SoftFailCondition { ExitStatus = SoftFailExitStatus.FromWildcard() }
            )
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("exit_status: 1", yaml);
        Assert.Contains("\"*\"", yaml);
        Assert.Contains("\"exit_status\": 1", json);
        Assert.Contains("\"exit_status\": \"*\"", json);
    }

    [Fact]
    public void SoftFail_OnTriggerStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy",
            SoftFail = true
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("soft_fail: true", yaml);
        Assert.Contains("\"soft_fail\": true", json);
    }

    [Fact]
    public void SoftFail_WrapperInstance_NeverSerializesValueProperty()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            SoftFail = SoftFail.FromBool(true)
        });

        var json = pipeline.ToJson();

        Assert.DoesNotContain("\"value\"", json);
        Assert.Contains("\"soft_fail\": true", json);
    }
}

public class CommandTests
{
    [Fact]
    public void Command_FromString_SerializesAsString()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("command: make test", yaml);
        Assert.Contains("\"command\": \"make test\"", json);
    }

    [Fact]
    public void Command_FromStringArray_SerializesAsList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = new[] { "make build", "make test" }
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- make build", yaml);
        Assert.Contains("- make test", yaml);
        Assert.Contains("\"make build\"", json);
        Assert.Contains("\"make test\"", json);
    }

    [Fact]
    public void Command_FromList_FactoryMethod()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = Command.FromList("echo hello", "echo world")
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- echo hello", yaml);
        Assert.Contains("- echo world", yaml);
        Assert.Contains("\"echo hello\"", json);
        Assert.Contains("\"echo world\"", json);
    }

    [Fact]
    public void Command_FromString_FactoryMethod()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = Command.FromString("make test")
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("command: make test", yaml);
        Assert.Contains("\"command\": \"make test\"", json);
    }

    [Fact]
    public void Command_CommandsAliasDelegatesToCommand()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Commands = "make test"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("command: make test", yaml);
        Assert.DoesNotContain("commands:", yaml);
    }

    [Fact]
    public void Command_WrapperInstance_NeverSerializesValueProperty()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test"
        });

        var json = pipeline.ToJson();

        Assert.DoesNotContain("\"value\"", json);
        Assert.Contains("\"command\": \"make test\"", json);
    }

    [Fact]
    public void Command_BooleanLikeString_IsQuotedInYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "true"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("\"true\"", yaml);
    }
}

public class BranchesTests
{
    [Fact]
    public void Branches_FromString_SerializesAsString()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = "main"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("branches: main", yaml);
        Assert.Contains("\"branches\": \"main\"", json);
    }

    [Fact]
    public void Branches_FromStringArray_SerializesAsList()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = new[] { "main", "develop" }
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("- main", yaml);
        Assert.Contains("- develop", yaml);
        Assert.Contains("\"main\"", json);
        Assert.Contains("\"develop\"", json);
    }

    [Fact]
    public void Branches_FromList_FactoryMethod()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = Branches.FromList("main", "staging", "production")
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("- main", yaml);
        Assert.Contains("- staging", yaml);
        Assert.Contains("- production", yaml);
    }

    [Fact]
    public void Branches_FromString_FactoryMethod()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = Branches.FromString("main")
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("branches: main", yaml);
    }

    [Fact]
    public void Branches_WrapperInstance_NeverSerializesValueProperty()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = "main"
        });

        var json = pipeline.ToJson();

        Assert.DoesNotContain("\"value\"", json);
        Assert.Contains("\"branches\": \"main\"", json);
    }

    [Fact]
    public void Branches_BooleanLikeString_IsQuotedInYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Branches = "yes"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("\"yes\"", yaml);
    }

    [Fact]
    public void Branches_OnBlockStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new BlockStep
        {
            Block = "Approve",
            Branches = "main"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("branches: main", yaml);
        Assert.Contains("\"branches\": \"main\"", json);
    }

    [Fact]
    public void Branches_OnWaitStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new WaitStep
        {
            Branches = "main"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("branches: main", yaml);
        Assert.Contains("\"branches\": \"main\"", json);
    }

    [Fact]
    public void Branches_OnInputStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new InputStep
        {
            Input = "Config",
            Branches = "main"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("branches: main", yaml);
        Assert.Contains("\"branches\": \"main\"", json);
    }

    [Fact]
    public void Branches_OnTriggerStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy",
            Branches = "main"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("branches: main", yaml);
        Assert.Contains("\"branches\": \"main\"", json);
    }
}

public class SkipTests
{
    [Fact]
    public void Skip_FromBool_SerializesAsBool()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Skip = true
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("skip: true", yaml);
        Assert.Contains("\"skip\": true", json);
    }

    [Fact]
    public void Skip_FromReason_SerializesAsString()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Skip = "Not needed on this branch"
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        Assert.Contains("skip: Not needed on this branch", yaml);
        Assert.Contains("\"skip\": \"Not needed on this branch\"", json);
    }

    [Fact]
    public void Skip_FromReasonThatLooksLikeYamlBool_IsQuotedInYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Skip = Skip.FromReason("true")
        });

        var yaml = pipeline.ToYaml();
        var json = pipeline.ToJson();

        // YAML must quote "true" to prevent parsing as boolean
        Assert.Contains("\"true\"", yaml);
        Assert.Contains("\"skip\": \"true\"", json);
    }

    [Fact]
    public void Skip_FromReasonYes_IsQuotedInYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Skip = Skip.FromReason("yes")
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("\"yes\"", yaml);
    }

    [Fact]
    public void Skip_WrapperInstance_NeverSerializesValueProperty()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Skip = Skip.FromReason("Skipped")
        });

        var json = pipeline.ToJson();

        Assert.DoesNotContain("\"value\"", json);
        Assert.Contains("\"skip\": \"Skipped\"", json);
    }

    [Fact]
    public void Skip_OnGroupStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new GroupStep
        {
            Group = "Tests",
            Skip = "Not on this branch",
            Steps = new List<IGroupStep>
            {
                new CommandStep { Label = "Test", Command = "make test" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("skip: Not on this branch", yaml);
    }

    [Fact]
    public void Skip_OnTriggerStep_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new TriggerStep
        {
            Trigger = "deploy",
            Skip = true
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("skip: true", yaml);
    }
}
