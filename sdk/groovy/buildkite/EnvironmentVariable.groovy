package buildkite

/**
 * Buildkite Environment Variables
 * A Groovy enum equivalent to the TypeScript EnvironmentVariable
 */
enum EnvironmentVariable {
    /**
     * Always `true`
     */
    BUILDKITE("BUILDKITE"),

    /**
     * The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.
     */
    BUILDKITE_AGENT_ACCESS_TOKEN("BUILDKITE_AGENT_ACCESS_TOKEN"),

    /**
     * The value of the `debug` agent configuration option.
     */
    BUILDKITE_AGENT_DEBUG("BUILDKITE_AGENT_DEBUG"),

    /**
     * The value of the `disconnect-after-job` agent configuration option.
     */
    BUILDKITE_AGENT_DISCONNECT_AFTER_JOB("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB"),

    /**
     * The value of the `disconnect-after-idle-timeout` agent configuration option.
     */
    BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT"),

    /**
     * The value of the `endpoint` agent configuration option.
     */
    BUILDKITE_AGENT_ENDPOINT("BUILDKITE_AGENT_ENDPOINT"),

    /**
     * A list of the experimental agent features that are currently enabled.
     */
    BUILDKITE_AGENT_EXPERIMENT("BUILDKITE_AGENT_EXPERIMENT"),

    /**
     * The value of the `health-check-addr` agent configuration option.
     */
    BUILDKITE_AGENT_HEALTH_CHECK_ADDR("BUILDKITE_AGENT_HEALTH_CHECK_ADDR"),

    /**
     * The UUID of the agent.
     */
    BUILDKITE_AGENT_ID("BUILDKITE_AGENT_ID"),

    /**
     * The value of each agent tag.
     */
    BUILDKITE_AGENT_META_DATA_("BUILDKITE_AGENT_META_DATA_"),

    /**
     * The name of the agent that ran the job.
     */
    BUILDKITE_AGENT_NAME("BUILDKITE_AGENT_NAME"),

    /**
     * The process ID of the agent.
     */
    BUILDKITE_AGENT_PID("BUILDKITE_AGENT_PID"),

    /**
     * The artifact paths to upload after the job, if any have been specified.
     */
    BUILDKITE_ARTIFACT_PATHS("BUILDKITE_ARTIFACT_PATHS"),

    /**
     * The path where artifacts will be uploaded.
     */
    BUILDKITE_ARTIFACT_UPLOAD_DESTINATION("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION"),

    /**
     * The path to the directory containing the `buildkite-agent` binary.
     */
    BUILDKITE_BIN_PATH("BUILDKITE_BIN_PATH"),

    /**
     * The branch being built.
     */
    BUILDKITE_BRANCH("BUILDKITE_BRANCH"),

    /**
     * The path where the agent has checked out your code for this build.
     */
    BUILDKITE_BUILD_CHECKOUT_PATH("BUILDKITE_BUILD_CHECKOUT_PATH"),

    /**
     * The name of the user who authored the commit being built.
     */
    BUILDKITE_BUILD_AUTHOR("BUILDKITE_BUILD_AUTHOR"),

    /**
     * The notification email of the user who authored the commit being built.
     */
    BUILDKITE_BUILD_AUTHOR_EMAIL("BUILDKITE_BUILD_AUTHOR_EMAIL"),

    /**
     * The name of the user who created the build.
     */
    BUILDKITE_BUILD_CREATOR("BUILDKITE_BUILD_CREATOR"),

    /**
     * The notification email of the user who created the build.
     */
    BUILDKITE_BUILD_CREATOR_EMAIL("BUILDKITE_BUILD_CREATOR_EMAIL"),

    /**
     * A colon separated list of non-private team slugs that the build creator belongs to.
     */
    BUILDKITE_BUILD_CREATOR_TEAMS("BUILDKITE_BUILD_CREATOR_TEAMS"),

    /**
     * The UUID of the build.
     */
    BUILDKITE_BUILD_ID("BUILDKITE_BUILD_ID"),

    /**
     * The build number.
     */
    BUILDKITE_BUILD_NUMBER("BUILDKITE_BUILD_NUMBER"),

    /**
     * The value of the `build-path` agent configuration option.
     */
    BUILDKITE_BUILD_PATH("BUILDKITE_BUILD_PATH"),

    /**
     * The url for this build on Buildkite.
     */
    BUILDKITE_BUILD_URL("BUILDKITE_BUILD_URL"),

    /**
     * The value of the `cancel-grace-period` agent configuration option in seconds.
     */
    BUILDKITE_CANCEL_GRACE_PERIOD("BUILDKITE_CANCEL_GRACE_PERIOD"),

    /**
     * The value of the `cancel-signal` agent configuration option.
     */
    BUILDKITE_CANCEL_SIGNAL("BUILDKITE_CANCEL_SIGNAL"),

    /**
     * Whether the build should perform a clean checkout.
     */
    BUILDKITE_CLEAN_CHECKOUT("BUILDKITE_CLEAN_CHECKOUT"),

    /**
     * The UUID value of the cluster.
     */
    BUILDKITE_CLUSTER_ID("BUILDKITE_CLUSTER_ID"),

    /**
     * The command that will be run for the job.
     */
    BUILDKITE_COMMAND("BUILDKITE_COMMAND"),

    /**
     * The opposite of the value of the `no-command-eval` agent configuration option.
     */
    BUILDKITE_COMMAND_EVAL("BUILDKITE_COMMAND_EVAL"),

    /**
     * The exit code from the last command run in the command hook.
     */
    BUILDKITE_COMMAND_EXIT_STATUS("BUILDKITE_COMMAND_EXIT_STATUS"),

    /**
     * The git commit object of the build.
     */
    BUILDKITE_COMMIT("BUILDKITE_COMMIT"),

    /**
     * The path to the agent config file.
     */
    BUILDKITE_CONFIG_PATH("BUILDKITE_CONFIG_PATH"),

    /**
     * The path to the file containing the job's environment variables.
     */
    BUILDKITE_ENV_FILE("BUILDKITE_ENV_FILE"),

    // More environment variables would be added here...

    /**
     * Always `true`.
     */
    CI("CI")

    private final String value

    private EnvironmentVariable(String value) {
        this.value = value
    }

    @Override
    String toString() {
        return value
    }
}
