using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class GroupStepTests
{
    [Fact]
    public void GroupStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new GroupStep
        {
            Group = ":test_tube: Tests",
            Steps = new List<IGroupStep>
            {
                new CommandStep { Label = "Unit Tests", Command = "dotnet test --filter Category=Unit" },
                new CommandStep { Label = "Integration Tests", Command = "dotnet test --filter Category=Integration" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":test_tube: Tests", yaml);
        Assert.Contains("Unit Tests", yaml);
        Assert.Contains("Integration Tests", yaml);
    }

    [Fact]
    public void GroupStep_WithKey_GeneratesKeyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new GroupStep
        {
            Group = "Build Group",
            Key = "build-group",
            Steps = new List<IGroupStep>
            {
                new CommandStep { Label = "Build", Command = "make" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("key: build-group", yaml);
    }

    [Fact]
    public void GroupStep_UsingFluentApi_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        
        var group = new GroupStep { Group = ":hammer: Build Steps" };
        group
            .AddStep(new CommandStep { Label = "Compile", Command = "dotnet build" })
            .AddStep(new CommandStep { Label = "Package", Command = "dotnet pack" });
        
        pipeline.AddStep(group);

        var yaml = pipeline.ToYaml();

        Assert.Contains(":hammer: Build Steps", yaml);
        Assert.Contains("Compile", yaml);
        Assert.Contains("Package", yaml);
    }

    [Fact]
    public void GroupStep_WithDependsOn_GeneratesDependencyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Key = "setup", Label = "Setup", Command = "make setup" });
        pipeline.AddStep(new GroupStep
        {
            Group = "Tests",
            DependsOn = "setup",
            Steps = new List<IGroupStep>
            {
                new CommandStep { Label = "Test", Command = "make test" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("depends_on: setup", yaml);
    }
}
