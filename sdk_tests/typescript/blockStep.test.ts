import { SchemaValidator } from "./schemaValidator";
import { StepBuilder, types } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("BlockStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render a minimal block step", () => {
        const pipeline = new StepBuilder()
            .addBlockStep({
                block: "My block step",
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });

    test("should render a block step", () => {
        const pipeline = new StepBuilder()
            .addBlockStep(StepArgs.createBlockStepArgs());

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
