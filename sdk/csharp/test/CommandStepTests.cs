using Buildkite.Sdk;
using Xunit;

namespace Buildkite.Sdk.Tests
{
    public class CommandStepTests
    {
        [Fact]
        public void CommandStep_WithCommand_ShouldSetCommand()
        {
            var step = new CommandStep("echo 'Hello, world!'");

            Assert.Equal("echo 'Hello, world!'", step.Command);
        }

        [Fact]
        public void CommandStep_WithLabelAndCommand_ShouldSetBoth()
        {
            var step = new CommandStep("test-label", "echo test");

            Assert.Equal("test-label", step.Label);
            Assert.Equal("echo test", step.Command);
        }

        [Fact]
        public void CommandStep_WithLabelAndCommands_ShouldSetCommandsList()
        {
            var commands = new List<string> { "echo test1", "echo test2" };
            var step = new CommandStep("test-label", commands);

            Assert.Equal("test-label", step.Label);
            Assert.Equal(commands, step.Commands);
        }

        [Fact]
        public void CommandStep_InPipeline_ShouldSerializeCorrectly()
        {
            var pipeline = new Pipeline();
            var step = new CommandStep
            {
                Label = "Build",
                Command = "dotnet build",
                ArtifactPaths = new List<string> { "dist/**/*" }
            };

            pipeline.AddStep(step);
            var json = pipeline.ToJson();

            Assert.Contains("\"label\": \"Build\"", json);
            Assert.Contains("\"command\": \"dotnet build\"", json);
            Assert.Contains("\"artifactPaths\"", json);
        }
    }
}