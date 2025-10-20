package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockStep(t *testing.T) {
	t.Run("should create a simple block step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(BlockStep{
			Block: Value("block"),
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "block": "block"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should create a complex block step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(BlockStep{
			AllowDependencyFailure: Value(true),
			BlockedState:           &BlockedState.Failed,
			Branches: []string{
				"main",
			},
			DependsOn: DependsOnStringArray{
				"build",
				"test",
			},
			Fields: []Field{
				InputTextField{
					Text: Value("Text Input"),
					Key:  Value("text-input"),
				},
			},
			ID:         Value("id"),
			Identifier: Value("identifier"),
			If:         Value("if"),
			Key:        Value("key"),
			Label:      Value("label"),
			Name:       Value("name"),
			Prompt:     Value("prompt"),
			Block:      Value("block"),
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "allow_dependency_failure": true,
            "block": "block",
            "blocked_state": "failed",
            "branches": [
                "main"
            ],
            "depends_on": [
                "build",
                "test"
            ],
            "fields": [
                {
                    "text": "Text Input",
                    "key": "text-input"
                }
            ],
            "id": "id",
            "identifier": "identifier",
            "if": "if",
            "key": "key",
            "label": "label",
            "name": "name",
            "prompt": "prompt"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
