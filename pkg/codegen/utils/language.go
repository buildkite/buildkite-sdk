package code_gen_utils

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type CodeGenLanguage interface {
	Name() string
	NewComment(comment string) string
	NewCommentWithIndent(comment string, indent int) string
	RenderImports(imports map[string]string) string
}

const typescriptName = "typescript"

type typescriptMetadata struct {
	commentIdentifier string
}

func (typescriptMetadata) Name() string {
	return typescriptName
}

func (typescriptMetadata) NewComment(comment string) string {
	return utils.CodeGen.Comment.TypeScript(comment, 0)
}

func (typescriptMetadata) NewCommentWithIndent(comment string, indent int) string {
	return utils.CodeGen.Comment.TypeScript(comment, indent)
}

func (typescriptMetadata) RenderImports(imports map[string]string) string {
	block := utils.CodeGen.NewCodeBlock()
	for pkgName, identifier := range imports {
		block = append(block, fmt.Sprintf("import * as %s from \"%s\";", pkgName, identifier))
	}
	return block.Display()
}

type goMetadata struct {
	commentIdentifier string
}

const goName = "go"

func (goMetadata) Name() string {
	return goName
}

func (goMetadata) NewComment(comment string) string {
	return utils.CodeGen.Comment.Go(comment, 0)
}

func (goMetadata) NewCommentWithIndent(comment string, indent int) string {
	return utils.CodeGen.Comment.Go(comment, indent)
}

func (goMetadata) RenderImports(imports map[string]string) string {
	importsBlock := utils.CodeGen.NewCodeBlock()
	for pkgName, identifier := range imports {
		pkgIdentifier := identifier
		if pkgName == identifier {
			pkgIdentifier = ""
		}

		importsBlock = append(importsBlock, fmt.Sprintf("    %s\"%s\"", pkgIdentifier, pkgName))
	}

	return utils.CodeGen.NewCodeBlock(
		"import (",
		importsBlock.Display(),
		")",
	).Display()
}

type codeGenLanguages struct{}

func (codeGenLanguages) TypeScript() typescriptMetadata {
	return typescriptMetadata{
		commentIdentifier: "//",
	}
}

func (codeGenLanguages) Go() goMetadata {
	return goMetadata{
		commentIdentifier: "//",
	}
}

var Languages = codeGenLanguages{}
