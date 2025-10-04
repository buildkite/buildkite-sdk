using Buildkite.Sdk;
using Xunit;
using System.Text.Json;

namespace Buildkite.Sdk.Tests
{
    public class PipelineTests
    {
        [Fact]
        public void Pipeline_EmptyPipeline_ShouldSerializeToEmptyObject()
        {
            var pipeline = new Pipeline();
            var json = pipeline.ToJson();

            // With NullValueHandling.Ignore, an empty pipeline should serialize to just {}
            Assert.Equal("{}", json);
        }

        [Fact]
        public void Pipeline_WithCommandStep_ShouldSerializeCorrectly()
        {
            var pipeline = new Pipeline();
            pipeline.AddStep(new CommandStep
            {
                Label = "some-label",
                Command = "echo 'Hello, world!'"
            });

            var json = pipeline.ToJson();

            Assert.Contains("\"label\": \"some-label\"", json);
            Assert.Contains("\"command\": \"echo 'Hello, world!'\"", json);
        }

        [Fact]
        public void Pipeline_WithAgents_ShouldSerializeAgentsCorrectly()
        {
            var pipeline = new Pipeline();
            pipeline.AddAgent("queue", "hosted");
            pipeline.AddStep(new CommandStep { Command = "echo test" });

            var json = pipeline.ToJson();

            Assert.Contains("\"agents\"", json);
            Assert.Contains("\"queue\": \"hosted\"", json);
        }

        [Fact]
        public void Pipeline_WithEnvironmentVariables_ShouldSerializeEnvCorrectly()
        {
            var pipeline = new Pipeline();
            pipeline.AddEnvironmentVariable("FOO", "bar");
            pipeline.AddStep(new CommandStep { Command = "echo test" });

            var json = pipeline.ToJson();

            Assert.Contains("\"env\"", json);
            Assert.Contains("\"foo\": \"bar\"", json); // CamelCase converts keys to lowercase
        }

        [Fact]
        public void Pipeline_ToYaml_ShouldGenerateValidYaml()
        {
            var pipeline = new Pipeline();
            pipeline.AddStep(new CommandStep
            {
                Label = "test",
                Command = "echo 'Hello, world!'"
            });

            var yaml = pipeline.ToYaml();

            Assert.Contains("steps:", yaml);
            Assert.Contains("label: test", yaml);
            Assert.Contains("command: echo 'Hello, world!'", yaml);
        }
    }
}