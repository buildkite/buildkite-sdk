package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

type testCommandStepAutomaticRetry struct {
	Retry buildkite.CommandStepAutomaticRetry `json:"retry"`
}

func TestCommandStepAutomaticRetry(t *testing.T) {
	t.Run("CommandStepAutomaticRetryEnum", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testCommandStepAutomaticRetry{
				Retry: buildkite.CommandStepAutomaticRetry{
					CommandStepAutomaticRetryEnum: &buildkite.CommandStepAutomaticRetryEnum{
						Bool: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":true}`)
		})

		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testCommandStepAutomaticRetry{
				Retry: buildkite.CommandStepAutomaticRetry{
					CommandStepAutomaticRetryEnum: &buildkite.CommandStepAutomaticRetryEnum{
						String: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":"true"}`)
		})
	})

	t.Run("AutomaticRetry", func(t *testing.T) {
		limit := 1
		val := testCommandStepAutomaticRetry{
			Retry: buildkite.CommandStepAutomaticRetry{
				AutomaticRetry: &buildkite.AutomaticRetry{
					Limit: &limit,
				},
			},
		}
		CheckResult(t, val, `{"retry":{"limit":1}}`)
	})

	t.Run("AutomaticRetryList", func(t *testing.T) {
		limit := 1
		list := []buildkite.AutomaticRetry{
			{
				Limit: &limit,
			},
		}
		val := testCommandStepAutomaticRetry{
			Retry: buildkite.CommandStepAutomaticRetry{
				AutomaticRetryList: &list,
			},
		}
		CheckResult(t, val, `{"retry":[{"limit":1}]}`)
	})
}

type testCommandStepManualRetryObject struct {
	Retry buildkite.CommandStepManualRetryObject `json:"retry"`
}

func TestCommandStepManualRetryObject(t *testing.T) {
	t.Run("Allowed", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testCommandStepManualRetryObject{
				Retry: buildkite.CommandStepManualRetryObject{
					Allowed: &buildkite.CommandStepManualRetryObjectAllowed{
						Bool: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":{"allowed":true}}`)
		})

		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testCommandStepManualRetryObject{
				Retry: buildkite.CommandStepManualRetryObject{
					Allowed: &buildkite.CommandStepManualRetryObjectAllowed{
						String: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":{"allowed":"true"}}`)
		})
	})

	t.Run("PermitOnPassed", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testCommandStepManualRetryObject{
				Retry: buildkite.CommandStepManualRetryObject{
					PermitOnPassed: &buildkite.CommandStepManualRetryObjectPermitOnPassed{
						Bool: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":{"permit_on_passed":true}}`)
		})

		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testCommandStepManualRetryObject{
				Retry: buildkite.CommandStepManualRetryObject{
					PermitOnPassed: &buildkite.CommandStepManualRetryObjectPermitOnPassed{
						String: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":{"permit_on_passed":"true"}}`)
		})
	})

	t.Run("Reason", func(t *testing.T) {
		value := "reason"
		val := testCommandStepManualRetryObject{
			Retry: buildkite.CommandStepManualRetryObject{
				Reason: &value,
			},
		}
		CheckResult(t, val, `{"retry":{"reason":"reason"}}`)
	})

	t.Run("All", func(t *testing.T) {
		allowed := true
		permit := "false"
		reason := "reason"
		val := testCommandStepManualRetryObject{
			Retry: buildkite.CommandStepManualRetryObject{
				Reason: &reason,
				PermitOnPassed: &buildkite.CommandStepManualRetryObjectPermitOnPassed{
					String: &permit,
				},
				Allowed: &buildkite.CommandStepManualRetryObjectAllowed{
					Bool: &allowed,
				},
			},
		}
		CheckResult(t, val, `{"retry":{"allowed":true,"permit_on_passed":"false","reason":"reason"}}`)
	})
}

type testCommandStepManualRetry struct {
	Retry buildkite.CommandStepManualRetry `json:"retry"`
}

func TestCommandStepManualRetry(t *testing.T) {
	t.Run("CommandStepManualRetryEnum", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testCommandStepManualRetry{
				Retry: buildkite.CommandStepManualRetry{
					CommandStepManualRetryEnum: &buildkite.CommandStepManualRetryEnum{
						Bool: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":true}`)
		})

		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testCommandStepManualRetry{
				Retry: buildkite.CommandStepManualRetry{
					CommandStepManualRetryEnum: &buildkite.CommandStepManualRetryEnum{
						String: &value,
					},
				},
			}
			CheckResult(t, val, `{"retry":"true"}`)
		})
	})

	t.Run("CommandStepManualRetryObject", func(t *testing.T) {
		allowed := true
		permit := "false"
		reason := "reason"
		val := testCommandStepManualRetry{
			Retry: buildkite.CommandStepManualRetry{
				CommandStepManualRetryObject: &buildkite.CommandStepManualRetryObject{
					Reason: &reason,
					PermitOnPassed: &buildkite.CommandStepManualRetryObjectPermitOnPassed{
						String: &permit,
					},
					Allowed: &buildkite.CommandStepManualRetryObjectAllowed{
						Bool: &allowed,
					},
				},
			},
		}
		CheckResult(t, val, `{"retry":{"allowed":true,"permit_on_passed":"false","reason":"reason"}}`)
	})
}
