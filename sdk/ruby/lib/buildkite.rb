require_relative "buildkite/version"
require_relative "environment"
require "json"
require "yaml"

module Buildkite
  class Error < StandardError; end

  # Here is a comment.
  class Pipeline
    def initialize
      @steps = []
      @agents = nil
      @env = nil
      @notify = nil
    end

    def add_agent(key, value)
      @agents = {} if @agents.nil?
      @agents[key] = value
    end

    def add_environment_variable(key, value)
      @env = {} if @env.nil?
      @env[key] = value
    end

    def add_notify(notify)
      @notify = notify
    end

    # Adds a step to the pipeline.
    #
    # @param [Buildkite::CommandStep, Buildkite::BlockStep] step
    #   The step to add, which can be either a CommandStep or a BlockStep.
    # @return [self]
    #   Returns the pipeline itself for chaining.
    #
    # @example Adding a CommandStep
    #   command_step = Buildkite::CommandStep.new(label: "Run tests", commands: ["bundle exec rspec"])
    #   pipeline.add_step(command_step)
    #
    # @example Adding a BlockStep
    #   block_step = Buildkite::BlockStep.new(label: "Manual approval", block: "Deploy to production")
    #   pipeline.add_step(block_step)
    def add_step(step)
      @steps << step
      self
    end

    def build
      pipeline = {
        "steps" => @steps
      }
      pipeline["agents"] = @agents if @agents != nil
      pipeline["env"] = @env if @env != nil
      pipeline["notify"] = @notify if @notify != nil
      return pipeline
    end

    def to_json(*_args)
      JSON.pretty_generate(self.build, indent: "    ")
    end

    def to_yaml
      self.build.to_yaml
    end
  end
end
