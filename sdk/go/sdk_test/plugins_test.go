package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testPluginsList struct {
	Plugins buildkite.PluginsList `json:"plugins"`
}

func TestPluginsList(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		value := "string"
		val := testPluginsList{
			Plugins: buildkite.PluginsList{
				{
					String: &value,
				},
			},
		}
		CheckResult(t, val, `{"plugins":["string"]}`)
	})

	t.Run("Object", func(t *testing.T) {
		val := testPluginsList{
			Plugins: buildkite.PluginsList{
				{
					PluginsList: &buildkite.PluginsListObject{
						"name": map[string]string{
							"foo": "bar",
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"plugins":[{"name":{"foo":"bar"}}]}`)
	})
}

type testPlugins struct {
	Plugins buildkite.Plugins `json:"plugins"`
}

func TestPlugins(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		plugin := "docker"
		list := buildkite.PluginsList{
			{
				String: &plugin,
			},
		}
		val := testPlugins{
			Plugins: buildkite.Plugins{
				PluginsList: &list,
			},
		}

		CheckResult(t, val, `{"plugins":["docker"]}`)
	})

	t.Run("Object", func(t *testing.T) {
		val := testPlugins{
			Plugins: buildkite.Plugins{
				PluginsObject: &buildkite.PluginsObject{
					"name": map[string]string{
						"foo": "bar",
					},
				},
			},
		}
		CheckResult(t, val, `{"plugins":{"name":{"foo":"bar"}}}`)
	})
}
