package main

import (
	"encoding/json"
	"fmt"
	"os"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

func run() error {
	// Create a new Buildkite Pipeline
	pipeline := bk.NewStepBuilder().
		AddCommand(&bk.Command{
			Commands: []string{
				"echo \"Hello World!\"",
			},
		})

	// Get the branch name of the current build
	branchName := bk.Environment.BUILDKITE_BRANCH()

	// Print out what branch we are on.
	if branchName == "main" {
		pipeline.AddCommand(&bk.Command{
			Commands: []string{
				`echo "I am on the main branch"`,
			},
		})
	} else {
		pipeline.AddCommand(&bk.Command{
			Commands: []string{
				fmt.Sprintf(`echo "I am on the %s branch"`, branchName),
			},
		})
	}

	str, err := json.Marshal(pipeline)
	if err != nil {
		return err
	}

	return os.WriteFile("pipeline.json", str, os.ModePerm)
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
