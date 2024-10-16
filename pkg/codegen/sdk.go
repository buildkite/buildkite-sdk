package codegen

import (
	"fmt"
	"path"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

const SDK_FOLDER = "sdk"

type LanguageTarget interface {
	FolderName() string
	Files(pipelineSchema schema.PipelineSchema) map[string]string
}

type generator struct {
	FS      utils.FileSystem
	Targets []LanguageTarget
}

func (g generator) GenerateSDKs(pipelineSchema schema.PipelineSchema) error {
	err := g.FS.NewDirectory(SDK_FOLDER)
	if err != nil {
		return fmt.Errorf("creating sdk folder: %v", err)
	}

	for _, target := range g.Targets {
		targetFolder := path.Join(SDK_FOLDER, target.FolderName())
		err = g.FS.NewDirectory(targetFolder)
		if err != nil {
			return err
		}

		files := target.Files(pipelineSchema)
		for name, contents := range files {
			out := path.Join(targetFolder, name)
			err = g.FS.NewFile(out, contents)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewGenerator(fs utils.FileSystem, targets []LanguageTarget) generator {
	return generator{
		FS:      fs,
		Targets: targets,
	}
}
