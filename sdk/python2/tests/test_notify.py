from buildkite_sdk import Pipeline, NotifyEmail, NotifyBasecamp, NotifySlack, NotifySlackObject, NotifyWebhook, NotifyPagerduty, NotifyGithubCommitStatus, NotifyGithubCommitStatusGithubCommitStatus, CommandStep, CommandStepDict
from .utils import TestRunner

class TestPipelineNotifyClass(TestRunner):
    def test_email(self):
        pipeline = Pipeline(
            notify=[
                NotifyEmail(email='dev@acmeinc.com')
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'email': 'dev@acmeinc.com'}]})

    def test_email_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyEmail(email='dev@acmeinc.com', pipeline_if="build.state == 'failed")
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'email': 'dev@acmeinc.com', 'if': "build.state == 'failed"}]})

    def test_basecamp_campfire(self):
        pipeline = Pipeline(
            notify=[
                NotifyBasecamp(basecamp_campfire='https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines')
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'}]})

    def test_basecamp_campfire_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyBasecamp(basecamp_campfire='https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', pipeline_if="build.state == 'failed")
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', 'if': "build.state == 'failed"}]})

    def test_slack(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack(slack='#channel')
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': '#channel'}]})

    def test_slack_if(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack(slack='#channel', pipeline_if="build.state == 'failed")
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': '#channel', 'if': "build.state == 'failed"}]})

    def test_slack_channels_list(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack(slack=NotifySlackObject(channels=['#one','#two']))
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': {'channels': ['#one','#two']}}]})

    def test_slack_message(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack(slack=NotifySlackObject(channels=['#one','#two'], message='a message'))
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': {'channels': ['#one','#two'], 'message': 'a message'}}]})

    def test_webhook(self):
        pipeline = Pipeline(
            notify=[
                NotifyWebhook(webhook='https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a')
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a'}]})

    def test_webhook_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyWebhook(webhook='https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a', pipeline_if="build.state == 'failed")
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a', 'if': "build.state == 'failed"}]})

    def test_pagerduty(self):
        pipeline = Pipeline(
            notify=[
                NotifyPagerduty(pagerduty_change_event='636d22Yourc0418Key3b49eee3e8')
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8'}]})

    def test_pagerduty_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyPagerduty(pagerduty_change_event='636d22Yourc0418Key3b49eee3e8', pipeline_if="build.state == 'failed")
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8', 'if': "build.state == 'failed"}]})

    def test_github_check_string(self):
        pipeline = Pipeline(
            notify=[
                'github_check'
            ]
        )
        self.validator.check_result(pipeline, {'notify': ['github_check']})

    def test_github_commit_status_string(self):
        pipeline = Pipeline(
            notify=[
                'github_commit_status'
            ]
        )
        self.validator.check_result(pipeline, {'notify': ['github_commit_status']})

    def test_github_commit_status(self):
        pipeline = Pipeline(
            notify=[
                NotifyGithubCommitStatus(
                    github_commit_status=NotifyGithubCommitStatusGithubCommitStatus(
                        context='my-custom-status'
                    )
                )
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'github_commit_status': {'context': 'my-custom-status'}}]})

    def test_github_commit_status_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyGithubCommitStatus(
                    pipeline_if="build.state == 'failed",
                    github_commit_status=NotifyGithubCommitStatusGithubCommitStatus(
                        context='my-custom-status'
                    )
                )
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'if': "build.state == 'failed", 'github_commit_status': {'context': 'my-custom-status'}}]})

class TestPipelineNotifyDict(TestRunner):
    def test_email(self):
        pipeline = Pipeline(
            notify=[
                NotifyEmail.from_dict({'email': 'dev@acmeinc.com'})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'email': 'dev@acmeinc.com'}]})

    def test_email_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyEmail.from_dict({'email': 'dev@acmeinc.com', 'if': "build.state == 'failed"})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'email': 'dev@acmeinc.com', 'if': "build.state == 'failed"}]})

    def test_basecamp_campfire(self):
        pipeline = Pipeline(
            notify=[
                NotifyBasecamp.from_dict({'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'}]})

    def test_basecamp_campfire_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyBasecamp.from_dict({'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', 'if': "build.state == 'failed"})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', 'if': "build.state == 'failed"}]})

    def test_slack(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack.from_dict({'slack': '#channel'})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': '#channel'}]})

    def test_slack_if(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack.from_dict({'slack': '#channel', 'if': "build.state == 'failed"})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': '#channel', 'if': "build.state == 'failed"}]})

    def test_slack_channels_list(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack.from_dict({'slack': {'channels': ['#one','#two']}})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': {'channels': ['#one','#two']}}]})

    def test_slack_message(self):
        pipeline = Pipeline(
            notify=[
                NotifySlack.from_dict({'slack': {'channels': ['#one','#two'], 'message': 'a message'}})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'slack': {'channels': ['#one','#two'], 'message': 'a message'}}]})

    def test_webhook(self):
        pipeline = Pipeline(
            notify=[
                NotifyWebhook.from_dict({'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a'})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a'}]})

    def test_webhook_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyWebhook.from_dict({'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a', 'if': "build.state == 'failed"})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'webhook': 'https://webhook.site/32raf257-168b-5aca-9067-3b410g78c23a', 'if': "build.state == 'failed"}]})

    def test_pagerduty(self):
        pipeline = Pipeline(
            notify=[
                NotifyPagerduty.from_dict({'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8'})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8'}]})

    def test_pagerduty_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyPagerduty.from_dict({'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8', 'if': "build.state == 'failed"})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'pagerduty_change_event': '636d22Yourc0418Key3b49eee3e8', 'if': "build.state == 'failed"}]})

    def test_github_commit_status(self):
        pipeline = Pipeline(
            notify=[
                NotifyGithubCommitStatus.from_dict({'github_commit_status': {'context': 'my-custom-status'}})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'github_commit_status': {'context': 'my-custom-status'}}]})

    def test_github_commit_status_if(self):
        pipeline = Pipeline(
            notify=[
                NotifyGithubCommitStatus.from_dict({'if': "build.state == 'failed", 'github_commit_status': {'context': 'my-custom-status'}})
            ]
        )
        self.validator.check_result(pipeline, {'notify': [{'if': "build.state == 'failed", 'github_commit_status': {'context': 'my-custom-status'}}]})

