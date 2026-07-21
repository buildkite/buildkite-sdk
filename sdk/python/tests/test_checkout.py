from buildkite_sdk import Checkout, CommandStep, Pipeline

from .utils import TestRunner


class TestCommandStepCheckout(TestRunner):
    def test_checkout_skip(self):
        # The exact repro from buildkite-sdk#323.
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout={"skip": False})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": {"skip": False}}]}
        )

    def test_checkout_skip_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout={"skip": "true"})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": {"skip": "true"}}]}
        )

    def test_checkout_model(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="ls", checkout=Checkout(submodules=False))]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": {"submodules": False}}]}
        )

    def test_checkout_depth(self):
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout={"depth": 50})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": {"depth": 50}}]}
        )

    def test_checkout_depth_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout={"depth": "50"})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": {"depth": "50"}}]}
        )

    def test_checkout_flags(self):
        checkout = {
            "flags": {"clone": "--depth 1 --single-branch", "fetch": "--prune --tags"}
        }
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout=checkout)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": checkout}]}
        )

    def test_checkout_sparse_single_path(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="ls", checkout={"sparse": {"paths": "src/"}})]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "ls", "checkout": {"sparse": {"paths": "src/"}}}]},
        )

    def test_checkout_sparse_path_list(self):
        checkout = {"sparse": {"paths": ["src/", "docs/"]}}
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout=checkout)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": checkout}]}
        )

    def test_checkout_ssh_secret_and_commit_verification(self):
        checkout = {"ssh_secret": "deploy_key", "commit_verification": "strict"}
        pipeline = Pipeline(steps=[CommandStep(command="ls", checkout=checkout)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "ls", "checkout": checkout}]}
        )


class TestPipelineCheckout(TestRunner):
    def test_pipeline_level_checkout(self):
        pipeline = Pipeline(
            checkout={"submodules": False, "depth": 50},
            steps=[CommandStep(command="ls")],
        )
        self.validator.check_result(
            pipeline,
            {
                "checkout": {"submodules": False, "depth": 50},
                "steps": [{"command": "ls"}],
            },
        )

    def test_set_checkout(self):
        pipeline = Pipeline(steps=[CommandStep(command="ls")])
        pipeline.set_checkout({"lfs": True})
        self.validator.check_result(
            pipeline,
            {"checkout": {"lfs": True}, "steps": [{"command": "ls"}]},
        )

    def test_pipeline_checkout_drops_step_only_ssh_secret(self):
        pipeline = Pipeline.model_validate(
            {
                "checkout": {"skip": True, "ssh_secret": "deploy_key"},
                "steps": [{"command": "ls"}],
            }
        )
        assert pipeline.to_dict()["checkout"] == {"skip": True}
