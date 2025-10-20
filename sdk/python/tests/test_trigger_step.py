from buildkite_sdk import Pipeline, TriggerStep
import json

def test_trigger_step_simple():
    pipeline = Pipeline()
    pipeline.add_step(TriggerStep(
        trigger="deploy"
    ))

    expected = {"steps": [{"trigger": "deploy"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_trigger_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({ "trigger": "deploy" })

    expected = {"steps": [{"trigger": "deploy"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
