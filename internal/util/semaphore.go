package util

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"time"
)

// ErrWaitExceeded is returned when maxWaitTime is exceeded while waiting on a permit.
var ErrWaitExceeded = errors.New("wait time exceeded")

// DynamicSemaphore is a semaphore that can be dynamically resized and allows FIFO blocking when full.
//
// This type is concurrency safe.
type DynamicSemaphore struct {
	mu sync.Mutex

	// Guarded by mu
	size    int
	used    int
	waiters list.List
}

func NewDynamicSemaphore(size int) *DynamicSemaphore { _ = "STUB: not implemented"; return nil }

// Acquire acquires a permit from the sempahore, blocking until one is made available via Release or the ctx is Done.
// Blocking callers are unblocked in FIFO order as permits are released.
func (s *DynamicSemaphore) Acquire(ctx context.Context) error {
	_ = "STUB: not implemented"

	// See if a permit is immediately available
	return nil
}

// Create a waiter for a permit

// AcquireWithMaxWait acquires a permit from the sempahore, blocking until one is made available via Release, the ctx is Done, or
// the maxWaitTime is hit. Blocking callers are unblocked in FIFO order as permits are released.
func (s *DynamicSemaphore) AcquireWithMaxWait(ctx context.Context, maxWaitTime time.Duration) error {
	_ = "STUB: not implemented"

	// See if a permit is immediately available
	return nil
}

// If semaphore is full and no wait is possible, return immediately

// Create a waiter for a permit

// drainWaiter drains a waiter when an Acquire attempt times out. Returns nil if the waiter is ready, else the err.
func (s *DynamicSemaphore) drainWaiter(waiter chan struct{}, waiterElem *list.Element, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore err if a waiter is ready

// Remove waiter if done prematurely

// TryAcquire acquires a permit if one is available, returning true, else returns false if no permits are available.
func (s *DynamicSemaphore) TryAcquire() bool { _ = "STUB: not implemented"; return false }

// Release releases a permit which unblocks a waiter if one exists.
// Panics if called without a corresponding call to Acquire or TryAcqyure.
func (s *DynamicSemaphore) Release() { _ = "STUB: not implemented"; return }

// Check for invalid state

// If we have waiters and capacity, wake up the next waiter

// SetSize resizes the semaphore. If the size is increased, waiters will be unblocked as needed.
func (s *DynamicSemaphore) SetSize(size int) { _ = "STUB: not implemented"; return }

// If capacity increased, wake up waiters that can now acquire

// Wakes a blocked waiter and acquires a permit.
// Requires s.mu to be locked before calling.
func (s *DynamicSemaphore) wakeAndAcquire(waiterElem *list.Element) {
	_ = "STUB: not implemented"
	return
}

func (s *DynamicSemaphore) IsFull() bool { _ = "STUB: not implemented"; return false }

// Waiters returns how many callers are blocked waiting for permits.
func (s *DynamicSemaphore) Waiters() int { _ = "STUB: not implemented"; return 0 }

// Used returns how many permits are currently in use.
func (s *DynamicSemaphore) Used() int { _ = "STUB: not implemented"; return 0 }
