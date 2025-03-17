package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

/**
 *   Notify Types
 */

// Slack
type NotifySlack interface {
	toSchema() *schema.IndecentSlack
}

type NotifySlackSimple string

func (n NotifySlackSimple) toSchema() *schema.IndecentSlack {
	val := string(n)
	return &schema.IndecentSlack{
		String: &val,
	}
}

type NotifySlackAdvanced schema.TentacledSlack

func (n NotifySlackAdvanced) toSchema() *schema.IndecentSlack {
	return &schema.IndecentSlack{
		TentacledSlack: &schema.TentacledSlack{
			Channels: n.Channels,
			Message:  n.Message,
		},
	}
}

// Notify interface
type Notify interface {
	toSchema() *schema.BlockStepNotify
}

// Step Notify
type StepNotify struct {
	Slack NotifySlack
}

func (c StepNotify) toSchema() *schema.BlockStepNotify {
	notify := &schema.BlockStepNotify{}

	if c.Slack != nil {
		notify.FluffyBuildNotify = &schema.FluffyBuildNotify{
			Slack: c.Slack.toSchema(),
		}
	}

	return notify
}

// Pipeline Notify
type NotifyGitHubCommitStatus struct {
	Context *string
}

type PipelineNotify struct {
	Email                *string                   `json:"email,omitempty"`
	If                   *string                   `json:"if,omitempty"`
	BasecampCampfire     *string                   `json:"basecamp_campfire,omitempty"`
	Slack                *NotifySlack              `json:"slack,omitempty"`
	Webhook              *string                   `json:"webhook,omitempty"`
	PagerdutyChangeEvent *string                   `json:"pagerduty_change_event,omitempty"`
	GitHubCommitStatus   *NotifyGitHubCommitStatus `json:"github_commit_status,omitempty"`
	GitHubCheck          map[string]interface{}    `json:"github_check,omitempty"`
}

func (p PipelineNotify) toSchema() *schema.BlockStepNotify {
	notify := &schema.FluffyBuildNotify{}

	if p.Email != nil {
		notify.Email = p.Email
	}

	if p.If != nil {
		notify.If = p.If
	}

	if p.BasecampCampfire != nil {
		notify.BasecampCampfire = p.BasecampCampfire
	}

	if p.Slack != nil {
		notify.Slack = p.toSchema().FluffyBuildNotify.Slack
	}

	if p.Webhook != nil {
		notify.Webhook = p.Webhook
	}

	if p.PagerdutyChangeEvent != nil {
		notify.PagerdutyChangeEvent = p.PagerdutyChangeEvent
	}

	if p.GitHubCommitStatus != nil {
		notify.GithubCommitStatus = (*schema.TentacledGithubCommitStatus)(p.GitHubCommitStatus)
	}

	if p.GitHubCheck != nil {
		notify.GithubCheck = p.GitHubCheck
	}

	return &schema.BlockStepNotify{
		FluffyBuildNotify: notify,
	}
}
