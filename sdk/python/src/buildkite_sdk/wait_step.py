from typing import Optional, List, Union
from buildkite_sdk.types import (
    DependsOn
)
from buildkite_sdk.schema import (
    WaitStep as _wait_step,
)

def WaitStep(
    wait: Optional[str] = "~",
    allow_dependency_failure: Optional[bool] = None,
    branches: Optional[Union[List[str], str]] = None,
    continue_on_failure: Optional[bool] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    wait_step_if: Optional[str] = None,
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
        wait_step_if=wait_step_if,
        key=key,
        label=label,
        name=name,
        type=None,
    )
