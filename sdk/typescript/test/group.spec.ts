import { createValidator, PipelineStepValidator } from "./utils";

describe("GroupStep", () => {
    let validatePipeline: PipelineStepValidator;
    beforeAll(async () => {
        const { step } = await createValidator();
        validatePipeline = step;
    });

    const simpleGroupPipeline = {
        group: "group",
        steps: [{ command: "command" }],
    };

    it("Simple", () => {
        validatePipeline(simpleGroupPipeline);
    });

    it("Id", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            id: "id",
        });
    });

    it("Identifier", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            identifier: "identifier",
        });
    });

    it("DependsOn", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            depends_on: "step",
        });
    });

    it("Key", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            key: "key",
        });
    });

    it("Wait", () => {
        validatePipeline({
            group: "group",
            steps: [
                "wait",
                { key: "waiter", type: "wait" },
                { wait: { key: "waiter2", type: "wait" } },
            ],
        });
    });

    it("Input", () => {
        validatePipeline({
            group: "group",
            steps: [
                "input",
                { input: "a label" },
                { key: "input", type: "input" },
                { input: { key: "input2", type: "input" } },
            ],
        });
    });

    it("If", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            if: "build.message !~ /skip tests/",
        });
    });

    it("AllowDependencyFailure", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            allow_dependency_failure: true,
        });
    });

    it("Notify", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            notify: [{ email: "dev@acmeinc.com" }],
        });
    });

    it("IfChanged", () => {
        validatePipeline({
            ...simpleGroupPipeline,
            if_changed: "*.txt",
        });
    });
});
