package codegen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type mockLanguageTarget struct{}

func (mockLanguageTarget) FolderName() string {
	return "mock"
}

func (mockLanguageTarget) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"mock-file.txt": "mock contents",
	}
}

func TestGenerator(t *testing.T) {

	t.Run("should generate the sdks", func(t *testing.T) {
		mockFS := utils.NewMockFS(t)
		mockFS.Mocks.AddMockItem("NewDirectory", []interface{}{"sdk"}, nil)
		mockFS.Mocks.AddMockItem("NewDirectory", []interface{}{"sdk/mock"}, nil)
		mockFS.Mocks.AddMockItem("NewFile", []interface{}{"sdk/mock/mock-file.txt", "mock contents"}, nil)

		gen := NewGenerator(mockFS, []LanguageTarget{
			mockLanguageTarget{},
		})

		err := gen.GenerateSDKs(schema.PipelinesSchema)
		assert.NoError(t, err)
	})
}
