package code_gen_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeGenFile(t *testing.T) {
	t.Run("should create a new file", func(t *testing.T) {
		file := File.New(Languages.Go())
		assert.Contains(t, file.header.Display(), "This file is auto-generated. Do not edit.")
		assert.Equal(t, 0, len(file.imports))
		assert.NotNil(t, file.language)
		assert.Equal(t, 0, len(file.code))
	})

	t.Run("should append to the header", func(t *testing.T) {
		file := File.New(Languages.Go())
		file.AppendToHeader("package buildkite")
		assert.Contains(t, file.header.Display(), "package buildkite")
	})

	t.Run("should add an import", func(t *testing.T) {
		file := File.New(Languages.Go())
		file.AddImport("fmt", "fmt")
		assert.Equal(t, 1, len(file.imports))
	})

	t.Run("should render imports", func(t *testing.T) {
		file := File.New(Languages.Go())
		file.AddImport("fmt", "fmt")
		result := file.RenderImports()
		assert.Equal(t, "import (\n    \"fmt\"\n)", result)
	})

	t.Run("should append code", func(t *testing.T) {
		file := File.New(Languages.Go())
		file.AppendCode("var something = 42")
		assert.Equal(t, 1, len(file.code))
	})

	t.Run("should render a comment", func(t *testing.T) {
		file := File.New(Languages.Go())
		assert.Equal(t, "// comment", file.NewComment("comment"))
	})

	t.Run("should render a comment with indent", func(t *testing.T) {
		file := File.New(Languages.Go())
		assert.Equal(t, "  // comment", file.NewCommentWithIndent("comment", 2))
	})

	t.Run("should render a file to a string", func(t *testing.T) {
		file := File.New(Languages.Go())
		file.AppendToHeader("package buildkite")
		file.AddImport("fmt", "fmt")
		file.AppendCode("var something = 42")
		assert.Equal(t, "// This file is auto-generated. Do not edit.\npackage buildkite\nimport (\n    \"fmt\"\n)\nvar something = 42", file.Render())
	})
}
