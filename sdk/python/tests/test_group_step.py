from buildkite_sdk import Pipeline, GroupStep, CommandStep, WaitStep
import json

def test_simple_group_step():
    pipeline = Pipeline()
    pipeline.add_step(GroupStep(
        group="my-group",
        steps=[
            CommandStep(commands="command1"),
            WaitStep(),
            CommandStep(commands="command2")
        ]
    ))

    expected = {"steps": [{
        "group": "my-group",
        "steps": [
            {"commands": "command1"},
            {"wait": "~"},
            {"commands": "command2"},
        ],
    }]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_group_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({
        "group":"my-group",
        "steps": [
            { "commands": "command1" },
            { "wait": "~" },
            { "commands": "command2" },
        ],
    })

    expected = {"steps": [{
        "group": "my-group",
        "steps": [
            {"commands": "command1"},
            {"wait": "~"},
            {"commands": "command2"},
        ],
    }]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
