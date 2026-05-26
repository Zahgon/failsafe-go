package adaptivelimiter

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/influxdata/tdigest"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/internal/util"
	"github.com/failsafe-go/failsafe-go/priority"
)

var (
	// ErrExceeded is returned when an execution exceeds the current limit.
	ErrExceeded = errors.New("limit exceeded")

	// Limit thresholding and adjustment functions
	alphaFunc    = util.Log10Func(3)
	betaFunc     = util.Log10Func(6)
	increaseFunc = util.Log10Func(1)
	decreaseFunc = util.Log10Func(1)
)

const (
	warmupSamples   = 10
	smoothedSamples = 5
)

// AdaptiveLimiter is an adaptive concurrency limiter that adjusts its limit up or down based on execution time trends:
//   - When recent execution times are trending up relative to baseline execution times, the concurrency limit is decreased.
//   - When recent execution times are trending down relative to baseline execution times, the concurrency limit is increased.
//
// To accomplish this, recent execution times are tracked and regularly compared to a weighted moving average of
// baseline execution times. Limit increases are additionally controlled to ensure they don't increase execution times.
// Any executions in excess of the limit will be rejected with ErrExceeded.
//
// By default, during overload, an AdaptiveLimiter will converge on a concurrency limit that represents the capacity of
// the machine it's running on, and avoids having executions queue. Since enforcing a limit without allowing for
// queueing is too strict in some cases and may cause unexpected rejections, optional queueing of executions when the
// limiter is full can be enabled via WithQueueing.
//
// Prioritized rejections can be enabled via BuildPrioritized, which accepts a Prioritizer that regularly determines a
// rejection threshold based on recent queueing across limiters.
//
// R is the execution result type. This type is concurrency safe.
type AdaptiveLimiter[R any] interface {
	failsafe.ResultAgnosticPolicy[R]
	Metrics

	// AcquirePermit attempts to acquire a permit to perform an execution via the limiter, waiting until one is available or
	// the execution is canceled. Returns [context.Canceled] if the ctx is canceled. Callers must call Record or Drop to
	// release a successfully acquired permit back to the limiter. ctx may be nil.
	AcquirePermit(context.Context) (Permit, error)

	// AcquirePermitWithMaxWait attempts to acquire a permit to perform an execution via the limiter, waiting until one is
	// available, the execution is canceled, or the maxWaitTime is exceeded.
	AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) (Permit, error)

	// TryAcquirePermit attempts to acquire a permit to perform an execution via the limiter, returning whether the Permit
	// was acquired or not. This method will never block, and only considers whether the limiter is full - it does not allow
	// queueing. Callers must call Record or Drop to release a successfully acquired permit back to the limiter.
	TryAcquirePermit() (Permit, bool)

	// CanAcquirePermit returns whether it's currently possible to acquire a permit.
	CanAcquirePermit() bool

	// Reset resets the limiter to its initial limit.
	Reset()
}

// Metrics provides info about the adaptive limiter.
//
// R is the execution result type. This type is concurrency safe.
type Metrics interface {
	// Limit returns the concurrent execution limit, as calculated by the adaptive limiter.
	Limit() int

	// Inflight returns the current number of inflight executions. The limit is adjusted using the max inflight requests for
	// a sampling period, which may be higher than the amount at a given point in time.
	Inflight() int

	// MaxInflight returns the max number of inflight executions during the most recent sampling period. This is used to
	// adjust the limit at the end of the sampling period, and should reflect how the current limit was set.
	MaxInflight() int

	// Queued returns the current number of queued executions when the limiter is full.
	Queued() int
}

// Permit is a permit to perform an execution that must be completed by calling Record or Drop.
type Permit interface {
	// Record records an execution completion and releases a permit back to the limiter. The execution duration will be used
	// to influence the limiter.
	Record()

	// Drop releases an execution permit back to the limiter without recording a completion. This should be used when an
	// execution completes prematurely, such as via a timeout, and we don't want the execution duration to influence the
	// limiter.
	Drop()
}

