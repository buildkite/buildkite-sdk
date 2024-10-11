package utils

import (
	"fmt"
	"strings"
)

func AddIndent(str string, indent int) string {
	return fmt.Sprintf("%s%s", strings.Repeat(" ", indent), str)
}

func Capitalize(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func SnakeCaseToTitleCase(str string) string {
	result := ""
	parts := strings.Split(str, "_")
	for _, part := range parts {
		result += Capitalize(strings.ToLower(part))
	}

	return result
}

func SnakeCaseToCamelCase(str string) string {
	result := ""
	parts := strings.Split(str, "_")
	for i, part := range parts {
		if i == 0 {
			result += strings.ToLower(part)
			continue
		}

		result += Capitalize(strings.ToLower(part))
	}

	return result
}
