import * as buildkite from "../src/index";

describe("toJSON()", () => {
    it("should render the pipeline steps", () => {
        const pipeline = new buildkite.Pipeline();

        pipeline.addStep({
            command: "echo 'Hello, world!'",
        });

        expect(pipeline.toJSON()).toBe(
            JSON.stringify(
                {
                    steps: [{ command: "echo 'Hello, world!'" }],
                },
                null,
                4
            )
        );
    });
});
