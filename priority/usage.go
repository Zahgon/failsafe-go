package priority

import (
	"container/list"
	"context"
	"sync"
	"time"

	"github.com/failsafe-go/failsafe-go/internal/util"
)

const (
	// UserKey is a key to use with a Context that stores a user ID.
	UserKey key = 2

	// The max level value for a priority class
	maxLevel = 99.0
)

// ContextWithUser returns a context with the userID stored with the UserKey.
func ContextWithUser(ctx context.Context, userID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// UserFromContext returns the userID from the context, else "".
func UserFromContext(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// UsageTracker tracks resource usage per user as execution usages for fair execution prioritization.
type UsageTracker interface {
	// RecordUsage calculates and records usage for the user.
	RecordUsage(userID string, usage int64)

	// GetUsage returns the total recorded usage for a user, returning the usage and true if the user exists in the tracker,
	// else 0 and false.
	GetUsage(userID string) (int64, bool)

	// GetLevel returns the priority level for a user based on their recent usage.
	GetLevel(userID string, priority Priority) int

	// Calibrate calibrates levels based on the distribution of usage across all users.
	Calibrate()
}

type usageTracker struct {
	clock              util.Clock
	newWindowFn        func() *usageWindow
	maxUsers           int
	expirationDuration time.Duration

	mu sync.RWMutex
	// Guarded by mu
	users map[string]*userEntry
	lru   *list.List
}

type userEntry struct {
	window     *usageWindow
	quantile   float64 // Negative value represents being uncalibrated
	lastActive time.Time
	lruElement *list.Element
}

// NewUsageTracker creates a new UsageTracker with the specified configuration. The UsageTracker will track up to the
// maxUsers, and track any recent usage within the usageWindow. If a user hasn't had activity in 2x the usageWindow,
// they're removed from the tracker.
func NewUsageTracker(windowDuration time.Duration, maxUsers int) UsageTracker {
	_ = "STUB: not implemented"
	return *new(UsageTracker)
}

func (ut *usageTracker) RecordUsage(userID string, usage int64) { _ = "STUB: not implemented"; return }

func (ut *usageTracker) GetLevel(userID string, priority Priority) int {
	_ = "STUB: not implemented"
	return 0
}

// Handle users that have no recorded usages

// Handle new entries with no recently recorded usages

// Handle uncalibrated user

func (ut *usageTracker) GetUsage(userID string) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Calibrate has an O(n log n) time complexity, where n is the number of users.
func (ut *usageTracker) Calibrate() { _ = "STUB: not implemented"; return }

// Update percentiles for all active users

// computeQuantile returns the quantile for a usage, among the sortedUsages.
func (ut *usageTracker) computeQuantile(usage int64, sortedUsages []int64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (ut *usageTracker) evictOldest() { _ = "STUB: not implemented"; return }

type usageStat struct {
	totalUsage int64
	samples    uint32
}

type usageWindow struct {
	util.BucketedWindow[usageStat]
}

func newUsageWindow(bucketCount int, thresholdingPeriod time.Duration, clock util.Clock) *usageWindow {
	_ = "STUB: not implemented"
	return nil
}

func (w *usageWindow) RecordUsage(usage int64) { _ = "STUB: not implemented"; return }

func (w *usageWindow) TotalUsage() int64 { _ = "STUB: not implemented"; return 0 }

func (w *usageWindow) Samples() uint32 { _ = "STUB: not implemented"; return 0 }
