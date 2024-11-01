import { SchemaValidator } from "./schemaValidator";
import { StepBuilder, types } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("TriggerStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render a minimal trigger step", () => {
        const pipeline = new StepBuilder()
            .addTriggerStep({
                trigger: "some-pipeline",
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });

    test("should render a trigger step", () => {
        const pipeline = new StepBuilder()
            .addTriggerStep(StepArgs.createTriggerStepArgs());

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
