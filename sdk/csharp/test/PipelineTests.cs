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
            Assert.Contains("\"command\": \"echo", json);
            Assert.Contains("Hello, world", json);
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
            Assert.Contains("\"FOO\": \"bar\"", json); // Dictionary keys are not affected by PropertyNamingPolicy
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

        [Fact]
        public void Pipeline_ToYaml_ShouldOmitUnsetProperties()
        {
            var pipeline = new Pipeline();
            pipeline.AddStep(new CommandStep
            {
                Label = "test",
                Command = "echo 'Hello, world!'"
            });

            var yaml = pipeline.ToYaml();

            // Should contain only set properties
            Assert.Contains("steps:", yaml);
            Assert.Contains("label: test", yaml);
            Assert.Contains("command: echo 'Hello, world!'", yaml);

            // Should NOT contain unset properties
            Assert.DoesNotContain("artifactPaths:", yaml);
            Assert.DoesNotContain("cache:", yaml);
            Assert.DoesNotContain("concurrency:", yaml);
            Assert.DoesNotContain("parallelism:", yaml);
            Assert.DoesNotContain("plugins:", yaml);
            Assert.DoesNotContain("agents:", yaml);
            Assert.DoesNotContain("env:", yaml);
            Assert.DoesNotContain("retry:", yaml);
        }
    }
}