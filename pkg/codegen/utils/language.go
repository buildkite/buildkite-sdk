package code_gen_utils

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type CodeGenLanguage interface {
	Name() string
	NewComment(comment string) string
	NewCommentWithIndent(comment string, indent int) string
	RenderImports(imports []codeGenImport) string
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

func (typescriptMetadata) RenderImports(imports []codeGenImport) string {
	block := utils.CodeGen.NewCodeBlock()
	for _, pkg := range imports {
		block = append(block, fmt.Sprintf("import * as %s from \"%s\";", pkg.pkgName, pkg.identifier))
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

func (goMetadata) RenderImports(imports []codeGenImport) string {
	importsBlock := utils.CodeGen.NewCodeBlock()
	for _, pkg := range imports {
		pkgIdentifier := pkg.identifier
		if pkg.pkgName == pkg.identifier {
			pkgIdentifier = ""
		}

		importsBlock = append(importsBlock, fmt.Sprintf("    %s\"%s\"", pkgIdentifier, pkg.pkgName))
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
