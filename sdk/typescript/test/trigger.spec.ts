import { createValidator, PipelineStepValidator } from './utils'

describe('TriggerStep', () => {
    let validatePipeline: PipelineStepValidator
    beforeAll(async () => {
        const { step } = await createValidator()
        validatePipeline = step
    })

    describe('NestingFormats', () => {
        it('Simple', () => {
            validatePipeline({
                trigger: "a-slug",
            })
        })

        it('Type', () => {
            validatePipeline({
                type: 'trigger',
                trigger: 'a-slug',
            })
        })
    })

    it('Async', () => {
        validatePipeline({
            trigger: 'a-slug',
            async: true,
        })
    })

    it('Branches', () => {
         validatePipeline({
            trigger: 'a-slug',
            branches: ['one','two']
        })
    })

    it('Build', () => {
         validatePipeline({
            trigger: 'a-slug',
            build: {
                branch: 'main',
                commit: 'a commit',
                env: { FOO: 'BAR' },
                message: 'a message',
                meta_data: { 'a-key': 'a-val' },
            },
        })
    })

    it('Id', () => {
        validatePipeline({
            trigger: 'a-slug',
            id: 'id',
        })
    })

    it('Identifier', () => {
        validatePipeline({
            trigger: 'a-slug',
            identifier: 'identifier',
        })
    })

    it('Label', () => {
        validatePipeline({
            trigger: 'a-slug',
            label: 'label',
        })
    })

    it('If', () => {
        validatePipeline({
            trigger: 'a-slug',
            if: 'build.message !~ /skip tests/',
        })
    })

    it('Key', () => {
        validatePipeline({
            trigger: 'a-slug',
            key: 'key',
        })
    })

    describe('DependsOn', () => {
        it('String', () => {
            validatePipeline({
                trigger: 'a-slug',
                depends_on: 'step',
            })
        })

        it('StringArray', () => {
            validatePipeline({
                trigger: 'a-slug',
                depends_on: ['one','two'],
            })
        })

        it('Object', () => {
            validatePipeline({
                trigger: 'a-slug',
                depends_on: [{ step: 'depend-on-me', allow_failure: true }],
            })
        })

        it('ObjectArray', () => {
            validatePipeline({
                trigger: 'a-slug',
                depends_on: [
                    { step: 'depend-on-me-1' },
                    { step: 'depend-on-me-2' },
                ],
            })
        })

        it('Mixed', () => {
            validatePipeline({
                trigger: 'a-slug',
                depends_on: [
                    'depend-on-me-1',
                    { step: 'depend-on-me-2' },
                ],
            })
        })
    })

    it('AllowDependencyFailure', () => {
        validatePipeline({
            trigger: 'a-slug',
            allow_dependency_failure: true,
        })
    })

    describe('Skip', () => {
        it('Boolean', () => {
            validatePipeline({
                trigger: 'a-slug',
                skip: true,
            })
        })

        it('String', () => {
            validatePipeline({
                trigger: 'a-slug',
                skip: 'reason',
            })
        })
    })

    describe('Softfail', () => {
        it('Boolean', () => {
            validatePipeline({
                trigger: 'a-slug',
                soft_fail: true,
            })
        })

        it('ExitStatusNumber', () => {
            validatePipeline({
                trigger: 'a-slug',
                soft_fail: [
                    { exit_status: -1 },
                ],
            })
        })

        it('ExistStatusString', () => {
            validatePipeline({
                trigger: 'a-slug',
                soft_fail: [
                    { exit_status: '*' },
                ],
            })
        })
    })

    it('IfChanged', () => {
        validatePipeline({
            trigger: 'a-slug',
            if_changed: '*.txt',
        })
    })
})
