package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	t.Run("Skip", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Skip: &buildkite.CheckoutSkip{Bool: buildkite.Value(true)},
			}
			CheckResult(t, testVal, `{"skip":true}`)
		})

		t.Run("String", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Skip: &buildkite.CheckoutSkip{String: buildkite.Value("false")},
			}
			CheckResult(t, testVal, `{"skip":"false"}`)
		})
	})

	t.Run("Submodules", func(t *testing.T) {
		testVal := buildkite.Checkout{
			Submodules: &buildkite.CheckoutSubmodules{Bool: buildkite.Value(false)},
		}
		CheckResult(t, testVal, `{"submodules":false}`)
	})

	t.Run("Lfs", func(t *testing.T) {
		testVal := buildkite.Checkout{
			Lfs: &buildkite.CheckoutLfs{Bool: buildkite.Value(true)},
		}
		CheckResult(t, testVal, `{"lfs":true}`)
	})

	t.Run("Depth", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Depth: &buildkite.CheckoutDepth{Int: buildkite.Value(50)},
			}
			CheckResult(t, testVal, `{"depth":50}`)
		})

		t.Run("String", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Depth: &buildkite.CheckoutDepth{String: buildkite.Value("50")},
			}
			CheckResult(t, testVal, `{"depth":"50"}`)
		})
	})

	t.Run("CommitVerification", func(t *testing.T) {
		testVal := buildkite.Checkout{
			CommitVerification: buildkite.Value(buildkite.CheckoutCommitVerificationValues["strict"]),
		}
		CheckResult(t, testVal, `{"commit_verification":"strict"}`)
	})

	t.Run("Flags", func(t *testing.T) {
		testVal := buildkite.Checkout{
			Flags: &buildkite.CheckoutFlags{
				Clone: buildkite.Value("--depth 1 --single-branch"),
				Fetch: buildkite.Value("--prune --tags"),
			},
		}
		CheckResult(t, testVal, `{"flags":{"clone":"--depth 1 --single-branch","fetch":"--prune --tags"}}`)
	})

	t.Run("Sparse", func(t *testing.T) {
		t.Run("SinglePath", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Sparse: &buildkite.CheckoutSparse{
					Paths: &buildkite.CheckoutSparsePaths{
						CheckoutSparsePath: buildkite.Value("src/"),
					},
				},
			}
			CheckResult(t, testVal, `{"sparse":{"paths":"src/"}}`)
		})

		t.Run("PathList", func(t *testing.T) {
			testVal := buildkite.Checkout{
				Sparse: &buildkite.CheckoutSparse{
					Paths: &buildkite.CheckoutSparsePaths{
						StringArray: []string{"src/", "docs/"},
					},
				},
			}
			CheckResult(t, testVal, `{"sparse":{"paths":["src/","docs/"]}}`)
		})
	})

	t.Run("SshSecret", func(t *testing.T) {
		testVal := buildkite.Checkout{
			SshSecret: buildkite.Value("deploy_key"),
		}
		CheckResult(t, testVal, `{"ssh_secret":"deploy_key"}`)
	})

	t.Run("CommandStep", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.AddStep(buildkite.CommandStep{
			Command: &buildkite.CommandStepCommand{
				String: buildkite.Value("ls"),
			},
			Checkout: &buildkite.Checkout{
				Skip: &buildkite.CheckoutSkip{Bool: buildkite.Value(false)},
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"steps":[{"checkout":{"skip":false},"command":"ls"}]}`, result)
	})

	t.Run("Pipeline", func(t *testing.T) {
		pipeline := buildkite.NewPipeline()
		pipeline.Checkout = &buildkite.PipelineCheckout{
			Submodules: &buildkite.PipelineCheckoutSubmodules{Bool: buildkite.Value(false)},
			Depth:      &buildkite.PipelineCheckoutDepth{Int: buildkite.Value(50)},
		}
		pipeline.AddStep(buildkite.CommandStep{
			Command: &buildkite.CommandStepCommand{
				String: buildkite.Value("ls"),
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"checkout":{"depth":50,"submodules":false},"steps":[{"command":"ls"}]}`, result)
	})
}
