// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    schema, err := UnmarshalSchema(bytes)
//    bytes, err = schema.Marshal()

package buildkite

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalSchema(data []byte) (Schema, error) {
	var r Schema
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Schema) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Schema struct {
	Agents *Agents                `json:"agents"`
	Env    map[string]interface{} `json:"env,omitempty"`
	Notify []SchemaNotify         `json:"notify,omitempty"`
	// A list of steps
	Steps []SchemaStep `json:"steps"`
}

type PurpleBuildNotify struct {
	Email                *string                   `json:"email,omitempty"`
	If                   *string                   `json:"if,omitempty"`
	BasecampCampfire     *string                   `json:"basecamp_campfire,omitempty"`
	Slack                *StickySlack              `json:"slack"`
	Webhook              *string                   `json:"webhook,omitempty"`
	PagerdutyChangeEvent *string                   `json:"pagerduty_change_event,omitempty"`
	GithubCommitStatus   *PurpleGithubCommitStatus `json:"github_commit_status,omitempty"`
	GithubCheck          *PurpleGithubCheck        `json:"github_check,omitempty"`
}

type PurpleGithubCheck struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type PurpleGithubCommitStatus struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type PurpleSlack struct {
	Channels []string `json:"channels,omitempty"`
	Message  *string  `json:"message,omitempty"`
}

// Waits for previous steps to pass before continuing
type GroupStepClass struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	// The label of the block step
	Block *Block `json:"block"`
	// The state that the build is set to when the build is blocked by this block step
	BlockedState *BlockedState  `json:"blocked_state,omitempty"`
	Branches     *Branches      `json:"branches"`
	DependsOn    *DependsOn     `json:"depends_on"`
	Fields       []Field        `json:"fields,omitempty"`
	ID           *string        `json:"id,omitempty"`
	Identifier   *string        `json:"identifier,omitempty"`
	If           *string        `json:"if,omitempty"`
	Key          *string        `json:"key,omitempty"`
	Label        *string        `json:"label"`
	Name         *string        `json:"name"`
	Prompt       *string        `json:"prompt,omitempty"`
	Type         *BlockStepType `json:"type,omitempty"`
	// The label of the input step
	Input  *Input  `json:"input"`
	Agents *Agents `json:"agents"`
	// The glob path/s of artifacts to upload once this step has finished running
	ArtifactPaths        []string                     `json:"artifact_paths"`
	Cache                *Cache                       `json:"cache"`
	CancelOnBuildFailing *AllowDependencyFailureUnion `json:"cancel_on_build_failing"`
	// The commands to run on the agent
	Command *CommandUnion `json:"command"`
	// The commands to run on the agent
	Commands *CommandUnion `json:"commands"`
	// The maximum number of jobs created from this step that are allowed to run at the same
	// time. If you use this attribute, you must also define concurrency_group.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// A unique name for the concurrency group that you are creating with the concurrency
	// attribute
	ConcurrencyGroup *string `json:"concurrency_group,omitempty"`
	// Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
	// this attribute, you must also define concurrency_group and concurrency.
	ConcurrencyMethod *ConcurrencyMethod     `json:"concurrency_method,omitempty"`
	Env               map[string]interface{} `json:"env,omitempty"`
	Matrix            *MatrixUnion           `json:"matrix"`
	// Array of notification options for this step
	Notify []BlockStepNotify `json:"notify,omitempty"`
	// The number of parallel jobs that will be created based on this step
	Parallelism *int64   `json:"parallelism,omitempty"`
	Plugins     *Plugins `json:"plugins"`
	// Priority of the job, higher priorities are assigned to agents
	Priority *int64 `json:"priority,omitempty"`
	// The conditions for retrying this step.
	Retry *Retry `json:"retry,omitempty"`
	// The signature of the command step, generally injected by agents at pipeline upload
	Signature *Signature `json:"signature,omitempty"`
	Skip      *Skip      `json:"skip"`
	SoftFail  *SoftFail  `json:"soft_fail"`
	// The number of minutes to time out a job
	TimeoutInMinutes *int64       `json:"timeout_in_minutes,omitempty"`
	Script           *CommandStep `json:"script,omitempty"`
	// Continue to the next steps, even if the previous group of steps fail
	ContinueOnFailure *AllowDependencyFailureUnion `json:"continue_on_failure"`
	// Waits for previous steps to pass before continuing
	Wait   *Label `json:"wait"`
	Waiter *Label `json:"waiter"`
	// Whether to continue the build without waiting for the triggered step to complete
	Async *AllowDependencyFailureUnion `json:"async"`
	// Properties of the build that will be created when the step is triggered
	Build *Build `json:"build,omitempty"`
	// The slug of the pipeline to create a build
	Trigger *Trigger `json:"trigger"`
	// The name to give to this group of steps
	Group *string `json:"group"`
	// A list of steps
	Steps []BlockStepStep `json:"steps,omitempty"`
}

