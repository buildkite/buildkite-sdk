package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	t.Run("Secrets", func(t *testing.T) {
		t.Run("StringArray", func(t *testing.T) {
			pipeline := buildkite.NewPipeline()
			pipeline.SetSecrets(&buildkite.Secrets{
				StringArray: []string{"MY_SECRET"},
			})

			result, err := pipeline.ToJSON()
			assert.NoError(t, err)
			assert.Equal(t, `{"secrets":["MY_SECRET"]}`, result)
		})

		t.Run("StringArray", func(t *testing.T) {
			pipeline := buildkite.NewPipeline()
			pipeline.SetSecrets(&buildkite.Secrets{
				Secrets: &buildkite.SecretsObject{"MY_SECRET": "API_TOKEN"},
			})

			result, err := pipeline.ToJSON()
			assert.NoError(t, err)
			assert.Equal(t, `{"secrets":{"MY_SECRET":"API_TOKEN"}}`, result)
		})
	})

	t.Run("AddAgent", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.AddAgent("foo", "bar")

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"agents":{"foo":"bar"}}`, result)
	})

	t.Run("AddEnvironmentVariable", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.AddEnvironmentVariable("FOO", "bar")

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"env":{"FOO":"bar"}}`, result)
	})

	t.Run("Notify", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.AddNotify(buildkite.BuildNotifyItem{
			NotifyEmail: &buildkite.NotifyEmail{
				Email: buildkite.Value("person@example.com"),
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"notify":[{"email":"person@example.com"}]}`, result)
	})

	t.Run("Steps", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.AddStep(buildkite.CommandStep{
			Command: &buildkite.CommandStepCommand{
				String: buildkite.Value("build.sh"),
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"steps":[{"command":"build.sh"}]}`, result)
	})
}
