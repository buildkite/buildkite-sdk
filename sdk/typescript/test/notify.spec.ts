import { createValidator, PipelineSchemaValidator, PipelineStepValidator } from './utils'

describe('Notify', () => {
    describe('PipelineNotify', () => {
        let validatePipeline: PipelineSchemaValidator
        beforeAll(async () => {
            const { pipeline } = await createValidator()
            validatePipeline = pipeline
        })

        it('Email', () => {
            validatePipeline({
                notify: [
                    { email: 'dev@acmeinc.com', if: `build.state == 'failed'` },
                ],
            })
        })

        it('BasecampCampfire', () => {
            validatePipeline({
                notify: [
                    { basecamp_campfire: 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', if: `build.state == 'failed'` },
                ],
            })
        })

        describe('Slack', () => {
            it('Simple', () => {
                validatePipeline({
                    notify: [
                        { slack: '#channel', if: `build.state == 'failed'` },
                    ],
                })
            })

            it('Detailed', () => {
                validatePipeline({
                    notify: [
                        {
                            slack: {
                                channels: ['important-business#announcements'],
                                message: 'CI announcement'
                            },
                            if: `build.state == 'failed'`,
                        },
                    ],
                })
            })
        })

        it('Webhook', () => {
            validatePipeline({
                notify: [
                    { webhook: 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a', if: `build.state == 'failed'` },
                ],
            })
        })

        it('PagerDuty', () => {
            validatePipeline({
                notify: [
                    { pagerduty_change_event: '636d22Yourc0418Key3b49eee3e8', if: `build.state == 'failed'` },
                ],
            })
        })

        describe('GithubCheck', () => {
            it('String', () => {
                validatePipeline({
                    notify: [
                        'github_check',
                    ],
                })
            })

            it('Object', () => {
                validatePipeline({
                    notify: [
                        { github_check: { foo: 'bar' } },
                    ],
                })
            })
        })

        describe('GithubCommitStatus', () => {
            it('String', () => {
                validatePipeline({
                    notify: [ 'github_commit_status' ],
                })
            })

            it('Object', () => {
                validatePipeline({
                    notify: [
                        { github_commit_status: { context: 'my-custom-status' }, if: `build.state == 'failed'` },
                    ],
                })
            })
        })
    })

    describe('CommandNotify', () => {
        let validatePipeline: PipelineStepValidator
        beforeAll(async () => {
            const { step } = await createValidator()
            validatePipeline = step
        })

        it('BasecampCampfire', () => {
            validatePipeline({
                command: "blah.sh",
                notify: [
                    { basecamp_campfire: 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', if: `build.state == 'failed'` },
                ],
            })
        })

        describe('Slack', () => {
            it('Simple', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [
                        { slack: '#channel', if: `build.state == 'failed'` },
                    ],
                })
            })

            it('Detailed', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [
                        {
                            slack: {
                                channels: ['important-business#announcements'],
                                message: 'CI announcement'
                            },
                            if: `build.state == 'failed'`,
                        },
                    ],
                })
            })
        })

        describe('GithubCheck', () => {
            it('String', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [
                        'github_check',
                    ],
                })
            })

            it('Object', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [
                        { github_check: { foo: 'bar' } },
                    ],
                })
            })
        })

        describe('GithubCommitStatus', () => {
            it('String', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [ 'github_commit_status' ],
                })
            })

            it('Object', () => {
                validatePipeline({
                    command: "blah.sh",
                    notify: [
                        { github_commit_status: { context: 'my-custom-status' }, if: `build.state == 'failed'` },
                    ],
                })
            })
        })
    })
})
