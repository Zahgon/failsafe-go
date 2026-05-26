package adaptivethrottler

import (
	"errors"
	"sync"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/internal/util"
	"github.com/failsafe-go/failsafe-go/policy"
	"github.com/failsafe-go/failsafe-go/priority"
)

// ErrExceeded is returned when an execution exceeds the current failure rate.
var ErrExceeded = errors.New("failure rate exceeded")

const executionPadding = 1

// AdaptiveThrottler throttles load probabalistically based on recent failures.
// This approach is described in the Google SRE book: https://sre.google/sre-book/handling-overload/#client-side-throttling-a7sYUg
type AdaptiveThrottler[R any] interface {
	failsafe.ResultAgnosticPolicy[R]
	Metrics

	// AcquirePermit attempts to acquire a permit to perform an execution via the throttler, returning ErrExceeded if one
	// could not be acquired.
	AcquirePermit() error

	// TryAcquirePermit attempts to acquire a permit to perform an execution via the throttler, returning whether one could be acquired.
	TryAcquirePermit() bool

	// RecordResult records an execution result as a success or failure based on the failure handling configuration.
	RecordResult(result R)

	// RecordError records an error as a success or failure based on the failure handling configuration.
	RecordError(err error)

	// RecordSuccess records an execution success.
	RecordSuccess()

	// RecordFailure records an execution failure.
	RecordFailure()
}

type Metrics interface {
	// RejectionRate returns the current rate, from 0 to 1, at which executions will be rejected, based on recent failures.
	RejectionRate() float64
}

/*
Builder builds AdaptiveThrottler instances.

This type is not concurrency safe.
*/
type Builder[R any] interface {
	failsafe.FailurePolicyBuilder[Builder[R], R]

	// WithFailureRateThreshold configures the failure rate threshold and thresholding period for the throttler. The
	// throttler will increase rejection probability when the failure rate exceeds this threshold over the specified time
	// period. The number of executions must also exceed the executionThreshold within the thresholdingPeriod
	// before any executions will be rejected.
	// Panics if failureRateThreshold < 0 or > 1.
	WithFailureRateThreshold(failureRateThreshold float64, executionThreshold uint, thresholdingPeriod time.Duration) Builder[R]

	// WithMaxRejectionRate configures the max allowed rejection rate, which defaults to .9.
	// Panics if maxRejectionRate < 0 or > 1.
	WithMaxRejectionRate(maxRejectionRate float64) Builder[R]

	// Build returns a new AdaptiveThrottler using the builder's configuration.
	Build() AdaptiveThrottler[R]

	// BuildPrioritized returns a new PrioritizedThrottler using the builder's configuration. This prioritizes rejections of
	// executions when throttling occurs. Rejections are performed using the Prioritizer, which sets a rejection threshold
	// based on all the throttlers being used by the Prioritizer. The Prioritizer can and should be shared across all
	// throttler instances that need to coordinate prioritization.
	//
	// Prioritized rejection is disabled by default, which means no executions will block when the throttler is full.
	BuildPrioritized(prioritizer priority.Prioritizer) PriorityThrottler[R]
}

type config[R any] struct {
	*policy.BaseFailurePolicy[R]

	maxRejectionRate     float64
	successRateThreshold float64
	executionThreshold   uint
	thresholdingPeriod   time.Duration
}

var _ Builder[any] = &config[any]{}

// NewWithDefaults returns a new AdaptiveThrottler with a failureRateThreshold of .1, a thresholdingPeriod of 1 minute,
// and a maxRejectionRate of .9. To configure additional options on am AdaptiveThrottler, use NewBuilder() instead.
func NewWithDefaults[R any]() AdaptiveThrottler[R] { _ = "STUB: not implemented"; return nil }

// NewBuilder returns an AdaptiveThrottler builder that defaults to a failureRateThreshold of .1, a thresholdingPeriod of 1 minute,
// and a maxRejectionRate of .9.
func NewBuilder[R any]() Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) HandleErrors(errs ...error) Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) HandleErrorTypes(errs ...any) Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) HandleResult(result R) Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) HandleIf(predicate func(R, error) bool) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnSuccess(listener func(event failsafe.ExecutionEvent[R])) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnFailure(listener func(event failsafe.ExecutionEvent[R])) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithFailureRateThreshold(failureRateThreshold float64, executionThreshold uint, thresholdingPeriod time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxRejectionRate(maxRejectionRate float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) Build() AdaptiveThrottler[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) BuildPrioritized(p priority.Prioritizer) PriorityThrottler[R] {
	_ = "STUB: not implemented"
	return nil
}

type adaptiveThrottler[R any] struct {
	config[R]
	mu sync.Mutex

	// Guarded by mu
	util.ExecutionStats
	rejectionRate float64
}

func (*adaptiveThrottler[R]) ResultAgnostic() { _ = "STUB: not implemented"; return }

func (t *adaptiveThrottler[R]) AcquirePermit() error { _ = "STUB: not implemented"; return nil }

// Check for successful acquisition

func (t *adaptiveThrottler[R]) TryAcquirePermit() bool { _ = "STUB: not implemented"; return false }

func (t *adaptiveThrottler[R]) RejectionRate() float64 { _ = "STUB: not implemented"; return 0 }

func (t *adaptiveThrottler[R]) RecordFailure() { _ = "STUB: not implemented"; return }

func (t *adaptiveThrottler[R]) RecordError(err error) { _ = "STUB: not implemented"; return }

func (t *adaptiveThrottler[R]) RecordResult(result R) { _ = "STUB: not implemented"; return }

func (t *adaptiveThrottler[R]) RecordSuccess() { _ = "STUB: not implemented"; return }

// Requires external locking.
func (t *adaptiveThrottler[R]) recordResult(result R, err error) { _ = "STUB: not implemented"; return }

func (t *adaptiveThrottler[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

// Computes a rejection rate as described in the SRE book: https://sre.google/sre-book/handling-overload/#client-side-throttling-a7sYUg
// The rejection rate ramps up rejections once the success rate falls below a threshold.
func computeRejectionRate(executions, successes, successRateThreshold, maxRejectionRate float64, executionThreshold uint) float64 {
	_ = "STUB: not implemented"
	return 0
}

// The max number of executions we should receive, given the successes and expected success rate threshold

// How many extra executions we processed beyond the max allowed
