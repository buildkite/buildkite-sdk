package main

import (
	"encoding/json"
	"fmt"
	"os"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
)

type dockerPluginArgs struct {
	Image       string   `json:"image"`
	Environment []string `json:"environment,omitempty"`
}

func runBranchBuild(pipeline *bk.StepBuilder) {
	pipeline.
		AddCommand(&bk.Command{
			Label: "Test",
			Commands: []string{
				"go test ./...",
			},
			Plugins: []map[string]interface{}{
				{
					"docker#v5.11.0": dockerPluginArgs{
						Image: "golang:1.22",
					},
				},
			},
		}).
		AddCommand(&bk.Command{
			Label: "Build and Install SDKs",
			Commands: []string{
				"./scripts/build_and_install.sh",
			},
			Plugins: []map[string]interface{}{
				{
					"docker#v5.11.0": dockerPluginArgs{
						Image: "golang:1.22",
					},
				},
			},
		})
}

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
				`echo "main build"`,
			},
		})
	} else {
		runBranchBuild(pipeline)
	}

	str, err := json.Marshal(pipeline)
	if err != nil {
		return err
	}

	fmt.Println(string(str))

	return os.WriteFile("pipeline.json", str, os.ModePerm)
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
