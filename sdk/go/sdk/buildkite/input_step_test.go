package buildkite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputStep(t *testing.T) {
	t.Run("should create a text input step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(InputStep{
			Input: Value("text"),
			Fields: []Field{
				InputTextField{
					Text: Value("Text Input"),
					Key:  Value("text-input"),
				},
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "fields": [
                {
                    "text": "Text Input",
                    "key": "text-input"
                }
            ],
            "input": "text"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should create a select input step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(InputStep{
			Input: Value("select"),
			Fields: []Field{
				InputSelectField{
					Select: Value("Select"),
					Key:    Value("select"),
					Options: []InputSelectFieldOption{
						{
							Value: "one",
							Label: "One",
						},
						{
							Value: "two",
							Label: "Two",
						},
					},
				},
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "fields": [
                {
                    "key": "select",
                    "select": "Select",
                    "options": [
                        {
                            "label": "One",
                            "value": "one"
                        },
                        {
                            "label": "Two",
                            "value": "two"
                        }
                    ]
                }
            ],
            "input": "select"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})

	t.Run("should create a complex input step", func(t *testing.T) {
		pipeline := NewPipeline()
		pipeline.AddStep(InputStep{
			AllowDependencyFailure: Value(true),
			Branches:               []string{"main"},
			DependsOn:              DependsOnString("build"),
			ID:                     Value("id"),
			Identifier:             Value("identifier"),
			If:                     Value("if"),
			Key:                    Value("key"),
			Label:                  Value("label"),
			Name:                   Value("name"),
			Prompt:                 Value("prompt"),
			Input:                  Value("complex"),
			Fields: []Field{
				InputTextField{
					Text:    Value("Text Input"),
					Key:     Value("text-input"),
					Hint:    Value("hint"),
					Default: Value("default"),
				},
				InputTextField{
					Text:     Value("Text Input Required"),
					Key:      Value("text-input-required"),
					Required: Value(true),
				},
				InputTextField{
					Text:     Value("Text Input Optional"),
					Key:      Value("text-input-optional"),
					Required: Value(false),
				},
				InputSelectField{
					Select: Value("Select"),
					Key:    Value("select"),
					Options: []InputSelectFieldOption{
						{
							Value: "one",
							Label: "One",
						},
						{
							Value: "two",
							Label: "Two",
						},
					},
					Hint:     Value("hint"),
					Required: Value(true),
					Default:  Value("default"),
					Multiple: Value(true),
				},
			},
		})

		result, err := pipeline.ToJSON()
		assert.NoError(t, err)

		expected := `{
    "steps": [
        {
            "allow_dependency_failure": true,
            "branches": [
                "main"
            ],
            "depends_on": "build",
            "fields": [
                {
                    "text": "Text Input",
                    "key": "text-input",
                    "hint": "hint",
                    "default": "default"
                },
                {
                    "text": "Text Input Required",
                    "key": "text-input-required",
                    "required": true
                },
                {
                    "text": "Text Input Optional",
                    "key": "text-input-optional",
                    "required": false
                },
                {
                    "key": "select",
                    "hint": "hint",
                    "required": true,
                    "default": "default",
                    "select": "Select",
                    "options": [
                        {
                            "label": "One",
                            "value": "one"
                        },
                        {
                            "label": "Two",
                            "value": "two"
                        }
                    ],
                    "multiple": true
                }
            ],
            "id": "id",
            "identifier": "identifier",
            "if": "if",
            "key": "key",
            "label": "label",
            "name": "name",
            "prompt": "prompt",
            "input": "complex"
        }
    ]
}`
		assert.Equal(t, expected, result)
	})
}
