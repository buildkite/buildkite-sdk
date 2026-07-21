import * as buildkite from "../src/index";
import {
    createValidator,
    PipelineStepValidator,
    PipelineSchemaValidator,
} from "./utils";

describe("Checkout", () => {
    let validateStep: PipelineStepValidator;
    let validatePipeline: PipelineSchemaValidator;
    beforeAll(async () => {
        const { step, pipeline } = await createValidator();
        validateStep = step;
        validatePipeline = pipeline;
    });

    describe("CommandStep", () => {
        it("Skip", () => {
            validateStep({ command: "ls", checkout: { skip: false } });
        });

        it("SkipString", () => {
            validateStep({ command: "ls", checkout: { skip: "true" } });
        });

        it("DepthNumber", () => {
            validateStep({ command: "ls", checkout: { depth: 50 } });
        });

        it("DepthString", () => {
            validateStep({ command: "ls", checkout: { depth: "50" } });
        });

        it("Flags", () => {
            validateStep({
                command: "ls",
                checkout: {
                    flags: {
                        clone: "--depth 1 --single-branch",
                        fetch: "--prune --tags",
                    },
                },
            });
        });

        it("SparseSinglePath", () => {
            validateStep({
                command: "ls",
                checkout: { sparse: { paths: "src/" } },
            });
        });

        it("SparsePathList", () => {
            validateStep({
                command: "ls",
                checkout: { sparse: { paths: ["src/", "docs/"] } },
            });
        });

        it("SshSecretAndCommitVerification", () => {
            validateStep({
                command: "ls",
                checkout: {
                    ssh_secret: "deploy_key",
                    commit_verification: "strict",
                },
            });
        });
    });

    describe("Pipeline", () => {
        it("PipelineLevelCheckout", () => {
            validatePipeline({
                checkout: { submodules: false, depth: 50 },
                steps: [{ command: "ls" }],
            });
        });

        it("should render pipeline-level checkout", () => {
            const pipeline = new buildkite.Pipeline();
            pipeline.setCheckout({ submodules: false, depth: 50 });
            pipeline.addStep({ command: "ls" });

            expect(pipeline.toJSON()).toBe(
                JSON.stringify(
                    {
                        checkout: { submodules: false, depth: 50 },
                        steps: [{ command: "ls" }],
                    },
                    null,
                    4
                )
            );
        });

        it("rejects the step-only ssh_secret at pipeline level", () => {
            const pipeline = new buildkite.Pipeline();
            // @ts-expect-error ssh_secret is step-level only
            pipeline.setCheckout({ ssh_secret: "deploy_key" });
        });

        it("should render step-level checkout", () => {
            const pipeline = new buildkite.Pipeline();
            pipeline.addStep({ command: "ls", checkout: { skip: false } });

            expect(pipeline.toJSON()).toBe(
                JSON.stringify(
                    {
                        steps: [{ command: "ls", checkout: { skip: false } }],
                    },
                    null,
                    4
                )
            );
        });
    });
});
