from buildkite_sdk import Pipeline, TriggerStep
import json

def test_simple_trigger_step():
    pipeline = Pipeline()
    pipeline.add_step(TriggerStep(
        trigger="deploy"
    ))

    expected = {"steps": [{"trigger": "deploy"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
