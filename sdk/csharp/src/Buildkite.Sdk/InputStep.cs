using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public class InputStep : PipelineStep
    {
        public string? Input { get; set; }
        public string? Prompt { get; set; }
        public List<Field>? Fields { get; set; }
        public BlockedState? BlockedState { get; set; }

        public InputStep()
        {
        }

        public InputStep(string input)
        {
            Input = input;
        }

        public InputStep(string label, string input)
        {
            Label = label;
            Input = input;
        }
    }
}