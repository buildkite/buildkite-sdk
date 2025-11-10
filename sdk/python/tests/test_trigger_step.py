from buildkite_sdk import Pipeline, TriggerStep, TriggerStepArgs, TriggerStepBuild, NestedTriggerStep, NestedTriggerStepArgs, DependsOnListObject, SoftFailObject
from .utils import TestRunner

class TestTriggerStepNestingTypesClass(TestRunner):
    def test_field(self):
        expected: TriggerStepArgs = {'trigger': 'a-slug'}
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_type(self):
        expected: TriggerStepArgs = {'trigger': 'a-slug', 'type': 'trigger'}
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', type='trigger')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested(self):
        expected: NestedTriggerStepArgs = {'trigger': {'trigger': 'a-slug'}}
        pipeline = Pipeline(
            steps=[
                NestedTriggerStep(trigger=TriggerStep(trigger='a-slug'))
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestTriggerStepNestingTypesDict(TestRunner):
    def test_field(self):
        expected: TriggerStepArgs = {'trigger': 'a-slug'}
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_type(self):
        expected: TriggerStepArgs = {'trigger': 'a-slug', 'type': 'trigger'}
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_nested(self):
        expected: NestedTriggerStepArgs = {'trigger': {'trigger': 'a-slug'}}
        pipeline = Pipeline(
            steps=[
                NestedTriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestTriggerStepClass(TestRunner):
    def test_async(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'async': True,
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', step_async=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_branches(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'branches': ['one','two']
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', branches=['one','two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_message(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'message': 'a message',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    build=TriggerStepBuild(message='a message')
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_commit(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'commit': 'a commit',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    build=TriggerStepBuild(commit='a commit')
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_branch(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'branch': 'a branch',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    build=TriggerStepBuild(branch='a branch')
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_meta_data(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'meta_data': {'a-key': 'a-val'}
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    build=TriggerStepBuild(meta_data={'a-key': 'a-val'})
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_env(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'env': {'SOME_ENV': 'some-val'}
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    build=TriggerStepBuild(env={'SOME_ENV': 'some-val'})
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_id(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'id': 'id'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', id='id')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'identifier': 'identifier'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', identifier='identifier')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_label(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'label': 'a label',
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', label='a label')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'if': "build.message !~ /skip tests/"
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', step_if="build.message !~ /skip tests/")
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'key': 'key'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', key='key')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': 'step'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', depends_on='step')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': ['one','two']
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', depends_on=['one','two'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_object_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        DependsOnListObject(step='two')
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(
                    trigger='a-slug',
                    depends_on=[
                        DependsOnListObject(step='one', allow_failure=True),
                        'two'
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'allow_dependency_failure': True
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', allow_dependency_failure=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_skip_bool(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'skip': True
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', skip=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_skip_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'skip': 'a reason'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', skip='a reason')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_bool(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': True,
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', soft_fail=True)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_exit_status_int(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': [{'exit_status': -1}]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', soft_fail=[SoftFailObject(exit_status=-1)])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_exit_status_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': [{'exit_status': '*'}]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', soft_fail=[SoftFailObject(exit_status='*')])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if_changed(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'if_changed': '*.txt'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep(trigger='a-slug', if_changed='*.txt')
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestTriggerStepArgs(TestRunner):
    def test_async(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'async': True,
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_branches(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'branches': ['one','two']
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_message(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'message': 'a message',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_commit(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'commit': 'a commit',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_branch(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'branch': 'a branch',
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_meta_data(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'meta_data': {'a-key': 'a-val'}
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_build_env(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'build': {
                'env': {'SOME_ENV': 'some-val'}
            }
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_id(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'id': 'id'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_identifier(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'identifier': 'identifier'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_label(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'label': 'a label',
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'if': "build.message !~ /skip tests/"
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_key(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'key': 'key'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': 'step'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_string_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': ['one','two']
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_object_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                {'step': 'two'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_depends_on_mixed_list(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'depends_on': [
                {'step': 'one', 'allow_failure': True},
                'two'
            ]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_allow_dependency_failure(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'allow_dependency_failure': True
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_skip_bool(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'skip': True
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_skip_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'skip': 'a reason'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_bool(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': True,
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_exit_status_int(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': [{'exit_status': -1}]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_soft_fail_exit_status_string(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'soft_fail': [{'exit_status': '*'}]
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_if_changed(self):
        expected: TriggerStepArgs = {
            'trigger': 'a-slug',
            'if_changed': '*.txt'
        }
        pipeline = Pipeline(
            steps=[
                TriggerStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
