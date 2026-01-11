package main

import (
	"fmt"

	"github.com/buildkite/buildkite-sdk/internal/gen/csharp"
	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
)

func generateCSharpTypes(
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	// Track written types and their dependencies for ordering
	writtenTypes := make(map[string]bool)
	allTypes := make(map[string]types.Value)
	allDeps := make(map[string][]string)

	// First pass: convert all definitions to values
	for _, name := range generator.Definitions.Keys() {
		prop, err := generator.Definitions.Get(name)
		if err != nil {
			return err
		}

		value, deps, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to value: %v", err)
		}

		allTypes[name] = value
		allDeps[name] = deps
	}

	// Second pass: write types in dependency order
	var writeType func(name string) error
	writeType = func(name string) error {
		if writtenTypes[name] {
			return nil
		}

		// Write dependencies first
		for _, dep := range allDeps[name] {
			if err := writeType(dep); err != nil {
				return err
			}
		}

		value, exists := allTypes[name]
		if !exists {
			return nil // Reference to external type
		}

		contents, err := value.CSharp()
		if err != nil {
			return fmt.Errorf("generating C# for [%s]: %v", name, err)
		}

		if contents == "" {
			writtenTypes[name] = true
			return nil
		}

		fileName := fmt.Sprintf("%s/%s.cs", outDir, utils.ToTitleCase(name))
		file := csharp.NewCSharpFile(
			"Buildkite.Sdk.Schema",
			fileName,
			[]string{
				"System.Collections.Generic",
				"System.Text.Json.Serialization",
			},
			contents,
		)

		if err := file.Write(); err != nil {
			return fmt.Errorf("writing file [%s]: %v", fileName, err)
		}

		writtenTypes[name] = true
		return nil
	}

	// Write all types
	for _, name := range generator.Definitions.Keys() {
		if err := writeType(name); err != nil {
			return err
		}
	}

	// Generate the main BuildkitePipeline class
	pipelineContents, err := generator.GenerateCSharpPipelineSchema()
	if err != nil {
		return fmt.Errorf("generating pipeline schema: %v", err)
	}

	pipelineFile := csharp.NewCSharpFile(
		"Buildkite.Sdk.Schema",
		fmt.Sprintf("%s/BuildkitePipeline.cs", outDir),
		[]string{
			"System.Collections.Generic",
			"System.Text.Json.Serialization",
		},
		pipelineContents,
	)

	if err := pipelineFile.Write(); err != nil {
		return fmt.Errorf("writing pipeline file: %v", err)
	}

	// Generate interfaces file
	interfacesContents := csharp.GenerateInterfaces()
	interfacesFile := csharp.NewCSharpFile(
		"Buildkite.Sdk.Schema",
		fmt.Sprintf("%s/Interfaces.cs", outDir),
		[]string{
			"System.Collections.Generic",
			"System.Text.Json.Serialization",
		},
		interfacesContents,
	)

	if err := interfacesFile.Write(); err != nil {
		return fmt.Errorf("writing interfaces file: %v", err)
	}

	return nil
}
