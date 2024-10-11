package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

// Imports
type ImportBlock struct {
	imports map[string]string
}

func (i *ImportBlock) AddImport(pkg string, name string) {
	i.imports[pkg] = name
}

func (i *ImportBlock) String() string {
	block := utils.CodeBlock{}

	for pkg, name := range i.imports {
		block = append(block,
			fmt.Sprintf("import * as %s from \"%s\";", name, pkg),
		)
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
	return &File{
		header: "// This file is auto-generated please do not edit",
		imports: &ImportBlock{
			imports: make(map[string]string),
		},
	}
}