/*
Builder defines base behavior for building AdaptiveLimiter instances.

This type is not concurrency safe.
*/
type Builder[R any] interface {
	// WithLimits configures min, max, and initial concurrency limits.
	//
	// The default values are 1, 200, and 20.
	// Panics if minLimit is not <= maxLimit or initialLimit is not between minLimit and maxLimit.
	WithLimits(minLimit uint, maxLimit uint, initialLimit uint) Builder[R]

	// WithMaxLimitFactor configures a maxLimitFactor which caps the limit as some multiple of the current inflight
	// executions. When the limiter is healthy, the limit will increase up to this multiple of the current inflight
	// executions. This is intended to provide headroom for bursts of executions that may happen during normal operation,
	// avoiding unnecessary rejections before the limiter has time to adjust. For example, with a maxLimitFactor of 5.0:
	//  - 1 inflight causes a max limit of 5
	//  - 10 inflight causes a max limit of 50
	//  - 100 inflight causes a max limit of 500
	//
	// The default value is 5, which means the limit will only rise to 5 times the inflight executions.
	// Panics if maxLimitFactor < 1.
	WithMaxLimitFactor(maxLimitFactor float64) Builder[R]

	// WithMaxLimitFactorDecay allows different effective maxLimitFactor values based on the current inflight requests. By
	// default, decay is disabled, and the max limit factor scales linearly. With a decay, the maxLimitFactor is reduced by
	// the decay for each 10x increase in inflight executions. For example, with maxLimitFactor=5 and decay=1.0:
	//  - 1 inflight causes a 5x factor (max limit of 5)
	//  - 10 inflight causes a 4x factor (max limit of 40)
	//  - 100 inflight causes a 3x factor (max limit of 300)
	//  - 1000 inflight causes a 2x factor (max limit of 2000)
	//
	// Higher maxLimitFactorDecay values make the limit scale more conservatively at higher loads.
	// When a decay is configured, the effective max limit factor will never drop below the minLimitFactor.
	// Panics if maxLimitFactorDecay < 0 or minLimitFactor < 1.
	WithMaxLimitFactorDecay(maxLimitFactorDecay, minLimitFactor float64) Builder[R]

	// WithMaxLimitFunc allows a max limit to be specified given the current inflight executions. When configured, a
	// maxLimitFunc replaces any configured max limit factor and decay.
	WithMaxLimitFunc(maxLimitFunc func(inflight int) float64) Builder[R]

	// WithMaxLimitStabilizationWindow configures a stabilization window that remembers the peak inflight executions
	// over the given duration. This prevents temporary dips in inflight from pulling down the max limit calculation.
	//
	// Without a stabilization window, oscillating inflight patterns can cause the limit to get stuck near the peak
	// inflight value, unable to grow to the configured maxLimitFactor headroom above it.
	//
	// Disabled by default.
	// Panics if window is negative.
	WithMaxLimitStabilizationWindow(window time.Duration) Builder[R]

	// WithRecentWindow configures how recent execution times are collected and summarized. These help the limiter determine
	// when execution times are trending up or down, relative to the baseline, which helps detect overload. The minDuration
	// and maxDuration define the time bounds for sample collection, while minSamples ensures enough data is collected
	// before adjusting the limit.
	//
	// The default values are 1s, 30s, and 50.
	// Panics if minDuration is not <= maxDuration.
	WithRecentWindow(minDuration time.Duration, maxDuration time.Duration, minSamples uint) Builder[R]

	// WithRecentQuantile configures the recentQuantile of recent execution times to consider when adjusting the concurrency limit.
	//
	// Defaults to 0.9 which uses p90 samples.
	// Panics if recentQuantile is not > 0 and < 1.
	WithRecentQuantile(quantile float64) Builder[R]

	// WithBaselineWindow configures how the baseline execution times are maintained and updated. The baseline represents
	// the long-term average execution times that recent execution times are compared against to detect overload. When the
	// recent window is filled an aggregated recent sample is added to the baseline window. The baselineAge controls how
	// many samples the baseline window remembers - smaller values make the baseline adapt faster to recent changes, while
	// larger values keep it more stable by retaining influence from older measurements. The default value is 10.
	WithBaselineWindow(baselineAge uint) Builder[R]

	// WithCorrelationWindow configures how many recent limit and execution time measurements are stored to detect whether
	// increases in limits correlate with increases in execution times, which causes the limit to be adjusted down.
	//
	// The default value is 50.
	WithCorrelationWindow(size uint) Builder[R]

	// WithQueueing enables additional queueing of executions when the limiter is full. Queueing allows short execution
	// spikes to be absorbed without strictly rejecting executions.
	//
	// When queueing is enabled and the limiter is full, executions are queued up to the current limit times the
	// initialRejectionFactor, after which they will gradually start to be rejected up to the limit times the
	// maxRejectionFactor. This allows rejection to gradually adjust based on how many executions are queued, relative to the
	// limit. Queueing allows short execution spikes to be absorbed without strictly rejecting executions.
	//
	// WithQueueing is disabled by default, which means no executions will queue when the limiter is full.
	// Panics if initialRejectionFactor or maxRejectionFactor are < 1 or if initialRejectionFactor is not <= maxRejectionFactor.
	WithQueueing(initialRejectionFactor, maxRejectionFactor float64) Builder[R]

	// WithMaxWaitTime configures the maxWaitTime to wait for a permit to be available, when the limiter is full.
	WithMaxWaitTime(maxWaitTime time.Duration) Builder[R]

	// OnLimitExceeded registers the listener to be called when the limit is exceeded.
	OnLimitExceeded(listener func(event failsafe.ExecutionEvent[R])) Builder[R]

	// OnLimitChanged configures a listener to be called with the limit changes.
	OnLimitChanged(listener func(event LimitChangedEvent)) Builder[R]

	// WithLogger configures a logger which provides debug logging of limit adjustments.
	WithLogger(logger *slog.Logger) Builder[R]

	// Build returns a new AdaptiveLimiter using the builder's configuration.
	Build() AdaptiveLimiter[R]

	// BuildPrioritized returns a new PrioritizedLimiter using the builder's configuration. This enables queueing and
	// prioritized rejections of executions when the limiter is full, where executions block while waiting for a permit.
	// Enabling this allows short execution spikes to be absorbed without strictly rejecting executions when the limiter is
	// full. Rejections are performed using the Prioritizer, which sets a rejection threshold based on the limits and queues
	// of all the limiters being used by the Prioritizer. The amount of queueing can be configured via WithQueueing, and
	// defaults to an initialRejectionFactor of 2 times the current limit and a maxRejectionFactor of 3 times the current
	// limit.
	//
	// The Prioritizer can and should be shared across all limiter instances that need to coordinate prioritization.
	// Prioritized rejection is disabled by default, which means no executions will block when the limiter is full.
	BuildPrioritized(prioritizer priority.Prioritizer) PriorityLimiter[R]
}

