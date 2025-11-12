import { createValidator, PipelineStepValidator } from "./utils";

describe("PipelineEnv", () => {
    let validatePipeline: PipelineStepValidator;
    beforeAll(async () => {
        const { step } = await createValidator();
        validatePipeline = step;
    });

    it("Env", () => {
        validatePipeline({
            env: {
                AN_ENV: "a value",
            },
            steps: [],
        });
    });
});
