from typing import List, Optional, TypedDict, Union

from buildkite_sdk.schema import BlockStep as _block_step
from buildkite_sdk.types import (
    BlockedStateEnum,
    DependsOn,
    SelectField,
    TextField,
)


class BlockStepArgs(TypedDict, total=False):
    block: Optional[str]
    allow_dependency_failure: Optional[bool]
    blocked_state: Optional[BlockedStateEnum]
    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    fields: Optional[List[Union[SelectField, TextField]]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]


def BlockStep(
    block: Optional[str],
    allow_dependency_failure: Optional[bool] = None,
    blocked_state: Optional[BlockedStateEnum] = None,
    branches: Optional[Union[List[str], str]] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    fields: Optional[List[Union[SelectField, TextField]]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    name: Optional[str] = None,
    prompt: Optional[str] = None,
) -> _block_step:
    return _block_step(
        allow_dependency_failure=allow_dependency_failure,
        block=block,
        blocked_state=blocked_state,
        branches=branches,
        depends_on=depends_on,
        fields=fields,
        id=id,
        identifier=identifier,
        step_if=step_if,
        key=key,
        label=label,
        name=name,
        prompt=prompt,
        type=None,
    )
