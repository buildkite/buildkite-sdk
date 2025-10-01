package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func TestGroupStep(t *testing.T) {
	t.Run("DependsOn", func(t *testing.T) {
		dependsOn := "step"
		val := buildkite.GroupStep{
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
		}
		CheckResult(t, val, `{"depends_on":"step"}`)
	})

	t.Run("Group", func(t *testing.T) {
		group := "group"
		val := buildkite.GroupStep{
			Group: &group,
		}
		CheckResult(t, val, `{"group":"group"}`)
	})

	t.Run("If", func(t *testing.T) {
		ifValue := "if"
		val := buildkite.GroupStep{
			If: &ifValue,
		}
		CheckResult(t, val, `{"if":"if"}`)
	})

	t.Run("IfChanged", func(t *testing.T) {
		ifChanged := "ifChanged"
		val := buildkite.GroupStep{
			IfChanged: &ifChanged,
		}
		CheckResult(t, val, `{"if_changed":"ifChanged"}`)
	})

	t.Run("Key", func(t *testing.T) {
		key := "key"
		val := buildkite.GroupStep{
			Key: &key,
		}
		CheckResult(t, val, `{"key":"key"}`)
	})

	t.Run("Identifier", func(t *testing.T) {
		identifier := "identifier"
		val := buildkite.GroupStep{
			Identifier: &identifier,
		}
		CheckResult(t, val, `{"identifier":"identifier"}`)
	})

	t.Run("Id", func(t *testing.T) {
		id := "id"
		val := buildkite.GroupStep{
			Id: &id,
		}
		CheckResult(t, val, `{"id":"id"}`)
	})

	t.Run("Label", func(t *testing.T) {
		label := "label"
		val := buildkite.GroupStep{
			Label: &label,
		}
		CheckResult(t, val, `{"label":"label"}`)
	})

	t.Run("Name", func(t *testing.T) {
		name := "name"
		val := buildkite.GroupStep{
			Name: &name,
		}
		CheckResult(t, val, `{"name":"name"}`)
	})

	t.Run("AllowDependencyFailure", func(t *testing.T) {
		allowDependencyFailure := true
		val := buildkite.GroupStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true}`)
	})

	t.Run("Notify", func(t *testing.T) {
		notifyWebhook := "url"
		val := buildkite.GroupStep{
			Notify: &buildkite.BuildNotify{
				{
					NotifyWebhook: &buildkite.NotifyWebhook{
						Webhook: &notifyWebhook,
					},
				},
			},
		}
		CheckResult(t, val, `{"notify":[{"webhook":"url"}]}`)
	})

	t.Run("Skip", func(t *testing.T) {
		skip := true
		val := buildkite.GroupStep{
			Skip: &buildkite.Skip{
				Bool: &skip,
			},
		}
		CheckResult(t, val, `{"skip":true}`)
	})

	t.Run("Steps", func(t *testing.T) {
		t.Run("BlockStep", func(t *testing.T) {
			block := "blockLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						BlockStep: &buildkite.BlockStep{
							Block: &block,
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"block":"blockLabel"}]}`)
		})

		t.Run("NestedBlockStep", func(t *testing.T) {
			block := "blockLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						NestedBlockStep: &buildkite.NestedBlockStep{
							Block: &buildkite.BlockStep{
								Block: &block,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"block":{"block":"blockLabel"}}]}`)
		})

		t.Run("StringBlockStep", func(t *testing.T) {
			block := buildkite.StringBlockStep("block")
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						StringBlockStep: &block,
					},
				},
			}
			CheckResult(t, val, `{"steps":["block"]}`)
		})

		t.Run("InputStep", func(t *testing.T) {
			input := "inputLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						InputStep: &buildkite.InputStep{
							Input: &input,
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"input":"inputLabel"}]}`)
		})

		t.Run("NestedInputStep", func(t *testing.T) {
			input := "inputLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						NestedInputStep: &buildkite.NestedInputStep{
							Input: &buildkite.InputStep{
								Input: &input,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"input":{"input":"inputLabel"}}]}`)
		})

		t.Run("StringInputStep", func(t *testing.T) {
			input := buildkite.StringInputStep("input")
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						StringInputStep: &input,
					},
				},
			}
			CheckResult(t, val, `{"steps":["input"]}`)
		})

		t.Run("CommandStep", func(t *testing.T) {
			command := "command"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						CommandStep: &buildkite.CommandStep{
							Command: &buildkite.CommandStepCommand{
								String: &command,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"command":"command"}]}`)
		})

		t.Run("NestedCommandStep", func(t *testing.T) {
			command := "command"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						NestedCommandStep: &buildkite.NestedCommandStep{
							Command: &buildkite.CommandStep{
								Command: &buildkite.CommandStepCommand{
									String: &command,
								},
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"command":{"command":"command"}}]}`)
		})

		t.Run("WaitStep", func(t *testing.T) {
			wait := "waitLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						WaitStep: &buildkite.WaitStep{
							Wait: &wait,
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"wait":"waitLabel"}]}`)
		})

		t.Run("NestedWaitStep", func(t *testing.T) {
			wait := "waitLabel"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						NestedWaitStep: &buildkite.NestedWaitStep{
							Wait: &buildkite.WaitStep{
								Wait: &wait,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"wait":{"wait":"waitLabel"}}]}`)
		})

		t.Run("StringWaitStep", func(t *testing.T) {
			wait := buildkite.StringWaitStep("wait")
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						StringWaitStep: &wait,
					},
				},
			}
			CheckResult(t, val, `{"steps":["wait"]}`)
		})

		t.Run("TriggerStep", func(t *testing.T) {
			trigger := "trigger"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						TriggerStep: &buildkite.TriggerStep{
							Trigger: &trigger,
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"trigger":"trigger"}]}`)
		})

		t.Run("NestedTriggerStep", func(t *testing.T) {
			trigger := "trigger"
			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						NestedTriggerStep: &buildkite.NestedTriggerStep{
							Trigger: &buildkite.TriggerStep{
								Trigger: &trigger,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"trigger":{"trigger":"trigger"}}]}`)
		})

		t.Run("MultipleStepTypes", func(t *testing.T) {
			block := "blockLabel"
			nestedBlock := "nestedBlockLabel"
			stringBlock := buildkite.StringBlockStep("block")
			input := "inputLabel"
			nestedInput := "nestedInputLabel"
			stringInput := buildkite.StringInputStep("input")
			command := "command"
			nestedCommand := "nestedCommand"
			wait := "waitLabel"
			nestedWait := "nestedWaitLabel"
			stringWait := buildkite.StringWaitStep("wait")
			trigger := "triggerLabel"
			nestedTrigger := "nestedTriggerLabel"

			val := buildkite.GroupStep{
				Steps: &buildkite.GroupSteps{
					{
						BlockStep: &buildkite.BlockStep{
							Block: &block,
						},
					},
					{
						NestedBlockStep: &buildkite.NestedBlockStep{
							Block: &buildkite.BlockStep{
								Block: &nestedBlock,
							},
						},
					},
					{
						StringBlockStep: &stringBlock,
					},
					{
						InputStep: &buildkite.InputStep{
							Input: &input,
						},
					},
					{
						NestedInputStep: &buildkite.NestedInputStep{
							Input: &buildkite.InputStep{
								Input: &nestedInput,
							},
						},
					},
					{
						StringInputStep: &stringInput,
					},
					{
						CommandStep: &buildkite.CommandStep{
							Command: &buildkite.CommandStepCommand{
								String: &command,
							},
						},
					},
					{
						NestedCommandStep: &buildkite.NestedCommandStep{
							Command: &buildkite.CommandStep{
								Command: &buildkite.CommandStepCommand{
									String: &nestedCommand,
								},
							},
						},
					},
					{
						WaitStep: &buildkite.WaitStep{
							Wait: &wait,
						},
					},
					{
						NestedWaitStep: &buildkite.NestedWaitStep{
							Wait: &buildkite.WaitStep{
								Wait: &nestedWait,
							},
						},
					},
					{
						StringWaitStep: &stringWait,
					},
					{
						TriggerStep: &buildkite.TriggerStep{
							Trigger: &trigger,
						},
					},
					{
						NestedTriggerStep: &buildkite.NestedTriggerStep{
							Trigger: &buildkite.TriggerStep{
								Trigger: &nestedTrigger,
							},
						},
					},
				},
			}
			CheckResult(t, val, `{"steps":[{"block":"blockLabel"},{"block":{"block":"nestedBlockLabel"}},"block",{"input":"inputLabel"},{"input":{"input":"nestedInputLabel"}},"input",{"command":"command"},{"command":{"command":"nestedCommand"}},{"wait":"waitLabel"},{"wait":{"wait":"nestedWaitLabel"}},"wait",{"trigger":"triggerLabel"},{"trigger":{"trigger":"nestedTriggerLabel"}}]}`)
		})
	})

	t.Run("All", func(t *testing.T) {
		dependsOn := "step"
		group := "group"
		ifValue := "if"
		ifChanged := "ifChanged"
		key := "key"
		identifier := "identifier"
		id := "id"
		label := "label"
		name := "name"
		allowDependencyFailure := true
		notifyWebhook := "url"
		skip := true
		commandStepCommand := "command"

		val := buildkite.GroupStep{
			AllowDependencyFailure: &buildkite.AllowDependencyFailure{
				Bool: &allowDependencyFailure,
			},
			DependsOn: &buildkite.DependsOn{
				String: &dependsOn,
			},
			Group:      &group,
			If:         &ifValue,
			IfChanged:  &ifChanged,
			Key:        &key,
			Identifier: &identifier,
			Id:         &id,
			Label:      &label,
			Name:       &name,
			Notify: &buildkite.BuildNotify{
				{
					NotifyWebhook: &buildkite.NotifyWebhook{
						Webhook: &notifyWebhook,
					},
				},
			},
			Skip: &buildkite.Skip{
				Bool: &skip,
			},
			Steps: &buildkite.GroupSteps{
				{
					CommandStep: &buildkite.CommandStep{
						Command: &buildkite.CommandStepCommand{
							String: &commandStepCommand,
						},
					},
				},
			},
		}
		CheckResult(t, val, `{"allow_dependency_failure":true,"depends_on":"step","group":"group","id":"id","identifier":"identifier","if":"if","if_changed":"ifChanged","key":"key","label":"label","name":"name","notify":[{"webhook":"url"}],"skip":true,"steps":[{"command":"command"}]}`)
	})
}
