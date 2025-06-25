from typing import Optional, List, Union, TypedDict
from collections.abc import Sequence
from buildkite_sdk.schema import PurpleStep
from buildkite_sdk.types import (
    DependsOn,
    NotifyEnum,
    StepNotify,
)
from buildkite_sdk.schema import (
    GroupStepClass as _group_step_class,
    BlockStep as _block_step,
    CommandStep as _command_step,
    InputStep as _input_step,
    TriggerStep as _trigger_step,
    WaitStep as _wait_step,
)
from .block_step import BlockStepArgs
from .command_step import CommandStepArgs
from .input_step import InputStepArgs
from .trigger_step import TriggerStepArgs
from .wait_step import WaitStepArgs


class GroupStepArgs(TypedDict):
    group: str
    steps: Sequence[
        Union[
            BlockStepArgs,
            CommandStepArgs,
            InputStepArgs,
            TriggerStepArgs,
            WaitStepArgs,
        ]
    ]
    allow_dependency_failure: Optional[bool]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    notify: Optional[List[Union[StepNotify, NotifyEnum]]]
    skip: Optional[Union[bool, str]]


def _step_to_purple_step(
    step: Union[
        _group_step_class,
        _block_step,
        _command_step,
        _input_step,
        _trigger_step,
        _wait_step,
    ],
):
    return PurpleStep(
        allow_dependency_failure=getattr(step, "allow_dependency_failure", None),
        depends_on=getattr(step, "depends_on", None),
        id=getattr(step, "id", None),
        identifier=getattr(step, "identifier", None),
        step_if=getattr(step, "step_if", None),
        key=getattr(step, "key", None),
        label=getattr(step, "label", None),
        name=getattr(step, "name", None),
        notify=getattr(step, "notify", None),
        skip=getattr(step, "skip", None),
        block=getattr(step, "block", None),
        blocked_state=getattr(step, "blocked_state", None),
        branches=getattr(step, "branches", None),
        fields=getattr(step, "fields", None),
        prompt=getattr(step, "prompt", None),
        type=None,
        input=getattr(step, "input", None),
        agents=getattr(step, "agents", None),
        artifact_paths=getattr(step, "artifact_paths", None),
        cache=getattr(step, "cache", None),
        cancel_on_build_failing=getattr(step, "cancel_on_build_failing", None),
        command=getattr(step, "command", None),
        commands=getattr(step, "commands", None),
        concurrency=getattr(step, "concurrency", None),
        concurrency_group=getattr(step, "concurrency_group", None),
        concurrency_method=getattr(step, "concurrency_method", None),
        env=getattr(step, "env", None),
        matrix=getattr(step, "matrix", None),
        parallelism=getattr(step, "parallelism", None),
        plugins=getattr(step, "plugins", None),
        priority=getattr(step, "priority", None),
        retry=getattr(step, "retry", None),
        signature=getattr(step, "signature", None),
        soft_fail=getattr(step, "soft_fail", None),
        timeout_in_minutes=getattr(step, "timeout_in_minutes", None),
        script=getattr(step, "script", None),
        continue_on_failure=getattr(step, "continue_on_failure", None),
        wait=getattr(step, "wait", None),
        waiter=getattr(step, "waiter", None),
        step_async=getattr(step, "step_async", None),
        build=getattr(step, "build", None),
        trigger=getattr(step, "trigger", None),
    )


def GroupStep(
    group: str,
    steps: Sequence[
        Union[
            _group_step_class,
            _block_step,
            _command_step,
            _input_step,
            _trigger_step,
            _wait_step,
        ]
    ],
    allow_dependency_failure: Optional[bool] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    name: Optional[str] = None,
    notify: Optional[List[Union[StepNotify, NotifyEnum]]] = None,
    skip: Optional[Union[bool, str]] = None,
) -> _group_step_class:
    group_steps = []
    for step in steps:
        group_steps.append(_step_to_purple_step(step))

    return _group_step_class(
        allow_dependency_failure=allow_dependency_failure,
        depends_on=depends_on,
        id=id,
        identifier=identifier,
        step_if=step_if,
        key=key,
        label=label,
        name=name,
        notify=notify,
        skip=skip,
        group=group,
        steps=group_steps,
        block=None,
        blocked_state=None,
        branches=None,
        fields=None,
        prompt=None,
        type=None,
        input=None,
        agents=None,
        artifact_paths=None,
        cache=None,
        cancel_on_build_failing=None,
        command=None,
        commands=None,
        concurrency=None,
        concurrency_group=None,
        concurrency_method=None,
        env=None,
        matrix=None,
        parallelism=None,
        plugins=None,
        priority=None,
        retry=None,
        signature=None,
        soft_fail=None,
        timeout_in_minutes=None,
        script=None,
        continue_on_failure=None,
        wait=None,
        waiter=None,
        step_async=None,
        build=None,
        trigger=None,
    )
