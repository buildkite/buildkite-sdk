import * as buildkite from "../src/index";

describe("toJSON()", () => {
    it("should render pipeline-level priority", () => {
        const pipeline = new buildkite.Pipeline();
        pipeline.setPriority(100);
        pipeline.addStep({ command: "echo 'Hello, world!'" });

        expect(pipeline.toJSON()).toBe(
            JSON.stringify(
                {
                    steps: [{ command: "echo 'Hello, world!'" }],
                    priority: 100,
                },
                null,
                4
            )
        );
    });

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
