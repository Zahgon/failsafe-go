package adaptivethrottler

import (
	"github.com/failsafe-go/failsafe-go/priority"
)

// NewPrioritizer returns a new Prioritizer.
func NewPrioritizer() priority.Prioritizer {
	_ = "STUB: not implemented"
	return *new(priority.Prioritizer)
}

// NewPrioritizerBuilder returns a new PrioritizerBuilder.
func NewPrioritizerBuilder() priority.PrioritizerBuilder {
	_ = "STUB: not implemented"
	return *new(priority.PrioritizerBuilder)
}

// Implements priority.RejectionStrategy.
type throttlerRejectionStrategy struct{}

// CombineStats combines throttler stats using a weighted rejection rate, where weights are based on the number of
// executions, to determine a weighted average rejection rate.
func (s *throttlerRejectionStrategy) CombineStats(statsFuncs []func() *throttlerStats) *throttlerStats {
	_ = "STUB: not implemented"
	return nil
}