// LimitChangedEvent indicates an AdaptiveLimiter's limit has changed.
type LimitChangedEvent struct {
	OldLimit uint
	NewLimit uint
}

type config[R any] struct {
	maxWaitTime time.Duration
	logger      *slog.Logger

	// Limit config
	minLimit, maxLimit          float64
	initialLimit                uint
	maxLimitFunc                func(inflight int) float64
	maxLimitFactor              float64
	maxLimitFactorDecay         float64
	minLimitFactor              float64
	maxLimitStabilizationWindow time.Duration

	// Windowing config
	recentWindowMinDuration time.Duration
	recentWindowMaxDuration time.Duration
	recentWindowMinSamples  uint
	recentQuantile          float64
	baselineWindowAge       uint
	correlationWindowSize   uint

	// Rejection config
	initialRejectionFactor float64
	maxRejectionFactor     float64

	// Listeners
	onLimitExceeded func(failsafe.ExecutionEvent[R])
	onLimitChanged  func(LimitChangedEvent)
}

var _ Builder[any] = &config[any]{}

// NewWithDefaults creates an AdaptiveLimiter with min, max, and initial limits of 1, 200, and 20 respectively, and a maxLimitFactor of 5.
// The recent window's min and max durations are 1 and 30 seconds respectively, and the min samples is 50.
// The baseline window age is 10 and the correlation window size is 50.
// To configure additional options on an AdaptiveLimiter, use NewBuilder() instead.
func NewWithDefaults[R any]() AdaptiveLimiter[R] { _ = "STUB: not implemented"; return nil }

