import * as yaml from "yaml";
import * as schema from "./schema";

export type CommandStep = schema.CommandStep;
export type WaitStep = schema.WaitStep;
export type InputStep = schema.InputStep;
export type TriggerStep = schema.TriggerStep;
export type BlockStep = schema.BlockStep;
export type GroupStep = schema.GroupStepClass;

export type PipelineStep =
    | CommandStep
    | WaitStep
    | InputStep
    | TriggerStep
    | BlockStep
    | GroupStep;

interface PipelineSchema {
    agents?: Record<string, any>;
    env?: Record<string, any>;
    notify?: (schema.PurpleBuildNotify | schema.NotifyEnum)[];
    steps?: PipelineStep[];
}

export class Pipeline {
    private agents: Record<string, any> = {};
    private env: Record<string, any> = {};
    private notify: (schema.PurpleBuildNotify | schema.NotifyEnum)[] = [];
    private steps: PipelineStep[] = [];

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
     * Add an environemnt variable
     * @param key
     * @param value
     */
    addEnvironmentVariable(key: string, value: any) {
        this.env[key] = value
    }

    /**
     * Add an notification
     * @param notify
     */
    addNotify(notify: schema.PurpleBuildNotify | schema.NotifyEnum) {
        this.notify.push(notify);
    }

    /**
     * Add a step to the pipeline.
     * @param step
     * @returns
     */
    addStep(step: PipelineStep) {
        this.steps.push(step);
        return this;
    }

    private createPipeline(): PipelineSchema {
        const pipeline: PipelineSchema = {}

        if (Object.keys(this.agents).length > 0) {
            pipeline.agents = this.agents
        }

        if (Object.keys(this.env).length > 0) {
            pipeline.env = this.env
        }

        if (Object.keys(this.notify).length > 0) {
            pipeline.notify = this.notify
        }

        if (Object.keys(this.steps).length > 0) {
            pipeline.steps = this.steps
        }

        return pipeline
    }

    toJSON() {
        return JSON.stringify(
            this.createPipeline(),
            null,
            4
        );
    }

    toYAML() {
        return yaml.stringify(this.createPipeline());
    }
}
