from buildkite_sdk import Pipeline, InputStep, NestedInputStep, InputStepDict, TextField, SelectField, SelectFieldOption, DependsOnListObject
from .utils import TestRunner

class TestInputStepNestingTypesClass(TestRunner):
    def test_string(self):
        pipeline = Pipeline(
            steps=['input']
        )
        self.validator.check_result(pipeline, {'steps': ['input']})

    def test_field(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label'}]})

    def test_nested(self):
        pipeline = Pipeline(
            steps=[
                NestedInputStep(input=InputStep(fields=[TextField(text='Field 1', key='field-1')]))
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': {'fields': [{'text': 'Field 1', 'key': 'field-1'}]}}]})

    def test_type(self):
        pipeline = Pipeline(
            steps=[
                InputStep(type='input', label='a label', fields=[TextField(text='Field 1', key='field-1')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'type': 'input', 'label': 'a label', 'fields': [{'text': 'Field 1', 'key': 'field-1'}]}]})

class TestInputStepNestingTypesDict(TestRunner):
    def test_field(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label'}]})

    def test_nested(self):
        pipeline = Pipeline(
            steps=[
                NestedInputStep.from_dict({'input': {'fields': [{'text': 'Field 1', 'key': 'field-1'}]}})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': {'fields': [{'text': 'Field 1', 'key': 'field-1'}]}}]})

    def test_type(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'type': 'input', 'label': 'a label', 'fields': [{'text': 'Field 1', 'key': 'field-1'}]})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'type': 'input', 'label': 'a label', 'fields': [{'text': 'Field 1', 'key': 'field-1'}]}]})

class TestInputStepClass(TestRunner):
    def test_branches(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', branches='branch')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'branches': 'branch'}]})

    def test_id(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', id='id')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'id': 'id'}]})

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', identifier='identifier')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'identifier': 'identifier'}]})

    def test_prompt(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', prompt='prompt')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'prompt': 'prompt'}]})

    def test_fields(self):
        expected: InputStepDict = {
            'input': 'A label',
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
                InputStep(
                    input='A label',
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
                            required=False,
                            default='select-2-option-1',
                            hint='Select 2 Hint',
                            options=[SelectFieldOption(label='Select 2 Option 1', value='select-2-option-1')]
                        ),
                    ]
                ),
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', step_if='build.message !~ /skip tests/')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'if': 'build.message !~ /skip tests/'}]})

    def test_key(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', key='key')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'key': 'key'}]})

    def test_depends_on_string(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', depends_on='step')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'depends_on': 'step'}]})

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', depends_on=['one','two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'depends_on': ['one','two']}]})

    def test_depends_on_object_list(self):
        expected: InputStepDict = {
            'input': 'a label',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ],
        }
        pipeline = Pipeline(
            steps=[
                InputStep(
                    input='a label',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        DependsOnListObject(step='two')
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: InputStepDict = {
            'input': 'a label',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ],
        }
        pipeline = Pipeline(
            steps=[
                InputStep(
                    input='a label',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        'two'
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', allow_dependency_failure=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allow_dependency_failure': True}]})

    def test_allowed_teams_string(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', allowed_teams='team')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allowed_teams': 'team'}]})

    def test_allowed_teams_string_list(self):
        pipeline = Pipeline(
            steps=[
                InputStep(input='a label', allowed_teams=['one','two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allowed_teams': ['one','two']}]})

    def test_multiple_fields(self):
        expected: InputStepDict = {
            'input': 'A label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'options': [
                        {
                            'label': 'Option 1',
                            'value': 'option-1'
                        },
                        {
                            'label': 'Option 2',
                            'value': 'option-2'
                        }
                    ]
                }
            ]
        }
        pipeline = Pipeline(
            steps=[
                InputStep(
                    input='A label',
                    fields=[
                        SelectField(
                            select='Multiple fields',
                            key='multiple-fields',
                            multiple=True,
                            options=[
                                SelectFieldOption(label='Option 1', value='option-1'),
                                SelectFieldOption(label='Option 2', value='option-2'),
                            ]
                        ),
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestInputStepDict(TestRunner):
    def test_branches(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'branches': 'branch'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'branches': 'branch'}]})

    def test_id(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'id': 'id'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'id': 'id'}]})

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'identifier': 'identifier'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'identifier': 'identifier'}]})

    def test_prompt(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'prompt': 'prompt'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'prompt': 'prompt'}]})

    def test_fields(self):
        expected: InputStepDict = {
            'input': 'A label',
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
                InputStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'if': 'build.message !~ /skip tests/'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'if': 'build.message !~ /skip tests/'}]})

    def test_key(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'key': 'key'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'key': 'key'}]})

    def test_depends_on_string(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'depends_on': 'step'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'depends_on': 'step'}]})

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'depends_on': ['one','two']})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'depends_on': ['one','two']}]})

    def test_depends_on_object_list(self):
        expected: InputStepDict = {
            'input': 'a label',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ],
        }
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: InputStepDict = {
            'input': 'a label',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ],
        }
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'allow_dependency_failure': True})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allow_dependency_failure': True}]})

    def test_allowed_teams_string(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'allowed_teams': 'team'})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allowed_teams': 'team'}]})

    def test_allowed_teams_string_list(self):
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict({'input': 'a label', 'allowed_teams': ['one','two']})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'input': 'a label', 'allowed_teams': ['one','two']}]})

    def test_multiple_fields(self):
        expected: InputStepDict = {
            'input': 'A label',
            'fields': [
                {
                    'select': 'Multiple fields',
                    'key': 'multiple-fields',
                    'multiple': True,
                    'options': [
                        {
                            'label': 'Option 1',
                            'value': 'option-1'
                        },
                        {
                            'label': 'Option 2',
                            'value': 'option-2'
                        }
                    ]
                }
            ]
        }
        pipeline = Pipeline(
            steps=[
                InputStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
