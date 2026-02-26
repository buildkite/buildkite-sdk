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

    // Build Information
    [Fact]
    public void BuildId_ReturnsValue()
    {
        SetEnv("BUILDKITE_BUILD_ID", "abc-123");
        Assert.Equal("abc-123", EnvironmentVariable.BuildId);
    }

    [Fact]
    public void BuildAuthor_ReturnsValue()
    {
        SetEnv("BUILDKITE_BUILD_AUTHOR", "Jane Doe");
        Assert.Equal("Jane Doe", EnvironmentVariable.BuildAuthor);
    }

    [Fact]
    public void BuildAuthorEmail_ReturnsValue()
    {
        SetEnv("BUILDKITE_BUILD_AUTHOR_EMAIL", "jane@example.com");
        Assert.Equal("jane@example.com", EnvironmentVariable.BuildAuthorEmail);
    }

    [Fact]
    public void BuildCreatorTeams_ReturnsValue()
    {
        SetEnv("BUILDKITE_BUILD_CREATOR_TEAMS", "deploy:ops");
        Assert.Equal("deploy:ops", EnvironmentVariable.BuildCreatorTeams);
    }

    [Fact]
    public void ClusterId_ReturnsValue()
    {
        SetEnv("BUILDKITE_CLUSTER_ID", "cluster-uuid-123");
        Assert.Equal("cluster-uuid-123", EnvironmentVariable.ClusterId);
    }

    // Repository Information
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
    public void RepoMirror_ReturnsValue()
    {
        SetEnv("BUILDKITE_REPO_MIRROR", "/var/lib/buildkite-agent/git-mirrors");
        Assert.Equal("/var/lib/buildkite-agent/git-mirrors", EnvironmentVariable.RepoMirror);
    }

    [Fact]
    public void Refspec_ReturnsValue()
    {
        SetEnv("BUILDKITE_REFSPEC", "+refs/heads/*:refs/remotes/origin/*");
        Assert.Equal("+refs/heads/*:refs/remotes/origin/*", EnvironmentVariable.Refspec);
    }

    [Fact]
    public void CleanCheckout_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_CLEAN_CHECKOUT", "true");
        Assert.True(EnvironmentVariable.CleanCheckout);
    }

    [Fact]
    public void CleanCheckout_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_CLEAN_CHECKOUT", null);
        Assert.False(EnvironmentVariable.CleanCheckout);
    }

    [Fact]
    public void GitCleanFlags_ReturnsValue()
    {
        SetEnv("BUILDKITE_GIT_CLEAN_FLAGS", "-ffxdq");
        Assert.Equal("-ffxdq", EnvironmentVariable.GitCleanFlags);
    }

    [Fact]
    public void GitCloneFlags_ReturnsValue()
    {
        SetEnv("BUILDKITE_GIT_CLONE_FLAGS", "--depth 1");
        Assert.Equal("--depth 1", EnvironmentVariable.GitCloneFlags);
    }

    // Pipeline Information
    [Fact]
    public void PipelineSlug_ReturnsValue()
    {
        SetEnv("BUILDKITE_PIPELINE_SLUG", "my-pipeline");
        Assert.Equal("my-pipeline", EnvironmentVariable.PipelineSlug);
    }

    [Fact]
    public void PipelineTeams_ReturnsValue()
    {
        SetEnv("BUILDKITE_PIPELINE_TEAMS", "backend:frontend");
        Assert.Equal("backend:frontend", EnvironmentVariable.PipelineTeams);
    }

    // Organization Information
    [Fact]
    public void OrganizationId_ReturnsValue()
    {
        SetEnv("BUILDKITE_ORGANIZATION_ID", "org-uuid-456");
        Assert.Equal("org-uuid-456", EnvironmentVariable.OrganizationId);
    }

    // Agent Information
    [Fact]
    public void AgentPid_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_PID", "12345");
        Assert.Equal("12345", EnvironmentVariable.AgentPid);
    }

    [Fact]
    public void AgentAccessToken_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_ACCESS_TOKEN", "token-abc");
        Assert.Equal("token-abc", EnvironmentVariable.AgentAccessToken);
    }

    [Fact]
    public void AgentExperiment_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_EXPERIMENT", "normalised-upload-paths,resolve-commit-after-checkout");
        Assert.Equal("normalised-upload-paths,resolve-commit-after-checkout", EnvironmentVariable.AgentExperiment);
    }

    [Fact]
    public void AgentDisconnectAfterJob_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB", "true");
        Assert.Equal("true", EnvironmentVariable.AgentDisconnectAfterJob);
    }

    [Fact]
    public void AgentEndpoint_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_ENDPOINT", "https://agent.buildkite.com/v3");
        Assert.Equal("https://agent.buildkite.com/v3", EnvironmentVariable.AgentEndpoint);
    }

    [Fact]
    public void AgentHealthCheckAddr_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_HEALTH_CHECK_ADDR", "localhost:8080");
        Assert.Equal("localhost:8080", EnvironmentVariable.AgentHealthCheckAddr);
    }

    [Fact]
    public void AgentDisconnectAfterIdleTimeout_ReturnsValue()
    {
        SetEnv("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT", "600");
        Assert.Equal("600", EnvironmentVariable.AgentDisconnectAfterIdleTimeout);
    }

    // Job Information
    [Fact]
    public void CommandExitStatus_ReturnsValue()
    {
        SetEnv("BUILDKITE_COMMAND_EXIT_STATUS", "0");
        Assert.Equal("0", EnvironmentVariable.CommandExitStatus);
    }

    [Fact]
    public void CommandEval_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_COMMAND_EVAL", "true");
        Assert.True(EnvironmentVariable.CommandEval);
    }

    [Fact]
    public void CommandEval_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_COMMAND_EVAL", null);
        Assert.False(EnvironmentVariable.CommandEval);
    }

    [Fact]
    public void JobLogTmpfile_ReturnsValue()
    {
        SetEnv("BUILDKITE_JOB_LOG_TMPFILE", "/tmp/buildkite-job-log-123");
        Assert.Equal("/tmp/buildkite-job-log-123", EnvironmentVariable.JobLogTmpfile);
    }

    [Fact]
    public void LastHookExitStatus_ReturnsValue()
    {
        SetEnv("BUILDKITE_LAST_HOOK_EXIT_STATUS", "0");
        Assert.Equal("0", EnvironmentVariable.LastHookExitStatus);
    }

    // Rebuild Information
    [Fact]
    public void Rebuilt_ReturnsValue()
    {
        SetEnv("BUILDKITE_REBUILT_FROM_BUILD_ID", "original-build-uuid");
        Assert.Equal("original-build-uuid", EnvironmentVariable.Rebuilt);
    }

    [Fact]
    public void RebuiltFromBuildNumber_ReturnsValue()
    {
        SetEnv("BUILDKITE_REBUILT_FROM_BUILD_NUMBER", "42");
        Assert.Equal("42", EnvironmentVariable.RebuiltFromBuildNumber);
    }

    // Pull Request Information
    [Fact]
    public void PullRequestDraft_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_PULL_REQUEST_DRAFT", "true");
        Assert.True(EnvironmentVariable.PullRequestDraft);
    }

    [Fact]
    public void PullRequestDraft_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_PULL_REQUEST_DRAFT", null);
        Assert.False(EnvironmentVariable.PullRequestDraft);
    }

    // Triggered Build Information
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

    [Fact]
    public void TriggeredFromBuildNumber_ReturnsValue()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER", "99");
        Assert.Equal("99", EnvironmentVariable.TriggeredFromBuildNumber);
    }

    [Fact]
    public void TriggeredFromBuildPipelineSlug_ReturnsValue()
    {
        SetEnv("BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG", "parent-pipeline");
        Assert.Equal("parent-pipeline", EnvironmentVariable.TriggeredFromBuildPipelineSlug);
    }

    // Paths
    [Fact]
    public void ConfigPath_ReturnsValue()
    {
        SetEnv("BUILDKITE_CONFIG_PATH", "/etc/buildkite-agent/buildkite-agent.cfg");
        Assert.Equal("/etc/buildkite-agent/buildkite-agent.cfg", EnvironmentVariable.ConfigPath);
    }

    [Fact]
    public void EnvFile_ReturnsValue()
    {
        SetEnv("BUILDKITE_ENV_FILE", "/tmp/job-env-123");
        Assert.Equal("/tmp/job-env-123", EnvironmentVariable.EnvFile);
    }

    // S3 Configuration
    [Fact]
    public void S3DefaultRegion_ReturnsValue()
    {
        SetEnv("BUILDKITE_S3_DEFAULT_REGION", "us-east-1");
        Assert.Equal("us-east-1", EnvironmentVariable.S3DefaultRegion);
    }

    [Fact]
    public void S3SseEnabled_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_S3_SSE_ENABLED", "true");
        Assert.True(EnvironmentVariable.S3SseEnabled);
    }

    [Fact]
    public void S3SseEnabled_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_S3_SSE_ENABLED", null);
        Assert.False(EnvironmentVariable.S3SseEnabled);
    }

    // Cancel Configuration
    [Fact]
    public void CancelGracePeriod_ReturnsValue()
    {
        SetEnv("BUILDKITE_CANCEL_GRACE_PERIOD", "10");
        Assert.Equal("10", EnvironmentVariable.CancelGracePeriod);
    }

    [Fact]
    public void CancelSignal_ReturnsValue()
    {
        SetEnv("BUILDKITE_CANCEL_SIGNAL", "SIGTERM");
        Assert.Equal("SIGTERM", EnvironmentVariable.CancelSignal);
    }

    // Plugin Information
    [Fact]
    public void PluginsEnabled_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_PLUGINS_ENABLED", "true");
        Assert.True(EnvironmentVariable.PluginsEnabled);
    }

    [Fact]
    public void PluginsEnabled_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_PLUGINS_ENABLED", null);
        Assert.False(EnvironmentVariable.PluginsEnabled);
    }

    [Fact]
    public void PluginConfiguration_ReturnsValue()
    {
        SetEnv("BUILDKITE_PLUGIN_CONFIGURATION", "{\"image\":\"node:18\"}");
        Assert.Equal("{\"image\":\"node:18\"}", EnvironmentVariable.PluginConfiguration);
    }

    [Fact]
    public void PluginName_ReturnsValue()
    {
        SetEnv("BUILDKITE_PLUGIN_NAME", "DOCKER");
        Assert.Equal("DOCKER", EnvironmentVariable.PluginName);
    }

    [Fact]
    public void PluginValidation_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_PLUGIN_VALIDATION", "true");
        Assert.True(EnvironmentVariable.PluginValidation);
    }

    // Unblocker Information
    [Fact]
    public void UnblockerTeams_ReturnsValue()
    {
        SetEnv("BUILDKITE_UNBLOCKER_TEAMS", "deploy:release");
        Assert.Equal("deploy:release", EnvironmentVariable.UnblockerTeams);
    }

    // GitHub Deployment Information
    [Fact]
    public void GithubDeploymentId_ReturnsValue()
    {
        SetEnv("BUILDKITE_GITHUB_DEPLOYMENT_ID", "deploy-123");
        Assert.Equal("deploy-123", EnvironmentVariable.GithubDeploymentId);
    }

    [Fact]
    public void GithubDeploymentEnvironment_ReturnsValue()
    {
        SetEnv("BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT", "production");
        Assert.Equal("production", EnvironmentVariable.GithubDeploymentEnvironment);
    }

    [Fact]
    public void GithubDeploymentTask_ReturnsValue()
    {
        SetEnv("BUILDKITE_GITHUB_DEPLOYMENT_TASK", "deploy");
        Assert.Equal("deploy", EnvironmentVariable.GithubDeploymentTask);
    }

    [Fact]
    public void GithubDeploymentPayload_ReturnsValue()
    {
        SetEnv("BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD", "{\"ref\":\"main\"}");
        Assert.Equal("{\"ref\":\"main\"}", EnvironmentVariable.GithubDeploymentPayload);
    }

    // Shell and Hooks
    [Fact]
    public void Shell_ReturnsValue()
    {
        SetEnv("BUILDKITE_SHELL", "/bin/bash -e -c");
        Assert.Equal("/bin/bash -e -c", EnvironmentVariable.Shell);
    }

    [Fact]
    public void LocalHooksEnabled_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_LOCAL_HOOKS_ENABLED", "true");
        Assert.True(EnvironmentVariable.LocalHooksEnabled);
    }

    [Fact]
    public void LocalHooksEnabled_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_LOCAL_HOOKS_ENABLED", null);
        Assert.False(EnvironmentVariable.LocalHooksEnabled);
    }

    [Fact]
    public void SshKeyscan_ReturnsTrue_WhenSetToTrue()
    {
        SetEnv("BUILDKITE_SSH_KEYSCAN", "true");
        Assert.True(EnvironmentVariable.SshKeyscan);
    }

    [Fact]
    public void SshKeyscan_ReturnsFalse_WhenUnset()
    {
        SetEnv("BUILDKITE_SSH_KEYSCAN", null);
        Assert.False(EnvironmentVariable.SshKeyscan);
    }

    // Tracing
    [Fact]
    public void TracingBackend_ReturnsValue()
    {
        SetEnv("BUILDKITE_TRACING_BACKEND", "datadog");
        Assert.Equal("datadog", EnvironmentVariable.TracingBackend);
    }

    // Miscellaneous
    [Fact]
    public void IgnoredEnv_ReturnsValue()
    {
        SetEnv("BUILDKITE_IGNORED_ENV", "BUILDKITE_AGENT_ACCESS_TOKEN");
        Assert.Equal("BUILDKITE_AGENT_ACCESS_TOKEN", EnvironmentVariable.IgnoredEnv);
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
    public void StringProperty_ReturnsNull_WhenUnset()
    {
        SetEnv("BUILDKITE_BUILD_ID", null);
        Assert.Null(EnvironmentVariable.BuildId);
    }
}
