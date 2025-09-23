package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testSoftFail struct {
	SoftFail buildkite.SoftFail `json:"soft_fail"`
}

func TestSoftFail(t *testing.T) {
	t.Run("SoftFailEnum", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testSoftFail{
				SoftFail: buildkite.SoftFail{
					SoftFailEnum: &buildkite.SoftFailEnum{
						Bool: &value,
					},
				},
			}
			CheckResult(t, val, `{"soft_fail":true}`)
		})

		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testSoftFail{
				SoftFail: buildkite.SoftFail{
					SoftFailEnum: &buildkite.SoftFailEnum{
						String: &value,
					},
				},
			}
			CheckResult(t, val, `{"soft_fail":"true"}`)
		})
	})

	t.Run("SoftFailList", func(t *testing.T) {
		exitStatus := 1
		list := []buildkite.SoftFailObject{
			{
				ExitStatus: &buildkite.SoftFailObjectExitStatus{
					Int: &exitStatus,
				},
			},
		}
		val := testSoftFail{
			SoftFail: buildkite.SoftFail{
				SoftFailList: &list,
			},
		}
		CheckResult(t, val, `{"soft_fail":[{"exit_status":1}]}`)
	})
}
