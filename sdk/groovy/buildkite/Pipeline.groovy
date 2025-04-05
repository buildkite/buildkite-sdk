package buildkite

import groovy.json.JsonOutput
import org.yaml.snakeyaml.Yaml

/**
 * Buildkite Pipeline SDK
 * A Groovy implementation of the Buildkite Pipeline SDK
 */
class Pipeline {
    private Map<String, Object> agents = [:]
    private Map<String, Object> env = [:]
    private List<Object> notify = []
    private List<Map<String, Object>> steps = []

    /**
     * Add an agent to target by tag
     * @param tagName The tag name
     * @param tagValue The tag value
     * @return This pipeline instance for chaining
     */
    Pipeline addAgent(String tagName, Object tagValue) {
        agents[tagName] = tagValue
        return this
    }

    /**
     * Add an environment variable
     * @param key The environment variable name
     * @param value The environment variable value
     * @return This pipeline instance for chaining
     */
    Pipeline addEnvironmentVariable(String key, Object value) {
        env[key] = value
        return this
    }

    /**
     * Add a notification
     * @param notification The notification to add
     * @return This pipeline instance for chaining
     */
    Pipeline addNotify(Object notification) {
        notify.add(notification)
        return this
    }

    /**
     * Add a step to the pipeline
     * @param step The step to add
     * @return This pipeline instance for chaining
     */
    Pipeline addStep(Map<String, Object> step) {
        steps.add(step)
        return this
    }

    /**
     * Build the pipeline object
     * @return The pipeline as a Map
     */
    private Map<String, Object> build() {
        Map<String, Object> pipeline = [:]

        if (!agents.isEmpty()) {
            pipeline.agents = agents
        }

        if (!env.isEmpty()) {
            pipeline.env = env
        }

        if (!notify.isEmpty()) {
            pipeline.notify = notify
        }

        if (!steps.isEmpty()) {
            pipeline.steps = steps
        }

        return pipeline
    }

    /**
     * Convert the pipeline to JSON
     * @return The pipeline as a JSON string
     */
    String toJSON() {
        return JsonOutput.prettyPrint(JsonOutput.toJson(build()))
    }

    /**
     * Convert the pipeline to YAML
     * @return The pipeline as a YAML string
     */
    String toYAML() {
        return new Yaml().dump(build())
    }
}
