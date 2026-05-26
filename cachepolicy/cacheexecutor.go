package cachepolicy

import (
	"context"

	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

// executor is a policy.Executor that handles failures according to a CachePolicy.
type executor[R any] struct {
	policy.BaseExecutor[R]
	*cachePolicy[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) PreExecute(exec policy.ExecutionInternal[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) PostExecute(exec policy.ExecutionInternal[R], er *common.PolicyResult[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) getCacheKey(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
