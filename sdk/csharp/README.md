# Buildkite SDK for C#

[![NuGet](https://img.shields.io/nuget/v/Buildkite.Sdk.svg)](https://www.nuget.org/packages/Buildkite.Sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A C# SDK for [Buildkite](https://buildkite.com) that makes it easy to generate dynamic pipeline configurations programmatically.

## Installation

```bash
dotnet add package Buildkite.Sdk
```

## Quick Start

```csharp
using Buildkite.Sdk;
using Buildkite.Sdk.Schema;

var pipeline = new Pipeline();

pipeline.AddStep(new CommandStep
{
    Label = ":hammer: Build",
    Command = "dotnet build"
});

pipeline.AddStep(new WaitStep());

pipeline.AddStep(new CommandStep
{
    Label = ":test_tube: Test",
    Command = "dotnet test"
});

// Output as YAML for `buildkite-agent pipeline upload`
Console.WriteLine(pipeline.ToYaml());
```

## Usage

### Command Steps

```csharp
pipeline.AddStep(new CommandStep
{
    Label = ":dotnet: Build",
    Key = "build",
    Command = "dotnet build --configuration Release",
    Agents = new AgentsObject { ["queue"] = "linux" },
    TimeoutInMinutes = 30
});
```

### Block Steps

```csharp
pipeline.AddStep(new BlockStep
{
    Block = ":rocket: Deploy to Production?",
    Prompt = "Are you sure?"
});
```

### Input Steps

```csharp
pipeline.AddStep(new InputStep
{
    Input = ":writing_hand: Release Details",
    Key = "release-info",
    Prompt = "Provide release information",
    Fields = new List<Field>
    {
        new TextField
        {
            Key = "release-notes",
            Text = "Release Notes",
            Hint = "Markdown supported",
            Required = true
        },
        new SelectField
        {
            Key = "environment",
            Select = "Target Environment",
            Options = new List<SelectOption>
            {
                new SelectOption { Label = "Staging", Value = "staging" },
                new SelectOption { Label = "Production", Value = "production" }
            },
            Default = "staging"
        }
    }
});
```

### Wait Steps

```csharp
pipeline.AddStep(new WaitStep());
pipeline.AddStep(new WaitStep { ContinueOnFailure = true });
```

### Trigger Steps

```csharp
pipeline.AddStep(new TriggerStep
{
    Trigger = "deploy-pipeline",
    Build = new TriggerBuild { Branch = "main" }
});
```

### Group Steps

```csharp
pipeline.AddStep(new GroupStep
{
    Group = ":test_tube: Tests",
    Steps = new List<IGroupStep>
    {
        new CommandStep { Label = "Unit", Command = "dotnet test" },
        new CommandStep { Label = "Integration", Command = "dotnet test --filter Integration" }
    }
});
```

### Plugins

```csharp
// Simple plugin (no configuration)
pipeline.AddStep(new CommandStep
{
    Label = ":docker: Build",
    Command = "build.sh",
    Plugins = new object[]
    {
        "docker-login#v3.0.0"
    }
});

// Plugin with configuration
pipeline.AddStep(new CommandStep
{
    Label = ":docker: Run",
    Command = "dotnet test",
    Plugins = new object[]
    {
        new Dictionary<string, object>
        {
            ["docker#v5.11.0"] = new { image = "mcr.microsoft.com/dotnet/sdk:9.0" }
        }
    }
});

// Mixing simple and configured plugins
pipeline.AddStep(new CommandStep
{
    Label = ":rocket: Deploy",
    Command = "deploy.sh",
    Plugins = new object[]
    {
        "docker-login#v3.0.0",
        new Dictionary<string, object>
        {
            ["ecr#v2.9.0"] = new { login = true, account_ids = "123456789" }
        }
    }
});
```

### Retry

```csharp
pipeline.AddStep(new CommandStep
{
    Label = ":test_tube: Tests",
    Command = "dotnet test",
    Retry = new Retry
    {
        Automatic = new AutomaticRetry
        {
            ExitStatus = "*",
            Limit = 2
        },
        Manual = new ManualRetry
        {
            PermitOnPassed = true,
            Reason = "Re-run if flaky"
        }
    }
});
```

### Soft Fail and Skip

```csharp
// Soft fail on any exit status
pipeline.AddStep(new CommandStep
{
    Label = ":lint: Optional Lint",
    Command = "dotnet format --verify-no-changes",
    SoftFail = true
});

// Soft fail on specific exit statuses
pipeline.AddStep(new CommandStep
{
    Label = ":warning: Flaky Test",
    Command = "dotnet test --filter Flaky",
    SoftFail = SoftFail.FromConditions(
        new SoftFailCondition { ExitStatus = 1 },
        new SoftFailCondition { ExitStatus = SoftFailExitStatus.FromWildcard() }
    )
});

// Skip with a reason
pipeline.AddStep(new CommandStep
{
    Label = ":no_entry: Disabled",
    Command = "echo skipped",
    Skip = "Temporarily disabled pending fix"
});
```

### Matrix

```csharp
// Simple matrix
pipeline.AddStep(new CommandStep
{
    Label = "Test {{matrix.runtime}}",
    Command = "dotnet test",
    Matrix = new Matrix
    {
        Setup = new Dictionary<string, List<string>>
        {
            ["runtime"] = new List<string> { "net8.0", "net9.0" },
            ["os"] = new List<string> { "linux", "windows" }
        }
    }
});

// Matrix with adjustments
pipeline.AddStep(new CommandStep
{
    Label = "Test {{matrix.runtime}} on {{matrix.os}}",
    Command = "dotnet test",
    Matrix = new Matrix
    {
        Setup = new Dictionary<string, List<string>>
        {
            ["runtime"] = new List<string> { "net8.0", "net9.0" },
            ["os"] = new List<string> { "linux", "windows" }
        },
        Adjustments = new List<MatrixAdjustment>
        {
            new MatrixAdjustment
            {
                With = new Dictionary<string, string>
                {
                    ["runtime"] = "net8.0",
                    ["os"] = "windows"
                },
                SoftFail = true
            }
        }
    }
});
```

### Notifications

```csharp
// Pipeline-level notifications
pipeline.AddNotify(new SlackNotification
{
    Slack = new SlackConfig
    {
        Channels = new List<string> { "#builds" },
        Message = "Build finished"
    },
    If = "build.state == 'failed'"
});

pipeline.AddNotify(new EmailNotification
{
    Email = "team@example.com",
    If = "build.state == 'failed'"
});

// Step-level notifications
pipeline.AddStep(new CommandStep
{
    Label = ":rocket: Deploy",
    Command = "deploy.sh",
    Notify = new List<INotification>
    {
        new SlackNotification { Slack = "#deploys" }
    }
});
```

### DependsOn

```csharp
// Single dependency
pipeline.AddStep(new CommandStep
{
    Label = ":test_tube: Test",
    Command = "dotnet test",
    DependsOn = "build"
});

// Multiple dependencies
pipeline.AddStep(new CommandStep
{
    Label = ":rocket: Deploy",
    Command = "deploy.sh",
    DependsOn = new string[] { "build", "test" }
});

// Allow a dependency to fail
pipeline.AddStep(new CommandStep
{
    Label = ":page_facing_up: Report",
    Command = "generate-report.sh",
    DependsOn = DependsOn.FromDependencies(
        new Dependency { Step = "test", AllowFailure = true }
    )
});
```

### If and IfChanged Conditions

```csharp
// Conditional step using a boolean expression
pipeline.AddStep(new CommandStep
{
    Label = ":rocket: Deploy",
    Command = "deploy.sh",
    If = "build.branch == 'main'"
});

// Run only when specific files change
pipeline.AddStep(new CommandStep
{
    Label = ":dotnet: Build",
    Command = "dotnet build",
    IfChanged = new[] { "src/**", "*.csproj" }
});
```

### Concurrency

```csharp
pipeline.AddStep(new CommandStep
{
    Label = ":rocket: Deploy",
    Command = "deploy.sh",
    Concurrency = 1,
    ConcurrencyGroup = "deploy/production",
    ConcurrencyMethod = "eager"
});
```

### Secrets

```csharp
// List of secret names
pipeline.AddStep(new CommandStep
{
    Label = ":closed_lock_with_key: Deploy",
    Command = "deploy.sh",
    Secrets = new[] { "npm-token", "deploy-key" }
});

// Map environment variables to secrets
pipeline.AddStep(new CommandStep
{
    Label = ":closed_lock_with_key: Publish",
    Command = "publish.sh",
    Secrets = new Dictionary<string, string>
    {
        ["NPM_TOKEN"] = "org/npm-token",
        ["DEPLOY_KEY"] = "org/deploy-key"
    }
});
```

### Cache

```csharp
pipeline.AddStep(new CommandStep
{
    Label = ":dotnet: Build",
    Command = "dotnet build",
    Cache = new Cache
    {
        Name = "nuget-packages",
        Paths = new List<string> { "~/.nuget/packages" },
        Size = "5g"
    }
});
```

### Artifact Paths

```csharp
pipeline.AddStep(new CommandStep
{
    Label = ":package: Package",
    Command = "dotnet publish -o ./out",
    ArtifactPaths = new[] { "out/**/*", "logs/*.txt" }
});
```

### Agents

```csharp
// Object format (key-value pairs)
pipeline.AddStep(new CommandStep
{
    Label = ":dotnet: Build",
    Command = "dotnet build",
    Agents = new AgentsObject { ["queue"] = "linux", ["dotnet"] = "true" }
});

// List format (key=value strings)
pipeline.AddStep(new CommandStep
{
    Label = ":dotnet: Build",
    Command = "dotnet build",
    Agents = new AgentsList { "queue=linux", "dotnet=true" }
});
```

### Environment Variables

```csharp
using Buildkite.Sdk;

var branch = EnvironmentVariable.Branch;
var commit = EnvironmentVariable.Commit;
var buildNumber = EnvironmentVariable.BuildNumber;
```

## Output Formats

```csharp
string yaml = pipeline.ToYaml();  // For buildkite-agent pipeline upload
string json = pipeline.ToJson();  // JSON format
```

## Development

```bash
dotnet build
dotnet test
dotnet pack -c Release
```

## License

MIT License - see [LICENSE](../../LICENSE) for details.
