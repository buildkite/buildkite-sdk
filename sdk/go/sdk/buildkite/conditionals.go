package buildkite

import (
	"errors"
	"fmt"

	bkconditional "github.com/buildkite/conditional"
)

// Condition validates a step-aware Buildkite conditional and returns it in the
// pointer form used by SDK `if` fields.
func Condition(expression string) (*string, error) {
	if err := validateConditional(expression, bkconditional.EntryPointBuildConditionWithStep); err != nil {
		return nil, err
	}

	return Value(expression), nil
}

// MustCondition validates a step-aware Buildkite conditional and panics if it
// is invalid.
func MustCondition(expression string) *string {
	condition, err := Condition(expression)
	if err != nil {
		panic(err)
	}

	return condition
}

// ValidateConditionals validates all `if` expressions reachable from the
// pipeline's steps and notifications.
func ValidateConditionals(p Pipeline) error {
	var errs []error

	if p.Notify != nil {
		validateBuildNotifications("notify", *p.Notify, bkconditional.EntryPointBuildNotification, &errs)
	}
	if p.Steps != nil {
		validatePipelineSteps("steps", *p.Steps, &errs)
	}

	return errors.Join(errs...)
}

type conditionalPathError struct {
	Path string
	Err  error
}

func (e *conditionalPathError) Error() string {
	return fmt.Sprintf("%s: %v", e.Path, e.Err)
}

func (e *conditionalPathError) Unwrap() error {
	return e.Err
}

func validateConditional(expression string, entryPoint bkconditional.EntryPoint) error {
	return bkconditional.Validate(expression, bkconditional.Context{EntryPoint: entryPoint})
}

func appendConditionalError(path string, expression *string, entryPoint bkconditional.EntryPoint, errs *[]error) {
	if expression == nil {
		return
	}

	if err := validateConditional(*expression, entryPoint); err != nil {
		*errs = append(*errs, &conditionalPathError{Path: path, Err: err})
	}
}

func validatePipelineSteps(path string, steps PipelineSteps, errs *[]error) {
	for i, step := range steps {
		stepPath := fmt.Sprintf("%s[%d]", path, i)
		switch {
		case step.BlockStep != nil:
			validateBlockStep(stepPath, *step.BlockStep, errs)
		case step.CommandStep != nil:
			validateCommandStep(stepPath, *step.CommandStep, errs)
		case step.GroupStep != nil:
			validateGroupStep(stepPath, *step.GroupStep, errs)
		case step.InputStep != nil:
			validateInputStep(stepPath, *step.InputStep, errs)
		case step.NestedBlockStep != nil:
			validateNestedBlockStep(stepPath, *step.NestedBlockStep, errs)
		case step.NestedCommandStep != nil:
			validateNestedCommandStep(stepPath, *step.NestedCommandStep, errs)
		case step.NestedInputStep != nil:
			validateNestedInputStep(stepPath, *step.NestedInputStep, errs)
		case step.NestedTriggerStep != nil:
			validateNestedTriggerStep(stepPath, *step.NestedTriggerStep, errs)
		case step.NestedWaitStep != nil:
			validateNestedWaitStep(stepPath, *step.NestedWaitStep, errs)
		case step.TriggerStep != nil:
			validateTriggerStep(stepPath, *step.TriggerStep, errs)
		case step.WaitStep != nil:
			validateWaitStep(stepPath, *step.WaitStep, errs)
		}
	}
}

func validateGroupSteps(path string, steps GroupSteps, errs *[]error) {
	for i, step := range steps {
		stepPath := fmt.Sprintf("%s[%d]", path, i)
		switch {
		case step.BlockStep != nil:
			validateBlockStep(stepPath, *step.BlockStep, errs)
		case step.CommandStep != nil:
			validateCommandStep(stepPath, *step.CommandStep, errs)
		case step.InputStep != nil:
			validateInputStep(stepPath, *step.InputStep, errs)
		case step.NestedBlockStep != nil:
			validateNestedBlockStep(stepPath, *step.NestedBlockStep, errs)
		case step.NestedCommandStep != nil:
			validateNestedCommandStep(stepPath, *step.NestedCommandStep, errs)
		case step.NestedInputStep != nil:
			validateNestedInputStep(stepPath, *step.NestedInputStep, errs)
		case step.NestedTriggerStep != nil:
			validateNestedTriggerStep(stepPath, *step.NestedTriggerStep, errs)
		case step.NestedWaitStep != nil:
			validateNestedWaitStep(stepPath, *step.NestedWaitStep, errs)
		case step.TriggerStep != nil:
			validateTriggerStep(stepPath, *step.TriggerStep, errs)
		case step.WaitStep != nil:
			validateWaitStep(stepPath, *step.WaitStep, errs)
		}
	}
}

