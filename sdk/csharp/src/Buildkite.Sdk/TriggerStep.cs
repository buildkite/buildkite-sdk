using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public class TriggerStep : PipelineStep
    {
        public string? Trigger { get; set; }
        public Build? Build { get; set; }
        public bool? Async { get; set; }

        public TriggerStep()
        {
        }

        public TriggerStep(string trigger)
        {
            Trigger = trigger;
        }

        public TriggerStep(string label, string trigger)
        {
            Label = label;
            Trigger = trigger;
        }
    }
}