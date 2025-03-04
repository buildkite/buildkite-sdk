from typing import Optional, Any, List, Union, Dict, TypeVar, Callable, Type, cast
from enum import Enum


T = TypeVar("T")
EnumT = TypeVar("EnumT", bound=Enum)


def from_str(x: Any) -> str:
    assert isinstance(x, str)
    return x


def from_none(x: Any) -> Any:
    assert x is None
    return x


def from_union(fs, x):
    for f in fs:
        try:
            return f(x)
        except:
            pass
    assert False


def from_list(f: Callable[[Any], T], x: Any) -> List[T]:
    assert isinstance(x, list)
    return [f(y) for y in x]


def to_class(c: Type[T], x: Any) -> dict:
    assert isinstance(x, c)
    return cast(Any, x).to_dict()


def from_bool(x: Any) -> bool:
    assert isinstance(x, bool)
    return x


def to_enum(c: Type[EnumT], x: Any) -> EnumT:
    assert isinstance(x, c)
    return x.value


def from_dict(f: Callable[[Any], T], x: Any) -> Dict[str, T]:
    assert isinstance(x, dict)
    return {k: f(v) for (k, v) in x.items()}


def from_int(x: Any) -> int:
    assert isinstance(x, int) and not isinstance(x, bool)
    return x


class PurpleGithubCheck:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "PurpleGithubCheck":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return PurpleGithubCheck(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class PurpleGithubCommitStatus:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "PurpleGithubCommitStatus":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return PurpleGithubCommitStatus(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class PurpleSlack:
    channels: Optional[List[str]]
    message: Optional[str]

    def __init__(self, channels: Optional[List[str]], message: Optional[str]) -> None:
        self.channels = channels
        self.message = message

    @staticmethod
    def from_dict(obj: Any) -> "PurpleSlack":
        assert isinstance(obj, dict)
        channels = from_union(
            [lambda x: from_list(from_str, x), from_none], obj.get("channels")
        )
        message = from_union([from_str, from_none], obj.get("message"))
        return PurpleSlack(channels, message)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.channels is not None:
            result["channels"] = from_union(
                [lambda x: from_list(from_str, x), from_none], self.channels
            )
        if self.message is not None:
            result["message"] = from_union([from_str, from_none], self.message)
        return result


class PurpleBuildNotify:
    email: Optional[str]
    build_notify_if: Optional[str]
    basecamp_campfire: Optional[str]
    slack: Optional[Union[PurpleSlack, str]]
    webhook: Optional[str]
    pagerduty_change_event: Optional[str]
    github_commit_status: Optional[PurpleGithubCommitStatus]
    github_check: Optional[PurpleGithubCheck]

    def __init__(
        self,
        email: Optional[str],
        build_notify_if: Optional[str],
        basecamp_campfire: Optional[str],
        slack: Optional[Union[PurpleSlack, str]],
        webhook: Optional[str],
        pagerduty_change_event: Optional[str],
        github_commit_status: Optional[PurpleGithubCommitStatus],
        github_check: Optional[PurpleGithubCheck],
    ) -> None:
        self.email = email
        self.build_notify_if = build_notify_if
        self.basecamp_campfire = basecamp_campfire
        self.slack = slack
        self.webhook = webhook
        self.pagerduty_change_event = pagerduty_change_event
        self.github_commit_status = github_commit_status
        self.github_check = github_check

    @staticmethod
    def from_dict(obj: Any) -> "PurpleBuildNotify":
        assert isinstance(obj, dict)
        email = from_union([from_str, from_none], obj.get("email"))
        build_notify_if = from_union([from_str, from_none], obj.get("if"))
        basecamp_campfire = from_union(
            [from_str, from_none], obj.get("basecamp_campfire")
        )
        slack = from_union(
            [PurpleSlack.from_dict, from_str, from_none], obj.get("slack")
        )
        webhook = from_union([from_str, from_none], obj.get("webhook"))
        pagerduty_change_event = from_union(
            [from_str, from_none], obj.get("pagerduty_change_event")
        )
        github_commit_status = from_union(
            [PurpleGithubCommitStatus.from_dict, from_none],
            obj.get("github_commit_status"),
        )
        github_check = from_union(
            [PurpleGithubCheck.from_dict, from_none], obj.get("github_check")
        )
        return PurpleBuildNotify(
            email,
            build_notify_if,
            basecamp_campfire,
            slack,
            webhook,
            pagerduty_change_event,
            github_commit_status,
            github_check,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.email is not None:
            result["email"] = from_union([from_str, from_none], self.email)
        if self.build_notify_if is not None:
            result["if"] = from_union([from_str, from_none], self.build_notify_if)
        if self.basecamp_campfire is not None:
            result["basecamp_campfire"] = from_union(
                [from_str, from_none], self.basecamp_campfire
            )
        if self.slack is not None:
            result["slack"] = from_union(
                [lambda x: to_class(PurpleSlack, x), from_str, from_none], self.slack
            )
        if self.webhook is not None:
            result["webhook"] = from_union([from_str, from_none], self.webhook)
        if self.pagerduty_change_event is not None:
            result["pagerduty_change_event"] = from_union(
                [from_str, from_none], self.pagerduty_change_event
            )
        if self.github_commit_status is not None:
            result["github_commit_status"] = from_union(
                [lambda x: to_class(PurpleGithubCommitStatus, x), from_none],
                self.github_commit_status,
            )
        if self.github_check is not None:
            result["github_check"] = from_union(
                [lambda x: to_class(PurpleGithubCheck, x), from_none], self.github_check
            )
        return result


class NotifyEnum(Enum):
    GITHUB_CHECK = "github_check"
    GITHUB_COMMIT_STATUS = "github_commit_status"


class AllowDependencyFailureEnum(Enum):
    FALSE = "false"
    TRUE = "true"


class BlockedState(Enum):
    """The state that the build is set to when the build is blocked by this block step"""

    FAILED = "failed"
    PASSED = "passed"
    RUNNING = "running"


class DependsOnClass:
    allow_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    step: Optional[str]

    def __init__(
        self,
        allow_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        step: Optional[str],
    ) -> None:
        self.allow_failure = allow_failure
        self.step = step

    @staticmethod
    def from_dict(obj: Any) -> "DependsOnClass":
        assert isinstance(obj, dict)
        allow_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("allow_failure")
        )
        step = from_union([from_str, from_none], obj.get("step"))
        return DependsOnClass(allow_failure, step)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_failure is not None:
            result["allow_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_failure,
            )
        if self.step is not None:
            result["step"] = from_union([from_str, from_none], self.step)
        return result


class Option:
    hint: Optional[str]
    """The text displayed directly under the select field’s label"""

    label: str
    """The text displayed on the select list item"""

    required: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether the field is required for form submission"""

    value: str
    """The value to be stored as meta-data"""

    def __init__(
        self,
        hint: Optional[str],
        label: str,
        required: Optional[Union[bool, AllowDependencyFailureEnum]],
        value: str,
    ) -> None:
        self.hint = hint
        self.label = label
        self.required = required
        self.value = value

    @staticmethod
    def from_dict(obj: Any) -> "Option":
        assert isinstance(obj, dict)
        hint = from_union([from_str, from_none], obj.get("hint"))
        label = from_str(obj.get("label"))
        required = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("required")
        )
        value = from_str(obj.get("value"))
        return Option(hint, label, required, value)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.hint is not None:
            result["hint"] = from_union([from_str, from_none], self.hint)
        result["label"] = from_str(self.label)
        if self.required is not None:
            result["required"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.required,
            )
        result["value"] = from_str(self.value)
        return result


class Field:
    """A list of input fields required to be filled out before unblocking the step"""

    default: Optional[Union[List[str], str]]
    """The value that is pre-filled in the text field
    
    The value of the option(s) that will be pre-selected in the dropdown
    """
    format: Optional[str]
    """The format must be a regular expression implicitly anchored to the beginning and end of
    the input and is functionally equivalent to the HTML5 pattern attribute.
    """
    hint: Optional[str]
    """The explanatory text that is shown after the label"""

    key: str
    """The meta-data key that stores the field's input"""

    required: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether the field is required for form submission"""

    text: Optional[str]
    """The text input name"""

    multiple: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether more than one option may be selected"""

    options: Optional[List[Option]]
    select: Optional[str]
    """The text input name"""

    def __init__(
        self,
        default: Optional[Union[List[str], str]],
        format: Optional[str],
        hint: Optional[str],
        key: str,
        required: Optional[Union[bool, AllowDependencyFailureEnum]],
        text: Optional[str],
        multiple: Optional[Union[bool, AllowDependencyFailureEnum]],
        options: Optional[List[Option]],
        select: Optional[str],
    ) -> None:
        self.default = default
        self.format = format
        self.hint = hint
        self.key = key
        self.required = required
        self.text = text
        self.multiple = multiple
        self.options = options
        self.select = select

    @staticmethod
    def from_dict(obj: Any) -> "Field":
        assert isinstance(obj, dict)
        default = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("default")
        )
        format = from_union([from_str, from_none], obj.get("format"))
        hint = from_union([from_str, from_none], obj.get("hint"))
        key = from_str(obj.get("key"))
        required = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("required")
        )
        text = from_union([from_str, from_none], obj.get("text"))
        multiple = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("multiple")
        )
        options = from_union(
            [lambda x: from_list(Option.from_dict, x), from_none], obj.get("options")
        )
        select = from_union([from_str, from_none], obj.get("select"))
        return Field(
            default, format, hint, key, required, text, multiple, options, select
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.default is not None:
            result["default"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.default
            )
        if self.format is not None:
            result["format"] = from_union([from_str, from_none], self.format)
        if self.hint is not None:
            result["hint"] = from_union([from_str, from_none], self.hint)
        result["key"] = from_str(self.key)
        if self.required is not None:
            result["required"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.required,
            )
        if self.text is not None:
            result["text"] = from_union([from_str, from_none], self.text)
        if self.multiple is not None:
            result["multiple"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.multiple,
            )
        if self.options is not None:
            result["options"] = from_union(
                [lambda x: from_list(lambda x: to_class(Option, x), x), from_none],
                self.options,
            )
        if self.select is not None:
            result["select"] = from_union([from_str, from_none], self.select)
        return result


class BlockType(Enum):
    BLOCK = "block"


class BlockStep:
    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    block: Optional[str]
    """The label of the block step"""

    blocked_state: Optional[BlockedState]
    """The state that the build is set to when the build is blocked by this block step"""

    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    fields: Optional[List[Field]]
    id: Optional[str]
    identifier: Optional[str]
    block_step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]
    type: Optional[BlockType]

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        block: Optional[str],
        blocked_state: Optional[BlockedState],
        branches: Optional[Union[List[str], str]],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        fields: Optional[List[Field]],
        id: Optional[str],
        identifier: Optional[str],
        block_step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        prompt: Optional[str],
        type: Optional[BlockType],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.block = block
        self.blocked_state = blocked_state
        self.branches = branches
        self.depends_on = depends_on
        self.fields = fields
        self.id = id
        self.identifier = identifier
        self.block_step_if = block_step_if
        self.key = key
        self.label = label
        self.name = name
        self.prompt = prompt
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> "BlockStep":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        block = from_union([from_str, from_none], obj.get("block"))
        blocked_state = from_union([BlockedState, from_none], obj.get("blocked_state"))
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        fields = from_union(
            [lambda x: from_list(Field.from_dict, x), from_none], obj.get("fields")
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        block_step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_str, from_none], obj.get("label"))
        name = from_union([from_str, from_none], obj.get("name"))
        prompt = from_union([from_str, from_none], obj.get("prompt"))
        type = from_union([BlockType, from_none], obj.get("type"))
        return BlockStep(
            allow_dependency_failure,
            block,
            blocked_state,
            branches,
            depends_on,
            fields,
            id,
            identifier,
            block_step_if,
            key,
            label,
            name,
            prompt,
            type,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.block is not None:
            result["block"] = from_union([from_str, from_none], self.block)
        if self.blocked_state is not None:
            result["blocked_state"] = from_union(
                [lambda x: to_enum(BlockedState, x), from_none], self.blocked_state
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.fields is not None:
            result["fields"] = from_union(
                [lambda x: from_list(lambda x: to_class(Field, x), x), from_none],
                self.fields,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.block_step_if is not None:
            result["if"] = from_union([from_str, from_none], self.block_step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_str, from_none], self.label)
        if self.name is not None:
            result["name"] = from_union([from_str, from_none], self.name)
        if self.prompt is not None:
            result["prompt"] = from_union([from_str, from_none], self.prompt)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(BlockType, x), from_none], self.type
            )
        return result


