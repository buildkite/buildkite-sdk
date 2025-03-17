from typing import Optional, Union, List, Dict, Any
from buildkite_sdk.schema import (
    Adjustment as _matrix_adjustment,
    AutomaticRetry as _automatic_retry,
    BlockedState as _blocked_state,
    Build as _build,
    CacheClass as _cache_class,
    ConcurrencyMethod as _concurrency_method,
    DependsOnClass as _depends_on_class,
    ExitStatusEnum as _exit_status_enum,
    Field as _field,
    FluffyBuildNotify as _fluffy_build_notify,
    ManualClass as _manual_retry,
    MatrixClass as _matrix_class,
    NotifyEnum as _notify_enum,
    Option as _select_option,
    PurpleBuildNotify as _purple_build_notify,
    PurpleGithubCommitStatus as _purple_github_commit_status,
    PurpleSlack as _purple_slack,
    Retry as _retry,
    SignalReason as _signal_reason,
    Signature as _signature,
    SoftFailElement as _soft_fail,
    TentacledSlack as _tentacled_slack,
)

BlockedStateEnum = _blocked_state
ConcurrencyMethod = _concurrency_method
ExitStatusEnum = _exit_status_enum
NotifyEnum = _notify_enum
SignalReasonEnum = _signal_reason


class Cache(_cache_class):
    def __init__(
        self,
        paths: List[str],
        name: Optional[str] = None,
        size: Optional[str] = None,
    ):
        super().__init__(name, paths, size)


class DependsOn(_depends_on_class):
    def __init__(
        self,
        step: str,
        allow_failure: Optional[bool] = None,
    ) -> None:
        super.__init__(
            self,
            allow_failure,
            step,
        )


class TextField(_field):
    def __init__(
        self,
        key: str,
        default: Optional[str] = None,
        hint: Optional[str] = None,
        required: Optional[bool] = None,
        text: Optional[str] = None,
    ) -> None:
        super().__init__(
            default=default,
            hint=hint,
            key=key,
            required=required,
            text=text,
            multiple=None,
            options=None,
            select=None,
            format=None,
        )


class SelectFieldOption(_select_option):
    def __init__(
        self,
        label: str,
        value: str,
        hint: Optional[str] = None,
        required: Optional[bool] = None,
    ):
        super().__init__(hint, label, required, value)


class SelectField(_field):
    def __init__(
        self,
        name: str,
        key: str,
        options: List[SelectFieldOption],
        default: Optional[str] = None,
        hint: Optional[str] = None,
        required: Optional[bool] = None,
        multiple: Optional[bool] = None,
    ) -> None:
        super().__init__(
            default,
            hint,
            key,
            required,
            multiple,
            options,
            select=name,
        )


class SoftFail(_soft_fail):
    def __init__(self, exit_status: Union[str, int]):
        super().__init__(exit_status)


class NotifySlack(_tentacled_slack):
    def __init__(self, channels: Optional[List[str]], message: Optional[str]) -> None:
        super().__init__(channels, message)

    def _to_pipeline_notify(self) -> _purple_slack:
        return _purple_slack(
            channels=self.channels,
            message=self.message,
        )


class StepNotify(_fluffy_build_notify):
    def __init__(
        self,
        slack: Optional[Union[NotifySlack, str]],
    ) -> None:
        super().__init__(
            slack,
            basecamp_campfire=None,
            build_notify_if=None,
            github_commit_status=None,
            github_check=None,
            email=None,
            webhook=None,
            pagerduty_change_event=None,
        )


class NotifyGitHubCommitStatus(_purple_github_commit_status):
    def __init__(self, context: Optional[str]):
        super().__init__(context=context)


class PipelineNotify(_purple_build_notify):
    def __init__(
        self,
        email: Optional[str],
        build_notify_if: Optional[str],
        basecamp_campfire: Optional[str],
        slack: Optional[NotifySlack],
        webhook: Optional[str],
        pagerduty_change_event: Optional[str],
        github_commit_status: Optional[NotifyGitHubCommitStatus],
        github_check: Optional[Dict[str, Any]],
    ) -> None:
        super().__init__(
            email=email,
            build_notify_if=build_notify_if,
            basecamp_campfire=basecamp_campfire,
            slack=slack._to_pipeline_notify(),
            webhook=webhook,
            pagerduty_change_event=pagerduty_change_event,
            github_commit_status=github_commit_status,
            github_check=github_check,
        )


# Matrix
class MatrixAdjustment(_matrix_adjustment):
    def __init__(
        self,
        adjustment_with: Union[List[Union[int, bool, str]], Dict[str, str]],
        skip: Optional[bool] = None,
        soft_fail: Optional[Union[SoftFail, bool]] = None,
    ) -> None:
        super().__init__(skip, soft_fail, adjustment_with)


class MatrixAdvanced(_matrix_class):
    def __init__(
        self,
        setup: Union[
            List[Union[int, bool, str]], Dict[str, List[Union[int, bool, str]]]
        ],
        adjustments: Optional[List[MatrixAdjustment]] = None,
    ) -> None:
        super().__init__(adjustments, setup)


# Retry
class AutomaticRetry(_automatic_retry):
    def __init__(
        self,
        exit_status: Optional[Union[int, List[int], ExitStatusEnum]] = None,
        limit: Optional[int] = None,
        signal: Optional[str] = None,
        signal_reason: Optional[SignalReasonEnum] = None,
    ) -> None:
        super().__init__(exit_status, limit, signal, signal_reason)


class ManualRetry(_manual_retry):
    def __init__(
        self,
        allowed: Optional[bool] = None,
        permit_on_passed: Optional[bool] = None,
        reason: Optional[str] = None,
    ) -> None:
        super().__init__(allowed, permit_on_passed, reason)


class Retry(_retry):
    def __init__(
        self,
        automatic: Optional[Union[bool, AutomaticRetry, List[AutomaticRetry]]],
        manual: Optional[Union[bool, ManualRetry]] = None,
    ) -> None:
        super().__init__(automatic, manual)


# Signature
class Signature(_signature):
    def __init__(
        self,
        algorithm: Optional[str],
        signed_fields: Optional[List[str]],
        value: Optional[str],
    ) -> None:
        super().__init__(algorithm, signed_fields, value)


# Build
class Build(_build):
    def __init__(
        self,
        branch: Optional[str],
        commit: Optional[str],
        env: Optional[Dict[str, Any]],
        message: Optional[str],
        meta_data: Optional[Dict[str, Any]],
    ) -> None:
        super().__init__(branch, commit, env, message, meta_data)
