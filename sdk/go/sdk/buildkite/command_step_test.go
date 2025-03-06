package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandStep(t *testing.T) {
	t.Run("should create a simple command step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(CommandStep{
			Commands: []string{
				"command",
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "commands": [
                "command"
            ]
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should create a complex command step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(CommandStep{
			Commands: []string{
				"command",
			},
			Agents: map[string]interface{}{
				"npm": true,
			},
			AllowDependencyFailure: Value(true),
			ArtifactPaths:          []string{"path"},
			Branches:               []string{"branch"},
			Cache:                  CacheString("cache"),
			CancelOnBuildFailing:   Value(true),
			Concurrency:            Value(int64(1)),
			ConcurrencyGroup:       Value("concurrency-group"),
			ConcurrencyMethod:      &ConcurrencyMethod.Eager,
			DependsOn:              DependsOnString("test"),
			Env:                    map[string]interface{}{"FOO": "bar"},
			ID:                     Value("id"),
			Identifier:             Value("identifier"),
			If:                     Value("if"),
			Key:                    Value("key"),
			Label:                  Value("label"),
			Matrix:                 MatrixSimple{"mac", "windows"},
			Name:                   Value("name"),
			Notify: []StepNotify{
				{
					Slack: NotifySlackSimple("#channel"),
				},
			},
			Parallelism: Value(int64(1)),
			Plugins:     map[string]interface{}{},
			Priority:    Value(int64(1)),
			Retry: RetrySimple{
				Manual: Value(false),
			},
			Signature: &Signature{
				Algorithm:    "algorithm",
				SignedFields: []string{"one", "two"},
				Value:        "value",
			},
			Skip:             Value(true),
			SoftFail:         SoftFailSimple(true),
			TimeoutInMinutes: Value(int64(1)),
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "allow_dependency_failure": true,
            "branches": [
                "branch"
            ],
            "depends_on": "test",
            "id": "id",
            "identifier": "identifier",
            "if": "if",
            "key": "key",
            "label": "label",
            "name": "name",
            "agents": {
                "npm": true
            },
            "artifact_paths": [
                "path"
            ],
            "cache": "cache",
            "cancel_on_build_failing": true,
            "commands": [
                "command"
            ],
            "concurrency": 1,
            "concurrency_group": "concurrency-group",
            "concurrency_method": "eager",
            "env": {
                "FOO": "bar"
            },
            "matrix": [
                "mac",
                "windows"
            ],
            "notify": [
                {
                    "slack": "#channel"
                }
            ],
            "parallelism": 1,
            "plugins": {},
            "priority": 1,
            "retry": {
                "automatic": null,
                "manual": false
            },
            "signature": {
                "algorithm": "algorithm",
                "signed_fields": [
                    "one",
                    "two"
                ],
                "value": "value"
            },
            "skip": true,
            "soft_fail": true,
            "timeout_in_minutes": 1
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
