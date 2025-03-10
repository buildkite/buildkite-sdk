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

export class Pipeline {
    private agents: Record<string, any> = {};
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
     * Add an notification
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

    toJSON() {
        return JSON.stringify(
            {
                agents: this.agents,
                notify: this.notify,
                steps: this.steps,
            },
            null,
            4
        );
    }

    toYAML() {
        return yaml.stringify({
            agents: this.agents,
            notify: this.notify,
            steps: this.steps,
        });
    }
}
