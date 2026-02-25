using Buildkite.Sdk;
using Buildkite.Sdk.Schema;
using Xunit;

namespace Buildkite.Sdk.Tests;

public class NotificationTests
{
    [Fact]
    public void Pipeline_WithEmailNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new EmailNotification
        {
            Email = "team@example.com",
            If = "build.state == 'failed'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("email: team@example.com", yaml);
        Assert.Contains("if: build.state == 'failed'", yaml);
    }

    [Fact]
    public void Pipeline_WithSlackStringChannel_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new SlackNotification
        {
            Slack = "#builds"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("slack: '#builds'", yaml);
    }

    [Fact]
    public void Pipeline_WithSlackConfig_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new SlackNotification
        {
            Slack = new SlackConfig
            {
                Channels = new List<string> { "#builds", "#ops" },
                Message = "Build finished"
            },
            If = "build.state == 'passed'"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("channels:", yaml);
        Assert.Contains("#builds", yaml);
        Assert.Contains("#ops", yaml);
        Assert.Contains("message: Build finished", yaml);
        Assert.Contains("if: build.state == 'passed'", yaml);
    }

    [Fact]
    public void Pipeline_WithWebhookNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new WebhookNotification
        {
            Webhook = "https://example.com/webhook"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("webhook: https://example.com/webhook", yaml);
    }

    [Fact]
    public void Pipeline_WithPagerDutyNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new PagerDutyNotification
        {
            PagerdutyChangeEvent = "abc123"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("pagerduty_change_event: abc123", yaml);
    }

    [Fact]
    public void Pipeline_WithBasecampNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new BasecampNotification
        {
            BasecampCampfire = "https://3.basecamp.com/1234/integrations/5678/buckets/1/chats/2"
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("basecamp_campfire: https://3.basecamp.com/1234/integrations/5678/buckets/1/chats/2", yaml);
    }

    [Fact]
    public void Pipeline_WithGitHubCommitStatusNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new GitHubCommitStatusNotification
        {
            GithubCommitStatus = new GitHubCommitStatusConfig
            {
                Context = "ci/build"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("github_commit_status:", yaml);
        Assert.Contains("context: ci/build", yaml);
    }

    [Fact]
    public void Pipeline_WithGitHubCheckNotification_GeneratesCorrectOutput()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep { Label = "Build", Command = "make" });
        pipeline.AddNotify(new GitHubCheckNotification
        {
            GithubCheck = new GitHubCheckConfig
            {
                Context = "ci/tests"
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("github_check:", yaml);
        Assert.Contains("context: ci/tests", yaml);
    }

    [Fact]
    public void CommandStep_WithNotify_GeneratesStepLevelNotification()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new CommandStep
        {
            Label = "Build",
            Command = "make",
            Notify = new List<INotification>
            {
                new EmailNotification { Email = "dev@example.com" },
                new SlackNotification { Slack = "#eng" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("notify:", yaml);
        Assert.Contains("email: dev@example.com", yaml);
        Assert.Contains("slack: '#eng'", yaml);
    }

    [Fact]
    public void GroupStep_WithNotify_GeneratesStepLevelNotification()
    {
        var pipeline = new Pipeline();
        pipeline.AddStep(new GroupStep
        {
            Group = "Deploy",
            Steps = new List<IGroupStep>
            {
                new CommandStep { Label = "Deploy", Command = "./deploy.sh" }
            },
            Notify = new List<INotification>
            {
                new WebhookNotification { Webhook = "https://example.com/hook" }
            }
        });

        var yaml = pipeline.ToYaml();

        Assert.Contains("notify:", yaml);
        Assert.Contains("webhook: https://example.com/hook", yaml);
    }
}
