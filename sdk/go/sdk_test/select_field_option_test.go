package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testSelectFieldOptionRequired struct {
	Required buildkite.SelectFieldOptionRequired `json:"required"`
}

func TestSelectFieldOption(t *testing.T) {
	t.Run("SelectFieldOptionRequired", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testSelectFieldOptionRequired{
				Required: buildkite.SelectFieldOptionRequired{
					String: &value,
				},
			}
			CheckResult(t, val, `{"required":"true"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			value := true
			val := testSelectFieldOptionRequired{
				Required: buildkite.SelectFieldOptionRequired{
					Bool: &value,
				},
			}
			CheckResult(t, val, `{"required":true}`)
		})
	})

	t.Run("SelectFieldOption", func(t *testing.T) {
		t.Run("Hint", func(t *testing.T) {
			hint := "hint"
			val := buildkite.SelectFieldOption{
				Hint: &hint,
			}
			CheckResult(t, val, `{"hint":"hint"}`)
		})

		t.Run("Label", func(t *testing.T) {
			label := "label"
			val := buildkite.SelectFieldOption{
				Label: &label,
			}
			CheckResult(t, val, `{"label":"label"}`)
		})

		t.Run("Required", func(t *testing.T) {
			required := true
			val := buildkite.SelectFieldOption{
				Required: &buildkite.SelectFieldOptionRequired{
					Bool: &required,
				},
			}
			CheckResult(t, val, `{"required":true}`)
		})

		t.Run("Value", func(t *testing.T) {
			value := "value"
			val := buildkite.SelectFieldOption{
				Value: &value,
			}
			CheckResult(t, val, `{"value":"value"}`)
		})

		t.Run("All", func(t *testing.T) {
			hint := "hint"
			label := "label"
			required := true
			value := "value"
			val := buildkite.SelectFieldOption{
				Hint:  &hint,
				Label: &label,
				Required: &buildkite.SelectFieldOptionRequired{
					Bool: &required,
				},
				Value: &value,
			}
			CheckResult(t, val, `{"hint":"hint","label":"label","required":true,"value":"value"}`)
		})
	})
}
