// To parse this data:
//
//   import { Convert, Schema } from "./file";
//
//   const schema = Convert.toSchema(json);
//
// These functions will throw an error if the JSON doesn't
// match the expected interface, even if the JSON is valid.

export interface Schema {
    agents?: string[] | { [key: string]: any };
    env?: { [key: string]: any };
    notify?: Array<PurpleBuildNotify | NotifyEnum>;
    /**
     * A list of steps
     */
    steps: Array<GroupStepClass | StringStep>;
    [property: string]: any;
}

export interface PurpleBuildNotify {
    email?: string;
    if?: string;
    basecamp_campfire?: string;
    slack?: PurpleSlack | string;
    webhook?: string;
    pagerduty_change_event?: string;
    github_commit_status?: PurpleGithubCommitStatus;
    github_check?: PurpleGithubCheck;
}

export interface PurpleGithubCheck {
    /**
     * GitHub commit status name
     */
    context?: string;
    [property: string]: any;
}

export interface PurpleGithubCommitStatus {
    /**
     * GitHub commit status name
     */
    context?: string;
}

export interface PurpleSlack {
    channels?: string[];
    message?: string;
    [property: string]: any;
}

export enum NotifyEnum {
    GithubCheck = "github_check",
    GithubCommitStatus = "github_commit_status",
}

/**
 * Waits for previous steps to pass before continuing
 */
export interface GroupStepClass {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * The label of the block step
     */
    block?: BlockStep | string;
    /**
     * The state that the build is set to when the build is blocked by this block step
     */
    blocked_state?: BlockedState;
    branches?: string[] | string;
    depends_on?: Array<DependsOnClass | string> | null | string;
    fields?: Field[];
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: null | string;
    name?: null | string;
    prompt?: string;
    type?: BlockStepType;
    /**
     * The label of the input step
     */
    input?: InputStep | string;
    agents?: string[] | { [key: string]: any };
    /**
     * The glob path/s of artifacts to upload once this step has finished running
     */
    artifact_paths?: string[] | string;
    cache?: string[] | CacheObject | string;
    cancel_on_build_failing?: boolean | AllowDependencyFailureEnum;
    /**
     * The commands to run on the agent
     */
    command?: string[] | CommandStep | string;
    /**
     * The commands to run on the agent
     */
    commands?: string[] | CommandStep | string;
    /**
     * The maximum number of jobs created from this step that are allowed to run at the same
     * time. If you use this attribute, you must also define concurrency_group.
     */
    concurrency?: number;
    /**
     * A unique name for the concurrency group that you are creating with the concurrency
     * attribute
     */
    concurrency_group?: string;
    /**
     * Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
     * this attribute, you must also define concurrency_group and concurrency.
     */
    concurrency_method?: ConcurrencyMethod;
    env?: { [key: string]: any };
    matrix?: Array<boolean | number | string> | MatrixObject;
    /**
     * Array of notification options for this step
     */
    notify?: Array<FluffyBuildNotify | NotifyEnum>;
    /**
     * The number of parallel jobs that will be created based on this step
     */
    parallelism?: number;
    plugins?: Array<{ [key: string]: any } | string> | { [key: string]: any };
    /**
     * Priority of the job, higher priorities are assigned to agents
     */
    priority?: number;
    /**
     * The conditions for retrying this step.
     */
    retry?: Retry;
    /**
     * The signature of the command step, generally injected by agents at pipeline upload
     */
    signature?: Signature;
    skip?: boolean | string;
    soft_fail?: SoftFailElement[] | boolean | AllowDependencyFailureEnum;
    /**
     * The number of minutes to time out a job
     */
    timeout_in_minutes?: number;
    script?: CommandStep;
    /**
     * Continue to the next steps, even if the previous group of steps fail
     */
    continue_on_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * Waits for previous steps to pass before continuing
     */
    wait?: WaitStep | null | string;
    waiter?: WaitStep | null | string;
    /**
     * Whether to continue the build without waiting for the triggered step to complete
     */
    async?: boolean | AllowDependencyFailureEnum;
    /**
     * Properties of the build that will be created when the step is triggered
     */
    build?: Build;
    /**
     * The slug of the pipeline to create a build
     */
    trigger?: TriggerStep | string;
    /**
     * The name to give to this group of steps
     */
    group?: null | string;
    /**
     * A list of steps
     */
    steps?: Array<PurpleStep | StringStep>;
}

export enum AllowDependencyFailureEnum {
    False = "false",
    True = "true",
}

export interface BlockStep {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * The label of the block step
     */
    block?: string;
    /**
     * The state that the build is set to when the build is blocked by this block step
     */
    blocked_state?: BlockedState;
    branches?: string[] | string;
    depends_on?: Array<DependsOnClass | string> | null | string;
    fields?: Field[];
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    prompt?: string;
    type?: BlockType;
}

/**
 * The state that the build is set to when the build is blocked by this block step
 */
export enum BlockedState {
    Failed = "failed",
    Passed = "passed",
    Running = "running",
}

export interface DependsOnClass {
    allow_failure?: boolean | AllowDependencyFailureEnum;
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
    required?: boolean | AllowDependencyFailureEnum;
    /**
     * The text input name
     */
    text?: string;
    /**
     * Whether more than one option may be selected
     */
    multiple?: boolean | AllowDependencyFailureEnum;
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
    required?: boolean | AllowDependencyFailureEnum;
    /**
     * The value to be stored as meta-data
     */
    value: string;
}

