package ratelimiter

import (
	"context"
	"errors"
	"time"

	"github.com/failsafe-go/failsafe-go"
)

// ErrExceeded is returned when an execution exceeds a configured rate limit.
var ErrExceeded = errors.New("rate limit exceeded")

/*
RateLimiter is a Policy that can control the rate of executions as a way of preventing system overload.

There are two types of rate limiting: smooth and bursty. NewSmooth rate limiting will evenly spread out execution requests
over-time, effectively smoothing out uneven execution request rates. NewBursty rate limiting allows potential bursts of
executions to occur, up to a configured max per time period.

Rate limiting is based on permits, which can be requested in order to perform rate limited execution. Permits are
automatically refreshed over time based on the rate limiter's configuration.

This type provides methods that block while waiting for permits to become available, and also methods that return
immediately. The blocking methods include:

  - AcquirePermit
  - AcquirePermits
  - AcquirePermitWithMaxWait
  - AcquirePermitsWithMaxWait

The methods that return immediately include:

  - TryAcquirePermit
  - TryAcquirePermits
  - ReservePermit
  - ReservePermits
  - TryReservePermit
  - TryReservePermits

This type provides methods that return ErrExceeded when permits cannot be acquired, and also methods that
return a bool. The Acquire methods all return ErrExceeded when permits cannot be acquired, and the TryAcquire
methods return a boolean.

The ReservePermit methods attempt to reserve permits and return an expected wait time before the permit can be used.
This helps integrate with scenarios where you need to wait externally.

R is the execution result type. This type is concurrency safe.
*/
type RateLimiter[R any] interface {
	failsafe.ResultAgnosticPolicy[R]

	// AcquirePermit attempts to acquire a permit to perform an execution against the rate limiter, waiting until one is
	// available or the ctx is canceled. Returns an error if the ctx is canceled.
	//
	// ctx may be nil.
	AcquirePermit(ctx context.Context) error

	// AcquirePermits attempts to acquire the requested permits to perform executions against the rate limiter, waiting until
	// they are available or the ctx is canceled. Returns an error if the ctx is canceled.
	//
	// ctx may be nil.
	AcquirePermits(ctx context.Context, permits uint) error

	// AcquirePermitWithMaxWait attempts to acquire a permit to perform an execution against the rate limiter, waiting up to
	// the maxWaitTime until one is available or the ctx is canceled. Returns ErrExceeded if a permit would not be
	// available in time. Returns an error if the context is canceled.
	//
	// ctx may be nil.
	AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) error

	// AcquirePermitsWithMaxWait attempts to acquire the requested permits to perform executions against the rate limiter,
	// waiting up to the maxWaitTime until they are available or the ctx is canceled. Returns ErrExceeded if the
	// permits would not be available in time. Returns an error if the context is canceled.
	//
	// ctx may be nil.
	AcquirePermitsWithMaxWait(ctx context.Context, requestedPermits uint, maxWaitTime time.Duration) error

	// ReservePermit reserves a permit to perform an execution against the rate limiter, and returns the time that the caller
	// is expected to wait before acting on the permit. Returns 0 if the permit is immediately available and no waiting is
	// needed.
	ReservePermit() time.Duration

	// ReservePermits reserves the permits to perform executions against the rate limiter, and returns the time that the
	// caller is expected to wait before acting on the permits. Returns 0 if the permits are immediately available and no
	// waiting is needed.
	ReservePermits(permits uint) time.Duration

	// TryAcquirePermit tries to acquire a permit to perform an execution against the rate limiter, returning immediately
	// without waiting.
	TryAcquirePermit() bool

	// TryAcquirePermits tries to acquire the requested permits to perform executions against the rate limiter, returning
	// immediately without waiting. Returns true if the permit was successfully acquired, else false.
	TryAcquirePermits(permits uint) bool

	// TryReservePermit tries to reserve a permit to perform an execution against the rate limiter, and returns the time that
	// the caller is expected to wait before acting on the permit, as long as it's less than the maxWaitTime.
	//
	//  - Returns the expected wait time for the permit if it was successfully reserved.
	//  - Returns 0 if the permit was successfully reserved and no waiting is needed.
	//  - Returns -1 if the permit was not reserved because the wait time would be greater than the maxWaitTime.
	TryReservePermit(maxWaitTime time.Duration) time.Duration

	// TryReservePermits tries to reserve the permits to perform executions against the rate limiter, and returns the time
	// that the caller is expected to wait before acting on the permits, as long as it's less than the maxWaitTime.
	//
	//  - Returns the expected wait time for the permit if it was successfully reserved.
	//  - Returns 0 if the permit was successfully reserved and no waiting is needed.
	//  - Returns -1 if the permit was not reserved because the wait time would be greater than the maxWaitTime.
	TryReservePermits(requestedPermits uint, maxWaitTime time.Duration) time.Duration
}

