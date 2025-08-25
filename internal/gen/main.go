package main

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
	"github.com/buildkite/pipeline-sdk/internal/gen/types"
	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

func main() {
	pipelineSchema, err := schema.ReadSchema()
	if err != nil {
		panic(fmt.Errorf("reading pipeline schema: %v", err))
	}

	generator := types.PipelineSchemaGenerator{
		Definitions: pipelineSchema.Definitions,
	}

	for name, prop := range pipelineSchema.Definitions {
		property, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			panic(fmt.Errorf("converting property definition to a value: %v", err))
		}

		fileName := fmt.Sprintf("definitions/%s.go", utils.CamelCaseToSnakeCase(name))

		contents, err := property.Go()
		if err != nil {
			panic(fmt.Errorf("generating go code for [%s]: %v", name, err))
		}

		file := utils.NewGoFile(
			"definitions",
			fileName,
			[]string{},
			utils.NewCodeBlock(
				contents,
			),
		)

		err = file.Write()
		if err != nil {
			panic(fmt.Errorf("writing file [%s]: %v", fileName, err))
		}
	}
}
