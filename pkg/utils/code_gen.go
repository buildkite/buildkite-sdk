package utils

import (
	"fmt"
	"strings"
)

// A CodeBlock contains lines of generated code.
type CodeBlock []string

func (c CodeBlock) Display() string {
	return strings.Join(c, "\n")
}

func (c CodeBlock) DisplayIndent(spaces int) string {
	indent := strings.Repeat(" ", spaces)
	return fmt.Sprintf("%s%s",
		indent,
		strings.Join(c, fmt.Sprintf("\n%s", indent)),
	)
}

type codeComment struct{}

func (codeComment) newComment(identifier, str string, indent int) string {
	parts := strings.Split(str, "\n")
	spaces := strings.Repeat(" ", indent)

	var result []string
	for i, part := range parts {
		// If the last line is empty don't include it.
		if i == (len(parts)-1) && part == "" {
			continue
		}

		result = append(result, fmt.Sprintf("%s%s %s", spaces, identifier, part))
	}

	return strings.Join(result, "\n")
}

func (c codeComment) TypeScript(str string, indent int) string {
	return c.newComment("//", str, indent)
}

func (c codeComment) Go(str string, indent int) string {
	return c.newComment("//", str, indent)
}

type codeGen struct {
	Comment codeComment
}

func (codeGen) NewCodeBlock(lines ...string) CodeBlock {
	block := CodeBlock{}
	for _, line := range lines {
		block = append(block, line)
	}

	return block
}

var CodeGen = codeGen{
	Comment: codeComment{},
}
