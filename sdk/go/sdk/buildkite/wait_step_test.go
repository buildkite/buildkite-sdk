package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaitStep(t *testing.T) {
	t.Run("should generate a simple wait step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(WaitStep{})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "wait": "~"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should generate an advanced wait step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(WaitStep{
			AllowDependencyFailure: Value(true),
			Branches: []string{
				"main",
			},
			ContinueOnFailure: Value(true),
			DependsOn: DependsOnStringArray{
				"build",
				"test",
			},
			ID:         Value("id"),
			Identifier: Value("identifier"),
			If:         Value("if"),
			Key:        Value("key"),
			Label:      Value("label"),
			Name:       Value("name"),
			Wait:       Value("wait"),
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "allow_dependency_failure": true,
            "branches": [
                "main"
            ],
            "depends_on": [
                "build",
                "test"
            ],
            "id": "id",
            "identifier": "identifier",
            "if": "if",
            "key": "key",
            "label": "label",
            "name": "name",
            "continue_on_failure": true,
            "wait": "wait"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
