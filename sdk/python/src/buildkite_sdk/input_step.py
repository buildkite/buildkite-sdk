from typing import List, Optional, TypedDict, Union

from buildkite_sdk.schema import InputStep as _input_step
from buildkite_sdk.types import DependsOn, SelectField, TextField


class InputStepArgs(TypedDict):
    input: str
    fields: List[Union[SelectField, TextField]]
    allow_dependency_failure: Optional[bool]
    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]


def InputStep(
    input: str,
    fields: List[Union[SelectField, TextField]],
    allow_dependency_failure: Optional[bool] = None,
    branches: Optional[Union[List[str], str]] = None,
    depends_on: Optional[Union[List[Union[DependsOn, str]], str]] = None,
    id: Optional[str] = None,
    identifier: Optional[str] = None,
    step_if: Optional[str] = None,
    key: Optional[str] = None,
    label: Optional[str] = None,
    name: Optional[str] = None,
    prompt: Optional[str] = None,
) -> _input_step:
    return _input_step(
        allow_dependency_failure=allow_dependency_failure,
        branches=branches,
        depends_on=depends_on,
        fields=fields,
        id=id,
        identifier=identifier,
        step_if=step_if,
        input=input,
        key=key,
        label=label,
        name=name,
        prompt=prompt,
        type=None,
    )
