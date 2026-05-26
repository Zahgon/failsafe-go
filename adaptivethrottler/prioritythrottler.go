package adaptivethrottler

import (
	"context"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/internal"
	"github.com/failsafe-go/failsafe-go/priority"
)

// PriorityThrottler is an adaptive throttler that throttles load probabalistically based on recent failures. When
// throttling is needed, it uses a Prioritizer to determine which priority levels should be rejected, allowing
// higher-priority executions to proceed while shedding lower-priority load.
//
// R is the execution result type. This type is concurrency safe.
type PriorityThrottler[R any] interface {
	failsafe.Policy[R]
	Metrics

	// AcquirePermit attempts to acquire a permit for an execution at the priority or level contained in the context,
	// returning ErrExceeded if one could not be acquired. A priority must be stored in the context using the PriorityKey,
	// or a level must be stored in the context using the LevelKey. The priority or level must be greater than the current
	// rejection threshold for admission. Levels must be between 0 and 499.
	//
	// Example usage:
	//   ctx := priority.ContextWithPriority(context.Background(), priority.High)
	//   permit, err := throttler.AcquirePermit(ctx)
	AcquirePermit(ctx context.Context) error

	// AcquirePermitWithPriority attempts to acquire a permit for an execution at the given priority, returning ErrExceeded
	// if one could not be acquired. The priority must be greater than the current rejection threshold for
	// admission.
	AcquirePermitWithPriority(priority priority.Priority) error

	// AcquirePermitWithLevel attempts to acquire a permit for an execution at the given level, returning ErrExceeded if one
	// could not be acquired. The level must be greater than the current rejection threshold for admission.
	AcquirePermitWithLevel(level int) error

	// TryAcquirePermit attempts to acquire a permit for an execution at the priority or level contained in the context,
	// returning one could be acquired. A priority must be stored in the context using the PriorityKey, or a level must be
	// stored in the context using the LevelKey. The priority or level must be greater than the current rejection threshold
	// for admission. Levels must be between 0 and 499.
	//
	// Example usage:
	//   ctx := priority.ContextWithPriority(context.Background(), priority.High)
	//   permit, err := throttler.AcquirePermit(ctx)
	TryAcquirePermit(ctx context.Context) bool

	// TryAcquirePermitWithPriority attempts to acquire a permit for an execution at the given priority, returning whether
	// one could be acquired. The priority must be greater than the current rejection threshold for admission.
	TryAcquirePermitWithPriority(priority priority.Priority) bool

	// TryAcquirePermitWithLevel attempts to acquire a permit for an execution at the given level, returning whether one
	// could be acquired. The level must be greater than the current rejection threshold for admission.
	TryAcquirePermitWithLevel(level int) bool

	// RecordResult records an execution result as a success or failure based on the failure handling configuration.
	RecordResult(result R)

	// RecordError records an error as a success or failure based on the failure handling configuration.
	RecordError(err error)

	// RecordSuccess records an execution success.
	RecordSuccess()

	// RecordFailure records an execution failure.
	RecordFailure()
}

type priorityThrottler[R any] struct {
	*adaptiveThrottler[R]
	prioritizer *internal.BasePrioritizer[*throttlerStats]
}

func (*priorityThrottler[R]) ResultAgnostic() { _ = "STUB: not implemented"; return }

func (t *priorityThrottler[R]) AcquirePermit(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *priorityThrottler[R]) AcquirePermitWithPriority(priority priority.Priority) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *priorityThrottler[R]) AcquirePermitWithLevel(level int) error {
	_ = "STUB: not implemented"
	// Try to acquire through prioritizer
	return nil
}

// Maintain min flow to prevent starvation

func (t *priorityThrottler[R]) TryAcquirePermit(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *priorityThrottler[R]) TryAcquirePermitWithPriority(priority priority.Priority) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *priorityThrottler[R]) TryAcquirePermitWithLevel(level int) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *priorityThrottler[R]) RejectionRate() float64 { _ = "STUB: not implemented"; return 0 }

func (t *priorityThrottler[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

// Implements Stats for throttler statistics.
type throttlerStats struct {
	executions       float64
	rejectionRate    float64
	maxRejectionRate float64
}

func (s *throttlerStats) ComputeRejectionRate() float64 { _ = "STUB: not implemented"; return 0 }

func (s *throttlerStats) DebugLogArgs() []any {
	_ = "STUB: not implemented"

	// Must be locked externally
	return nil
}

func (t *priorityThrottler[R]) getThrottlerStats() *throttlerStats {
	_ = "STUB: not implemented"
	return nil
}
