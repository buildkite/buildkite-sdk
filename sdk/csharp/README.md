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
