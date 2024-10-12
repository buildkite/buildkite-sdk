package go_code_gen

import "github.com/buildkite/pipeline-sdk/pkg/schema"

type GoSDK struct{}

func (GoSDK) FolderName() string {
	return "go"
}

func (GoSDK) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"enviornment.go":  newEnvironmentFile(pipelineSchema.Environment),
		"root.go":         newRootFile(),
		"step_builder.go": newStepBuilderFile(pipelineSchema),
		"types.go":        newTypesFile(pipelineSchema.Types, pipelineSchema.Steps),
		"utils.go":        newUtilsFiles(),
	}
}
