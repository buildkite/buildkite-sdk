package go_code_gen

import "github.com/buildkite/pipeline-sdk/pkg/schema"

type GoSDK struct{}

func (GoSDK) FolderName() string {
	return "go"
}

func (GoSDK) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"environment.go":  newEnvironmentFile(pipelineSchema.Environment),
		"step_builder.go": newStepBuilderFile(pipelineSchema.Steps),
		"types.go":        newTypesFile(pipelineSchema.Types, pipelineSchema.Steps),
		"utils.go":        newUtilsFiles(),
		"version.go":      newVersionFile(pipelineSchema.Version),
	}
}
