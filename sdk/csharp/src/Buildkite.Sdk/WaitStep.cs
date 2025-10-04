namespace Buildkite.Sdk
{
    public class WaitStep : PipelineStep
    {
        public WaitStep()
        {
        }

        public WaitStep(string label)
        {
            Label = label;
        }
    }
}