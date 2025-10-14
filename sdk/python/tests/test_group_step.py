from buildkite_sdk import Pipeline, GroupStep, GroupStepDict, CommandStep, WaitStep, NestedWaitStep, InputStep, NestedInputStep, NotifyEmail
from .utils import TestRunner

class TestGroupStepNestingTypesClass(TestRunner):
    def test_field(self):
        pipeline = Pipeline(
            steps=[
                GroupStep(group='Tests', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'group': 'Tests', 'steps': [{'command': 'test'}]}]})

    def test_label(self):
        expected: GroupStepDict = {
            'group': '~',
            'label': 'Tests',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(group='~', label='Tests', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestGroupStepNestingTypesDict(TestRunner):
    def test_field(self):
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict({'group': 'Tests', 'steps': [{'command': 'test'}]})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'group': 'Tests', 'steps': [{'command': 'test'}]}]})

    def test_label(self):
        expected: GroupStepDict = {
            'group': '~',
            'label': 'Tests',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestGroupStepClass(TestRunner):
    def test_id(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'id': 'id',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(group='Tests', id='id', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'identifier': 'identifier',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(group='Tests', identifier='identifier', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'depends_on': 'step',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(group='Tests', depends_on='step', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'key': 'key',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(group='Tests', key='key', steps=[CommandStep(command='test')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_wait(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [
                'wait',
                {'key': 'waiter', 'type': 'wait'},
                {'wait': { 'key': 'waiter2', 'type': 'wait' }}
            ],
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    steps=[
                        'wait',
                        WaitStep(key='waiter', type='wait'),
                        NestedWaitStep(wait=WaitStep(key='waiter2', type='wait'))
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_input(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [
                'input',
                {'input': 'a label'},
                {'key': 'input', 'type': 'input'},
                {'input': {'key': 'input2', 'type': 'input'}}
            ],
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    steps=[
                        'input',
                        InputStep(input='a label'),
                        InputStep(key='input', type='input'),
                        NestedInputStep(input=InputStep(key='input2', type='input'))
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected = {
            'group': 'Tests',
            'if': 'build.message !~ /skip tests/',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    step_if='build.message !~ /skip tests/',
                    steps=[CommandStep(command='test')]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'allow_dependency_failure': True,
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    allow_dependency_failure=True,
                    steps=[CommandStep(command='test')]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_notify(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [{'command': 'test'}],
            'notify': [{'email': 'dev@acmeinc.com'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    steps=[CommandStep(command='test')],
                    notify=[NotifyEmail(email='dev@acmeinc.com')]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if_changed(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [{'command': 'test'}],
            'if_changed': '*.txt'
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='Tests',
                    steps=[CommandStep(command='test')],
                    if_changed='*.txt'
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestGroupStepDict(TestRunner):
    def test_id(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'id': 'id',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'identifier': 'identifier',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'depends_on': 'step',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'key': 'key',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_wait(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [
                'wait',
                {'key': 'waiter', 'type': 'wait'},
                {'wait': { 'key': 'waiter2', 'type': 'wait' }}
            ],
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_input(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [
                'input',
                {'input': 'a label'},
                {'key': 'input', 'type': 'input'},
                {'input': {'key': 'input2', 'type': 'input'}}
            ],
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'if': 'build.message !~ /skip tests/',
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{
            'group': 'Tests',
            'if': 'build.message !~ /skip tests/',
            'steps': [{'command': 'test'}]
        }]})

    def test_allow_dependency_failure(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'allow_dependency_failure': True,
            'steps': [{'command': 'test'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_notify(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [{'command': 'test'}],
            'notify': [{'email': 'dev@acmeinc.com'}]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if_changed(self):
        expected: GroupStepDict = {
            'group': 'Tests',
            'steps': [{'command': 'test'}],
            'if_changed': '*.txt'
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
