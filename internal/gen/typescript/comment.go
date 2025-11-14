package typescript

import (
	"fmt"
	"strings"
)

const commentIndent = "    "

func NewTypeDocComment(comment string) string {
	if comment == "" {
		return ""
	}

	// Handle multi-line comments by splitting and indenting each line
	lines := strings.Split(comment, "\n")
	if len(lines) == 1 {
		return fmt.Sprintf("/**\n%s * %s\n%s */", commentIndent, comment, commentIndent)
	}

	// Multi-line comment handling
	var result strings.Builder
	result.WriteString("/**\n")
	for _, line := range lines {
		result.WriteString(fmt.Sprintf("%s * %s\n", commentIndent, line))
	}
	result.WriteString(fmt.Sprintf("%s */", commentIndent))
	return result.String()
}
