package sdk_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckResult(t *testing.T, value any, expected string) {
	result, err := json.Marshal(value)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}
