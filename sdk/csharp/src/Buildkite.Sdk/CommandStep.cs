using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public class CommandStep : PipelineStep
    {
        public List<string>? ArtifactPaths { get; set; }
        public string? ArtifactPathsString { get; set; }
        public object? Cache { get; set; }
        public bool? CancelOnBuildFailing { get; set; }
        public string? Command { get; set; }
        public List<string>? Commands { get; set; }
        public int? Concurrency { get; set; }
        public string? ConcurrencyGroup { get; set; }
        public ConcurrencyMethod? ConcurrencyMethod { get; set; }
        public object? Matrix { get; set; }
        public int? Parallelism { get; set; }
        public object? Plugins { get; set; }

        public CommandStep()
        {
        }

        public CommandStep(string command)
        {
            Command = command;
        }

        public CommandStep(string label, string command)
        {
            Label = label;
            Command = command;
        }

        public CommandStep(string label, List<string> commands)
        {
            Label = label;
            Commands = commands;
        }
    }
}