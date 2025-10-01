from buildkite_sdk import Pipeline, BlockStep, BlockStepDict, NestedBlockStep, TextField, SelectField, SelectFieldOption, DependsOnListObject, GroupStep
from .utils import TestRunner

class TestBlockStepNestingTypesClass(TestRunner):
    def test_block_step_string(self):
        pipeline = Pipeline(
            steps=[
                'block',
            ]
        )
        self.validator.check_result(pipeline, {'steps': ['block']})

    def test_block_step_label(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='label',
                ),
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label'}]})

    def test_block_step_nested(self):
        pipeline = Pipeline(
            steps=[
                NestedBlockStep(
                    block=BlockStep(
                        block='label',
                    )
                ),
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': {'block': 'label'}}]})

    def test_block_step_type(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    type='block',
                    label='label',
                ),
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'label': 'label', 'type': 'block'}]})

class TestBlockStepNestingTypesDict(TestRunner):
    def test_block_step_label(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'label'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label'}]})

    def test_block_step_nested(self):
        pipeline = Pipeline(
            steps=[
                NestedBlockStep.from_dict({ 'block': { 'block': 'label' }})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': {'block': 'label'}}]})

    def test_block_step_type(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'type': 'block', 'label': 'label'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'label': 'label', 'type': 'block'}]})

class TestBlockStepClass(TestRunner):
    def test_branches(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='label', branches='branch')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'branches': 'branch'}]})

    def test_id(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='label', id='id')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'id': 'id'}]})

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='label', identifier='identifier')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'identifier': 'identifier'}]})

    def test_prompt(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='label', prompt='prompt')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'prompt': 'prompt'}]})

    def test_fields(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='A label',
                    prompt='A prompt',
                    fields=[
                        TextField(text='Field 1', key='field-1'),
                        TextField(text='Field 2', key='field-2', required=False, default='Field 2 Default', hint='Field 2 Hint'),
                        SelectField(
                            select='Select 1',
                            key='select-1',
                            options=[
                                SelectFieldOption(label='Select 1 Option 1', value='select-1-option-1'),
                                SelectFieldOption(label='Select 1 Option 2', value='select-1-option-2')
                            ]
                        ),
                        SelectField(
                            select='Select 2',
                            key='select-2',
                            hint='Select 2 Hint',
                            required=False,
                            default='select-2-option-1',
                            options=[
                                SelectFieldOption(label='Select 2 Option 1', value='select-2-option-1'),
                            ],
                        ),
                    ]
                )
            ]
        )
        expected = {
            'block': 'A label',
            'prompt': 'A prompt',
            'fields': [
                {
                    'text': 'Field 1',
                    'key': 'field-1'
                },
                {
                    'text': 'Field 2',
                    'key': 'field-2',
                    'required': False,
                    'default': 'Field 2 Default',
                    'hint': 'Field 2 Hint'
                },
                {
                    'select': 'Select 1',
                    'key': 'select-1',
                    'options': [
                        {
                            'label': 'Select 1 Option 1',
                            'value': 'select-1-option-1'
                        },
                        {
                            'label': 'Select 1 Option 2',
                            'value': 'select-1-option-2'
                        }
                    ]
                },
                {
                    'select': 'Select 2',
                    'key': 'select-2',
                    'hint': 'Select 2 Hint',
                    'required': False,
                    'default': 'select-2-option-1',
                    'options': [
                        {
                            'label': 'Select 2 Option 1',
                            'value': 'select-2-option-1'
                        }
                    ]
                }
            ]
        }
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', pipeline_if='build.message !~ /skip tests/')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'if': 'build.message !~ /skip tests/'}]})

    def test_key(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', key='key')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'key': 'key'}]})

    def test_depends_on_string(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', depends_on='step')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': 'step'}]})

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', depends_on=['one','two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': ['one','two']}]})

    def test_depends_on_object_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='a label',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        DependsOnListObject(step='two'),
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': [{'step': 'one', 'allow_failure': True},{'step':'two'}]}]})

    def test_depends_on_mixed_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='a label',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        'two',
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': [{'step': 'one', 'allow_failure': True},'two']}]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', allow_dependency_failure=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allow_dependency_failure': True}]})

    def test_multiple_fields(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='a label',
                    fields=[
                        SelectField(
                            select='Multiple fields',
                            key='multiple-fields',
                            multiple=True,
                            options=[
                                SelectFieldOption(label='Option 1', value='option-1'),
                                SelectFieldOption(label='Option 2', value='option-2')
                            ]
                        ),
                    ],
                )
            ]
        )
        expected = {'steps': [{
            'block': 'a label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'options': [
                        {'label': 'Option 1', 'value': 'option-1'},
                        {'label': 'Option 2', 'value': 'option-2'},
                    ],
                },
            ],
        }]}
        self.validator.check_result(pipeline, expected)

    def test_multiple_fields_with_default(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(
                    block='a label',
                    fields=[
                        SelectField(
                            select='Multiple fields',
                            key='multiple-fields',
                            multiple=True,
                            default=['option-1', 'option-2'],
                            options=[
                                SelectFieldOption(label='Option 1', value='option-1'),
                                SelectFieldOption(label='Option 2', value='option-2')
                            ]
                        ),
                    ],
                )
            ]
        )
        expected = {'steps': [{
            'block': 'a label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'default': ['option-1', 'option-2'],
                    'options': [
                        {'label': 'Option 1', 'value': 'option-1'},
                        {'label': 'Option 2', 'value': 'option-2'},
                    ],
                },
            ],
        }]}
        self.validator.check_result(pipeline, expected)

    def test_allowed_teams_string(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', allowed_teams='team')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allowed_teams': 'team'}]})

    def test_allowed_teams_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep(block='a label', allowed_teams=['one', 'two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allowed_teams': ['one', 'two']}]})

    def test_group_nesting(self):
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    steps=[
                        BlockStep(block='a label')
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'group': 'Tests', 'steps': [{'block': 'a label'}]}]})

