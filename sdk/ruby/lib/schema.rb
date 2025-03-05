# This code may look unusually verbose for Ruby (and it is), but
# it performs some subtle and complex validation of JSON data.
#
# To parse this JSON, add 'dry-struct' and 'dry-types' gems, then do:
#
#   schema = Schema.from_json! "{…}"
#   puts schema.steps.first
#
# If from_json! succeeds, the value returned matches the schema.

require "json"
require "dry-types"
require "dry-struct"

module Types
  include Dry.Types(default: :nominal)

  Integer                    = Strict::Integer
  Nil                        = Strict::Nil
  Bool                       = Strict::Bool
  Hash                       = Strict::Hash
  String                     = Strict::String
  NotifyEnum                 = Strict::String.enum("github_check", "github_commit_status")
  AllowDependencyFailureEnum = Strict::String.enum("false", "true")
  BlockType                  = Strict::String.enum("block")
  BlockedState               = Strict::String.enum("failed", "passed", "running")
  ExitStatusEnum             = Strict::String.enum("*")
  SignalReason               = Strict::String.enum("agent_refused", "agent_stop", "cancel", "*", "none",
                                                   "process_run_error", "signature_rejected")
  ScriptType                 = Strict::String.enum("command", "commands", "script")
  ConcurrencyMethod          = Strict::String.enum("eager", "ordered")
  InputType                  = Strict::String.enum("input")
  BlockStepType              = Strict::String.enum("block", "command", "commands", "input", "script", "trigger",
                                                   "wait", "waiter")
  TriggerType                = Strict::String.enum("trigger")
  WaiterType                 = Strict::String.enum("wait", "waiter")
  StringStep                 = Strict::String.enum("block", "input", "wait", "waiter")
end

