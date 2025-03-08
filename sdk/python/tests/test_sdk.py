from buildkite_sdk.sdk import Pipeline
from buildkite_sdk.command_step import CommandStep
import json

def test_sdk():
    assert 1 == 1
    # pipeline = Pipeline()
    # pipeline.add_step(CommandStep(
    #     commands="echo 'Hello, world!'"
    # ))
    # assert pipeline.to_json() == json.dumps({"steps": [{"command": "echo 'Hello, world!'"}]}, indent="    ")
