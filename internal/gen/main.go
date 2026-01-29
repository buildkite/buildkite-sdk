package main

import (
	"fmt"
	"os"
	"path"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
)

func generateTypes(outDir, language string) error {
	pipelineSchema, err := schema.ReadSchema()
	if err != nil {
		return fmt.Errorf("reading pipeline schema: %v", err)
	}

	generator := types.NewPipelineSchemaGenerator(pipelineSchema)
	switch language {
	case "ts":
		return generateTypeScriptTypes(generator, outDir)
	case "py":
		return generatePythonTypes(generator, outDir)
	case "go":
		return generateGoTypes(generator, outDir)
	case "csharp":
		return generateCSharpTypes(generator, outDir)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		panic("incorrect amount of arguments provided")
	}

	language := os.Args[1]
	outPath := os.Args[2]

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outDir := path.Join(dir, outPath)
	err = generateTypes(
		outDir,
		language,
	)
	if err != nil {
		panic(err)
	}
}
