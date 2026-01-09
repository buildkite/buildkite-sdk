using System;
using System.IO;
using Buildkite.Sdk;
using Buildkite.Sdk.Schema;

var pipeline = new Pipeline();

pipeline.AddStep(new CommandStep
{
    Label = "some-label",
    Command = "echo 'Hello, world!'"
});

Directory.CreateDirectory("../../out/apps/csharp");

var json = pipeline.ToJson();
File.WriteAllText("../../out/apps/csharp/pipeline.json", json);

var yaml = pipeline.ToYaml();
File.WriteAllText("../../out/apps/csharp/pipeline.yaml", yaml);
