package main

import (
	"fmt"
	"path"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

func generatePythonTypes(
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

		contents, err := property.Python()
		if err != nil {
			return fmt.Errorf("generating files contents for [%s]", name)
		}

		codeBlock.AddLines(contents, "")
	}

	file := utils.NewPythonFile(path.Join(outDir, "schema.py"), codeBlock)
	err := file.Write()
	if err != nil {
		return fmt.Errorf("writing python schema file: %v", err)
	}

	return nil
}
