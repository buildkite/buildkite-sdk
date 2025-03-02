import { Pipeline } from "@buildkite/buildkite-sdk";
import * as fs from "fs";

const pipeline = new Pipeline();

pipeline.addStep({
    label: "some-label",
    command: "echo 'Hello, world!'",
});

fs.mkdirSync("./out/apps/typescript", { recursive: true });

fs.writeFileSync(
    "./out/apps/typescript/pipeline.json",
    pipeline.toJSON(),
    "utf-8"
);

fs.writeFileSync(
    "./out/apps/typescript/pipeline.yaml",
    pipeline.toYAML(),
    "utf-8"
);
