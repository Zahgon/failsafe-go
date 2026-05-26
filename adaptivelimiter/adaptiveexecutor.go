package adaptivelimiter

import (
	"context"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/common"
	"github.com/failsafe-go/failsafe-go/policy"
)

type blockingLimiter[R any] interface {
	canAcquirePermit(ctx context.Context) bool

	AcquirePermit(ctx context.Context) (Permit, error)

	AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) (Permit, error)

	configRef() *config[R]
}

// executor is a policy.Executor that handles failures according to an AdaptiveLimiter or PriorityLimiter.
type executor[R any] struct {
	policy.BaseExecutor[R]
	blockingLimiter[R]
}

var _ policy.Executor[any] = &executor[any]{}

func (e *executor[R]) Apply(innerFn func(failsafe.Execution[R]) *common.PolicyResult[R]) func(failsafe.Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Check for cancellation while waiting for a permit

// Handle exceeded

// Check for cancellation during execution
