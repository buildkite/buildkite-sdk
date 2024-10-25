import Ajv from "ajv";
import type { ValidateFunction } from "ajv";
import { StepBuilder } from 'buildkite-pipeline-sdk';

const ajv = new Ajv({ allErrors: true });

export class SchemaValidator {
    private validator: ValidateFunction | null = null;

    public async fetchSchema() {
        const res = await fetch("https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json");
        const body = await res.json();
        const { fileMatch, ...schema } = body;
        this.validator = ajv.compile(schema);
    }

    public validate(stepBuilder: StepBuilder) {
        if (!this.validator) {
            return false;
        }

        const valid = this.validator({ steps: stepBuilder.steps });
        if (!valid) {
            return false;
        }

        return true;
    }
}