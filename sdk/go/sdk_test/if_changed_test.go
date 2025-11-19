package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testIfChanged struct {
	IfChanged buildkite.IfChanged `json:"if_changed"`
}

func TestIfChanged(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		testVal := testIfChanged{
			IfChanged: buildkite.IfChanged{
				String: buildkite.Value("*.txt"),
			},
		}
		CheckResult(t, testVal, `{"if_changed":"*.txt"}`)
	})

	t.Run("StringArray", func(t *testing.T) {
		testVal := testIfChanged{
			IfChanged: buildkite.IfChanged{
				StringArray: []string{"*.txt"},
			},
		}
		CheckResult(t, testVal, `{"if_changed":["*.txt"]}`)
	})

	t.Run("ObjectString", func(t *testing.T) {
		testVal := testIfChanged{
			IfChanged: buildkite.IfChanged{
				IfChanged: &buildkite.IfChangedObject{
					Include: &buildkite.IfChangedObjectInclude{
						String: buildkite.Value("*.txt"),
					},
					Exclude: &buildkite.IfChangedObjectExclude{
						String: buildkite.Value("*.md"),
					},
				},
			},
		}
		CheckResult(t, testVal, `{"if_changed":{"exclude":"*.md","include":"*.txt"}}`)
	})

	t.Run("ObjectStringArray", func(t *testing.T) {
		testVal := testIfChanged{
			IfChanged: buildkite.IfChanged{
				IfChanged: &buildkite.IfChangedObject{
					Include: &buildkite.IfChangedObjectInclude{
						StringArray: []string{"*.txt"},
					},
					Exclude: &buildkite.IfChangedObjectExclude{
						StringArray: []string{"*.md"},
					},
				},
			},
		}
		CheckResult(t, testVal, `{"if_changed":{"exclude":["*.md"],"include":["*.txt"]}}`)
	})
}
