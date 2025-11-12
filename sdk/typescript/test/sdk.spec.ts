import * as buildkite from "../src/index";

describe("toJSON()", () => {
    it("should render the pipeline steps", () => {
        const pipeline = new buildkite.Pipeline();

        pipeline.setSecrets(["MY_SECRET"]);
        pipeline.addAgent("os", "mac");
        pipeline.addEnvironmentVariable("FOO", "BAR");
        pipeline.addNotify({ email: "person@example.com" });

        pipeline.addStep({
            command: "echo 'Hello, world!'",
        });

        expect(pipeline.toJSON()).toBe(
            JSON.stringify(
                {
                    secrets: ["MY_SECRET"],
                    agents: { os: "mac" },
                    env: { FOO: "BAR" },
                    notify: [{ email: "person@example.com" }],
                    steps: [{ command: "echo 'Hello, world!'" }],
                },
                null,
                4
            )
        );
    });
});
