using Buildkite.Sdk;
using Xunit;

namespace Buildkite.Sdk.Tests
{
    public class GroupStepTests
    {
        [Fact]
        public void GroupStep_WithGroup_ShouldSetGroup()
        {
            var step = new GroupStep("Build Steps");

            Assert.Equal("Build Steps", step.Group);
            Assert.NotNull(step.Steps);
            Assert.Empty(step.Steps);
        }

        [Fact]
        public void GroupStep_WithLabelAndGroup_ShouldSetBoth()
        {
            var step = new GroupStep("Build", "Build Steps");

            Assert.Equal("Build", step.Label);
            Assert.Equal("Build Steps", step.Group);
        }

        [Fact]
        public void GroupStep_AddStep_ShouldAddStepToCollection()
        {
            var group = new GroupStep("Test Group");
            var commandStep = new CommandStep("echo test");

            group.AddStep(commandStep);

            Assert.Single(group.Steps!);
            Assert.Equal(commandStep, group.Steps![0]);
        }

        [Fact]
        public void GroupStep_InPipeline_ShouldSerializeCorrectly()
        {
            var pipeline = new Pipeline();
            var group = new GroupStep
            {
                Label = "Build and Test",
                Group = "Build Group"
            };

            group.AddStep(new CommandStep("Build", "dotnet build"));
            group.AddStep(new CommandStep("Test", "dotnet test"));

            pipeline.AddStep(group);
            var json = pipeline.ToJson();

            Assert.Contains("\"label\": \"Build and Test\"", json);
            Assert.Contains("\"group\": \"Build Group\"", json);
            Assert.Contains("\"steps\"", json);
        }
    }
}