package main

import (
	"fmt"
	"path"

	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/typescript"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

func generateTypeScriptTypes(
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	// Create a new code block.
	codeBlock := utils.NewCodeBlock()

	// Generate the type for each definition in the schema.
	for _, name := range generator.Definitions.Keys() {
		prop, err := generator.Definitions.Get(name)
		if err != nil {
			return fmt.Errorf("getting definition: %v", err)
		}

		property, _, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		codeBlock.AddLines(property.TypeScript(), "")
	}

	// Create an interface for the properties
	pipelineInterface := typescript.NewTypeScriptInterface("BuildkitePipeline", "", false)
	for _, name := range generator.Properties.Keys() {
		prop, err := generator.Properties.Get(name)
		if err != nil {
			return fmt.Errorf("getting property: %v", err)
		}

		structType := utils.CamelCaseToTitleCase(prop.Ref.Name())
		pipelineInterface.AddItem(name, structType, "", false)
	}

	codeBlock.AddLines(
		pipelineInterface.Write(),
	)

	// Write out the schema file.
	file := typescript.NewTypeScriptFile(path.Join(outDir, "schema.ts"), codeBlock.String())
	err := file.Write()
	if err != nil {
		return fmt.Errorf("writing ts schema file: %v", err)
	}

	return nil
}
