package adaptivethrottler

import (
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to an AdaptiveThrottler.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*adaptiveThrottler[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) PreExecute(_ policy.ExecutionInternal[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) OnSuccess(exec policy.ExecutionInternal[R], result *common.PolicyResult[R]) {
	_ = "STUB: not implemented"
	return
}

func (e *executor[R]) OnFailure(exec policy.ExecutionInternal[R], result *common.PolicyResult[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}
