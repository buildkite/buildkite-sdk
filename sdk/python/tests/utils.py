from buildkite_sdk import Pipeline
from typing import Dict, Any
from jsoncomparison import Compare, NO_DIFF
import unittest
import requests
import jsonschema
import json
import yaml


class PipelineValidator:
    schema: Any

    def __init__(self) -> None:
        response = requests.get(
            "https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json"
        )
        response.raise_for_status()
        self.schema = response.json()

    def validate_pipeline_against_schema(self, pipeline_json: Dict[str, Any]) -> bool:
        try:
            jsonschema.validate(instance=pipeline_json, schema=self.schema)
            return True
        except Exception as e:
            return False

    def check_result(self, pipeline: Pipeline, expected: Dict[str, Any]):
        actual = pipeline.to_dict()

        is_valid = self.validate_pipeline_against_schema(actual)
        pipeline_string = pipeline.to_json()
        assert is_valid == True, f"Pipeline is not valid: {pipeline_string}"

        comparison_result = Compare().check(expected, actual)
        assert comparison_result == NO_DIFF, (
            f"Differences found: {json.dumps(comparison_result, indent=4)}"
        )

        expectedYaml = yaml.dump(expected)
        actualYaml = pipeline.to_yaml()
        assert expectedYaml == actualYaml, (
            f"YAML results do not match:\n{expectedYaml}\n{actualYaml}"
        )


class TestRunner(unittest.TestCase):
    validator = PipelineValidator()
