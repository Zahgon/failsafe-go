package fallback

import (
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a Fallback.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*fallback[R]
}

var _ policy.Executor[any] = &executor[any]{}

// Apply performs an execution by calling the innerFn, applying a fallback if it fails, and calling post-execute.
func (e *executor[R]) Apply(innerFn func(failsafe.Execution[R]) *common.PolicyResult[R]) func(failsafe.Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Check for cancellation during execution

// Call fallback fn

// Check for cancellation during fallback
