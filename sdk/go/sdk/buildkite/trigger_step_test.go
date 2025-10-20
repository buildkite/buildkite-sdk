package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriggerStep(t *testing.T) {
	t.Run("should generate a simple trigger step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(TriggerStep{
			Trigger: "deploy",
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "trigger": "deploy"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should generate an advanced trigger step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(TriggerStep{
			Trigger:                "deploy",
			AllowDependencyFailure: Value(true),
			Async:                  Value(true),
			Branches:               []string{"main"},
			Build: &Build{
				Branch:   Value("branch"),
				Commit:   Value("commit"),
				Env:      map[string]interface{}{"foo": "bar"},
				Message:  Value("message"),
				MetaData: map[string]interface{}{"dom": "toretto"},
			},
			DependsOn:  DependsOnStringArray{"build", "test"},
			ID:         Value("id"),
			Identifier: Value("identifier"),
			If:         Value("if"),
			Key:        Value("key"),
			Label:      Value("label"),
			Name:       Value("name"),
			Skip:       Value(true),
			SoftFail:   SoftFailSimple(true),
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
            "skip": true,
            "soft_fail": true,
            "async": true,
            "build": {
                "branch": "branch",
                "commit": "commit",
                "env": {
                    "foo": "bar"
                },
                "message": "message",
                "meta_data": {
                    "dom": "toretto"
                }
            },
            "trigger": "deploy"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
