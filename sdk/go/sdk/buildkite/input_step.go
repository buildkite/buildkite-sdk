package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

// Input Fields
type Field interface {
	toSchema() schema.Field
}

type InputTextField struct {
	Text     *string
	Key      *string
	Hint     *string
	Required *bool
	Default  *string
}

func (i InputTextField) toSchema() schema.Field {
	field := schema.Field{
		Text: i.Text,
		Key:  *i.Key,
		Hint: i.Hint,
	}

	if i.Required != nil {
		field.Required = &schema.AllowDependencyFailureUnion{
			Bool: i.Required,
		}
	}

	if i.Default != nil {
		field.Default = &schema.Branches{
			String: i.Default,
		}
	}

	return field
}

type InputSelectFieldOption struct {
	Label string
	Value string
}

type InputSelectField struct {
	Select   *string
	Options  []InputSelectFieldOption
	Key      *string
	Hint     *string
	Required *bool
	Default  *string
	Multiple *bool
}

func (i InputSelectField) toSchema() schema.Field {
	opts := make([]schema.Option, len(i.Options))
	for i, opt := range i.Options {
		opts[i] = schema.Option{
			Label: opt.Label,
			Value: opt.Value,
		}
	}

	field := schema.Field{
		Select:  i.Select,
		Key:     *i.Key,
		Hint:    i.Hint,
		Options: opts,
	}

	if i.Required != nil {
		field.Required = &schema.AllowDependencyFailureUnion{
			Bool: i.Required,
		}
	}

	if i.Default != nil {
		field.Default = &schema.Branches{
			String: i.Default,
		}
	}

	if i.Multiple != nil {
		field.Multiple = &schema.AllowDependencyFailureUnion{
			Bool: i.Multiple,
		}
	}

	return field
}

type InputStep struct {
	AllowDependencyFailure *bool
	Branches               []string
	DependsOn              DependsOn
	Fields                 []Field
	ID                     *string
	Identifier             *string
	If                     *string
	Input                  *string
	Key                    *string
	Label                  *string
	Name                   *string
	Prompt                 *string
}

func (step InputStep) toPipelineStep() *PipelineStep {
	fields := make([]schema.Field, len(step.Fields))
	for i, field := range step.Fields {
		fields[i] = field.toSchema()
	}

	inputStep := &PipelineStep{
		Input: &schema.Input{
			String: step.Input,
		},

		Fields:     fields,
		ID:         step.ID,
		Identifier: step.Identifier,
		If:         step.If,
		Key:        step.Key,
		Label:      step.Label,
		Name:       step.Name,
		Prompt:     step.Prompt,
	}

	if step.AllowDependencyFailure != nil {
		inputStep.AllowDependencyFailure = &schema.AllowDependencyFailureUnion{
			Bool: step.AllowDependencyFailure,
		}
	}

	if step.Branches != nil {
		inputStep.Branches = &schema.Branches{
			StringArray: step.Branches,
		}
	}

	if step.DependsOn != nil {
		inputStep.DependsOn = step.DependsOn.toSchema()
	}

	return inputStep
}
