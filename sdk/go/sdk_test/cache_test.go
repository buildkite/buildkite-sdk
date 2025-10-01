package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testCache struct {
	Cache buildkite.Cache `json:"cache"`
}

func TestCache(t *testing.T) {
	t.Run("CacheObject", func(t *testing.T) {
		t.Run("Paths", func(t *testing.T) {
			testVal := buildkite.CacheObject{
				Paths: []string{"one", "two"},
			}
			CheckResult(t, testVal, `{"paths":["one","two"]}`)
		})

		t.Run("Size", func(t *testing.T) {
			size := "string"
			testVal := buildkite.CacheObject{
				Size: &size,
			}
			CheckResult(t, testVal, `{"size":"string"}`)
		})

		t.Run("Name", func(t *testing.T) {
			name := "name"
			testVal := buildkite.CacheObject{
				Name: &name,
			}
			CheckResult(t, testVal, `{"name":"name"}`)
		})

		t.Run("All", func(t *testing.T) {
			name := "name"
			size := "string"
			testVal := buildkite.CacheObject{
				Paths: []string{"one", "two"},
				Size:  &size,
				Name:  &name,
			}
			CheckResult(t, testVal, `{"name":"name","paths":["one","two"],"size":"string"}`)
		})
	})

	t.Run("Cache", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			val := "string"
			testVal := testCache{
				Cache: buildkite.Cache{
					String: &val,
				},
			}
			CheckResult(t, testVal, `{"cache":"string"}`)
		})

		t.Run("StringArray", func(t *testing.T) {
			val := []string{"one", "two"}
			testVal := testCache{
				Cache: buildkite.Cache{
					StringArray: val,
				},
			}
			CheckResult(t, testVal, `{"cache":["one","two"]}`)
		})

		t.Run("Object", func(t *testing.T) {
			size := "string"
			testVal := testCache{
				Cache: buildkite.Cache{
					Cache: &buildkite.CacheObject{
						Size: &size,
					},
				},
			}

			CheckResult(t, testVal, `{"cache":{"size":"string"}}`)
		})
	})
}
