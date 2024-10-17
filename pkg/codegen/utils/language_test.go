package code_gen_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	t.Run("typescript", func(t *testing.T) {
		ts := Languages.TypeScript()

		t.Run("should return the name", func(t *testing.T) {
			assert.Equal(t, "typescript", ts.Name())
		})

		t.Run("should create a comment", func(t *testing.T) {
			assert.Equal(t, "// comment", ts.NewComment("comment"))
		})

		t.Run("should create an indented comment", func(t *testing.T) {
			assert.Equal(t, "  // comment", ts.NewCommentWithIndent("comment", 2))
		})

		t.Run("should render imports", func(t *testing.T) {
			imports := []codeGenImport{
				{pkgName: "fs", identifier: "fs"},
			}
			result := ts.RenderImports(imports)
			assert.Equal(t, "import * as fs from \"fs\";", result)
		})
	})

	t.Run("go", func(t *testing.T) {
		golang := Languages.Go()

		t.Run("should return the name", func(t *testing.T) {
			assert.Equal(t, "go", golang.Name())
		})

		t.Run("should create a comment", func(t *testing.T) {
			assert.Equal(t, "// comment", golang.NewComment("comment"))
		})

		t.Run("should create an indented comment", func(t *testing.T) {
			assert.Equal(t, "  // comment", golang.NewCommentWithIndent("comment", 2))
		})

		t.Run("should render imports", func(t *testing.T) {
			imports := []codeGenImport{
				{pkgName: "fmt"},
			}
			result := golang.RenderImports(imports)
			assert.Equal(t, "import (\n    \"fmt\"\n)", result)
		})
	})
}
