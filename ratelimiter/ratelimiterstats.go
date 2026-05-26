package ratelimiter

import (
	"sync"
	"time"

	"github.com/failsafe-go/failsafe-go/internal/util"
)

type stats interface {
	// acquirePermits eagerly acquires requestedPermits and returns the time that must be waited in order to use the permits,
	// else returns -1 if the wait time would exceed the maxWaitTime. A maxWaitTime of -1 indicates no max wait.
	acquirePermits(requestedPermits int, maxWaitTime time.Duration) time.Duration

	reset()
}

// A rate limiter implementation that evenly distributes permits over time, based on the max permits per period. This
// implementation focuses on the interval between permits, and tracks the next interval in which a permit is free.
type smoothStats[R any] struct {
	*config[R]
	stopwatch util.Stopwatch
	mu        sync.Mutex

	// Guarded by mu
	// The amount of time, relative to the start time, that the next permit will be free.
	// Will be a multiple of the config.interval.
	nextFreePermitTime time.Duration
}

func (s *smoothStats[R]) acquirePermits(requestedPermits int, maxWaitTime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// If a permit is currently available

// Time at the start of the current interval

func (s *smoothStats[R]) reset() { _ = "STUB: not implemented"; return }

// A rate limiter implementation that allows bursts of executions, up to the max permits per period. This implementation
// tracks the current period and available permits, which can go into a deficit. A deficit of available permits will
// cause wait times for callers that can be several periods long, depending on the size of the deficit and the number of
// requested permits.
type burstyStats[R any] struct {
	*config[R]
	stopwatch util.Stopwatch
	mu        sync.Mutex

	// Available permits. Can be negative during a deficit.
	// Guarded by mu
	availablePermits int
	currentPeriod    int
}

func (s *burstyStats[R]) acquirePermits(requestedPermits int, maxWaitTime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Update current period and available permits

// Do not wait for an additional period if we're not using any permits from it

// The time to wait until the beginning of the next period that will have free permits

func (s *burstyStats[R]) reset() { _ = "STUB: not implemented"; return }

// exceedsMaxWaitTime returns whether the waitTime would exceed the maxWaitTime, else false if maxWaitTime is -1.
func exceedsMaxWaitTime(waitTime time.Duration, maxWaitTime time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}
