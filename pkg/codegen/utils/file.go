package code_gen_utils

import (
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

type CodeGenFile struct {
	language CodeGenLanguage
	header   utils.CodeBlock
	imports  map[string]string
	code     utils.CodeBlock
}

func (c *CodeGenFile) AppendToHeader(str string) {
	c.header = append(c.header, str)
}

func (c *CodeGenFile) AddImport(pkgName, identifier string) {
	c.imports[pkgName] = identifier
}

func (c CodeGenFile) RenderImports() string {
	imports := utils.CodeGen.NewCodeBlock()
	imports = append(imports, c.language.RenderImports(c.imports))
	return imports.Display()
}

func (c *CodeGenFile) AppendCode(code string) {
	c.code = append(c.code, code)
}

func (c CodeGenFile) NewComment(comment string) string {
	return c.language.NewComment(comment)
}

func (c CodeGenFile) NewCommentWithIndent(comment string, indent int) string {
	return c.language.NewCommentWithIndent(comment, indent)
}

func (c CodeGenFile) Render() string {
	return utils.CodeGen.NewCodeBlock(
		c.header.Display(),
		c.RenderImports(),
		c.code.Display(),
	).Display()
}

type file struct{}

func (file) New(language CodeGenLanguage) *CodeGenFile {
	return &CodeGenFile{
		language: language,
		header: utils.CodeGen.NewCodeBlock(
			language.NewComment("This file is auto-generated. Do not edit."),
		),
		imports: make(map[string]string),
		code:    utils.CodeGen.NewCodeBlock(),
	}
}

var File = file{}
