package testutil

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/failsafe-go/failsafe-go"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInvalidState    = errors.New("invalid state")
	ErrConnecting      = errors.New("connection error")

	NoopFn = func() error { return nil }

	GetFalseFn = func() (bool, error) { return false, nil }

	GetTrueFn = func() (bool, error) { return true, nil }
)

type CustomError struct{ Msg string }

func (e CustomError) Error() string { _ = "STUB: not implemented"; return "" }

type CompositeError struct {
	Cause error
}

func (e CompositeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CompositeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type MultiError []error

func (e MultiError) Error() string { _ = "STUB: not implemented"; return "" }

func (e MultiError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

func RunFn(err error) func(failsafe.Execution[any]) error { _ = "STUB: not implemented"; return nil }

func GetFn[R any](result R, err error) func(failsafe.Execution[R]) (R, error) {
	_ = "STUB: not implemented"
	return nil
}

// ErrorNTimesThenReturn returns a stub function that returns the err errorTimes and then returns the results.
// Can be used with Executor.GetWithExecution.
func ErrorNTimesThenReturn[R any](err error, errorTimes int, results ...R) (fn func(failsafe.Execution[R]) (R, error), resetFn func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ErrorNTimesThenPanic returns a stub function that returns the err errorTimes and then panics with the panicValue.
// Can be used with Executor.GetWithExecution.
func ErrorNTimesThenPanic[R any](err error, errorTimes int, panicValue any) func(failsafe.Execution[R]) (R, error) {
	_ = "STUB: not implemented"
	return nil
}

// ErrorNTimesThenError returns a stub function that returns the err errorTimes and then returns the finalError.
// Can be used with Executor.GetWithExecution.
func ErrorNTimesThenError[R any](err error, errorTimes int, finalError error) func(failsafe.Execution[R]) (R, error) {
	_ = "STUB: not implemented"
	return nil
}

func SlowNTimesThenReturn[R any](t *testing.T, slowTimes int, sleepTime time.Duration, delayedResult R, fastResult R) func(failsafe.Execution[R]) (R, error) {
	_ = "STUB: not implemented"
	return nil
}

type TestExecution[R any] struct {
	TheLastResult R
	TheAttempts   int
	TheRetries    int
	TheHedges     int
}

func (e TestExecution[R]) Attempts() int { _ = "STUB: not implemented"; return 0 }

func (e TestExecution[R]) Executions() int { _ = "STUB: not implemented"; return 0 }

func (e TestExecution[R]) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e TestExecution[R]) Retries() int { _ = "STUB: not implemented"; return 0 }

func (e TestExecution[R]) Hedges() int { _ = "STUB: not implemented"; return 0 }

func (e TestExecution[R]) IsFirstAttempt() bool { _ = "STUB: not implemented"; return false }

func (e TestExecution[R]) IsRetry() bool { _ = "STUB: not implemented"; return false }

func (e TestExecution[R]) ElapsedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e TestExecution[R]) IsHedge() bool { _ = "STUB: not implemented"; return false }

func (e TestExecution[R]) LastResult() R { _ = "STUB: not implemented"; return *new(R) }

func (e TestExecution[R]) LastError() error { _ = "STUB: not implemented"; return nil }

func (e TestExecution[R]) AttemptStartTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (e TestExecution[R]) ElapsedAttemptTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e TestExecution[R]) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (e TestExecution[R]) IsCanceled() bool { _ = "STUB: not implemented"; return false }

func (e TestExecution[R]) Canceled() <-chan struct{} { _ = "STUB: not implemented"; return nil }
