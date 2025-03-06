package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCommandStep(t *testing.T) {
	pipeline := NewPipeline()

	pipeline.AddStep(CommandStep{
		Label: Value("some-label"),
		Commands: []string{
			"echo 'Hello, world!'",
		},
	})

	actual, err := pipeline.ToJSON()
	assert.NoError(t, err)

	expected := `{
    "steps": [
        {
            "label": "some-label",
            "commands": [
                "echo 'Hello, world!'"
            ]
        }
    ]
}`
	if actual != expected {
		t.Errorf("Expected '%s' to be '%s'", actual, expected)
	}
}
