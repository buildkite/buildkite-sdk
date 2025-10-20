from typing import Optional, List, Union, TypedDict
from buildkite_sdk.types import (
    Build,
    DependsOn,
    SoftFail,
)
from buildkite_sdk.schema import (
    TriggerStep as _trigger_step,
)


class TriggerStepArgs(TypedDict):
    trigger: str
    allow_dependency_failure: Optional[bool]
    trigger_step_async: Optional[bool]
    branches: Optional[Union[List[str], str]]
    build: Optional[Build]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFail]]]


def TriggerStep(
    trigger: str,
    allow_dependency_failure: Optional[bool] = None,
    trigger_step_async: Optional[bool] = None,
    branches: Optional[Union[List[str], str]] = None,
    build: Optional[Build] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    name: Optional[str] = None,
    skip: Optional[Union[bool, str]] = None,
    soft_fail: Optional[Union[bool, List[SoftFail]]] = None,
) -> _trigger_step:
    return _trigger_step(
        trigger=trigger,
        allow_dependency_failure=allow_dependency_failure,
        trigger_step_async=trigger_step_async,
        branches=branches,
        build=build,
        depends_on=depends_on,
        id=id,
        identifier=identifier,
        step_if=step_if,
        key=key,
        label=label,
        name=name,
        skip=skip,
        soft_fail=soft_fail,
        type=None,
    )
