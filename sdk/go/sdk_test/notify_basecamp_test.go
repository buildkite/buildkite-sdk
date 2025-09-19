package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

func TestNotifyBasecamp(t *testing.T) {
	t.Run("Basecamp", func(t *testing.T) {
		value := "string"
		val := buildkite.NotifyBasecamp{
			BasecampCampfire: &value,
		}
		CheckResult(t, val, `{"basecamp_campfire":"string"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifVal := "string"
		val := buildkite.NotifyBasecamp{
			If: &ifVal,
		}
		CheckResult(t, val, `{"if":"string"}`)
	})

	t.Run("All", func(t *testing.T) {
		value := "value"
		ifVal := "if"
		val := buildkite.NotifyBasecamp{
			BasecampCampfire: &value,
			If:               &ifVal,
		}
		CheckResult(t, val, `{"basecamp_campfire":"value","if":"if"}`)
	})
}
