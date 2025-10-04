# Buildkite SDK for .NET

A .NET SDK for [Buildkite](https://buildkite.com)! 🪁

## Installation

Add the package to your project:

```bash
dotnet add package Buildkite.Sdk
```

## Usage

```csharp
using Buildkite.Sdk;

var pipeline = new Pipeline();

pipeline.AddStep(new CommandStep
{
    Label = "Test",
    Command = "dotnet test"
});

pipeline.AddStep(new CommandStep
{
    Label = "Build",
    Command = "dotnet build --configuration Release"
});

// Output as YAML
Console.WriteLine(pipeline.ToYaml());

// Output as JSON
Console.WriteLine(pipeline.ToJson());
```

## Documentation

Visit the [Buildkite SDK documentation](https://buildkite.com/docs/pipelines/configure/dynamic-pipelines/sdk) for more details.
