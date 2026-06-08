package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCondition(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		condition, err := buildkite.Condition(`build.branch == "main"`)
		require.NoError(t, err)
		require.NotNil(t, condition)
		assert.Equal(t, `build.branch == "main"`, *condition)
	})

	t.Run("Invalid", func(t *testing.T) {
		condition, err := buildkite.Condition(`build.branch == `)
		require.Error(t, err)
		assert.Nil(t, condition)
		assert.ErrorContains(t, err, "no prefix parse function for EOF found")
	})

	t.Run("RejectsStepVariables", func(t *testing.T) {
		condition, err := buildkite.Condition(`step.outcome == "passed"`)
		require.Error(t, err)
		assert.Nil(t, condition)
		assert.ErrorContains(t, err, `step variables are not available for entry point "build_condition"`)
	})
}

func TestMustCondition(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		condition := buildkite.MustCondition(`build.branch == "main"`)
		require.NotNil(t, condition)
		assert.Equal(t, `build.branch == "main"`, *condition)
	})

	t.Run("Invalid", func(t *testing.T) {
		assert.Panics(t, func() {
			buildkite.MustCondition(`build.branch == `)
		})
	})
}

func TestValidateConditionals(t *testing.T) {
	t.Run("ValidatesStepAndNotificationConditionals", func(t *testing.T) {
		pipeline := buildkite.Pipeline{
			Notify: &buildkite.BuildNotify{
				{
					NotifyEmail: &buildkite.NotifyEmail{
						Email: buildkite.Value("alerts@example.com"),
						If:    buildkite.Value(`build.branch == "main"`),
					},
				},
			},
			Steps: &buildkite.PipelineSteps{
				{
					CommandStep: &buildkite.CommandStep{
						If: buildkite.MustCondition(`build.branch == "main"`),
						Notify: &buildkite.CommandStepNotify{
							{
								NotifySlack: &buildkite.NotifySlack{
									If: buildkite.Value(`step.outcome == "soft_failed"`),
									Slack: &buildkite.NotifySlackSlack{
										String: buildkite.Value("#builds"),
									},
								},
							},
						},
					},
				},
				{
					GroupStep: &buildkite.GroupStep{
						If: buildkite.MustCondition(`build.branch == "main"`),
						Notify: &buildkite.BuildNotify{
							{
								NotifyWebhook: &buildkite.NotifyWebhook{
									If:      buildkite.Value(`step.outcome == "hard_failed"`),
									Webhook: buildkite.Value("https://example.com/builds"),
								},
							},
						},
						Steps: &buildkite.GroupSteps{
							{
								NestedTriggerStep: &buildkite.NestedTriggerStep{
									Trigger: &buildkite.TriggerStep{
										If: buildkite.MustCondition(`build.tag == null`),
									},
								},
							},
						},
					},
				},
			},
		}

		require.NoError(t, buildkite.ValidateConditionals(pipeline))
	})

	t.Run("ReportsEntrypointSpecificValidationErrors", func(t *testing.T) {
		pipeline := buildkite.Pipeline{
			Steps: &buildkite.PipelineSteps{
				{
					CommandStep: &buildkite.CommandStep{
						If: buildkite.Value(`step.label == "Deploy"`),
					},
				},
				{
					GroupStep: &buildkite.GroupStep{
						Notify: &buildkite.BuildNotify{
							{
								NotifyEmail: &buildkite.NotifyEmail{
									Email: buildkite.Value("alerts@example.com"),
									If:    buildkite.Value(`step.outcome == "failed"`),
								},
							},
						},
					},
				},
			},
		}

		err := buildkite.ValidateConditionals(pipeline)
		require.Error(t, err)
		assert.ErrorContains(t, err, `steps[0].if: validation: step variables are not available for entry point "build_condition"`)
		assert.ErrorContains(t, err, "steps[1].notify[0].if: validation: \"failed\" is not a valid `step.outcome`")
	})

	t.Run("MatchesSerializedStepArmOrder", func(t *testing.T) {
		pipeline := buildkite.Pipeline{
			Steps: &buildkite.PipelineSteps{
				{
					CommandStep: &buildkite.CommandStep{
						If: buildkite.Value(`build.branch == "main"`),
					},
					NestedBlockStep: &buildkite.NestedBlockStep{
						Block: &buildkite.BlockStep{
							If: buildkite.Value(`build.branch == `),
						},
					},
				},
			},
		}

		err := buildkite.ValidateConditionals(pipeline)
		require.Error(t, err)
		assert.ErrorContains(t, err, `steps[0].block.if: parse: no prefix parse function for EOF found`)
		assert.NotContains(t, err.Error(), `steps[0].if`)
	})

	t.Run("MatchesSerializedNotificationArmOrder", func(t *testing.T) {
		pipeline := buildkite.Pipeline{
			Notify: &buildkite.BuildNotify{
				{
					NotifyBasecamp: &buildkite.NotifyBasecamp{
						BasecampCampfire: buildkite.Value("general"),
						If:               buildkite.Value(`build.branch == "main"`),
					},
					NotifyEmail: &buildkite.NotifyEmail{
						Email: buildkite.Value("alerts@example.com"),
						If:    buildkite.Value(`build.branch == `),
					},
				},
			},
		}

		err := buildkite.ValidateConditionals(pipeline)
		require.Error(t, err)
		assert.ErrorContains(t, err, `notify[0].if: parse: no prefix parse function for EOF found`)
	})
}