export enum BlockType {
    Block = "block",
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

export interface CacheObject {
    name?: string;
    paths: string[];
    size?: string;
    [property: string]: any;
}

export interface CommandStep {
    agents?: string[] | { [key: string]: any };
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * The glob path/s of artifacts to upload once this step has finished running
     */
    artifact_paths?: string[] | string;
    branches?: string[] | string;
    cache?: string[] | CacheObject | string;
    cancel_on_build_failing?: boolean | AllowDependencyFailureEnum;
    /**
     * The commands to run on the agent
     */
    command?: string[] | string;
    /**
     * The commands to run on the agent
     */
    commands?: string[] | string;
    /**
     * The maximum number of jobs created from this step that are allowed to run at the same
     * time. If you use this attribute, you must also define concurrency_group.
     */
    concurrency?: number;
    /**
     * A unique name for the concurrency group that you are creating with the concurrency
     * attribute
     */
    concurrency_group?: string;
    /**
     * Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
     * this attribute, you must also define concurrency_group and concurrency.
     */
    concurrency_method?: ConcurrencyMethod;
    depends_on?: Array<DependsOnClass | string> | null | string;
    env?: { [key: string]: any };
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    matrix?: Array<boolean | number | string> | MatrixObject;
    name?: string;
    /**
     * Array of notification options for this step
     */
    notify?: Array<NotifyClass | NotifyEnum>;
    /**
     * The number of parallel jobs that will be created based on this step
     */
    parallelism?: number;
    plugins?: Array<{ [key: string]: any } | string> | { [key: string]: any };
    /**
     * Priority of the job, higher priorities are assigned to agents
     */
    priority?: number;
    /**
     * The conditions for retrying this step.
     */
    retry?: Retry;
    /**
     * The signature of the command step, generally injected by agents at pipeline upload
     */
    signature?: Signature;
    skip?: boolean | string;
    soft_fail?: SoftFailElement[] | boolean | AllowDependencyFailureEnum;
    /**
     * The number of minutes to time out a job
     */
    timeout_in_minutes?: number;
    type?: ScriptType;
}

/**
 * Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
 * this attribute, you must also define concurrency_group and concurrency.
 */
export enum ConcurrencyMethod {
    Eager = "eager",
    Ordered = "ordered",
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
    [property: string]: any;
}

/**
 * An adjustment to a Build Matrix
 */
export interface Adjustment {
    skip?: boolean | string;
    soft_fail?: SoftFailElement[] | boolean | AllowDependencyFailureEnum;
    with: Array<boolean | number | string> | { [key: string]: string };
    [property: string]: any;
}

export interface SoftFailElement {
    /**
     * The exit status number that will cause this job to soft-fail
     */
    exit_status?: ExitStatusEnum | number;
    [property: string]: any;
}

export enum ExitStatusEnum {
    Empty = "*",
}

export interface NotifyClass {
    basecamp_campfire?: string;
    if?: string;
    slack?: SlackClass | string;
    github_commit_status?: FluffyGithubCommitStatus;
    github_check?: FluffyGithubCheck;
}

export interface FluffyGithubCheck {
    /**
     * GitHub commit status name
     */
    context?: string;
    [property: string]: any;
}

export interface FluffyGithubCommitStatus {
    /**
     * GitHub commit status name
     */
    context?: string;
}

export interface SlackClass {
    channels?: string[];
    message?: string;
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
        | AutomaticRetry[]
        | boolean
        | AutomaticRetry
        | AllowDependencyFailureEnum;
    /**
     * Whether to allow a job to be retried manually
     */
    manual?: boolean | ManualClass | AllowDependencyFailureEnum;
}

export interface AutomaticRetry {
    /**
     * The exit status number that will cause this job to retry
     */
    exit_status?: number[] | ExitStatusEnum | number;
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

export interface ManualClass {
    /**
     * Whether or not this job can be retried manually
     */
    allowed?: boolean | AllowDependencyFailureEnum;
    /**
     * Whether or not this job can be retried after it has passed
     */
    permit_on_passed?: boolean | AllowDependencyFailureEnum;
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
    [property: string]: any;
}

export enum ScriptType {
    Command = "command",
    Commands = "commands",
    Script = "script",
}

export interface InputStep {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    branches?: string[] | string;
    depends_on?: Array<DependsOnClass | string> | null | string;
    fields?: Field[];
    id?: string;
    identifier?: string;
    if?: string;
    /**
     * The label of the input step
     */
    input?: string;
    key?: string;
    label?: string;
    name?: string;
    prompt?: string;
    type?: InputType;
}

export enum InputType {
    Input = "input",
}

export interface FluffyBuildNotify {
    basecamp_campfire?: string;
    if?: string;
    slack?: FluffySlack | string;
    github_commit_status?: TentacledGithubCommitStatus;
    github_check?: TentacledGithubCheck;
    email?: string;
    webhook?: string;
    pagerduty_change_event?: string;
}

export interface TentacledGithubCheck {
    /**
     * GitHub commit status name
     */
    context?: string;
    [property: string]: any;
}

export interface TentacledGithubCommitStatus {
    /**
     * GitHub commit status name
     */
    context?: string;
}

export interface FluffySlack {
    channels?: string[];
    message?: string;
    [property: string]: any;
}

/**
 * Waits for previous steps to pass before continuing
 */
export interface PurpleStep {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * The label of the block step
     */
    block?: BlockStep | string;
    /**
     * The state that the build is set to when the build is blocked by this block step
     */
    blocked_state?: BlockedState;
    branches?: string[] | string;
    depends_on?: Array<DependsOnClass | string> | null | string;
    fields?: Field[];
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: null | string;
    name?: null | string;
    prompt?: string;
    type?: BlockStepType;
    /**
     * The label of the input step
     */
    input?: InputStep | string;
    agents?: string[] | { [key: string]: any };
    /**
     * The glob path/s of artifacts to upload once this step has finished running
     */
    artifact_paths?: string[] | string;
    cache?: string[] | CacheObject | string;
    cancel_on_build_failing?: boolean | AllowDependencyFailureEnum;
    /**
     * The commands to run on the agent
     */
    command?: string[] | CommandStep | string;
    /**
     * The commands to run on the agent
     */
    commands?: string[] | CommandStep | string;
    /**
     * The maximum number of jobs created from this step that are allowed to run at the same
     * time. If you use this attribute, you must also define concurrency_group.
     */
    concurrency?: number;
    /**
     * A unique name for the concurrency group that you are creating with the concurrency
     * attribute
     */
    concurrency_group?: string;
    /**
     * Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
     * this attribute, you must also define concurrency_group and concurrency.
     */
    concurrency_method?: ConcurrencyMethod;
    env?: { [key: string]: any };
    matrix?: Array<boolean | number | string> | MatrixObject;
    /**
     * Array of notification options for this step
     */
    notify?: Array<NotifyClass | NotifyEnum>;
    /**
     * The number of parallel jobs that will be created based on this step
     */
    parallelism?: number;
    plugins?: Array<{ [key: string]: any } | string> | { [key: string]: any };
    /**
     * Priority of the job, higher priorities are assigned to agents
     */
    priority?: number;
    /**
     * The conditions for retrying this step.
     */
    retry?: Retry;
    /**
     * The signature of the command step, generally injected by agents at pipeline upload
     */
    signature?: Signature;
    skip?: boolean | string;
    soft_fail?: SoftFailElement[] | boolean | AllowDependencyFailureEnum;
    /**
     * The number of minutes to time out a job
     */
    timeout_in_minutes?: number;
    script?: CommandStep;
    /**
     * Continue to the next steps, even if the previous group of steps fail
     */
    continue_on_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * Waits for previous steps to pass before continuing
     */
    wait?: WaitStep | null | string;
    waiter?: WaitStep | null | string;
    /**
     * Whether to continue the build without waiting for the triggered step to complete
     */
    async?: boolean | AllowDependencyFailureEnum;
    /**
     * Properties of the build that will be created when the step is triggered
     */
    build?: Build;
    /**
     * The slug of the pipeline to create a build
     */
    trigger?: TriggerStep | string;
}

export interface TriggerStep {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    /**
     * Whether to continue the build without waiting for the triggered step to complete
     */
    async?: boolean | AllowDependencyFailureEnum;
    branches?: string[] | string;
    /**
     * Properties of the build that will be created when the step is triggered
     */
    build?: Build;
    depends_on?: Array<DependsOnClass | string> | null | string;
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    skip?: boolean | string;
    soft_fail?: SoftFailElement[] | boolean | AllowDependencyFailureEnum;
    /**
     * The slug of the pipeline to create a build
     */
    trigger: string;
    type?: TriggerType;
}

export enum TriggerType {
    Trigger = "trigger",
}

export enum BlockStepType {
    Block = "block",
    Command = "command",
    Commands = "commands",
    Input = "input",
    Script = "script",
    Trigger = "trigger",
    Wait = "wait",
    Waiter = "waiter",
}

/**
 * Waits for previous steps to pass before continuing
 */
export interface WaitStep {
    allow_dependency_failure?: boolean | AllowDependencyFailureEnum;
    branches?: string[] | string;
    /**
     * Continue to the next steps, even if the previous group of steps fail
     */
    continue_on_failure?: boolean | AllowDependencyFailureEnum;
    depends_on?: Array<DependsOnClass | string> | null | string;
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: null | string;
    name?: null | string;
    type?: WaitType;
    /**
     * Waits for previous steps to pass before continuing
     */
    wait?: null | string;
    waiter?: null | string;
}

export enum WaitType {
    Wait = "wait",
    Waiter = "waiter",
}

/**
 * Pauses the execution of a build and waits on a user to unblock it
 *
 * Waits for previous steps to pass before continuing
 */
export enum StringStep {
    Block = "block",
    Input = "input",
    Wait = "wait",
    Waiter = "waiter",
}

// Converts JSON strings to/from your types
// and asserts the results of JSON.parse at runtime
export class Convert {
    public static toSchema(json: string): Schema {
        return cast(JSON.parse(json), r("Schema"));
    }

