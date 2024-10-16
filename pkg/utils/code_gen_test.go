package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeBlock(t *testing.T) {
	t.Run("should display a code block", func(t *testing.T) {
		block := CodeGen.NewCodeBlock()
		assert.Equal(t, 0, len(block))
		assert.Equal(t, "", block.Display())
	})

	t.Run("should display a code block with an existing line", func(t *testing.T) {
		block := CodeGen.NewCodeBlock("one")
		assert.Equal(t, 1, len(block))
		assert.Equal(t, "one", block.Display())
	})

	t.Run("should display a code block with multiple existing lines", func(t *testing.T) {
		block := CodeGen.NewCodeBlock("one", "two", "three")
		assert.Equal(t, 3, len(block))
		assert.Equal(t, "one\ntwo\nthree", block.Display())
	})

	t.Run("should display a code block with indent", func(t *testing.T) {
		block := CodeGen.NewCodeBlock("one", "two", "three")
		assert.Equal(t, 3, len(block))
		assert.Equal(t, "  one\n  two\n  three", block.DisplayIndent(2))
	})
}

func TestNewCodeComment(t *testing.T) {
	t.Run("should generate a code comment", func(t *testing.T) {
		comment := CodeGen.Comment.newComment("//", "cool comment", 0)
		assert.Equal(t, "// cool comment", comment)
	})

	t.Run("should generate a code comment with multiple lines", func(t *testing.T) {
		comment := CodeGen.Comment.newComment("//", "cool comment\nanother line", 0)
		assert.Equal(t, "// cool comment\n// another line", comment)
	})

	t.Run("should generate a code comment with indent", func(t *testing.T) {
		comment := CodeGen.Comment.newComment("//", "cool comment", 2)
		assert.Equal(t, "  // cool comment", comment)
	})

	t.Run("should generate a code comment with multiple lines with indent", func(t *testing.T) {
		comment := CodeGen.Comment.newComment("//", "cool comment\nanother line", 2)
		assert.Equal(t, "  // cool comment\n  // another line", comment)
	})

	t.Run("should generate a code comment for typescript", func(t *testing.T) {
		comment := CodeGen.Comment.TypeScript("typescript", 0)
		assert.Equal(t, "// typescript", comment)
	})

	t.Run("should generate a code comment for go", func(t *testing.T) {
		comment := CodeGen.Comment.TypeScript("go", 0)
		assert.Equal(t, "// go", comment)
	})
}
