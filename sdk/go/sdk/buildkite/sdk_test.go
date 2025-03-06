package buildkite

import (
	"testing"

	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"
)

func TestAddCommandStep(t *testing.T) {
	pipeline := NewPipeline()

	label := "some-label"
	command := "echo 'Hello, world!'"
	pipeline.AddCommandStep(schema.CommandStep{
		Label: &label,
		Command: &schema.CommandUnion{
			String: &command,
		},
	})

	actual, _ := pipeline.ToJSON()
	expected := `{
    "steps": [
        {
            "label": "some-label",
            "command": "echo 'Hello, world!'"
        }
    ]
}`
	if actual != expected {
		t.Errorf("Expected '%s' to be '%s'", actual, expected)
	}
}
