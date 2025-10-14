using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public class BlockStep : PipelineStep
    {
        public string? Block { get; set; }
        public string? Prompt { get; set; }
        public List<Field>? Fields { get; set; }
        public BlockedState? BlockedState { get; set; }

        public BlockStep()
        {
        }

        public BlockStep(string block)
        {
            Block = block;
        }

        public BlockStep(string label, string block)
        {
            Label = label;
            Block = block;
        }
    }
}
