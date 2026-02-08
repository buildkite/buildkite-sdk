from __future__ import annotations

import json
import sys

if sys.version_info >= (3, 12):
    from typing import Any, List, NotRequired, Optional, TypedDict
else:
    from typing import Any, List, Optional

    from typing_extensions import NotRequired, TypedDict

import yaml
from pydantic import BaseModel

from .schema import (
    Agents,
    BlockStep,
    BlockStepArgs,
    BuildNotify,
    CommandStep,
    CommandStepArgs,
    Env,
    GroupStep,
    GroupStepArgs,
    Image,
    InputStep,
    InputStepArgs,
    NestedBlockStep,
    NestedBlockStepArgs,
    NestedCommandStep,
    NestedCommandStepArgs,
    NestedInputStep,
    NestedInputStepArgs,
    NestedTriggerStep,
    NestedTriggerStepArgs,
    NestedWaitStep,
    NestedWaitStepArgs,
    Secrets,
    StringBlockStep,
    StringInputStep,
    StringWaitStep,
    TriggerStep,
    TriggerStepArgs,
    WaitStep,
    WaitStepArgs,
)

Step = (
    BlockStepArgs
    | BlockStep
    | NestedBlockStepArgs
    | NestedBlockStep
    | StringBlockStep
    | InputStepArgs
    | InputStep
    | NestedInputStepArgs
    | NestedInputStep
    | StringInputStep
    | CommandStepArgs
    | CommandStep
    | NestedCommandStepArgs
    | NestedCommandStep
    | WaitStepArgs
    | WaitStep
    | NestedWaitStepArgs
    | NestedWaitStep
    | StringWaitStep
    | TriggerStepArgs
    | TriggerStep
    | NestedTriggerStepArgs
    | NestedTriggerStep
    | GroupStepArgs
    | GroupStep
)


class PipelineArgs(TypedDict):
    env: NotRequired[Env]
    agents: NotRequired[Agents]
    notify: NotRequired[BuildNotify]
    image: NotRequired[Image]
    secrets: NotRequired[Secrets]
    steps: List[Step]


class Pipeline(BaseModel):
    env: Optional[Env] = None
    agents: Optional[Agents] = None
    notify: Optional[BuildNotify] = None
    image: Optional[Image] = None
    secrets: Optional[Secrets] = None
    steps: List[Step] = []

    @classmethod
    def from_dict(cls, data: PipelineArgs) -> Pipeline:
        return cls(**data)

    def set_secrets(self, secrets: Secrets) -> None:
        self.secrets = secrets

    def add_agent(self, key: str, value: Any) -> None:
        if self.agents is None:
            self.agents = {}

        if isinstance(self.agents, List):
            self.agents.append(f"{key}={value}")
            return

        self.agents[key] = value

    def add_environment_variable(self, key: str, value: Any) -> None:
        if self.env is None:
            self.env = {}
        self.env[key] = value

    def add_notify(self, notify: BuildNotify) -> None:
        self.notify = notify

    def add_step(self, step: Step) -> None:
        self.steps.append(step)

    def to_dict(self) -> dict[str, Any]:
        return self.model_dump(
            by_alias=True,
            exclude_none=True,
            mode="json",
        )

    def to_json(self) -> str:
        """Serialize the pipeline as a JSON string."""
        pipeline_json = self.to_dict()
        return json.dumps(pipeline_json, indent=4)

    def to_yaml(self) -> str:
        """Serialize the pipeline as a YAML string."""
        result: str = yaml.dump(self.to_dict())
        return result
