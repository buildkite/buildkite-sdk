import { createValidator, PipelineStepValidator } from './utils'

describe('WaitStep', () => {
    let validatePipeline: PipelineStepValidator
    beforeAll(async () => {
        const { step } = await createValidator()
        validatePipeline = step
    })

    describe('NestingTypes', () => {
        it('StringWait', () => {
            validatePipeline('wait')
        })

        it('StringWaiter', () => {
            validatePipeline('waiter')
        })

        it('WaitType', () => {
            validatePipeline({ type: 'wait' })
        })

        it('WaiterType', () => {
            validatePipeline({ type: 'waiter' })
        })

        it('SimpleWait', () => {
            validatePipeline({ wait: '~' })
        })

        it('NestedWait', () => {
            validatePipeline({ wait: { continue_on_failure: true } })
        })

        it('NestedWaiter', () => {
            validatePipeline({ waiter: { continue_on_failure: true } })
        })
    })

    it('Id', () => {
        validatePipeline({
            type: 'waiter',
            id: 'id',
        })
    })

    it('Identifier', () => {
        validatePipeline({
            type: 'waiter',
            identifier: 'identifier',
        })
    })

    it('If', () => {
        validatePipeline({
            type: 'waiter',
            if: 'build.message !~ /skip tests/',
        })
    })

    it('Key', () => {
        validatePipeline({
            type: 'waiter',
            key: 'key',
        })
    })

    describe('DependsOn', () => {
        it('String', () => {
            validatePipeline({
                type: 'waiter',
                depends_on: 'step',
            })
        })

        it('StringArray', () => {
            validatePipeline({
                type: 'waiter',
                depends_on: ['one','two'],
            })
        })

        it('Object', () => {
            validatePipeline({
                type: 'waiter',
                depends_on: [{ step: 'depend-on-me', allow_failure: true }],
            })
        })

        it('ObjectArray', () => {
            validatePipeline({
                type: 'waiter',
                depends_on: [
                    { step: 'depend-on-me-1' },
                    { step: 'depend-on-me-2' },
                ],
            })
        })

        it('Mixed', () => {
            validatePipeline({
                type: 'waiter',
                depends_on: [
                    'depend-on-me-1',
                    { step: 'depend-on-me-2' },
                ],
            })
        })
    })

    it('AllowDependencyFailure', () => {
        validatePipeline({
            type: 'waiter',
            allow_dependency_failure: true,
        })
    })
})
