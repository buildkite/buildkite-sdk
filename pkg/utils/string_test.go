package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIndent(t *testing.T) {
	t.Run("should indent a string", func(t *testing.T) {
		result := String.AddIndent("test", 4)
		assert.Equal(t, "    test", result)
	})
}

func TestCapitalize(t *testing.T) {
	t.Run("should capitalize a string", func(t *testing.T) {
		result := String.Capitalize("test")
		assert.Equal(t, "Test", result)
	})
}

func TestSnakeCaseToTitleCase(t *testing.T) {
	testValues := map[string]string{
		"test_value":        "TestValue",
		"longer_test_value": "LongerTestValue",
		"short":             "Short",
	}

	for value, expectation := range testValues {
		t.Run(fmt.Sprintf("should transform %s to title case", value), func(t *testing.T) {
			result := String.SnakeCaseToTitleCase(value)
			assert.Equal(t, expectation, result)
		})
	}
}

func TestCamelCaseToTitleCase(t *testing.T) {
	testValues := map[string]string{
		"test_value":        "testValue",
		"longer_test_value": "longerTestValue",
		"short":             "short",
	}

	for value, expectation := range testValues {
		t.Run(fmt.Sprintf("should transform %s to camel case", value), func(t *testing.T) {
			result := String.SnakeCaseToCamelCase(value)
			assert.Equal(t, expectation, result)
		})
	}
}
