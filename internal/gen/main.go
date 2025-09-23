package main

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/iancoleman/orderedmap"
)

func generateTypes(outDir, language string) error {
	pipelineSchema, err := schema.ReadSchema()
	if err != nil {
		return fmt.Errorf("reading pipeline schema: %v", err)
	}

	definitions := orderedmap.New()
	for key, prop := range pipelineSchema.Definitions {
		definitions.Set(key, prop)
	}
	definitions.SortKeys(sort.Strings)

	properties := orderedmap.New()
	for key, prop := range pipelineSchema.Properties {
		properties.Set(key, prop)
	}
	properties.SortKeys(sort.Strings)

	generator := types.PipelineSchemaGenerator{
		Definitions: definitions,
		Properties:  properties,
	}

	if language == "ts" {
		return generateTypeScriptTypes(generator, outDir)
	}

	if language == "py" {
		return generatePythonTypes(generator, outDir)
	}

	if language == "go" {
		return generateGoTypes(generator, outDir)
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
