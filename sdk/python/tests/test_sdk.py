from buildkite_sdk.sdk import Pipeline
from buildkite_sdk.command_step import CommandStep
import json

def test_sdk():
    pipeline = Pipeline()
    pipeline.add_agent("queue", "hosted")
    pipeline.add_environment_variable("FOO", "bar")
    pipeline.add_step({ "commands": "echo 'Hello, world!'" })
    assert pipeline.to_json() == json.dumps({
        "steps": [{"commands": "echo 'Hello, world!'"}],
        "agents": {
            "queue": "hosted"
        },
        "env": {
            "FOO": "bar"
        }
    }, indent="    ")
