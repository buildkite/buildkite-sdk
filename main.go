package main

import (
	"fmt"
	"os"

	"github.com/buildkite/pipeline-sdk/pkg/codegen"
	"github.com/buildkite/pipeline-sdk/pkg/schema"
)

func main() {
	err := codegen.GenerateSDKs(schema.PipelinesSchema)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
