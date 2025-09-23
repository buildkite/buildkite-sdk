package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testCommandStepCommand struct {
	Command buildkite.CommandStepCommand `json:"command"`
}

func TestCommandStepCommand(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		value := "string"
		val := testCommandStepCommand{
			Command: buildkite.CommandStepCommand{
				String: &value,
			},
		}
		CheckResult(t, val, `{"command":"string"}`)
	})

	t.Run("StringArray", func(t *testing.T) {
		value := []string{"one", "two"}
		val := testCommandStepCommand{
			Command: buildkite.CommandStepCommand{
				StringArray: value,
			},
		}
		CheckResult(t, val, `{"command":["one","two"]}`)
	})
}