func validateBlockStep(path string, step BlockStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
}

func validateCommandStep(path string, step CommandStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
	if step.Notify != nil {
		validateCommandNotifications(path+".notify", *step.Notify, errs)
	}
}

func validateGroupStep(path string, step GroupStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
	if step.Notify != nil {
		validateBuildNotifications(path+".notify", *step.Notify, bkconditional.EntryPointStepNotification, errs)
	}
	if step.Steps != nil {
		validateGroupSteps(path+".steps", *step.Steps, errs)
	}
}

func validateInputStep(path string, step InputStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
}

func validateTriggerStep(path string, step TriggerStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
}

func validateWaitStep(path string, step WaitStep, errs *[]error) {
	appendConditionalError(path+".if", step.If, bkconditional.EntryPointBuildConditionWithStep, errs)
}

func validateNestedBlockStep(path string, step NestedBlockStep, errs *[]error) {
	if step.Block != nil {
		validateBlockStep(path+".block", *step.Block, errs)
	}
}

func validateNestedCommandStep(path string, step NestedCommandStep, errs *[]error) {
	if step.Command != nil {
		validateCommandStep(path+".command", *step.Command, errs)
	}
	if step.Commands != nil {
		validateCommandStep(path+".commands", *step.Commands, errs)
	}
	if step.Script != nil {
		validateCommandStep(path+".script", *step.Script, errs)
	}
}

func validateNestedInputStep(path string, step NestedInputStep, errs *[]error) {
	if step.Input != nil {
		validateInputStep(path+".input", *step.Input, errs)
	}
}

func validateNestedTriggerStep(path string, step NestedTriggerStep, errs *[]error) {
	if step.Trigger != nil {
		validateTriggerStep(path+".trigger", *step.Trigger, errs)
	}
}

func validateNestedWaitStep(path string, step NestedWaitStep, errs *[]error) {
	if step.Wait != nil {
		validateWaitStep(path+".wait", *step.Wait, errs)
	}
	if step.Waiter != nil {
		validateWaitStep(path+".waiter", *step.Waiter, errs)
	}
}

func validateBuildNotifications(path string, notifications BuildNotify, entryPoint bkconditional.EntryPoint, errs *[]error) {
	for i, notification := range notifications {
		notificationPath := fmt.Sprintf("%s[%d]", path, i)
		switch {
		case notification.NotifyBasecamp != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyBasecamp.If, entryPoint, errs)
		case notification.NotifyEmail != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyEmail.If, entryPoint, errs)
		case notification.NotifyGithubCommitStatus != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyGithubCommitStatus.If, entryPoint, errs)
		case notification.NotifyPagerduty != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyPagerduty.If, entryPoint, errs)
		case notification.NotifySlack != nil:
			appendConditionalError(notificationPath+".if", notification.NotifySlack.If, entryPoint, errs)
		case notification.NotifyWebhook != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyWebhook.If, entryPoint, errs)
		}
	}
}

func validateCommandNotifications(path string, notifications CommandStepNotify, errs *[]error) {
	for i, notification := range notifications {
		notificationPath := fmt.Sprintf("%s[%d]", path, i)
		switch {
		case notification.NotifyBasecamp != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyBasecamp.If, bkconditional.EntryPointStepNotification, errs)
		case notification.NotifyGithubCommitStatus != nil:
			appendConditionalError(notificationPath+".if", notification.NotifyGithubCommitStatus.If, bkconditional.EntryPointStepNotification, errs)
		case notification.NotifySlack != nil:
			appendConditionalError(notificationPath+".if", notification.NotifySlack.If, bkconditional.EntryPointStepNotification, errs)
		}
	}
}
