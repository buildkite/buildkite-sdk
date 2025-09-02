package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

func TestSelectField(t *testing.T) {
	t.Run("SelectFieldDefault", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			value := "string"
			val := buildkite.SelectField{
				Default: &buildkite.SelectFieldDefault{
					String: &value,
				},
			}
			CheckResult(t, val, `{"default":"string"}`)
		})

		t.Run("StringArray", func(t *testing.T) {
			value := []string{"one", "two"}
			val := buildkite.SelectField{
				Default: &buildkite.SelectFieldDefault{
					StringArray: value,
				},
			}
			CheckResult(t, val, `{"default":["one","two"]}`)
		})
	})

	t.Run("SelectFieldMultiple", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			value := "true"
			val := buildkite.SelectField{
				Multiple: &buildkite.SelectFieldMultiple{
					String: &value,
				},
			}
			CheckResult(t, val, `{"multiple":"true"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			value := true
			val := buildkite.SelectField{
				Multiple: &buildkite.SelectFieldMultiple{
					Bool: &value,
				},
			}
			CheckResult(t, val, `{"multiple":true}`)
		})
	})

	t.Run("SelectFieldRequired", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			value := "true"
			val := buildkite.SelectField{
				Required: &buildkite.SelectFieldRequired{
					String: &value,
				},
			}
			CheckResult(t, val, `{"required":"true"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			value := true
			val := buildkite.SelectField{
				Required: &buildkite.SelectFieldRequired{
					Bool: &value,
				},
			}
			CheckResult(t, val, `{"required":true}`)
		})
	})

	t.Run("All", func(t *testing.T) {
		defaultVal := "value"
		hint := "hint"
		key := "key"
		selectVal := "select"
		multiple := false
		required := false
		optionVal := "optionValue"
		options := []buildkite.SelectFieldOption{
			{Value: &optionVal},
		}

		val := buildkite.SelectField{
			Default: &buildkite.SelectFieldDefault{
				String: &defaultVal,
			},
			Hint:   &hint,
			Key:    &key,
			Select: &selectVal,
			Multiple: &buildkite.SelectFieldMultiple{
				Bool: &multiple,
			},
			Required: &buildkite.SelectFieldRequired{
				Bool: &required,
			},
			Options: options,
		}
		CheckResult(t, val, `{"default":"value","hint":"hint","key":"key","multiple":false,"options":[{"value":"optionValue"}],"required":false,"select":"select"}`)
	})
}
