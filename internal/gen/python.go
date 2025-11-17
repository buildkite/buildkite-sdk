package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/buildkite/buildkite-sdk/internal/gen/types"
	"github.com/buildkite/buildkite-sdk/internal/gen/utils"
	"github.com/iancoleman/orderedmap"
)

type pythonProperty struct {
	Property     types.Value
	Dependencies []string
	Written      bool
}

type pythonTypes struct {
	Properties *orderedmap.OrderedMap
}

func (p *pythonTypes) Add(name string, prop types.Value, dependencies []string) {
	p.Properties.Set(name, pythonProperty{
		Property:     prop,
		Dependencies: dependencies,
	})
}

func (p *pythonTypes) Get(name string) pythonProperty {
	prop, _ := p.Properties.Get(name)
	return prop.(pythonProperty)
}

func (p *pythonTypes) Generate(name string) (string, error) {
	prop := p.Get(name)
	if prop.Written {
		fmt.Printf("%s has already been written\n", name)
		return "", nil
	}

	codeBlock := utils.NewCodeBlock()

	fmt.Printf("Generating %s: [%s]\n", name, strings.Join(prop.Dependencies, ","))
	for _, dependency := range prop.Dependencies {
		if dependency == "" {
			continue
		}

		fmt.Printf("Generating dependency [%s]: %s\n", name, dependency)
		dependencyCode, err := p.Generate(dependency)
		if err != nil {
			return "", fmt.Errorf("getting dependency for [%s]: %s", name, dependency)
		}

		if dependencyCode != "" {
			codeBlock.AddLines(dependencyCode)
		}
	}

	contents, err := prop.Property.Python()
	if err != nil {
		return "", fmt.Errorf("generating files contents for [%s]", name)
	}

	p.Properties.Set(name, pythonProperty{
		Property:     prop.Property,
		Dependencies: prop.Dependencies,
		Written:      true,
	})

	codeBlock.AddLines(contents, "")
	return codeBlock.String(), nil
}

func generatePythonTypes(
	generator types.PipelineSchemaGenerator,
	outDir string,
) error {
	codeBlock := utils.NewCodeBlock()
	typeNames := generator.Definitions.Keys()
	types := &pythonTypes{
		Properties: orderedmap.New(),
	}

	for _, name := range typeNames {
		prop, err := generator.Definitions.Get(name)
		if err != nil {
			return err
		}

		property, dependencies, err := generator.PropertyDefinitionToValue(name, prop)
		if err != nil {
			return fmt.Errorf("converting property definition to a value: %v", err)
		}

		types.Add(name, property, dependencies)
	}

	for _, key := range types.Properties.Keys() {
		lines, err := types.Generate(key)
		if err != nil {
			return fmt.Errorf("generating types: %v", err)
		}

		if lines != "" {
			codeBlock.AddLines(lines)
		}
	}

	file := utils.NewPythonFile(path.Join(outDir, "schema.py"), codeBlock)
	err := file.Write()
	if err != nil {
		return fmt.Errorf("writing python schema file: %v", err)
	}
	return nil
}
