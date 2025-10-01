package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestNotifyPagerduty(t *testing.T) {
	t.Run("Pagerduty", func(t *testing.T) {
		changeEvent := "event"
		val := buildkite.NotifyPagerduty{
			PagerdutyChangeEvent: &changeEvent,
		}
		CheckResult(t, val, `{"pagerduty_change_event":"event"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifVal := "string"
		val := buildkite.NotifyPagerduty{
			If: &ifVal,
		}
		CheckResult(t, val, `{"if":"string"}`)
	})

	t.Run("All", func(t *testing.T) {
		changeEvent := "event"
		ifVal := "if"
		val := buildkite.NotifyPagerduty{
			PagerdutyChangeEvent: &changeEvent,
			If:                   &ifVal,
		}
		CheckResult(t, val, `{"if":"if","pagerduty_change_event":"event"}`)
	})
}
