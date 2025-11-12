import { createValidator, PipelineStepValidator } from "./utils";

describe("Matrix", () => {
    let validatePipeline: PipelineStepValidator;
    beforeAll(async () => {
        const { step } = await createValidator();
        validatePipeline = step;
    });

    it("Simple", () => {
        validatePipeline({
            command: "echo {{matrix}}",
            label: "{{matrix}}",
            matrix: ["one", "two"],
        });
    });

    describe("Detailed", () => {
        it("Simple", () => {
            validatePipeline({
                command: "echo {{matrix}}",
                label: "{{matrix}}",
                matrix: {
                    setup: ["one", "two"],
                    adjustments: [{ with: ["three"], skip: true }],
                },
            });
        });

        it("Complex", () => {
            validatePipeline({
                command: "echo {{matrix.color}} {{matrix.shape}}",
                label: "{{matrix.color}} {{matrix.shape}}",
                matrix: {
                    setup: {
                        color: ["green", "blue"],
                        shape: ["triangle", "hexagon"],
                    },
                    adjustments: [
                        {
                            with: { color: "blue", shape: "triangle" },
                            skip: true,
                        },
                        {
                            with: { color: "green", shape: "triangle" },
                            skip: "look, hexagons are just better",
                        },
                        { with: { color: "purple", shape: "hexagon" } },
                    ],
                },
            });
        });
    });
});
