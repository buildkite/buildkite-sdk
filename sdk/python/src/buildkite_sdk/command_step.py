from typing import Any, Dict, List, Optional, TypedDict, Union

from buildkite_sdk.schema import CommandStep as _command_step
from buildkite_sdk.types import (
    Cache,
    ConcurrencyMethod,
    DependsOn,
    MatrixAdvanced,
    NotifyEnum,
    Retry,
    Signature,
    SoftFail,
    StepNotify,
)


class CommandStepArgs(TypedDict, total=False):
    commands: Optional[Union[List[str], str]]
    command: Optional[Union[List[str], str]]
    agents: Optional[Union[Dict[str, Any], List[str]]]
    allow_dependency_failure: Optional[bool]
    artifact_paths: Optional[Union[List[str], str]]
    branches: Optional[Union[List[str], str]]
    cache: Optional[Union[List[str], Cache, str]]
    cancel_on_build_failing: Optional[bool]
    concurrency: Optional[int]
    concurrency_group: Optional[str]
    concurrency_method: Optional[ConcurrencyMethod]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    env: Optional[Dict[str, Any]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    matrix: Optional[Union[List[Union[int, bool, str]], MatrixAdvanced]]
    name: Optional[str]
    notify: Optional[List[Union[StepNotify, NotifyEnum]]]
    parallelism: Optional[int]
    plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]]
    priority: Optional[int]
    retry: Optional[Retry]
    signature: Optional[Signature]
    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFail]]]
    timeout_in_minutes: Optional[int]


def CommandStep(
    commands: Optional[Union[List[str], str]] = None,
    command: Optional[Union[List[str], str]] = None,
    agents: Optional[Union[Dict[str, Any], List[str]]] = None,
    allow_dependency_failure: Optional[bool] = None,
    artifact_paths: Optional[Union[List[str], str]] = None,
    branches: Optional[Union[List[str], str]] = None,
    cache: Optional[Union[List[str], Cache, str]] = None,
    cancel_on_build_failing: Optional[bool] = None,
    concurrency: Optional[int] = None,
    concurrency_group: Optional[str] = None,
    concurrency_method: Optional[ConcurrencyMethod] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    env: Optional[Dict[str, Any]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    matrix: Optional[Union[List[Union[int, bool, str]], MatrixAdvanced]] = None,
    name: Optional[str] = None,
    notify: Optional[List[Union[StepNotify, NotifyEnum]]] = None,
    parallelism: Optional[int] = None,
    plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]] = None,
    priority: Optional[int] = None,
    retry: Optional[Retry] = None,
    signature: Optional[Signature] = None,
    skip: Optional[Union[bool, str]] = None,
    soft_fail: Optional[Union[bool, List[SoftFail]]] = None,
    timeout_in_minutes: Optional[int] = None,
) -> _command_step:
    return _command_step(
        agents=agents,
        allow_dependency_failure=allow_dependency_failure,
        artifact_paths=artifact_paths,
        branches=branches,
        cache=cache,
        cancel_on_build_failing=cancel_on_build_failing,
        command=command,
        commands=commands,
        concurrency=concurrency,
        concurrency_group=concurrency_group,
        concurrency_method=concurrency_method,
        depends_on=depends_on,
        env=env,
        id=id,
        identifier=identifier,
        step_if=step_if,
        key=key,
        label=label,
        matrix=matrix,
        name=name,
        notify=notify,
        parallelism=parallelism,
        plugins=plugins,
        priority=priority,
        retry=retry,
        signature=signature,
        skip=skip,
        soft_fail=soft_fail,
        timeout_in_minutes=timeout_in_minutes,
        type=None,
    )
