package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

type k schema.BlockStepNotify

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
