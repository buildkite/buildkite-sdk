package main

import (
	"fmt"
	"os"
	"path"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
	"github.com/buildkite/pipeline-sdk/internal/gen/types"
	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
)

func generateTypeScriptTypes(
	pipelineSchema schema.PipelineSchema,
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	codeBlock := utils.NewCodeBlock()

	for name, prop := range pipelineSchema.Definitions {
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

	file := utils.NewTypeScriptFile(path.Join(outDir, "schema.ts"), nil, codeBlock)
	err := file.Write()
	if err != nil {
		return fmt.Errorf("writing ts schema file: %v", err)
	}

	return nil
}

func generateTypes(outDir, language string) error {
	pipelineSchema, err := schema.ReadSchema()
	if err != nil {
		return fmt.Errorf("reading pipeline schema: %v", err)
	}

	generator := types.PipelineSchemaGenerator{
		Definitions: pipelineSchema.Definitions,
		Properties:  pipelineSchema.Properties,
	}

	if language == "ts" {
		return generateTypeScriptTypes(pipelineSchema, generator, outDir)
	}

	for name, prop := range pipelineSchema.Definitions {
		property, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		var file utils.FileWriter
		var fileName string
		switch language {
		case "go":
			fileName := fmt.Sprintf("%s/%s.go", outDir, utils.CamelCaseToSnakeCase(name))
			contents, err := property.Go()
			if err != nil {
				return fmt.Errorf("generating files contents for [%s]", fileName)
			}
			file = utils.NewGoFile(
				"buildkite",
				fileName,
				[]string{},
				utils.NewCodeBlock(
					contents,
				),
			)
		case "ts":
			fileName := fmt.Sprintf("%s/%s.ts", outDir, name)
			contents, err := property.TypeScript()
			if err != nil {
				return fmt.Errorf("generating files contents for [%s]", fileName)
			}
			file = utils.NewTypeScriptFile(fileName, nil, utils.NewCodeBlock(contents))
		default:
			return fmt.Errorf("unsupported language provided [%s]", language)
		}

		err = file.Write()
		if err != nil {
			return fmt.Errorf("writing file [%s]: %v", fileName, err)
		}
	}

	// pipelineSchemaString, err := generator.GeneratePipelineSchema()
	// if err != nil {
	// 	return err
	// }

	// var pipelineFileName string
	// switch language {
	// case "go":
	// 	pipelineFileName = fmt.Sprintf("%s/pipeline.go", outDir)
	// default:
	// 	return fmt.Errorf("unsupported language provided [%s]", language)
	// }

	// file := utils.NewGoFile(
	// 	"buildkite",
	// 	pipelineFileName,
	// 	[]string{},
	// 	utils.NewCodeBlock(
	// 		pipelineSchemaString,
	// 	),
	// )

	// err = file.Write()
	// if err != nil {
	// 	return fmt.Errorf("writing file [%s]: %v", pipelineFileName, err)
	// }

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
