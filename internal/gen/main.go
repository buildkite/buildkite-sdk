package main

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/buildkite/pipeline-sdk/internal/gen/schema"
	"github.com/buildkite/pipeline-sdk/internal/gen/types"
	"github.com/buildkite/pipeline-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
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

		codeBlock.AddLines(contents)
	}

	file := utils.NewPythonFile(path.Join(outDir, "schema.py"), codeBlock)
	err := file.Write()
	if err != nil {
		return fmt.Errorf("writing python schema file: %v", err)
	}

	return nil
}

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

	// Old
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