// NewBuilder creates a Builder for execution result type R.
// The min, max, and initial limits default to 1, 200, and 20 respectively, and the maxLimitFactor to 5.
// The recent window's min and max durations default to 1 and 30 seconds respectively, and the min samples to 50.
// The baseline window age defaults to 10 and the correlation window size to 50.
func NewBuilder[R any]() Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) WithLimits(minLimit uint, maxLimit uint, initialLimit uint) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxLimitFactor(maxLimitFactor float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxLimitFactorDecay(maxLimitFactorDecay, minLimitFactor float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxLimitFunc(maxLimitFunc func(inflight int) float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxLimitStabilizationWindow(window time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithRecentWindow(minDuration time.Duration, maxDuration time.Duration, minSamples uint) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithRecentQuantile(quantile float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithBaselineWindow(baselineAge uint) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithCorrelationWindow(size uint) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithQueueing(initialRejectionFactor, maxRejectionFactor float64) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxWaitTime(maxWaitTime time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithLogger(logger *slog.Logger) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnLimitExceeded(listener func(event failsafe.ExecutionEvent[R])) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnLimitChanged(listener func(event LimitChangedEvent)) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) Build() AdaptiveLimiter[R] { _ = "STUB: not implemented"; return nil }

// Wait indefinitely for queued executions

func (c *config[R]) BuildPrioritized(p priority.Prioritizer) PriorityLimiter[R] {
	_ = "STUB: not implemented"
	return nil
}

type limitChange int

const (
	increase limitChange = iota
	decrease
	hold
)

type tdigestSample struct {
	MinRTT      time.Duration
	MaxInflight int
	Size        uint
	*tdigest.TDigest
}

func (td *tdigestSample) Add(rtt time.Duration, inflight int) { _ = "STUB: not implemented"; return }

func (td *tdigestSample) Reset() { _ = "STUB: not implemented"; return }

type adaptiveLimiter[R any] struct {
	config[R]

	// Mutable state
	semaphore *util.DynamicSemaphore
	mu        sync.RWMutex

	// Guarded by mu
	limit                 float64        // The current concurrency limit
	maxInflightWindow     util.MaxWindow // Tracks the max inflight over a stabilization window
	recentRTT             tdigestSample  // Recent execution times
	lastMaxInflight       int            // The max inflight requests for the last sampling period
	medianFilter          util.MovingMedian
	smoothedRecentRTT     util.MovingAverage
	baselineRTT           util.MovingAverage     // Tracks baseline execution time
	nextUpdateTime        time.Time              // Tracks when the limit can next be updated
	throughputCorrelation util.CorrelationWindow // Tracks the correlation between concurrency and throughput
	rttCorrelation        util.CorrelationWindow // Tracks the correlation between concurrency and round trip times (RTT)
}

func (*adaptiveLimiter[R]) ResultAgnostic() { _ = "STUB: not implemented"; return }

func (l *adaptiveLimiter[R]) AcquirePermit(ctx context.Context) (Permit, error) {
	_ = "STUB: not implemented"
	return *new(Permit), nil
}

func (l *adaptiveLimiter[R]) AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) (Permit, error) {
	_ = "STUB: not implemented"
	return *new(Permit), nil
}

func (l *adaptiveLimiter[R]) TryAcquirePermit() (Permit, bool) {
	_ = "STUB: not implemented"
	return *new(Permit), false
}

func (l *adaptiveLimiter[R]) newPermit() Permit { _ = "STUB: not implemented"; return *new(Permit) }

func (l *adaptiveLimiter[R]) CanAcquirePermit() bool { _ = "STUB: not implemented"; return false }

func (l *adaptiveLimiter[R]) Limit() int { _ = "STUB: not implemented"; return 0 }

func (l *adaptiveLimiter[R]) Inflight() int { _ = "STUB: not implemented"; return 0 }

func (l *adaptiveLimiter[R]) MaxInflight() int { _ = "STUB: not implemented"; return 0 }

func (l *adaptiveLimiter[R]) Queued() int { _ = "STUB: not implemented"; return 0 }

func (l *adaptiveLimiter[R]) Reset() { _ = "STUB: not implemented"; return }

// Records the duration of a completed execution, updating the concurrency limit if the recentRTT window is full.
func (l *adaptiveLimiter[R]) record(now time.Time, rtt time.Duration, inflight int, dropped bool) {
	_ = "STUB: not implemented"
	return
}

// updateLimit updates the concurrency limit based on the gradient between the recentRTT and historical baselineRTT.
// A stability check prevents unnecessary decreases during steady state.
// A correlation adjustment prevents upward drift during overload.
func (l *adaptiveLimiter[R]) updateLimit(recentRTT float64, inflight int, now time.Time) {
	_ = "STUB: not implemented"
	// Update baseline RTT and calculate the queue size
	// This is the primary signal that we threshold off of to detect overload
	return
}

// Calculate throughput correlation, throughput CV, and RTT correlation
// These are the secondary signals that we threshold off of to detect overload
// Convert to RPS

// Additional values for thresholding the limit

// alpha is the queueSize threshold below which we increase
// beta is the queueSize threshold above which we decrease

// Get the max inflight over the stabilization window

// Decrease gradually to avoid noise if inflights fluctuate

// Clamp the limit based on absolute min and max

func computeChange(queueSize, alpha, beta int, overloaded bool, throughputCorr, throughputCV, rttCorr float64) (change limitChange, reason string) {
	_ = "STUB: not implemented"
	return *

	// This condition handles severe overload where recent RTT significantly exceeds the baseline
	new(limitChange), ""
}

// This condition prevents runaway limit increases during moderate overload where inflight is increasing but throughput is decreasing

// This condition prevents runaway limit increases during moderate overload where throughputCorr is weak and rttCorr is high
// This indicates overload since latency is increasing with inflight, but throughput is not

// This condition prevents runaway limit increases during moderate overload where throughputCV low and rttCorr is high
// This indicates overload since latency is increasing with inflight, but throughput is not

// If our queue size is sufficiently small, increase until we detect overload

// If queueSize is between alpha and beta, leave the limit unchanged

func (l *adaptiveLimiter[R]) logLimit(direction, reason string, limit float64, gradient float64, queueSize, inflight int, recentRTT, baselineRTT, rttCorr, throughput, throughputCorr, throughputCV float64) {
	_ = "STUB: not implemented"
	return
}

// computeMaxLimit computes the max limit using a provided function, else based on max limit factor, with optional
// logarithmic decay.
func (l *adaptiveLimiter[R]) computeMaxLimit(inflight int) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Apply logarithmic decay, where the factor decreases by the decay amount for each order of magnitude increase in inflights

func (l *adaptiveLimiter[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

func (l *adaptiveLimiter[R]) canAcquirePermit(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *adaptiveLimiter[R]) configRef() *config[R] { _ = "STUB: not implemented"; return nil }

type recordingPermit[R any] struct {
	clock           util.Clock
	limiter         *adaptiveLimiter[R]
	startTime       time.Time
	currentInflight int
	userID          string
	usageTracker    priority.UsageTracker
}

func (p *recordingPermit[R]) Record() { _ = "STUB: not implemented"; return }

func (p *recordingPermit[R]) RecordUsage(userID string, usage int64) {
	_ = "STUB: not implemented"
	return
}

func (p *recordingPermit[R]) Drop() { _ = "STUB: not implemented"; return }
