# Buildkite Pipeline SDK

This repo contains the code to generate language SDKs for Buildkite Pipelines.

## Build and Install

To build and install the SDKs, you just need to run the build and install script:

```
$ ./scripts/build_and_install.sh
```

## TypeScript Example

```typescript
import * as bk from "buildkite-pipline-sdk";

let pipeline = bk.stepBuilder
    .addCommandStep({
        commands: [ "echo \"Hello World!\"" ],
    });

const branchName = bk.environment.branch();
if (branchName === "main") {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "I am on the main branch"` ],
    })
} else {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "This isn't the main its the ${branchName} branch"` ],
    })
}

pipeline.write();
```

## Go Example

```go
package main

import (
	"fmt"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func run() error {
	// Create a new Buildkite Pipeline
	pipeline := bk.NewStepBuilder().
		AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				"echo \"Hello World!\"",
			},
		})

	// Get the branch name of the current build
	branchName := bk.Environment.Branch()

	// Print out what branch we are on.
	if branchName == "main" {
		pipeline.AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				`echo "I am on the main branch"`,
			},
		})
	} else {
		pipeline.AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				fmt.Sprintf(`echo "I am on the %s branch"`, branchName),
			},
		})
	}

	return pipeline.Print()
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
```
