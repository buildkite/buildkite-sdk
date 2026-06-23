package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestNotifyGithubCheck(t *testing.T) {
	name := "my-check"
	val := buildkite.NotifyGithubCheck{
		GithubCheck: &buildkite.NotifyGithubCheckGithubCheck{
			Name: &name,
		},
	}
	CheckResult(t, val, `{"github_check":{"name":"my-check"}}`)
}
