from buildkite_sdk.sdk import Pipeline
import json

def test_sdk():
    pipeline = Pipeline()
    pipeline.add_command_step({ "command": "echo 'Hello, world!'" })
    assert pipeline.to_json() == json.dumps({"steps": [{"command": "echo 'Hello, world!'"}]}, indent="    ")
