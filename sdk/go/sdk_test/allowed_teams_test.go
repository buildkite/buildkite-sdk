package sdk_test

import (
	"encoding/json"
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
	"github.com/stretchr/testify/assert"
)

type testAllowedTeams struct {
	AllowedTeams buildkite.AllowedTeams `json:"allowed_teams"`
}

func TestAllowedTeams(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		val := "string"
		testVal := testAllowedTeams{
			AllowedTeams: buildkite.AllowedTeams{
				String: &val,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"allowed_teams\":\"string\"}", string(result))
	})

	t.Run("StringArray", func(t *testing.T) {
		val := []string{"string"}
		testVal := testAllowedTeams{
			AllowedTeams: buildkite.AllowedTeams{
				StringArray: val,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"allowed_teams\":[\"string\"]}", string(result))
	})
}
