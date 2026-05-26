package circuitbreaker

import (
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/internal/util"
)

// The default number of buckets to aggregate time-based stats into.
const defaultBucketCount = 10

func newStats[R any](config *config[R], supportsTimeBased bool, capacity uint) util.ExecutionStats {
	_ = "STUB: not implemented"
	return *new(util.ExecutionStats)
}

// State of a CircuitBreaker.
// Implementations are not concurrency safe and must be guarded externally.
type circuitState[R any] interface {
	util.ExecutionStats
	state() State
	remainingDelay() time.Duration
	tryAcquirePermit() bool
	checkThresholdAndReleasePermit(exec failsafe.Execution[R])
}

type closedState[R any] struct {
	breaker *circuitBreaker[R]
	util.ExecutionStats
}

func newClosedState[R any](breaker *circuitBreaker[R]) *closedState[R] {
	_ = "STUB: not implemented"
	return nil
}

func (s *closedState[R]) state() State { _ = "STUB: not implemented"; return *new(State) }

func (s *closedState[R]) remainingDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *closedState[R]) tryAcquirePermit() bool {
	_ = "STUB: not implemented"

	// Checks to see if the executions and failure thresholds have been exceeded, opening the circuit if so.
	return false
}

func (s *closedState[R]) checkThresholdAndReleasePermit(exec failsafe.Execution[R]) {
	_ = "STUB: not implemented"
	// Execution threshold can only be set for time based thresholding
	return
}

// Failure rate threshold can only be set for time based thresholding

type openState[R any] struct {
	breaker *circuitBreaker[R]
	util.ExecutionStats
	startTime int64
	delay     time.Duration
}

func newOpenState[R any](breaker *circuitBreaker[R], previousState circuitState[R], delay time.Duration) *openState[R] {
	_ = "STUB: not implemented"
	return nil
}

func (s *openState[R]) state() State { _ = "STUB: not implemented"; return *new(State) }

func (s *openState[R]) remainingDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *openState[R]) tryAcquirePermit() bool { _ = "STUB: not implemented"; return false }

func (s *openState[R]) checkThresholdAndReleasePermit(_ failsafe.Execution[R]) {
	_ = "STUB: not implemented"
	return
}

type halfOpenState[R any] struct {
	breaker *circuitBreaker[R]
	util.ExecutionStats
	permittedExecutions uint
}

func newHalfOpenState[R any](breaker *circuitBreaker[R]) *halfOpenState[R] {
	_ = "STUB: not implemented"
	return nil
}

func (s *halfOpenState[R]) state() State { _ = "STUB: not implemented"; return *new(State) }

func (s *halfOpenState[R]) remainingDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *halfOpenState[R]) tryAcquirePermit() bool { _ = "STUB: not implemented"; return false }

/*
Checks to determine if a threshold has been met and the circuit should be opened or closed.
  - If a success threshold is configured, the circuit is opened or closed based on whether the ratio was exceeded.
  - Else the circuit is opened or closed based on whether the failure threshold was exceeded.

A permit is released before returning.
*/
func (s *halfOpenState[R]) checkThresholdAndReleasePermit(exec failsafe.Execution[R]) {
	_ = "STUB: not implemented"
	return
}

// Failure rate threshold can only be set for time based thresholding

// Execution threshold can only be set for time based thresholding
