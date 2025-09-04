package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go2/sdk"
)

func TestNotifyGithubCheck(t *testing.T) {
	githubCheck := map[string]interface{}{"foo": "bar"}
	val := buildkite.NotifyGithubCheck{
		GithubCheck: &githubCheck,
	}
	CheckResult(t, val, `{"github_check":{"foo":"bar"}}`)
}
