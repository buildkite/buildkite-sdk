package schema_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeName(t *testing.T) {
	testValue := AttributeName("really_cool_test_value")

	t.Run("should transform to title case", func(t *testing.T) {
		assert.Equal(t, "ReallyCoolTestValue", testValue.TitleCase())
	})

	t.Run("should transform to camel case", func(t *testing.T) {
		assert.Equal(t, "reallyCoolTestValue", testValue.CamelCase())
	})
}
