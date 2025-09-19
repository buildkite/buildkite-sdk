package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/pipeline-sdk/sdk/go/sdk/buildkite"
)

func TestInputStep(t *testing.T) {
	t.Run("AllowDepedencyFailure", func(t *testing.T) {
		allowDependencyFailure := true
		val := buildkite.InputStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("AllowedTeams", func(t *testing.T) {
		allowedTeams := "allowedTeams"
		val := buildkite.InputStep{
			AllowedTeams: &buildkite.AllowedTeams{
				String: &allowedTeams,
			},
		}
		CheckResult(t, val, `{"allowed_teams":"allowedTeams"}`)
	})

	t.Run("Input", func(t *testing.T) {
		input := "input"
		val := buildkite.InputStep{
			Input: &input,
		}
		CheckResult(t, val, `{"input":"input"}`)
	})

	t.Run("Branches", func(t *testing.T) {
		branches := "branch"
		val := buildkite.BlockStep{
			Branches: &buildkite.Branches{
				String: &branches,
			},
		}
		CheckResult(t, val, `{"branches":"branch"}`)
	})

	t.Run("DependsOn", func(t *testing.T) {
		dependsOn := "step"
		val := buildkite.InputStep{
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
		}
		CheckResult(t, val, `{"depends_on":"step"}`)
	})

	t.Run("Fields", func(t *testing.T) {
		text := "textField"
		fields := []buildkite.FieldsUnion{
			{
				TextField: &buildkite.TextField{
					Text: &text,
				},
			},
		}
		val := buildkite.InputStep{
			Fields: &fields,
		}
		CheckResult(t, val, `{"fields":[{"text":"textField"}]}`)
	})

	t.Run("Id", func(t *testing.T) {
		id := "id"
		val := buildkite.InputStep{
			Id: &id,
		}
		CheckResult(t, val, `{"id":"id"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		identifier := "identifier"
		val := buildkite.InputStep{
			Identifier: &identifier,
		}
		CheckResult(t, val, `{"identifier":"identifier"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifValue := "if"
		val := buildkite.InputStep{
			If: &ifValue,
		}
		CheckResult(t, val, `{"if":"if"}`)
	})

	t.Run("Key", func(t *testing.T) {
		key := "key"
		val := buildkite.InputStep{
			Key: &key,
		}
		CheckResult(t, val, `{"key":"key"}`)
	})

	t.Run("Label", func(t *testing.T) {
		label := "label"
		val := buildkite.InputStep{
			Label: &label,
		}
		CheckResult(t, val, `{"label":"label"}`)
	})

	t.Run("Name", func(t *testing.T) {
		name := "name"
		val := buildkite.InputStep{
			Name: &name,
		}
		CheckResult(t, val, `{"name":"name"}`)
	})

	t.Run("Prompt", func(t *testing.T) {
		prompt := "prompt"
		val := buildkite.InputStep{
			Prompt: &prompt,
		}
		CheckResult(t, val, `{"prompt":"prompt"}`)
	})

	t.Run("Type", func(t *testing.T) {
		typeVal := buildkite.InputStepTypeValues["input"]
		val := buildkite.InputStep{
			Type: &typeVal,
		}
		CheckResult(t, val, `{"type":"input"}`)
	})

	t.Run("All", func(t *testing.T) {
		allowDependencyFailure := true
		allowedTeams := "allowedTeams"
		input := "input"
		branches := "branch"
		dependsOn := "step"
		text := "textField"
		fields := []buildkite.FieldsUnion{
			{
				TextField: &buildkite.TextField{
					Text: &text,
				},
			},
		}
		id := "id"
		identifier := "identifier"
		ifValue := "if"
		key := "key"
		label := "label"
		name := "name"
		prompt := "prompt"
		typeVal := buildkite.InputStepTypeValues["input"]

		val := buildkite.InputStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			AllowedTeams: &buildkite.AllowedTeams{
				String: &allowedTeams,
			},
			Branches: &buildkite.Branches{
				String: &branches,
			},
			Input: &input,
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			Fields:     &fields,
			Id:         &id,
			Identifier: &identifier,
			If:         &ifValue,
			Key:        &key,
			Label:      &label,
			Name:       &name,
			Prompt:     &prompt,
			Type:       &typeVal,
		}
		CheckResult(t, val, `{"allow_dependency_failure":true,"allowed_teams":"allowedTeams","branches":"branch","depends_on":"step","fields":[{"text":"textField"}],"id":"id","identifier":"identifier","if":"if","input":"input","key":"key","label":"label","name":"name","prompt":"prompt","type":"input"}`)
	})
}

func TestNestedInputStep(t *testing.T) {
	allowDependencyFailure := true
	allowedTeams := "allowedTeams"
	input := "input"
	branches := "branch"
	dependsOn := "step"
	text := "textField"
	fields := []buildkite.FieldsUnion{
		{
			TextField: &buildkite.TextField{
				Text: &text,
			},
		},
	}
	id := "id"
	identifier := "identifier"
	ifValue := "if"
	key := "key"
	label := "label"
	name := "name"
	prompt := "prompt"
	typeVal := buildkite.InputStepTypeValues["input"]

	val := buildkite.NestedInputStep{
		Input: &buildkite.InputStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			AllowedTeams: &buildkite.AllowedTeams{
				String: &allowedTeams,
			},
			Branches: &buildkite.Branches{
				String: &branches,
			},
			Input: &input,
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			Fields:     &fields,
			Id:         &id,
			Identifier: &identifier,
			If:         &ifValue,
			Key:        &key,
			Label:      &label,
			Name:       &name,
			Prompt:     &prompt,
			Type:       &typeVal,
		},
	}
	CheckResult(t, val, `{"input":{"allow_dependency_failure":true,"allowed_teams":"allowedTeams","branches":"branch","depends_on":"step","fields":[{"text":"textField"}],"id":"id","identifier":"identifier","if":"if","input":"input","key":"key","label":"label","name":"name","prompt":"prompt","type":"input"}}`)
}
