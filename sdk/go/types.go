// This file is auto-generated. Do not edit.
package buildkite
import (
    "encoding/json"
)
type Build struct {
    // The message for the build. Supports emoji.
    Message string `json:"message,omitempty"`

    // The commit hash for the build.
    Commit string `json:"commit,omitempty"`

    // The branch for the build.
    Branch string `json:"branch,omitempty"`

    // A map of meta-data for the build.
    MetaData map[string]string `json:"meta_data,omitempty"`

    // A map of environment variables for the build.
    Env map[string]string `json:"env,omitempty"`

}

type BlockedState string
const (
    PASSED BlockedState = "passed"
    FAILED BlockedState = "failed"
    RUNNING BlockedState = "running"
)

type Fields interface {
    MarshalJSON() ([]byte, error)
}
func (x *TextInput) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}
func (x *SelectInputAttribute) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}

type Retry struct {
    // Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.
    Automatic bool `json:"automatic,omitempty"`

    // Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.
    Manual bool `json:"manual,omitempty"`

}

type SelectInputOption struct {
    // The text displayed for the option.
    Label string `json:"label,omitempty"`

    // The value to be stored as meta-data (to be later retrieved using the buildkite-agent meta-data command).
    Value string `json:"value,omitempty"`

}

type SelectInputAttribute struct {
    // The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.
    Key string `json:"key,omitempty"`

    // The label for the select input.
    Select string `json:"select,omitempty"`

    // The explanatory text that is shown after the label.
    Hint string `json:"hint,omitempty"`

    // A boolean value that defines whether the field is required for form submission.
    Required bool `json:"required,omitempty"`

    // The value that is pre-filled in the text field.
    Default string `json:"default,omitempty"`

    // A boolean value that defines whether multiple options may be selected. When multiple options are selected, they are delimited in the meta-data field by a line break
    Multiple bool `json:"multiple,omitempty"`

    // The list of select field options. For 6 or less options they'll be displayed as radio buttons, otherwise they'll be displayed in a dropdown box. If selecting multiple options is permitted the options will be displayed as checkboxes.
    Options []SelectInputOption `json:"options,omitempty"`

}

type TextInput struct {
    // The meta-data key that stores the field's input (using the buildkite-agent meta-data command). The key may only contain alphanumeric characters, slashes, dashes, or underscores.
    Key string `json:"key,omitempty"`

    // The label for the text input.
    Text string `json:"text,omitempty"`

    // The explanatory text that is shown after the label.
    Hint string `json:"hint,omitempty"`

    // A boolean value that defines whether the field is required for form submission.
    Required bool `json:"required,omitempty"`

    // The value that is pre-filled in the text field.
    Default string `json:"default,omitempty"`

}

type Steps interface {
    MarshalJSON() ([]byte, error)
}
func (x *Block) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}
func (x *Command) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}
func (x *Input) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}
func (x *Trigger) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}
func (x *Wait) MarshalJSON() ([]byte, error) {
    return json.Marshal(*x)
}

type WaitLabel string
func (WaitLabel) MarshalJSON() ([]byte,error) {
 return json.Marshal(nil) 
}

