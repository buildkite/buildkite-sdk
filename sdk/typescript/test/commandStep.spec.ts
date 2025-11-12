import { createValidator, PipelineStepValidator } from './utils'

describe('CommandStep', () => {
    let validatePipeline: PipelineStepValidator
    beforeAll(async () => {
        const { step } = await createValidator()
        validatePipeline = step
    })

    describe('Nesting', () => {
        it('ScriptType', () => {
            validatePipeline({ type: 'script' })
        })

        it('Command', () => {
            validatePipeline({ command: 'bash.sh' })
        })

        it('CommandType', () => {
            validatePipeline({ type: 'command', command: 'bash.sh' })
        })

        it('CommandsType', () => {
            validatePipeline({ type: 'commands', command: 'bash.sh' })
        })

        it('NestedCommand', () => {
            validatePipeline({ command: { command: 'bash.sh' }})
        })

        it('NestedCommands', () => {
            validatePipeline({ commands: { command: 'bash.sh' }})
        })

        it('NestedScript', () => {
            validatePipeline({ script: { command: 'bash.sh' }})
        })
    })

    it('MultipleCommands', () => {
        validatePipeline({
            command: ["one", "two"]
        })
    })

    it('PluginOnly', () => {
        validatePipeline({ plugins: [{'a-plugin#v1.0.0': {run: 'app'}}] })
    })

    describe('Agents', () => {
        it('Object', () => {
             validatePipeline({
                command: 'test',
                agents: {
                    os: 'macOS'
                },
            })
        })

        it('List', () => {
            validatePipeline({
                command: 'test',
                agents: [ 'os=macOS' ],
            })
        })
    })

    describe('ArtifactPath', () => {
        it('Array', () => {
            validatePipeline({
                command: 'test',
                artifact_paths: [
                    'one',
                    'two',
                ],
            })
        })

        it('String', () => {
            validatePipeline({
                command: 'test',
                artifact_paths: 'one',
            })
        })
    })

    it('Branches', () => {
        validatePipeline({
            command: 'test',
            branches: 'branch',
        })
    })

    it('Concurrency', () => {
        validatePipeline({
            command: 'test',
            concurrency: 1,
            concurrency_group: 'my-group',
            concurrency_method: 'eager',
        })
    })

    it('Env', () => {
        validatePipeline({
            command: 'test',
            env: { AN_ENV: 'value' },
        })
    })

    it('Id', () => {
        validatePipeline({
            command: 'test',
            id: 'id',
        })
    })

    it('Identifier', () => {
        validatePipeline({
            command: 'test',
            identifier: 'id',
        })
    })

    it('Label', () => {
        validatePipeline({
            command: 'test',
            label: 'a label',
        })
    })

    it('Parallelism', () => {
        validatePipeline({
            command: 'test',
            parallelism: 42,
        })
    })

    describe('Plugins', () => {
        it('Object', () => {
            validatePipeline({
                command: 'test',
                plugins: {'a-plugin#v1.0.0': {run: 'app'}},
            })
        })

        it('StringArray', () => {
            validatePipeline({
                command: 'test',
                plugins: ['a-plugin#v1.0.0']
            })
        })

        it('ObjectArray', () => {
            validatePipeline({
                command: 'test',
                plugins: [{'a-plugin#v1.0.0': {run: 'app'}}],
            })
        })
    })

    describe('Retry', () => {
        it('ObjectArray', () => {
            validatePipeline({
                command: 'test',
                retry: {
                    automatic: [
                        { exit_status: -1, signal_reason: 'none' },
                        { signal: 'kill' },
                        { exit_status: 255 },
                        { exit_status: 3, limit: 3 },
                    ],
                },
            })
        })

        it('ExitStatus', () => {
            validatePipeline({
                command: 'test',
                retry: {
                    automatic: {
                        exit_status: -1,
                    }
                }
            })
        })

        it('ExitStatusArray', () => {
            validatePipeline({
                command: 'test',
                retry: {
                    automatic: {
                        exit_status: [1, 2, 3],
                    },
                },
            })
        })

        it('Boolean', () => {
            validatePipeline({
                command: 'test',
                retry: {
                    automatic: true,
                },
            })
        })
    })

    describe('Secrets', () => {
        it('StringArray', () => validatePipeline({
            command: 'test',
            secrets: ['MY_SECRET'],
        }))

        it('Object', () => validatePipeline({
            command: 'test',
            secrets: {
                'MY_SECRET': 'API_TOKEN',
            },
        }))
    })

    describe('Skip', () => {
        it('Boolean', () => {
            validatePipeline({
                command: 'test',
                skip: true,
            })
        })

        it('String', () => {
            validatePipeline({
                command: 'test',
                skip: 'reason',
            })
        })
    })

    it('TimeoutInMinutes', () => {
        validatePipeline({
            command: 'test',
            timeout_in_minutes: 1,
        })
    })

    describe('SoftFail', () => {
        it('Boolean', () => {
            validatePipeline({
                command: 'test',
                soft_fail: true,
            })
        })

        it('ObjectNumber', () => {
            validatePipeline({
                command: 'test',
                soft_fail: [
                    { exit_status: -1 },
                ],
            })
        })

        it('ObjectString', () => {
            validatePipeline({
                command: 'test',
                soft_fail: [
                    { exit_status: '*' },
                ],
            })
        })
    })

    it('If', () => {
        validatePipeline({
            command: 'test',
            if: 'build.message !~ /skip tests/',
        })
    })

    it('Key', () => {
        validatePipeline({
            command: 'test',
            key: 'key',
        })
    })

    describe('DependsOn', () => {
        it('String', () => {
            validatePipeline({
                command: 'test',
                depends_on: 'depend-on-me',
            })
        })

        it('StringArray', () => {
            validatePipeline({
                command: 'test',
                depends_on: [
                    'depend-on-me-1',
                    'depend-on-me-2',
                ],
            })
        })

        it('Object', () => {
            validatePipeline({
                command: 'test',
                depends_on: [
                    { step: 'depend-on-me', allow_failure: true },
                ],
            })
        })

        it('ObjectArray', () => {
            validatePipeline({
                command: 'test',
                depends_on: [
                    { step: 'depend-on-me-1' },
                    { step: 'depend-on-me-2' },
                ],
            })
        })

        it('Mixed', () => {
            validatePipeline({
                command: 'test',
                depends_on: [
                    'depend-on-me-1',
                    { step: 'depend-on-me-2' },
                ],
            })
        })
    })

    it('AllowDependencyFailure', () => {
        validatePipeline({
            command: 'test',
            allow_dependency_failure: true,
        })
    })

    it('Priority', () => {
        validatePipeline({
            command: 'test',
            priority: 100,
        })
    })

    it('CancelOnBuildFailing', () => {
        validatePipeline({
            command: 'test',
            cancel_on_build_failing: false,
        })
    })

    it('Matrix', () => {
        validatePipeline({
            command: 'echo {{matrix}}',
            env: { FOO: 'bar' },
            plugins: [
                { 'docker#v3.0.0': { image: 'alpine', 'always-pull': true }},
            ],
            matrix: ['one','two','three'],
            signature: {
                value: 'not a real signature value',
                algorithm: 'HS256',
                signed_fields: [
                    'command',
                    'env::FOO',
                    'plugins',
                    'matrix',
                ],
            },
        })
    })

    describe('Cache', () => {
        it('Array', () => {
            validatePipeline({
                command: 'test',
                cache: [
                    'dist/',
                    './src/target/',
                ],
            })

            validatePipeline({
                command: 'test',
                cache: 'dist/',
            })
        })
    })

    it('IfChanged', () => {
        validatePipeline({
            command: 'test',
            if_changed: '*.txt',
        })
    })
})