class Build:
    """Properties of the build that will be created when the step is triggered"""

    branch: Optional[str]
    """The branch for the build"""

    commit: Optional[str]
    """The commit hash for the build"""

    env: Optional[Dict[str, Any]]
    message: Optional[str]
    """The message for the build (supports emoji)"""

    meta_data: Optional[Dict[str, Any]]
    """Meta-data for the build"""

    def __init__(
        self,
        branch: Optional[str],
        commit: Optional[str],
        env: Optional[Dict[str, Any]],
        message: Optional[str],
        meta_data: Optional[Dict[str, Any]],
    ) -> None:
        self.branch = branch
        self.commit = commit
        self.env = env
        self.message = message
        self.meta_data = meta_data

    @staticmethod
    def from_dict(obj: Any) -> "Build":
        assert isinstance(obj, dict)
        branch = from_union([from_str, from_none], obj.get("branch"))
        commit = from_union([from_str, from_none], obj.get("commit"))
        env = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("env")
        )
        message = from_union([from_str, from_none], obj.get("message"))
        meta_data = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("meta_data")
        )
        return Build(branch, commit, env, message, meta_data)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.branch is not None:
            result["branch"] = from_union([from_str, from_none], self.branch)
        if self.commit is not None:
            result["commit"] = from_union([from_str, from_none], self.commit)
        if self.env is not None:
            result["env"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.env
            )
        if self.message is not None:
            result["message"] = from_union([from_str, from_none], self.message)
        if self.meta_data is not None:
            result["meta_data"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.meta_data
            )
        return result


class CacheClass:
    name: Optional[str]
    paths: List[str]
    size: Optional[str]

    def __init__(
        self, name: Optional[str], paths: List[str], size: Optional[str]
    ) -> None:
        self.name = name
        self.paths = paths
        self.size = size

    @staticmethod
    def from_dict(obj: Any) -> "CacheClass":
        assert isinstance(obj, dict)
        name = from_union([from_str, from_none], obj.get("name"))
        paths = from_list(from_str, obj.get("paths"))
        size = from_union([from_str, from_none], obj.get("size"))
        return CacheClass(name, paths, size)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.name is not None:
            result["name"] = from_union([from_str, from_none], self.name)
        result["paths"] = from_list(from_str, self.paths)
        if self.size is not None:
            result["size"] = from_union([from_str, from_none], self.size)
        return result


class ConcurrencyMethod(Enum):
    """Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
    this attribute, you must also define concurrency_group and concurrency.
    """

    EAGER = "eager"
    ORDERED = "ordered"


class ExitStatusEnum(Enum):
    EMPTY = "*"


class SoftFailElement:
    exit_status: Optional[Union[int, ExitStatusEnum]]
    """The exit status number that will cause this job to soft-fail"""

    def __init__(self, exit_status: Optional[Union[int, ExitStatusEnum]]) -> None:
        self.exit_status = exit_status

    @staticmethod
    def from_dict(obj: Any) -> "SoftFailElement":
        assert isinstance(obj, dict)
        exit_status = from_union(
            [from_int, ExitStatusEnum, from_none], obj.get("exit_status")
        )
        return SoftFailElement(exit_status)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.exit_status is not None:
            result["exit_status"] = from_union(
                [from_int, lambda x: to_enum(ExitStatusEnum, x), from_none],
                self.exit_status,
            )
        return result


class Adjustment:
    """An adjustment to a Build Matrix"""

    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]]
    adjustment_with: Union[List[Union[int, bool, str]], Dict[str, str]]

    def __init__(
        self,
        skip: Optional[Union[bool, str]],
        soft_fail: Optional[
            Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]
        ],
        adjustment_with: Union[List[Union[int, bool, str]], Dict[str, str]],
    ) -> None:
        self.skip = skip
        self.soft_fail = soft_fail
        self.adjustment_with = adjustment_with

    @staticmethod
    def from_dict(obj: Any) -> "Adjustment":
        assert isinstance(obj, dict)
        skip = from_union([from_bool, from_str, from_none], obj.get("skip"))
        soft_fail = from_union(
            [
                from_bool,
                lambda x: from_list(SoftFailElement.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("soft_fail"),
        )
        adjustment_with = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                lambda x: from_dict(from_str, x),
            ],
            obj.get("with"),
        )
        return Adjustment(skip, soft_fail, adjustment_with)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.skip is not None:
            result["skip"] = from_union([from_bool, from_str, from_none], self.skip)
        if self.soft_fail is not None:
            result["soft_fail"] = from_union(
                [
                    from_bool,
                    lambda x: from_list(lambda x: to_class(SoftFailElement, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.soft_fail,
            )
        result["with"] = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                lambda x: from_dict(from_str, x),
            ],
            self.adjustment_with,
        )
        return result


class MatrixClass:
    """Configuration for multi-dimension Build Matrix"""

    adjustments: Optional[List[Adjustment]]
    """List of Build Matrix adjustments"""

    setup: Union[List[Union[int, bool, str]], Dict[str, List[Union[int, bool, str]]]]

    def __init__(
        self,
        adjustments: Optional[List[Adjustment]],
        setup: Union[
            List[Union[int, bool, str]], Dict[str, List[Union[int, bool, str]]]
        ],
    ) -> None:
        self.adjustments = adjustments
        self.setup = setup

    @staticmethod
    def from_dict(obj: Any) -> "MatrixClass":
        assert isinstance(obj, dict)
        adjustments = from_union(
            [lambda x: from_list(Adjustment.from_dict, x), from_none],
            obj.get("adjustments"),
        )
        setup = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                lambda x: from_dict(
                    lambda x: from_list(
                        lambda x: from_union([from_int, from_bool, from_str], x), x
                    ),
                    x,
                ),
            ],
            obj.get("setup"),
        )
        return MatrixClass(adjustments, setup)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.adjustments is not None:
            result["adjustments"] = from_union(
                [lambda x: from_list(lambda x: to_class(Adjustment, x), x), from_none],
                self.adjustments,
            )
        result["setup"] = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                lambda x: from_dict(
                    lambda x: from_list(
                        lambda x: from_union([from_int, from_bool, from_str], x), x
                    ),
                    x,
                ),
            ],
            self.setup,
        )
        return result


