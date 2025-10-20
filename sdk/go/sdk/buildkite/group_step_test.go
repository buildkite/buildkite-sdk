package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupStep(t *testing.T) {
	t.Run("should create a simple group pipeline", func(t *testing.T) {
		pipeline := NewPipeline()

		cmdStep := CommandStep{
			Commands: []string{"command"},
		}

		pipeline.AddStep(GroupStep{
			Steps: []GroupStepStep{
				cmdStep,
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "group": "~",
            "steps": [
                {
                    "commands": [
                        "command"
                    ]
                }
            ]
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should create a complext group pipeline", func(t *testing.T) {
		pipeline := NewPipeline()

		cmdStep := CommandStep{
			Commands: []string{"command"},
		}

		pipeline.AddStep(GroupStep{
			Steps: []GroupStepStep{
				cmdStep,
			},
			Group:                  Value("group"),
			If:                     Value("if"),
			Key:                    Value("key"),
			ID:                     Value("id"),
			Identifier:             Value("identifier"),
			Label:                  Value("label"),
			Name:                   Value("name"),
			Skip:                   Value(true),
			AllowDependencyFailure: Value(false),
			DependsOn:              DependsOnString("test"),
			Notify: []StepNotify{
				{
					Slack: NotifySlackSimple("#channel"),
				},
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "allow_dependency_failure": false,
            "depends_on": "test",
            "id": "id",
            "identifier": "identifier",
            "if": "if",
            "key": "key",
            "label": "label",
            "name": "name",
            "notify": [
                {
                    "slack": "#channel"
                }
            ],
            "skip": true,
            "group": "group",
            "steps": [
                {
                    "commands": [
                        "command"
                    ]
                }
            ]
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should handle multiple steps", func(t *testing.T) {
		pipeline := NewPipeline()

		cmdStep1 := CommandStep{
			Commands: []string{"command1"},
		}

		waitStep := WaitStep{}

		cmdStep2 := CommandStep{
			Commands: []string{"command2"},
		}

		pipeline.AddStep(GroupStep{
			Steps: []GroupStepStep{
				cmdStep1,
				waitStep,
				cmdStep2,
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "group": "~",
            "steps": [
                {
                    "commands": [
                        "command1"
                    ]
                },
                {
                    "wait": "~"
                },
                {
                    "commands": [
                        "command2"
                    ]
                }
            ]
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
