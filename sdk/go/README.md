# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/pipeline-sdk)

A Ruby SDK for [Buildkite](https://buildkite.com)! 🪁

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
	command := "echo 'Hello, world!"

	pipeline.AddCommandStep(buildkite.CommandStep{
		Command: &buildkite.CommandUnion{
			String: &command,
		},
	})

	fmt.Println(pipeline.ToJSON())
	fmt.Println(pipeline.ToYAML())
}
```
