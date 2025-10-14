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
        public List<DependsOn>? DependsOn { get; set; }
        public string? DependsOnString { get; set; }
        public Dictionary<string, object>? Env { get; set; }
        public string? Id { get; set; }
        public string? Identifier { get; set; }
        public string? If { get; set; }
        public string? Key { get; set; }
        public string? Label { get; set; }
        public string? Name { get; set; }
        public List<Notify>? Notify { get; set; }
        public int? Priority { get; set; }
        public Retry? Retry { get; set; }
        public Signature? Signature { get; set; }
        public object? Skip { get; set; }
        public object? SoftFail { get; set; }
        public int? TimeoutInMinutes { get; set; }
    }
}