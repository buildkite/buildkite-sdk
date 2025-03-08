from typing import Optional, List, Union
from buildkite_sdk.types import (
    BlockedStateEnum,
    DependsOn,
    SelectField,
    TextField,
)
from buildkite_sdk.schema import (
    BlockStep as _block_step
)

def BlockStep(
    block: Optional[str],
    allow_dependency_failure: Optional[bool] = None,
    blocked_state: Optional[BlockedStateEnum] = None,
    branches: Optional[Union[List[str], str]] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    fields: Optional[List[Union[SelectField, TextField]]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    block_step_if: Optional[str] = None,
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
        block_step_if=block_step_if,
        key=key,
        label=label,
        name=name,
        prompt=prompt,
        type=None,
    )
