from buildkite_sdk import (
    Pipeline,
    NotifyEmail,
    BlockStepArgs,
    CommandStepArgs,
    GroupStepArgs,
    InputStepArgs,
    TriggerStepArgs,
    WaitStepArgs,
)
from .utils import TestRunner


class TestPipelineClass(TestRunner):
    def test_from_dict(self):
        pipeline = Pipeline.from_dict({"steps": [{"command": "run.sh"}]})
        self.validator.check_result(pipeline, {"steps": [{"command": "run.sh"}]})

    def test_set_secrets_list(self):
        pipeline = Pipeline()
        pipeline.set_secrets(["MY_SECRET"])
        self.validator.check_result(pipeline, {"steps": [], "secrets": ["MY_SECRET"]})

    def test_set_secrets_object(self):
        pipeline = Pipeline()
        pipeline.set_secrets({"MY_SECRET": "API_TOKEN"})
        self.validator.check_result(pipeline, {"steps": [], "secrets": {"MY_SECRET": "API_TOKEN"}})

    def test_add_agent_list(self):
        pipeline = Pipeline(agents=[])
        pipeline.add_agent("foo", "bar")
        self.validator.check_result(pipeline, {"steps": [], "agents": ["foo=bar"]})

    def test_add_agent_object(self):
        pipeline = Pipeline()
        pipeline.add_agent("foo", "bar")
        self.validator.check_result(pipeline, {"steps": [], "agents": {"foo": "bar"}})

    def test_add_environment_variable(self):
        pipeline = Pipeline()
        pipeline.add_environment_variable("FOO", "bar")
        self.validator.check_result(pipeline, {"steps": [], "env": {"FOO": "bar"}})

    def test_add_notify(self):
        pipeline = Pipeline()
        pipeline.add_notify([NotifyEmail(email="person@example.com")])
        self.validator.check_result(
            pipeline, {"steps": [], "notify": [{"email": "person@example.com"}]}
        )

    def test_add_step(self):
        pipeline = Pipeline()
        pipeline.add_step({"command": "run.sh"})
        self.validator.check_result(pipeline, {"steps": [{"command": "run.sh"}]})

    def test_add_multiple_steps(self):
        block_step: BlockStepArgs = {"block": "block label"}
        command_step: CommandStepArgs = {"command": "run.sh"}
        group_step: GroupStepArgs = {
            "group": "group label",
            "steps": [{"command": "group.sh"}],
        }
        input_step: InputStepArgs = {"input": "input label"}
        trigger_step: TriggerStepArgs = {"trigger": "trigger label"}
        wait_step: WaitStepArgs = {"wait": "wait label"}

        pipeline = Pipeline()
        pipeline.add_step(block_step)
        pipeline.add_step(command_step)
        pipeline.add_step(group_step)
        pipeline.add_step(input_step)
        pipeline.add_step(trigger_step)
        pipeline.add_step(wait_step)

        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    block_step,
                    command_step,
                    group_step,
                    input_step,
                    trigger_step,
                    wait_step,
                ]
            },
        )