class Agents < Dry::Struct
  attribute :anything_map, Types::Hash.meta(of: Types::Any).optional
  attribute :string_array, Types.Array(Types::String).optional

  def self.from_dynamic!(d)
    begin
      value = Types::Hash[d].transform_values { |v| Types::Any[v] }
      return new(anything_map: value, string_array: nil) if schema[:anything_map].right.valid? value
    rescue StandardError
    end
    return new(string_array: d, anything_map: nil) if schema[:string_array].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !anything_map.nil?
      anything_map
    elsif !string_array.nil?
      string_array
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class GithubCommitStatus1 < Dry::Struct
  # GitHub commit status name
  attribute :context, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      context: d["context"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "context" => context
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack1 < Dry::Struct
  attribute :channels, Types.Array(Types::String).optional
  attribute :message,  Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      channels: d["channels"],
      message: d["message"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "channels" => channels,
      "message" => message
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack4 < Dry::Struct
  attribute :slack1, Slack1.optional
  attribute :string, Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = Slack1.from_dynamic!(d)
      return new(slack1: value, string: nil) if schema[:slack1].right.valid? value
    rescue StandardError
    end
    return new(string: d, slack1: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !slack1.nil?
      slack1.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class BuildNotify1 < Dry::Struct
  attribute :email,                  Types::String.optional
  attribute :build_notify_if,        Types::String.optional
  attribute :basecamp_campfire,      Types::String.optional
  attribute :slack,                  Types.Instance(Slack4).optional
  attribute :webhook,                Types::String.optional
  attribute :pagerduty_change_event, Types::String.optional
  attribute :github_commit_status,   GithubCommitStatus1.optional
  attribute :github_check,           Types::Hash.meta(of: Types::Any).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      email: d["email"],
      build_notify_if: d["if"],
      basecamp_campfire: d["basecamp_campfire"],
      slack: d["slack"] ? Slack4.from_dynamic!(d["slack"]) : nil,
      webhook: d["webhook"],
      pagerduty_change_event: d["pagerduty_change_event"],
      github_commit_status: d["github_commit_status"] ? GithubCommitStatus1.from_dynamic!(d["github_commit_status"]) : nil,
      github_check: Types::Hash.optional[d["github_check"]]&.transform_values { |v| Types::Any[v] }
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "email" => email,
      "if" => build_notify_if,
      "basecamp_campfire" => basecamp_campfire,
      "slack" => slack&.to_dynamic,
      "webhook" => webhook,
      "pagerduty_change_event" => pagerduty_change_event,
      "github_commit_status" => github_commit_status&.to_dynamic,
      "github_check" => github_check
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module NotifyEnum
  GithubCheck        = "github_check".freeze
  GithubCommitStatus = "github_commit_status".freeze
end

# Array of notification options for this step
class SchemaNotify < Dry::Struct
  attribute :build_notify1, BuildNotify1.optional
  attribute :enum,          Types::NotifyEnum.optional

  def self.from_dynamic!(d)
    begin
      value = BuildNotify1.from_dynamic!(d)
      return new(build_notify1: value, enum: nil) if schema[:build_notify1].right.valid? value
    rescue StandardError
    end
    return new(enum: d, build_notify1: nil) if schema[:enum].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !build_notify1.nil?
      build_notify1.to_dynamic
    elsif !enum.nil?
      enum
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module AllowDependencyFailureEnum
  False = "false".freeze
  True  = "true".freeze
end

class AllowDependencyFailureUnion < Dry::Struct
  attribute :bool, Types::Bool.optional
  attribute :enum, Types::AllowDependencyFailureEnum.optional

  def self.from_dynamic!(d)
    return new(bool: d, enum: nil) if schema[:bool].right.valid? d
    return new(enum: d, bool: nil) if schema[:enum].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !bool.nil?
      bool
    elsif !enum.nil?
      enum
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Which branches will include this step in their builds
#
# The value of the option(s) that will be pre-selected in the dropdown
#
# The glob path/s of artifacts to upload once this step has finished running
#
# The commands to run on the agent
class Branches < Dry::Struct
  attribute :string,       Types::String.optional
  attribute :string_array, Types.Array(Types::String).optional

  def self.from_dynamic!(d)
    return new(string: d, string_array: nil) if schema[:string].right.valid? d
    return new(string_array: d, string: nil) if schema[:string_array].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !string.nil?
      string
    elsif !string_array.nil?
      string_array
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module BlockType
  Block = "block".freeze
end

# The state that the build is set to when the build is blocked by this block step
module BlockedState
  Failed  = "failed".freeze
  Passed  = "passed".freeze
  Running = "running".freeze
end

class DependsOnClass < Dry::Struct
  attribute :allow_failure, Types.Instance(AllowDependencyFailureUnion).optional
  attribute :step,          Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_failure: d["allow_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_failure"]) : nil,
      step: d["step"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_failure" => allow_failure&.to_dynamic,
      "step" => step
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class DependsOnElement < Dry::Struct
  attribute :depends_on_class, DependsOnClass.optional
  attribute :string,           Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = DependsOnClass.from_dynamic!(d)
      return new(depends_on_class: value, string: nil) if schema[:depends_on_class].right.valid? value
    rescue StandardError
    end
    return new(string: d, depends_on_class: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !depends_on_class.nil?
      depends_on_class.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The step keys for a step to depend on
class DependsOn < Dry::Struct
  attribute :null,        Types::Nil.optional
  attribute :string,      Types::String.optional
  attribute :union_array, Types.Array(Types.Instance(DependsOnElement)).optional

  def self.from_dynamic!(d)
    return new(null: d, union_array: nil, string: nil) if schema[:null].right.valid? d
    return new(string: d, union_array: nil, null: nil) if schema[:string].right.valid? d

    begin
      value = d.map { |x| DependsOnElement.from_dynamic!(x) }
      return new(union_array: value, null: nil, string: nil) if schema[:union_array].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !string.nil?
      string
    elsif !union_array.nil?
      union_array.map(&:to_dynamic)
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Option < Dry::Struct
  # The text displayed directly under the select field’s label
  attribute :hint, Types::String.optional

  # The text displayed on the select list item
  attribute :label, Types::String

  # Whether the field is required for form submission
  attribute :required, Types.Instance(AllowDependencyFailureUnion).optional

  # The value to be stored as meta-data
  attribute :value, Types::String

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      hint: d["hint"],
      label: d.fetch("label"),
      required: d["required"] ? AllowDependencyFailureUnion.from_dynamic!(d["required"]) : nil,
      value: d.fetch("value")
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "hint" => hint,
      "label" => label,
      "required" => required&.to_dynamic,
      "value" => value
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# A list of input fields required to be filled out before unblocking the step
class Field < Dry::Struct
  # The value that is pre-filled in the text field
  #
  # The value of the option(s) that will be pre-selected in the dropdown
  attribute :field_default, Types.Instance(Branches).optional

  # The format must be a regular expression implicitly anchored to the beginning and end of
  # the input and is functionally equivalent to the HTML5 pattern attribute.
  attribute :field_format, Types::String.optional

  # The explanatory text that is shown after the label
  attribute :hint, Types::String.optional

  # The meta-data key that stores the field's input
  attribute :key, Types::String

  # Whether the field is required for form submission
  attribute :required, Types.Instance(AllowDependencyFailureUnion).optional

  # The text input name
  attribute :text, Types::String.optional

  # Whether more than one option may be selected
  attribute :multiple, Types.Instance(AllowDependencyFailureUnion).optional

  attribute :field_options, Types.Array(Option).optional

  # The text input name
  attribute :field_select, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      field_default: d["default"] ? Branches.from_dynamic!(d["default"]) : nil,
      field_format: d["format"],
      hint: d["hint"],
      key: d.fetch("key"),
      required: d["required"] ? AllowDependencyFailureUnion.from_dynamic!(d["required"]) : nil,
      text: d["text"],
      multiple: d["multiple"] ? AllowDependencyFailureUnion.from_dynamic!(d["multiple"]) : nil,
      field_options: d["options"]&.map { |x| Option.from_dynamic!(x) },
      field_select: d["select"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "default" => field_default&.to_dynamic,
      "format" => field_format,
      "hint" => hint,
      "key" => key,
      "required" => required&.to_dynamic,
      "text" => text,
      "multiple" => multiple&.to_dynamic,
      "options" => field_options&.map(&:to_dynamic),
      "select" => field_select
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class BlockStep < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # The label of the block step
  attribute :block, Types::String.optional

  # The state that the build is set to when the build is blocked by this block step
  attribute :blocked_state, Types::BlockedState.optional

  attribute :branches,        Types.Instance(Branches).optional
  attribute :depends_on,      Types.Instance(DependsOn).optional
  attribute :fields,          Types.Array(Field).optional
  attribute :id,              Types::String.optional
  attribute :identifier,      Types::String.optional
  attribute :block_step_if,   Types::String.optional
  attribute :key,             Types::String.optional
  attribute :label,           Types::String.optional
  attribute :block_step_name, Types::String.optional
  attribute :prompt,          Types::String.optional
  attribute :block_step_type, Types::BlockType.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      block: d["block"],
      blocked_state: d["blocked_state"],
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      fields: d["fields"]&.map { |x| Field.from_dynamic!(x) },
      id: d["id"],
      identifier: d["identifier"],
      block_step_if: d["if"],
      key: d["key"],
      label: d["label"],
      block_step_name: d["name"],
      prompt: d["prompt"],
      block_step_type: d["type"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "block" => block,
      "blocked_state" => blocked_state,
      "branches" => branches&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "fields" => fields&.map(&:to_dynamic),
      "id" => id,
      "identifier" => identifier,
      "if" => block_step_if,
      "key" => key,
      "label" => label,
      "name" => block_step_name,
      "prompt" => prompt,
      "type" => block_step_type
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Block < Dry::Struct
  attribute :block_step, BlockStep.optional
  attribute :string,     Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = BlockStep.from_dynamic!(d)
      return new(block_step: value, string: nil) if schema[:block_step].right.valid? value
    rescue StandardError
    end
    return new(string: d, block_step: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !block_step.nil?
      block_step.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Properties of the build that will be created when the step is triggered
class Build < Dry::Struct
  # The branch for the build
  attribute :branch, Types::String.optional

  # The commit hash for the build
  attribute :commit, Types::String.optional

  attribute :env, Types::Hash.meta(of: Types::Any).optional

  # The message for the build (supports emoji)
  attribute :message, Types::String.optional

  # Meta-data for the build
  attribute :meta_data, Types::Hash.meta(of: Types::Any).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      branch: d["branch"],
      commit: d["commit"],
      env: Types::Hash.optional[d["env"]]&.transform_values { |v| Types::Any[v] },
      message: d["message"],
      meta_data: Types::Hash.optional[d["meta_data"]]&.transform_values { |v| Types::Any[v] }
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "branch" => branch,
      "commit" => commit,
      "env" => env,
      "message" => message,
      "meta_data" => meta_data
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class CacheClass < Dry::Struct
  attribute :cache_name, Types::String.optional
  attribute :paths,      Types.Array(Types::String)
  attribute :size,       Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      cache_name: d["name"],
      paths: d.fetch("paths"),
      size: d["size"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "name" => cache_name,
      "paths" => paths,
      "size" => size
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The paths for the caches to be used in the step
class Cache < Dry::Struct
  attribute :cache_class,  CacheClass.optional
  attribute :string,       Types::String.optional
  attribute :string_array, Types.Array(Types::String).optional

  def self.from_dynamic!(d)
    begin
      value = CacheClass.from_dynamic!(d)
      return new(cache_class: value, string_array: nil, string: nil) if schema[:cache_class].right.valid? value
    rescue StandardError
    end
    return new(string: d, string_array: nil, cache_class: nil) if schema[:string].right.valid? d
    return new(string_array: d, cache_class: nil, string: nil) if schema[:string_array].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !cache_class.nil?
      cache_class.to_dynamic
    elsif !string.nil?
      string
    elsif !string_array.nil?
      string_array
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module ExitStatusEnum
  Empty = "*".freeze
end

# The exit status number that will cause this job to retry
class AutomaticRetryExitStatus < Dry::Struct
  attribute :enum,          Types::ExitStatusEnum.optional
  attribute :integer,       Types::Integer.optional
  attribute :integer_array, Types.Array(Types::Integer).optional

  def self.from_dynamic!(d)
    return new(enum: d, integer_array: nil, integer: nil) if schema[:enum].right.valid? d
    return new(integer: d, integer_array: nil, enum: nil) if schema[:integer].right.valid? d
    return new(integer_array: d, enum: nil, integer: nil) if schema[:integer_array].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !enum.nil?
      enum
    elsif !integer.nil?
      integer
    elsif !integer_array.nil?
      integer_array
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The exit signal reason, if any, that may be retried
module SignalReason
  AgentRefused      = "agent_refused".freeze
  AgentStop         = "agent_stop".freeze
  Cancel            = "cancel".freeze
  Empty             = "*".freeze
  None              = "none".freeze
  ProcessRunError   = "process_run_error".freeze
  SignatureRejected = "signature_rejected".freeze
end

class AutomaticRetry < Dry::Struct
  # The exit status number that will cause this job to retry
  attribute :exit_status, Types.Instance(AutomaticRetryExitStatus).optional

  # The number of times this job can be retried
  attribute :limit, Types::Integer.optional

  # The exit signal, if any, that may be retried
  attribute :signal, Types::String.optional

  # The exit signal reason, if any, that may be retried
  attribute :signal_reason, Types::SignalReason.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      exit_status: d["exit_status"] ? AutomaticRetryExitStatus.from_dynamic!(d["exit_status"]) : nil,
      limit: d["limit"],
      signal: d["signal"],
      signal_reason: d["signal_reason"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "exit_status" => exit_status&.to_dynamic,
      "limit" => limit,
      "signal" => signal,
      "signal_reason" => signal_reason
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Whether to allow a job to retry automatically. If set to true, the retry conditions are
# set to the default value.
class Automatic < Dry::Struct
  attribute :automatic_retry,       AutomaticRetry.optional
  attribute :automatic_retry_array, Types.Array(AutomaticRetry).optional
  attribute :bool,                  Types::Bool.optional
  attribute :enum,                  Types::AllowDependencyFailureEnum.optional

  def self.from_dynamic!(d)
    begin
      value = AutomaticRetry.from_dynamic!(d)
      if schema[:automatic_retry].right.valid? value
        return new(automatic_retry: value, automatic_retry_array: nil, bool: nil, enum: nil)
      end
    rescue StandardError
    end
    begin
      value = d.map { |x| AutomaticRetry.from_dynamic!(x) }
      if schema[:automatic_retry_array].right.valid? value
        return new(automatic_retry_array: value, bool: nil, automatic_retry: nil, enum: nil)
      end
    rescue StandardError
    end
    return new(bool: d, automatic_retry_array: nil, automatic_retry: nil, enum: nil) if schema[:bool].right.valid? d
    return new(enum: d, automatic_retry_array: nil, bool: nil, automatic_retry: nil) if schema[:enum].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !automatic_retry.nil?
      automatic_retry.to_dynamic
    elsif !automatic_retry_array.nil?
      automatic_retry_array.map(&:to_dynamic)
    elsif !bool.nil?
      bool
    elsif !enum.nil?
      enum
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class ManualClass < Dry::Struct
  # Whether or not this job can be retried manually
  attribute :allowed, Types.Instance(AllowDependencyFailureUnion).optional

  # Whether or not this job can be retried after it has passed
  attribute :permit_on_passed, Types.Instance(AllowDependencyFailureUnion).optional

  # A string that will be displayed in a tooltip on the Retry button in Buildkite. This will
  # only be displayed if the allowed attribute is set to false.
  attribute :reason, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allowed: d["allowed"] ? AllowDependencyFailureUnion.from_dynamic!(d["allowed"]) : nil,
      permit_on_passed: d["permit_on_passed"] ? AllowDependencyFailureUnion.from_dynamic!(d["permit_on_passed"]) : nil,
      reason: d["reason"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allowed" => allowed&.to_dynamic,
      "permit_on_passed" => permit_on_passed&.to_dynamic,
      "reason" => reason
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Whether to allow a job to be retried manually
class ManualUnion < Dry::Struct
  attribute :bool,         Types::Bool.optional
  attribute :enum,         Types::AllowDependencyFailureEnum.optional
  attribute :manual_class, ManualClass.optional

  def self.from_dynamic!(d)
    return new(bool: d, manual_class: nil, enum: nil) if schema[:bool].right.valid? d
    return new(enum: d, bool: nil, manual_class: nil) if schema[:enum].right.valid? d

    begin
      value = ManualClass.from_dynamic!(d)
      return new(manual_class: value, bool: nil, enum: nil) if schema[:manual_class].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !bool.nil?
      bool
    elsif !enum.nil?
      enum
    elsif !manual_class.nil?
      manual_class.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The conditions for retrying this step.
class Retry < Dry::Struct
  # Whether to allow a job to retry automatically. If set to true, the retry conditions are
  # set to the default value.
  attribute :automatic, Types.Instance(Automatic).optional

  # Whether to allow a job to be retried manually
  attribute :manual, Types.Instance(ManualUnion).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      automatic: d["automatic"] ? Automatic.from_dynamic!(d["automatic"]) : nil,
      manual: d["manual"] ? ManualUnion.from_dynamic!(d["manual"]) : nil
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "automatic" => automatic&.to_dynamic,
      "manual" => manual&.to_dynamic
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module ScriptType
  Command  = "command".freeze
  Commands = "commands".freeze
  Script   = "script".freeze
end

# Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
# this attribute, you must also define concurrency_group and concurrency.
module ConcurrencyMethod
  Eager   = "eager".freeze
  Ordered = "ordered".freeze
end

# List of elements for simple single-dimension Build Matrix
#
# List of existing or new elements for single-dimension Build Matrix
#
# List of elements for single-dimension Build Matrix
#
# List of elements for this Build Matrix dimension
class MatrixElement < Dry::Struct
  attribute :bool,    Types::Bool.optional
  attribute :integer, Types::Integer.optional
  attribute :string,  Types::String.optional

  def self.from_dynamic!(d)
    return new(bool: d, integer: nil, string: nil) if schema[:bool].right.valid? d
    return new(integer: d, bool: nil, string: nil) if schema[:integer].right.valid? d
    return new(string: d, bool: nil, integer: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !bool.nil?
      bool
    elsif !integer.nil?
      integer
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class With < Dry::Struct
  attribute :string_map,  Types::Hash.meta(of: Types::String).optional
  attribute :union_array, Types.Array(Types.Instance(MatrixElement)).optional

  def self.from_dynamic!(d)
    begin
      value = Types::Hash[d].transform_values { |v| Types::String[v] }
      return new(string_map: value, union_array: nil) if schema[:string_map].right.valid? value
    rescue StandardError
    end
    begin
      value = d.map { |x| MatrixElement.from_dynamic!(x) }
      return new(union_array: value, string_map: nil) if schema[:union_array].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !string_map.nil?
      string_map
    elsif !union_array.nil?
      union_array.map(&:to_dynamic)
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Whether this step should be skipped. Passing a string provides a reason for skipping this
# command
class Skip < Dry::Struct
  attribute :bool,   Types::Bool.optional
  attribute :string, Types::String.optional

  def self.from_dynamic!(d)
    return new(bool: d, string: nil) if schema[:bool].right.valid? d
    return new(string: d, bool: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !bool.nil?
      bool
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The exit status number that will cause this job to soft-fail
class SoftFailExitStatus < Dry::Struct
  attribute :enum,    Types::ExitStatusEnum.optional
  attribute :integer, Types::Integer.optional

  def self.from_dynamic!(d)
    return new(enum: d, integer: nil) if schema[:enum].right.valid? d
    return new(integer: d, enum: nil) if schema[:integer].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !enum.nil?
      enum
    elsif !integer.nil?
      integer
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class SoftFailElement < Dry::Struct
  # The exit status number that will cause this job to soft-fail
  attribute :exit_status, Types.Instance(SoftFailExitStatus).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      exit_status: d["exit_status"] ? SoftFailExitStatus.from_dynamic!(d["exit_status"]) : nil
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "exit_status" => exit_status&.to_dynamic
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The conditions for marking the step as a soft-fail.
class SoftFail < Dry::Struct
  attribute :bool,                    Types::Bool.optional
  attribute :enum,                    Types::AllowDependencyFailureEnum.optional
  attribute :soft_fail_element_array, Types.Array(SoftFailElement).optional

  def self.from_dynamic!(d)
    return new(bool: d, soft_fail_element_array: nil, enum: nil) if schema[:bool].right.valid? d
    return new(enum: d, soft_fail_element_array: nil, bool: nil) if schema[:enum].right.valid? d

    begin
      value = d.map { |x| SoftFailElement.from_dynamic!(x) }
      if schema[:soft_fail_element_array].right.valid? value
        return new(soft_fail_element_array: value, bool: nil, enum: nil)
      end
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !bool.nil?
      bool
    elsif !enum.nil?
      enum
    elsif !soft_fail_element_array.nil?
      soft_fail_element_array.map(&:to_dynamic)
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# An adjustment to a Build Matrix
class Adjustment < Dry::Struct
  attribute :skip,            Types.Instance(Skip).optional
  attribute :soft_fail,       Types.Instance(SoftFail).optional
  attribute :adjustment_with, Types.Instance(With)

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      skip: d["skip"] ? Skip.from_dynamic!(d["skip"]) : nil,
      soft_fail: d["soft_fail"] ? SoftFail.from_dynamic!(d["soft_fail"]) : nil,
      adjustment_with: With.from_dynamic!(d.fetch("with"))
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "skip" => skip&.to_dynamic,
      "soft_fail" => soft_fail&.to_dynamic,
      "with" => adjustment_with.to_dynamic
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Setup < Dry::Struct
  attribute :union_array,     Types.Array(Types.Instance(MatrixElement)).optional
  attribute :union_array_map, Types::Hash.meta(of: Types.Array(Types.Instance(MatrixElement))).optional

  def self.from_dynamic!(d)
    begin
      value = d.map { |x| MatrixElement.from_dynamic!(x) }
      return new(union_array: value, union_array_map: nil) if schema[:union_array].right.valid? value
    rescue StandardError
    end
    begin
      value = Types::Hash[d].transform_values { |v| v.map { |x| MatrixElement.from_dynamic!(x) } }
      return new(union_array_map: value, union_array: nil) if schema[:union_array_map].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !union_array.nil?
      union_array.map(&:to_dynamic)
    elsif !union_array_map.nil?
      union_array_map.transform_values { |v| v.map(&:to_dynamic) }
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Configuration for multi-dimension Build Matrix
class MatrixClass < Dry::Struct
  # List of Build Matrix adjustments
  attribute :adjustments, Types.Array(Adjustment).optional

  attribute :setup, Types.Instance(Setup)

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      adjustments: d["adjustments"]&.map { |x| Adjustment.from_dynamic!(x) },
      setup: Setup.from_dynamic!(d.fetch("setup"))
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "adjustments" => adjustments&.map(&:to_dynamic),
      "setup" => setup.to_dynamic
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class MatrixUnion < Dry::Struct
  attribute :matrix_class, MatrixClass.optional
  attribute :union_array,  Types.Array(Types.Instance(MatrixElement)).optional

  def self.from_dynamic!(d)
    begin
      value = MatrixClass.from_dynamic!(d)
      return new(matrix_class: value, union_array: nil) if schema[:matrix_class].right.valid? value
    rescue StandardError
    end
    begin
      value = d.map { |x| MatrixElement.from_dynamic!(x) }
      return new(union_array: value, matrix_class: nil) if schema[:union_array].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !matrix_class.nil?
      matrix_class.to_dynamic
    elsif !union_array.nil?
      union_array.map(&:to_dynamic)
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class GithubCommitStatus2 < Dry::Struct
  # GitHub commit status name
  attribute :context, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      context: d["context"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "context" => context
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack2 < Dry::Struct
  attribute :channels, Types.Array(Types::String).optional
  attribute :message,  Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      channels: d["channels"],
      message: d["message"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "channels" => channels,
      "message" => message
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack5 < Dry::Struct
  attribute :slack2, Slack2.optional
  attribute :string, Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = Slack2.from_dynamic!(d)
      return new(slack2: value, string: nil) if schema[:slack2].right.valid? value
    rescue StandardError
    end
    return new(string: d, slack2: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !slack2.nil?
      slack2.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class NotifyClass < Dry::Struct
  attribute :basecamp_campfire,    Types::String.optional
  attribute :notify_if,            Types::String.optional
  attribute :slack,                Types.Instance(Slack5).optional
  attribute :github_commit_status, GithubCommitStatus2.optional
  attribute :github_check,         Types::Hash.meta(of: Types::Any).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      basecamp_campfire: d["basecamp_campfire"],
      notify_if: d["if"],
      slack: d["slack"] ? Slack5.from_dynamic!(d["slack"]) : nil,
      github_commit_status: d["github_commit_status"] ? GithubCommitStatus2.from_dynamic!(d["github_commit_status"]) : nil,
      github_check: Types::Hash.optional[d["github_check"]]&.transform_values { |v| Types::Any[v] }
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "basecamp_campfire" => basecamp_campfire,
      "if" => notify_if,
      "slack" => slack&.to_dynamic,
      "github_commit_status" => github_commit_status&.to_dynamic,
      "github_check" => github_check
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class NotifyElement < Dry::Struct
  attribute :enum,         Types::NotifyEnum.optional
  attribute :notify_class, NotifyClass.optional

  def self.from_dynamic!(d)
    return new(enum: d, notify_class: nil) if schema[:enum].right.valid? d

    begin
      value = NotifyClass.from_dynamic!(d)
      return new(notify_class: value, enum: nil) if schema[:notify_class].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !enum.nil?
      enum
    elsif !notify_class.nil?
      notify_class.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Array of plugins for this step
class Plugin < Dry::Struct
  attribute :anything_map, Types::Hash.meta(of: Types::Any).optional
  attribute :string,       Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = Types::Hash[d].transform_values { |v| Types::Any[v] }
      return new(anything_map: value, string: nil) if schema[:anything_map].right.valid? value
    rescue StandardError
    end
    return new(string: d, anything_map: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !anything_map.nil?
      anything_map
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Plugins < Dry::Struct
  attribute :anything_map, Types::Hash.meta(of: Types::Any).optional
  attribute :union_array,  Types.Array(Types.Instance(Plugin)).optional

  def self.from_dynamic!(d)
    begin
      value = Types::Hash[d].transform_values { |v| Types::Any[v] }
      return new(anything_map: value, union_array: nil) if schema[:anything_map].right.valid? value
    rescue StandardError
    end
    begin
      value = d.map { |x| Plugin.from_dynamic!(x) }
      return new(union_array: value, anything_map: nil) if schema[:union_array].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !anything_map.nil?
      anything_map
    elsif !union_array.nil?
      union_array.map(&:to_dynamic)
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# The signature of the command step, generally injected by agents at pipeline upload
class Signature < Dry::Struct
  # The algorithm used to generate the signature
  attribute :algorithm, Types::String.optional

  # The fields that were signed to form the signature value
  attribute :signed_fields, Types.Array(Types::String).optional

  # The signature value, a JWS compact signature with a detached body
  attribute :value, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      algorithm: d["algorithm"],
      signed_fields: d["signed_fields"],
      value: d["value"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "algorithm" => algorithm,
      "signed_fields" => signed_fields,
      "value" => value
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class CommandStep < Dry::Struct
  attribute :agents,                   Types.Instance(Agents).optional
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # The glob path/s of artifacts to upload once this step has finished running
  attribute :artifact_paths, Types.Instance(Branches).optional

  attribute :branches,                Types.Instance(Branches).optional
  attribute :cache,                   Types.Instance(Cache).optional
  attribute :cancel_on_build_failing, Types.Instance(AllowDependencyFailureUnion).optional

  # The commands to run on the agent
  attribute :command, Types.Instance(Branches).optional

  # The commands to run on the agent
  attribute :commands, Types.Instance(Branches).optional

  # The maximum number of jobs created from this step that are allowed to run at the same
  # time. If you use this attribute, you must also define concurrency_group.
  attribute :concurrency, Types::Integer.optional

  # A unique name for the concurrency group that you are creating with the concurrency
  # attribute
  attribute :concurrency_group, Types::String.optional

  # Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
  # this attribute, you must also define concurrency_group and concurrency.
  attribute :concurrency_method, Types::ConcurrencyMethod.optional

  attribute :depends_on,        Types.Instance(DependsOn).optional
  attribute :env,               Types::Hash.meta(of: Types::Any).optional
  attribute :id,                Types::String.optional
  attribute :identifier,        Types::String.optional
  attribute :command_step_if,   Types::String.optional
  attribute :key,               Types::String.optional
  attribute :label,             Types::String.optional
  attribute :matrix,            Types.Instance(MatrixUnion).optional
  attribute :command_step_name, Types::String.optional

  # Array of notification options for this step
  attribute :notify, Types.Array(Types.Instance(NotifyElement)).optional

  # The number of parallel jobs that will be created based on this step
  attribute :parallelism, Types::Integer.optional

  attribute :plugins, Types.Instance(Plugins).optional

  # Priority of the job, higher priorities are assigned to agents
  attribute :priority, Types::Integer.optional

  # The conditions for retrying this step.
  attribute :command_step_retry, Retry.optional

  # The signature of the command step, generally injected by agents at pipeline upload
  attribute :signature, Signature.optional

  attribute :skip,      Types.Instance(Skip).optional
  attribute :soft_fail, Types.Instance(SoftFail).optional

  # The number of minutes to time out a job
  attribute :timeout_in_minutes, Types::Integer.optional

  attribute :command_step_type, Types::ScriptType.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      agents: d["agents"] ? Agents.from_dynamic!(d["agents"]) : nil,
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      artifact_paths: d["artifact_paths"] ? Branches.from_dynamic!(d["artifact_paths"]) : nil,
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      cache: d["cache"] ? Cache.from_dynamic!(d["cache"]) : nil,
      cancel_on_build_failing: d["cancel_on_build_failing"] ? AllowDependencyFailureUnion.from_dynamic!(d["cancel_on_build_failing"]) : nil,
      command: d["command"] ? Branches.from_dynamic!(d["command"]) : nil,
      commands: d["commands"] ? Branches.from_dynamic!(d["commands"]) : nil,
      concurrency: d["concurrency"],
      concurrency_group: d["concurrency_group"],
      concurrency_method: d["concurrency_method"],
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      env: Types::Hash.optional[d["env"]]&.transform_values { |v| Types::Any[v] },
      id: d["id"],
      identifier: d["identifier"],
      command_step_if: d["if"],
      key: d["key"],
      label: d["label"],
      matrix: d["matrix"] ? MatrixUnion.from_dynamic!(d["matrix"]) : nil,
      command_step_name: d["name"],
      notify: d["notify"]&.map { |x| NotifyElement.from_dynamic!(x) },
      parallelism: d["parallelism"],
      plugins: d["plugins"] ? Plugins.from_dynamic!(d["plugins"]) : nil,
      priority: d["priority"],
      command_step_retry: d["retry"] ? Retry.from_dynamic!(d["retry"]) : nil,
      signature: d["signature"] ? Signature.from_dynamic!(d["signature"]) : nil,
      skip: d["skip"] ? Skip.from_dynamic!(d["skip"]) : nil,
      soft_fail: d["soft_fail"] ? SoftFail.from_dynamic!(d["soft_fail"]) : nil,
      timeout_in_minutes: d["timeout_in_minutes"],
      command_step_type: d["type"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "agents" => agents&.to_dynamic,
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "artifact_paths" => artifact_paths&.to_dynamic,
      "branches" => branches&.to_dynamic,
      "cache" => cache&.to_dynamic,
      "cancel_on_build_failing" => cancel_on_build_failing&.to_dynamic,
      "command" => command&.to_dynamic,
      "commands" => commands&.to_dynamic,
      "concurrency" => concurrency,
      "concurrency_group" => concurrency_group,
      "concurrency_method" => concurrency_method,
      "depends_on" => depends_on&.to_dynamic,
      "env" => env,
      "id" => id,
      "identifier" => identifier,
      "if" => command_step_if,
      "key" => key,
      "label" => label,
      "matrix" => matrix&.to_dynamic,
      "name" => command_step_name,
      "notify" => notify&.map(&:to_dynamic),
      "parallelism" => parallelism,
      "plugins" => plugins&.to_dynamic,
      "priority" => priority,
      "retry" => command_step_retry&.to_dynamic,
      "signature" => signature&.to_dynamic,
      "skip" => skip&.to_dynamic,
      "soft_fail" => soft_fail&.to_dynamic,
      "timeout_in_minutes" => timeout_in_minutes,
      "type" => command_step_type
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class CommandUnion < Dry::Struct
  attribute :command_step, CommandStep.optional
  attribute :string,       Types::String.optional
  attribute :string_array, Types.Array(Types::String).optional

  def self.from_dynamic!(d)
    begin
      value = CommandStep.from_dynamic!(d)
      return new(command_step: value, string_array: nil, string: nil) if schema[:command_step].right.valid? value
    rescue StandardError
    end
    return new(string: d, string_array: nil, command_step: nil) if schema[:string].right.valid? d
    return new(string_array: d, command_step: nil, string: nil) if schema[:string_array].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !command_step.nil?
      command_step.to_dynamic
    elsif !string.nil?
      string
    elsif !string_array.nil?
      string_array
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module InputType
  Input = "input".freeze
end

class InputStep < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional
  attribute :branches,                 Types.Instance(Branches).optional
  attribute :depends_on,               Types.Instance(DependsOn).optional
  attribute :fields,                   Types.Array(Field).optional
  attribute :id,                       Types::String.optional
  attribute :identifier,               Types::String.optional
  attribute :input_step_if,            Types::String.optional

  # The label of the input step
  attribute :input, Types::String.optional

  attribute :key,             Types::String.optional
  attribute :label,           Types::String.optional
  attribute :input_step_name, Types::String.optional
  attribute :prompt,          Types::String.optional
  attribute :input_step_type, Types::InputType.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      fields: d["fields"]&.map { |x| Field.from_dynamic!(x) },
      id: d["id"],
      identifier: d["identifier"],
      input_step_if: d["if"],
      input: d["input"],
      key: d["key"],
      label: d["label"],
      input_step_name: d["name"],
      prompt: d["prompt"],
      input_step_type: d["type"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "branches" => branches&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "fields" => fields&.map(&:to_dynamic),
      "id" => id,
      "identifier" => identifier,
      "if" => input_step_if,
      "input" => input,
      "key" => key,
      "label" => label,
      "name" => input_step_name,
      "prompt" => prompt,
      "type" => input_step_type
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Input < Dry::Struct
  attribute :input_step, InputStep.optional
  attribute :string,     Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = InputStep.from_dynamic!(d)
      return new(input_step: value, string: nil) if schema[:input_step].right.valid? value
    rescue StandardError
    end
    return new(string: d, input_step: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !input_step.nil?
      input_step.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class GithubCommitStatus3 < Dry::Struct
  # GitHub commit status name
  attribute :context, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      context: d["context"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "context" => context
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack3 < Dry::Struct
  attribute :channels, Types.Array(Types::String).optional
  attribute :message,  Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      channels: d["channels"],
      message: d["message"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "channels" => channels,
      "message" => message
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Slack6 < Dry::Struct
  attribute :slack3, Slack3.optional
  attribute :string, Types::String.optional

  def self.from_dynamic!(d)
    begin
      value = Slack3.from_dynamic!(d)
      return new(slack3: value, string: nil) if schema[:slack3].right.valid? value
    rescue StandardError
    end
    return new(string: d, slack3: nil) if schema[:string].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !slack3.nil?
      slack3.to_dynamic
    elsif !string.nil?
      string
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class BuildNotify2 < Dry::Struct
  attribute :basecamp_campfire,      Types::String.optional
  attribute :build_notify_if,        Types::String.optional
  attribute :slack,                  Types.Instance(Slack6).optional
  attribute :github_commit_status,   GithubCommitStatus3.optional
  attribute :github_check,           Types::Hash.meta(of: Types::Any).optional
  attribute :email,                  Types::String.optional
  attribute :webhook,                Types::String.optional
  attribute :pagerduty_change_event, Types::String.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      basecamp_campfire: d["basecamp_campfire"],
      build_notify_if: d["if"],
      slack: d["slack"] ? Slack6.from_dynamic!(d["slack"]) : nil,
      github_commit_status: d["github_commit_status"] ? GithubCommitStatus3.from_dynamic!(d["github_commit_status"]) : nil,
      github_check: Types::Hash.optional[d["github_check"]]&.transform_values { |v| Types::Any[v] },
      email: d["email"],
      webhook: d["webhook"],
      pagerduty_change_event: d["pagerduty_change_event"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "basecamp_campfire" => basecamp_campfire,
      "if" => build_notify_if,
      "slack" => slack&.to_dynamic,
      "github_commit_status" => github_commit_status&.to_dynamic,
      "github_check" => github_check,
      "email" => email,
      "webhook" => webhook,
      "pagerduty_change_event" => pagerduty_change_event
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class BlockStepNotify < Dry::Struct
  attribute :build_notify2, BuildNotify2.optional
  attribute :enum,          Types::NotifyEnum.optional

  def self.from_dynamic!(d)
    begin
      value = BuildNotify2.from_dynamic!(d)
      return new(build_notify2: value, enum: nil) if schema[:build_notify2].right.valid? value
    rescue StandardError
    end
    return new(enum: d, build_notify2: nil) if schema[:enum].right.valid? d

    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !build_notify2.nil?
      build_notify2.to_dynamic
    elsif !enum.nil?
      enum
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module BlockStepType
  Block    = "block".freeze
  Command  = "command".freeze
  Commands = "commands".freeze
  Input    = "input".freeze
  Script   = "script".freeze
  Trigger  = "trigger".freeze
  Wait     = "wait".freeze
  Waiter   = "waiter".freeze
end

module TriggerType
  Trigger = "trigger".freeze
end

class TriggerStep < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # Whether to continue the build without waiting for the triggered step to complete
  attribute :async, Types.Instance(AllowDependencyFailureUnion).optional

  attribute :branches, Types.Instance(Branches).optional

  # Properties of the build that will be created when the step is triggered
  attribute :build, Build.optional

  attribute :depends_on,        Types.Instance(DependsOn).optional
  attribute :id,                Types::String.optional
  attribute :identifier,        Types::String.optional
  attribute :trigger_step_if,   Types::String.optional
  attribute :key,               Types::String.optional
  attribute :label,             Types::String.optional
  attribute :trigger_step_name, Types::String.optional
  attribute :skip,              Types.Instance(Skip).optional
  attribute :soft_fail,         Types.Instance(SoftFail).optional

  # The slug of the pipeline to create a build
  attribute :trigger, Types::String

  attribute :trigger_step_type, Types::TriggerType.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      async: d["async"] ? AllowDependencyFailureUnion.from_dynamic!(d["async"]) : nil,
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      build: d["build"] ? Build.from_dynamic!(d["build"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      id: d["id"],
      identifier: d["identifier"],
      trigger_step_if: d["if"],
      key: d["key"],
      label: d["label"],
      trigger_step_name: d["name"],
      skip: d["skip"] ? Skip.from_dynamic!(d["skip"]) : nil,
      soft_fail: d["soft_fail"] ? SoftFail.from_dynamic!(d["soft_fail"]) : nil,
      trigger: d.fetch("trigger"),
      trigger_step_type: d["type"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "async" => async&.to_dynamic,
      "branches" => branches&.to_dynamic,
      "build" => build&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "id" => id,
      "identifier" => identifier,
      "if" => trigger_step_if,
      "key" => key,
      "label" => label,
      "name" => trigger_step_name,
      "skip" => skip&.to_dynamic,
      "soft_fail" => soft_fail&.to_dynamic,
      "trigger" => trigger,
      "type" => trigger_step_type
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Trigger < Dry::Struct
  attribute :string,       Types::String.optional
  attribute :trigger_step, TriggerStep.optional

  def self.from_dynamic!(d)
    return new(string: d, trigger_step: nil) if schema[:string].right.valid? d

    begin
      value = TriggerStep.from_dynamic!(d)
      return new(trigger_step: value, string: nil) if schema[:trigger_step].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !string.nil?
      string
    elsif !trigger_step.nil?
      trigger_step.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

module WaiterType
  Wait   = "wait".freeze
  Waiter = "waiter".freeze
end

# Waits for previous steps to pass before continuing
class WaitStep < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional
  attribute :branches,                 Types.Instance(Branches).optional

  # Continue to the next steps, even if the previous group of steps fail
  attribute :continue_on_failure, Types.Instance(AllowDependencyFailureUnion).optional

  attribute :depends_on,     Types.Instance(DependsOn).optional
  attribute :id,             Types::String.optional
  attribute :identifier,     Types::String.optional
  attribute :wait_step_if,   Types::String.optional
  attribute :key,            Types::String.optional
  attribute :label,          Types::String.optional.optional
  attribute :wait_step_name, Types::String.optional.optional
  attribute :wait_step_type, Types::WaiterType.optional

  # Waits for previous steps to pass before continuing
  attribute :wait, Types::String.optional.optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      continue_on_failure: d["continue_on_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["continue_on_failure"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      id: d["id"],
      identifier: d["identifier"],
      wait_step_if: d["if"],
      key: d["key"],
      label: d["label"],
      wait_step_name: d["name"],
      wait_step_type: d["type"],
      wait: d["wait"]
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "branches" => branches&.to_dynamic,
      "continue_on_failure" => continue_on_failure&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "id" => id,
      "identifier" => identifier,
      "if" => wait_step_if,
      "key" => key,
      "label" => label,
      "name" => wait_step_name,
      "type" => wait_step_type,
      "wait" => wait
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Label < Dry::Struct
  attribute :null,      Types::Nil.optional
  attribute :string,    Types::String.optional
  attribute :wait_step, WaitStep.optional

  def self.from_dynamic!(d)
    return new(null: d, wait_step: nil, string: nil) if schema[:null].right.valid? d
    return new(string: d, wait_step: nil, null: nil) if schema[:string].right.valid? d

    begin
      value = WaitStep.from_dynamic!(d)
      return new(wait_step: value, null: nil, string: nil) if schema[:wait_step].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !string.nil?
      string
    elsif !wait_step.nil?
      wait_step.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Waits for previous steps to pass before continuing
class Step1 < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # The label of the block step
  attribute :block, Types.Instance(Block).optional

  # The state that the build is set to when the build is blocked by this block step
  attribute :blocked_state, Types::BlockedState.optional

  attribute :branches,   Types.Instance(Branches).optional
  attribute :depends_on, Types.Instance(DependsOn).optional
  attribute :fields,     Types.Array(Field).optional
  attribute :id,         Types::String.optional
  attribute :identifier, Types::String.optional
  attribute :step_if,    Types::String.optional
  attribute :key,        Types::String.optional
  attribute :label,      Types::String.optional.optional
  attribute :step_name,  Types::String.optional.optional
  attribute :prompt,     Types::String.optional
  attribute :step_type,  Types::BlockStepType.optional

  # The label of the input step
  attribute :input, Types.Instance(Input).optional

  attribute :agents, Types.Instance(Agents).optional

  # The glob path/s of artifacts to upload once this step has finished running
  attribute :artifact_paths, Types.Instance(Branches).optional

  attribute :cache,                   Types.Instance(Cache).optional
  attribute :cancel_on_build_failing, Types.Instance(AllowDependencyFailureUnion).optional

  # The commands to run on the agent
  attribute :command, Types.Instance(CommandUnion).optional

  # The commands to run on the agent
  attribute :commands, Types.Instance(CommandUnion).optional

  # The maximum number of jobs created from this step that are allowed to run at the same
  # time. If you use this attribute, you must also define concurrency_group.
  attribute :concurrency, Types::Integer.optional

  # A unique name for the concurrency group that you are creating with the concurrency
  # attribute
  attribute :concurrency_group, Types::String.optional

  # Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
  # this attribute, you must also define concurrency_group and concurrency.
  attribute :concurrency_method, Types::ConcurrencyMethod.optional

  attribute :env,    Types::Hash.meta(of: Types::Any).optional
  attribute :matrix, Types.Instance(MatrixUnion).optional

  # Array of notification options for this step
  attribute :notify, Types.Array(Types.Instance(NotifyElement)).optional

  # The number of parallel jobs that will be created based on this step
  attribute :parallelism, Types::Integer.optional

  attribute :plugins, Types.Instance(Plugins).optional

  # Priority of the job, higher priorities are assigned to agents
  attribute :priority, Types::Integer.optional

  # The conditions for retrying this step.
  attribute :step_retry, Retry.optional

  # The signature of the command step, generally injected by agents at pipeline upload
  attribute :signature, Signature.optional

  attribute :skip,      Types.Instance(Skip).optional
  attribute :soft_fail, Types.Instance(SoftFail).optional

  # The number of minutes to time out a job
  attribute :timeout_in_minutes, Types::Integer.optional

  attribute :script, CommandStep.optional

  # Continue to the next steps, even if the previous group of steps fail
  attribute :continue_on_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # Waits for previous steps to pass before continuing
  attribute :wait, Types.Instance(Label).optional

  attribute :waiter, WaitStep.optional

  # Whether to continue the build without waiting for the triggered step to complete
  attribute :async, Types.Instance(AllowDependencyFailureUnion).optional

  # Properties of the build that will be created when the step is triggered
  attribute :build, Build.optional

  # The slug of the pipeline to create a build
  attribute :trigger, Types.Instance(Trigger).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      block: d["block"] ? Block.from_dynamic!(d["block"]) : nil,
      blocked_state: d["blocked_state"],
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      fields: d["fields"]&.map { |x| Field.from_dynamic!(x) },
      id: d["id"],
      identifier: d["identifier"],
      step_if: d["if"],
      key: d["key"],
      label: d["label"],
      step_name: d["name"],
      prompt: d["prompt"],
      step_type: d["type"],
      input: d["input"] ? Input.from_dynamic!(d["input"]) : nil,
      agents: d["agents"] ? Agents.from_dynamic!(d["agents"]) : nil,
      artifact_paths: d["artifact_paths"] ? Branches.from_dynamic!(d["artifact_paths"]) : nil,
      cache: d["cache"] ? Cache.from_dynamic!(d["cache"]) : nil,
      cancel_on_build_failing: d["cancel_on_build_failing"] ? AllowDependencyFailureUnion.from_dynamic!(d["cancel_on_build_failing"]) : nil,
      command: d["command"] ? CommandUnion.from_dynamic!(d["command"]) : nil,
      commands: d["commands"] ? CommandUnion.from_dynamic!(d["commands"]) : nil,
      concurrency: d["concurrency"],
      concurrency_group: d["concurrency_group"],
      concurrency_method: d["concurrency_method"],
      env: Types::Hash.optional[d["env"]]&.transform_values { |v| Types::Any[v] },
      matrix: d["matrix"] ? MatrixUnion.from_dynamic!(d["matrix"]) : nil,
      notify: d["notify"]&.map { |x| NotifyElement.from_dynamic!(x) },
      parallelism: d["parallelism"],
      plugins: d["plugins"] ? Plugins.from_dynamic!(d["plugins"]) : nil,
      priority: d["priority"],
      step_retry: d["retry"] ? Retry.from_dynamic!(d["retry"]) : nil,
      signature: d["signature"] ? Signature.from_dynamic!(d["signature"]) : nil,
      skip: d["skip"] ? Skip.from_dynamic!(d["skip"]) : nil,
      soft_fail: d["soft_fail"] ? SoftFail.from_dynamic!(d["soft_fail"]) : nil,
      timeout_in_minutes: d["timeout_in_minutes"],
      script: d["script"] ? CommandStep.from_dynamic!(d["script"]) : nil,
      continue_on_failure: d["continue_on_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["continue_on_failure"]) : nil,
      wait: d["wait"] ? Label.from_dynamic!(d["wait"]) : nil,
      waiter: d["waiter"] ? WaitStep.from_dynamic!(d["waiter"]) : nil,
      async: d["async"] ? AllowDependencyFailureUnion.from_dynamic!(d["async"]) : nil,
      build: d["build"] ? Build.from_dynamic!(d["build"]) : nil,
      trigger: d["trigger"] ? Trigger.from_dynamic!(d["trigger"]) : nil
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "block" => block&.to_dynamic,
      "blocked_state" => blocked_state,
      "branches" => branches&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "fields" => fields&.map(&:to_dynamic),
      "id" => id,
      "identifier" => identifier,
      "if" => step_if,
      "key" => key,
      "label" => label,
      "name" => step_name,
      "prompt" => prompt,
      "type" => step_type,
      "input" => input&.to_dynamic,
      "agents" => agents&.to_dynamic,
      "artifact_paths" => artifact_paths&.to_dynamic,
      "cache" => cache&.to_dynamic,
      "cancel_on_build_failing" => cancel_on_build_failing&.to_dynamic,
      "command" => command&.to_dynamic,
      "commands" => commands&.to_dynamic,
      "concurrency" => concurrency,
      "concurrency_group" => concurrency_group,
      "concurrency_method" => concurrency_method,
      "env" => env,
      "matrix" => matrix&.to_dynamic,
      "notify" => notify&.map(&:to_dynamic),
      "parallelism" => parallelism,
      "plugins" => plugins&.to_dynamic,
      "priority" => priority,
      "retry" => step_retry&.to_dynamic,
      "signature" => signature&.to_dynamic,
      "skip" => skip&.to_dynamic,
      "soft_fail" => soft_fail&.to_dynamic,
      "timeout_in_minutes" => timeout_in_minutes,
      "script" => script&.to_dynamic,
      "continue_on_failure" => continue_on_failure&.to_dynamic,
      "wait" => wait&.to_dynamic,
      "waiter" => waiter&.to_dynamic,
      "async" => async&.to_dynamic,
      "build" => build&.to_dynamic,
      "trigger" => trigger&.to_dynamic
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Pauses the execution of a build and waits on a user to unblock it
#
# Waits for previous steps to pass before continuing
module StringStep
  Block  = "block".freeze
  Input  = "input".freeze
  Wait   = "wait".freeze
  Waiter = "waiter".freeze
end

class BlockStepStep < Dry::Struct
  attribute :enum,  Types::StringStep.optional
  attribute :step1, Step1.optional

  def self.from_dynamic!(d)
    return new(enum: d, step1: nil) if schema[:enum].right.valid? d

    begin
      value = Step1.from_dynamic!(d)
      return new(step1: value, enum: nil) if schema[:step1].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !enum.nil?
      enum
    elsif !step1.nil?
      step1.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

# Waits for previous steps to pass before continuing
class GroupStepClass < Dry::Struct
  attribute :allow_dependency_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # The label of the block step
  attribute :block, Types.Instance(Block).optional

  # The state that the build is set to when the build is blocked by this block step
  attribute :blocked_state, Types::BlockedState.optional

  attribute :branches,   Types.Instance(Branches).optional
  attribute :depends_on, Types.Instance(DependsOn).optional
  attribute :fields,     Types.Array(Field).optional
  attribute :id,         Types::String.optional
  attribute :identifier, Types::String.optional
  attribute :step_if,    Types::String.optional
  attribute :key,        Types::String.optional
  attribute :label,      Types::String.optional.optional
  attribute :step_name,  Types::String.optional.optional
  attribute :prompt,     Types::String.optional
  attribute :step_type,  Types::BlockStepType.optional

  # The label of the input step
  attribute :input, Types.Instance(Input).optional

  attribute :agents, Types.Instance(Agents).optional

  # The glob path/s of artifacts to upload once this step has finished running
  attribute :artifact_paths, Types.Instance(Branches).optional

  attribute :cache,                   Types.Instance(Cache).optional
  attribute :cancel_on_build_failing, Types.Instance(AllowDependencyFailureUnion).optional

  # The commands to run on the agent
  attribute :command, Types.Instance(CommandUnion).optional

  # The commands to run on the agent
  attribute :commands, Types.Instance(CommandUnion).optional

  # The maximum number of jobs created from this step that are allowed to run at the same
  # time. If you use this attribute, you must also define concurrency_group.
  attribute :concurrency, Types::Integer.optional

  # A unique name for the concurrency group that you are creating with the concurrency
  # attribute
  attribute :concurrency_group, Types::String.optional

  # Control command order, allowed values are 'ordered' (default) and 'eager'.  If you use
  # this attribute, you must also define concurrency_group and concurrency.
  attribute :concurrency_method, Types::ConcurrencyMethod.optional

  attribute :env,    Types::Hash.meta(of: Types::Any).optional
  attribute :matrix, Types.Instance(MatrixUnion).optional

  # Array of notification options for this step
  attribute :notify, Types.Array(Types.Instance(BlockStepNotify)).optional

  # The number of parallel jobs that will be created based on this step
  attribute :parallelism, Types::Integer.optional

  attribute :plugins, Types.Instance(Plugins).optional

  # Priority of the job, higher priorities are assigned to agents
  attribute :priority, Types::Integer.optional

  # The conditions for retrying this step.
  attribute :step_retry, Retry.optional

  # The signature of the command step, generally injected by agents at pipeline upload
  attribute :signature, Signature.optional

  attribute :skip,      Types.Instance(Skip).optional
  attribute :soft_fail, Types.Instance(SoftFail).optional

  # The number of minutes to time out a job
  attribute :timeout_in_minutes, Types::Integer.optional

  attribute :script, CommandStep.optional

  # Continue to the next steps, even if the previous group of steps fail
  attribute :continue_on_failure, Types.Instance(AllowDependencyFailureUnion).optional

  # Waits for previous steps to pass before continuing
  attribute :wait, Types.Instance(Label).optional

  attribute :waiter, WaitStep.optional

  # Whether to continue the build without waiting for the triggered step to complete
  attribute :async, Types.Instance(AllowDependencyFailureUnion).optional

  # Properties of the build that will be created when the step is triggered
  attribute :build, Build.optional

  # The slug of the pipeline to create a build
  attribute :trigger, Types.Instance(Trigger).optional

  # The name to give to this group of steps
  attribute :group, Types::String.optional.optional

  # A list of steps
  attribute :steps, Types.Array(Types.Instance(BlockStepStep)).optional

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      allow_dependency_failure: d["allow_dependency_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["allow_dependency_failure"]) : nil,
      block: d["block"] ? Block.from_dynamic!(d["block"]) : nil,
      blocked_state: d["blocked_state"],
      branches: d["branches"] ? Branches.from_dynamic!(d["branches"]) : nil,
      depends_on: d["depends_on"] ? DependsOn.from_dynamic!(d["depends_on"]) : nil,
      fields: d["fields"]&.map { |x| Field.from_dynamic!(x) },
      id: d["id"],
      identifier: d["identifier"],
      step_if: d["if"],
      key: d["key"],
      label: d["label"],
      step_name: d["name"],
      prompt: d["prompt"],
      step_type: d["type"],
      input: d["input"] ? Input.from_dynamic!(d["input"]) : nil,
      agents: d["agents"] ? Agents.from_dynamic!(d["agents"]) : nil,
      artifact_paths: d["artifact_paths"] ? Branches.from_dynamic!(d["artifact_paths"]) : nil,
      cache: d["cache"] ? Cache.from_dynamic!(d["cache"]) : nil,
      cancel_on_build_failing: d["cancel_on_build_failing"] ? AllowDependencyFailureUnion.from_dynamic!(d["cancel_on_build_failing"]) : nil,
      command: d["command"] ? CommandUnion.from_dynamic!(d["command"]) : nil,
      commands: d["commands"] ? CommandUnion.from_dynamic!(d["commands"]) : nil,
      concurrency: d["concurrency"],
      concurrency_group: d["concurrency_group"],
      concurrency_method: d["concurrency_method"],
      env: Types::Hash.optional[d["env"]]&.transform_values { |v| Types::Any[v] },
      matrix: d["matrix"] ? MatrixUnion.from_dynamic!(d["matrix"]) : nil,
      notify: d["notify"]&.map { |x| BlockStepNotify.from_dynamic!(x) },
      parallelism: d["parallelism"],
      plugins: d["plugins"] ? Plugins.from_dynamic!(d["plugins"]) : nil,
      priority: d["priority"],
      step_retry: d["retry"] ? Retry.from_dynamic!(d["retry"]) : nil,
      signature: d["signature"] ? Signature.from_dynamic!(d["signature"]) : nil,
      skip: d["skip"] ? Skip.from_dynamic!(d["skip"]) : nil,
      soft_fail: d["soft_fail"] ? SoftFail.from_dynamic!(d["soft_fail"]) : nil,
      timeout_in_minutes: d["timeout_in_minutes"],
      script: d["script"] ? CommandStep.from_dynamic!(d["script"]) : nil,
      continue_on_failure: d["continue_on_failure"] ? AllowDependencyFailureUnion.from_dynamic!(d["continue_on_failure"]) : nil,
      wait: d["wait"] ? Label.from_dynamic!(d["wait"]) : nil,
      waiter: d["waiter"] ? WaitStep.from_dynamic!(d["waiter"]) : nil,
      async: d["async"] ? AllowDependencyFailureUnion.from_dynamic!(d["async"]) : nil,
      build: d["build"] ? Build.from_dynamic!(d["build"]) : nil,
      trigger: d["trigger"] ? Trigger.from_dynamic!(d["trigger"]) : nil,
      group: d["group"],
      steps: d["steps"]&.map { |x| BlockStepStep.from_dynamic!(x) }
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "allow_dependency_failure" => allow_dependency_failure&.to_dynamic,
      "block" => block&.to_dynamic,
      "blocked_state" => blocked_state,
      "branches" => branches&.to_dynamic,
      "depends_on" => depends_on&.to_dynamic,
      "fields" => fields&.map(&:to_dynamic),
      "id" => id,
      "identifier" => identifier,
      "if" => step_if,
      "key" => key,
      "label" => label,
      "name" => step_name,
      "prompt" => prompt,
      "type" => step_type,
      "input" => input&.to_dynamic,
      "agents" => agents&.to_dynamic,
      "artifact_paths" => artifact_paths&.to_dynamic,
      "cache" => cache&.to_dynamic,
      "cancel_on_build_failing" => cancel_on_build_failing&.to_dynamic,
      "command" => command&.to_dynamic,
      "commands" => commands&.to_dynamic,
      "concurrency" => concurrency,
      "concurrency_group" => concurrency_group,
      "concurrency_method" => concurrency_method,
      "env" => env,
      "matrix" => matrix&.to_dynamic,
      "notify" => notify&.map(&:to_dynamic),
      "parallelism" => parallelism,
      "plugins" => plugins&.to_dynamic,
      "priority" => priority,
      "retry" => step_retry&.to_dynamic,
      "signature" => signature&.to_dynamic,
      "skip" => skip&.to_dynamic,
      "soft_fail" => soft_fail&.to_dynamic,
      "timeout_in_minutes" => timeout_in_minutes,
      "script" => script&.to_dynamic,
      "continue_on_failure" => continue_on_failure&.to_dynamic,
      "wait" => wait&.to_dynamic,
      "waiter" => waiter&.to_dynamic,
      "async" => async&.to_dynamic,
      "build" => build&.to_dynamic,
      "trigger" => trigger&.to_dynamic,
      "group" => group,
      "steps" => steps&.map(&:to_dynamic)
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class SchemaStep < Dry::Struct
  attribute :enum,             Types::StringStep.optional
  attribute :group_step_class, GroupStepClass.optional

  def self.from_dynamic!(d)
    return new(enum: d, group_step_class: nil) if schema[:enum].right.valid? d

    begin
      value = GroupStepClass.from_dynamic!(d)
      return new(group_step_class: value, enum: nil) if schema[:group_step_class].right.valid? value
    rescue StandardError
    end
    raise "Invalid union"
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    if !enum.nil?
      enum
    elsif !group_step_class.nil?
      group_step_class.to_dynamic
    end
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end

class Schema < Dry::Struct
  attribute :agents, Types.Instance(Agents).optional
  attribute :env,    Types::Hash.meta(of: Types::Any).optional
  attribute :notify, Types.Array(Types.Instance(SchemaNotify)).optional

  # A list of steps
  attribute :steps, Types.Array(Types.Instance(SchemaStep))

  def self.from_dynamic!(d)
    d = Types::Hash[d]
    new(
      agents: d["agents"] ? Agents.from_dynamic!(d["agents"]) : nil,
      env: Types::Hash.optional[d["env"]]&.transform_values { |v| Types::Any[v] },
      notify: d["notify"]&.map { |x| SchemaNotify.from_dynamic!(x) },
      steps: d.fetch("steps").map { |x| SchemaStep.from_dynamic!(x) }
    )
  end

  def self.from_json!(json)
    from_dynamic!(JSON.parse(json))
  end

  def to_dynamic
    {
      "agents" => agents&.to_dynamic,
      "env" => env,
      "notify" => notify&.map(&:to_dynamic),
      "steps" => steps.map(&:to_dynamic)
    }
  end

  def to_json(options = nil)
    JSON.generate(to_dynamic, options)
  end
end