class FluffyGithubCheck:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "FluffyGithubCheck":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return FluffyGithubCheck(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class FluffyGithubCommitStatus:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "FluffyGithubCommitStatus":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return FluffyGithubCommitStatus(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class FluffySlack:
    channels: Optional[List[str]]
    message: Optional[str]

    def __init__(self, channels: Optional[List[str]], message: Optional[str]) -> None:
        self.channels = channels
        self.message = message

    @staticmethod
    def from_dict(obj: Any) -> "FluffySlack":
        assert isinstance(obj, dict)
        channels = from_union(
            [lambda x: from_list(from_str, x), from_none], obj.get("channels")
        )
        message = from_union([from_str, from_none], obj.get("message"))
        return FluffySlack(channels, message)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.channels is not None:
            result["channels"] = from_union(
                [lambda x: from_list(from_str, x), from_none], self.channels
            )
        if self.message is not None:
            result["message"] = from_union([from_str, from_none], self.message)
        return result


class NotifyClass:
    basecamp_campfire: Optional[str]
    notify_if: Optional[str]
    slack: Optional[Union[FluffySlack, str]]
    github_commit_status: Optional[FluffyGithubCommitStatus]
    github_check: Optional[FluffyGithubCheck]

    def __init__(
        self,
        basecamp_campfire: Optional[str],
        notify_if: Optional[str],
        slack: Optional[Union[FluffySlack, str]],
        github_commit_status: Optional[FluffyGithubCommitStatus],
        github_check: Optional[FluffyGithubCheck],
    ) -> None:
        self.basecamp_campfire = basecamp_campfire
        self.notify_if = notify_if
        self.slack = slack
        self.github_commit_status = github_commit_status
        self.github_check = github_check

    @staticmethod
    def from_dict(obj: Any) -> "NotifyClass":
        assert isinstance(obj, dict)
        basecamp_campfire = from_union(
            [from_str, from_none], obj.get("basecamp_campfire")
        )
        notify_if = from_union([from_str, from_none], obj.get("if"))
        slack = from_union(
            [FluffySlack.from_dict, from_str, from_none], obj.get("slack")
        )
        github_commit_status = from_union(
            [FluffyGithubCommitStatus.from_dict, from_none],
            obj.get("github_commit_status"),
        )
        github_check = from_union(
            [FluffyGithubCheck.from_dict, from_none], obj.get("github_check")
        )
        return NotifyClass(
            basecamp_campfire, notify_if, slack, github_commit_status, github_check
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.basecamp_campfire is not None:
            result["basecamp_campfire"] = from_union(
                [from_str, from_none], self.basecamp_campfire
            )
        if self.notify_if is not None:
            result["if"] = from_union([from_str, from_none], self.notify_if)
        if self.slack is not None:
            result["slack"] = from_union(
                [lambda x: to_class(FluffySlack, x), from_str, from_none], self.slack
            )
        if self.github_commit_status is not None:
            result["github_commit_status"] = from_union(
                [lambda x: to_class(FluffyGithubCommitStatus, x), from_none],
                self.github_commit_status,
            )
        if self.github_check is not None:
            result["github_check"] = from_union(
                [lambda x: to_class(FluffyGithubCheck, x), from_none], self.github_check
            )
        return result


class SignalReason(Enum):
    """The exit signal reason, if any, that may be retried"""

    AGENT_REFUSED = "agent_refused"
    AGENT_STOP = "agent_stop"
    CANCEL = "cancel"
    EMPTY = "*"
    NONE = "none"
    PROCESS_RUN_ERROR = "process_run_error"
    SIGNATURE_REJECTED = "signature_rejected"


class AutomaticRetry:
    exit_status: Optional[Union[int, List[int], ExitStatusEnum]]
    """The exit status number that will cause this job to retry"""

    limit: Optional[int]
    """The number of times this job can be retried"""

    signal: Optional[str]
    """The exit signal, if any, that may be retried"""

    signal_reason: Optional[SignalReason]
    """The exit signal reason, if any, that may be retried"""

    def __init__(
        self,
        exit_status: Optional[Union[int, List[int], ExitStatusEnum]],
        limit: Optional[int],
        signal: Optional[str],
        signal_reason: Optional[SignalReason],
    ) -> None:
        self.exit_status = exit_status
        self.limit = limit
        self.signal = signal
        self.signal_reason = signal_reason

    @staticmethod
    def from_dict(obj: Any) -> "AutomaticRetry":
        assert isinstance(obj, dict)
        exit_status = from_union(
            [from_int, lambda x: from_list(from_int, x), ExitStatusEnum, from_none],
            obj.get("exit_status"),
        )
        limit = from_union([from_int, from_none], obj.get("limit"))
        signal = from_union([from_str, from_none], obj.get("signal"))
        signal_reason = from_union([SignalReason, from_none], obj.get("signal_reason"))
        return AutomaticRetry(exit_status, limit, signal, signal_reason)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.exit_status is not None:
            result["exit_status"] = from_union(
                [
                    from_int,
                    lambda x: from_list(from_int, x),
                    lambda x: to_enum(ExitStatusEnum, x),
                    from_none,
                ],
                self.exit_status,
            )
        if self.limit is not None:
            result["limit"] = from_union([from_int, from_none], self.limit)
        if self.signal is not None:
            result["signal"] = from_union([from_str, from_none], self.signal)
        if self.signal_reason is not None:
            result["signal_reason"] = from_union(
                [lambda x: to_enum(SignalReason, x), from_none], self.signal_reason
            )
        return result


class ManualClass:
    allowed: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether or not this job can be retried manually"""

    permit_on_passed: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether or not this job can be retried after it has passed"""

    reason: Optional[str]
    """A string that will be displayed in a tooltip on the Retry button in Buildkite. This will
    only be displayed if the allowed attribute is set to false.
    """

    def __init__(
        self,
        allowed: Optional[Union[bool, AllowDependencyFailureEnum]],
        permit_on_passed: Optional[Union[bool, AllowDependencyFailureEnum]],
        reason: Optional[str],
    ) -> None:
        self.allowed = allowed
        self.permit_on_passed = permit_on_passed
        self.reason = reason

    @staticmethod
    def from_dict(obj: Any) -> "ManualClass":
        assert isinstance(obj, dict)
        allowed = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("allowed")
        )
        permit_on_passed = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("permit_on_passed"),
        )
        reason = from_union([from_str, from_none], obj.get("reason"))
        return ManualClass(allowed, permit_on_passed, reason)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allowed is not None:
            result["allowed"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allowed,
            )
        if self.permit_on_passed is not None:
            result["permit_on_passed"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.permit_on_passed,
            )
        if self.reason is not None:
            result["reason"] = from_union([from_str, from_none], self.reason)
        return result


class Retry:
    """The conditions for retrying this step."""

    automatic: Optional[
        Union[bool, AutomaticRetry, List[AutomaticRetry], AllowDependencyFailureEnum]
    ]
    """Whether to allow a job to retry automatically. If set to true, the retry conditions are
    set to the default value.
    """
    manual: Optional[Union[bool, ManualClass, AllowDependencyFailureEnum]]
    """Whether to allow a job to be retried manually"""

    def __init__(
        self,
        automatic: Optional[
            Union[
                bool, AutomaticRetry, List[AutomaticRetry], AllowDependencyFailureEnum
            ]
        ],
        manual: Optional[Union[bool, ManualClass, AllowDependencyFailureEnum]],
    ) -> None:
        self.automatic = automatic
        self.manual = manual

    @staticmethod
    def from_dict(obj: Any) -> "Retry":
        assert isinstance(obj, dict)
        automatic = from_union(
            [
                from_bool,
                AutomaticRetry.from_dict,
                lambda x: from_list(AutomaticRetry.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("automatic"),
        )
        manual = from_union(
            [from_bool, ManualClass.from_dict, AllowDependencyFailureEnum, from_none],
            obj.get("manual"),
        )
        return Retry(automatic, manual)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.automatic is not None:
            result["automatic"] = from_union(
                [
                    from_bool,
                    lambda x: to_class(AutomaticRetry, x),
                    lambda x: from_list(lambda x: to_class(AutomaticRetry, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.automatic,
            )
        if self.manual is not None:
            result["manual"] = from_union(
                [
                    from_bool,
                    lambda x: to_class(ManualClass, x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.manual,
            )
        return result


class Signature:
    """The signature of the command step, generally injected by agents at pipeline upload"""

    algorithm: Optional[str]
    """The algorithm used to generate the signature"""

    signed_fields: Optional[List[str]]
    """The fields that were signed to form the signature value"""

    value: Optional[str]
    """The signature value, a JWS compact signature with a detached body"""

    def __init__(
        self,
        algorithm: Optional[str],
        signed_fields: Optional[List[str]],
        value: Optional[str],
    ) -> None:
        self.algorithm = algorithm
        self.signed_fields = signed_fields
        self.value = value

    @staticmethod
    def from_dict(obj: Any) -> "Signature":
        assert isinstance(obj, dict)
        algorithm = from_union([from_str, from_none], obj.get("algorithm"))
        signed_fields = from_union(
            [lambda x: from_list(from_str, x), from_none], obj.get("signed_fields")
        )
        value = from_union([from_str, from_none], obj.get("value"))
        return Signature(algorithm, signed_fields, value)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.algorithm is not None:
            result["algorithm"] = from_union([from_str, from_none], self.algorithm)
        if self.signed_fields is not None:
            result["signed_fields"] = from_union(
                [lambda x: from_list(from_str, x), from_none], self.signed_fields
            )
        if self.value is not None:
            result["value"] = from_union([from_str, from_none], self.value)
        return result


class CommandType(Enum):
    COMMAND = "command"
    COMMANDS = "commands"
    SCRIPT = "script"


class CommandStep:
    agents: Optional[Union[Dict[str, Any], List[str]]]
    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    artifact_paths: Optional[Union[List[str], str]]
    """The glob path/s of artifacts to upload once this step has finished running"""

    branches: Optional[Union[List[str], str]]
    cache: Optional[Union[List[str], CacheClass, str]]
    cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]]
    command: Optional[Union[List[str], str]]
    """The commands to run on the agent"""

    commands: Optional[Union[List[str], str]]
    """The commands to run on the agent"""

    concurrency: Optional[int]
    """The maximum number of jobs created from this step that are allowed to run at the same
    time. If you use this attribute, you must also define concurrency_group.
    """
    concurrency_group: Optional[str]
    """A unique name for the concurrency group that you are creating with the concurrency
    attribute
    """
    concurrency_method: Optional[ConcurrencyMethod]
    """Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
    this attribute, you must also define concurrency_group and concurrency.
    """
    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    env: Optional[Dict[str, Any]]
    id: Optional[str]
    identifier: Optional[str]
    command_step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]]
    name: Optional[str]
    notify: Optional[List[Union[NotifyClass, NotifyEnum]]]
    """Array of notification options for this step"""

    parallelism: Optional[int]
    """The number of parallel jobs that will be created based on this step"""

    plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]]
    priority: Optional[int]
    """Priority of the job, higher priorities are assigned to agents"""

    retry: Optional[Retry]
    """The conditions for retrying this step."""

    signature: Optional[Signature]
    """The signature of the command step, generally injected by agents at pipeline upload"""

    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]]
    timeout_in_minutes: Optional[int]
    """The number of minutes to time out a job"""

    type: Optional[CommandType]

    def __init__(
        self,
        agents: Optional[Union[Dict[str, Any], List[str]]],
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        artifact_paths: Optional[Union[List[str], str]],
        branches: Optional[Union[List[str], str]],
        cache: Optional[Union[List[str], CacheClass, str]],
        cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]],
        command: Optional[Union[List[str], str]],
        commands: Optional[Union[List[str], str]],
        concurrency: Optional[int],
        concurrency_group: Optional[str],
        concurrency_method: Optional[ConcurrencyMethod],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        env: Optional[Dict[str, Any]],
        id: Optional[str],
        identifier: Optional[str],
        command_step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]],
        name: Optional[str],
        notify: Optional[List[Union[NotifyClass, NotifyEnum]]],
        parallelism: Optional[int],
        plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]],
        priority: Optional[int],
        retry: Optional[Retry],
        signature: Optional[Signature],
        skip: Optional[Union[bool, str]],
        soft_fail: Optional[
            Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]
        ],
        timeout_in_minutes: Optional[int],
        type: Optional[CommandType],
    ) -> None:
        self.agents = agents
        self.allow_dependency_failure = allow_dependency_failure
        self.artifact_paths = artifact_paths
        self.branches = branches
        self.cache = cache
        self.cancel_on_build_failing = cancel_on_build_failing
        self.command = command
        self.commands = commands
        self.concurrency = concurrency
        self.concurrency_group = concurrency_group
        self.concurrency_method = concurrency_method
        self.depends_on = depends_on
        self.env = env
        self.id = id
        self.identifier = identifier
        self.command_step_if = command_step_if
        self.key = key
        self.label = label
        self.matrix = matrix
        self.name = name
        self.notify = notify
        self.parallelism = parallelism
        self.plugins = plugins
        self.priority = priority
        self.retry = retry
        self.signature = signature
        self.skip = skip
        self.soft_fail = soft_fail
        self.timeout_in_minutes = timeout_in_minutes
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> "CommandStep":
        assert isinstance(obj, dict)
        agents = from_union(
            [
                lambda x: from_dict(lambda x: x, x),
                lambda x: from_list(from_str, x),
                from_none,
            ],
            obj.get("agents"),
        )
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        artifact_paths = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none],
            obj.get("artifact_paths"),
        )
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        cache = from_union(
            [
                lambda x: from_list(from_str, x),
                CacheClass.from_dict,
                from_str,
                from_none,
            ],
            obj.get("cache"),
        )
        cancel_on_build_failing = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("cancel_on_build_failing"),
        )
        command = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("command")
        )
        commands = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("commands")
        )
        concurrency = from_union([from_int, from_none], obj.get("concurrency"))
        concurrency_group = from_union(
            [from_str, from_none], obj.get("concurrency_group")
        )
        concurrency_method = from_union(
            [ConcurrencyMethod, from_none], obj.get("concurrency_method")
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        env = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("env")
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        command_step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_str, from_none], obj.get("label"))
        matrix = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                MatrixClass.from_dict,
                from_none,
            ],
            obj.get("matrix"),
        )
        name = from_union([from_str, from_none], obj.get("name"))
        notify = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([NotifyClass.from_dict, NotifyEnum], x), x
                ),
                from_none,
            ],
            obj.get("notify"),
        )
        parallelism = from_union([from_int, from_none], obj.get("parallelism"))
        plugins = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union(
                        [lambda x: from_dict(lambda x: x, x), from_str], x
                    ),
                    x,
                ),
                lambda x: from_dict(lambda x: x, x),
                from_none,
            ],
            obj.get("plugins"),
        )
        priority = from_union([from_int, from_none], obj.get("priority"))
        retry = from_union([Retry.from_dict, from_none], obj.get("retry"))
        signature = from_union([Signature.from_dict, from_none], obj.get("signature"))
        skip = from_union([from_bool, from_str, from_none], obj.get("skip"))
        soft_fail = from_union(
            [
                from_bool,
                lambda x: from_list(SoftFailElement.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("soft_fail"),
        )
        timeout_in_minutes = from_union(
            [from_int, from_none], obj.get("timeout_in_minutes")
        )
        type = from_union([CommandType, from_none], obj.get("type"))
        return CommandStep(
            agents,
            allow_dependency_failure,
            artifact_paths,
            branches,
            cache,
            cancel_on_build_failing,
            command,
            commands,
            concurrency,
            concurrency_group,
            concurrency_method,
            depends_on,
            env,
            id,
            identifier,
            command_step_if,
            key,
            label,
            matrix,
            name,
            notify,
            parallelism,
            plugins,
            priority,
            retry,
            signature,
            skip,
            soft_fail,
            timeout_in_minutes,
            type,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.agents is not None:
            result["agents"] = from_union(
                [
                    lambda x: from_dict(lambda x: x, x),
                    lambda x: from_list(from_str, x),
                    from_none,
                ],
                self.agents,
            )
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.artifact_paths is not None:
            result["artifact_paths"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none],
                self.artifact_paths,
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.cache is not None:
            result["cache"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CacheClass, x),
                    from_str,
                    from_none,
                ],
                self.cache,
            )
        if self.cancel_on_build_failing is not None:
            result["cancel_on_build_failing"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.cancel_on_build_failing,
            )
        if self.command is not None:
            result["command"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.command
            )
        if self.commands is not None:
            result["commands"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.commands
            )
        if self.concurrency is not None:
            result["concurrency"] = from_union([from_int, from_none], self.concurrency)
        if self.concurrency_group is not None:
            result["concurrency_group"] = from_union(
                [from_str, from_none], self.concurrency_group
            )
        if self.concurrency_method is not None:
            result["concurrency_method"] = from_union(
                [lambda x: to_enum(ConcurrencyMethod, x), from_none],
                self.concurrency_method,
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.env is not None:
            result["env"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.env
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.command_step_if is not None:
            result["if"] = from_union([from_str, from_none], self.command_step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_str, from_none], self.label)
        if self.matrix is not None:
            result["matrix"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union([from_int, from_bool, from_str], x), x
                    ),
                    lambda x: to_class(MatrixClass, x),
                    from_none,
                ],
                self.matrix,
            )
        if self.name is not None:
            result["name"] = from_union([from_str, from_none], self.name)
        if self.notify is not None:
            result["notify"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [
                                lambda x: to_class(NotifyClass, x),
                                lambda x: to_enum(NotifyEnum, x),
                            ],
                            x,
                        ),
                        x,
                    ),
                    from_none,
                ],
                self.notify,
            )
        if self.parallelism is not None:
            result["parallelism"] = from_union([from_int, from_none], self.parallelism)
        if self.plugins is not None:
            result["plugins"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: from_dict(lambda x: x, x), from_str], x
                        ),
                        x,
                    ),
                    lambda x: from_dict(lambda x: x, x),
                    from_none,
                ],
                self.plugins,
            )
        if self.priority is not None:
            result["priority"] = from_union([from_int, from_none], self.priority)
        if self.retry is not None:
            result["retry"] = from_union(
                [lambda x: to_class(Retry, x), from_none], self.retry
            )
        if self.signature is not None:
            result["signature"] = from_union(
                [lambda x: to_class(Signature, x), from_none], self.signature
            )
        if self.skip is not None:
            result["skip"] = from_union([from_bool, from_str, from_none], self.skip)
        if self.soft_fail is not None:
            result["soft_fail"] = from_union(
                [
                    from_bool,
                    lambda x: from_list(lambda x: to_class(SoftFailElement, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.soft_fail,
            )
        if self.timeout_in_minutes is not None:
            result["timeout_in_minutes"] = from_union(
                [from_int, from_none], self.timeout_in_minutes
            )
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(CommandType, x), from_none], self.type
            )
        return result


class InputType(Enum):
    INPUT = "input"


class InputStep:
    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    fields: Optional[List[Field]]
    id: Optional[str]
    identifier: Optional[str]
    input_step_if: Optional[str]
    input: Optional[str]
    """The label of the input step"""

    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]
    type: Optional[InputType]

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        branches: Optional[Union[List[str], str]],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        fields: Optional[List[Field]],
        id: Optional[str],
        identifier: Optional[str],
        input_step_if: Optional[str],
        input: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        prompt: Optional[str],
        type: Optional[InputType],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.branches = branches
        self.depends_on = depends_on
        self.fields = fields
        self.id = id
        self.identifier = identifier
        self.input_step_if = input_step_if
        self.input = input
        self.key = key
        self.label = label
        self.name = name
        self.prompt = prompt
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> "InputStep":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        fields = from_union(
            [lambda x: from_list(Field.from_dict, x), from_none], obj.get("fields")
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        input_step_if = from_union([from_str, from_none], obj.get("if"))
        input = from_union([from_str, from_none], obj.get("input"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_str, from_none], obj.get("label"))
        name = from_union([from_str, from_none], obj.get("name"))
        prompt = from_union([from_str, from_none], obj.get("prompt"))
        type = from_union([InputType, from_none], obj.get("type"))
        return InputStep(
            allow_dependency_failure,
            branches,
            depends_on,
            fields,
            id,
            identifier,
            input_step_if,
            input,
            key,
            label,
            name,
            prompt,
            type,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.fields is not None:
            result["fields"] = from_union(
                [lambda x: from_list(lambda x: to_class(Field, x), x), from_none],
                self.fields,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.input_step_if is not None:
            result["if"] = from_union([from_str, from_none], self.input_step_if)
        if self.input is not None:
            result["input"] = from_union([from_str, from_none], self.input)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_str, from_none], self.label)
        if self.name is not None:
            result["name"] = from_union([from_str, from_none], self.name)
        if self.prompt is not None:
            result["prompt"] = from_union([from_str, from_none], self.prompt)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(InputType, x), from_none], self.type
            )
        return result


