package main

import (
	"log"
	"os"

	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func value[T any](val T) *T {
	return &val
}

func main() {
	pipeline := buildkite.NewPipeline()

	pipeline.AddStep(buildkite.CommandStep{
		Label: value("some-label"),
		Commands: &buildkite.CommandStepCommand{
			String: value("echo 'Hello, world!"),
		},
	})

	err := os.MkdirAll("../../out/apps/go", 0755) // 0755 sets permissions (read, write, execute for owner; read and execute for others)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	json, err := pipeline.ToJSON()
	if err != nil {
		log.Fatalf("Failed to serialize JSON: %v", err)
	}

	jsonFile, err := os.Create("../../out/apps/go/pipeline.json")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.WriteString(json)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	yaml, err := pipeline.ToYAML()
	if err != nil {
		log.Fatalf("Failed to serialize JSON: %v", err)
	}

	yamlFile, err := os.Create("../../out/apps/go/pipeline.yaml")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer yamlFile.Close()

	_, err = yamlFile.WriteString(yaml)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}
