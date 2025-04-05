package buildkite

/**
 * Defines types for Buildkite steps
 */
class StepTypes {
    /**
     * Represents a command step in a Buildkite pipeline
     */
    static class CommandStep extends LinkedHashMap<String, Object> {
        CommandStep(Map<String, Object> properties) {
            super(properties)
        }
    }

    /**
     * Represents a wait step in a Buildkite pipeline
     */
    static class WaitStep extends LinkedHashMap<String, Object> {
        WaitStep(Map<String, Object> properties = [:]) {
            super(properties)
            this.wait = null
        }
    }

    /**
     * Represents an input step in a Buildkite pipeline
     */
    static class InputStep extends LinkedHashMap<String, Object> {
        InputStep(String label, Map<String, Object> properties = [:]) {
            super(properties)
            this.input = label
        }
    }

    /**
     * Represents a trigger step in a Buildkite pipeline
     */
    static class TriggerStep extends LinkedHashMap<String, Object> {
        TriggerStep(String pipeline, Map<String, Object> properties = [:]) {
            super(properties)
            this.trigger = pipeline
        }
    }

    /**
     * Represents a block step in a Buildkite pipeline
     */
    static class BlockStep extends LinkedHashMap<String, Object> {
        BlockStep(String label, Map<String, Object> properties = [:]) {
            super(properties)
            this.block = label
        }
    }

    /**
     * Represents a group step in a Buildkite pipeline
     */
    static class GroupStep extends LinkedHashMap<String, Object> {
        GroupStep(String label, List<Map<String, Object>> steps, Map<String, Object> properties = [:]) {
            super(properties)
            this.group = label
            this.steps = steps
        }
    }
}
