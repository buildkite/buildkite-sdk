using System.Collections.Generic;

namespace Buildkite.Sdk
{
    public abstract class PipelineStep
    {
        public List<string>? Agents { get; set; }
        public Dictionary<string, object>? AgentsDictionary { get; set; }
        public bool? AllowDependencyFailure { get; set; }
        public List<string>? Branches { get; set; }
        public string? BranchesString { get; set; }
        public bool? CancelOnBuildFailing { get; set; }
        public List<string>? DependsOn { get; set; }
        public string? DependsOnString { get; set; }
        public Dictionary<string, object>? Env { get; set; }
        public string? Id { get; set; }
        public string? Identifier { get; set; }
        public string? If { get; set; }
        public string? Key { get; set; }
        public string? Label { get; set; }
        public string? Name { get; set; }
        public List<object>? Notify { get; set; }
        public int? Priority { get; set; }
        public object? Retry { get; set; }
        public object? Signature { get; set; }
        public object? Skip { get; set; }
        public object? SoftFail { get; set; }
        public int? TimeoutInMinutes { get; set; }
    }
}