type BlockStep struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	// The label of the block step
	Block *string `json:"block,omitempty"`
	// The state that the build is set to when the build is blocked by this block step
	BlockedState *BlockedState `json:"blocked_state,omitempty"`
	Branches     *Branches     `json:"branches"`
	DependsOn    *DependsOn    `json:"depends_on"`
	Fields       []Field       `json:"fields,omitempty"`
	ID           *string       `json:"id,omitempty"`
	Identifier   *string       `json:"identifier,omitempty"`
	If           *string       `json:"if,omitempty"`
	Key          *string       `json:"key,omitempty"`
	Label        *string       `json:"label,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Prompt       *string       `json:"prompt,omitempty"`
	Type         *BlockType    `json:"type,omitempty"`
}

type DependsOnClass struct {
	AllowFailure *AllowDependencyFailureUnion `json:"allow_failure"`
	Step         *string                      `json:"step,omitempty"`
}

// A list of input fields required to be filled out before unblocking the step
type Field struct {
	// The value that is pre-filled in the text field
	//
	// The value of the option(s) that will be pre-selected in the dropdown
	Default *Branches `json:"default"`
	// The format must be a regular expression implicitly anchored to the beginning and end of
	// the input and is functionally equivalent to the HTML5 pattern attribute.
	Format *string `json:"format,omitempty"`
	// The explanatory text that is shown after the label
	Hint *string `json:"hint,omitempty"`
	// The meta-data key that stores the field's input
	Key string `json:"key"`
	// Whether the field is required for form submission
	Required *AllowDependencyFailureUnion `json:"required"`
	// The text input name
	Text *string `json:"text,omitempty"`
	// Whether more than one option may be selected
	Multiple *AllowDependencyFailureUnion `json:"multiple"`
	Options  []Option                     `json:"options,omitempty"`
	// The text input name
	Select *string `json:"select,omitempty"`
}

type Option struct {
	// The text displayed directly under the select field’s label
	Hint *string `json:"hint,omitempty"`
	// The text displayed on the select list item
	Label string `json:"label"`
	// Whether the field is required for form submission
	Required *AllowDependencyFailureUnion `json:"required"`
	// The value to be stored as meta-data
	Value string `json:"value"`
}

// Properties of the build that will be created when the step is triggered
type Build struct {
	// The branch for the build
	Branch *string `json:"branch,omitempty"`
	// The commit hash for the build
	Commit *string                `json:"commit,omitempty"`
	Env    map[string]interface{} `json:"env,omitempty"`
	// The message for the build (supports emoji)
	Message *string `json:"message,omitempty"`
	// Meta-data for the build
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
}

type CacheClass struct {
	Name  *string  `json:"name,omitempty"`
	Paths []string `json:"paths"`
	Size  *string  `json:"size,omitempty"`
}

type CommandStep struct {
	Agents                 *Agents                      `json:"agents"`
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	// The glob path/s of artifacts to upload once this step has finished running
	ArtifactPaths        []string                     `json:"artifact_paths"`
	Branches             *Branches                    `json:"branches"`
	Cache                *Cache                       `json:"cache"`
	CancelOnBuildFailing *AllowDependencyFailureUnion `json:"cancel_on_build_failing"`
	// The commands to run on the agent
	Command *CommandUnion `json:"command"`
	// The commands to run on the agent
	Commands *CommandUnion `json:"commands"`
	// The maximum number of jobs created from this step that are allowed to run at the same
	// time. If you use this attribute, you must also define concurrency_group.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// A unique name for the concurrency group that you are creating with the concurrency
	// attribute
	ConcurrencyGroup *string `json:"concurrency_group,omitempty"`
	// Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
	// this attribute, you must also define concurrency_group and concurrency.
	ConcurrencyMethod *ConcurrencyMethod     `json:"concurrency_method,omitempty"`
	DependsOn         *DependsOn             `json:"depends_on"`
	Env               map[string]interface{} `json:"env,omitempty"`
	ID                *string                `json:"id,omitempty"`
	Identifier        *string                `json:"identifier,omitempty"`
	If                *string                `json:"if,omitempty"`
	Key               *string                `json:"key,omitempty"`
	Label             *string                `json:"label,omitempty"`
	Matrix            *MatrixUnion           `json:"matrix"`
	Name              *string                `json:"name,omitempty"`
	// Array of notification options for this step
	Notify []NotifyElement `json:"notify,omitempty"`
	// The number of parallel jobs that will be created based on this step
	Parallelism *int64   `json:"parallelism,omitempty"`
	Plugins     *Plugins `json:"plugins"`
	// Priority of the job, higher priorities are assigned to agents
	Priority *int64 `json:"priority,omitempty"`
	// The conditions for retrying this step.
	Retry *Retry `json:"retry,omitempty"`
	// The signature of the command step, generally injected by agents at pipeline upload
	Signature *Signature `json:"signature,omitempty"`
	Skip      *Skip      `json:"skip"`
	SoftFail  *SoftFail  `json:"soft_fail"`
	// The number of minutes to time out a job
	TimeoutInMinutes *int64      `json:"timeout_in_minutes,omitempty"`
	Type             *ScriptType `json:"type,omitempty"`
}

// Configuration for multi-dimension Build Matrix
type MatrixClass struct {
	// List of Build Matrix adjustments
	Adjustments []Adjustment `json:"adjustments,omitempty"`
	Setup       *Setup       `json:"setup"`
}

// An adjustment to a Build Matrix
type Adjustment struct {
	Skip     *Skip     `json:"skip"`
	SoftFail *SoftFail `json:"soft_fail"`
	With     *With     `json:"with"`
}

type SoftFailElement struct {
	// The exit status number that will cause this job to soft-fail
	ExitStatus *SoftFailExitStatus `json:"exit_status"`
}

type NotifyClass struct {
	BasecampCampfire   *string                   `json:"basecamp_campfire,omitempty"`
	If                 *string                   `json:"if,omitempty"`
	Slack              *IndigoSlack              `json:"slack"`
	GithubCommitStatus *FluffyGithubCommitStatus `json:"github_commit_status,omitempty"`
	GithubCheck        *FluffyGithubCheck        `json:"github_check,omitempty"`
}

type FluffyGithubCheck struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type FluffyGithubCommitStatus struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type FluffySlack struct {
	Channels []string `json:"channels,omitempty"`
	Message  *string  `json:"message,omitempty"`
}

// The conditions for retrying this step.
type Retry struct {
	// Whether to allow a job to retry automatically. If set to true, the retry conditions are
	// set to the default value.
	Automatic *Automatic `json:"automatic"`
	// Whether to allow a job to be retried manually
	Manual *ManualUnion `json:"manual"`
}

type AutomaticRetry struct {
	// The exit status number that will cause this job to retry
	ExitStatus *AutomaticRetryExitStatus `json:"exit_status"`
	// The number of times this job can be retried
	Limit *int64 `json:"limit,omitempty"`
	// The exit signal, if any, that may be retried
	Signal *string `json:"signal,omitempty"`
	// The exit signal reason, if any, that may be retried
	SignalReason *SignalReasonEnum `json:"signal_reason,omitempty"`
}

type ManualClass struct {
	// Whether or not this job can be retried manually
	Allowed *AllowDependencyFailureUnion `json:"allowed"`
	// Whether or not this job can be retried after it has passed
	PermitOnPassed *AllowDependencyFailureUnion `json:"permit_on_passed"`
	// A string that will be displayed in a tooltip on the Retry button in Buildkite. This will
	// only be displayed if the allowed attribute is set to false.
	Reason *string `json:"reason,omitempty"`
}

// The signature of the command step, generally injected by agents at pipeline upload
type Signature struct {
	// The algorithm used to generate the signature
	Algorithm *string `json:"algorithm,omitempty"`
	// The fields that were signed to form the signature value
	SignedFields []string `json:"signed_fields,omitempty"`
	// The signature value, a JWS compact signature with a detached body
	Value *string `json:"value,omitempty"`
}

type InputStep struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	Branches               *Branches                    `json:"branches"`
	DependsOn              *DependsOn                   `json:"depends_on"`
	Fields                 []Field                      `json:"fields,omitempty"`
	ID                     *string                      `json:"id,omitempty"`
	Identifier             *string                      `json:"identifier,omitempty"`
	If                     *string                      `json:"if,omitempty"`
	// The label of the input step
	Input  *string    `json:"input,omitempty"`
	Key    *string    `json:"key,omitempty"`
	Label  *string    `json:"label,omitempty"`
	Name   *string    `json:"name,omitempty"`
	Prompt *string    `json:"prompt,omitempty"`
	Type   *InputType `json:"type,omitempty"`
}

type FluffyBuildNotify struct {
	BasecampCampfire     *string                      `json:"basecamp_campfire,omitempty"`
	If                   *string                      `json:"if,omitempty"`
	Slack                *IndecentSlack               `json:"slack"`
	GithubCommitStatus   *TentacledGithubCommitStatus `json:"github_commit_status,omitempty"`
	GithubCheck          *TentacledGithubCheck        `json:"github_check,omitempty"`
	Email                *string                      `json:"email,omitempty"`
	Webhook              *string                      `json:"webhook,omitempty"`
	PagerdutyChangeEvent *string                      `json:"pagerduty_change_event,omitempty"`
}

type TentacledGithubCheck struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type TentacledGithubCommitStatus struct {
	// GitHub commit status name
	Context *string `json:"context,omitempty"`
}

type TentacledSlack struct {
	Channels []string `json:"channels,omitempty"`
	Message  *string  `json:"message,omitempty"`
}

// Waits for previous steps to pass before continuing
type PurpleStep struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	// The label of the block step
	Block *Block `json:"block"`
	// The state that the build is set to when the build is blocked by this block step
	BlockedState *BlockedState  `json:"blocked_state,omitempty"`
	Branches     *Branches      `json:"branches"`
	DependsOn    *DependsOn     `json:"depends_on"`
	Fields       []Field        `json:"fields,omitempty"`
	ID           *string        `json:"id,omitempty"`
	Identifier   *string        `json:"identifier,omitempty"`
	If           *string        `json:"if,omitempty"`
	Key          *string        `json:"key,omitempty"`
	Label        *string        `json:"label"`
	Name         *string        `json:"name"`
	Prompt       *string        `json:"prompt,omitempty"`
	Type         *BlockStepType `json:"type,omitempty"`
	// The label of the input step
	Input  *Input  `json:"input"`
	Agents *Agents `json:"agents"`
	// The glob path/s of artifacts to upload once this step has finished running
	ArtifactPaths        []string                     `json:"artifact_paths"`
	Cache                *Cache                       `json:"cache"`
	CancelOnBuildFailing *AllowDependencyFailureUnion `json:"cancel_on_build_failing"`
	// The commands to run on the agent
	Command *CommandUnion `json:"command"`
	// The commands to run on the agent
	Commands *CommandUnion `json:"commands"`
	// The maximum number of jobs created from this step that are allowed to run at the same
	// time. If you use this attribute, you must also define concurrency_group.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// A unique name for the concurrency group that you are creating with the concurrency
	// attribute
	ConcurrencyGroup *string `json:"concurrency_group,omitempty"`
	// Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
	// this attribute, you must also define concurrency_group and concurrency.
	ConcurrencyMethod *ConcurrencyMethod     `json:"concurrency_method,omitempty"`
	Env               map[string]interface{} `json:"env,omitempty"`
	Matrix            *MatrixUnion           `json:"matrix"`
	// Array of notification options for this step
	Notify []NotifyElement `json:"notify,omitempty"`
	// The number of parallel jobs that will be created based on this step
	Parallelism *int64   `json:"parallelism,omitempty"`
	Plugins     *Plugins `json:"plugins"`
	// Priority of the job, higher priorities are assigned to agents
	Priority *int64 `json:"priority,omitempty"`
	// The conditions for retrying this step.
	Retry *Retry `json:"retry,omitempty"`
	// The signature of the command step, generally injected by agents at pipeline upload
	Signature *Signature `json:"signature,omitempty"`
	Skip      *Skip      `json:"skip"`
	SoftFail  *SoftFail  `json:"soft_fail"`
	// The number of minutes to time out a job
	TimeoutInMinutes *int64       `json:"timeout_in_minutes,omitempty"`
	Script           *CommandStep `json:"script,omitempty"`
	// Continue to the next steps, even if the previous group of steps fail
	ContinueOnFailure *AllowDependencyFailureUnion `json:"continue_on_failure"`
	// Waits for previous steps to pass before continuing
	Wait   *Label `json:"wait"`
	Waiter *Label `json:"waiter"`
	// Whether to continue the build without waiting for the triggered step to complete
	Async *AllowDependencyFailureUnion `json:"async"`
	// Properties of the build that will be created when the step is triggered
	Build *Build `json:"build,omitempty"`
	// The slug of the pipeline to create a build
	Trigger *Trigger `json:"trigger"`
}

type TriggerStep struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	// Whether to continue the build without waiting for the triggered step to complete
	Async    *AllowDependencyFailureUnion `json:"async"`
	Branches *Branches                    `json:"branches"`
	// Properties of the build that will be created when the step is triggered
	Build      *Build     `json:"build,omitempty"`
	DependsOn  *DependsOn `json:"depends_on"`
	ID         *string    `json:"id,omitempty"`
	Identifier *string    `json:"identifier,omitempty"`
	If         *string    `json:"if,omitempty"`
	Key        *string    `json:"key,omitempty"`
	Label      *string    `json:"label,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Skip       *Skip      `json:"skip"`
	SoftFail   *SoftFail  `json:"soft_fail"`
	// The slug of the pipeline to create a build
	Trigger string       `json:"trigger"`
	Type    *TriggerType `json:"type,omitempty"`
}

