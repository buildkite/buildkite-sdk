from buildkite_sdk import Pipeline, PipelineDict
from .utils import TestRunner

class TestEnv(TestRunner):
    def test_class(self):
        pipeline = Pipeline(
            env={'FOO': 'bar'},
            steps=[]
        )
        self.validator.check_result(pipeline, {'steps': [], 'env': {'FOO': 'bar'}})

    def test_dict(self):
        expected: PipelineDict = {
            'env': {'FOO': 'bar'},
            'steps': []
        }
        pipeline = Pipeline.from_dict(expected)
        self.validator.check_result(pipeline, {'steps': [], 'env': {'FOO': 'bar'}})
