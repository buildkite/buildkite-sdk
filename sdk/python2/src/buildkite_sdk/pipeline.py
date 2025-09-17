from .schema import Env, Agents, BlockStepDict, BuildNotify, Image, PipelineSteps, BlockStep, NestedBlockStep,StringBlockStep,InputStep,NestedInputStep,StringInputStep,CommandStep,NestedCommandStep,WaitStep,NestedWaitStep,StringWaitStep,TriggerStep,NestedTriggerStep,GroupStep
from typing import Optional, List, Any, Literal
from pydantic import BaseModel
import json

type Step = BlockStepDict | Literal['block'] | BlockStep #| NestedBlockStep | InputStep | NestedInputStep | StringInputStep | CommandStep | NestedCommandStep | WaitStep | NestedWaitStep | StringWaitStep | TriggerStep | NestedTriggerStep | GroupStep

class Pipeline(BaseModel):
    env: Optional[Env] = None
    agents: Optional[Agents] = None
    notify: Optional[BuildNotify] = None
    image: Optional[Image] = None
    steps: Optional[List[Step]] = None

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
        if self.steps == None:
            self.steps = []
        self.steps.append(step)

    def to_json(self):
        return self.model_dump(
            by_alias=True,
            exclude_none=True,
            exclude_unset=True,
        )

    def to_json_string(self):
        pipeline_json = self.to_json()
        return json.dumps(pipeline_json)
