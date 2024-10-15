package main

import (
	"fmt"
	"os"

	"github.com/buildkite/pipeline-sdk/pkg/codegen"
	go_code_gen "github.com/buildkite/pipeline-sdk/pkg/codegen/go"
	typescript_code_gen "github.com/buildkite/pipeline-sdk/pkg/codegen/typescript"
	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

func main() {
	gen := codegen.NewGenerator(utils.FS, []codegen.LanguageTarget{
		typescript_code_gen.TypeScriptSDK{},
		go_code_gen.GoSDK{},
	})

	err := gen.GenerateSDKs(schema.PipelinesSchema)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
