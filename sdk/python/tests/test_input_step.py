from buildkite_sdk import Pipeline, InputStep, TextField
import json

def test_input_step_simple():
    pipeline = Pipeline()
    pipeline.add_step(InputStep(
        input="My Input",
        fields=[
            TextField(
                key="my-input-key"
            )
        ],
    ))

    expected = {"steps": [{ "fields": [{ "key": "my-input-key" }], "input": "My Input" }]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")

def test_input_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({
        "fields": [{ "key": "my-input-key" }],
        "input": "My Input"
    })

    expected = {"steps": [{ "fields": [{ "key": "my-input-key" }], "input": "My Input" }]}
    assert pipeline.to_json() == json.dumps(expected, indent="    ")
