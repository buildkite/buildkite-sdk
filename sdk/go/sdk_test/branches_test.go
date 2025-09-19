package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testBranches struct {
	Branches buildkite.Branches `json:"branches"`
}

func TestBranches(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		val := "string"
		testVal := testBranches{
			Branches: buildkite.Branches{
				String: &val,
			},
		}
		CheckResult(t, testVal, `{"branches":"string"}`)
	})

	t.Run("StringArray", func(t *testing.T) {
		val := []string{"one", "two"}
		testVal := testBranches{
			Branches: buildkite.Branches{
				StringArray: val,
			},
		}
		CheckResult(t, testVal, `{"branches":["one","two"]}`)
	})
}
