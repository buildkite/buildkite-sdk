from buildkite_sdk import Pipeline, CommandStep
import json

def test_command_step_multiple():
    pipeline = Pipeline()
    pipeline.add_step(CommandStep(
        commands=["echo 'Hello, world!'"]
    ))

    expected = {"steps": [{"commands": ["echo 'Hello, world!'"]}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_command_step_simple():
    pipeline = Pipeline()
    pipeline.add_step(CommandStep(
        command="echo 'Hello, world!'"
    ))

    expected = {"steps": [{"command": "echo 'Hello, world!'"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_command_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({ "commands": "echo 'Hello, world!'" })

    expected = {"steps": [{"commands": "echo 'Hello, world!'"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
