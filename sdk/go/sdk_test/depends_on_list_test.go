package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testDependsOnListObjectAllowFailure struct {
	DependsOnListObjectAllowFailure buildkite.DependsOnListObjectAllowFailure `json:"allow_failure"`
}

type testDependsOnList struct {
	DependsOn buildkite.DependsOnList `json:"depends_on"`
}

func TestDependsOnList(t *testing.T) {
	t.Run("DependsOnListObjectAllowFailure", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			val := "string"
			testVal := testDependsOnListObjectAllowFailure{
				DependsOnListObjectAllowFailure: buildkite.DependsOnListObjectAllowFailure{
					String: &val,
				},
			}
			CheckResult(t, testVal, `{"allow_failure":"string"}`)
		})

		t.Run("Bool", func(t *testing.T) {
			val := true
			testVal := testDependsOnListObjectAllowFailure{
				DependsOnListObjectAllowFailure: buildkite.DependsOnListObjectAllowFailure{
					Bool: &val,
				},
			}
			CheckResult(t, testVal, `{"allow_failure":true}`)
		})
	})

	t.Run("DependsOnListObject", func(t *testing.T) {
		t.Run("Step", func(t *testing.T) {
			step := "step"
			testVal := buildkite.DependsOnListObject{
				Step: &step,
			}
			CheckResult(t, testVal, `{"step":"step"}`)
		})

		t.Run("AllowFailure", func(t *testing.T) {
			val := true
			testVal := buildkite.DependsOnListObject{
				AllowFailure: &buildkite.DependsOnListObjectAllowFailure{
					Bool: &val,
				},
			}
			CheckResult(t, testVal, `{"allow_failure":true}`)
		})

		t.Run("All", func(t *testing.T) {
			step := "step"
			val := true
			testVal := buildkite.DependsOnListObject{
				Step: &step,
				AllowFailure: &buildkite.DependsOnListObjectAllowFailure{
					Bool: &val,
				},
			}
			CheckResult(t, testVal, `{"allow_failure":true,"step":"step"}`)
		})
	})

	t.Run("DependsOnList", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			one := "one"
			two := "two"
			testVal := testDependsOnList{
				DependsOn: []buildkite.DependsOnListUnion{
					{String: &one},
					{String: &two},
				},
			}
			CheckResult(t, testVal, `{"depends_on":["one","two"]}`)
		})

		t.Run("Object", func(t *testing.T) {
			one := "step1"
			two := "step2"
			testVal := testDependsOnList{
				DependsOn: []buildkite.DependsOnListUnion{
					{
						DependsOnList: &buildkite.DependsOnListObject{
							Step: &one,
						},
					},
					{
						DependsOnList: &buildkite.DependsOnListObject{
							Step: &two,
						},
					},
				},
			}
			CheckResult(t, testVal, `{"depends_on":[{"step":"step1"},{"step":"step2"}]}`)
		})

		t.Run("Mixed", func(t *testing.T) {
			one := "one"
			two := "step2"
			testVal := testDependsOnList{
				DependsOn: []buildkite.DependsOnListUnion{
					{String: &one},
					{
						DependsOnList: &buildkite.DependsOnListObject{
							Step: &two,
						},
					},
				},
			}
			CheckResult(t, testVal, `{"depends_on":["one",{"step":"step2"}]}`)
		})
	})
}
