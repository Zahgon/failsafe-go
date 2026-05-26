package internal

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/failsafe-go/failsafe-go/priority"
)

// BasePrioritizerConfig provides a base for implementing a PrioritizerBuilder.
type BasePrioritizerConfig[S Stats] struct {
	logger             *slog.Logger
	LevelTracker       priority.LevelTracker
	UsageTracker       priority.UsageTracker
	Strategy           RejectionStrategy[S]
	onThresholdChanged func(event priority.ThresholdChangedEvent)
}

var _ priority.PrioritizerBuilder = &BasePrioritizerConfig[Stats]{}

func (c *BasePrioritizerConfig[S]) WithLevelTracker(levelTracker priority.LevelTracker) priority.PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(priority.PrioritizerBuilder)
}

func (c *BasePrioritizerConfig[S]) WithLogger(logger *slog.Logger) priority.PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(priority.PrioritizerBuilder)
}

func (c *BasePrioritizerConfig[S]) OnThresholdChanged(listener func(event priority.ThresholdChangedEvent)) priority.PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(priority.PrioritizerBuilder)
}

func (c *BasePrioritizerConfig[S]) Build() priority.Prioritizer {
	_ = "STUB: not implemented"
	return *new(priority.Prioritizer)
}

// TODO copy base fields

type Stats interface {
	// ComputeRejectionRate returns the rate at which future executions should be rejected, given the window.
	ComputeRejectionRate() float64

	// DebugLogArgs returns any args you'd like to include in the prioritizer's debug logs.
	DebugLogArgs() []any
}

type RejectionStrategy[S any] interface {
	CombineStats(statsFuncs []func() S) (totalStats S)
}

// BasePrioritizer provides a base implementation of a Prioritizer.
type BasePrioritizer[S Stats] struct {
	BasePrioritizerConfig[S]

	// Mutable state
	mu              sync.Mutex
	statsFuncs      []func() S // Guarded by mu
	numStats        atomic.Int32
	rejectionRate   float64 // Guarded by mu
	RejectionThresh atomic.Int32
}

func (p *BasePrioritizer[S]) Register(statsFunc func() S) { _ = "STUB: not implemented"; return }

func (p *BasePrioritizer[S]) RejectionRate() float64 { _ = "STUB: not implemented"; return 0 }

func (p *BasePrioritizer[S]) RejectionThreshold() int { _ = "STUB: not implemented"; return 0 }

func (p *BasePrioritizer[S]) RegisteredPolicies() int { _ = "STUB: not implemented"; return 0 }

// Calibrate computes a combined rejection rate and threshold based on all registered statsFuncs.
func (p *BasePrioritizer[S]) Calibrate() { _ = "STUB: not implemented"; return }

// LevelFromContext gets a level based on usage from the user in the context, else returns LevelFromContext, else 0.
func (p *BasePrioritizer[S]) LevelFromContext(ctx context.Context) int {
	_ = "STUB: not implemented"
	return 0
}

// LevelFromContextWithPriority gets a level based on usage from the user in the context and the priority, else returns
// LevelFromContext, else a random level from the priority.
func (p *BasePrioritizer[S]) LevelFromContextWithPriority(ctx context.Context, prt priority.Priority) int {
	_ = "STUB: not implemented"
	return 0
}

func (p *BasePrioritizer[S]) ScheduleCalibrations(ctx context.Context, interval time.Duration) context.CancelFunc {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc)
}
