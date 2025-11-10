package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestBlockStep(t *testing.T) {
	t.Run("AllowDependencyFailure", func(t *testing.T) {
		value := true
		val := buildkite.BlockStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &value,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("AllowedTeams", func(t *testing.T) {
		value := "string"
		val := buildkite.BlockStep{
			AllowedTeams: &buildkite.AllowedTeams{
				String: &value,
			},
		}
		CheckResult(t, val, `{"allowed_teams":"string"}`)
	})

	t.Run("Block", func(t *testing.T) {
		value := "string"
		val := buildkite.BlockStep{
			Block: &value,
		}
		CheckResult(t, val, `{"block":"string"}`)
	})

	t.Run("BlockedState", func(t *testing.T) {
		value := buildkite.BlockStepBlockedStateValues["passed"]
		val := buildkite.BlockStep{
			BlockedState: &value,
		}
		CheckResult(t, val, `{"blocked_state":"passed"}`)
	})

	t.Run("Branches", func(t *testing.T) {
		value := "branch"
		val := buildkite.BlockStep{
			Branches: &buildkite.Branches{
				String: &value,
			},
		}
		CheckResult(t, val, `{"branches":"branch"}`)
	})

	t.Run("DependsOn", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			DependsOn: &buildkite.DependsOn{
				String: &value,
			},
		}
		CheckResult(t, val, `{"depends_on":"value"}`)
	})

	t.Run("Fields", func(t *testing.T) {
		text := "textField"
		fields := []buildkite.FieldsItem{
			{
				TextField: &buildkite.TextField{
					Text: &text,
				},
			},
		}
		val := buildkite.BlockStep{
			Fields: &fields,
		}
		CheckResult(t, val, `{"fields":[{"text":"textField"}]}`)
	})

	t.Run("Id", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Id: &value,
		}
		CheckResult(t, val, `{"id":"value"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Identifier: &value,
		}
		CheckResult(t, val, `{"identifier":"value"}`)
	})

	t.Run("If", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			If: &value,
		}
		CheckResult(t, val, `{"if":"value"}`)
	})

	t.Run("Key", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Key: &value,
		}
		CheckResult(t, val, `{"key":"value"}`)
	})

	t.Run("Label", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Label: &value,
		}
		CheckResult(t, val, `{"label":"value"}`)
	})

	t.Run("Name", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Name: &value,
		}
		CheckResult(t, val, `{"name":"value"}`)
	})

	t.Run("Prompt", func(t *testing.T) {
		value := "value"
		val := buildkite.BlockStep{
			Prompt: &value,
		}
		CheckResult(t, val, `{"prompt":"value"}`)
	})

	t.Run("Type", func(t *testing.T) {
		value := buildkite.BlockStepTypeValues["block"]
		val := buildkite.BlockStep{
			Type: &value,
		}
		CheckResult(t, val, `{"type":"block"}`)
	})

	t.Run("All", func(t *testing.T) {
		allowDependencyFailure := true
		allowedTeams := "allowedTeams"
		block := "block"
		blockedState := buildkite.BlockStepBlockedStateValues["passed"]
		branches := "branch"
		dependsOn := "step"
		fieldsText := "textField"
		fields := []buildkite.FieldsItem{
			{
				TextField: &buildkite.TextField{
					Text: &fieldsText,
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
		typeValue := buildkite.BlockStepTypeValues["block"]

		val := buildkite.BlockStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			AllowedTeams: &buildkite.AllowedTeams{
				String: &allowedTeams,
			},
			Block:        &block,
			BlockedState: &blockedState,
			Branches: &buildkite.Branches{
				String: &branches,
			},
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
			Type:       &typeValue,
		}
		CheckResult(t, val, `{"allow_dependency_failure":true,"allowed_teams":"allowedTeams","block":"block","blocked_state":"passed","branches":"branch","depends_on":"step","fields":[{"text":"textField"}],"id":"id","identifier":"identifier","if":"if","key":"key","label":"label","name":"name","prompt":"prompt","type":"block"}`)

	})
}

func TestNestedBlockStep(t *testing.T) {
	allowDependencyFailure := true
	allowedTeams := "allowedTeams"
	block := "block"
	blockedState := buildkite.BlockStepBlockedStateValues["passed"]
	branches := "branch"
	dependsOn := "step"
	fieldsText := "textField"
	fields := []buildkite.FieldsItem{
		{
			TextField: &buildkite.TextField{
				Text: &fieldsText,
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
	typeValue := buildkite.BlockStepTypeValues["block"]

	val := buildkite.NestedBlockStep{
		Block: &buildkite.BlockStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			AllowedTeams: &buildkite.AllowedTeams{
				String: &allowedTeams,
			},
			Block:        &block,
			BlockedState: &blockedState,
			Branches: &buildkite.Branches{
				String: &branches,
			},
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
			Type:       &typeValue,
		},
	}
	CheckResult(t, val, `{"block":{"allow_dependency_failure":true,"allowed_teams":"allowedTeams","block":"block","blocked_state":"passed","branches":"branch","depends_on":"step","fields":[{"text":"textField"}],"id":"id","identifier":"identifier","if":"if","key":"key","label":"label","name":"name","prompt":"prompt","type":"block"}}`)

}
