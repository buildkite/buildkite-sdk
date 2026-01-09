namespace Buildkite.Sdk.Schema;

/// <summary>
/// Query rules to target specific agents as key-value pairs.
/// </summary>
public class AgentsObject : Dictionary<string, string>
{
    public AgentsObject() : base() { }
    public AgentsObject(IDictionary<string, string> dictionary) : base(dictionary) { }
}

/// <summary>
/// Query rules to target specific agents as a list of strings in "key=value" format.
/// </summary>
public class AgentsList : List<string>
{
    public AgentsList() : base() { }
    public AgentsList(IEnumerable<string> collection) : base(collection) { }
}