type Block struct {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // The label for this block step.
    Block string `json:"block,omitempty"`

    // The state that the build is set to when the build is blocked by this block step. The default is passed. When the blocked_state of a block step is set to failed, the step that triggered it will be stuck in the running state until it is manually unblocked. Default: passed Values: passed, failed, running
    BlockedState *BlockedState `json:"blocked_state,omitempty"`
    // The branch pattern defining which branches will include this block step in their builds.
    Branches string `json:"branches,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // An input step is used to collect information from a user.
    Fields []Fields `json:"fields,omitempty"`
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // A unique string to identify the block step.
    Key string `json:"key,omitempty"`

    // The instructional message displayed in the dialog box when the unblock step is activated.
    Prompt string `json:"prompt,omitempty"`

}

type Command struct {
    // A map of agent tag keys to values to target specific agents for this step.
    Agents map[string]string `json:"agents,omitempty"`

    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // The glob path or paths of artifacts to upload from this step.
    ArtifactPaths []string `json:"artifact_paths,omitempty"`

    // The branch pattern defining which branches will include this block step in their builds.
    Branches string `json:"branches,omitempty"`

    // Setting this attribute to true cancels the job as soon as the build is marked as failing.
    CancelOnBuildFailing bool `json:"cancel_on_build_failing,omitempty"`

    // The shell command to run during this step.
    Commands []string `json:"commands,omitempty"`

    // The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.
    Concurrency int `json:"concurrency,omitempty"`

    // A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.
    ConcurrencyGroup string `json:"concurrency_group,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // A map of environment variables for this step.
    Env map[string]string `json:"env,omitempty"`

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // A unique string to identify the block step.
    Key string `json:"key,omitempty"`

    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    Label string `json:"label,omitempty"`

    // An array of values to be used in the matrix expansion.
    Matrix []string `json:"matrix,omitempty"`

    // The number of parallel jobs that will be created based on this step.
    Parallelism int `json:"parallelism,omitempty"`

    // An array of plugins for this step.
    Plugins []map[string]interface{} `json:"plugins,omitempty"`

    // Adjust the priority for a specific job, as a positive or negative integer.
    Priority int `json:"priority,omitempty"`

    // The conditions for retrying this step.
    Retry *Retry `json:"retry,omitempty"`
    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    Skip bool `json:"skip,omitempty"`

    // Allow specified non-zero exit statuses not to fail the build.
    SoftFail map[string]int `json:"soft_fail,omitempty"`

    // Allow all non-zero exit statuses not to fail the build.
    SoftFailAll bool `json:"soft_fail_all,omitempty"`

    // The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.
    TimeoutInMinutes int `json:"timeout_in_minutes,omitempty"`

}

type Group struct {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // Name of the group in the UI. In YAML, if you don't want a label, pass a `~`. Can also be provided in the `label` attribute if `null` is provided to the `group` attribute.
    Group string `json:"group,omitempty"`

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // A unique string to identify the block step.
    Key string `json:"key,omitempty"`

    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    Label string `json:"label,omitempty"`

    // Allows you to trigger build notifications to different services. You can also choose to conditionally send notifications based on pipeline events.
    Notify string `json:"notify,omitempty"`

    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    Skip bool `json:"skip,omitempty"`

    // A list of steps in the group; at least 1 step is required. Allowed step types: wait, trigger, command/commands, block, input.
    Steps []Steps `json:"steps,omitempty"`
}

type Input struct {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // The branch pattern defining which branches will include this block step in their builds.
    Branches string `json:"branches,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // An input step is used to collect information from a user.
    Fields []Fields `json:"fields,omitempty"`
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // The label for this input step.
    Input string `json:"input,omitempty"`

    // A unique string to identify the block step.
    Key string `json:"key,omitempty"`

    // The instructional message displayed in the dialog box when the unblock step is activated.
    Prompt string `json:"prompt,omitempty"`

}

type Trigger struct {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // If set to true the step will immediately continue, regardless of the success of the triggered build. If set to false the step will wait for the triggered build to complete and continue only if the triggered build passed.
    // Note that when async is set to true, as long as the triggered build starts, the original pipeline will show that as successful. The original pipeline does not get updated after subsequent steps or after the triggered build completes.
    Async bool `json:"async,omitempty"`

    // The branch pattern defining which branches will include this block step in their builds.
    Branches string `json:"branches,omitempty"`

    // An optional map of attributes for the triggered build. Available attributes: branch, commit, env, message, meta_data
    Build *Build `json:"build,omitempty"`
    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    Label string `json:"label,omitempty"`

    // Whether to skip this step or not. Passing a string provides a reason for skipping this command. Passing an empty string is equivalent to false.
    Skip bool `json:"skip,omitempty"`

    // Allow specified non-zero exit statuses not to fail the build.
    SoftFail map[string]int `json:"soft_fail,omitempty"`

    // The slug of the pipeline to create a build. You can find it in the URL of your pipeline, and it corresponds to the name of the pipeline, converted to kebab-case.
    Trigger string `json:"trigger,omitempty"`

}

type Wait struct {
    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // Run the next step, even if the previous step has failed.
    ContinueOnFailure bool `json:"continue_on_failure,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // When providing options for the wait step, you will need to set this value to "~".
    Wait *WaitLabel `json:"wait"`
}
