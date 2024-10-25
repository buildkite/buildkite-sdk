import { SchemaValidator } from "./schemaValidator";
import { StepBuilder } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("WaitStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render a minimal wait step", () => {
        const pipeline = new StepBuilder()
            .addWaitStep({
                wait: null,
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });

    test("should render a wait step", () => {
        const pipeline = new StepBuilder()
            .addWaitStep(StepArgs.createWaitStepArgs());

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
