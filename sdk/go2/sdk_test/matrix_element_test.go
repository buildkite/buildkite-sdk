package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

type testMatrixElement struct {
	Element buildkite.MatrixElement `json:"element"`
}

func TestMatrixElement(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		value := "string"
		val := testMatrixElement{
			Element: buildkite.MatrixElement{
				String: &value,
			},
		}
		CheckResult(t, val, `{"element":"string"}`)
	})

	t.Run("Integer", func(t *testing.T) {
		value := 1
		val := testMatrixElement{
			Element: buildkite.MatrixElement{
				Int: &value,
			},
		}
		CheckResult(t, val, `{"element":1}`)
	})

	t.Run("Bool", func(t *testing.T) {
		value := true
		val := testMatrixElement{
			Element: buildkite.MatrixElement{
				Bool: &value,
			},
		}
		CheckResult(t, val, `{"element":true}`)
	})
}
