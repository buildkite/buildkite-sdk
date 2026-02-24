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

        var json = pipeline.ToJson();

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
    [Fact]
    public void SoftFail_FromBool_SerializesAsBool()
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
        // JSON always quotes strings, so this is straightforward
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
