package buildkite

import "github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"

// Input Fields
type field struct {
	Text     *string                  `json:"text,omitempty"`
	Key      *string                  `json:"key,omitempty"`
	Hint     *string                  `json:"hint,omitempty"`
	Required *bool                    `json:"required,omitempty"`
	Default  *string                  `json:"default,omitempty"`
	Select   *string                  `json:"select,omitempty"`
	Options  []InputSelectFieldOption `json:"options,omitempty"`
	Multiple *bool                    `json:"multiple,omitempty"`
}

type Field interface {
	toSchema() field
}

type InputTextField struct {
	Text     *string
	Key      *string
	Hint     *string
	Required *bool
	Default  *string
}

func (i InputTextField) toSchema() field {
	field := field{
		Text:     i.Text,
		Key:      i.Key,
		Hint:     i.Hint,
		Required: i.Required,
		Default:  i.Default,
	}

	return field
}

type InputSelectFieldOption struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
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

func (i InputSelectField) toSchema() field {
	field := field{
		Select:   i.Select,
		Hint:     i.Hint,
		Options:  i.Options,
		Key:      i.Key,
		Required: i.Required,
		Default:  i.Default,
		Multiple: i.Multiple,
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

func (step InputStep) ToPipelineStep() *PipelineStep {
	fields := make([]field, len(step.Fields))
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
