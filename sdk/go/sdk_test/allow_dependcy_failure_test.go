package sdk_test

import (
	"encoding/json"
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
	"github.com/stretchr/testify/assert"
)

type testAllowDependencyFailure struct {
	AllowDependencyFailure buildkite.AllowDependencyFailure `json:"allowDependencyFailure"`
}

func TestAllowDependencyFailureEnum(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		val := "true"
		testVal := testAllowDependencyFailure{
			AllowDependencyFailure: buildkite.AllowDependencyFailure{
				String: &val,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"allowDependencyFailure\":\"true\"}", string(result))
	})

	t.Run("Boolean", func(t *testing.T) {
		val := true
		testVal := testAllowDependencyFailure{
			AllowDependencyFailure: buildkite.AllowDependencyFailure{
				Bool: &val,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"allowDependencyFailure\":true}", string(result))
	})
}
