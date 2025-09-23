package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestWaitStep(t *testing.T) {
	t.Run("AllowDependencyFailure", func(t *testing.T) {
		allowDependencyFailure := true
		val := buildkite.WaitStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("Branches", func(t *testing.T) {
		branches := []string{"one", "two"}
		val := buildkite.WaitStep{
			Branches: &buildkite.Branches{
				StringArray: branches,
			},
		}
		CheckResult(t, val, `{"branches":["one","two"]}`)
	})

	t.Run("ContinueOnFailure", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			continueOnFailure := "true"
			val := buildkite.WaitStep{
				ContinueOnFailure: &buildkite.WaitStepContinueOnFailure{
					String: &continueOnFailure,
				},
			}
			CheckResult(t, val, `{"continue_on_failure":"true"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			continueOnFailure := true
			val := buildkite.WaitStep{
				ContinueOnFailure: &buildkite.WaitStepContinueOnFailure{
					Bool: &continueOnFailure,
				},
			}
			CheckResult(t, val, `{"continue_on_failure":true}`)
		})
	})

	t.Run("DependsOn", func(t *testing.T) {
		dependsOn := "step"
		val := buildkite.WaitStep{
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
		}
		CheckResult(t, val, `{"depends_on":"step"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifValue := "if"
		val := buildkite.WaitStep{
			If: &ifValue,
		}
		CheckResult(t, val, `{"if":"if"}`)
	})

	t.Run("Key", func(t *testing.T) {
		key := "key"
		val := buildkite.WaitStep{
			Key: &key,
		}
		CheckResult(t, val, `{"key":"key"}`)
	})

	t.Run("Label", func(t *testing.T) {
		label := "label"
		val := buildkite.WaitStep{
			Label: &label,
		}
		CheckResult(t, val, `{"label":"label"}`)
	})

	t.Run("Name", func(t *testing.T) {
		name := "name"
		val := buildkite.WaitStep{
			Name: &name,
		}
		CheckResult(t, val, `{"name":"name"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		identifier := "identifier"
		val := buildkite.WaitStep{
			Identifier: &identifier,
		}
		CheckResult(t, val, `{"identifier":"identifier"}`)
	})

	t.Run("Id", func(t *testing.T) {
		id := "id"
		val := buildkite.WaitStep{
			Id: &id,
		}
		CheckResult(t, val, `{"id":"id"}`)
	})

	t.Run("Type", func(t *testing.T) {
		typeValue := buildkite.WaitStepTypeValues["wait"]
		val := buildkite.WaitStep{
			Type: &typeValue,
		}
		CheckResult(t, val, `{"type":"wait"}`)
	})

	t.Run("Wait", func(t *testing.T) {
		wait := "wait"
		val := buildkite.WaitStep{
			Wait: &wait,
		}
		CheckResult(t, val, `{"wait":"wait"}`)
	})

	t.Run("All", func(t *testing.T) {
		allowDependencyFailure := true
		branches := []string{"one", "two"}
		continueOnFailure := true
		dependsOn := "step"
		ifValue := "if"
		key := "key"
		label := "label"
		name := "name"
		identifier := "identifier"
		id := "id"
		typeValue := buildkite.WaitStepTypeValues["wait"]
		wait := "wait"

		val := buildkite.WaitStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			Branches: &buildkite.Branches{
				StringArray: branches,
			},
			ContinueOnFailure: &buildkite.WaitStepContinueOnFailure{
				Bool: &continueOnFailure,
			},
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			If:         &ifValue,
			Key:        &key,
			Label:      &label,
			Name:       &name,
			Identifier: &identifier,
			Id:         &id,
			Type:       &typeValue,
			Wait:       &wait,
		}
		CheckResult(t, val, `{"allow_dependency_failure":true,"branches":["one","two"],"continue_on_failure":true,"depends_on":"step","id":"id","identifier":"identifier","if":"if","key":"key","label":"label","name":"name","type":"wait","wait":"wait"}`)
	})
}

func TestNestedWaitStep(t *testing.T) {
	t.Run("Wait", func(t *testing.T) {
		wait := "wait"
		val := buildkite.NestedWaitStep{
			Wait: &buildkite.WaitStep{
				Wait: &wait,
			},
		}
		CheckResult(t, val, `{"wait":{"wait":"wait"}}`)
	})

	t.Run("Wait", func(t *testing.T) {
		wait := "wait"
		val := buildkite.NestedWaitStep{
			Waiter: &buildkite.WaitStep{
				Wait: &wait,
			},
		}
		CheckResult(t, val, `{"waiter":{"wait":"wait"}}`)
	})
}
