package codegen

import (
	"fmt"
	"os"
	"path"

	typescript_code_gen "github.com/buildkite/pipeline-sdk/pkg/codegen/typescript"
	"github.com/buildkite/pipeline-sdk/pkg/schema"
)

const SDK_FOLDER = "sdk"

type LanguageTarget interface {
	FolderName() string
	Files(pipelineSchema schema.PipelineSchema) (map[string]string, error)
}

var targets = []LanguageTarget{
	typescript_code_gen.TypeScriptSDK{},
	// go_code_gen.GoSDK{},
}

func GenerateSDKs(pipelineSchema schema.PipelineSchema) error {
	err := os.Mkdir(SDK_FOLDER, os.ModePerm)
	if err != nil {
		return fmt.Errorf("creating sdk folder: %v", err)
	}

	for _, target := range targets {
		targetFolder := path.Join(SDK_FOLDER, target.FolderName())
		err = os.Mkdir(targetFolder, os.ModePerm)
		if err != nil {
			return err
		}

		files, err := target.Files(pipelineSchema)
		if err != nil {
			return fmt.Errorf("generating files: %v", err)
		}

		for name, contents := range files {
			out := path.Join(targetFolder, name)
			err = os.WriteFile(out, []byte(contents), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