class TestCommandNotifyClass(TestRunner):
    def test_basecamp_campfire(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'}
            ],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifyBasecamp(basecamp_campfire='https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines')
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_basecamp_campfire_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {
                    'if': "build.state == 'failed",
                    'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'
                }
            ],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifyBasecamp(basecamp_campfire='https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines', pipeline_if="build.state == 'failed")
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': '#channel'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifySlack(slack='#channel')
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': '#channel', 'if': "build.state == 'failed'"}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifySlack(slack='#channel', pipeline_if="build.state == 'failed'")
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_channels(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': {'channels': ['#one', '#two']}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifySlack(slack=NotifySlackObject(channels=['#one','#two']))
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_message(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': {'channels': ['#one', '#two'], 'message': 'a message'}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifySlack(slack=NotifySlackObject(channels=['#one','#two'], message='a message'))
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_check(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': ['github_check']
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=['github_check']
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_string(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': ['github_commit_status']
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=['github_commit_status']
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_object(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'github_commit_status': {'context': 'my-context'}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifyGithubCommitStatus(
                            github_commit_status=NotifyGithubCommitStatusGithubCommitStatus(
                                context='my-context'
                            )
                        )
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'github_commit_status': {'context': 'my-context'}, 'if': "build.state == 'failed'"}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command='bash.sh',
                    notify=[
                        NotifyGithubCommitStatus(
                            pipeline_if="build.state == 'failed'",
                            github_commit_status=NotifyGithubCommitStatusGithubCommitStatus(
                                context='my-context'
                            )
                        )
                    ]
                )
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

class TestCommandNotifyDict(TestRunner):
    def test_basecamp_campfire(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'}
            ],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_basecamp_campfire_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {
                    'if': "build.state == 'failed",
                    'basecamp_campfire': 'https://3.basecamp.com/1234567/integrations/qwertyuiop/buckets/1234567/chats/1234567/lines'
                }
            ],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': '#channel'}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': '#channel', 'if': "build.state == 'failed'"}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_channels(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': {'channels': ['#one', '#two']}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_slack_message(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'slack': {'channels': ['#one', '#two'], 'message': 'a message'}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_check(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': ['github_check']
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_string(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': ['github_commit_status']
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_object(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'github_commit_status': {'context': 'my-context'}}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})

    def test_github_commit_status_if(self):
        expected: CommandStepDict = {
            'command': 'bash.sh',
            'notify': [
                {'github_commit_status': {'context': 'my-context'}, 'if': "build.state == 'failed'"}
            ]
        }
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(expected)
            ]
        )
        self.validator.check_result(pipeline, {'steps': [expected]})
