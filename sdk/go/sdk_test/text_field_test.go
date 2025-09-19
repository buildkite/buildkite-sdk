package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

type testTextFieldRequired struct {
	Required buildkite.TextFieldRequired `json:"required"`
}

func TestTextField(t *testing.T) {
	t.Run("TextFieldRequired", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			value := "true"
			val := testTextFieldRequired{
				Required: buildkite.TextFieldRequired{
					String: &value,
				},
			}
			CheckResult(t, val, `{"required":"true"}`)
		})

		t.Run("Boolean", func(t *testing.T) {
			value := true
			val := testTextFieldRequired{
				Required: buildkite.TextFieldRequired{
					Bool: &value,
				},
			}
			CheckResult(t, val, `{"required":true}`)
		})
	})

	t.Run("TextField", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			defaultVal := "default"
			val := buildkite.TextField{
				Default: &defaultVal,
			}
			CheckResult(t, val, `{"default":"default"}`)
		})

		t.Run("Format", func(t *testing.T) {
			format := "format"
			val := buildkite.TextField{
				Format: &format,
			}
			CheckResult(t, val, `{"format":"format"}`)
		})

		t.Run("Hint", func(t *testing.T) {
			hint := "hint"
			val := buildkite.TextField{
				Hint: &hint,
			}
			CheckResult(t, val, `{"hint":"hint"}`)
		})

		t.Run("Key", func(t *testing.T) {
			key := "key"
			val := buildkite.TextField{
				Key: &key,
			}
			CheckResult(t, val, `{"key":"key"}`)
		})

		t.Run("Required", func(t *testing.T) {
			required := "true"
			val := buildkite.TextField{
				Required: &buildkite.TextFieldRequired{
					String: &required,
				},
			}
			CheckResult(t, val, `{"required":"true"}`)
		})

		t.Run("Text", func(t *testing.T) {
			text := "text"
			val := buildkite.TextField{
				Text: &text,
			}
			CheckResult(t, val, `{"text":"text"}`)
		})

		t.Run("All", func(t *testing.T) {
			defaultVal := "default"
			format := "format"
			hint := "hint"
			key := "key"
			required := "true"
			text := "text"
			val := buildkite.TextField{
				Default: &defaultVal,
				Format:  &format,
				Hint:    &hint,
				Key:     &key,
				Required: &buildkite.TextFieldRequired{
					String: &required,
				},
				Text: &text,
			}
			CheckResult(t, val, `{"default":"default","format":"format","hint":"hint","key":"key","required":"true","text":"text"}`)
		})
	})
}
