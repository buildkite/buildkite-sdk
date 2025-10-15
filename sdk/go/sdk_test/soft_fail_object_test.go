package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestSoftFailObject(t *testing.T) {
	t.Run("SoftFailObjectExitStatusEnum", func(t *testing.T) {
		value := buildkite.SoftFailObjectExitStatusEnumValues["*"]
		val := buildkite.SoftFailObject{
			ExitStatus: &buildkite.SoftFailObjectExitStatus{
				SoftFailObjectExitStatusEnum: &value,
			},
		}
		CheckResult(t, val, `{"exit_status":"*"}`)
	})

	t.Run("Int", func(t *testing.T) {
		value := 1
		val := buildkite.SoftFailObject{
			ExitStatus: &buildkite.SoftFailObjectExitStatus{
				Int: &value,
			},
		}
		CheckResult(t, val, `{"exit_status":1}`)
	})
}
