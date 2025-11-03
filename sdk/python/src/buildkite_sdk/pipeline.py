from .schema import (
    Env,
    Agents,
    BuildNotify,
    Image,
    BlockStepArgs,
    BlockStep,
    NestedBlockStepArgs,
    NestedBlockStep,
    StringBlockStep,
    InputStepArgs,
    InputStep,
    NestedInputStepArgs,
    NestedInputStep,
    StringInputStep,
    CommandStepArgs,
    CommandStep,
    NestedCommandStepArgs,
    NestedCommandStep,
    WaitStepArgs,
    WaitStep,
    NestedWaitStepArgs,
    NestedWaitStep,
    StringWaitStep,
    TriggerStepArgs,
    TriggerStep,
    NestedTriggerStepArgs,
    NestedTriggerStep,
    GroupStepArgs,
    GroupStep
)
from typing import Optional, List, Any, TypedDict, NotRequired, Union
from pydantic import BaseModel
import json
import yaml

type Step = Union[
    BlockStepArgs,
    BlockStep,
    NestedBlockStepArgs,
    NestedBlockStep,
    StringBlockStep,
    InputStepArgs,
    InputStep,
    NestedInputStepArgs,
    NestedInputStep,
    StringInputStep,
    CommandStepArgs,
    CommandStep,
    NestedCommandStepArgs,
    NestedCommandStep,
    WaitStepArgs,
    WaitStep,
    NestedWaitStepArgs,
    NestedWaitStep,
    StringWaitStep,
    TriggerStepArgs,
    TriggerStep,
    NestedTriggerStepArgs,
    NestedTriggerStep,
    GroupStepArgs,
    GroupStep
]

class PipelineArgs(TypedDict):
    env: NotRequired[Env]
    agents: NotRequired[Agents]
    notify: NotRequired[BuildNotify]
    image: NotRequired[Image]
    steps: List[Step]

class Pipeline(BaseModel):
    env: Optional[Env] = None
    agents: Optional[Agents] = None
    notify: Optional[BuildNotify] = None
    image: Optional[Image] = None
    steps: List[Step] = []

    @classmethod
    def from_dict(cls, data: PipelineArgs):
        return cls(**data)

    def add_agent(self, key: str, value: Any):
        if self.agents == None:
            self.agents = {}

        if isinstance(self.agents, List):
            self.agents.append(f"{key}={value}")
        else:
            self.agents[key] = value

    def add_environment_variable(self, key: str, value: Any):
        if self.env == None:
            self.env = {}
        self.env[key] = value

    def add_notify(self, notify: BuildNotify):
        self.notify = notify

    def add_step(self, step: Step):
        self.steps.append(step)

    def to_dict(self):
        return self.model_dump(
            by_alias=True,
            exclude_none=True,
        )

    def to_json(self):
        """Serialize the pipeline as a JSON string."""
        pipeline_json = self.to_dict()
        return json.dumps(pipeline_json, indent=4)

    def to_yaml(self):
        """Serialize the pipeline as a YAML string."""
        print(self.to_dict())
        return yaml.dump(self.to_dict())
