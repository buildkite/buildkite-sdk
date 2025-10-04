using Buildkite.Sdk;
using Xunit;

namespace Buildkite.Sdk.Tests
{
    public class BlockStepTests
    {
        [Fact]
        public void BlockStep_WithBlock_ShouldSetBlock()
        {
            var step = new BlockStep("Deploy to production");

            Assert.Equal("Deploy to production", step.Block);
        }

        [Fact]
        public void BlockStep_WithLabelAndBlock_ShouldSetBoth()
        {
            var step = new BlockStep("Manual Deploy", "Deploy to production");

            Assert.Equal("Manual Deploy", step.Label);
            Assert.Equal("Deploy to production", step.Block);
        }

        [Fact]
        public void BlockStep_InPipeline_ShouldSerializeCorrectly()
        {
            var pipeline = new Pipeline();
            var step = new BlockStep
            {
                Label = "Release gate",
                Block = "Deploy to production",
                Prompt = "Ready to deploy?"
            };

            pipeline.AddStep(step);
            var json = pipeline.ToJson();

            Assert.Contains("\"label\": \"Release gate\"", json);
            Assert.Contains("\"block\": \"Deploy to production\"", json);
            Assert.Contains("\"prompt\": \"Ready to deploy?\"", json);
        }
    }
}
