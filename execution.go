package failsafe

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/failsafe-go/failsafe-go/common"
)

// ExecutionInfo contains execution info.
type ExecutionInfo interface {
	// Context returns the context configured for the execution, else context.Background if none was configured. For
	// executions involving a timeout or hedge, each attempt will get a separate child context.
	Context() context.Context

	// Attempts returns the number of execution attempts so far, including attempts that are currently in progress and
	// attempts that were blocked before being executed, such as by a CircuitBreaker or RateLimiter. These can include an initial
	// execution along with retries and hedges.
	Attempts() int

	// Executions returns the number of completed executions. Executions that are blocked, such as when a CircuitBreaker is
	// open, are not counted.
	Executions() int

	// Retries returns the number of retries so far, including retries that are currently in progress.
	Retries() int

	// Hedges returns the number of hedges that have been executed so far, including hedges that are currently in progress.
	Hedges() int

	// StartTime returns the time that the initial execution attempt started at.
	StartTime() time.Time

	// ElapsedTime returns the elapsed time since initial execution attempt began.
	ElapsedTime() time.Duration
}

// ExecutionAttempt contains information for an execution attempt.
type ExecutionAttempt[R any] interface {
	ExecutionInfo

	// LastResult returns the result, if any, from the last execution attempt.
	LastResult() R

	// LastError returns the error, if any, from the last execution attempt.
	LastError() error

	// IsFirstAttempt returns true when Attempts is 1, meaning this is the first execution attempt.
	IsFirstAttempt() bool

	// IsRetry returns true when Attempts is > 1, meaning the execution is being retried.
	IsRetry() bool

	// IsHedge returns true when the execution is part of a hedged attempt.
	IsHedge() bool

	// AttemptStartTime returns the time that the most recent execution attempt started at.
	AttemptStartTime() time.Time

	// ElapsedAttemptTime returns the elapsed time since the last execution attempt began.
	ElapsedAttemptTime() time.Duration
}

// Execution contains information about an execution.
type Execution[R any] interface {
	ExecutionAttempt[R]

	// IsCanceled returns whether the execution has been canceled by an external Context or a timeout.Timeout.
	IsCanceled() bool

	// Canceled returns a channel that is closed when the execution is canceled, either by an external Context or a
	// timeout.Timeout.
	Canceled() <-chan struct{}
}

type execution[R any] struct {
	// Shared state across instances
	mu         *sync.Mutex
	startTime  time.Time
	attempts   *atomic.Uint32
	retries    *atomic.Uint32
	hedges     *atomic.Uint32
	executions *atomic.Uint32

	// Partly shared cancellation state
	ctx            context.Context
	cancelFunc     context.CancelFunc
	deferCancel    *atomic.Bool
	canceledResult **common.PolicyResult[R]

	// Per execution state
	attemptStartTime time.Time
	isHedge          bool
	lastResult       R     // The last error that occurred, else the zero value for R.
	lastError        error // The last error that occurred, else nil.
}

var _ Execution[any] = &execution[any]{}
var _ ExecutionInfo = &execution[any]{}

func (e *execution[R]) Attempts() int { _ = "STUB: not implemented"; return 0 }

func (e *execution[R]) Executions() int { _ = "STUB: not implemented"; return 0 }

func (e *execution[R]) Retries() int { _ = "STUB: not implemented"; return 0 }

func (e *execution[R]) Hedges() int { _ = "STUB: not implemented"; return 0 }

func (e *execution[R]) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e *execution[R]) IsFirstAttempt() bool { _ = "STUB: not implemented"; return false }

func (e *execution[R]) IsRetry() bool { _ = "STUB: not implemented"; return false }

func (e *execution[R]) IsHedge() bool { _ = "STUB: not implemented"; return false }

func (e *execution[R]) ElapsedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e *execution[R]) LastResult() R { _ = "STUB: not implemented"; return *new(R) }

func (e *execution[R]) LastError() error { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (e *execution[R]) AttemptStartTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (e *execution[_]) ElapsedAttemptTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e *execution[_]) IsCanceled() bool { _ = "STUB: not implemented"; return false }

func (e *execution[_]) Canceled() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) RecordResult(result *common.PolicyResult[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	// Lock to guard against a race with a Timeout canceling the execution
	return nil
}

func (e *execution[R]) InitializeRetry() *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	// Lock to guard against a race with a Timeout canceling the execution
	return nil
}

func (e *execution[R]) Cancel(result *common.PolicyResult[R]) { _ = "STUB: not implemented"; return }

func (e *execution[R]) DeferCancel() func() { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) IsCanceledWithResult() (bool, *common.PolicyResult[R]) {
	_ = "STUB: not implemented"
	return false, nil
}

// isCanceledWithResult must be locked externally
func (e *execution[R]) isCanceledWithResult() (bool, *common.PolicyResult[R]) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *execution[R]) CopyWithResult(result *common.PolicyResult[R]) Execution[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *execution[R]) CopyForCancellable() Execution[R] { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) CopyForHedge() Execution[R] { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) copy() *execution[R] { _ = "STUB: not implemented"; return nil }

func (e *execution[R]) record() { _ = "STUB: not implemented"; return }

func newExecution[R any](ctx context.Context) *execution[R] { _ = "STUB: not implemented"; return nil }

// executionAnyWrapper adapts execution[R] to execution[any], allowing Policy[any] to be used in compositions with Policy[R].
type executionAnyWrapper[R any] struct {
	*execution[R]
}

func (w *executionAnyWrapper[R]) LastResult() any { _ = "STUB: not implemented"; return *new(any) }

func (w *executionAnyWrapper[R]) RecordResult(anyResult *common.PolicyResult[any]) *common.PolicyResult[any] {
	_ = "STUB: not implemented"
	return nil
}

func (w *executionAnyWrapper[R]) InitializeRetry() *common.PolicyResult[any] {
	_ = "STUB: not implemented"
	return nil
}

func (w *executionAnyWrapper[R]) Cancel(anyResult *common.PolicyResult[any]) {
	_ = "STUB: not implemented"
	return
}

func (w *executionAnyWrapper[R]) IsCanceledWithResult() (bool, *common.PolicyResult[any]) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w *executionAnyWrapper[R]) CopyWithResult(anyResult *common.PolicyResult[any]) Execution[any] {
	_ = "STUB: not implemented"
	return nil
}

func (w *executionAnyWrapper[R]) CopyForCancellable() Execution[any] {
	_ = "STUB: not implemented"
	return nil
}

func (w *executionAnyWrapper[R]) CopyForHedge() Execution[any] {
	_ = "STUB: not implemented"
	return nil
}

func (w *executionAnyWrapper[R]) DeferCancel() func() { _ = "STUB: not implemented"; return nil }

func resultToAny[R any](result *common.PolicyResult[R]) *common.PolicyResult[any] {
	_ = "STUB: not implemented"
	return nil
}

func resultFromAny[R any](anyResult *common.PolicyResult[any]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}
