import { createValidator, PipelineStepValidator } from "./utils";

describe("BlockStep", () => {
    let validatePipeline: PipelineStepValidator;
    beforeAll(async () => {
        const { step } = await createValidator();
        validatePipeline = step;
    });

    describe("Nesting formats", () => {
        it("String", () => {
            validatePipeline("block");
        });

        it("Label", () => {
            validatePipeline({ block: "blockLabel" });
        });

        it("Nested", () => {
            validatePipeline({
                block: { fields: [{ text: "text", key: "key" }] },
            });
        });

        it("Type", () => {
            validatePipeline({
                type: "block",
                label: "label",
                fields: [{ text: "text", key: "key" }],
            });
        });
    });

    describe("All the options", () => {
        it("Simple", () => {
            validatePipeline({ block: "A label" });
        });

        it("Branches", () => {
            validatePipeline({
                block: "A label",
                branches: "main",
            });
        });

        it("Id", () => {
            validatePipeline({
                block: "A label",
                id: "an-id",
            });
        });

        it("Identifier", () => {
            validatePipeline({
                block: "A label",
                identifier: "identifier",
            });
        });

        it("Prompt", () => {
            validatePipeline({
                block: "A label",
                prompt: "A prompt",
            });
        });

        it("Fields", () => {
            validatePipeline({
                block: "A label",
                prompt: "A prompt",
                fields: [
                    {
                        text: "Field 1",
                        key: "field-1",
                    },
                    {
                        text: "Field 2",
                        key: "field-2",
                        required: false,
                        default: "Field 2 Default",
                        hint: "Field 2 Hint",
                    },
                    {
                        select: "Select 1",
                        key: "select-1",
                        options: [
                            {
                                label: "Select 1 Option 1",
                                value: "select-1-option-1",
                            },
                            {
                                label: "Select 1 Option 2",
                                value: "select-1-option-2",
                            },
                        ],
                    },
                    {
                        select: "Select 2",
                        key: "select-2",
                        hint: "Select 2 Hint",
                        required: false,
                        default: "select-2-option-1",
                        options: [
                            {
                                label: "Select 2 Option 1",
                                value: "select-2-option-1",
                            },
                        ],
                    },
                ],
            });
        });

        it("If", () => {
            validatePipeline({
                block: "A label",
                if: "build.message !~ /skip tests/",
            });
        });

        it("Key", () => {
            validatePipeline({
                block: "A label",
                key: "important-step",
            });
        });

        describe("DependsOn", () => {
            it("String", () => {
                validatePipeline({
                    block: "A label",
                    depends_on: "depend-on-me",
                });
            });

            it("StringArray", () => {
                validatePipeline({
                    block: "A label",
                    depends_on: ["depend-on-me-1", "depend-on-me-2"],
                });
            });

            it("Object", () => {
                validatePipeline({
                    block: "A label",
                    depends_on: [{ step: "depend-on-me", allow_failure: true }],
                });
            });

            it("ObjectArray", () => {
                validatePipeline({
                    block: "A label",
                    depends_on: [
                        { step: "depend-on-me-1" },
                        { step: "depend-on-me-2" },
                    ],
                });
            });

            it("Mixed", () => {
                validatePipeline({
                    block: "A label",
                    depends_on: ["depend-on-me-1", { step: "depend-on-me-2" }],
                });
            });
        });

        it("AllowDependencyFailure", () => {
            validatePipeline({
                block: "A label",
                allow_dependency_failure: true,
            });
        });

        it("MultipleSelectFields", () => {
            validatePipeline({
                block: "A label",
                fields: [
                    {
                        select: "Multiple fields",
                        key: "multiple-fields",
                        multiple: true,
                        options: [
                            { label: "Option 1", value: "option-1" },
                            { label: "Option 2", value: "option-2" },
                        ],
                    },
                ],
            });
        });

        it("MultipleSelectFieldsDefaults", () => {
            validatePipeline({
                block: "A label",
                fields: [
                    {
                        select: "Multiple fields with multiple default",
                        key: "multiple-fields",
                        multiple: true,
                        default: ["option-1", "option-2"],
                        options: [
                            { label: "Option 1", value: "option-1" },
                            { label: "Option 2", value: "option-2" },
                        ],
                    },
                ],
            });
        });

        describe("AllowedTeams", () => {
            it("String", () => {
                validatePipeline({
                    block: "A label",
                    allowed_teams: "team-slug",
                });
            });

            it("StringArray", () => {
                validatePipeline({
                    block: "A label",
                    allowed_teams: [
                        "team-slug",
                        "9b30da58-3dac-4dbb-89e9-b9748ca76445",
                    ],
                });
            });
        });

        it("GroupNested", () => {
            validatePipeline({
                group: "Test",
                steps: [{ block: "A label" }],
            });
        });
    });
});
