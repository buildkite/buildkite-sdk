from buildkite_sdk import Pipeline, WaitStep
import json

def test_no_argument_wait_step():
    pipeline = Pipeline()
    pipeline.add_step(WaitStep())

    expected = {"steps": [{"wait": "~"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_wait_step_simple():
    pipeline = Pipeline()
    pipeline.add_step(WaitStep(
        wait="My wait step"
    ))

    expected = {"steps": [{"wait": "My wait step"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_wait_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({ "wait": "My wait step" })

    expected = {"steps": [{"wait": "My wait step"}]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
