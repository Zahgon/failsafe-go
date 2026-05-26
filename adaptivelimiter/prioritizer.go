package adaptivelimiter

import (
	"github.com/failsafe-go/failsafe-go/internal"
	"github.com/failsafe-go/failsafe-go/priority"
)

// NewPrioritizer returns a new Prioritizer.
func NewPrioritizer() priority.Prioritizer {
	_ = "STUB: not implemented"
	return *new(priority.Prioritizer)
}

// NewPrioritizerBuilder returns a new PrioritizerBuilder.
func NewPrioritizerBuilder() PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(PrioritizerBuilder)
}

// PrioritizerBuilder builds Prioritizer instances.
//
// This type is not concurrency safe.
type PrioritizerBuilder interface {
	priority.PrioritizerBuilder

	// WithUsageTracker configures a usage tracker to use with the prioritizer. The usage tracker tracks usage as total
	// execution duration, to enforce fairness when rejecting executions.
	WithUsageTracker(usageTracker priority.UsageTracker) PrioritizerBuilder
}

// BasePrioritizerConfig provides a base for implementing a PrioritizerBuilder.
type prioritizerConfig struct {
	*internal.BasePrioritizerConfig[*queueStats]
}

func (c *prioritizerConfig) WithUsageTracker(usageTracker priority.UsageTracker) PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(PrioritizerBuilder)
}

// Implements priority.RejectionStrategy.
type queueRejectionStrategy struct{}

func (s *queueRejectionStrategy) CombineStats(statsFuncs []func() *queueStats) *queueStats {
	_ = "STUB: not implemented"
	return nil
}
