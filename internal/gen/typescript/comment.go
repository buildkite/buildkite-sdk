package typescript

import (
	"fmt"
	"strings"
)

func NewTypeDocCommentWithIndent(comment string, indentSize int) string {
	if comment == "" {
		return ""
	}

	indent := strings.Repeat(" ", indentSize)

	// Handle multi-line comments by splitting and indenting each line
	lines := strings.Split(comment, "\n")
	if len(lines) == 1 {
		return fmt.Sprintf("%s/**\n%s * %s\n%s */", indent, indent, comment, indent)
	}

	// Multi-line comment handling
	var result strings.Builder
	result.WriteString("/**\n")
	for _, line := range lines {
		result.WriteString(fmt.Sprintf("%s * %s\n", indent, line))
	}
	result.WriteString(fmt.Sprintf("%s */", indent))
	return result.String()
}

func NewTypeDocComment(comment string) string {
	return NewTypeDocCommentWithIndent(comment, 0)
}
