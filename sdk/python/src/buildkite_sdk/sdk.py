from buildkite_sdk.schema import CommandStep, WaitStep, BlockStep, GroupStepClass, TriggerStep, InputStep
from buildkite_sdk.environment import Environment
import json
import yaml

"""
A pipeline.
"""


class Pipeline:
    """
    A description of the class.
    """

    def __init__(self):
        """A description of the constructor."""
        self.steps = []
        """I guess this is where we define the steps?"""

    def add_command_step(self, props: CommandStep):
        """Add a command step to the pipeline."""
        self.steps.append(props)

    def add_wait_step(self, props: WaitStep):
        """Add a wait step to the pipeline"""
        self.steps.append(props)

    def add_block_step(self, props: BlockStep):
        """Add a block step to the pipeline"""
        self.steps.append(props)

    def add_input_step(self, props: InputStep):
        """Add a input step to the pipeline"""
        self.steps.append(props)

    def add_trigger_step(self, props: TriggerStep):
        """Add a trigger step to the pipeline"""
        self.steps.append(props)

    def add_group_step(self, props: GroupStepClass):
        """Add a group step to the pipeline"""
        self.steps.append(props)

    def to_json(self):
        """Serialize the pipeline as a JSON string."""
        return json.dumps(self.__dict__, indent=4)

    def to_yaml(self):
        """Serialize the pipeline as a YAML string."""
        return yaml.dump(self.__dict__)
