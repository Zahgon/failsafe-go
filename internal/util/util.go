package util

import (
	"context"
	"reflect"
	"time"
)

type number interface {
	~int | ~int64 | ~uint | ~uint64
}

func noop(_ error) { _ = "STUB: not implemented"; return }

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// ErrorTypesMatch indicates whether the err or any unwrapped causes of the err are assignable to the target type. This is
// similar to the test that errors.As performs, but does not actually assign a value and allows a non-pointer target.
// This method also allows a non-pointer target for an error that's implemented with pointer receivers.
// Panics if target is nil or not an error.
func ErrorTypesMatch(err error, target any) bool { _ = "STUB: not implemented"; return false }

// If targetType is not an error, convert it to a pointer and check again

func errorAs(err error, targetType reflect.Type) bool { _ = "STUB: not implemented"; return false }

// MergeContexts returns a context that is canceled when either ctx1 or ctx2 are Done.
func MergeContexts(ctx1, ctx2 context.Context) (context.Context, context.CancelCauseFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelCauseFunc)
}

// mergedContext wraps two parent contexts and checks both for values.
type mergedContext struct {
	context.Context // For Done(), Err()
	ctx1, ctx2      context.Context
}

// Value checks ctx1 first, then ctx2.
func (m *mergedContext) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

// Deadline returns the earliest deadline from both parent contexts.
func (m *mergedContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// AppliesToAny returns true if any of the biPredicates evaluate to true for the values.
func AppliesToAny[A any, B any](biPredicates []func(A, B) bool, value1 A, value2 B) bool {
	_ = "STUB: not implemented"
	return false
}

// RoundDown returns the input rounded down to the nearest interval.
func RoundDown[T number](input T, interval T) T { _ = "STUB: not implemented"; return *new(T) }

func RandomDelayInRange[T number](delayMin T, delayMax T, random float64) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func RandomDelay[T number](delay T, jitter T, random float64) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func RandomDelayFactor[T number](delay T, jitterFactor float64, random float64) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Smooth returns a value that is decreased by some portion of the oldValue, and increased by some portion of the
// newValue, based on the factor.
func Smooth(oldValue, newValue, factor float64) float64 { _ = "STUB: not implemented"; return 0 }

var log10Values []int

func init() {
	for i := 0; i < 100; i++ {
		log10Values = append(log10Values, 1)
	}
	for i := 100; i < 1000; i++ {
		log10Values = append(log10Values, 2)
	}
}

func Log10Func(factor int) func(limit int) int { _ = "STUB: not implemented"; return nil }

type Clock interface {
	Now() time.Time
}

var WallClock = &wallClock{}

type wallClock struct {
}

func (wc *wallClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type Stopwatch interface {
	ElapsedTime() time.Duration

	Reset()
}

type wallClockStopwatch struct {
	startTime time.Time
}

func NewStopwatch() Stopwatch { _ = "STUB: not implemented"; return *new(Stopwatch) }

func (s *wallClockStopwatch) ElapsedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *wallClockStopwatch) Reset() { _ = "STUB: not implemented"; return }

func Round(v float64) float64 { _ = "STUB: not implemented"; return 0 }
