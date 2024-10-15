package utils

import (
	"fmt"
	"strings"
)

type stringFunctions struct{}

func (stringFunctions) AddIndent(str string, indent int) string {
	return fmt.Sprintf("%s%s", strings.Repeat(" ", indent), str)
}

func (stringFunctions) Capitalize(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func (stringFunctions) SnakeCaseToTitleCase(str string) string {
	result := ""
	parts := strings.Split(str, "_")
	for _, part := range parts {
		result += String.Capitalize(strings.ToLower(part))
	}

	return result
}

func (stringFunctions) SnakeCaseToCamelCase(str string) string {
	result := ""
	parts := strings.Split(str, "_")
	for i, part := range parts {
		if i == 0 {
			result += strings.ToLower(part)
			continue
		}

		result += String.Capitalize(strings.ToLower(part))
	}

	return result
}

var String = stringFunctions{}
