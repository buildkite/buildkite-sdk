import { SchemaValidator } from "./schemaValidator";
import { StepBuilder, types } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("InputStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render an input step", () => {
        const pipeline = new StepBuilder()
            .addInputStep(StepArgs.createInputStepArgs());

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });

    test("should render a text and select input", () => {
        const pipeline = new StepBuilder()
            .addInputStep({
                input: "My input step",
                fields: [
                    {
                        text: "One",
                        key: "one",
                    },
                    {
                        key: "two",
                        options: [
                            {
                                label: "Item",
                                value: "item",
                            },
                        ],
                    },
                ],
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
