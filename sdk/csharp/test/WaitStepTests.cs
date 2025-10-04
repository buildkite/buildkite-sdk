using Buildkite.Sdk;
using Xunit;

namespace Buildkite.Sdk.Tests
{
    public class WaitStepTests
    {
        [Fact]
        public void WaitStep_DefaultConstructor_ShouldCreateEmptyStep()
        {
            var step = new WaitStep();

            Assert.Null(step.Label);
        }

        [Fact]
        public void WaitStep_WithLabel_ShouldSetLabel()
        {
            var step = new WaitStep("Wait for tests");

            Assert.Equal("Wait for tests", step.Label);
        }

        [Fact]
        public void WaitStep_InPipeline_ShouldSerializeCorrectly()
        {
            var pipeline = new Pipeline();
            var step = new WaitStep
            {
                Label = "Wait for previous steps"
            };

            pipeline.AddStep(step);
            var json = pipeline.ToJson();

            Assert.Contains("\"label\": \"Wait for previous steps\"", json);
        }
    }
}