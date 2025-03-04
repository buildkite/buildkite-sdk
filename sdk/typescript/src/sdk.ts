import * as yaml from "yaml";
import * as schema from "./schema";
import * as env from "./environment";
export * as environmentVariables from "./environment";

type PipelineStep =
    | schema.CommandStep
    | schema.WaitStep
    | schema.InputStep
    | schema.TriggerStep
    | schema.BlockStep
    | schema.GroupStepClass;

// export class Environment {
//     static get(key: env.Environment) {
//         env.Environment.BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT;
//         return env.Environment[key];
//     }
// }

export class Pipeline {
    private steps: PipelineStep[] = [];

    constructor() {
        this.steps = [];
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
                steps: this.steps,
            },
            null,
            4
        );
    }

    toYAML() {
        return yaml.stringify({
            steps: this.steps,
        });
    }
}
