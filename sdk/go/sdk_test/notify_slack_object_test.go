package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

func TestNotifySlackObject(t *testing.T) {
	t.Run("SlackObject", func(t *testing.T) {
		message := "hi"
		val := buildkite.NotifySlackObject{
			Message: &message,
		}
		CheckResult(t, val, `{"message":"hi"}`)
	})

	t.Run("If", func(t *testing.T) {
		channels := []string{"one", "two"}
		val := buildkite.NotifySlackObject{
			Channels: channels,
		}
		CheckResult(t, val, `{"channels":["one","two"]}`)
	})

	t.Run("All", func(t *testing.T) {
		message := "hi"
		channels := []string{"one", "two"}
		val := buildkite.NotifySlackObject{
			Message:  &message,
			Channels: channels,
		}
		CheckResult(t, val, `{"channels":["one","two"],"message":"hi"}`)
	})
}
