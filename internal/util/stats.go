package util

import (
	"time"

	"github.com/bits-and-blooms/bitset"
)

// ExecutionStats for tracking execution results.
// Implementations are not concurrency safe and must be guarded externally.
type ExecutionStats interface {
	ExecutionCount() uint
	FailureCount() uint
	FailureRate() float64
	SuccessCount() uint
	SuccessRate() float64
	RecordFailure()
	RecordSuccess()
	Reset()
}

// countingStats is a ExecutionStats implementation that counts execution results using a BitSet.
type countingStats struct {
	bitSet       *bitset.BitSet
	head         uint
	occupiedBits uint
	successes    uint
	failures     uint
}

func NewCountingStats(size uint) ExecutionStats {
	_ = "STUB: not implemented"
	return *new(ExecutionStats)
}

/*
Sets the value of the next bit in the bitset, returning the previous value, else -1 if no previous value was set for the bit.

value is true if positive/success, false if negative/failure
*/
func (c *countingStats) setNext(value bool) int { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) ExecutionCount() uint { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) FailureCount() uint { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) FailureRate() float64 { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) SuccessCount() uint { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) SuccessRate() float64 { _ = "STUB: not implemented"; return 0 }

func (c *countingStats) RecordFailure() { _ = "STUB: not implemented"; return }

func (c *countingStats) RecordSuccess() { _ = "STUB: not implemented"; return }

func (c *countingStats) Reset() { _ = "STUB: not implemented"; return }

type timedStat struct {
	successes uint
	failures  uint
}

// timedStats is an ExecutionStats implementation that counts execution results within a time period, and Buckets
// results to minimize overhead.
type timedStats struct {
	BucketedWindow[timedStat]
}

func NewTimedStats(bucketCount int, thresholdingPeriod time.Duration, clock Clock) ExecutionStats {
	_ = "STUB: not implemented"
	return *new(ExecutionStats)
}

func (s *timedStats) ExecutionCount() uint { _ = "STUB: not implemented"; return 0 }

func (s *timedStats) FailureCount() uint { _ = "STUB: not implemented"; return 0 }

func (s *timedStats) FailureRate() float64 { _ = "STUB: not implemented"; return 0 }

func (s *timedStats) SuccessCount() uint { _ = "STUB: not implemented"; return 0 }

func (s *timedStats) SuccessRate() float64 { _ = "STUB: not implemented"; return 0 }

func (s *timedStats) RecordFailure() { _ = "STUB: not implemented"; return }

func (s *timedStats) RecordSuccess() { _ = "STUB: not implemented"; return }