// Waits for previous steps to pass before continuing
type WaitStep struct {
	AllowDependencyFailure *AllowDependencyFailureUnion `json:"allow_dependency_failure"`
	Branches               *Branches                    `json:"branches"`
	// Continue to the next steps, even if the previous group of steps fail
	ContinueOnFailure *AllowDependencyFailureUnion `json:"continue_on_failure"`
	DependsOn         *DependsOn                   `json:"depends_on"`
	ID                *string                      `json:"id,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	If                *string                      `json:"if,omitempty"`
	Key               *string                      `json:"key,omitempty"`
	Label             *string                      `json:"label"`
	Name              *string                      `json:"name"`
	Type              *WaitType                    `json:"type,omitempty"`
	// Waits for previous steps to pass before continuing
	Wait   *string `json:"wait"`
	Waiter *string `json:"waiter"`
}

type NotifyEnum string

const (
	GithubCheck        NotifyEnum = "github_check"
	GithubCommitStatus NotifyEnum = "github_commit_status"
)

type AllowDependencyFailureEnum string

const (
	False AllowDependencyFailureEnum = "false"
	True  AllowDependencyFailureEnum = "true"
)

// The state that the build is set to when the build is blocked by this block step
type BlockedState string

const (
	Failed  BlockedState = "failed"
	Passed  BlockedState = "passed"
	Running BlockedState = "running"
)

type BlockType string

const (
	PurpleBlock BlockType = "block"
)

// Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
// this attribute, you must also define concurrency_group and concurrency.
type ConcurrencyMethod string

const (
	Eager   ConcurrencyMethod = "eager"
	Ordered ConcurrencyMethod = "ordered"
)

type ExitStatusEnum string

const (
	ExitStatus ExitStatusEnum = "*"
)

// The exit signal reason, if any, that may be retried
type SignalReasonEnum string

const (
	AgentRefused      SignalReasonEnum = "agent_refused"
	AgentStop         SignalReasonEnum = "agent_stop"
	Cancel            SignalReasonEnum = "cancel"
	None              SignalReasonEnum = "none"
	ProcessRunError   SignalReasonEnum = "process_run_error"
	SignalReason      SignalReasonEnum = "*"
	SignatureRejected SignalReasonEnum = "signature_rejected"
)

type ScriptType string

const (
	PurpleCommand  ScriptType = "command"
	PurpleCommands ScriptType = "commands"
	PurpleScript   ScriptType = "script"
)

type InputType string

const (
	PurpleInput InputType = "input"
)

type TriggerType string

const (
	PurpleTrigger TriggerType = "trigger"
)

type BlockStepType string

const (
	FluffyBlock    BlockStepType = "block"
	FluffyCommand  BlockStepType = "command"
	FluffyCommands BlockStepType = "commands"
	FluffyInput    BlockStepType = "input"
	FluffyScript   BlockStepType = "script"
	FluffyTrigger  BlockStepType = "trigger"
	PurpleWait     BlockStepType = "wait"
	PurpleWaiter   BlockStepType = "waiter"
)

type WaitType string

const (
	FluffyWait   WaitType = "wait"
	FluffyWaiter WaitType = "waiter"
)

// Pauses the execution of a build and waits on a user to unblock it
//
// Waits for previous steps to pass before continuing
type StringStep string

const (
	StringStepBlock  StringStep = "block"
	StringStepInput  StringStep = "input"
	StringStepWait   StringStep = "wait"
	StringStepWaiter StringStep = "waiter"
)

type Agents struct {
	AnythingMap map[string]interface{}
	StringArray []string
}

func (x *Agents) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Agents) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

// Array of notification options for this step
type SchemaNotify struct {
	Enum              *NotifyEnum
	PurpleBuildNotify *PurpleBuildNotify
}

func (x *SchemaNotify) UnmarshalJSON(data []byte) error {
	x.PurpleBuildNotify = nil
	x.Enum = nil
	var c PurpleBuildNotify
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleBuildNotify = &c
	}
	return nil
}

func (x *SchemaNotify) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.PurpleBuildNotify != nil, x.PurpleBuildNotify, false, nil, x.Enum != nil, x.Enum, false)
}

type StickySlack struct {
	PurpleSlack *PurpleSlack
	String      *string
}

func (x *StickySlack) UnmarshalJSON(data []byte) error {
	x.PurpleSlack = nil
	var c PurpleSlack
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleSlack = &c
	}
	return nil
}

func (x *StickySlack) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleSlack != nil, x.PurpleSlack, false, nil, false, nil, false)
}

type SchemaStep struct {
	Enum           *StringStep
	GroupStepClass *GroupStepClass
}

func (x *SchemaStep) UnmarshalJSON(data []byte) error {
	x.GroupStepClass = nil
	x.Enum = nil
	var c GroupStepClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.GroupStepClass = &c
	}
	return nil
}

func (x *SchemaStep) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.GroupStepClass != nil, x.GroupStepClass, false, nil, x.Enum != nil, x.Enum, false)
}

type AllowDependencyFailureUnion struct {
	Bool *bool
	Enum *AllowDependencyFailureEnum
}

func (x *AllowDependencyFailureUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AllowDependencyFailureUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// Which branches will include this step in their builds
//
// # The value of the option(s) that will be pre-selected in the dropdown
//
// The glob path/s of artifacts to upload once this step has finished running
//
// The commands to run on the agent
type Branches struct {
	String      *string
	StringArray []string
}

func (x *Branches) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Branches) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type Block struct {
	BlockStep *BlockStep
	String    *string
}

func (x *Block) UnmarshalJSON(data []byte) error {
	x.BlockStep = nil
	var c BlockStep
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.BlockStep = &c
	}
	return nil
}

func (x *Block) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.BlockStep != nil, x.BlockStep, false, nil, false, nil, false)
}

// The step keys for a step to depend on
type DependsOn struct {
	String     *string
	UnionArray []DependsOnElement
}

func (x *DependsOn) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *DependsOn) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, true)
}

type DependsOnElement struct {
	DependsOnClass *DependsOnClass
	String         *string
}

func (x *DependsOnElement) UnmarshalJSON(data []byte) error {
	x.DependsOnClass = nil
	var c DependsOnClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.DependsOnClass = &c
	}
	return nil
}

func (x *DependsOnElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.DependsOnClass != nil, x.DependsOnClass, false, nil, false, nil, false)
}

// The paths for the caches to be used in the step
type Cache struct {
	CacheClass  *CacheClass
	String      *string
	StringArray []string
}

func (x *Cache) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.CacheClass = nil
	var c CacheClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CacheClass = &c
	}
	return nil
}

func (x *Cache) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, x.CacheClass != nil, x.CacheClass, false, nil, false, nil, false)
}

type CommandUnion struct {
	CommandStep *CommandStep
	String      *string
	StringArray []string
}

func (x *CommandUnion) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.CommandStep = nil
	var c CommandStep
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CommandStep = &c
	}
	return nil
}

func (x *CommandUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, x.CommandStep != nil, x.CommandStep, false, nil, false, nil, false)
}

type MatrixUnion struct {
	MatrixClass *MatrixClass
	UnionArray  []MatrixElement
}

func (x *MatrixUnion) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.MatrixClass = nil
	var c MatrixClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.UnionArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.MatrixClass = &c
	}
	return nil
}

func (x *MatrixUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.UnionArray != nil, x.UnionArray, x.MatrixClass != nil, x.MatrixClass, false, nil, false, nil, false)
}

// List of elements for simple single-dimension Build Matrix
//
// # List of existing or new elements for single-dimension Build Matrix
//
// # List of elements for single-dimension Build Matrix
//
// List of elements for this Build Matrix dimension
type MatrixElement struct {
	Bool    *bool
	Integer *int64
	String  *string
}

func (x *MatrixElement) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MatrixElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

// Whether this step should be skipped. Passing a string provides a reason for skipping this
// command
type Skip struct {
	Bool   *bool
	String *string
}

func (x *Skip) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Skip) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

// The conditions for marking the step as a soft-fail.
type SoftFail struct {
	Bool                 *bool
	Enum                 *AllowDependencyFailureEnum
	SoftFailElementArray []SoftFailElement
}

func (x *SoftFail) UnmarshalJSON(data []byte) error {
	x.SoftFailElementArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.SoftFailElementArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *SoftFail) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.SoftFailElementArray != nil, x.SoftFailElementArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// The exit status number that will cause this job to soft-fail
type SoftFailExitStatus struct {
	Enum    *ExitStatusEnum
	Integer *int64
}

func (x *SoftFailExitStatus) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *SoftFailExitStatus) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

type With struct {
	StringMap  map[string]string
	UnionArray []MatrixElement
}

func (x *With) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.UnionArray, false, nil, true, &x.StringMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *With) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, x.StringMap != nil, x.StringMap, false, nil, false)
}

type Setup struct {
	UnionArray    []MatrixElement
	UnionArrayMap map[string][]MatrixElement
}

func (x *Setup) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.UnionArrayMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.UnionArray, false, nil, true, &x.UnionArrayMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Setup) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, x.UnionArrayMap != nil, x.UnionArrayMap, false, nil, false)
}

type NotifyElement struct {
	Enum        *NotifyEnum
	NotifyClass *NotifyClass
}

func (x *NotifyElement) UnmarshalJSON(data []byte) error {
	x.NotifyClass = nil
	x.Enum = nil
	var c NotifyClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.NotifyClass = &c
	}
	return nil
}

func (x *NotifyElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.NotifyClass != nil, x.NotifyClass, false, nil, x.Enum != nil, x.Enum, false)
}

type IndigoSlack struct {
	FluffySlack *FluffySlack
	String      *string
}

func (x *IndigoSlack) UnmarshalJSON(data []byte) error {
	x.FluffySlack = nil
	var c FluffySlack
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffySlack = &c
	}
	return nil
}

func (x *IndigoSlack) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffySlack != nil, x.FluffySlack, false, nil, false, nil, false)
}

type Plugins struct {
	AnythingMap map[string]interface{}
	UnionArray  []Plugin
}

func (x *Plugins) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.UnionArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Plugins) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

// Array of plugins for this step
type Plugin struct {
	AnythingMap map[string]interface{}
	String      *string
}

func (x *Plugin) UnmarshalJSON(data []byte) error {
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Plugin) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

// Whether to allow a job to retry automatically. If set to true, the retry conditions are
// set to the default value.
type Automatic struct {
	AutomaticRetry      *AutomaticRetry
	AutomaticRetryArray []AutomaticRetry
	Bool                *bool
	Enum                *AllowDependencyFailureEnum
}

func (x *Automatic) UnmarshalJSON(data []byte) error {
	x.AutomaticRetryArray = nil
	x.AutomaticRetry = nil
	x.Enum = nil
	var c AutomaticRetry
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.AutomaticRetryArray, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.AutomaticRetry = &c
	}
	return nil
}

func (x *Automatic) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.AutomaticRetryArray != nil, x.AutomaticRetryArray, x.AutomaticRetry != nil, x.AutomaticRetry, false, nil, x.Enum != nil, x.Enum, false)
}

// The exit status number that will cause this job to retry
type AutomaticRetryExitStatus struct {
	Enum         *ExitStatusEnum
	Integer      *int64
	IntegerArray []int64
}

func (x *AutomaticRetryExitStatus) UnmarshalJSON(data []byte) error {
	x.IntegerArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, nil, true, &x.IntegerArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AutomaticRetryExitStatus) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, nil, x.IntegerArray != nil, x.IntegerArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// Whether to allow a job to be retried manually
type ManualUnion struct {
	Bool        *bool
	Enum        *AllowDependencyFailureEnum
	ManualClass *ManualClass
}

func (x *ManualUnion) UnmarshalJSON(data []byte) error {
	x.ManualClass = nil
	x.Enum = nil
	var c ManualClass
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.ManualClass = &c
	}
	return nil
}

func (x *ManualUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, x.ManualClass != nil, x.ManualClass, false, nil, x.Enum != nil, x.Enum, false)
}

type Input struct {
	InputStep *InputStep
	String    *string
}

func (x *Input) UnmarshalJSON(data []byte) error {
	x.InputStep = nil
	var c InputStep
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.InputStep = &c
	}
	return nil
}

func (x *Input) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.InputStep != nil, x.InputStep, false, nil, false, nil, false)
}

type BlockStepNotify struct {
	Enum              *NotifyEnum
	FluffyBuildNotify *FluffyBuildNotify
}

func (x *BlockStepNotify) UnmarshalJSON(data []byte) error {
	x.FluffyBuildNotify = nil
	x.Enum = nil
	var c FluffyBuildNotify
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyBuildNotify = &c
	}
	return nil
}

func (x *BlockStepNotify) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.FluffyBuildNotify != nil, x.FluffyBuildNotify, false, nil, x.Enum != nil, x.Enum, false)
}

type IndecentSlack struct {
	String         *string
	TentacledSlack *TentacledSlack
}

func (x *IndecentSlack) UnmarshalJSON(data []byte) error {
	x.TentacledSlack = nil
	var c TentacledSlack
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledSlack = &c
	}
	return nil
}

func (x *IndecentSlack) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledSlack != nil, x.TentacledSlack, false, nil, false, nil, false)
}

type BlockStepStep struct {
	Enum       *StringStep
	PurpleStep *PurpleStep
}

func (x *BlockStepStep) UnmarshalJSON(data []byte) error {
	x.PurpleStep = nil
	x.Enum = nil
	var c PurpleStep
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleStep = &c
	}
	return nil
}

func (x *BlockStepStep) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.PurpleStep != nil, x.PurpleStep, false, nil, x.Enum != nil, x.Enum, false)
}

type Trigger struct {
	String      *string
	TriggerStep *TriggerStep
}

func (x *Trigger) UnmarshalJSON(data []byte) error {
	x.TriggerStep = nil
	var c TriggerStep
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TriggerStep = &c
	}
	return nil
}

func (x *Trigger) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TriggerStep != nil, x.TriggerStep, false, nil, false, nil, false)
}

type Label struct {
	String   *string
	WaitStep *WaitStep
}

func (x *Label) UnmarshalJSON(data []byte) error {
	x.WaitStep = nil
	var c WaitStep
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
		x.WaitStep = &c
	}
	return nil
}

func (x *Label) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.WaitStep != nil, x.WaitStep, false, nil, false, nil, true)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
