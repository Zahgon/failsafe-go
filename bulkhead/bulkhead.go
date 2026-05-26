package bulkhead

import (
	"context"
	"errors"
	"time"

	"github.com/failsafe-go/failsafe-go"
)

// ErrFull is returned when an execution is attempted against a Bulkhead that is full.
var ErrFull = errors.New("bulkhead full")

// Bulkhead is a policy that restricts concurrent executions as a way of preventing system overload.
//
// R is the execution result type. This type is concurrency safe.
type Bulkhead[R any] interface {
	failsafe.ResultAgnosticPolicy[R]

	// AcquirePermit attempts to acquire a permit to perform an execution within the Bulkhead, waiting until one is
	// available or the execution is canceled. Returns context.Canceled if the ctx is canceled. Callers should call
	// ReleasePermit to release a successfully acquired permit back to the Bulkhead.
	//
	// ctx may be nil.
	AcquirePermit(ctx context.Context) error

	// AcquirePermitWithMaxWait attempts to acquire a permit to perform an execution within the Bulkhead, waiting up to the
	// maxWaitTime until one is available or the ctx is canceled. Returns ErrFull if a permit could not be acquired
	// in time. Returns context.Canceled if the ctx is canceled. Callers should call ReleasePermit to release a successfully
	// acquired permit back to the Bulkhead.
	//
	// ctx may be nil.
	AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) error

	// TryAcquirePermit tries to acquire a permit to perform an execution within the Bulkhead, returning immediately without
	// waiting. Returns true if the permit was acquired, else false. Callers should call ReleasePermit to release a
	// successfully acquired permit back to the Bulkhead.
	TryAcquirePermit() bool

	// ReleasePermit releases an execution permit back to the Bulkhead.
	ReleasePermit()
}

// Builder builds Bulkhead instances.
//
// R is the execution result type. This type is not concurrency safe.
type Builder[R any] interface {
	// WithMaxWaitTime configures the maxWaitTime to wait for permits to be available.
	WithMaxWaitTime(maxWaitTime time.Duration) Builder[R]

	// OnFull registers the listener to be called when the bulkhead is full.
	OnFull(listener func(event failsafe.ExecutionEvent[R])) Builder[R]

	// Build returns a new Bulkhead using the builder's configuration.
	Build() Bulkhead[R]
}

type config[R any] struct {
	maxConcurrency uint
	maxWaitTime    time.Duration
	onFull         func(failsafe.ExecutionEvent[R])
}

var _ Builder[any] = &config[any]{}

// New returns a new Bulkhead for execution result type R and the maxConcurrency.
func New[R any](maxConcurrency uint) Bulkhead[R] { _ = "STUB: not implemented"; return nil }

// NewBuilder returns a Builder for execution result type R which builds Timeouts for the timeoutDelay.
func NewBuilder[R any](maxConcurrency uint) Builder[R] { _ = "STUB: not implemented"; return nil }

func (c *config[R]) WithMaxWaitTime(maxWaitTime time.Duration) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) OnFull(listener func(event failsafe.ExecutionEvent[R])) Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func (c *config[R]) Build() Bulkhead[R] { _ = "STUB: not implemented"; return nil }

// TODO copy base fields

type bulkhead[R any] struct {
	config[R]
	semaphore chan struct{}
}

func (*bulkhead[R]) ResultAgnostic() { _ = "STUB: not implemented"; return }

func (b *bulkhead[R]) AcquirePermit(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bulkhead[R]) AcquirePermitWithMaxWait(ctx context.Context, maxWaitTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Initial attempt, in case permit is immediately available or context is done, so we don't race with a timer

// Second attempt with timer

func (b *bulkhead[R]) TryAcquirePermit() bool { _ = "STUB: not implemented"; return false }

func (b *bulkhead[R]) ReleasePermit() { _ = "STUB: not implemented"; return }

func (b *bulkhead[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }
