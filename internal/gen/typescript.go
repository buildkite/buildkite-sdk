package main

import (
	"fmt"
	"path"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

func generateTypeScriptTypes(
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	codeBlock := utils.NewCodeBlock()

	for _, name := range generator.Definitions.Keys() {
		val, _ := generator.Definitions.Get(name)
		prop := val.(schema.PropertyDefinition)

		property, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		contents, err := property.TypeScript()
		if err != nil {
			return fmt.Errorf("generating files contents for [%s]", name)
		}

		codeBlock.AddLines(contents)
	}

	pipelineInterface := utils.NewTypeScriptInterface("BuildkitePipeline")
	for _, name := range generator.Properties.Keys() {
		val, _ := generator.Properties.Get(name)
		prop := val.(schema.SchemaProperty)

		structType := utils.CamelCaseToTitleCase(prop.Ref.Name())
		pipelineInterface.AddItem(name, structType, false)
	}

	pipelineString, err := pipelineInterface.Write()
	if err != nil {
		return fmt.Errorf("generating pipeline interface: %v", err)
	}
	codeBlock.AddLines(pipelineString)

	file := utils.NewTypeScriptFile(path.Join(outDir, "schema.ts"), nil, codeBlock)
	err = file.Write()
	if err != nil {
		return fmt.Errorf("writing ts schema file: %v", err)
	}

	return nil
}
