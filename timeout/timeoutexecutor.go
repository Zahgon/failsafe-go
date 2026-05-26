package timeout

import (
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a Timeout.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*timeout[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) Apply(innerFn func(failsafe.Execution[R]) *common.PolicyResult[R]) func(failsafe.Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	// This func sets up a race between a timeout and the innerFn returning
	return nil
}

// Create child context

// Sets the timeoutResult, overwriting any previously set result for the execution. This is correct, because while an
// execution may have completed, inner policies such as fallbacks may still be processing that result, in which case
// it's still important to interrupt them with a timeout.

// Store result and ctxCancel timeout context if needed

func (e *executor[R]) IsFailure(_ R, err error) bool { _ = "STUB: not implemented"; return false }
