import { types } from "buildkite-pipeline-sdk";

class StepArgs {
    createBlockStepArgs(): types.Block {
        return {
            block: "My block step",
            allow_dependency_failure: false,
            blocked_state: types.BlockedState.FAILED,
            branches: "branch",
            depends_on: ["build"],
            fields: [
                {
                    key: "one",
                    text: "One",
                },
            ],
            if: `branch == "main"`,
            key: "block",
            prompt: "prompt",
        };
    }

    createCommandStepArgs(): types.Command {
        return {
            agents: {queue: "mac"},
            allow_dependency_failure: false,
            artifact_paths: ["result.json"],
            branches: "branch",
            cancel_on_build_failing: false,
            commands: ["run.sh"],
            concurrency: 1,
            concurrency_group: "build",
            depends_on: ["test"],
            env: {FOO: "BAR"},
            if: `branch == "main"`,
            key: "key",
            label: "label",
            matrix: ["mac"],
            parallelism: 1,
            plugins: [
                {"plugin": {}},
            ],
            priority: 1,
            retry: {
                automatic: true,
            },
            soft_fail: false,
            timeout_in_minutes: 10,
        };
    }

    createInputStepArgs(): types.Input {
        return {
            input: "My input step",
            allow_dependency_failure: false,
            branches: "branch",
            depends_on: ["build"],
            if: `branch == "main"`,
            key: "input",
            prompt: "Tell me somethig",
            fields: [
                {
                    text: "One",
                    key: "one",
                },
            ],
        };
    }

    createTriggerStepArgs(): types.Trigger {
        return {
            trigger: "some-pipeline",
            allow_dependency_failure: false,
            async: true,
            branches: "branch",
            depends_on: ["build"],
            if: `branch == "main"`,
            label: "label",
            skip: false,
            soft_fail: false,
        };
    }

    createWaitStepArgs(): types.Wait {
        return {
            wait: null,
            allow_dependency_failure: true,
            continue_on_failure: true,
            depends_on: ["build"],
            if: "branch == \"main\"",
        };
    }
}

export default new StepArgs();