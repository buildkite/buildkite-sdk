from buildkite_sdk.schema import (
    BlockStep as _block_step,
    CommandStep as _command_step,
    GroupStepClass as _group_step,
    InputStep as _input_step,
    TriggerStep as _trigger_step,
    WaitStep as _wait_step,
)
from buildkite_sdk.environment import Environment
from typing import Union
import json
import yaml

class Pipeline:
    """
    A pipeline.
    """

    def __init__(self):
        """A description of the constructor."""
        self.steps = []
        """I guess this is where we define the steps?"""

    def add_step(
        self,
        props: Union[
            _block_step, _command_step, _group_step, _input_step, _trigger_step, _wait_step
        ],
    ):
        """Add a command step to the pipeline."""
        self.steps.append(props.to_dict())

    def to_json(self):
        """Serialize the pipeline as a JSON string."""
        return json.dumps(self.__dict__, indent=4)

    def to_yaml(self):
        """Serialize the pipeline as a YAML string."""
        return yaml.dump(self.__dict__)
