from typing import Optional, List, Union, TypedDict
from buildkite_sdk.types import DependsOn
from buildkite_sdk.schema import (
    WaitStep as _wait_step,
)


class WaitStepArgs(TypedDict):
    wait: str
    allow_dependency_failure: Optional[bool]
    branches: Optional[Union[List[str], str]]
    continue_on_failure: Optional[bool]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]


def WaitStep(
    wait: Optional[str] = "~",
    allow_dependency_failure: Optional[bool] = None,
    branches: Optional[Union[List[str], str]] = None,
    continue_on_failure: Optional[bool] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    name: Optional[str] = None,
) -> _wait_step:
    return _wait_step(
        wait=wait,
        allow_dependency_failure=allow_dependency_failure,
        branches=branches,
        continue_on_failure=continue_on_failure,
        depends_on=depends_on,
        id=id,
        identifier=identifier,
        step_if=step_if,
        key=key,
        label=label,
        name=name,
        type=None,
    )
