// Package policytesting is needed to avoid a circular dependency with the policy package.
package policytesting

import (
	"sync/atomic"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/adaptivelimiter"
	"github.com/failsafe-go/failsafe-go/budget"
	"github.com/failsafe-go/failsafe-go/bulkhead"
	"github.com/failsafe-go/failsafe-go/cachepolicy"
	"github.com/failsafe-go/failsafe-go/circuitbreaker"
	"github.com/failsafe-go/failsafe-go/fallback"
	"github.com/failsafe-go/failsafe-go/hedgepolicy"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
	"github.com/failsafe-go/failsafe-go/timeout"
)

func Reset[R any](p failsafe.Policy[R]) { _ = "STUB: not implemented"; return }

func WithRetryStats[R any](rp retrypolicy.Builder[R], stats *Stats) retrypolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithRetryLogs[R any](rp retrypolicy.Builder[R]) retrypolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithRetryStatsAndLogs[R any](rp retrypolicy.Builder[R], stats *Stats) retrypolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func withRetryStatsAndLogs[R any](rp retrypolicy.Builder[R], stats *Stats, withLogging bool) retrypolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithAdaptiveLimiterStatsAndLogs[R any](l adaptivelimiter.Builder[R], stats *Stats, withLogging bool) adaptivelimiter.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithBreakerStats[R any](cb circuitbreaker.Builder[R], stats *Stats) circuitbreaker.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithBreakerLogs[R any](cb circuitbreaker.Builder[R]) circuitbreaker.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func withBreakerStatsAndLogs[R any](cb circuitbreaker.Builder[R], stats *Stats, withLogging bool) circuitbreaker.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithTimeoutStatsAndLogs[R any](to timeout.Builder[R], stats *Stats) timeout.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithFallbackStatsAndLogs[R any](fb fallback.Builder[R], stats *Stats) fallback.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithHedgeStatsAndLogs[R any](hp hedgepolicy.Builder[R], stats *Stats) hedgepolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithBulkheadStatsAndLogs[R any](bh bulkhead.Builder[R], stats *Stats, withLogging bool) bulkhead.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func WithBudgetStatsAndLogs(b budget.Builder, stats *Stats, withLogging bool) budget.Builder {
	_ = "STUB: not implemented"
	return *new(budget.Builder)
}

func WithCacheStats[R any](cp cachepolicy.Builder[R], stats *Stats) cachepolicy.Builder[R] {
	_ = "STUB: not implemented"
	return nil
}

func withStatsAndLogs[P any, R any](policy failsafe.FailurePolicyBuilder[P, R], stats *Stats, withLogging bool) {
	_ = "STUB: not implemented"
	return
}

type Stats struct {
	executions atomic.Int32
	successes  atomic.Int32
	failures   atomic.Int32

	// Retry specific stats
	retries         atomic.Int32
	retriesExceeded atomic.Int32
	aborts          atomic.Int32

	limitsExceeded atomic.Int32
	limitsChanged  atomic.Int32

	// Hedge specific stats
	hedges atomic.Int32

	// Bulkhead specific stats
	fulls atomic.Int32

	// Budget stats
	budgetExceeded atomic.Int32

	// Cache specific stats
	caches      atomic.Int32
	cacheHits   atomic.Int32
	cacheMisses atomic.Int32
	cachedCount atomic.Int32
}

func (s *Stats) Executions() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Successes() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Failures() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Retries() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) RetriesExceeded() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) LimitsChanged() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) LimitsExceeded() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Hedges() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Aborts() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Fulls() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) BudgetExceededs() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) CacheHits() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) CacheMisses() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Caches() int { _ = "STUB: not implemented"; return 0 }

func (s *Stats) Reset() { _ = "STUB: not implemented"; return }

// Retry specific stats

// AdaptiveLimiter specific stats

// Hedge specific stats

// Bulkhead specific stats

// Budget specific stats

// Cache specific stats
