package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestNotifySlack(t *testing.T) {
	t.Run("Slack", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			channel := "#general"
			val := buildkite.NotifySlack{
				Slack: &buildkite.NotifySlackSlack{
					String: &channel,
				},
			}
			CheckResult(t, val, `{"slack":"#general"}`)
		})

		t.Run("SlackObject", func(t *testing.T) {
			message := "hi"
			channels := []string{"one", "two"}
			val := buildkite.NotifySlack{
				Slack: &buildkite.NotifySlackSlack{
					NotifySlackObject: &buildkite.NotifySlackObject{
						Message:  &message,
						Channels: channels,
					},
				},
			}
			CheckResult(t, val, `{"slack":{"channels":["one","two"],"message":"hi"}}`)
		})
	})

	t.Run("If", func(t *testing.T) {
		ifVal := "string"
		val := buildkite.NotifySlack{
			If: &ifVal,
		}
		CheckResult(t, val, `{"if":"string"}`)
	})

	t.Run("All", func(t *testing.T) {
		channel := "#general"
		ifVal := "string"
		val := buildkite.NotifySlack{
			If: &ifVal,
			Slack: &buildkite.NotifySlackSlack{
				String: &channel,
			},
		}
		CheckResult(t, val, `{"if":"string","slack":"#general"}`)
	})
}
