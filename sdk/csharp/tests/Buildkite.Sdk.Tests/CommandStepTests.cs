using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class CommandStepTests
{
    [Fact]
    public void CommandStep_WithBasicProperties_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = ":dotnet: Build",
            Key = "build",
            Command = "dotnet build --configuration Release"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains(":dotnet: Build", yaml);
        Assert.Contains("key: build", yaml);
        Assert.Contains("command: dotnet build --configuration Release", yaml);
    }

    [Fact]
    public void CommandStep_WithAgents_GeneratesAgentConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Agents = new AgentsObject { ["queue"] = "linux", ["os"] = "ubuntu" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("queue: linux", yaml);
        Assert.Contains("os: ubuntu", yaml);
    }

    [Fact]
    public void CommandStep_WithEnv_GeneratesEnvConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, string> { ["NODE_ENV"] = "production" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("NODE_ENV: production", yaml);
    }

    [Fact]
    public void CommandStep_WithMultipleEnvVars_GeneratesAllEnvVars()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, string>
            {
                ["NODE_ENV"] = "production",
                ["CI"] = "true",
                ["BUILD_NUMBER"] = "42"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("NODE_ENV: production", yaml);
        Assert.Contains("CI: true", yaml);
        Assert.Contains("BUILD_NUMBER: 42", yaml);
    }

    [Fact]
    public void CommandStep_WithEnv_PlacesEnvUnderStep()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, string>
            {
                ["NODE_ENV"] = "production",
                ["CI"] = "true"
            }
        });

        var yaml = pipeline.ToYaml();

        // Verify env block appears as a child of the step, not at pipeline level
        var lines = yaml.Split('\n');
        var envLineIdx = Array.FindIndex(lines, l => l.TrimEnd() == "  env:");
        Assert.True(envLineIdx >= 0, "Expected indented 'env:' under step");
        Assert.StartsWith("    NODE_ENV:", lines[envLineIdx + 1].TrimEnd());
    }

    [Fact]
    public void CommandStep_WithEmptyEnv_SerializesAsEmptyMapping()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, string>()
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("env: {}", yaml);
    }

    [Fact]
    public void CommandStep_WithNullEnv_OmitsEnvFromOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = null
        });

        var yaml = pipeline.ToYaml();

        Assert.DoesNotContain("env", yaml);
    }

    [Fact]
    public void CommandStep_WithEmptyStringEnvValue_PreservesEntry()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Env = new Dictionary<string, string> { ["EMPTY_VAR"] = "" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("env:", yaml);
        Assert.Contains("EMPTY_VAR:", yaml);
    }

    [Fact]
    public void CommandStep_WithParallelism_GeneratesParallelConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "dotnet test",
            Parallelism = 5
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("parallelism: 5", yaml);
    }

    [Fact]
    public void CommandStep_WithTimeout_GeneratesTimeoutConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Long running task",
            Command = "./long-task.sh",
            TimeoutInMinutes = 60
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("timeout_in_minutes: 60", yaml);
    }

    [Fact]
    public void CommandStep_WithDependsOn_GeneratesDependencyConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Key = "build",
            Label = "Build",
            Command = "make build"
        });
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            DependsOn = "build"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("depends_on: build", yaml);
    }

    [Fact]
    public void CommandStep_WithConditional_GeneratesIfConfig()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            If = "build.branch == 'main'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("if: build.branch == 'main'", yaml);
    }

    [Fact]
    public void CommandStep_WithStringPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[] { "docker#v5.0.0", "ecr#v2.0.0" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("ecr#v2.0.0", yaml);
    }

    [Fact]
    public void CommandStep_WithConfiguredPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[]
            {
                new Dictionary<string, object>
                {
                    ["docker#v5.0.0"] = new { image = "node:18" }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("image: node:18", yaml);
    }

    [Fact]
    public void CommandStep_WithMixedPlugins_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Plugins = new object[]
            {
                "artifacts#v1.0.0",
                new Dictionary<string, object>
                {
                    ["docker#v5.0.0"] = new { image = "node:18", environment = new[] { "NODE_ENV=production" } }
                },
                new Dictionary<string, object>
                {
                    ["ecr#v2.0.0"] = new { login = true }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("plugins:", yaml);
        Assert.Contains("artifacts#v1.0.0", yaml);
        Assert.Contains("docker#v5.0.0", yaml);
        Assert.Contains("image: node:18", yaml);
        Assert.Contains("ecr#v2.0.0", yaml);
        Assert.Contains("login: true", yaml);
    }

    [Fact]
    public void CommandStep_WithSecretsAsList_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            Secrets = new[] { "my-secret", "another-secret" }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("secrets:", yaml);
        Assert.Contains("my-secret", yaml);
        Assert.Contains("another-secret", yaml);
    }

    [Fact]
    public void CommandStep_WithSecretsAsMap_SerializesCorrectly()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Deploy",
            Command = "./deploy.sh",
            Secrets = new Dictionary<string, string>
            {
                ["MY_SECRET"] = "org/secret-name",
                ["API_KEY"] = "org/api-key"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("secrets:", yaml);
        Assert.Contains("MY_SECRET: org/secret-name", yaml);
        Assert.Contains("API_KEY: org/api-key", yaml);
    }

    [Fact]
    public void CommandStep_WithAutomaticRetry_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Retry = new Retry
            {
                Automatic = new AutomaticRetry { ExitStatus = -1, Limit = 3 }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("retry:", yaml);
        Assert.Contains("automatic:", yaml);
        Assert.Contains("exit_status: -1", yaml);
        Assert.Contains("limit: 3", yaml);
    }

    [Fact]
    public void CommandStep_WithManualRetry_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Retry = new Retry
            {
                Manual = new ManualRetry
                {
                    Allowed = false,
                    PermitOnPassed = true,
                    Reason = "Requires approval"
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("retry:", yaml);
        Assert.Contains("manual:", yaml);
        Assert.Contains("allowed: false", yaml);
        Assert.Contains("permit_on_passed: true", yaml);
        Assert.Contains("reason: Requires approval", yaml);
    }

    [Fact]
    public void CommandStep_WithCache_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "npm install",
            Cache = new Cache
            {
                Name = "node-modules",
                Paths = new List<string> { "node_modules", ".cache" },
                Size = "5g"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("cache:", yaml);
        Assert.Contains("name: node-modules", yaml);
        Assert.Contains("node_modules", yaml);
        Assert.Contains(".cache", yaml);
        Assert.Contains("size: 5g", yaml);
    }

    [Fact]
    public void CommandStep_WithMatrix_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Matrix = new Matrix
            {
                Setup = new Dictionary<string, List<string>>
                {
                    ["os"] = new() { "linux", "windows" },
                    ["arch"] = new() { "amd64", "arm64" }
                },
                Adjustments = new List<MatrixAdjustment>
                {
                    new()
                    {
                        With = new Dictionary<string, string>
                        {
                            ["os"] = "windows",
                            ["arch"] = "arm64"
                        },
                        Skip = true
                    }
                }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("matrix:", yaml);
        Assert.Contains("setup:", yaml);
        Assert.Contains("os:", yaml);
        Assert.Contains("linux", yaml);
        Assert.Contains("windows", yaml);
        Assert.Contains("adjustments:", yaml);
        Assert.Contains("skip: true", yaml);
    }

    [Fact]
    public void CommandStep_WithSignature_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Signed step",
            Command = "echo hello",
            Signature = new Signature
            {
                Algorithm = "HS512",
                SignedFields = new List<string> { "command", "env" },
                Value = "abc123"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("signature:", yaml);
        Assert.Contains("algorithm: HS512", yaml);
        Assert.Contains("signed_fields:", yaml);
        Assert.Contains("value: abc123", yaml);
    }

    [Fact]
    public void CommandStep_WithAutomaticAndManualRetry_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Retry = new Retry
            {
                Automatic = new AutomaticRetry { ExitStatus = 1, Limit = 2 },
                Manual = new ManualRetry { Allowed = true, PermitOnPassed = false }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("retry:", yaml);
        Assert.Contains("automatic:", yaml);
        Assert.Contains("exit_status: 1", yaml);
        Assert.Contains("manual:", yaml);
        Assert.Contains("allowed: true", yaml);
        Assert.Contains("permit_on_passed: false", yaml);
    }

    [Fact]
    public void CommandStep_WithAutomaticRetryWildcardExitStatus_GeneratesCorrectYaml()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Test",
            Command = "make test",
            Retry = new Retry
            {
                Automatic = new AutomaticRetry { ExitStatus = "*", Limit = 2 }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("automatic:", yaml);
        Assert.Contains("exit_status: '*'", yaml);
        Assert.Contains("limit: 2", yaml);
    }
}
