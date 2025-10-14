from .schema import Env, Agents, BlockStepDict, BuildNotify, Image, BlockStep, NestedBlockStep,StringBlockStep,InputStep,NestedInputStep,StringInputStep,CommandStep,NestedCommandStep,WaitStep,NestedWaitStep,StringWaitStep,TriggerStep,NestedTriggerStep,GroupStep
from typing import Optional, List, Any, TypedDict, NotRequired
from pydantic import BaseModel
import json
import yaml

type Step = BlockStepDict | StringBlockStep | BlockStep | NestedBlockStep | InputStep | NestedInputStep | StringInputStep | CommandStep | NestedCommandStep | WaitStep | NestedWaitStep | StringWaitStep | TriggerStep | NestedTriggerStep | GroupStep

class PipelineDict(TypedDict):
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
    def from_dict(cls, data: PipelineDict):
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
        """Serialize the pipeline as a JSON string."""
        return self.model_dump(
            by_alias=True,
            exclude_none=True,
        )

    def to_json(self):
        pipeline_json = self.to_dict()
        return json.dumps(pipeline_json)

    def to_yaml_string(self):
        """Serialize the pipeline as a YAML string."""
        return yaml.dump(self)
