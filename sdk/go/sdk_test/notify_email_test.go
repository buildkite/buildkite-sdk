package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

func TestNotifyEmail(t *testing.T) {
	t.Run("Email", func(t *testing.T) {
		email := "string"
		val := buildkite.NotifyEmail{
			Email: &email,
		}
		CheckResult(t, val, `{"email":"string"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifVal := "string"
		val := buildkite.NotifyEmail{
			If: &ifVal,
		}
		CheckResult(t, val, `{"if":"string"}`)
	})

	t.Run("All", func(t *testing.T) {
		email := "email"
		ifVal := "if"
		val := buildkite.NotifyEmail{
			Email: &email,
			If:    &ifVal,
		}
		CheckResult(t, val, `{"email":"email","if":"if"}`)
	})
}
