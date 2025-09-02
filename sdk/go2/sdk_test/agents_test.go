package sdk_test

import (
	"encoding/json"
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
	"github.com/stretchr/testify/assert"
)

type testAgents struct {
	Agents buildkite.Agents `json:"agents"`
}

func TestAgents(t *testing.T) {
	t.Run("AgentsList", func(t *testing.T) {
		agents := []string{"one", "two"}
		testVal := testAgents{
			Agents: buildkite.Agents{
				AgentsList: &agents,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"agents\":[\"one\",\"two\"]}", string(result))
	})

	t.Run("AgentsObject", func(t *testing.T) {
		agents := map[string]string{
			"one": "two",
		}
		testVal := testAgents{
			Agents: buildkite.Agents{
				AgentsObject: &agents,
			},
		}

		result, err := json.Marshal(testVal)
		assert.NoError(t, err)
		assert.Equal(t, "{\"agents\":{\"one\":\"two\"}}", string(result))
	})
}
