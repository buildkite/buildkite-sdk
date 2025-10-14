using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public class GroupStep : PipelineStep
    {
        public string? Group { get; set; }
        public List<PipelineStep>? Steps { get; set; }
        public List<Notify>? Notify { get; set; }

        public GroupStep()
        {
        }

        public GroupStep(string group)
        {
            Group = group;
            Steps = new List<PipelineStep>();
        }

        public GroupStep(string label, string group)
        {
            Label = label;
            Group = group;
            Steps = new List<PipelineStep>();
        }

        public GroupStep AddStep(PipelineStep step)
        {
            Steps ??= new List<PipelineStep>();
            Steps.Add(step);
            return this;
        }
    }
}