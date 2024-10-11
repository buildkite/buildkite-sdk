package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
)

type TypeScriptSDK struct{}

func (TypeScriptSDK) FolderName() string {
	return "typescript"
}

func (TypeScriptSDK) Files(pipelineSchema schema.PipelineSchema) (map[string]string, error) {
	tsTypes, err := NewTypesFile(pipelineSchema.Types, pipelineSchema.Steps)
	if err != nil {
		return nil, fmt.Errorf("generating typds.d.ts: %v", err)
	}

	return map[string]string{
		"environment.ts": newEnvironmentFile(pipelineSchema.Environment),
		"index.ts":       newIndexFile(),
		"package.json":   newPackageJSONFile(),
		"stepBuilder.ts": newStepBuilderFile(pipelineSchema),
		"types.ts":       tsTypes,
	}, nil
}
