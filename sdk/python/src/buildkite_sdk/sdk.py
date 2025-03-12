from .schema import (
    BlockStep as _block_step,
    CommandStep as _command_step,
    GroupStepClass as _group_step,
    InputStep as _input_step,
    TriggerStep as _trigger_step,
    WaitStep as _wait_step,
)
from .environment import Environment
from .types import PipelineNotify
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
        self.agents = None
        self.env = None
        self.notify = None
        """I guess this is where we define the steps?"""

    def add_agent(self, key: str, value: any):
        if self.agents == None:
            self.agents = {}
        self.agents[key] = value

    def add_environment_variable(self, key: str, value: any):
        if self.env == None:
            self.env = {}
        self.env[key] = value

    def add_notify(self, notify: PipelineNotify):
        self.notify = notify

    def add_step(
        self,
        props: Union[
            _block_step,
            _command_step,
            _group_step,
            _input_step,
            _trigger_step,
            _wait_step,
        ],
    ):
        """Add a command step to the pipeline."""
        self.steps.append(props.to_dict())

    def to_pipeline(self):
        pipeline = {}
        pipeline["steps"] = self.steps

        if self.agents != None:
            pipeline["agents"] = self.agents
        if self.env != None:
            pipeline["env"] = self.env
        if self.notify != None:
            pipeline["notify"] = self.notify

        return pipeline

    def to_json(self):
        """Serialize the pipeline as a JSON string."""
        return json.dumps(self.to_pipeline(), indent=4)

    def to_yaml(self):
        """Serialize the pipeline as a YAML string."""
        return yaml.dump(self.to_pipeline())
