package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestTriggerStep(t *testing.T) {
	t.Run("AllowDependencyFailure", func(t *testing.T) {
		allowDependencyFailure := true
		val := buildkite.TriggerStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("Async", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			async := "true"
			val := buildkite.TriggerStep{
				Async: &buildkite.TriggerStepAsync{
					String: &async,
				},
			}
			CheckResult(t, val, `{"async":"true"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			async := true
			val := buildkite.TriggerStep{
				Async: &buildkite.TriggerStepAsync{
					Bool: &async,
				},
			}
			CheckResult(t, val, `{"async":true}`)
		})
	})

	t.Run("Branches", func(t *testing.T) {
		branches := "branch"
		val := buildkite.TriggerStep{
			Branches: &buildkite.Branches{
				String: &branches,
			},
		}
		CheckResult(t, val, `{"branches":"branch"}`)
	})

	t.Run("Build", func(t *testing.T) {
		t.Run("Branch", func(t *testing.T) {
			branch := "branch"
			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					Branch: &branch,
				},
			}
			CheckResult(t, val, `{"build":{"branch":"branch"}}`)
		})

		t.Run("Commit", func(t *testing.T) {
			commit := "commit"
			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					Commit: &commit,
				},
			}
			CheckResult(t, val, `{"build":{"commit":"commit"}}`)
		})

		t.Run("Env", func(t *testing.T) {
			env := map[string]interface{}{"foo": "bar"}
			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					Env: &env,
				},
			}
			CheckResult(t, val, `{"build":{"env":{"foo":"bar"}}}`)
		})

		t.Run("Message", func(t *testing.T) {
			message := "message"
			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					Message: &message,
				},
			}
			CheckResult(t, val, `{"build":{"message":"message"}}`)
		})

		t.Run("MetaData", func(t *testing.T) {
			metadata := map[string]interface{}{"foo": "bar"}
			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					MetaData: &metadata,
				},
			}
			CheckResult(t, val, `{"build":{"meta_data":{"foo":"bar"}}}`)
		})

		t.Run("All", func(t *testing.T) {
			branch := "branch"
			commit := "commit"
			env := map[string]interface{}{"foo": "bar"}
			message := "message"
			metadata := map[string]interface{}{"foo": "bar"}

			val := buildkite.TriggerStep{
				Build: &buildkite.TriggerStepBuild{
					Branch:   &branch,
					Commit:   &commit,
					Env:      &env,
					Message:  &message,
					MetaData: &metadata,
				},
			}

			CheckResult(t, val, `{"build":{"branch":"branch","commit":"commit","env":{"foo":"bar"},"message":"message","meta_data":{"foo":"bar"}}}`)
		})
	})

	t.Run("DependsOn", func(t *testing.T) {
		dependsOn := "step"
		val := buildkite.TriggerStep{
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
		}
		CheckResult(t, val, `{"depends_on":"step"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifValue := "if"
		val := buildkite.TriggerStep{
			If: &ifValue,
		}
		CheckResult(t, val, `{"if":"if"}`)
	})

	t.Run("IfChanged", func(t *testing.T) {
		val := buildkite.GroupStep{
			IfChanged: &buildkite.IfChanged{
				String: buildkite.Value("ifChanged"),
			},
		}
		CheckResult(t, val, `{"if_changed":"ifChanged"}`)
	})

	t.Run("Key", func(t *testing.T) {
		key := "key"
		val := buildkite.TriggerStep{
			Key: &key,
		}
		CheckResult(t, val, `{"key":"key"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		identifier := "identifier"
		val := buildkite.TriggerStep{
			Identifier: &identifier,
		}
		CheckResult(t, val, `{"identifier":"identifier"}`)
	})

	t.Run("Id", func(t *testing.T) {
		id := "id"
		val := buildkite.TriggerStep{
			Id: &id,
		}
		CheckResult(t, val, `{"id":"id"}`)
	})

	t.Run("Label", func(t *testing.T) {
		label := "label"
		val := buildkite.TriggerStep{
			Label: &label,
		}
		CheckResult(t, val, `{"label":"label"}`)
	})

	t.Run("Name", func(t *testing.T) {
		name := "name"
		val := buildkite.TriggerStep{
			Name: &name,
		}
		CheckResult(t, val, `{"name":"name"}`)
	})

	t.Run("Type", func(t *testing.T) {
		typeValue := buildkite.TriggerStepTypeValues["trigger"]
		val := buildkite.TriggerStep{
			Type: &typeValue,
		}
		CheckResult(t, val, `{"type":"trigger"}`)
	})

	t.Run("Trigger", func(t *testing.T) {
		trigger := "trigger"
		val := buildkite.TriggerStep{
			Trigger: &trigger,
		}
		CheckResult(t, val, `{"trigger":"trigger"}`)
	})

	t.Run("Skip", func(t *testing.T) {
		skip := true
		val := buildkite.TriggerStep{
			Skip: &buildkite.Skip{
				Bool: &skip,
			},
		}
		CheckResult(t, val, `{"skip":true}`)
	})

	t.Run("Softfail", func(t *testing.T) {
		softFail := "true"
		val := buildkite.TriggerStep{
			SoftFail: &buildkite.SoftFail{
				SoftFailEnum: &buildkite.SoftFailEnum{
					String: &softFail,
				},
			},
		}
		CheckResult(t, val, `{"soft_fail":"true"}`)
	})

	t.Run("All", func(t *testing.T) {
		allowDependencyFailure := true
		async := true
		branches := "branch"
		buildCommit := "commit"
		dependsOn := "step"
		ifValue := "if"
		key := "key"
		identifier := "identifier"
		id := "id"
		label := "label"
		name := "name"
		typeValue := buildkite.TriggerStepTypeValues["trigger"]
		trigger := "trigger"
		skip := true
		softFail := "true"

		val := buildkite.TriggerStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			Async: &buildkite.TriggerStepAsync{
				Bool: &async,
			},
			Branches: &buildkite.Branches{
				String: &branches,
			},
			Build: &buildkite.TriggerStepBuild{
				Commit: &buildCommit,
			},
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			If: &ifValue,
			IfChanged: &buildkite.IfChanged{
				String: buildkite.Value("ifChanged"),
			},
			Key:        &key,
			Identifier: &identifier,
			Id:         &id,
			Label:      &label,
			Name:       &name,
			Type:       &typeValue,
			Trigger:    &trigger,
			Skip: &buildkite.Skip{
				Bool: &skip,
			},
			SoftFail: &buildkite.SoftFail{
				SoftFailEnum: &buildkite.SoftFailEnum{
					String: &softFail,
				},
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true,"async":true,"branches":"branch","build":{"commit":"commit"},"depends_on":"step","id":"id","identifier":"identifier","if":"if","if_changed":"ifChanged","key":"key","label":"label","name":"name","skip":true,"soft_fail":"true","trigger":"trigger","type":"trigger"}`)
	})
}

func TestNestedTriggerStep(t *testing.T) {
	trigger := "trigger"
	val := buildkite.NestedTriggerStep{
		Trigger: &buildkite.TriggerStep{
			Trigger: &trigger,
		},
	}
	CheckResult(t, val, `{"trigger":{"trigger":"trigger"}}`)
}
