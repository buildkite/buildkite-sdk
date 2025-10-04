namespace Buildkite.Sdk
{
    public class WaitStep : PipelineStep
    {
        public string? Wait { get; set; }
        public bool? ContinueOnFailure { get; set; }

        public WaitStep()
        {
        }

        public WaitStep(string label)
        {
            Label = label;
        }
    }
}