import { Pipeline, PipelineSteps, BuildkitePipeline } from '../src/index'
import { Ajv, ValidateFunction } from 'ajv'

export type PipelineStepValidator = (step: PipelineSteps[0]) => void
export type PipelineSchemaValidator = (pipeline: BuildkitePipeline) => void

interface PipelineValidator {
    step: PipelineStepValidator
    pipeline: PipelineSchemaValidator
}

export async function createValidator(): Promise<PipelineValidator> {
    const response = await fetch('https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json')
    const rawSchema = await response.text()
    const { fileMatch, ...schema } = JSON.parse(rawSchema)

    const ajv = new Ajv({ allErrors: true })
    const validator = ajv.compile({
        definitions: schema['definitions'],
    })

    return {
        step: (step: PipelineSteps[0]) => {
            const pipeline = new Pipeline()
            pipeline.addStep(step)
            expect(validator(pipeline.toJSON())).toBe(true)
        },
        pipeline: (pipeline: BuildkitePipeline) => {
            const buildkitePipeline = new Pipeline()
            buildkitePipeline.setPipeline(pipeline)
            expect(validator(buildkitePipeline.toJSON())).toBe(true)
        }
    }
}
