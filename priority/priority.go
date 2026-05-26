package priority

import (
	"context"
	"sync"
)

// Priority is an execution priority.
type Priority int

const (
	VeryLow Priority = iota
	Low
	Medium
	High
	VeryHigh
)

const totalLevels = 500

// RandomLevel returns a random level for the Priority.
func (p Priority) RandomLevel() int { _ = "STUB: not implemented"; return 0 }

// AddTo returns the ctx with the priority added to it as a value with the PriorityKey.
func (p Priority) AddTo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// MaxLevel returns the max level for the priority.
func (p Priority) MaxLevel() int { _ = "STUB: not implemented"; return 0 }

// MinLevel returns the min level for the priority.
func (p Priority) MinLevel() int { _ = "STUB: not implemented"; return 0 }

func (p Priority) levelRange() levelRange { _ = "STUB: not implemented"; return *new(levelRange) }

// levelRange provides a wider range of levels that allow for rejecting a subset of executions within a Priority.
type levelRange struct {
	lower, upper int
}

var priorityLevelRanges = map[Priority]levelRange{
	VeryLow:  {0, 99},
	Low:      {100, 199},
	Medium:   {200, 299},
	High:     {300, 399},
	VeryHigh: {400, 499},
}

type key int

// PriorityKey is a key to use with a Context that stores the priority value.
const PriorityKey key = 0

// LevelKey is a key to use with a Context that stores the level value.
const LevelKey key = 1

// ContextWithPriority returns a context with the priority value stored with the PriorityKey.
func ContextWithPriority(ctx context.Context, priority Priority) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ContextWithLevel returns a context with the level value stored with the LevelKey.
func ContextWithLevel(ctx context.Context, level int) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext returns the priority from the context, else -1.
func FromContext(ctx context.Context) Priority { _ = "STUB: not implemented"; return *new(Priority) }

// LevelFromContext returns a level for the level contained within the given context, else if a priority is contained
// within the context, a random level is generated within that priority, else -1 is returned.
func LevelFromContext(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

// LevelTracker tracks priority levels for executions, which can be used to prioritize rejections.
type LevelTracker interface {
	// RecordLevel records an execution having been accepted for the level.
	RecordLevel(level int)

	// GetLevel returns the level that falls at the quantile among all recorded levels in the tracker, else returns 0 if no
	// levels have been recorded.
	GetLevel(quantile float64) int
}

type windowedLevelTracker struct {
	mu          sync.Mutex
	window      []int // records recent levels
	levelCounts []int // current counts of each level
	head        int
	filled      bool
}

// NewLevelTracker creates a LevelTracker that stores the last windowSize recorded levels.
func NewLevelTracker(windowSize int) LevelTracker {
	_ = "STUB: not implemented"
	return *new(LevelTracker)
}

func (lt *windowedLevelTracker) RecordLevel(level int) { _ = "STUB: not implemented"; return }

// Remove old value from counts

// Add new value to counts

// Advance head

func (lt *windowedLevelTracker) GetLevel(quantile float64) int { _ = "STUB: not implemented"; return 0 }

// Determine how many recorded levels we need to find to match the quantile

// Count the levels until we hit the desired quantile
