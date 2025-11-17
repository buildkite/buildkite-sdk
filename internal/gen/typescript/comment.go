package typescript

import (
	"fmt"
	"strings"
)

func NewTypeDocComment(comment string) string {
	if comment == "" {
		return ""
	}

	// Handle multi-line comments by splitting and indenting each line
	lines := strings.Split(comment, "\n")
	if len(lines) == 1 {
		return fmt.Sprintf("/**\n * %s\n */", comment)
	}

	// Multi-line comment handling
	var result strings.Builder
	result.WriteString("/**\n")
	for _, line := range lines {
		result.WriteString(fmt.Sprintf(" * %s\n", line))
	}
	result.WriteString(" */")
	return result.String()
}
