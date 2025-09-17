from buildkite_sdk import Pipeline, BlockStep
import json

class TestNestingTypes:
    # def test_block_step_string(self):
    #     pipeline = Pipeline(
    #         steps=[
    #             'block',
    #         ]
    #     )

    #     pipeline_json = pipeline.to_json_string()
    #     expected = json.dumps({'steps': ['block']})
    #     assert pipeline_json == expected

    # def test_block_step_label(self):
    #     pipeline = Pipeline(
    #         steps=[
    #             BlockStep(
    #                 block='label',
    #             ),
    #         ]
    #     )
    #     pipeline_json = pipeline.to_json_string()
    #     expected = json.dumps({'steps': [{'block': 'label'}]})
    #     assert pipeline_json == expected

    # def test_block_step_nested(self):
    #     pipeline = Pipeline(
    #         steps=[
    #             NestedBlockStep(
    #                 block=BlockStep(
    #                     block='label',
    #                 )
    #             ),
    #         ]
    #     )
    #     pipeline_json = pipeline.to_json_string()
    #     expected = json.dumps({'steps': [{'block': {'block': 'label'}}]})
    #     assert pipeline_json == expected

    # def test_block_step_type(self):
    #     pipeline = Pipeline(
    #         steps=[
    #             BlockStep(
    #                 type='block',
    #                 label='label',
    #             ),
    #         ]
    #     )
    #     pipeline_json = pipeline.to_json_string()
    #     expected = json.dumps({'steps': [{'label': 'label', 'type': 'block'}]})
    #     assert pipeline_json == expected

    def test_block_step_dict(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({ 'allow_dependency_failure': True, 'block': 'a label' })
            ]
        )

        pipeline_json = pipeline.to_json_string()
        expected = json.dumps({'steps': [{'allow_dependency_failure': True, 'block': 'a label'}]})
        assert pipeline_json == expected
