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
	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func main() {
	pipeline := buildkite.Pipeline{}

	pipeline.AddStep(buildkite.CommandStep{
		Command: &buildkite.CommandStepCommand{
			String: buildkite.Value("echo 'Hello, world!"),
		},
	})

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
