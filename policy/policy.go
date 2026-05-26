package policy

import (
	"time"

	"github.com/failsafe-go/failsafe-go"
)

type key int

// CheckOnlyKey is a key to use with a Context that indicates a policy should check if capacity is available without
// actually reserving it.
const CheckOnlyKey key = 0

// BaseFailurePolicy provides a base for implementing FailurePolicyBuilder.
type BaseFailurePolicy[R any] struct {
	// Indicates whether errors are checked by a configured failure condition
	errorsChecked bool
	// Conditions that determine whether an execution is a failure
	failureConditions []func(result R, err error) bool
	onSuccess         func(failsafe.ExecutionEvent[R])
	onFailure         func(failsafe.ExecutionEvent[R])
}

func (p *BaseFailurePolicy[R]) HandleErrors(errs ...error) { _ = "STUB: not implemented"; return }

func (p *BaseFailurePolicy[R]) HandleErrorTypes(errs ...any) { _ = "STUB: not implemented"; return }

func (p *BaseFailurePolicy[R]) HandleResult(result R) { _ = "STUB: not implemented"; return }

func (p *BaseFailurePolicy[R]) HandleIf(predicate func(R, error) bool) {
	_ = "STUB: not implemented"
	return
}

func (p *BaseFailurePolicy[R]) OnSuccess(listener func(event failsafe.ExecutionEvent[R])) {
	_ = "STUB: not implemented"
	return
}

func (p *BaseFailurePolicy[R]) OnFailure(listener func(event failsafe.ExecutionEvent[R])) {
	_ = "STUB: not implemented"
	return
}

func (p *BaseFailurePolicy[R]) IsFailure(result R, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// Fail by default if an error exists and was not checked by a condition

// BaseDelayablePolicy provides a base for implementing DelayablePolicyBuilder.
type BaseDelayablePolicy[R any] struct {
	Delay     time.Duration
	DelayFunc failsafe.DelayFunc[R]
}

func (d *BaseDelayablePolicy[R]) WithDelay(delay time.Duration) { _ = "STUB: not implemented"; return }

func (d *BaseDelayablePolicy[R]) WithDelayFunc(delayFunc failsafe.DelayFunc[R]) {
	_ = "STUB: not implemented"
	return

	// ComputeDelay returns a computed delay else -1 if no delay could be computed.
}

func (d *BaseDelayablePolicy[R]) ComputeDelay(exec failsafe.ExecutionAttempt[R]) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// BaseAbortablePolicy provides a base for implementing policies that can be aborted or canceled.
type BaseAbortablePolicy[R any] struct {
	// Conditions that determine whether the policy should be aborted
	abortConditions []func(result R, err error) bool
}

func (c *BaseAbortablePolicy[R]) AbortOnResult(result R) { _ = "STUB: not implemented"; return }

func (c *BaseAbortablePolicy[R]) AbortOnErrors(errs ...error) { _ = "STUB: not implemented"; return }

func (c *BaseAbortablePolicy[R]) AbortOnErrorTypes(errs ...any) { _ = "STUB: not implemented"; return }

func (c *BaseAbortablePolicy[R]) AbortIf(predicate func(R, error) bool) {
	_ = "STUB: not implemented"
	return
}

func (c *BaseAbortablePolicy[R]) IsConfigured() bool { _ = "STUB: not implemented"; return false }

func (c *BaseAbortablePolicy[R]) IsAbortable(result R, err error) bool {
	_ = "STUB: not implemented"
	return false
}
