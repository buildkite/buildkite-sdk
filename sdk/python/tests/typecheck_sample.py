"""Usage that must pass Pyright; checked by test_typecheck.py, never executed.

Not prefixed with test_ so pytest does not collect it.
"""

from buildkite_sdk import Checkout, CommandStep, Pipeline

# The Pyright repro from https://github.com/buildkite/buildkite-sdk/issues/323
Pipeline(steps=[CommandStep(command="ls", checkout={"skip": False})])

# Model form
Pipeline(steps=[CommandStep(command="ls", checkout=Checkout(submodules=False))])

# Nested dict forms
CommandStep(
    command="ls",
    checkout={
        "depth": 50,
        "flags": {"clone": "--depth 1 --single-branch"},
        "sparse": {"paths": ["src/", "docs/"]},
        "ssh_secret": "deploy_key",
    },
)

# Pipeline-level checkout; ssh_secret is step-level only and is not part of
# PipelineCheckout, so passing it here would fail this type check.
Pipeline(
    checkout={"submodules": False, "depth": 50},
    steps=[CommandStep(command="ls")],
)
Pipeline(steps=[CommandStep(command="ls")]).set_checkout({"lfs": True})
