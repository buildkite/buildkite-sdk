package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testDependsOn struct {
	DependsOn buildkite.DependsOn `json:"depends_on"`
}

func TestDependsOn(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		val := "string"
		testVal := testDependsOn{
			DependsOn: buildkite.DependsOn{
				String: &val,
			},
		}
		CheckResult(t, testVal, `{"depends_on":"string"}`)
	})

	t.Run("DependsOnList", func(t *testing.T) {
		t.Run("Simple", func(t *testing.T) {
			one := "one"
			two := "two"
			val := []buildkite.DependsOnListItem{
				{String: &one},
				{String: &two},
			}
			testVal := testDependsOn{
				DependsOn: buildkite.DependsOn{
					DependsOnList: &val,
				},
			}
			CheckResult(t, testVal, `{"depends_on":["one","two"]}`)
		})

		t.Run("Mixed", func(t *testing.T) {
			one := "one"
			two := "step2"
			val := []buildkite.DependsOnListItem{
				{String: &one},
				{
					DependsOnList: &buildkite.DependsOnListObject{
						Step: &two,
					},
				},
			}
			testVal := testDependsOn{
				DependsOn: buildkite.DependsOn{
					DependsOnList: &val,
				},
			}
			CheckResult(t, testVal, `{"depends_on":["one",{"step":"step2"}]}`)
		})
	})
}