class TentacledGithubCheck:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "TentacledGithubCheck":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return TentacledGithubCheck(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class TentacledGithubCommitStatus:
    context: Optional[str]
    """GitHub commit status name"""

    def __init__(self, context: Optional[str]) -> None:
        self.context = context

    @staticmethod
    def from_dict(obj: Any) -> "TentacledGithubCommitStatus":
        assert isinstance(obj, dict)
        context = from_union([from_str, from_none], obj.get("context"))
        return TentacledGithubCommitStatus(context)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.context is not None:
            result["context"] = from_union([from_str, from_none], self.context)
        return result


class TentacledSlack:
    channels: Optional[List[str]]
    message: Optional[str]

    def __init__(self, channels: Optional[List[str]], message: Optional[str]) -> None:
        self.channels = channels
        self.message = message

    @staticmethod
    def from_dict(obj: Any) -> "TentacledSlack":
        assert isinstance(obj, dict)
        channels = from_union(
            [lambda x: from_list(from_str, x), from_none], obj.get("channels")
        )
        message = from_union([from_str, from_none], obj.get("message"))
        return TentacledSlack(channels, message)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.channels is not None:
            result["channels"] = from_union(
                [lambda x: from_list(from_str, x), from_none], self.channels
            )
        if self.message is not None:
            result["message"] = from_union([from_str, from_none], self.message)
        return result


class FluffyBuildNotify:
    basecamp_campfire: Optional[str]
    build_notify_if: Optional[str]
    slack: Optional[Union[TentacledSlack, str]]
    github_commit_status: Optional[TentacledGithubCommitStatus]
    github_check: Optional[TentacledGithubCheck]
    email: Optional[str]
    webhook: Optional[str]
    pagerduty_change_event: Optional[str]

    def __init__(
        self,
        basecamp_campfire: Optional[str],
        build_notify_if: Optional[str],
        slack: Optional[Union[TentacledSlack, str]],
        github_commit_status: Optional[TentacledGithubCommitStatus],
        github_check: Optional[TentacledGithubCheck],
        email: Optional[str],
        webhook: Optional[str],
        pagerduty_change_event: Optional[str],
    ) -> None:
        self.basecamp_campfire = basecamp_campfire
        self.build_notify_if = build_notify_if
        self.slack = slack
        self.github_commit_status = github_commit_status
        self.github_check = github_check
        self.email = email
        self.webhook = webhook
        self.pagerduty_change_event = pagerduty_change_event

    @staticmethod
    def from_dict(obj: Any) -> "FluffyBuildNotify":
        assert isinstance(obj, dict)
        basecamp_campfire = from_union(
            [from_str, from_none], obj.get("basecamp_campfire")
        )
        build_notify_if = from_union([from_str, from_none], obj.get("if"))
        slack = from_union(
            [TentacledSlack.from_dict, from_str, from_none], obj.get("slack")
        )
        github_commit_status = from_union(
            [TentacledGithubCommitStatus.from_dict, from_none],
            obj.get("github_commit_status"),
        )
        github_check = from_union(
            [TentacledGithubCheck.from_dict, from_none], obj.get("github_check")
        )
        email = from_union([from_str, from_none], obj.get("email"))
        webhook = from_union([from_str, from_none], obj.get("webhook"))
        pagerduty_change_event = from_union(
            [from_str, from_none], obj.get("pagerduty_change_event")
        )
        return FluffyBuildNotify(
            basecamp_campfire,
            build_notify_if,
            slack,
            github_commit_status,
            github_check,
            email,
            webhook,
            pagerduty_change_event,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.basecamp_campfire is not None:
            result["basecamp_campfire"] = from_union(
                [from_str, from_none], self.basecamp_campfire
            )
        if self.build_notify_if is not None:
            result["if"] = from_union([from_str, from_none], self.build_notify_if)
        if self.slack is not None:
            result["slack"] = from_union(
                [lambda x: to_class(TentacledSlack, x), from_str, from_none], self.slack
            )
        if self.github_commit_status is not None:
            result["github_commit_status"] = from_union(
                [lambda x: to_class(TentacledGithubCommitStatus, x), from_none],
                self.github_commit_status,
            )
        if self.github_check is not None:
            result["github_check"] = from_union(
                [lambda x: to_class(TentacledGithubCheck, x), from_none],
                self.github_check,
            )
        if self.email is not None:
            result["email"] = from_union([from_str, from_none], self.email)
        if self.webhook is not None:
            result["webhook"] = from_union([from_str, from_none], self.webhook)
        if self.pagerduty_change_event is not None:
            result["pagerduty_change_event"] = from_union(
                [from_str, from_none], self.pagerduty_change_event
            )
        return result


class TriggerType(Enum):
    TRIGGER = "trigger"


class TriggerStep:
    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    trigger_step_async: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether to continue the build without waiting for the triggered step to complete"""

    branches: Optional[Union[List[str], str]]
    build: Optional[Build]
    """Properties of the build that will be created when the step is triggered"""

    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    trigger_step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]]
    trigger: str
    """The slug of the pipeline to create a build"""

    type: Optional[TriggerType]

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        trigger_step_async: Optional[Union[bool, AllowDependencyFailureEnum]],
        branches: Optional[Union[List[str], str]],
        build: Optional[Build],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        id: Optional[str],
        identifier: Optional[str],
        trigger_step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        skip: Optional[Union[bool, str]],
        soft_fail: Optional[
            Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]
        ],
        trigger: str,
        type: Optional[TriggerType],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.trigger_step_async = trigger_step_async
        self.branches = branches
        self.build = build
        self.depends_on = depends_on
        self.id = id
        self.identifier = identifier
        self.trigger_step_if = trigger_step_if
        self.key = key
        self.label = label
        self.name = name
        self.skip = skip
        self.soft_fail = soft_fail
        self.trigger = trigger
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> "TriggerStep":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        trigger_step_async = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("async")
        )
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        build = from_union([Build.from_dict, from_none], obj.get("build"))
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        trigger_step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_str, from_none], obj.get("label"))
        name = from_union([from_str, from_none], obj.get("name"))
        skip = from_union([from_bool, from_str, from_none], obj.get("skip"))
        soft_fail = from_union(
            [
                from_bool,
                lambda x: from_list(SoftFailElement.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("soft_fail"),
        )
        trigger = from_str(obj.get("trigger"))
        type = from_union([TriggerType, from_none], obj.get("type"))
        return TriggerStep(
            allow_dependency_failure,
            trigger_step_async,
            branches,
            build,
            depends_on,
            id,
            identifier,
            trigger_step_if,
            key,
            label,
            name,
            skip,
            soft_fail,
            trigger,
            type,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.trigger_step_async is not None:
            result["async"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.trigger_step_async,
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.build is not None:
            result["build"] = from_union(
                [lambda x: to_class(Build, x), from_none], self.build
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.trigger_step_if is not None:
            result["if"] = from_union([from_str, from_none], self.trigger_step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_str, from_none], self.label)
        if self.name is not None:
            result["name"] = from_union([from_str, from_none], self.name)
        if self.skip is not None:
            result["skip"] = from_union([from_bool, from_str, from_none], self.skip)
        if self.soft_fail is not None:
            result["soft_fail"] = from_union(
                [
                    from_bool,
                    lambda x: from_list(lambda x: to_class(SoftFailElement, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.soft_fail,
            )
        result["trigger"] = from_str(self.trigger)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(TriggerType, x), from_none], self.type
            )
        return result


class BlockStepType(Enum):
    BLOCK = "block"
    COMMAND = "command"
    COMMANDS = "commands"
    INPUT = "input"
    SCRIPT = "script"
    TRIGGER = "trigger"
    WAIT = "wait"
    WAITER = "waiter"


class WaitType(Enum):
    WAIT = "wait"
    WAITER = "waiter"


class WaitStep:
    """Waits for previous steps to pass before continuing"""

    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    branches: Optional[Union[List[str], str]]
    continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Continue to the next steps, even if the previous group of steps fail"""

    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    id: Optional[str]
    identifier: Optional[str]
    wait_step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    type: Optional[WaitType]
    wait: Optional[str]
    """Waits for previous steps to pass before continuing"""

    waiter: Optional[str]

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        branches: Optional[Union[List[str], str]],
        continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        id: Optional[str],
        identifier: Optional[str],
        wait_step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        type: Optional[WaitType],
        wait: Optional[str],
        waiter: Optional[str],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.branches = branches
        self.continue_on_failure = continue_on_failure
        self.depends_on = depends_on
        self.id = id
        self.identifier = identifier
        self.wait_step_if = wait_step_if
        self.key = key
        self.label = label
        self.name = name
        self.type = type
        self.wait = wait
        self.waiter = waiter

    @staticmethod
    def from_dict(obj: Any) -> "WaitStep":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        continue_on_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("continue_on_failure"),
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        wait_step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_none, from_str], obj.get("label"))
        name = from_union([from_none, from_str], obj.get("name"))
        type = from_union([WaitType, from_none], obj.get("type"))
        wait = from_union([from_none, from_str], obj.get("wait"))
        waiter = from_union([from_none, from_str], obj.get("waiter"))
        return WaitStep(
            allow_dependency_failure,
            branches,
            continue_on_failure,
            depends_on,
            id,
            identifier,
            wait_step_if,
            key,
            label,
            name,
            type,
            wait,
            waiter,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.continue_on_failure is not None:
            result["continue_on_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.continue_on_failure,
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.wait_step_if is not None:
            result["if"] = from_union([from_str, from_none], self.wait_step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_none, from_str], self.label)
        if self.name is not None:
            result["name"] = from_union([from_none, from_str], self.name)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(WaitType, x), from_none], self.type
            )
        if self.wait is not None:
            result["wait"] = from_union([from_none, from_str], self.wait)
        if self.waiter is not None:
            result["waiter"] = from_union([from_none, from_str], self.waiter)
        return result


