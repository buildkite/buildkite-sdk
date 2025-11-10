package main

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/schema"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

func generateGoTypes(
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	for _, name := range generator.Definitions.Keys() {
		val, _ := generator.Definitions.Get(name)
		prop := val.(schema.PropertyDefinition)

		property, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		fileName := fmt.Sprintf("%s/%s.go", outDir, utils.CamelCaseToSnakeCase(name))
		contents, err := property.Go()
		if err != nil {
			return fmt.Errorf("generating files contents for [%s]: %v", fileName, err)
		}

		file := utils.NewGoFile(
			"buildkite",
			fileName,
			[]string{},
			utils.NewCodeBlock(
				contents,
			),
		)

		err = file.Write()
		if err != nil {
			return fmt.Errorf("writing file [%s]: %v", fileName, err)
		}
	}

	pipelineSchemaString, err := generator.GeneratePipelineSchema()
	if err != nil {
		return err
	}

	pipelineFileName := fmt.Sprintf("%s/pipeline.go", outDir)
	file := utils.NewGoFile(
		"buildkite",
		pipelineFileName,
		[]string{},
		utils.NewCodeBlock(
			pipelineSchemaString,
		),
	)

	err = file.Write()
	if err != nil {
		return fmt.Errorf("writing file [%s]: %v", pipelineFileName, err)
	}

	return nil
}
