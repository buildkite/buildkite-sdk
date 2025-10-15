package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AI Generated
func CamelCaseToSnakeCase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func TitleCaseToSnakeCase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func CamelCaseToTitleCase(str string) string {
	caser := cases.Title(language.English, cases.NoLower)
	return caser.String(str)
}

func DashCaseToTitleCase(str string) string {
	words := strings.Split(str, "_")
	var titleCaseWords []string

	for _, word := range words {
		if word == "" {
			continue
		}
		firstChar := string(unicode.ToUpper(rune(word[0])))
		restOfWord := strings.ToLower(word[1:])
		titleCaseWords = append(titleCaseWords, firstChar+restOfWord)
	}

	return strings.Join(titleCaseWords, "")
}
