package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestCommandStep(t *testing.T) {
	t.Run("Agents", func(t *testing.T) {
		agents := buildkite.AgentsList{"agent"}
		val := buildkite.CommandStep{
			Agents: &buildkite.Agents{
				AgentsList: &agents,
			},
		}
		CheckResult(t, val, `{"agents":["agent"]}`)
	})

	t.Run("AllowDependencyFailure", func(t *testing.T) {
		allowDependencyFailure := true
		val := buildkite.CommandStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("ArtifactPaths", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			artifactPath := "path"
			val := buildkite.CommandStep{
				ArtifactPaths: &buildkite.CommandStepArtifactPaths{
					String: &artifactPath,
				},
			}
			CheckResult(t, val, `{"artifact_paths":"path"}`)
		})

		t.Run("StringArray", func(t *testing.T) {
			artifactPaths := []string{"one", "two"}
			val := buildkite.CommandStep{
				ArtifactPaths: &buildkite.CommandStepArtifactPaths{
					StringArray: artifactPaths,
				},
			}
			CheckResult(t, val, `{"artifact_paths":["one","two"]}`)
		})
	})

	t.Run("Branches", func(t *testing.T) {
		branches := "branch"
		val := buildkite.CommandStep{
			Branches: &buildkite.Branches{
				String: &branches,
			},
		}
		CheckResult(t, val, `{"branches":"branch"}`)
	})

	t.Run("Cache", func(t *testing.T) {
		cache := "cache"
		val := buildkite.CommandStep{
			Cache: &buildkite.Cache{
				String: &cache,
			},
		}
		CheckResult(t, val, `{"cache":"cache"}`)
	})

	t.Run("CancelOnBuildFailing", func(t *testing.T) {
		cancelOnBuildFailing := true
		val := buildkite.CommandStep{
			CancelOnBuildFailing: &buildkite.CancelOnBuildFailing{
				Bool: &cancelOnBuildFailing,
			},
		}
		CheckResult(t, val, `{"cancel_on_build_failing":true}`)
	})

	t.Run("Command", func(t *testing.T) {
		command := "command"
		val := buildkite.CommandStep{
			Command: &buildkite.CommandStepCommand{
				String: &command,
			},
		}
		CheckResult(t, val, `{"command":"command"}`)
	})

	t.Run("Commands", func(t *testing.T) {
		commands := []string{"one", "two"}
		val := buildkite.CommandStep{
			Commands: &buildkite.CommandStepCommand{
				StringArray: commands,
			},
		}
		CheckResult(t, val, `{"commands":["one","two"]}`)
	})

	t.Run("Concurrency", func(t *testing.T) {
		concurrency := 1
		val := buildkite.CommandStep{
			Concurrency: &concurrency,
		}
		CheckResult(t, val, `{"concurrency":1}`)
	})

	t.Run("ConcurrencyGroup", func(t *testing.T) {
		concurrencyGroup := "group"
		val := buildkite.CommandStep{
			ConcurrencyGroup: &concurrencyGroup,
		}
		CheckResult(t, val, `{"concurrency_group":"group"}`)
	})

	t.Run("ConcurrencyMethod", func(t *testing.T) {
		concurrencyMethod := buildkite.CommandStepConcurrencyMethodValues["ordered"]
		val := buildkite.CommandStep{
			ConcurrencyMethod: &concurrencyMethod,
		}
		CheckResult(t, val, `{"concurrency_method":"ordered"}`)
	})

	t.Run("DependsOn", func(t *testing.T) {
		dependsOn := "step"
		val := buildkite.CommandStep{
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
		}
		CheckResult(t, val, `{"depends_on":"step"}`)
	})

	t.Run("Env", func(t *testing.T) {
		env := buildkite.Env{"foo": "bar"}
		val := buildkite.CommandStep{
			Env: &env,
		}
		CheckResult(t, val, `{"env":{"foo":"bar"}}`)
	})

	t.Run("If", func(t *testing.T) {
		ifValue := "ifValue"
		val := buildkite.CommandStep{
			If: &ifValue,
		}
		CheckResult(t, val, `{"if":"ifValue"}`)
	})

	t.Run("IfChanged", func(t *testing.T) {
		val := buildkite.GroupStep{
			IfChanged: &buildkite.IfChanged{
				String: buildkite.Value("ifChanged"),
			},
		}
		CheckResult(t, val, `{"if_changed":"ifChanged"}`)
	})

	t.Run("Key", func(t *testing.T) {
		key := "key"
		val := buildkite.CommandStep{
			Key: &key,
		}
		CheckResult(t, val, `{"key":"key"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		identifier := "identifier"
		val := buildkite.CommandStep{
			Identifier: &identifier,
		}
		CheckResult(t, val, `{"identifier":"identifier"}`)
	})

	t.Run("Id", func(t *testing.T) {
		id := "id"
		val := buildkite.CommandStep{
			Id: &id,
		}
		CheckResult(t, val, `{"id":"id"}`)
	})

	t.Run("Image", func(t *testing.T) {
		image := "image"
		val := buildkite.CommandStep{
			Image: &image,
		}
		CheckResult(t, val, `{"image":"image"}`)
	})

	t.Run("Label", func(t *testing.T) {
		label := "label"
		val := buildkite.CommandStep{
			Label: &label,
		}
		CheckResult(t, val, `{"label":"label"}`)
	})

	t.Run("Signature", func(t *testing.T) {
		t.Run("Algorithm", func(t *testing.T) {
			algorithm := "algorithm"
			val := buildkite.CommandStep{
				Signature: &buildkite.CommandStepSignature{
					Algorithm: &algorithm,
				},
			}
			CheckResult(t, val, `{"signature":{"algorithm":"algorithm"}}`)
		})

		t.Run("Value", func(t *testing.T) {
			value := "value"
			val := buildkite.CommandStep{
				Signature: &buildkite.CommandStepSignature{
					Value: &value,
				},
			}
			CheckResult(t, val, `{"signature":{"value":"value"}}`)
		})

		t.Run("SignedFields", func(t *testing.T) {
			signedFields := []string{"one", "two"}
			val := buildkite.CommandStep{
				Signature: &buildkite.CommandStepSignature{
					SignedFields: signedFields,
				},
			}
			CheckResult(t, val, `{"signature":{"signed_fields":["one","two"]}}`)
		})

		t.Run("All", func(t *testing.T) {
			algorithm := "algorithm"
			value := "value"
			signedFields := []string{"one", "two"}
			val := buildkite.CommandStep{
				Signature: &buildkite.CommandStepSignature{
					Algorithm:    &algorithm,
					Value:        &value,
					SignedFields: signedFields,
				},
			}
			CheckResult(t, val, `{"signature":{"algorithm":"algorithm","signed_fields":["one","two"],"value":"value"}}`)
		})
	})

	t.Run("Matrix", func(t *testing.T) {
		element := "value"
		list := buildkite.MatrixElementList{
			{
				String: &element,
			},
		}
		val := buildkite.CommandStep{
			Matrix: &buildkite.Matrix{
				MatrixElementList: &list,
			},
		}
		CheckResult(t, val, `{"matrix":["value"]}`)
	})

	t.Run("Name", func(t *testing.T) {
		name := "name"
		val := buildkite.CommandStep{
			Name: &name,
		}
		CheckResult(t, val, `{"name":"name"}`)
	})

	t.Run("Notify", func(t *testing.T) {
		notifySlack := "#channel"
		val := buildkite.CommandStep{
			Notify: &buildkite.CommandStepNotify{
				{
					NotifySlack: &buildkite.NotifySlack{
						Slack: &buildkite.NotifySlackSlack{
							String: &notifySlack,
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"slack":"#channel"}]}`)
	})

	t.Run("Parallelism", func(t *testing.T) {
		parallelism := 2
		val := buildkite.CommandStep{
			Parallelism: &parallelism,
		}
		CheckResult(t, val, `{"parallelism":2}`)
	})

	t.Run("Plugins", func(t *testing.T) {
		val := buildkite.CommandStep{
			Plugins: &buildkite.Plugins{
				PluginsList: &buildkite.PluginsList{
					{
						PluginsList: &buildkite.PluginsListObject{
							"docker": map[string]interface{}{
								"foo": "bar",
							},
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"plugins":[{"docker":{"foo":"bar"}}]}`)
	})

	t.Run("Secrets", func(t *testing.T) {
		t.Run("StringArray", func(t *testing.T) {
			val := buildkite.CommandStep{
				Secrets: &buildkite.Secrets{
					StringArray: []string{"MY_SECRET"},
				},
			}
			CheckResult(t, val, `{"secrets":["MY_SECRET"]}`)
		})

		t.Run("StringArray", func(t *testing.T) {
			val := buildkite.CommandStep{
				Secrets: &buildkite.Secrets{
					Secrets: &buildkite.SecretsObject{"MY_SECRET": "API_TOKEN"},
				},
			}
			CheckResult(t, val, `{"secrets":{"MY_SECRET":"API_TOKEN"}}`)
		})
	})

	t.Run("Softfail", func(t *testing.T) {
		softFail := true
		val := buildkite.CommandStep{
			SoftFail: &buildkite.SoftFail{
				SoftFailEnum: &buildkite.SoftFailEnum{
					Bool: &softFail,
				},
			},
		}
		CheckResult(t, val, `{"soft_fail":true}`)
	})

	t.Run("Retry", func(t *testing.T) {
		t.Run("Automatic", func(t *testing.T) {
			limit := 1
			list := []buildkite.AutomaticRetry{
				{
					Limit: &limit,
				},
			}
			val := buildkite.CommandStep{
				Retry: &buildkite.CommandStepRetry{
					Automatic: &buildkite.CommandStepAutomaticRetry{
						AutomaticRetryList: &list,
					},
				},
			}
			CheckResult(t, val, `{"retry":{"automatic":[{"limit":1}]}}`)
		})

		t.Run("Manual", func(t *testing.T) {
			value := true
			val := buildkite.CommandStep{
				Retry: &buildkite.CommandStepRetry{
					Manual: &buildkite.CommandStepManualRetry{
						CommandStepManualRetryEnum: &buildkite.CommandStepManualRetryEnum{
							Bool: &value,
						},
					},
				},
			}
			CheckResult(t, val, `{"retry":{"manual":true}}`)
		})
	})

	t.Run("Skip", func(t *testing.T) {
		skip := "true"
		val := buildkite.CommandStep{
			Skip: &buildkite.Skip{
				String: &skip,
			},
		}
		CheckResult(t, val, `{"skip":"true"}`)
	})

	t.Run("TimeoutInMinutes", func(t *testing.T) {
		timeout := 2
		val := buildkite.CommandStep{
			TimeoutInMinutes: &timeout,
		}
		CheckResult(t, val, `{"timeout_in_minutes":2}`)
	})

	t.Run("Type", func(t *testing.T) {
		commandType := buildkite.CommandStepTypeValues["command"]
		val := buildkite.CommandStep{
			Type: &commandType,
		}
		CheckResult(t, val, `{"type":"command"}`)
	})

	t.Run("Priority", func(t *testing.T) {
		priority := 1
		val := buildkite.CommandStep{
			Priority: &priority,
		}
		CheckResult(t, val, `{"priority":1}`)
	})

	t.Run("All", func(t *testing.T) {
		agents := buildkite.AgentsList{"agent"}
		allowDependencyFailure := true
		artifactPaths := []string{"one", "two"}
		branches := "branch"
		cache := "cache"
		cancelOnBuildFailing := true
		command := "command"
		commands := []string{"one", "two"}
		concurrency := 1
		concurrencyGroup := "group"
		concurrencyMethod := buildkite.CommandStepConcurrencyMethodValues["ordered"]
		dependsOn := "step"
		env := buildkite.Env{"foo": "bar"}
		ifValue := "ifValue"
		key := "key"
		identifier := "identifier"
		id := "id"
		image := "image"
		label := "label"
		algorithm := "algorithm"
		value := "value"
		signedFields := []string{"one", "two"}
		signature := buildkite.CommandStepSignature{
			Algorithm:    &algorithm,
			Value:        &value,
			SignedFields: signedFields,
		}
		matrixElement := "value"
		matrixList := buildkite.MatrixElementList{
			{
				String: &matrixElement,
			},
		}
		name := "name"
		notifySlack := "#channel"
		parallelism := 2
		softFail := true
		manualRetry := true
		skip := "true"
		timeout := 2
		commandType := buildkite.CommandStepTypeValues["command"]
		priority := 1

		val := buildkite.CommandStep{
			Agents: &buildkite.Agents{
				AgentsList: &agents,
			},
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			ArtifactPaths: &buildkite.CommandStepArtifactPaths{
				StringArray: artifactPaths,
			},
			Branches: &buildkite.Branches{
				String: &branches,
			},
			Cache: &buildkite.Cache{
				String: &cache,
			},
			CancelOnBuildFailing: &buildkite.CancelOnBuildFailing{
				Bool: &cancelOnBuildFailing,
			},
			Command: &buildkite.CommandStepCommand{
				String: &command,
			},
			Commands: &buildkite.CommandStepCommand{
				StringArray: commands,
			},
			Concurrency:       &concurrency,
			ConcurrencyGroup:  &concurrencyGroup,
			ConcurrencyMethod: &concurrencyMethod,
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			Env: &env,
			If:  &ifValue,
			IfChanged: &buildkite.IfChanged{
				String: buildkite.Value("ifChanged"),
			},
			Key:        &key,
			Identifier: &identifier,
			Id:         &id,
			Image:      &image,
			Label:      &label,
			Signature:  &signature,
			Matrix: &buildkite.Matrix{
				MatrixElementList: &matrixList,
			},
			Name: &name,
			Notify: &buildkite.CommandStepNotify{
				{
					NotifySlack: &buildkite.NotifySlack{
						Slack: &buildkite.NotifySlackSlack{
							String: &notifySlack,
						},
					},
				},
			},
			Parallelism: &parallelism,
			Plugins: &buildkite.Plugins{
				PluginsList: &buildkite.PluginsList{
					{
						PluginsList: &buildkite.PluginsListObject{
							"docker": map[string]interface{}{
								"foo": "bar",
							},
						},
					},
				},
			},
			SoftFail: &buildkite.SoftFail{
				SoftFailEnum: &buildkite.SoftFailEnum{
					Bool: &softFail,
				},
			},
			Retry: &buildkite.CommandStepRetry{
				Manual: &buildkite.CommandStepManualRetry{
					CommandStepManualRetryEnum: &buildkite.CommandStepManualRetryEnum{
						Bool: &manualRetry,
					},
				},
			},
			Skip: &buildkite.Skip{
				String: &skip,
			},
			TimeoutInMinutes: &timeout,
			Type:             &commandType,
			Priority:         &priority,
		}
		CheckResult(t, val, `{"agents":["agent"],"allow_dependency_failure":true,"artifact_paths":["one","two"],"branches":"branch","cache":"cache","cancel_on_build_failing":true,"command":"command","commands":["one","two"],"concurrency":1,"concurrency_group":"group","concurrency_method":"ordered","depends_on":"step","env":{"foo":"bar"},"id":"id","identifier":"identifier","if":"ifValue","if_changed":"ifChanged","image":"image","key":"key","label":"label","matrix":["value"],"name":"name","notify":[{"slack":"#channel"}],"parallelism":2,"plugins":[{"docker":{"foo":"bar"}}],"priority":1,"retry":{"manual":true},"signature":{"algorithm":"algorithm","signed_fields":["one","two"],"value":"value"},"skip":"true","soft_fail":true,"timeout_in_minutes":2,"type":"command"}`)
	})
}

func TestNestedCommandStep(t *testing.T) {
	t.Run("Command", func(t *testing.T) {
		command := "command"
		val := buildkite.NestedCommandStep{
			Command: &buildkite.CommandStep{
				Command: &buildkite.CommandStepCommand{
					String: &command,
				},
			},
		}
		CheckResult(t, val, `{"command":{"command":"command"}}`)
	})

	t.Run("Commands", func(t *testing.T) {
		command := "command"
		val := buildkite.NestedCommandStep{
			Commands: &buildkite.CommandStep{
				Command: &buildkite.CommandStepCommand{
					String: &command,
				},
			},
		}
		CheckResult(t, val, `{"commands":{"command":"command"}}`)
	})

	t.Run("Script", func(t *testing.T) {
		command := "command"
		val := buildkite.NestedCommandStep{
			Script: &buildkite.CommandStep{
				Command: &buildkite.CommandStepCommand{
					String: &command,
				},
			},
		}
		CheckResult(t, val, `{"script":{"command":"command"}}`)
	})
}