class TestBlockStepDict(TestRunner):
    def test_branches(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'label', 'branches': 'branch'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'branches': 'branch'}]})

    def test_id(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'label', 'id': 'id'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'id': 'id'}]})

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'label', 'identifier': 'identifier'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'identifier': 'identifier'}]})

    def test_prompt(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'label', 'prompt': 'prompt'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'label', 'prompt': 'prompt'}]})

    def test_fields(self):
        expected: BlockStepDict = {
            'block': 'A label',
            'prompt': 'A prompt',
            'fields': [
                {
                    'text': 'Field 1',
                    'key': 'field-1'
                },
                {
                    'text': 'Field 2',
                    'key': 'field-2',
                    'required': False,
                    'default': 'Field 2 Default',
                    'hint': 'Field 2 Hint'
                },
                {
                    'select': 'Select 1',
                    'key': 'select-1',
                    'options': [
                        {
                            'label': 'Select 1 Option 1',
                            'value': 'select-1-option-1'
                        },
                        {
                            'label': 'Select 1 Option 2',
                            'value': 'select-1-option-2'
                        }
                    ]
                },
                {
                    'select': 'Select 2',
                    'key': 'select-2',
                    'hint': 'Select 2 Hint',
                    'required': False,
                    'default': 'select-2-option-1',
                    'options': [
                        {
                            'label': 'Select 2 Option 1',
                            'value': 'select-2-option-1'
                        }
                    ]
                }
            ]
        }
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'if': 'build.message !~ /skip tests/'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'if': 'build.message !~ /skip tests/'}]})

    def test_key(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'key': 'key'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'key': 'key'}]})

    def test_depends_on_string(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'depends_on': 'step'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': 'step'}]})

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'depends_on': ['one','two']})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': ['one','two']}]})

    def test_depends_on_object_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({
                    'block': 'a label',
                    'depends_on': [
                        {'step': 'one', 'allow_failure': True},
                        {'step': 'two'}
                    ],
                })
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': [{'step': 'one', 'allow_failure': True},{'step':'two'}]}]})

    def test_depends_on_mixed_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({
                    'block': 'a label',
                    'depends_on': [
                        {'step': 'one', 'allow_failure': True},
                        'two'
                    ],
                })
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'depends_on': [{'step': 'one', 'allow_failure': True},'two']}]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'allow_dependency_failure': True})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allow_dependency_failure': True}]})

    def test_multiple_fields(self):
        expected: BlockStepDict = {
            'block': 'a label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'options': [
                        {'label': 'Option 1', 'value': 'option-1'},
                        {'label': 'Option 2', 'value': 'option-2'},
                    ],
                },
            ],
        }
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_multiple_fields_with_default(self):
        expected: BlockStepDict = {
            'block': 'a label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'default': ['option-1', 'option-2'],
                    'options': [
                        {'label': 'Option 1', 'value': 'option-1'},
                        {'label': 'Option 2', 'value': 'option-2'},
                    ],
                },
            ],
        }
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allowed_teams_string(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'allowed_teams': 'team'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allowed_teams': 'team'}]})

    def test_allowed_teams_list(self):
        pipeline = Pipeline(
            steps=[
                BlockStep.from_dict({'block': 'a label', 'allowed_teams': ['one', 'two']})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'block': 'a label', 'allowed_teams': ['one', 'two']}]})

    def test_group_nesting(self):
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict({'group': 'Tests', 'steps': [{'block': 'a label'}]})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'group': 'Tests', 'steps': [{'block': 'a label'}]}]})
