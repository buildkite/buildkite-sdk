package typescript_code_gen

import (
	"github.com/buildkite/pipeline-sdk/pkg/schema"
)

type TypeScriptSDK struct{}

func (TypeScriptSDK) FolderName() string {
	return "typescript"
}

func (TypeScriptSDK) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"environment.ts": newEnvironmentFile(pipelineSchema.Environment).Render(),
		"index.ts":       newIndexFile(),
		"package.json":   newPackageJSONFile(),
		"stepBuilder.ts": newStepBuilderFile(pipelineSchema),
		"types.ts":       newTypesFile(pipelineSchema.Types, pipelineSchema.Steps),
	}
}
