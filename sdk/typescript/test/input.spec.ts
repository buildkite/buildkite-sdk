import { createValidator, PipelineStepValidator } from './utils'

describe('InputStep', () => {
    let validatePipeline: PipelineStepValidator
    beforeAll(async () => {
        const { step } = await createValidator()
        validatePipeline = step
    })

    describe('NestingTypes', () => {
        it('String', () => {
            validatePipeline('input')
        })

        it('Simple', () => {
            validatePipeline({ input: 'a label' })
        })

        it('Nested', () => {
            validatePipeline({
                input: {
                    fields: [
                        { text: 'Field 1', key: 'field-1' },
                    ],
                },
            })
        })

        it('Type', () => {
            validatePipeline({
                type: 'input',
                label: 'a label',
                fields: [
                    { text: 'Field 1', key: 'field-1' },
                ],
            })
        })
    })

    it('AllowedTeams', () => {
        validatePipeline({
            input: 'a label',
            allowed_teams: 'team',
        })
    })

    it('Branches', () => {
        validatePipeline({
            input: 'a label',
            branches: 'main',
        })
    })

    it('Id', () => {
        validatePipeline({
            input: 'a label',
            id: 'id',
        })
    })

    it('Identifier', () => {
        validatePipeline({
            input: 'a label',
            identifier: 'identifier',
        })
    })

    it('Prompt', () => {
        validatePipeline({
            input: 'a label',
            prompt: 'prompt',
        })
    })

    it('Fields', () => {
        validatePipeline({
            input: 'a label',
            fields: [
                {
                    text: 'Field 1',
                    key: 'field-1',
                },
                {
                    text: 'Field 2',
                    key: 'field-2',
                    required: false,
                    default: 'Field 2 Default',
                    hint: 'Field 2 Hint',
                },
                {
                    select: 'Select 1',
                    key: 'select-1',
                    multiple: true,
                    options: [
                        { label: 'Select 1 Option 1', value: 'select-1-option-1' },
                        { label: 'Select 1 Option 2', value: 'select-1-option-2' },
                    ],
                },
                {
                    select: 'Select 2',
                    key: 'select-2',
                    hint: 'Select 2 Hint',
                    required: false,
                    default: 'select-2-option-1',
                    options: [
                        { label: 'Select 2 Option 1', value: 'select-2-option-1' },
                    ],
                },
            ],
        })
    })

    it('If', () => {
        validatePipeline({
            input: 'a label',
            if: 'build.message !~ /skip tests/',
        })
    })

    it('Key', () => {
        validatePipeline({
            input: 'a label',
            key: 'key',
        })
    })

    describe('DependsOn', () => {
        it('String', () => {
            validatePipeline({
                input: 'a label',
                depends_on: 'depend-on-me',
            })
        })

        it('StringArray', () => {
            validatePipeline({
                input: 'a label',
                depends_on: [
                    'depend-on-me-1',
                    'depend-on-me-2',
                ],
            })
        })

        it('Object', () => {
            validatePipeline({
                input: 'a label',
                depends_on: [
                    { step: 'depend-on-me', allow_failure: true },
                ],
            })
        })

        it('ObjectArray', () => {
            validatePipeline({
                input: 'a label',
                depends_on: [
                    { step: 'depend-on-me-1' },
                    { step: 'depend-on-me-2' },
                ],
            })
        })

        it('Mixed', () => {
            validatePipeline({
                input: 'a label',
                depends_on: [
                    'depend-on-me-1',
                    { step: 'depend-on-me-2' },
                ],
            })
        })
    })

    it('AllowDependencyFailure', () => {
        validatePipeline({
            input: 'a label',
            allow_dependency_failure: true,
        })
    })
})
