package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testFields struct {
	Fields buildkite.Fields `json:"fields"`
}

func TestFields(t *testing.T) {
	t.Run("TextField", func(t *testing.T) {
		textOne := "textFieldOne"
		textTwo := "textFieldTwo"
		val := testFields{
			Fields: []buildkite.FieldsUnion{
				{
					TextField: &buildkite.TextField{
						Text: &textOne,
					},
				},
				{
					TextField: &buildkite.TextField{
						Text: &textTwo,
					},
				},
			},
		}
		CheckResult(t, val, `{"fields":[{"text":"textFieldOne"},{"text":"textFieldTwo"}]}`)
	})

	t.Run("SelectField", func(t *testing.T) {
		selectOne := "selectFieldOne"
		selectTwo := "selectFieldTwo"
		val := testFields{
			Fields: []buildkite.FieldsUnion{
				{
					SelectField: &buildkite.SelectField{
						Select: &selectOne,
					},
				},
				{
					SelectField: &buildkite.SelectField{
						Select: &selectTwo,
					},
				},
			},
		}
		CheckResult(t, val, `{"fields":[{"select":"selectFieldOne"},{"select":"selectFieldTwo"}]}`)
	})
}
