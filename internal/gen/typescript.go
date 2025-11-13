package main

import (
	"fmt"
	"path"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
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
		val, _ := generator.Definitions.Get(name)
		prop := val.(schema.PropertyDefinition)

		property, _, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		contents, err := property.TypeScript()
		if err != nil {
			return fmt.Errorf("generating files contents for [%s]", name)
		}

		codeBlock.AddLines(contents, "")
	}

	// Create an interface for the properties
	pipelineInterface := typescript.NewTypeScriptInterface("BuildkitePipeline", "", false)
	for _, name := range generator.Properties.Keys() {
		val, _ := generator.Properties.Get(name)
		prop := val.(schema.SchemaProperty)

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
