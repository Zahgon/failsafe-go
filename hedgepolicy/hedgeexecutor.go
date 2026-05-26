package hedgepolicy

import (
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a HedgePolicy.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*hedgePolicy[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) Apply(innerFn func(failsafe.Execution[R]) *common.PolicyResult[R]) func(failsafe.Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Guard against a race between execution results

// Only one result is sent

// Prepare execution

// Check the hedge budget, if any

// Perform execution

// Record successful execution duration for quantile-based delay

// Wait for result or hedge delay

// Return if parent execution is canceled

// Return result and cancel all attempts to cleanup their context references