class PurpleStep:
    """Waits for previous steps to pass before continuing"""

    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    block: Optional[Union[str, BlockStep]]
    """The label of the block step"""

    blocked_state: Optional[BlockedState]
    """The state that the build is set to when the build is blocked by this block step"""

    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    fields: Optional[List[Field]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]
    type: Optional[BlockStepType]
    input: Optional[Union[str, InputStep]]
    """The label of the input step"""

    agents: Optional[Union[Dict[str, Any], List[str]]]
    artifact_paths: Optional[Union[List[str], str]]
    """The glob path/s of artifacts to upload once this step has finished running"""

    cache: Optional[Union[List[str], CacheClass, str]]
    cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]]
    command: Optional[Union[List[str], CommandStep, str]]
    """The commands to run on the agent"""

    commands: Optional[Union[List[str], CommandStep, str]]
    """The commands to run on the agent"""

    concurrency: Optional[int]
    """The maximum number of jobs created from this step that are allowed to run at the same
    time. If you use this attribute, you must also define concurrency_group.
    """
    concurrency_group: Optional[str]
    """A unique name for the concurrency group that you are creating with the concurrency
    attribute
    """
    concurrency_method: Optional[ConcurrencyMethod]
    """Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
    this attribute, you must also define concurrency_group and concurrency.
    """
    env: Optional[Dict[str, Any]]
    matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]]
    notify: Optional[List[Union[NotifyClass, NotifyEnum]]]
    """Array of notification options for this step"""

    parallelism: Optional[int]
    """The number of parallel jobs that will be created based on this step"""

    plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]]
    priority: Optional[int]
    """Priority of the job, higher priorities are assigned to agents"""

    retry: Optional[Retry]
    """The conditions for retrying this step."""

    signature: Optional[Signature]
    """The signature of the command step, generally injected by agents at pipeline upload"""

    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]]
    timeout_in_minutes: Optional[int]
    """The number of minutes to time out a job"""

    script: Optional[CommandStep]
    continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Continue to the next steps, even if the previous group of steps fail"""

    wait: Optional[Union[WaitStep, str]]
    """Waits for previous steps to pass before continuing"""

    waiter: Optional[Union[WaitStep, str]]
    step_async: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether to continue the build without waiting for the triggered step to complete"""

    build: Optional[Build]
    """Properties of the build that will be created when the step is triggered"""

    trigger: Optional[Union[str, TriggerStep]]
    """The slug of the pipeline to create a build"""

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        block: Optional[Union[str, BlockStep]],
        blocked_state: Optional[BlockedState],
        branches: Optional[Union[List[str], str]],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        fields: Optional[List[Field]],
        id: Optional[str],
        identifier: Optional[str],
        step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        prompt: Optional[str],
        type: Optional[BlockStepType],
        input: Optional[Union[str, InputStep]],
        agents: Optional[Union[Dict[str, Any], List[str]]],
        artifact_paths: Optional[Union[List[str], str]],
        cache: Optional[Union[List[str], CacheClass, str]],
        cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]],
        command: Optional[Union[List[str], CommandStep, str]],
        commands: Optional[Union[List[str], CommandStep, str]],
        concurrency: Optional[int],
        concurrency_group: Optional[str],
        concurrency_method: Optional[ConcurrencyMethod],
        env: Optional[Dict[str, Any]],
        matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]],
        notify: Optional[List[Union[NotifyClass, NotifyEnum]]],
        parallelism: Optional[int],
        plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]],
        priority: Optional[int],
        retry: Optional[Retry],
        signature: Optional[Signature],
        skip: Optional[Union[bool, str]],
        soft_fail: Optional[
            Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]
        ],
        timeout_in_minutes: Optional[int],
        script: Optional[CommandStep],
        continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        wait: Optional[Union[WaitStep, str]],
        waiter: Optional[Union[WaitStep, str]],
        step_async: Optional[Union[bool, AllowDependencyFailureEnum]],
        build: Optional[Build],
        trigger: Optional[Union[str, TriggerStep]],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.block = block
        self.blocked_state = blocked_state
        self.branches = branches
        self.depends_on = depends_on
        self.fields = fields
        self.id = id
        self.identifier = identifier
        self.step_if = step_if
        self.key = key
        self.label = label
        self.name = name
        self.prompt = prompt
        self.type = type
        self.input = input
        self.agents = agents
        self.artifact_paths = artifact_paths
        self.cache = cache
        self.cancel_on_build_failing = cancel_on_build_failing
        self.command = command
        self.commands = commands
        self.concurrency = concurrency
        self.concurrency_group = concurrency_group
        self.concurrency_method = concurrency_method
        self.env = env
        self.matrix = matrix
        self.notify = notify
        self.parallelism = parallelism
        self.plugins = plugins
        self.priority = priority
        self.retry = retry
        self.signature = signature
        self.skip = skip
        self.soft_fail = soft_fail
        self.timeout_in_minutes = timeout_in_minutes
        self.script = script
        self.continue_on_failure = continue_on_failure
        self.wait = wait
        self.waiter = waiter
        self.step_async = step_async
        self.build = build
        self.trigger = trigger

    @staticmethod
    def from_dict(obj: Any) -> "PurpleStep":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        block = from_union([from_str, BlockStep.from_dict, from_none], obj.get("block"))
        blocked_state = from_union([BlockedState, from_none], obj.get("blocked_state"))
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        fields = from_union(
            [lambda x: from_list(Field.from_dict, x), from_none], obj.get("fields")
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_none, from_str], obj.get("label"))
        name = from_union([from_none, from_str], obj.get("name"))
        prompt = from_union([from_str, from_none], obj.get("prompt"))
        type = from_union([BlockStepType, from_none], obj.get("type"))
        input = from_union([from_str, InputStep.from_dict, from_none], obj.get("input"))
        agents = from_union(
            [
                lambda x: from_dict(lambda x: x, x),
                lambda x: from_list(from_str, x),
                from_none,
            ],
            obj.get("agents"),
        )
        artifact_paths = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none],
            obj.get("artifact_paths"),
        )
        cache = from_union(
            [
                lambda x: from_list(from_str, x),
                CacheClass.from_dict,
                from_str,
                from_none,
            ],
            obj.get("cache"),
        )
        cancel_on_build_failing = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("cancel_on_build_failing"),
        )
        command = from_union(
            [
                lambda x: from_list(from_str, x),
                CommandStep.from_dict,
                from_str,
                from_none,
            ],
            obj.get("command"),
        )
        commands = from_union(
            [
                lambda x: from_list(from_str, x),
                CommandStep.from_dict,
                from_str,
                from_none,
            ],
            obj.get("commands"),
        )
        concurrency = from_union([from_int, from_none], obj.get("concurrency"))
        concurrency_group = from_union(
            [from_str, from_none], obj.get("concurrency_group")
        )
        concurrency_method = from_union(
            [ConcurrencyMethod, from_none], obj.get("concurrency_method")
        )
        env = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("env")
        )
        matrix = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                MatrixClass.from_dict,
                from_none,
            ],
            obj.get("matrix"),
        )
        notify = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([NotifyClass.from_dict, NotifyEnum], x), x
                ),
                from_none,
            ],
            obj.get("notify"),
        )
        parallelism = from_union([from_int, from_none], obj.get("parallelism"))
        plugins = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union(
                        [lambda x: from_dict(lambda x: x, x), from_str], x
                    ),
                    x,
                ),
                lambda x: from_dict(lambda x: x, x),
                from_none,
            ],
            obj.get("plugins"),
        )
        priority = from_union([from_int, from_none], obj.get("priority"))
        retry = from_union([Retry.from_dict, from_none], obj.get("retry"))
        signature = from_union([Signature.from_dict, from_none], obj.get("signature"))
        skip = from_union([from_bool, from_str, from_none], obj.get("skip"))
        soft_fail = from_union(
            [
                from_bool,
                lambda x: from_list(SoftFailElement.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("soft_fail"),
        )
        timeout_in_minutes = from_union(
            [from_int, from_none], obj.get("timeout_in_minutes")
        )
        script = from_union([CommandStep.from_dict, from_none], obj.get("script"))
        continue_on_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("continue_on_failure"),
        )
        wait = from_union([from_none, WaitStep.from_dict, from_str], obj.get("wait"))
        waiter = from_union(
            [from_none, WaitStep.from_dict, from_str], obj.get("waiter")
        )
        step_async = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("async")
        )
        build = from_union([Build.from_dict, from_none], obj.get("build"))
        trigger = from_union(
            [from_str, TriggerStep.from_dict, from_none], obj.get("trigger")
        )
        return PurpleStep(
            allow_dependency_failure,
            block,
            blocked_state,
            branches,
            depends_on,
            fields,
            id,
            identifier,
            step_if,
            key,
            label,
            name,
            prompt,
            type,
            input,
            agents,
            artifact_paths,
            cache,
            cancel_on_build_failing,
            command,
            commands,
            concurrency,
            concurrency_group,
            concurrency_method,
            env,
            matrix,
            notify,
            parallelism,
            plugins,
            priority,
            retry,
            signature,
            skip,
            soft_fail,
            timeout_in_minutes,
            script,
            continue_on_failure,
            wait,
            waiter,
            step_async,
            build,
            trigger,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.block is not None:
            result["block"] = from_union(
                [from_str, lambda x: to_class(BlockStep, x), from_none], self.block
            )
        if self.blocked_state is not None:
            result["blocked_state"] = from_union(
                [lambda x: to_enum(BlockedState, x), from_none], self.blocked_state
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.fields is not None:
            result["fields"] = from_union(
                [lambda x: from_list(lambda x: to_class(Field, x), x), from_none],
                self.fields,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.step_if is not None:
            result["if"] = from_union([from_str, from_none], self.step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_none, from_str], self.label)
        if self.name is not None:
            result["name"] = from_union([from_none, from_str], self.name)
        if self.prompt is not None:
            result["prompt"] = from_union([from_str, from_none], self.prompt)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(BlockStepType, x), from_none], self.type
            )
        if self.input is not None:
            result["input"] = from_union(
                [from_str, lambda x: to_class(InputStep, x), from_none], self.input
            )
        if self.agents is not None:
            result["agents"] = from_union(
                [
                    lambda x: from_dict(lambda x: x, x),
                    lambda x: from_list(from_str, x),
                    from_none,
                ],
                self.agents,
            )
        if self.artifact_paths is not None:
            result["artifact_paths"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none],
                self.artifact_paths,
            )
        if self.cache is not None:
            result["cache"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CacheClass, x),
                    from_str,
                    from_none,
                ],
                self.cache,
            )
        if self.cancel_on_build_failing is not None:
            result["cancel_on_build_failing"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.cancel_on_build_failing,
            )
        if self.command is not None:
            result["command"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CommandStep, x),
                    from_str,
                    from_none,
                ],
                self.command,
            )
        if self.commands is not None:
            result["commands"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CommandStep, x),
                    from_str,
                    from_none,
                ],
                self.commands,
            )
        if self.concurrency is not None:
            result["concurrency"] = from_union([from_int, from_none], self.concurrency)
        if self.concurrency_group is not None:
            result["concurrency_group"] = from_union(
                [from_str, from_none], self.concurrency_group
            )
        if self.concurrency_method is not None:
            result["concurrency_method"] = from_union(
                [lambda x: to_enum(ConcurrencyMethod, x), from_none],
                self.concurrency_method,
            )
        if self.env is not None:
            result["env"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.env
            )
        if self.matrix is not None:
            result["matrix"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union([from_int, from_bool, from_str], x), x
                    ),
                    lambda x: to_class(MatrixClass, x),
                    from_none,
                ],
                self.matrix,
            )
        if self.notify is not None:
            result["notify"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [
                                lambda x: to_class(NotifyClass, x),
                                lambda x: to_enum(NotifyEnum, x),
                            ],
                            x,
                        ),
                        x,
                    ),
                    from_none,
                ],
                self.notify,
            )
        if self.parallelism is not None:
            result["parallelism"] = from_union([from_int, from_none], self.parallelism)
        if self.plugins is not None:
            result["plugins"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: from_dict(lambda x: x, x), from_str], x
                        ),
                        x,
                    ),
                    lambda x: from_dict(lambda x: x, x),
                    from_none,
                ],
                self.plugins,
            )
        if self.priority is not None:
            result["priority"] = from_union([from_int, from_none], self.priority)
        if self.retry is not None:
            result["retry"] = from_union(
                [lambda x: to_class(Retry, x), from_none], self.retry
            )
        if self.signature is not None:
            result["signature"] = from_union(
                [lambda x: to_class(Signature, x), from_none], self.signature
            )
        if self.skip is not None:
            result["skip"] = from_union([from_bool, from_str, from_none], self.skip)
        if self.soft_fail is not None:
            result["soft_fail"] = from_union(
                [
                    from_bool,
                    lambda x: from_list(lambda x: to_class(SoftFailElement, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.soft_fail,
            )
        if self.timeout_in_minutes is not None:
            result["timeout_in_minutes"] = from_union(
                [from_int, from_none], self.timeout_in_minutes
            )
        if self.script is not None:
            result["script"] = from_union(
                [lambda x: to_class(CommandStep, x), from_none], self.script
            )
        if self.continue_on_failure is not None:
            result["continue_on_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.continue_on_failure,
            )
        if self.wait is not None:
            result["wait"] = from_union(
                [from_none, lambda x: to_class(WaitStep, x), from_str], self.wait
            )
        if self.waiter is not None:
            result["waiter"] = from_union(
                [from_none, lambda x: to_class(WaitStep, x), from_str], self.waiter
            )
        if self.step_async is not None:
            result["async"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.step_async,
            )
        if self.build is not None:
            result["build"] = from_union(
                [lambda x: to_class(Build, x), from_none], self.build
            )
        if self.trigger is not None:
            result["trigger"] = from_union(
                [from_str, lambda x: to_class(TriggerStep, x), from_none], self.trigger
            )
        return result


class StringStep(Enum):
    """Pauses the execution of a build and waits on a user to unblock it

    Waits for previous steps to pass before continuing
    """

    BLOCK = "block"
    INPUT = "input"
    WAIT = "wait"
    WAITER = "waiter"


class GroupStepClass:
    """Waits for previous steps to pass before continuing"""

    allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    block: Optional[Union[str, BlockStep]]
    """The label of the block step"""

    blocked_state: Optional[BlockedState]
    """The state that the build is set to when the build is blocked by this block step"""

    branches: Optional[Union[List[str], str]]
    depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]]
    fields: Optional[List[Field]]
    id: Optional[str]
    identifier: Optional[str]
    step_if: Optional[str]
    key: Optional[str]
    label: Optional[str]
    name: Optional[str]
    prompt: Optional[str]
    type: Optional[BlockStepType]
    input: Optional[Union[str, InputStep]]
    """The label of the input step"""

    agents: Optional[Union[Dict[str, Any], List[str]]]
    artifact_paths: Optional[Union[List[str], str]]
    """The glob path/s of artifacts to upload once this step has finished running"""

    cache: Optional[Union[List[str], CacheClass, str]]
    cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]]
    command: Optional[Union[List[str], CommandStep, str]]
    """The commands to run on the agent"""

    commands: Optional[Union[List[str], CommandStep, str]]
    """The commands to run on the agent"""

    concurrency: Optional[int]
    """The maximum number of jobs created from this step that are allowed to run at the same
    time. If you use this attribute, you must also define concurrency_group.
    """
    concurrency_group: Optional[str]
    """A unique name for the concurrency group that you are creating with the concurrency
    attribute
    """
    concurrency_method: Optional[ConcurrencyMethod]
    """Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
    this attribute, you must also define concurrency_group and concurrency.
    """
    env: Optional[Dict[str, Any]]
    matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]]
    notify: Optional[List[Union[FluffyBuildNotify, NotifyEnum]]]
    """Array of notification options for this step"""

    parallelism: Optional[int]
    """The number of parallel jobs that will be created based on this step"""

    plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]]
    priority: Optional[int]
    """Priority of the job, higher priorities are assigned to agents"""

    retry: Optional[Retry]
    """The conditions for retrying this step."""

    signature: Optional[Signature]
    """The signature of the command step, generally injected by agents at pipeline upload"""

    skip: Optional[Union[bool, str]]
    soft_fail: Optional[Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]]
    timeout_in_minutes: Optional[int]
    """The number of minutes to time out a job"""

    script: Optional[CommandStep]
    continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Continue to the next steps, even if the previous group of steps fail"""

    wait: Optional[Union[WaitStep, str]]
    """Waits for previous steps to pass before continuing"""

    waiter: Optional[Union[WaitStep, str]]
    step_async: Optional[Union[bool, AllowDependencyFailureEnum]]
    """Whether to continue the build without waiting for the triggered step to complete"""

    build: Optional[Build]
    """Properties of the build that will be created when the step is triggered"""

    trigger: Optional[Union[str, TriggerStep]]
    """The slug of the pipeline to create a build"""

    group: Optional[str]
    """The name to give to this group of steps"""

    steps: Optional[List[Union[PurpleStep, StringStep]]]
    """A list of steps"""

    def __init__(
        self,
        allow_dependency_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        block: Optional[Union[str, BlockStep]],
        blocked_state: Optional[BlockedState],
        branches: Optional[Union[List[str], str]],
        depends_on: Optional[Union[List[Union[DependsOnClass, str]], str]],
        fields: Optional[List[Field]],
        id: Optional[str],
        identifier: Optional[str],
        step_if: Optional[str],
        key: Optional[str],
        label: Optional[str],
        name: Optional[str],
        prompt: Optional[str],
        type: Optional[BlockStepType],
        input: Optional[Union[str, InputStep]],
        agents: Optional[Union[Dict[str, Any], List[str]]],
        artifact_paths: Optional[Union[List[str], str]],
        cache: Optional[Union[List[str], CacheClass, str]],
        cancel_on_build_failing: Optional[Union[bool, AllowDependencyFailureEnum]],
        command: Optional[Union[List[str], CommandStep, str]],
        commands: Optional[Union[List[str], CommandStep, str]],
        concurrency: Optional[int],
        concurrency_group: Optional[str],
        concurrency_method: Optional[ConcurrencyMethod],
        env: Optional[Dict[str, Any]],
        matrix: Optional[Union[List[Union[int, bool, str]], MatrixClass]],
        notify: Optional[List[Union[FluffyBuildNotify, NotifyEnum]]],
        parallelism: Optional[int],
        plugins: Optional[Union[List[Union[Dict[str, Any], str]], Dict[str, Any]]],
        priority: Optional[int],
        retry: Optional[Retry],
        signature: Optional[Signature],
        skip: Optional[Union[bool, str]],
        soft_fail: Optional[
            Union[bool, List[SoftFailElement], AllowDependencyFailureEnum]
        ],
        timeout_in_minutes: Optional[int],
        script: Optional[CommandStep],
        continue_on_failure: Optional[Union[bool, AllowDependencyFailureEnum]],
        wait: Optional[Union[WaitStep, str]],
        waiter: Optional[Union[WaitStep, str]],
        step_async: Optional[Union[bool, AllowDependencyFailureEnum]],
        build: Optional[Build],
        trigger: Optional[Union[str, TriggerStep]],
        group: Optional[str],
        steps: Optional[List[Union[PurpleStep, StringStep]]],
    ) -> None:
        self.allow_dependency_failure = allow_dependency_failure
        self.block = block
        self.blocked_state = blocked_state
        self.branches = branches
        self.depends_on = depends_on
        self.fields = fields
        self.id = id
        self.identifier = identifier
        self.step_if = step_if
        self.key = key
        self.label = label
        self.name = name
        self.prompt = prompt
        self.type = type
        self.input = input
        self.agents = agents
        self.artifact_paths = artifact_paths
        self.cache = cache
        self.cancel_on_build_failing = cancel_on_build_failing
        self.command = command
        self.commands = commands
        self.concurrency = concurrency
        self.concurrency_group = concurrency_group
        self.concurrency_method = concurrency_method
        self.env = env
        self.matrix = matrix
        self.notify = notify
        self.parallelism = parallelism
        self.plugins = plugins
        self.priority = priority
        self.retry = retry
        self.signature = signature
        self.skip = skip
        self.soft_fail = soft_fail
        self.timeout_in_minutes = timeout_in_minutes
        self.script = script
        self.continue_on_failure = continue_on_failure
        self.wait = wait
        self.waiter = waiter
        self.step_async = step_async
        self.build = build
        self.trigger = trigger
        self.group = group
        self.steps = steps

    @staticmethod
    def from_dict(obj: Any) -> "GroupStepClass":
        assert isinstance(obj, dict)
        allow_dependency_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("allow_dependency_failure"),
        )
        block = from_union([from_str, BlockStep.from_dict, from_none], obj.get("block"))
        blocked_state = from_union([BlockedState, from_none], obj.get("blocked_state"))
        branches = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none], obj.get("branches")
        )
        depends_on = from_union(
            [
                from_none,
                lambda x: from_list(
                    lambda x: from_union([DependsOnClass.from_dict, from_str], x), x
                ),
                from_str,
            ],
            obj.get("depends_on"),
        )
        fields = from_union(
            [lambda x: from_list(Field.from_dict, x), from_none], obj.get("fields")
        )
        id = from_union([from_str, from_none], obj.get("id"))
        identifier = from_union([from_str, from_none], obj.get("identifier"))
        step_if = from_union([from_str, from_none], obj.get("if"))
        key = from_union([from_str, from_none], obj.get("key"))
        label = from_union([from_none, from_str], obj.get("label"))
        name = from_union([from_none, from_str], obj.get("name"))
        prompt = from_union([from_str, from_none], obj.get("prompt"))
        type = from_union([BlockStepType, from_none], obj.get("type"))
        input = from_union([from_str, InputStep.from_dict, from_none], obj.get("input"))
        agents = from_union(
            [
                lambda x: from_dict(lambda x: x, x),
                lambda x: from_list(from_str, x),
                from_none,
            ],
            obj.get("agents"),
        )
        artifact_paths = from_union(
            [lambda x: from_list(from_str, x), from_str, from_none],
            obj.get("artifact_paths"),
        )
        cache = from_union(
            [
                lambda x: from_list(from_str, x),
                CacheClass.from_dict,
                from_str,
                from_none,
            ],
            obj.get("cache"),
        )
        cancel_on_build_failing = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("cancel_on_build_failing"),
        )
        command = from_union(
            [
                lambda x: from_list(from_str, x),
                CommandStep.from_dict,
                from_str,
                from_none,
            ],
            obj.get("command"),
        )
        commands = from_union(
            [
                lambda x: from_list(from_str, x),
                CommandStep.from_dict,
                from_str,
                from_none,
            ],
            obj.get("commands"),
        )
        concurrency = from_union([from_int, from_none], obj.get("concurrency"))
        concurrency_group = from_union(
            [from_str, from_none], obj.get("concurrency_group")
        )
        concurrency_method = from_union(
            [ConcurrencyMethod, from_none], obj.get("concurrency_method")
        )
        env = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("env")
        )
        matrix = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([from_int, from_bool, from_str], x), x
                ),
                MatrixClass.from_dict,
                from_none,
            ],
            obj.get("matrix"),
        )
        notify = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([FluffyBuildNotify.from_dict, NotifyEnum], x),
                    x,
                ),
                from_none,
            ],
            obj.get("notify"),
        )
        parallelism = from_union([from_int, from_none], obj.get("parallelism"))
        plugins = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union(
                        [lambda x: from_dict(lambda x: x, x), from_str], x
                    ),
                    x,
                ),
                lambda x: from_dict(lambda x: x, x),
                from_none,
            ],
            obj.get("plugins"),
        )
        priority = from_union([from_int, from_none], obj.get("priority"))
        retry = from_union([Retry.from_dict, from_none], obj.get("retry"))
        signature = from_union([Signature.from_dict, from_none], obj.get("signature"))
        skip = from_union([from_bool, from_str, from_none], obj.get("skip"))
        soft_fail = from_union(
            [
                from_bool,
                lambda x: from_list(SoftFailElement.from_dict, x),
                AllowDependencyFailureEnum,
                from_none,
            ],
            obj.get("soft_fail"),
        )
        timeout_in_minutes = from_union(
            [from_int, from_none], obj.get("timeout_in_minutes")
        )
        script = from_union([CommandStep.from_dict, from_none], obj.get("script"))
        continue_on_failure = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none],
            obj.get("continue_on_failure"),
        )
        wait = from_union([from_none, WaitStep.from_dict, from_str], obj.get("wait"))
        waiter = from_union(
            [from_none, WaitStep.from_dict, from_str], obj.get("waiter")
        )
        step_async = from_union(
            [from_bool, AllowDependencyFailureEnum, from_none], obj.get("async")
        )
        build = from_union([Build.from_dict, from_none], obj.get("build"))
        trigger = from_union(
            [from_str, TriggerStep.from_dict, from_none], obj.get("trigger")
        )
        group = from_union([from_none, from_str], obj.get("group"))
        steps = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([PurpleStep.from_dict, StringStep], x), x
                ),
                from_none,
            ],
            obj.get("steps"),
        )
        return GroupStepClass(
            allow_dependency_failure,
            block,
            blocked_state,
            branches,
            depends_on,
            fields,
            id,
            identifier,
            step_if,
            key,
            label,
            name,
            prompt,
            type,
            input,
            agents,
            artifact_paths,
            cache,
            cancel_on_build_failing,
            command,
            commands,
            concurrency,
            concurrency_group,
            concurrency_method,
            env,
            matrix,
            notify,
            parallelism,
            plugins,
            priority,
            retry,
            signature,
            skip,
            soft_fail,
            timeout_in_minutes,
            script,
            continue_on_failure,
            wait,
            waiter,
            step_async,
            build,
            trigger,
            group,
            steps,
        )

    def to_dict(self) -> dict:
        result: dict = {}
        if self.allow_dependency_failure is not None:
            result["allow_dependency_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.allow_dependency_failure,
            )
        if self.block is not None:
            result["block"] = from_union(
                [from_str, lambda x: to_class(BlockStep, x), from_none], self.block
            )
        if self.blocked_state is not None:
            result["blocked_state"] = from_union(
                [lambda x: to_enum(BlockedState, x), from_none], self.blocked_state
            )
        if self.branches is not None:
            result["branches"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none], self.branches
            )
        if self.depends_on is not None:
            result["depends_on"] = from_union(
                [
                    from_none,
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: to_class(DependsOnClass, x), from_str], x
                        ),
                        x,
                    ),
                    from_str,
                ],
                self.depends_on,
            )
        if self.fields is not None:
            result["fields"] = from_union(
                [lambda x: from_list(lambda x: to_class(Field, x), x), from_none],
                self.fields,
            )
        if self.id is not None:
            result["id"] = from_union([from_str, from_none], self.id)
        if self.identifier is not None:
            result["identifier"] = from_union([from_str, from_none], self.identifier)
        if self.step_if is not None:
            result["if"] = from_union([from_str, from_none], self.step_if)
        if self.key is not None:
            result["key"] = from_union([from_str, from_none], self.key)
        if self.label is not None:
            result["label"] = from_union([from_none, from_str], self.label)
        if self.name is not None:
            result["name"] = from_union([from_none, from_str], self.name)
        if self.prompt is not None:
            result["prompt"] = from_union([from_str, from_none], self.prompt)
        if self.type is not None:
            result["type"] = from_union(
                [lambda x: to_enum(BlockStepType, x), from_none], self.type
            )
        if self.input is not None:
            result["input"] = from_union(
                [from_str, lambda x: to_class(InputStep, x), from_none], self.input
            )
        if self.agents is not None:
            result["agents"] = from_union(
                [
                    lambda x: from_dict(lambda x: x, x),
                    lambda x: from_list(from_str, x),
                    from_none,
                ],
                self.agents,
            )
        if self.artifact_paths is not None:
            result["artifact_paths"] = from_union(
                [lambda x: from_list(from_str, x), from_str, from_none],
                self.artifact_paths,
            )
        if self.cache is not None:
            result["cache"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CacheClass, x),
                    from_str,
                    from_none,
                ],
                self.cache,
            )
        if self.cancel_on_build_failing is not None:
            result["cancel_on_build_failing"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.cancel_on_build_failing,
            )
        if self.command is not None:
            result["command"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CommandStep, x),
                    from_str,
                    from_none,
                ],
                self.command,
            )
        if self.commands is not None:
            result["commands"] = from_union(
                [
                    lambda x: from_list(from_str, x),
                    lambda x: to_class(CommandStep, x),
                    from_str,
                    from_none,
                ],
                self.commands,
            )
        if self.concurrency is not None:
            result["concurrency"] = from_union([from_int, from_none], self.concurrency)
        if self.concurrency_group is not None:
            result["concurrency_group"] = from_union(
                [from_str, from_none], self.concurrency_group
            )
        if self.concurrency_method is not None:
            result["concurrency_method"] = from_union(
                [lambda x: to_enum(ConcurrencyMethod, x), from_none],
                self.concurrency_method,
            )
        if self.env is not None:
            result["env"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.env
            )
        if self.matrix is not None:
            result["matrix"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union([from_int, from_bool, from_str], x), x
                    ),
                    lambda x: to_class(MatrixClass, x),
                    from_none,
                ],
                self.matrix,
            )
        if self.notify is not None:
            result["notify"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [
                                lambda x: to_class(FluffyBuildNotify, x),
                                lambda x: to_enum(NotifyEnum, x),
                            ],
                            x,
                        ),
                        x,
                    ),
                    from_none,
                ],
                self.notify,
            )
        if self.parallelism is not None:
            result["parallelism"] = from_union([from_int, from_none], self.parallelism)
        if self.plugins is not None:
            result["plugins"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [lambda x: from_dict(lambda x: x, x), from_str], x
                        ),
                        x,
                    ),
                    lambda x: from_dict(lambda x: x, x),
                    from_none,
                ],
                self.plugins,
            )
        if self.priority is not None:
            result["priority"] = from_union([from_int, from_none], self.priority)
        if self.retry is not None:
            result["retry"] = from_union(
                [lambda x: to_class(Retry, x), from_none], self.retry
            )
        if self.signature is not None:
            result["signature"] = from_union(
                [lambda x: to_class(Signature, x), from_none], self.signature
            )
        if self.skip is not None:
            result["skip"] = from_union([from_bool, from_str, from_none], self.skip)
        if self.soft_fail is not None:
            result["soft_fail"] = from_union(
                [
                    from_bool,
                    lambda x: from_list(lambda x: to_class(SoftFailElement, x), x),
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.soft_fail,
            )
        if self.timeout_in_minutes is not None:
            result["timeout_in_minutes"] = from_union(
                [from_int, from_none], self.timeout_in_minutes
            )
        if self.script is not None:
            result["script"] = from_union(
                [lambda x: to_class(CommandStep, x), from_none], self.script
            )
        if self.continue_on_failure is not None:
            result["continue_on_failure"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.continue_on_failure,
            )
        if self.wait is not None:
            result["wait"] = from_union(
                [from_none, lambda x: to_class(WaitStep, x), from_str], self.wait
            )
        if self.waiter is not None:
            result["waiter"] = from_union(
                [from_none, lambda x: to_class(WaitStep, x), from_str], self.waiter
            )
        if self.step_async is not None:
            result["async"] = from_union(
                [
                    from_bool,
                    lambda x: to_enum(AllowDependencyFailureEnum, x),
                    from_none,
                ],
                self.step_async,
            )
        if self.build is not None:
            result["build"] = from_union(
                [lambda x: to_class(Build, x), from_none], self.build
            )
        if self.trigger is not None:
            result["trigger"] = from_union(
                [from_str, lambda x: to_class(TriggerStep, x), from_none], self.trigger
            )
        if self.group is not None:
            result["group"] = from_union([from_none, from_str], self.group)
        if self.steps is not None:
            result["steps"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [
                                lambda x: to_class(PurpleStep, x),
                                lambda x: to_enum(StringStep, x),
                            ],
                            x,
                        ),
                        x,
                    ),
                    from_none,
                ],
                self.steps,
            )
        return result


class Schema:
    agents: Optional[Union[Dict[str, Any], List[str]]]
    env: Optional[Dict[str, Any]]
    notify: Optional[List[Union[PurpleBuildNotify, NotifyEnum]]]
    steps: List[Union[GroupStepClass, StringStep]]
    """A list of steps"""

    def __init__(
        self,
        agents: Optional[Union[Dict[str, Any], List[str]]],
        env: Optional[Dict[str, Any]],
        notify: Optional[List[Union[PurpleBuildNotify, NotifyEnum]]],
        steps: List[Union[GroupStepClass, StringStep]],
    ) -> None:
        self.agents = agents
        self.env = env
        self.notify = notify
        self.steps = steps

    @staticmethod
    def from_dict(obj: Any) -> "Schema":
        assert isinstance(obj, dict)
        agents = from_union(
            [
                lambda x: from_dict(lambda x: x, x),
                lambda x: from_list(from_str, x),
                from_none,
            ],
            obj.get("agents"),
        )
        env = from_union(
            [lambda x: from_dict(lambda x: x, x), from_none], obj.get("env")
        )
        notify = from_union(
            [
                lambda x: from_list(
                    lambda x: from_union([PurpleBuildNotify.from_dict, NotifyEnum], x),
                    x,
                ),
                from_none,
            ],
            obj.get("notify"),
        )
        steps = from_list(
            lambda x: from_union([GroupStepClass.from_dict, StringStep], x),
            obj.get("steps"),
        )
        return Schema(agents, env, notify, steps)

    def to_dict(self) -> dict:
        result: dict = {}
        if self.agents is not None:
            result["agents"] = from_union(
                [
                    lambda x: from_dict(lambda x: x, x),
                    lambda x: from_list(from_str, x),
                    from_none,
                ],
                self.agents,
            )
        if self.env is not None:
            result["env"] = from_union(
                [lambda x: from_dict(lambda x: x, x), from_none], self.env
            )
        if self.notify is not None:
            result["notify"] = from_union(
                [
                    lambda x: from_list(
                        lambda x: from_union(
                            [
                                lambda x: to_class(PurpleBuildNotify, x),
                                lambda x: to_enum(NotifyEnum, x),
                            ],
                            x,
                        ),
                        x,
                    ),
                    from_none,
                ],
                self.notify,
            )
        result["steps"] = from_list(
            lambda x: from_union(
                [
                    lambda x: to_class(GroupStepClass, x),
                    lambda x: to_enum(StringStep, x),
                ],
                x,
            ),
            self.steps,
        )
        return result


def schema_from_dict(s: Any) -> Schema:
    return Schema.from_dict(s)


def schema_to_dict(x: Schema) -> Any:
    return to_class(Schema, x)
