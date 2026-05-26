package testutil

import (
	"context"
	"testing"

	"github.com/failsafe-go/failsafe-go"
)

type WhenRun[R any] func(execution failsafe.Execution[R]) error
type WhenGet[R any] func(execution failsafe.Execution[R]) (R, error)

type Resetable interface {
	Reset()
}

type Tester[R any] struct {
	T         *testing.T
	BeforeFn  func()
	AfterFn   func()
	ContextFn func() context.Context
	Executor  failsafe.Executor[R]
	run       WhenRun[R]
	get       WhenGet[R]
}

func Test[R any](t *testing.T) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) Before(fn func()) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) After(fn func()) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) Context(fn func() context.Context) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) Reset(stats ...Resetable) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) With(policies ...failsafe.Policy[R]) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) WithAny(policy failsafe.ResultAgnosticPolicy[any]) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) Compose(policy failsafe.Policy[R]) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) ComposeAny(policy failsafe.ResultAgnosticPolicy[any]) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) WithExecutor(executor failsafe.Executor[R]) *Tester[R] {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tester[R]) Run(when WhenRun[R]) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) Get(when WhenGet[R]) *Tester[R] { _ = "STUB: not implemented"; return nil }

func (t *Tester[R]) Exec(async bool) (R, error) { _ = "STUB: not implemented"; return *new(R), nil }

func (t *Tester[R]) AssertSuccess(expectedAttempts int, expectedExecutions int, expectedResult R, then ...func()) {
	_ = "STUB: not implemented"
	return
}

func (t *Tester[R]) AssertSuccessError(expectedAttempts int, expectedExecutions int, expectedError error, then ...func()) {
	_ = "STUB: not implemented"
	return
}

func (t *Tester[R]) AssertFailure(expectedAttempts int, expectedExecutions int, expectedError error, then ...func()) {
	_ = "STUB: not implemented"
	return
}

func (t *Tester[R]) AssertFailureAs(expectedAttempts int, expectedExecutions int, expectedError error, then ...func()) {
	_ = "STUB: not implemented"
	return
}

func (t *Tester[R]) exec(async bool) (R, error, AssertFunc[R]) {
	_ = "STUB: not implemented"
	return *new(R), nil, nil
}

func (t *Tester[R]) assertResult(expectedAttempts int, expectedExecutions int, expectedResult R, expectedError error, expectedSuccess bool, errorAs bool, then ...func()) {
	_ = "STUB: not implemented"
	return
}

// Run sync

// Run async

type AssertFunc[R any] func(expectedAttempts int, expectedExecutions int, expectedResult R, result R, expectedErr error, err error, expectedSuccess bool, expectedFailure bool, errorAs bool, thens ...func())

func PrepareTest[R any](t *testing.T, beforeFn func(), contextFn func() context.Context, executor failsafe.Executor[R]) (executorFn func() failsafe.Executor[R], assertFn AssertFunc[R]) {
	_ = "STUB: not implemented"
	return nil, nil
}
