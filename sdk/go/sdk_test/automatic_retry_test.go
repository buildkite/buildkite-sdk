package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testAutomaticRetry struct {
	AutomaticRetry buildkite.AutomaticRetry `json:"automatic_retry,omitempty"`
}

type testAutomaticRetryExitStatus struct {
	AutomaticRetryExitStatus buildkite.AutomaticRetryExitStatus `json:"status"`
}

func TestAutomaticRetry(t *testing.T) {
	t.Run("AutomaticRetryExitStatus", func(t *testing.T) {
		t.Run("AutomaticRetryExitStatusEnum", func(t *testing.T) {
			status := buildkite.AutomaticRetryExitStatusEnumValues["*"]
			testVal := testAutomaticRetryExitStatus{
				AutomaticRetryExitStatus: buildkite.AutomaticRetryExitStatus{
					AutomaticRetryExitStatusEnum: &status,
				},
			}

			CheckResult(t, testVal, "{\"status\":\"*\"}")
		})

		t.Run("Int", func(t *testing.T) {
			status := 1
			testVal := testAutomaticRetryExitStatus{
				AutomaticRetryExitStatus: buildkite.AutomaticRetryExitStatus{
					Int: &status,
				},
			}

			CheckResult(t, testVal, "{\"status\":1}")
		})

		t.Run("IntArray", func(t *testing.T) {
			status := []int{1, 2}
			testVal := testAutomaticRetryExitStatus{
				AutomaticRetryExitStatus: buildkite.AutomaticRetryExitStatus{
					IntArray: status,
				},
			}

			CheckResult(t, testVal, "{\"status\":[1,2]}")
		})
	})

	t.Run("AutomaticRetry", func(t *testing.T) {
		t.Run("ExitStatus", func(t *testing.T) {
			exitStatus := 1
			testVal := testAutomaticRetry{
				AutomaticRetry: buildkite.AutomaticRetry{
					ExitStatus: &buildkite.AutomaticRetryExitStatus{
						Int: &exitStatus,
					},
				},
			}

			CheckResult(t, testVal, `{"automatic_retry":{"exit_status":1}}`)
		})

		t.Run("Limit", func(t *testing.T) {
			limit := 1
			testVal := testAutomaticRetry{
				AutomaticRetry: buildkite.AutomaticRetry{
					Limit: &limit,
				},
			}

			CheckResult(t, testVal, `{"automatic_retry":{"limit":1}}`)
		})

		t.Run("Signal", func(t *testing.T) {
			signal := "string"
			testVal := testAutomaticRetry{
				AutomaticRetry: buildkite.AutomaticRetry{
					Signal: &signal,
				},
			}

			CheckResult(t, testVal, `{"automatic_retry":{"signal":"string"}}`)
		})

		t.Run("SignalReason", func(t *testing.T) {
			signalReason := buildkite.AutomaticRetrySignalReasonValues["none"]
			testVal := testAutomaticRetry{
				AutomaticRetry: buildkite.AutomaticRetry{
					SignalReason: &signalReason,
				},
			}

			CheckResult(t, testVal, `{"automatic_retry":{"signal_reason":"none"}}`)
		})

		t.Run("All", func(t *testing.T) {
			exitStatus := 1
			limit := 2
			signal := "string"
			signalReason := buildkite.AutomaticRetrySignalReasonValues["none"]
			testVal := testAutomaticRetry{
				AutomaticRetry: buildkite.AutomaticRetry{
					ExitStatus: &buildkite.AutomaticRetryExitStatus{
						Int: &exitStatus,
					},
					Limit:        &limit,
					Signal:       &signal,
					SignalReason: &signalReason,
				},
			}

			CheckResult(t, testVal, `{"automatic_retry":{"exit_status":1,"limit":2,"signal":"string","signal_reason":"none"}}`)
		})
	})
}
