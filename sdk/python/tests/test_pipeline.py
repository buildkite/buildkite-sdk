from buildkite_sdk import Pipeline, NotifyEmail
from .utils import TestRunner

class TestPipelineClass(TestRunner):
    def test_from_dict(self):
        pipeline = Pipeline.from_dict({"steps": [{"command": "run.sh"}]})
        self.validator.check_result(pipeline, {"steps": [{"command": "run.sh"}]})

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
        self.validator.check_result(pipeline, {"steps": [], "notify": [{"email": "person@example.com"}]})

    def test_add_step(self):
        pipeline = Pipeline()
        pipeline.add_step({"command": "run.sh"})
        self.validator.check_result(pipeline, {"steps": [{"command": "run.sh"}]})
