import { SchemaValidator } from "./schemaValidator";
import { StepBuilder } from "buildkite-pipeline-sdk";
import StepArgs from "./stepArgs";

describe("GroupStep", () => {
    const validator = new SchemaValidator();
    beforeAll(async () => {
        await validator.fetchSchema();
    });

    test("should render a group step", () => {
        const pipeline = new StepBuilder()
            .addGroupStep({
                allow_dependency_failure: false,
                depends_on: ["build"],
                group: "My group",
                if: `branch == "main"`,
                key: "key",
                label: "",
                notify: ["github_commit_status"],
                skip: false,
                steps: [
                    StepArgs.createBlockStepArgs(),
                    StepArgs.createCommandStepArgs(),
                    StepArgs.createInputStepArgs(),
                    StepArgs.createTriggerStepArgs(),
                    StepArgs.createWaitStepArgs(),
                ],
            });

        expect(
            validator.validate(pipeline)
        ).toBeTruthy();
    });
});
