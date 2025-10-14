from buildkite_sdk import Pipeline, WaitStep, WaitStepDict, NestedWaitStep, NestedWaitStepDict, DependsOnListObject
from .utils import TestRunner

class TestWaitStepNestingTypesClass(TestRunner):
    def test_string_wait(self):
        pipeline = Pipeline(steps=['wait'])
        self.validator.check_result(pipeline, {'steps': ['wait']})

    def test_string_waiter(self):
        pipeline = Pipeline(steps=['waiter'])
        self.validator.check_result(pipeline, {'steps': ['waiter']})

    def test_type_wait(self):
        expected: WaitStepDict = {'type': 'wait'}
        pipeline = Pipeline(
            steps=[
                WaitStep(type='wait')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_type_waiter(self):
        expected: WaitStepDict = {'type': 'waiter'}
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_wait_field(self):
        expected: WaitStepDict = {'wait': '~', 'continue_on_failure': True}
        pipeline = Pipeline(
            steps=[
                WaitStep(wait='~', continue_on_failure=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested_wait(self):
        expected: NestedWaitStepDict = {
            'wait': {'continue_on_failure': True}
        }
        pipeline = Pipeline(
            steps=[
                NestedWaitStep(wait=WaitStep(continue_on_failure=True))
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested_waiter(self):
        expected: NestedWaitStepDict = {
            'waiter': {'continue_on_failure': True}
        }
        pipeline = Pipeline(
            steps=[
                NestedWaitStep(waiter=WaitStep(continue_on_failure=True))
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestWaitStepNestingTypesDict(TestRunner):
    def test_type_wait(self):
        expected: WaitStepDict = {'type': 'wait'}
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_type_waiter(self):
        expected: WaitStepDict = {'type': 'waiter'}
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_wait_field(self):
        expected: WaitStepDict = {'wait': '~', 'continue_on_failure': True}
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested_wait(self):
        expected: NestedWaitStepDict = {
            'wait': {'continue_on_failure': True}
        }
        pipeline = Pipeline(
            steps=[
                NestedWaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested_waiter(self):
        expected: NestedWaitStepDict = {
            'waiter': {'continue_on_failure': True}
        }
        pipeline = Pipeline(
            steps=[
                NestedWaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestWaitStepClass(TestRunner):
    def test_id(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'id': 'id'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', id='id')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'identifier': 'identifier'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', identifier='identifier')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'if': 'build.message !~ /skip tests/'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', step_if='build.message !~ /skip tests/')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'key': 'key'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', key='key')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': 'step'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', depends_on='step')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': ['one', 'two']
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', depends_on=['one', 'two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_object_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(
                    type='waiter',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        DependsOnListObject(step='two')
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ]
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(
                    type='waiter',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        'two'
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'allow_dependency_failure': True
        }
        pipeline = Pipeline(
            steps=[
                WaitStep(type='waiter', allow_dependency_failure=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestWaitStepDict(TestRunner):
    def test_id(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'id': 'id'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'identifier': 'identifier'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'if': 'build.message !~ /skip tests/'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'key': 'key'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': 'step'
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': ['one', 'two']
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_object_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ]
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        expected: WaitStepDict = {
            'type': 'waiter',
            'allow_dependency_failure': True
        }
        pipeline = Pipeline(
            steps=[
                WaitStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
