package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

// Imports
type ImportBlock struct {
	imports utils.CodeBlock
}

func (i *ImportBlock) AddImports(names ...string) {
	for _, name := range names {
		i.imports = append(i.imports, fmt.Sprintf("    \"%s\"", name))
	}
}

func (i *ImportBlock) String() string {
	if len(i.imports) == 0 {
		return ""
	}

	block := utils.CodeBlock{
		"import (",
		i.imports.Display(),
		")",
	}

	return block.Display()
}

// File
type File struct {
	header  string
	imports *ImportBlock
	code    utils.CodeBlock
}

func (f *File) String() string {
	file := utils.CodeBlock{
		f.header,
		f.imports.String(),
		f.code.Display(),
	}

	return file.Display()
}

func NewFile() *File {
	header := utils.CodeBlock{
		"// This file is auto-generated please do not edit",
		"",
		"package buildkite",
	}

	return &File{
		header:  header.Display(),
		imports: &ImportBlock{},
	}
}