    public static schemaToJson(value: Schema): string {
        return JSON.stringify(uncast(value, r("Schema")), null, 2);
    }
}

function invalidValue(typ: any, val: any, key: any, parent: any = ""): never {
    const prettyTyp = prettyTypeName(typ);
    const parentText = parent ? ` on ${parent}` : "";
    const keyText = key ? ` for key "${key}"` : "";
    throw Error(
        `Invalid value${keyText}${parentText}. Expected ${prettyTyp} but got ${JSON.stringify(
            val
        )}`
    );
}

function prettyTypeName(typ: any): string {
    if (Array.isArray(typ)) {
        if (typ.length === 2 && typ[0] === undefined) {
            return `an optional ${prettyTypeName(typ[1])}`;
        } else {
            return `one of [${typ
                .map((a) => {
                    return prettyTypeName(a);
                })
                .join(", ")}]`;
        }
    } else if (typeof typ === "object" && typ.literal !== undefined) {
        return typ.literal;
    } else {
        return typeof typ;
    }
}

function jsonToJSProps(typ: any): any {
    if (typ.jsonToJS === undefined) {
        const map: any = {};
        typ.props.forEach(
            (p: any) => (map[p.json] = { key: p.js, typ: p.typ })
        );
        typ.jsonToJS = map;
    }
    return typ.jsonToJS;
}

function jsToJSONProps(typ: any): any {
    if (typ.jsToJSON === undefined) {
        const map: any = {};
        typ.props.forEach(
            (p: any) => (map[p.js] = { key: p.json, typ: p.typ })
        );
        typ.jsToJSON = map;
    }
    return typ.jsToJSON;
}

function transform(
    val: any,
    typ: any,
    getProps: any,
    key: any = "",
    parent: any = ""
): any {
    function transformPrimitive(typ: string, val: any): any {
        if (typeof typ === typeof val) return val;
        return invalidValue(typ, val, key, parent);
    }

    function transformUnion(typs: any[], val: any): any {
        // val must validate against one typ in typs
        const l = typs.length;
        for (let i = 0; i < l; i++) {
            const typ = typs[i];
            try {
                return transform(val, typ, getProps);
            } catch (_) {}
        }
        return invalidValue(typs, val, key, parent);
    }

    function transformEnum(cases: string[], val: any): any {
        if (cases.indexOf(val) !== -1) return val;
        return invalidValue(
            cases.map((a) => {
                return l(a);
            }),
            val,
            key,
            parent
        );
    }

    function transformArray(typ: any, val: any): any {
        // val must be an array with no invalid elements
        if (!Array.isArray(val))
            return invalidValue(l("array"), val, key, parent);
        return val.map((el) => transform(el, typ, getProps));
    }

    function transformDate(val: any): any {
        if (val === null) {
            return null;
        }
        const d = new Date(val);
        if (isNaN(d.valueOf())) {
            return invalidValue(l("Date"), val, key, parent);
        }
        return d;
    }

    function transformObject(
        props: { [k: string]: any },
        additional: any,
        val: any
    ): any {
        if (val === null || typeof val !== "object" || Array.isArray(val)) {
            return invalidValue(l(ref || "object"), val, key, parent);
        }
        const result: any = {};
        Object.getOwnPropertyNames(props).forEach((key) => {
            const prop = props[key];
            const v = Object.prototype.hasOwnProperty.call(val, key)
                ? val[key]
                : undefined;
            result[prop.key] = transform(v, prop.typ, getProps, key, ref);
        });
        Object.getOwnPropertyNames(val).forEach((key) => {
            if (!Object.prototype.hasOwnProperty.call(props, key)) {
                result[key] = transform(
                    val[key],
                    additional,
                    getProps,
                    key,
                    ref
                );
            }
        });
        return result;
    }

    if (typ === "any") return val;
    if (typ === null) {
        if (val === null) return val;
        return invalidValue(typ, val, key, parent);
    }
    if (typ === false) return invalidValue(typ, val, key, parent);
    let ref: any = undefined;
    while (typeof typ === "object" && typ.ref !== undefined) {
        ref = typ.ref;
        typ = typeMap[typ.ref];
    }
    if (Array.isArray(typ)) return transformEnum(typ, val);
    if (typeof typ === "object") {
        return typ.hasOwnProperty("unionMembers")
            ? transformUnion(typ.unionMembers, val)
            : typ.hasOwnProperty("arrayItems")
            ? transformArray(typ.arrayItems, val)
            : typ.hasOwnProperty("props")
            ? transformObject(getProps(typ), typ.additional, val)
            : invalidValue(typ, val, key, parent);
    }
    // Numbers can be parsed by Date but shouldn't be.
    if (typ === Date && typeof val !== "number") return transformDate(val);
    return transformPrimitive(typ, val);
}

function cast<T>(val: any, typ: any): T {
    return transform(val, typ, jsonToJSProps);
}

function uncast<T>(val: T, typ: any): any {
    return transform(val, typ, jsToJSONProps);
}

function l(typ: any) {
    return { literal: typ };
}

function a(typ: any) {
    return { arrayItems: typ };
}

function u(...typs: any[]) {
    return { unionMembers: typs };
}

function o(props: any[], additional: any) {
    return { props, additional };
}

function m(additional: any) {
    return { props: [], additional };
}

function r(name: string) {
    return { ref: name };
}

const typeMap: any = {
    Schema: o(
        [
            {
                json: "agents",
                js: "agents",
                typ: u(undefined, u(a(""), m("any"))),
            },
            { json: "env", js: "env", typ: u(undefined, m("any")) },
            {
                json: "notify",
                js: "notify",
                typ: u(
                    undefined,
                    a(u(r("PurpleBuildNotify"), r("NotifyEnum")))
                ),
            },
            {
                json: "steps",
                js: "steps",
                typ: a(u(r("GroupStepClass"), r("StringStep"))),
            },
        ],
        "any"
    ),
    PurpleBuildNotify: o(
        [
            { json: "email", js: "email", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            {
                json: "basecamp_campfire",
                js: "basecamp_campfire",
                typ: u(undefined, ""),
            },
            {
                json: "slack",
                js: "slack",
                typ: u(undefined, u(r("PurpleSlack"), "")),
            },
            { json: "webhook", js: "webhook", typ: u(undefined, "") },
            {
                json: "pagerduty_change_event",
                js: "pagerduty_change_event",
                typ: u(undefined, ""),
            },
            {
                json: "github_commit_status",
                js: "github_commit_status",
                typ: u(undefined, r("PurpleGithubCommitStatus")),
            },
            {
                json: "github_check",
                js: "github_check",
                typ: u(undefined, r("PurpleGithubCheck")),
            },
        ],
        false
    ),
    PurpleGithubCheck: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        "any"
    ),
    PurpleGithubCommitStatus: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        false
    ),
    PurpleSlack: o(
        [
            { json: "channels", js: "channels", typ: u(undefined, a("")) },
            { json: "message", js: "message", typ: u(undefined, "") },
        ],
        "any"
    ),
    GroupStepClass: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "block",
                js: "block",
                typ: u(undefined, u(r("BlockStep"), "")),
            },
            {
                json: "blocked_state",
                js: "blocked_state",
                typ: u(undefined, r("BlockedState")),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "fields", js: "fields", typ: u(undefined, a(r("Field"))) },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, u(null, "")) },
            { json: "name", js: "name", typ: u(undefined, u(null, "")) },
            { json: "prompt", js: "prompt", typ: u(undefined, "") },
            { json: "type", js: "type", typ: u(undefined, r("BlockStepType")) },
            {
                json: "input",
                js: "input",
                typ: u(undefined, u(r("InputStep"), "")),
            },
            {
                json: "agents",
                js: "agents",
                typ: u(undefined, u(a(""), m("any"))),
            },
            {
                json: "artifact_paths",
                js: "artifact_paths",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "cache",
                js: "cache",
                typ: u(undefined, u(a(""), r("CacheObject"), "")),
            },
            {
                json: "cancel_on_build_failing",
                js: "cancel_on_build_failing",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "command",
                js: "command",
                typ: u(undefined, u(a(""), r("CommandStep"), "")),
            },
            {
                json: "commands",
                js: "commands",
                typ: u(undefined, u(a(""), r("CommandStep"), "")),
            },
            { json: "concurrency", js: "concurrency", typ: u(undefined, 0) },
            {
                json: "concurrency_group",
                js: "concurrency_group",
                typ: u(undefined, ""),
            },
            {
                json: "concurrency_method",
                js: "concurrency_method",
                typ: u(undefined, r("ConcurrencyMethod")),
            },
            { json: "env", js: "env", typ: u(undefined, m("any")) },
            {
                json: "matrix",
                js: "matrix",
                typ: u(undefined, u(a(u(true, 0, "")), r("MatrixObject"))),
            },
            {
                json: "notify",
                js: "notify",
                typ: u(
                    undefined,
                    a(u(r("FluffyBuildNotify"), r("NotifyEnum")))
                ),
            },
            { json: "parallelism", js: "parallelism", typ: u(undefined, 0) },
            {
                json: "plugins",
                js: "plugins",
                typ: u(undefined, u(a(u(m("any"), "")), m("any"))),
            },
            { json: "priority", js: "priority", typ: u(undefined, 0) },
            { json: "retry", js: "retry", typ: u(undefined, r("Retry")) },
            {
                json: "signature",
                js: "signature",
                typ: u(undefined, r("Signature")),
            },
            { json: "skip", js: "skip", typ: u(undefined, u(true, "")) },
            {
                json: "soft_fail",
                js: "soft_fail",
                typ: u(
                    undefined,
                    u(
                        a(r("SoftFailElement")),
                        true,
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            {
                json: "timeout_in_minutes",
                js: "timeout_in_minutes",
                typ: u(undefined, 0),
            },
            {
                json: "script",
                js: "script",
                typ: u(undefined, r("CommandStep")),
            },
            {
                json: "continue_on_failure",
                js: "continue_on_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "wait",
                js: "wait",
                typ: u(undefined, u(r("WaitStep"), null, "")),
            },
            {
                json: "waiter",
                js: "waiter",
                typ: u(undefined, u(r("WaitStep"), null, "")),
            },
            {
                json: "async",
                js: "async",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "build", js: "build", typ: u(undefined, r("Build")) },
            {
                json: "trigger",
                js: "trigger",
                typ: u(undefined, u(r("TriggerStep"), "")),
            },
            { json: "group", js: "group", typ: u(undefined, u(null, "")) },
            {
                json: "steps",
                js: "steps",
                typ: u(undefined, a(u(r("PurpleStep"), r("StringStep")))),
            },
        ],
        false
    ),
    BlockStep: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "block", js: "block", typ: u(undefined, "") },
            {
                json: "blocked_state",
                js: "blocked_state",
                typ: u(undefined, r("BlockedState")),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "fields", js: "fields", typ: u(undefined, a(r("Field"))) },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, "") },
            { json: "name", js: "name", typ: u(undefined, "") },
            { json: "prompt", js: "prompt", typ: u(undefined, "") },
            { json: "type", js: "type", typ: u(undefined, r("BlockType")) },
        ],
        false
    ),
    DependsOnClass: o(
        [
            {
                json: "allow_failure",
                js: "allow_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "step", js: "step", typ: u(undefined, "") },
        ],
        false
    ),
    Field: o(
        [
            { json: "default", js: "default", typ: u(undefined, u(a(""), "")) },
            { json: "format", js: "format", typ: u(undefined, "") },
            { json: "hint", js: "hint", typ: u(undefined, "") },
            { json: "key", js: "key", typ: "" },
            {
                json: "required",
                js: "required",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "text", js: "text", typ: u(undefined, "") },
            {
                json: "multiple",
                js: "multiple",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "options",
                js: "options",
                typ: u(undefined, a(r("Option"))),
            },
            { json: "select", js: "select", typ: u(undefined, "") },
        ],
        false
    ),
    Option: o(
        [
            { json: "hint", js: "hint", typ: u(undefined, "") },
            { json: "label", js: "label", typ: "" },
            {
                json: "required",
                js: "required",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "value", js: "value", typ: "" },
        ],
        false
    ),
    Build: o(
        [
            { json: "branch", js: "branch", typ: u(undefined, "") },
            { json: "commit", js: "commit", typ: u(undefined, "") },
            { json: "env", js: "env", typ: u(undefined, m("any")) },
            { json: "message", js: "message", typ: u(undefined, "") },
            { json: "meta_data", js: "meta_data", typ: u(undefined, m("any")) },
        ],
        false
    ),
    CacheObject: o(
        [
            { json: "name", js: "name", typ: u(undefined, "") },
            { json: "paths", js: "paths", typ: a("") },
            { json: "size", js: "size", typ: u(undefined, "") },
        ],
        "any"
    ),
    CommandStep: o(
        [
            {
                json: "agents",
                js: "agents",
                typ: u(undefined, u(a(""), m("any"))),
            },
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "artifact_paths",
                js: "artifact_paths",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "cache",
                js: "cache",
                typ: u(undefined, u(a(""), r("CacheObject"), "")),
            },
            {
                json: "cancel_on_build_failing",
                js: "cancel_on_build_failing",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "command", js: "command", typ: u(undefined, u(a(""), "")) },
            {
                json: "commands",
                js: "commands",
                typ: u(undefined, u(a(""), "")),
            },
            { json: "concurrency", js: "concurrency", typ: u(undefined, 0) },
            {
                json: "concurrency_group",
                js: "concurrency_group",
                typ: u(undefined, ""),
            },
            {
                json: "concurrency_method",
                js: "concurrency_method",
                typ: u(undefined, r("ConcurrencyMethod")),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "env", js: "env", typ: u(undefined, m("any")) },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, "") },
            {
                json: "matrix",
                js: "matrix",
                typ: u(undefined, u(a(u(true, 0, "")), r("MatrixObject"))),
            },
            { json: "name", js: "name", typ: u(undefined, "") },
            {
                json: "notify",
                js: "notify",
                typ: u(undefined, a(u(r("NotifyClass"), r("NotifyEnum")))),
            },
            { json: "parallelism", js: "parallelism", typ: u(undefined, 0) },
            {
                json: "plugins",
                js: "plugins",
                typ: u(undefined, u(a(u(m("any"), "")), m("any"))),
            },
            { json: "priority", js: "priority", typ: u(undefined, 0) },
            { json: "retry", js: "retry", typ: u(undefined, r("Retry")) },
            {
                json: "signature",
                js: "signature",
                typ: u(undefined, r("Signature")),
            },
            { json: "skip", js: "skip", typ: u(undefined, u(true, "")) },
            {
                json: "soft_fail",
                js: "soft_fail",
                typ: u(
                    undefined,
                    u(
                        a(r("SoftFailElement")),
                        true,
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            {
                json: "timeout_in_minutes",
                js: "timeout_in_minutes",
                typ: u(undefined, 0),
            },
            { json: "type", js: "type", typ: u(undefined, r("ScriptType")) },
        ],
        false
    ),
    MatrixObject: o(
        [
            {
                json: "adjustments",
                js: "adjustments",
                typ: u(undefined, a(r("Adjustment"))),
            },
            {
                json: "setup",
                js: "setup",
                typ: u(a(u(true, 0, "")), m(a(u(true, 0, "")))),
            },
        ],
        "any"
    ),
    Adjustment: o(
        [
            { json: "skip", js: "skip", typ: u(undefined, u(true, "")) },
            {
                json: "soft_fail",
                js: "soft_fail",
                typ: u(
                    undefined,
                    u(
                        a(r("SoftFailElement")),
                        true,
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            { json: "with", js: "with", typ: u(a(u(true, 0, "")), m("")) },
        ],
        "any"
    ),
    SoftFailElement: o(
        [
            {
                json: "exit_status",
                js: "exit_status",
                typ: u(undefined, u(r("ExitStatusEnum"), 0)),
            },
        ],
        "any"
    ),
    NotifyClass: o(
        [
            {
                json: "basecamp_campfire",
                js: "basecamp_campfire",
                typ: u(undefined, ""),
            },
            { json: "if", js: "if", typ: u(undefined, "") },
            {
                json: "slack",
                js: "slack",
                typ: u(undefined, u(r("SlackClass"), "")),
            },
            {
                json: "github_commit_status",
                js: "github_commit_status",
                typ: u(undefined, r("FluffyGithubCommitStatus")),
            },
            {
                json: "github_check",
                js: "github_check",
                typ: u(undefined, r("FluffyGithubCheck")),
            },
        ],
        false
    ),
    FluffyGithubCheck: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        "any"
    ),
    FluffyGithubCommitStatus: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        false
    ),
    SlackClass: o(
        [
            { json: "channels", js: "channels", typ: u(undefined, a("")) },
            { json: "message", js: "message", typ: u(undefined, "") },
        ],
        false
    ),
    Retry: o(
        [
            {
                json: "automatic",
                js: "automatic",
                typ: u(
                    undefined,
                    u(
                        a(r("AutomaticRetry")),
                        true,
                        r("AutomaticRetry"),
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            {
                json: "manual",
                js: "manual",
                typ: u(
                    undefined,
                    u(true, r("ManualClass"), r("AllowDependencyFailureEnum"))
                ),
            },
        ],
        false
    ),
    AutomaticRetry: o(
        [
            {
                json: "exit_status",
                js: "exit_status",
                typ: u(undefined, u(a(0), r("ExitStatusEnum"), 0)),
            },
            { json: "limit", js: "limit", typ: u(undefined, 0) },
            { json: "signal", js: "signal", typ: u(undefined, "") },
            {
                json: "signal_reason",
                js: "signal_reason",
                typ: u(undefined, r("SignalReason")),
            },
        ],
        false
    ),
    ManualClass: o(
        [
            {
                json: "allowed",
                js: "allowed",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "permit_on_passed",
                js: "permit_on_passed",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "reason", js: "reason", typ: u(undefined, "") },
        ],
        false
    ),
    Signature: o(
        [
            { json: "algorithm", js: "algorithm", typ: u(undefined, "") },
            {
                json: "signed_fields",
                js: "signed_fields",
                typ: u(undefined, a("")),
            },
            { json: "value", js: "value", typ: u(undefined, "") },
        ],
        "any"
    ),
    InputStep: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "fields", js: "fields", typ: u(undefined, a(r("Field"))) },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "input", js: "input", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, "") },
            { json: "name", js: "name", typ: u(undefined, "") },
            { json: "prompt", js: "prompt", typ: u(undefined, "") },
            { json: "type", js: "type", typ: u(undefined, r("InputType")) },
        ],
        false
    ),
    FluffyBuildNotify: o(
        [
            {
                json: "basecamp_campfire",
                js: "basecamp_campfire",
                typ: u(undefined, ""),
            },
            { json: "if", js: "if", typ: u(undefined, "") },
            {
                json: "slack",
                js: "slack",
                typ: u(undefined, u(r("FluffySlack"), "")),
            },
            {
                json: "github_commit_status",
                js: "github_commit_status",
                typ: u(undefined, r("TentacledGithubCommitStatus")),
            },
            {
                json: "github_check",
                js: "github_check",
                typ: u(undefined, r("TentacledGithubCheck")),
            },
            { json: "email", js: "email", typ: u(undefined, "") },
            { json: "webhook", js: "webhook", typ: u(undefined, "") },
            {
                json: "pagerduty_change_event",
                js: "pagerduty_change_event",
                typ: u(undefined, ""),
            },
        ],
        false
    ),
    TentacledGithubCheck: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        "any"
    ),
    TentacledGithubCommitStatus: o(
        [{ json: "context", js: "context", typ: u(undefined, "") }],
        false
    ),
    FluffySlack: o(
        [
            { json: "channels", js: "channels", typ: u(undefined, a("")) },
            { json: "message", js: "message", typ: u(undefined, "") },
        ],
        "any"
    ),
    PurpleStep: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "block",
                js: "block",
                typ: u(undefined, u(r("BlockStep"), "")),
            },
            {
                json: "blocked_state",
                js: "blocked_state",
                typ: u(undefined, r("BlockedState")),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "fields", js: "fields", typ: u(undefined, a(r("Field"))) },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, u(null, "")) },
            { json: "name", js: "name", typ: u(undefined, u(null, "")) },
            { json: "prompt", js: "prompt", typ: u(undefined, "") },
            { json: "type", js: "type", typ: u(undefined, r("BlockStepType")) },
            {
                json: "input",
                js: "input",
                typ: u(undefined, u(r("InputStep"), "")),
            },
            {
                json: "agents",
                js: "agents",
                typ: u(undefined, u(a(""), m("any"))),
            },
            {
                json: "artifact_paths",
                js: "artifact_paths",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "cache",
                js: "cache",
                typ: u(undefined, u(a(""), r("CacheObject"), "")),
            },
            {
                json: "cancel_on_build_failing",
                js: "cancel_on_build_failing",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "command",
                js: "command",
                typ: u(undefined, u(a(""), r("CommandStep"), "")),
            },
            {
                json: "commands",
                js: "commands",
                typ: u(undefined, u(a(""), r("CommandStep"), "")),
            },
            { json: "concurrency", js: "concurrency", typ: u(undefined, 0) },
            {
                json: "concurrency_group",
                js: "concurrency_group",
                typ: u(undefined, ""),
            },
            {
                json: "concurrency_method",
                js: "concurrency_method",
                typ: u(undefined, r("ConcurrencyMethod")),
            },
            { json: "env", js: "env", typ: u(undefined, m("any")) },
            {
                json: "matrix",
                js: "matrix",
                typ: u(undefined, u(a(u(true, 0, "")), r("MatrixObject"))),
            },
            {
                json: "notify",
                js: "notify",
                typ: u(undefined, a(u(r("NotifyClass"), r("NotifyEnum")))),
            },
            { json: "parallelism", js: "parallelism", typ: u(undefined, 0) },
            {
                json: "plugins",
                js: "plugins",
                typ: u(undefined, u(a(u(m("any"), "")), m("any"))),
            },
            { json: "priority", js: "priority", typ: u(undefined, 0) },
            { json: "retry", js: "retry", typ: u(undefined, r("Retry")) },
            {
                json: "signature",
                js: "signature",
                typ: u(undefined, r("Signature")),
            },
            { json: "skip", js: "skip", typ: u(undefined, u(true, "")) },
            {
                json: "soft_fail",
                js: "soft_fail",
                typ: u(
                    undefined,
                    u(
                        a(r("SoftFailElement")),
                        true,
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            {
                json: "timeout_in_minutes",
                js: "timeout_in_minutes",
                typ: u(undefined, 0),
            },
            {
                json: "script",
                js: "script",
                typ: u(undefined, r("CommandStep")),
            },
            {
                json: "continue_on_failure",
                js: "continue_on_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "wait",
                js: "wait",
                typ: u(undefined, u(r("WaitStep"), null, "")),
            },
            {
                json: "waiter",
                js: "waiter",
                typ: u(undefined, u(r("WaitStep"), null, "")),
            },
            {
                json: "async",
                js: "async",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            { json: "build", js: "build", typ: u(undefined, r("Build")) },
            {
                json: "trigger",
                js: "trigger",
                typ: u(undefined, u(r("TriggerStep"), "")),
            },
        ],
        false
    ),
    TriggerStep: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "async",
                js: "async",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            { json: "build", js: "build", typ: u(undefined, r("Build")) },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, "") },
            { json: "name", js: "name", typ: u(undefined, "") },
            { json: "skip", js: "skip", typ: u(undefined, u(true, "")) },
            {
                json: "soft_fail",
                js: "soft_fail",
                typ: u(
                    undefined,
                    u(
                        a(r("SoftFailElement")),
                        true,
                        r("AllowDependencyFailureEnum")
                    )
                ),
            },
            { json: "trigger", js: "trigger", typ: "" },
            { json: "type", js: "type", typ: u(undefined, r("TriggerType")) },
        ],
        false
    ),
    WaitStep: o(
        [
            {
                json: "allow_dependency_failure",
                js: "allow_dependency_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "branches",
                js: "branches",
                typ: u(undefined, u(a(""), "")),
            },
            {
                json: "continue_on_failure",
                js: "continue_on_failure",
                typ: u(undefined, u(true, r("AllowDependencyFailureEnum"))),
            },
            {
                json: "depends_on",
                js: "depends_on",
                typ: u(undefined, u(a(u(r("DependsOnClass"), "")), null, "")),
            },
            { json: "id", js: "id", typ: u(undefined, "") },
            { json: "identifier", js: "identifier", typ: u(undefined, "") },
            { json: "if", js: "if", typ: u(undefined, "") },
            { json: "key", js: "key", typ: u(undefined, "") },
            { json: "label", js: "label", typ: u(undefined, u(null, "")) },
            { json: "name", js: "name", typ: u(undefined, u(null, "")) },
            { json: "type", js: "type", typ: u(undefined, r("WaitType")) },
            { json: "wait", js: "wait", typ: u(undefined, u(null, "")) },
            { json: "waiter", js: "waiter", typ: u(undefined, u(null, "")) },
        ],
        false
    ),
    NotifyEnum: ["github_check", "github_commit_status"],
    AllowDependencyFailureEnum: ["false", "true"],
    BlockedState: ["failed", "passed", "running"],
    BlockType: ["block"],
    ConcurrencyMethod: ["eager", "ordered"],
    ExitStatusEnum: ["*"],
    SignalReason: [
        "agent_refused",
        "agent_stop",
        "cancel",
        "*",
        "none",
        "process_run_error",
        "signature_rejected",
    ],
    ScriptType: ["command", "commands", "script"],
    InputType: ["input"],
    TriggerType: ["trigger"],
    BlockStepType: [
        "block",
        "command",
        "commands",
        "input",
        "script",
        "trigger",
        "wait",
        "waiter",
    ],
    WaitType: ["wait", "waiter"],
    StringStep: ["block", "input", "wait", "waiter"],
};
