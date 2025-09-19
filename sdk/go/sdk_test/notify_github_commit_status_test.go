package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

func TestNotifyGithubCommitStatus(t *testing.T) {
	t.Run("NotifyGithubCommitStatusGithubCommitStatus", func(t *testing.T) {
		context := "name"
		val := buildkite.NotifyGithubCommitStatusGithubCommitStatus{
			Context: &context,
		}
		CheckResult(t, val, `{"context":"name"}`)
	})

	t.Run("NotifyGithubCommitStatus", func(t *testing.T) {
		t.Run("GithubCommitStatus", func(t *testing.T) {
			context := "name"
			val := buildkite.NotifyGithubCommitStatus{
				GithubCommitStatus: &buildkite.NotifyGithubCommitStatusGithubCommitStatus{
					Context: &context,
				},
			}
			CheckResult(t, val, `{"github_commit_status":{"context":"name"}}`)
		})

		t.Run("If", func(t *testing.T) {
			ifVal := "string"
			val := buildkite.NotifyGithubCommitStatus{
				If: &ifVal,
			}
			CheckResult(t, val, `{"if":"string"}`)
		})

		t.Run("GithubCommitStatus", func(t *testing.T) {
			context := "name"
			ifVal := "string"
			val := buildkite.NotifyGithubCommitStatus{
				If: &ifVal,
				GithubCommitStatus: &buildkite.NotifyGithubCommitStatusGithubCommitStatus{
					Context: &context,
				},
			}
			CheckResult(t, val, `{"github_commit_status":{"context":"name"},"if":"string"}`)
		})
	})
}
