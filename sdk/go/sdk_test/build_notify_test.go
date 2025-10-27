package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testBuildNotify struct {
	Notify buildkite.BuildNotify `json:"notify"`
}

func TestBuildNotify(t *testing.T) {
	t.Run("NotifySimple", func(t *testing.T) {
		value := buildkite.NotifySimpleValues["github_check"]
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifySimple: &value,
				},
			},
		}
		CheckResult(t, val, `{"notify":["github_check"]}`)
	})

	t.Run("NotifyEmail", func(t *testing.T) {
		email := "email"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyEmail: &buildkite.NotifyEmail{
						Email: &email,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"email":"email"}]}`)
	})

	t.Run("NotifyBasecamp", func(t *testing.T) {
		value := "string"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyBasecamp: &buildkite.NotifyBasecamp{
						BasecampCampfire: &value,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"basecamp_campfire":"string"}]}`)
	})

	t.Run("NotifySlack", func(t *testing.T) {
		channel := "#general"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifySlack: &buildkite.NotifySlack{
						Slack: &buildkite.NotifySlackSlack{
							String: &channel,
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"slack":"#general"}]}`)
	})

	t.Run("NotifyWebhook", func(t *testing.T) {
		webhook := "url"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyWebhook: &buildkite.NotifyWebhook{
						Webhook: &webhook,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"webhook":"url"}]}`)
	})

	t.Run("NotifyPagerduty", func(t *testing.T) {
		changeEvent := "event"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyPagerduty: &buildkite.NotifyPagerduty{
						PagerdutyChangeEvent: &changeEvent,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"pagerduty_change_event":"event"}]}`)
	})

	t.Run("NotifyGithubCommitStatus", func(t *testing.T) {
		context := "ctx"
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyGithubCommitStatus: &buildkite.NotifyGithubCommitStatus{
						GithubCommitStatus: &buildkite.NotifyGithubCommitStatusGithubCommitStatus{
							Context: &context,
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"github_commit_status":{"context":"ctx"}}]}`)
	})

	t.Run("NotifyGithubCheck", func(t *testing.T) {
		value := map[string]interface{}{"foo": "bar"}
		val := testBuildNotify{
			Notify: []buildkite.BuildNotifyItem{
				{
					NotifyGithubCheck: &buildkite.NotifyGithubCheck{
						GithubCheck: &value,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"github_check":{"foo":"bar"}}]}`)
	})
}
