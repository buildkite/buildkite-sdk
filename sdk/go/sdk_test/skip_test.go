package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testSkip struct {
	Skip buildkite.Skip `json:"skip"`
}

func TestSkip(t *testing.T) {
	t.Run("Bool", func(t *testing.T) {
		value := true
		val := testSkip{
			Skip: buildkite.Skip{
				Bool: &value,
			},
		}
		CheckResult(t, val, `{"skip":true}`)
	})

	t.Run("String", func(t *testing.T) {
		value := "string"
		val := testSkip{
			Skip: buildkite.Skip{
				String: &value,
			},
		}
		CheckResult(t, val, `{"skip":"string"}`)
	})
}
