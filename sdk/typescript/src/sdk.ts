import * as yaml from "yaml";
import { PipelineNotify, NotifyEnum } from "./types";
import { BlockStep } from "./blockStep";
import { CommandStep } from "./commandStep";
import { GroupStep } from "./groupStep";
import { InputStep } from "./inputStep";
import { TriggerStep } from "./triggerStep";
import { WaitStep } from "./waitStep";
export { EnvironmentVariable } from "./environment";

export type {
    BlockStep,
    CommandStep,
    GroupStep,
    InputStep,
    TriggerStep,
    WaitStep,
};

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
    notify?: (PipelineNotify | NotifyEnum)[];
    steps?: PipelineStep[];
}

export class Pipeline {
    public agents: Record<string, any> = {};
    public env: Record<string, any> = {};
    public notify: (PipelineNotify | NotifyEnum)[] = [];
    public steps: PipelineStep[] = [];

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
    addNotify(notify: PipelineNotify | NotifyEnum) {
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

    private build(): PipelineSchema {
        const pipeline: PipelineSchema = {};

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