/*
Builder builds RateLimiter instances.

R is the execution result type. This type is not concurrency safe.
*/
type Builder[R any] interface {
	// WithMaxWaitTime configures the maxWaitTime to wait for permits to be available. If permits cannot be acquired before
	// the maxWaitTime is exceeded, then the rate limiter will return ErrExceeded.
	//
	// This setting only applies when the resulting RateLimiter is used with the failsafe.Executor APIs. It does not
	// apply when the RateLimiter is used in a standalone way.
	WithMaxWaitTime(maxWaitTime time.Duration) Builder[R]

	// OnRateLimitExceeded registers the listener to be called when the rate limit is exceeded.
	OnRateLimitExceeded(listener func(failsafe.ExecutionEvent[R])) Builder[R]

	// Build returns a new RateLimiter using the builder's configuration.
	Build() RateLimiter[R]
}

type config[R any] struct {
	// Common
	maxWaitTime         time.Duration
	onRateLimitExceeded func(failsafe.ExecutionEvent[R])

	// Smooth
	interval time.Duration

	// Bursty
	periodPermits int
	period        time.Duration
}

/*
NewSmooth returns a smooth RateLimiter for execution result type R and the maxExecutions and period, which control how
frequently an execution is permitted. The individual execution rate is computed as period / maxExecutions. For example,
with maxExecutions of 100 and a period of 1000 millis, individual executions will be permitted at a max rate of one
every 10 millis. The returned RateLimiter will have a max wait time of 0.

Executions are performed with no delay until they exceed the max rate, after which they are rejected.
*/
func NewSmooth[R any](maxExecutions uint, period time.Duration) RateLimiter[R] {
	_ = "STUB: not implemented"
	return nil
}

/*
NewSmoothWithMaxRate returns a smooth RateLimiter for execution result type R and the maxRate, which controls how
frequently an execution is permitted. For example, a maxRate of 10*time.Millisecond would allow up to one execution
every 10 milliseconds. The returned RateLimiter will have a max wait time of 0.

Executions are performed with no delay until they exceed the max rate, after which they are rejected.
*/
func NewSmoothWithMaxRate[R any](maxRate time.Duration) RateLimiter[R] {
	_ = "STUB: not implemented"
	return nil
}

/*
NewSmoothBuilder returns a smooth Builder for execution result type R and the maxExecutions and period, which
control how frequently an execution is permitted. The individual execution rate is computed as period / maxExecutions.
For example, with maxExecutions of 100 and a period of 1000 millis, individual executions will be permitted at a max
rate of one every 10 millis.

By default, the returned Builder will have a max wait time of 0.

Executions are performed with no delay until they exceed the max rate, after which they are either rejected or
will block and wait until the max wait time is exceeded.
*/
func NewSmoothBuilder[R any](maxExecutions uint, period time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

/*
NewSmoothBuilderWithMaxRate returns a smooth Builder for execution result type R and the maxRate, which controls
how frequently an execution is permitted. For example, a maxRate of 10*time.Millisecond would allow up to one execution
every 10 milliseconds.

By default, the returned Builder will have a max wait time of 0.

Executions are performed with no delay until they exceed the maxRate, after which they are either rejected or will block
and wait until the max wait time is exceeded.
*/
func NewSmoothBuilderWithMaxRate[R any](maxRate time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

/*
NewBursty returns a bursty RateLimiter for execution result type R and the maxExecutions per period. For example, a
maxExecutions value of 100 with a period of 1 second would allow up to 100 executions every 1 second. The returned
RateLimiter will have a max wait time of 0.

Executions are performed with no delay until they exceed the max rate, after which they are rejected.
*/
func NewBursty[R any](maxExecutions uint, period time.Duration) RateLimiter[R] {
	_ = "STUB: not implemented"
	return nil
}

/*
NewBurstyBuilder returns a bursty Builder for execution result type R and the maxExecutions per period. For
example, a maxExecutions value of 100 with a period of 1 second would allow up to 100 executions every 1 second.

By default, the returned Builder will have a max wait time of 0.

Executions are performed with no delay up until the maxExecutions are reached for the current period, after which
executions are either rejected or will block and wait until the max wait time is exceeded.
*/
func NewBurstyBuilder[R any](maxExecutions uint, period time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) WithMaxWaitTime(maxWaitTime time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnRateLimitExceeded(listener func(event failsafe.ExecutionEvent[R])) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) Build() RateLimiter[R] { _ = "STUB: not implemented"; return nil }

// TODO copy base fields

type rateLimiter[R any] struct {
	config[R]
	stats stats
}

func (*rateLimiter[R]) ResultAgnostic() { _ = "STUB: not implemented"; return }

func (r *rateLimiter[R]) AcquirePermit(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rateLimiter[R]) AcquirePermits(ctx context.Context, permits uint) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rateLimiter[R]) AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rateLimiter[R]) AcquirePermitsWithMaxWait(ctx context.Context, requestedPermits uint, maxWaitTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rateLimiter[R]) ReservePermit() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rateLimiter[R]) ReservePermits(permits uint) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rateLimiter[R]) TryAcquirePermit() bool { _ = "STUB: not implemented"; return false }

func (r *rateLimiter[R]) TryAcquirePermits(permits uint) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *rateLimiter[R]) TryReservePermit(maxWaitTime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rateLimiter[R]) TryReservePermits(requestedPermits uint, maxWaitTime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rateLimiter[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

func (r *rateLimiter[R]) Reset() { _ = "STUB: not implemented"; return }
