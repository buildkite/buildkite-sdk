# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/buildkite-sdk)

A Go SDK for [Buildkite](https://buildkite.com)! 🪁

## Usage

Install the package:

```bash
go get github.com/buildkite/buildkite-sdk/sdk/go
```

Use it in your program:

```go
package main

import (
	"fmt"
	"log"

	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func main() {
	pipeline := buildkite.Pipeline{}

	pipeline.AddStep(buildkite.CommandStep{
		Command: &buildkite.CommandStepCommand{
			String: buildkite.Value("echo 'Hello, world!'"),
		},
		If: buildkite.MustCondition(`build.branch == "main"`),
	})

	if err := buildkite.ValidateConditionals(pipeline); err != nil {
		log.Fatalf("Failed to validate conditionals: %v", err)
	}

	json, err := pipeline.ToJSON()
	if err != nil {
		log.Fatalf("Failed to serialize JSON: %v", err)
	}

	fmt.Println(json)

	yaml, err := pipeline.ToYAML()
	if err != nil {
		log.Fatalf("Failed to serialize YAML: %v", err)
	}

	fmt.Println(yaml)
}
```

`Condition` returns the same `*string` shape used by generated SDK `if` fields,
and `ValidateConditionals` walks a pipeline's step and notification conditions
before serialization:

```go
condition, err := buildkite.Condition(`build.branch == "main"`)
if err != nil {
	log.Fatalf("Invalid condition: %v", err)
}

step := buildkite.CommandStep{
	Command: &buildkite.CommandStepCommand{
		String: buildkite.Value("./deploy.sh"),
	},
	If: condition,
}
```
