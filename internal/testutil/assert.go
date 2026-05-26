package testutil

import (
	"testing"
	"time"

	"github.com/failsafe-go/failsafe-go"
)

func AssertDuration(t *testing.T, expectedDuration int, actualDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (w *Waiter) AssertEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func WaitAndAssertCanceled[R any](t *testing.T, waitDuration time.Duration, exec failsafe.Execution[R]) {
	_ = "STUB: not implemented"
	return
}
