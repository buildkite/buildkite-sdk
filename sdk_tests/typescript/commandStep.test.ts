import { SchemaValidator } from "./schemaValidator";
import { StepBuilder, types } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("CommandStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render a minimal command step", () => {
        const pipeline = new StepBuilder()
            .addCommandStep({
                commands: ["./run.sh"],
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });

    test("should render a command step", () => {
        const pipeline = new StepBuilder()
            .addCommandStep(StepArgs.createCommandStepArgs());

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
