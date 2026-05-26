package failsafe

import (
	"errors"
	"sync/atomic"

	"github.com/failsafe-go/failsafe-go/common"
)

// ErrExecutionCanceled indicates that an execution was canceled by ExecutionResult.Cancel.
var ErrExecutionCanceled = errors.New("execution canceled")

// ExecutionResult provides the result of an asynchronous execution.
type ExecutionResult[R any] interface {
	// Done is a channel that is closed when the execution is done and the result can be retrieved via Get, Result, or Error.
	Done() <-chan any

	// IsDone returns whether the execution is done and the result can be retrieved via Get.
	IsDone() bool

	// Get returns the execution result and error, else the default values, blocking until the execution is done.
	Get() (R, error)

	// Result returns the execution result else its default value, blocking until the execution is done.
	Result() R

	// Error returns the execution error else nil, blocking until the execution is done.
	Error() error

	// Cancel cancels the execution if it is not already done, with ErrExecutionCanceled as the error. If a Context was
	// configured with the execution, a child context will be created for the execution and canceled as well.
	Cancel()
}

type executionResult[R any] struct {
	*execution[R]
	cancelFunc func()
	doneChan   chan any
	done       atomic.Bool
	result     atomic.Pointer[*common.PolicyResult[R]]
}

func (e *executionResult[R]) record(result *common.PolicyResult[R]) {
	_ = "STUB: not implemented"
	return
}

func (e *executionResult[R]) Done() <-chan any { _ = "STUB: not implemented"; return nil }

func (e *executionResult[R]) IsDone() bool { _ = "STUB: not implemented"; return false }

func (e *executionResult[R]) Get() (R, error) { _ = "STUB: not implemented"; return *new(R), nil }

func (e *executionResult[R]) Result() R { _ = "STUB: not implemented"; return *new(R) }

func (e *executionResult[R]) Error() error { _ = "STUB: not implemented"; return nil }

func (e *executionResult[R]) Cancel() {
	_ = "STUB: not implemented"
	// Propagate cancelation to contexts
	return
}
