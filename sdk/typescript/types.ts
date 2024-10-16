// This file is auto-generated. Do not edit.

export interface Build {
    // The message for the build. Supports emoji.
    message?: string;
    // The commit hash for the build.
    commit?: string;
    // The branch for the build.
    branch?: string;
    // A map of meta-data for the build.
    metaData?: Record<string, string>;
    // A map of environment variables for the build.
    env?: Record<string, string>;
}

export enum BlockedState {
    PASSED = "passed",
    FAILED = "failed",
    RUNNING = "running",
}


type Fields = (TextInput | SelectInputAttribute)[]

export interface Retry {
    // Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.
    automatic?: boolean;
    // Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.
    manual?: boolean;
}

export interface SelectInputOption {
    // The text displayed for the option.
    label?: string;
    // The value to be stored as meta-data (to be later retrieved using the buildkite-agent meta-data command).
    value?: string;
}

export interface SelectInputAttribute {
    // The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.
    key?: string;
    // The label for the select input.
    select?: string;
    // The explanatory text that is shown after the label.
    hint?: string;
    // A boolean value that defines whether the field is required for form submission.
    required?: boolean;
    // The value that is pre-filled in the text field.
    default?: string;
    // A boolean value that defines whether multiple options may be selected. When multiple options are selected, they are delimited in the meta-data field by a line break
    multiple?: boolean;
    // The list of select field options. For 6 or less options they'll be displayed as radio buttons, otherwise they'll be displayed in a dropdown box. If selecting multiple options is permitted the options will be displayed as checkboxes.
    options?: SelectInputOption;
}

export interface TextInput {
    // The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.
    key?: string;
    // The label for the text input.
    text?: string;
    // The explanatory text that is shown after the label.
    hint?: string;
    // A boolean value that defines whether the field is required for form submission.
    required?: boolean;
    // The value that is pre-filled in the text field.
    default?: string;
}

type Steps = (Block | Command | Input | Trigger | Wait)[]

export interface Block {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // The label for this block step.
    block?: string;
    // The state that the build is set to when the build is blocked by this block step. The default is passed. When the blocked_state of a block step is set to failed, the step that triggered it will be stuck in the running state until it is manually unblocked. Default: passed Values: passed, failed, running
    blockedState?: BlockedState;
    // The branch pattern defining which branches will include this block step in their builds.
    branches?: string;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // An input step is used to collect information from a user.
    fields?: Fields;
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // A unique string to identify the block step.
    key?: string;
    // The instructional message displayed in the dialog box when the unblock step is activated.
    prompt?: string;
}

export interface Command {
    // A map of agent tag keys to values to target specific agents for this step.
    agents?: Record<string, string>;
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // The glob path or paths of artifacts to upload from this step.
    artifactPaths?: string[];
    // The branch pattern defining which branches will include this block step in their builds.
    branches?: string;
    // Setting this attribute to true cancels the job as soon as the build is marked as failing.
    cancelOnBuildFailing?: boolean;
    // The shell command to run during this step.
    commands?: string[];
    // The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.
    concurrency?: number;
    // A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.
    concurrencyGroup?: string;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // A map of environment variables for this step.
    env?: Record<string, string>;
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // A unique string to identify the block step.
    key?: string;
    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    label?: string;
    // An array of values to be used in the matrix expansion.
    matrix?: string[];
    // The number of parallel jobs that will be created based on this step.
    parallelism?: number;
    // An array of plugins for this step.
    plugins?: Record<string, any>[];
    // Adjust the priority for a specific job, as a positive or negative integer.
    priority?: number;
    // The conditions for retrying this step.
    retry?: Retry;
    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    skip?: boolean;
    // Allow specified non-zero exit statuses not to fail the build.
    softFail?: Record<string, number>;
    // Allow all non-zero exit statuses not to fail the build.
    softFailAll?: boolean;
    // The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.
    timeoutInMinutes?: number;
}

export interface Group {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // Name of the group in the UI. In YAML, if you don't want a label, pass a `~`. Can also be provided in the `label` attribute if `null` is provided to the `group` attribute.
    group?: string;
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // A unique string to identify the block step.
    key?: string;
    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    label?: string;
    // Allows you to trigger build notifications to different services. You can also choose to conditionally send notifications based on pipeline events.
    notify?: string;
    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    skip?: boolean;
    // A list of steps in the group; at least 1 step is required. Allowed step types: wait, trigger, command/commands, block, input.
    steps?: Steps;
}

export interface Input {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // The branch pattern defining which branches will include this block step in their builds.
    branches?: string;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // An input step is used to collect information from a user.
    fields?: Fields;
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // The label for this input step.
    input?: string;
    // A unique string to identify the block step.
    key?: string;
    // The instructional message displayed in the dialog box when the unblock step is activated.
    prompt?: string;
}

export interface Trigger {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // If set to true the step will immediately continue, regardless of the success of the triggered build. If set to false the step will wait for the triggered build to complete and continue only if the triggered build passed.
// Note that when async is set to true, as long as the triggered build starts, the original pipeline will show that as successful. The original pipeline does not get updated after subsequent steps or after the triggered build completes.
    async?: boolean;
    // The branch pattern defining which branches will include this block step in their builds.
    branches?: string;
    // An optional map of attributes for the triggered build. Available attributes: branch, commit, env, message, meta_data
    build?: Build;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    label?: string;
    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    skip?: boolean;
    // Allow specified non-zero exit statuses not to fail the build.
    softFail?: Record<string, number>;
    // The slug of the pipeline to create a build. You can find it in the URL of your pipeline, and it corresponds to the name of the pipeline, converted to kebab-case.
    trigger?: string;
}

export interface Wait {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;
    // Run the next step, even if the previous step has failed.
    continueOnFailure?: boolean;
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;
    // When providing options for the wait step, you will need to set this value to "~".
    wait: string;
}
