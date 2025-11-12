import * as yaml from "yaml";
import * as schema from './types/schema'
export { EnvironmentVariable } from "./environment";

export class Pipeline {
    public agents: schema.AgentsObject = {};
    public env: schema.Env = {};
    public notify: schema.BuildNotify = [];
    public steps: schema.PipelineSteps = [];
    public secrets: schema.Secrets = [];

    /**
     * Set the pipeline
     * @param pipeline
     * @returns
     */
    setPipeline(pipeline: schema.BuildkitePipeline) {
        if (pipeline.agents) {
            this.agents = pipeline.agents
        }

        if (pipeline.env) {
            this.env = pipeline.env
        }

        if (pipeline.notify) {
            this.notify = pipeline.notify
        }

        if (pipeline.steps) {
            this.steps = pipeline.steps
        }
    }

    /**
     * Set the secrets for the pipeline
     * @param secrets
     * @returns
     */
    setSecrets(secrets: schema.Secrets) {
        this.secrets = secrets
    }

    /**
     * Add an agent to target by tag
     * @param tagName
     * @param tagValue
     * @returns
     */
    addAgent(tagName: string, tagValue: string) {
        this.agents[tagName] = tagValue;
    }

    /**
     * Add an environment variable
     * @param key
     * @param value
     */
    addEnvironmentVariable(key: string, value: any) {
        this.env[key] = value;
    }

    /**
     * Add an notification
     * @param notify
     */
    addNotify(notify: schema.BuildNotify[0]) {
        this.notify.push(notify);
    }

    /**
     * Add a step to the pipeline.
     * @param step
     * @returns
     */
    addStep(step: schema.PipelineSteps[0]) {
        this.steps.push(step);
        return this;
    }

    private build(): schema.BuildkitePipeline {
        const pipeline: schema.BuildkitePipeline = {};

        if (Object.keys(this.secrets).length > 0) {
            pipeline.secrets = this.secrets;
        }

        if (Object.keys(this.agents).length > 0) {
            pipeline.agents = this.agents;
        }

        if (Object.keys(this.env).length > 0) {
            pipeline.env = this.env;
        }

        if (Object.keys(this.notify).length > 0) {
            pipeline.notify = this.notify;
        }

        if (Object.keys(this.steps).length > 0) {
            pipeline.steps = this.steps;
        }

        return pipeline;
    }

    toJSON() {
        return JSON.stringify(this.build(), null, 4);
    }

    toYAML() {
        return yaml.stringify(this.build());
    }
}
