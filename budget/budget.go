package budget

import (
	"errors"
	"sync/atomic"

	"github.com/failsafe-go/failsafe-go"
)

// ErrExceeded is returned when an execution attempt exceeds the budget.
var ErrExceeded = errors.New("budget exceeded")

// Budget restricts concurrent executions as a way of preventing system overload.
//
// This type is concurrency safe.
type Budget interface {
	// RetryRate returns the current rate of retries relative to total executions, from 0 to 1.
	RetryRate() float64

	// HedgeRate returns the current rate of hedges relative to total executions, from 0 to 1.
	HedgeRate() float64
}

// Builder builds Budget instances.
//
// This type is not concurrency safe.
type Builder interface {
	// WithMaxRate configures the max rate of inflight executions that can be retries and/or hedges.
	WithMaxRate(maxRate float64) Builder

	// WithMinConcurrency configures the min number of budgeted retries and/or hedges that can be executed, regardless of
	// the total number of inflight executions.
	WithMinConcurrency(minConcurrency uint) Builder

	// OnBudgetExceeded registers the listener to be called when the budget is exceeded.
	OnBudgetExceeded(listener func(ExceededEvent)) Builder

	// Build returns a new Budget using the builder's configuration.
	Build() Budget
}

// ExecutionType indicates the type of execution used by the budget.
type ExecutionType string

const (
	// RetryExecution indicates a retry execution was used with the budget.
	RetryExecution ExecutionType = "retry"
	// HedgeExecution indicates a hedge execution was used with the budget.
	HedgeExecution ExecutionType = "hedge"
)

// ExceededEvent indicates a budget limit has been exceeded.
type ExceededEvent struct {
	// ExecutionType indicates the type of execution that exceeded the budget.
	ExecutionType ExecutionType

	// ExecutionInfo provides information about the execution that caused the budget to be exceeded.
	failsafe.ExecutionInfo

	// Budget provides access to the current budget state.
	Budget
}

type config struct {
	maxRate          float64
	minConcurrency   uint
	onBudgetExceeded func(ExceededEvent)
}

var _ Builder = &config{}

// New returns a new budget with a default maxRate of .2 and minConcurrency of 3.
func New() Budget { _ = "STUB: not implemented"; return *new(Budget) }

// NewBuilder returns a TypeBuilder for execution result type R which builds Budgets with a default maxRate of .2 and
// minConcurrency of 3.
func NewBuilder() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *config) WithMaxRate(maxRate float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (c *config) WithMinConcurrency(minConcurrency uint) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (c *config) OnBudgetExceeded(listener func(ExceededEvent)) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (c *config) Build() Budget { _ = "STUB: not implemented"; return *new(Budget) }

// TODO copy base fields

type budget struct {
	config

	executions atomic.Int32
	retries    atomic.Int32
	hedges     atomic.Int32
}

func (b *budget) TryAcquireRetryPermit() bool { _ = "STUB: not implemented"; return false }

func (b *budget) TryAcquireHedgePermit() bool { _ = "STUB: not implemented"; return false }

func (b *budget) ReleaseRetryPermit() { _ = "STUB: not implemented"; return }

func (b *budget) ReleaseHedgePermit() { _ = "STUB: not implemented"; return }

func (b *budget) RetryRate() float64 { _ = "STUB: not implemented"; return 0 }

func (b *budget) HedgeRate() float64 { _ = "STUB: not implemented"; return 0 }

func (b *budget) OnBudgetExceeded(executionType ExecutionType, info failsafe.ExecutionInfo) {
	_ = "STUB: not implemented"
	return
}
