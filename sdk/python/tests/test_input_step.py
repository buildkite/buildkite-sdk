from buildkite_sdk import Pipeline, InputStep, TextField
import json

def test_simple_input_step():
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
