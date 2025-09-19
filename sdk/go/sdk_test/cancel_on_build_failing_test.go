package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

type testCancelOnBuildFailing struct {
	CancelOnBuildFailing buildkite.CancelOnBuildFailing `json:"cancel_on_build_failing"`
}

func TestCancelOnBuildFailingEnum(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		val := "true"
		testVal := testCancelOnBuildFailing{
			CancelOnBuildFailing: buildkite.CancelOnBuildFailing{
				String: &val,
			},
		}
		CheckResult(t, testVal, `{"cancel_on_build_failing":"true"}`)
	})

	t.Run("Boolean", func(t *testing.T) {
		val := true
		testVal := testCancelOnBuildFailing{
			CancelOnBuildFailing: buildkite.CancelOnBuildFailing{
				Bool: &val,
			},
		}
		CheckResult(t, testVal, `{"cancel_on_build_failing":true}`)
	})
}
