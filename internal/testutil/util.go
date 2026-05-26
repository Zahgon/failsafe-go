package testutil

import (
	"context"
	"sync/atomic"
	"time"
)

var CanceledContextFn = func() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func ContextFn(ctx context.Context) func() context.Context { _ = "STUB: not implemented"; return nil }

// ContextWithCancel returns a function that provides a context that is canceled after the sleepTime.
func ContextWithCancel(sleepTime time.Duration) func() context.Context {
	_ = "STUB: not implemented"
	return nil
}

type TestClock struct {
	Time time.Time
}

func NewTestClock(millis int) *TestClock { _ = "STUB: not implemented"; return nil }

func (tc *TestClock) SetTime(millis int) { _ = "STUB: not implemented"; return }

func (tc *TestClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type TestStopwatch struct {
	CurrentTime int64
}

func (t *TestStopwatch) ElapsedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (t *TestStopwatch) Reset() { _ = "STUB: not implemented"; return }

func Timed(fn func()) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

type Waiter struct {
	count atomic.Int32
	done  chan struct{}
}

func NewWaiter() *Waiter { _ = "STUB: not implemented"; return nil }

func (w *Waiter) Await(expectedResumes int) { _ = "STUB: not implemented"; return }

func (w *Waiter) AwaitWithTimeout(expectedResumes int, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (w *Waiter) Resume() { _ = "STUB: not implemented"; return }

func MillisToNanos(millis int) int64 { _ = "STUB: not implemented"; return 0 }

func GetPrioritizerRejectionThreshold(prioritizer any) *atomic.Int32 {
	_ = "STUB: not implemented"
	return nil
}

func GetBudgetExecutions(budget any) *atomic.Int32 { _ = "STUB: not implemented"; return nil }
