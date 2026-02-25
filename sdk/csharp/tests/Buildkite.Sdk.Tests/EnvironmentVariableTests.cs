using Buildkite.Sdk;
using Xunit;

namespace Buildkite.Sdk.Tests;

[Collection("EnvironmentVariables")]
public class EnvironmentVariableTests : IDisposable
{
    private readonly List<string> _setVars = new();

    private void SetEnv(string name, string? value)
    {
        System.Environment.SetEnvironmentVariable(name, value);
        _setVars.Add(name);
    }

    public void Dispose()
    {
        foreach (var name in _setVars)
            System.Environment.SetEnvironmentVariable(name, null);
    }

    [Fact]
    public void BuildId_ReturnsValue()
    {
        SetEnv("BUILDKITE_BUILD_ID", "abc-123");
        Assert.Equal("abc-123", EnvironmentVariable.BuildId);
    }

    [Fact]
    public void Branch_ReturnsValue()
    {
        SetEnv("BUILDKITE_BRANCH", "main");
        Assert.Equal("main", EnvironmentVariable.Branch);
    }

    [Fact]
    public void Commit_ReturnsValue()
    {
        SetEnv("BUILDKITE_COMMIT", "deadbeef");
        Assert.Equal("deadbeef", EnvironmentVariable.Commit);
    }

    [Fact]
    public void PipelineSlug_ReturnsValue()
    {
        SetEnv("BUILDKITE_PIPELINE_SLUG", "my-pipeline");
        Assert.Equal("my-pipeline", EnvironmentVariable.PipelineSlug);
    }

    [Fact]
    public void StringProperty_ReturnsNull_WhenUnset()
    {
        SetEnv("BUILDKITE_BUILD_ID", null);
        Assert.Null(EnvironmentVariable.BuildId);
    }

    [Fact]
    public void IsBuildkite_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE", "true");
        Assert.True(EnvironmentVariable.IsBuildkite);
    }

    [Fact]
    public void IsBuildkite_ReturnsFalse_WhenSetToOtherValue()
    {
        SetEnv("BUILDKITE", "false");
        Assert.False(EnvironmentVariable.IsBuildkite);
    }

    [Fact]
    public void IsBuildkite_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE", null);
        Assert.False(EnvironmentVariable.IsBuildkite);
    }

    [Fact]
    public void IsCi_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("CI", "true");
        Assert.True(EnvironmentVariable.IsCi);
    }

    [Fact]
    public void IsCi_ReturnsFalse_WhenUnset()
    {
        SetEnv("CI", null);
        Assert.False(EnvironmentVariable.IsCi);
    }

    [Fact]
    public void AgentDebug_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_AGENT_DEBUG", "true");
        Assert.True(EnvironmentVariable.AgentDebug);
    }

    [Fact]
    public void AgentDebug_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_AGENT_DEBUG", null);
        Assert.False(EnvironmentVariable.AgentDebug);
    }

    [Fact]
    public void GitSubmodules_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_GIT_SUBMODULES", "true");
        Assert.True(EnvironmentVariable.GitSubmodules);
    }

    [Fact]
    public void GitSubmodules_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_GIT_SUBMODULES", null);
        Assert.False(EnvironmentVariable.GitSubmodules);
    }

    [Fact]
    public void TriggeredFromBuildId_ReturnsTrue_WhenSet()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_ID", "some-build-id");
        Assert.True(EnvironmentVariable.TriggeredFromBuildId);
    }

    [Fact]
    public void TriggeredFromBuildId_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_ID", null);
        Assert.False(EnvironmentVariable.TriggeredFromBuildId);
    }

    [Fact]
    public void TriggeredFromBuildId_ReturnsFalse_WhenEmpty()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_ID", "");
        Assert.False(EnvironmentVariable.TriggeredFromBuildId);
    }

    [Fact]
    public void TriggeredFromBuildIdValue_ReturnsStringValue()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_ID", "abc-123");
        Assert.Equal("abc-123", EnvironmentVariable.TriggeredFromBuildIdValue);
        Assert.True(EnvironmentVariable.TriggeredFromBuildId);
    }
}
