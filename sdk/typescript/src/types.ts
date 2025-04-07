/**
 * The state that the build is set to when the build is blocked by this block step
 */
export enum BlockedState {
    Failed = "failed",
    Passed = "passed",
    Running = "running",
}

/**
 * Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
 * this attribute, you must also define concurrency_group and concurrency.
 */
export enum ConcurrencyMethod {
    Eager = "eager",
    Ordered = "ordered",
}

export enum NotifyEnum {
    GithubCheck = "github_check",
    GithubCommitStatus = "github_commit_status",
}

export interface DependsOn {
    allow_failure?: boolean;
    step?: string;
}

/**
 * A list of input fields required to be filled out before unblocking the step
 */
export interface Field {
    /**
     * The value that is pre-filled in the text field
     *
     * The value of the option(s) that will be pre-selected in the dropdown
     */
    default?: string[] | string;
    /**
     * The format must be a regular expression implicitly anchored to the beginning and end of
     * the input and is functionally equivalent to the HTML5 pattern attribute.
     */
    format?: string;
    /**
     * The explanatory text that is shown after the label
     */
    hint?: string;
    /**
     * The meta-data key that stores the field's input
     */
    key: string;
    /**
     * Whether the field is required for form submission
     */
    required?: boolean;
    /**
     * The text input name
     */
    text?: string;
    /**
     * Whether more than one option may be selected
     */
    multiple?: boolean;
    options?: Option[];
    /**
     * The text input name
     */
    select?: string;
}

export interface Option {
    /**
     * The text displayed directly under the select field’s label
     */
    hint?: string;
    /**
     * The text displayed on the select list item
     */
    label: string;
    /**
     * Whether the field is required for form submission
     */
    required?: boolean;
    /**
     * The value to be stored as meta-data
     */
    value: string;
}

export interface CacheObject {
    name?: string;
    paths: string[];
    size?: string;
}

export interface SoftFail {
    /**
     * The exit status number that will cause this job to soft-fail
     */
    exit_status?: '*' | number;
}

/**
 * An adjustment to a Build Matrix
 */
export interface Adjustment {
    skip?: boolean | string;
    soft_fail?: SoftFail[] | boolean;
    with: Array<boolean | number | string> | { [key: string]: string };
}

/**
 * Configuration for multi-dimension Build Matrix
 */
export interface MatrixObject {
    /**
     * List of Build Matrix adjustments
     */
    adjustments?: Adjustment[];
    setup:
        | Array<boolean | number | string>
        | { [key: string]: Array<boolean | number | string> };
}

export interface GithubCommitStatus {
    /**
     * GitHub commit status name
     */
    context?: string;
}

export interface NotifySlack {
    channels?: string[];
    message?: string;
}

export interface Notify {
    basecamp_campfire?: string;
    if?: string;
    slack?: NotifySlack | string;
    github_commit_status?: GithubCommitStatus;
    github_check?: { [key: string]: any };
}

export interface PipelineNotify {
    email?: string;
    if?: string;
    basecamp_campfire?: string;
    slack?: NotifySlack | string;
    webhook?: string;
    pagerduty_change_event?: string;
    github_commit_status?: GithubCommitStatus;
    github_check?: { [key: string]: any };
}

/**
 * The conditions for retrying this step.
 */
export interface Retry {
    /**
     * Whether to allow a job to retry automatically. If set to true, the retry conditions are
     * set to the default value.
     */
    automatic?:
        | boolean
        | AutomaticRetry
        | AutomaticRetry[];
    /**
     * Whether to allow a job to be retried manually
     */
    manual?: boolean | ManualRetry;
}

export interface AutomaticRetry {
    /**
     * The exit status number that will cause this job to retry
     */
    exit_status?: number[] | '*' | number;
    /**
     * The number of times this job can be retried
     */
    limit?: number;
    /**
     * The exit signal, if any, that may be retried
     */
    signal?: string;
    /**
     * The exit signal reason, if any, that may be retried
     */
    signal_reason?: SignalReason;
}

/**
 * The exit signal reason, if any, that may be retried
 */
export enum SignalReason {
    AgentRefused = "agent_refused",
    AgentStop = "agent_stop",
    Cancel = "cancel",
    Empty = "*",
    None = "none",
    ProcessRunError = "process_run_error",
    SignatureRejected = "signature_rejected",
}

export interface ManualRetry {
    /**
     * Whether or not this job can be retried manually
     */
    allowed?: boolean;
    /**
     * Whether or not this job can be retried after it has passed
     */
    permit_on_passed?: boolean;
    /**
     * A string that will be displayed in a tooltip on the Retry button in Buildkite. This will
     * only be displayed if the allowed attribute is set to false.
     */
    reason?: string;
}

/**
 * The signature of the command step, generally injected by agents at pipeline upload
 */
export interface Signature {
    /**
     * The algorithm used to generate the signature
     */
    algorithm?: string;
    /**
     * The fields that were signed to form the signature value
     */
    signed_fields?: string[];
    /**
     * The signature value, a JWS compact signature with a detached body
     */
    value?: string;
}

/**
 * Properties of the build that will be created when the step is triggered
 */
export interface Build {
    /**
     * The branch for the build
     */
    branch?: string;
    /**
     * The commit hash for the build
     */
    commit?: string;
    env?: { [key: string]: any };
    /**
     * The message for the build (supports emoji)
     */
    message?: string;
    /**
     * Meta-data for the build
     */
    meta_data?: { [key: string]: any };
}
