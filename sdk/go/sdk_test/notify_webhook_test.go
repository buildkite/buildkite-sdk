package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

func TestNotifyWebhook(t *testing.T) {
	t.Run("Webhook", func(t *testing.T) {
		url := "string"
		val := buildkite.NotifyWebhook{
			Webhook: &url,
		}
		CheckResult(t, val, `{"webhook":"string"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifVal := "string"
		val := buildkite.NotifyWebhook{
			If: &ifVal,
		}
		CheckResult(t, val, `{"if":"string"}`)
	})

	t.Run("All", func(t *testing.T) {
		webhook := "string"
		ifVal := "if"
		val := buildkite.NotifyWebhook{
			Webhook: &webhook,
			If:      &ifVal,
		}
		CheckResult(t, val, `{"if":"if","webhook":"string"}`)
	})
}
