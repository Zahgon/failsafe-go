package bulkhead

import (
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a Bulkhead.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*bulkhead[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) PreExecute(exec policy.ExecutionInternal[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Check for cancellation while waiting for a permit

func (e *executor[R]) PostExecute(_ policy.ExecutionInternal[R], result *common.PolicyResult[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}
