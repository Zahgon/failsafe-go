package adaptivelimiter

import (
	"context"
	"time"
)

// queueingLimiter wraps an adaptiveLimiter and queues some portion of executions when the adaptiveLimiter is full.
type queueingLimiter[R any] struct {
	*adaptiveLimiter[R]
}

func (l *queueingLimiter[R]) AcquirePermit(ctx context.Context) (Permit, error) {
	_ = "STUB: not implemented"
	return *new(Permit), nil
}

// Acquire a permit, blocking if needed

func (l *queueingLimiter[R]) AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) (Permit, error) {
	_ = "STUB: not implemented"
	return *new(Permit), nil
}

// Acquire a permit, blocking if needed

// TryAcquirePermit for a queueingLimiter adds no new behavior since it needs to return immediately, even if the
// semaphore is full, regardless of the queue size.
func (l *queueingLimiter[R]) TryAcquirePermit() (Permit, bool) {
	_ = "STUB: not implemented"
	return *new(Permit), false
}

// CanAcquirePermit returns whether a permit can be acquired based on the semaphore or the queue.
func (l *queueingLimiter[R]) CanAcquirePermit() bool {
	_ = "STUB: not implemented"
	// Check with semaphore
	return false
}

// Check with queue

func (l *queueingLimiter[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

func (l *queueingLimiter[R]) canAcquirePermit(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *queueingLimiter[R]) configRef() *config[R] { _ = "STUB: not implemented"; return nil }

func (l *queueingLimiter[R]) getQueueStats() *queueStats { _ = "STUB: not implemented"; return nil }

// Implements priority.Stats.
type queueStats struct {
	limit              int
	queued             int
	rejectionThreshold int
	maxQueue           int
}

func (s *queueStats) ComputeRejectionRate() float64 { _ = "STUB: not implemented"; return 0 }

func (s *queueStats) DebugLogArgs() []any { _ = "STUB: not implemented"; return nil }
