package retrypolicy

import (
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a RetryPolicy.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*retryPolicy[R]

	// Mutable state
	failedAttempts  int
	retriesExceeded bool
	lastDelay       time.Duration // The last backoff delay time
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) Apply(innerFn func(failsafe.Execution[R]) *common.PolicyResult[R]) func(failsafe.Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Perform the execution

// Check for cancellation during execution

// Record result and check for cancellation during PostExecute

// Delay

// Prepare for next iteration and check for cancellation during delay

// Check the retry budget, if any

// OnFailure updates failedAttempts and retriesExceeded, and calls event listeners
func (e *executor[R]) OnFailure(exec policy.ExecutionInternal[R], result *common.PolicyResult[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Call listeners

// getDelay updates lastDelay and returns the new delay
func (e *executor[R]) getDelay(exec failsafe.ExecutionAttempt[R]) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e *executor[R]) getFixedOrRandomDelay(exec failsafe.ExecutionAttempt[R]) time.Duration {
	_ = "STUB: not implemented"

	// Adjust for backoffs
	return *new(time.Duration)
}

func (e *executor[R]) adjustForJitter(delay time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e *executor[R]) adjustForMaxDuration(delay time.Duration, elapsed time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
