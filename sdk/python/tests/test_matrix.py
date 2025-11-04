from buildkite_sdk import Pipeline, CommandStep, CommandStepArgs, MatrixObject, MatrixAdjustments, GroupStep, GroupStepArgs
from .utils import TestRunner

class TestMatrixClass(TestRunner):
    def test_simple(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='echo {{matrix}}',
                    label='{{matrix}}',
                    matrix=['one','two'],
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'command': 'echo {{matrix}}', 'label': '{{matrix}}', 'matrix': ['one','two']}]})

    def test_single_adjustements(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='echo {{matrix}}',
                    label='{{matrix}}',
                    matrix=MatrixObject(
                        setup=['one','two'],
                        adjustments=[
                            MatrixAdjustments(matrix_with=['three'], skip=True)
                        ]
                    ),
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{
            'command': 'echo {{matrix}}',
            'label': '{{matrix}}',
            'matrix': {
                'setup': ['one','two'],
                'adjustments': [
                    {'with': ['three'], 'skip': True},
                ],
            }
        }]})

    def test_multiple_adjustments(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='echo {{matrix.color}} {{matrix.shape}}',
                    label='{{matrix.color}} {{matrix.shape}}',
                    matrix=MatrixObject(
                        setup={
                            'color': ['green','blue'],
                            'shape': ['triangle','hexagon']
                        },
                        adjustments=[
                            MatrixAdjustments(
                                matrix_with={
                                    'color': 'blue',
                                    'shape': 'triangle'
                                },
                                skip=True
                            ),
                            MatrixAdjustments(
                                matrix_with={
                                    'color': 'green',
                                    'shape': 'triangle'
                                },
                                skip='look, hexagons are just better'
                            ),
                            MatrixAdjustments(
                                matrix_with={
                                    'color': 'purple',
                                    'shape': 'octagon'
                                },
                            )
                        ]
                    ),
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{
            'command': 'echo {{matrix.color}} {{matrix.shape}}',
            'label': '{{matrix.color}} {{matrix.shape}}',
            'matrix': {
                'setup': {
                    'color': ['green','blue'],
                    'shape': ['triangle','hexagon']
                },
                'adjustments': [
                    {
                        'with': {
                            'color': 'blue',
                            'shape': 'triangle'
                        },
                        'skip': True
                    },
                    {
                        'with': {
                            'color': 'green',
                            'shape': 'triangle'
                        },
                        'skip': 'look, hexagons are just better'
                    },
                    {
                        'with': {
                            'color': 'purple',
                            'shape': 'octagon'
                        }
                    }
                ]
            }
        }]})

    def test_group(self):
        expected: GroupStepArgs = {
            'group': 'matrices',
            'steps': [
                {
                    'command': 'echo {{matrix}}',
                    'label': '{{matrix}}',
                    'matrix': ['one','two']
                },
                {
                    'command': 'echo {{matrix.color}} {{matrix.shape}}',
                    'label': '{{matrix.color}} {{matrix.shape}}',
                    'matrix': {
                        'setup': {
                            'color': ['green','blue'],
                            'shape': ['triangle','hexagon']
                        }
                    }
                }
            ]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep(
                    group='matrices',
                    steps=[
                        CommandStep(
                            command='echo {{matrix}}',
                            label='{{matrix}}',
                            matrix=['one','two'],
                        ),
                        CommandStep(
                            command='echo {{matrix.color}} {{matrix.shape}}',
                            label='{{matrix.color}} {{matrix.shape}}',
                            matrix=MatrixObject(
                                setup={
                                    'color': ['green','blue'],
                                    'shape': ['triangle','hexagon']
                                }
                            )
                        ),
                    ]
                ),
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestMatrixDict(TestRunner):
    def test_simple(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict({'command': 'echo {{matrix}}', 'label': '{{matrix}}', 'matrix': ['one','two']})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'command': 'echo {{matrix}}', 'label': '{{matrix}}', 'matrix': ['one','two']}]})

    def test_single_adjustements(self):
        expected: CommandStepArgs = {
            'command': 'echo {{matrix}}',
            'label': '{{matrix}}',
            'matrix': {
                'setup': ['one','two'],
                'adjustments': [
                    {'with': ['three'], 'skip': True},
                ],
            }
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_multiple_adjustments(self):
        expected: CommandStepArgs = {
            'command': 'echo {{matrix.color}} {{matrix.shape}}',
            'label': '{{matrix.color}} {{matrix.shape}}',
            'matrix': {
                'setup': {
                    'color': ['green','blue'],
                    'shape': ['triangle','hexagon']
                },
                'adjustments': [
                    {
                        'with': {
                            'color': 'blue',
                            'shape': 'triangle'
                        },
                        'skip': True
                    },
                    {
                        'with': {
                            'color': 'green',
                            'shape': 'triangle'
                        },
                        'skip': 'look, hexagons are just better'
                    },
                    {
                        'with': {
                            'color': 'purple',
                            'shape': 'octagon'
                        }
                    }
                ]
            }
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_group(self):
        expected: GroupStepArgs = {
            'group': 'matrices',
            'steps': [
                {
                    'command': 'echo {{matrix}}',
                    'label': '{{matrix}}',
                    'matrix': ['one','two']
                },
                {
                    'command': 'echo {{matrix.color}} {{matrix.shape}}',
                    'label': '{{matrix.color}} {{matrix.shape}}',
                    'matrix': {
                        'setup': {
                            'color': ['green','blue'],
                            'shape': ['triangle','hexagon']
                        }
                    }
                }
            ]
        }
        pipeline = Pipeline(
            steps=[
                GroupStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
