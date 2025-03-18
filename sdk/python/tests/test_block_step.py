from buildkite_sdk import Pipeline, BlockStep
import json

def test_block_step_simple():
    pipeline = Pipeline()
    pipeline.add_step(BlockStep(
        block="my block step"
    ))

    assert pipeline.to_json() == json.dumps({"steps": [{"block": "my block step"}]}, indent="    ")

def test_block_step_typed_dict():
    pipeline = Pipeline()
    pipeline.add_step({ "block": "my block step" })

    assert pipeline.to_json() == json.dumps({"steps": [{"block": "my block step"}]}, indent="    ")
