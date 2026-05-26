package failsafe

import (
	"context"

	"github.com/failsafe-go/failsafe-go/common"
)

// Executor handles failures according to configured policies. See [With] for details on creating an Executor.
//
// R is the result type. This type is concurrency safe.
type Executor[R any] interface {
	// Compose returns the Executor with the currently configured policies composed around the innerPolicy.
	Compose(innerPolicy Policy[R]) Executor[R]

	// ComposeAny returns the Executor with the currently configured policies composed around the innerPolicy. This method
	// allows composing policies that have any result type with policies that have a specific result type.
	ComposeAny(innerPolicy ResultAgnosticPolicy[any]) Executor[R]

	// Context returns the configured Context, else context.Background() by default.
	Context() context.Context

	// WithContext returns a new copy of the Executor with the ctx configured. Any executions created with the resulting
	// Executor will be canceled when the ctx is done. Executions can cooperate with cancellation by checking
	// Execution.Canceled or Execution.IsCanceled.
	WithContext(ctx context.Context) Executor[R]

	// OnDone registers the listener to be called when an execution is done.
	OnDone(listener func(ExecutionDoneEvent[R])) Executor[R]

	// OnSuccess registers the listener to be called when an execution is successful. If multiple policies, are configured,
	// this handler is called when execution is done and all policies succeed. If all policies do not succeed, then the
	// OnFailure registered listener is called instead.
	OnSuccess(listener func(ExecutionDoneEvent[R])) Executor[R]

	// OnFailure registers the listener to be called when an execution fails. This occurs when the execution fails according
	// to some policy, and all policies have been exceeded.
	OnFailure(listener func(ExecutionDoneEvent[R])) Executor[R]

	// Run executes the fn until successful or until the configured policies are exceeded.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	Run(fn func() error) error

	// RunWithExecution executes the fn until successful or until the configured policies are exceeded, while providing an
	// Execution to the fn.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	RunWithExecution(fn func(exec Execution[R]) error) error

	// Get executes the fn until a successful result is returned or the configured policies are exceeded.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	Get(fn func() (R, error)) (R, error)

	// GetWithExecution executes the fn until a successful result is returned or the configured policies are exceeded, while
	// providing an Execution to the fn.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	GetWithExecution(fn func(exec Execution[R]) (R, error)) (R, error)

	// RunAsync executes the fn in a goroutine until successful or until the configured policies are exceeded.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	RunAsync(fn func() error) ExecutionResult[R]

	// RunAsyncWithExecution executes the fn in a goroutine until successful or until the configured policies are exceeded,
	// while providing an Execution to the fn.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	RunAsyncWithExecution(fn func(exec Execution[R]) error) ExecutionResult[R]

	// GetAsync executes the fn in a goroutine until a successful result is returned or the configured policies are
	// exceeded.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	GetAsync(fn func() (R, error)) ExecutionResult[R]

	// GetAsyncWithExecution executes the fn in a goroutine until a successful result is returned or the configured policies
	// are exceeded, while providing an Execution to the fn.
	//
	// Any panic causes the execution to stop immediately without calling any event listeners.
	GetAsyncWithExecution(fn func(exec Execution[R]) (R, error)) ExecutionResult[R]
}

type executor[R any] struct {
	policies  []Policy[R]
	ctx       context.Context
	onDone    func(ExecutionDoneEvent[R])
	onSuccess func(ExecutionDoneEvent[R])
	onFailure func(ExecutionDoneEvent[R])
}

// With creates and returns a new Executor for result type R that will handle failures according to the given
// policies. The policies are composed around a func and will handle its results in reverse order. For example, consider:
//
//	failsafe.With(fallback, retryPolicy, circuitBreaker).Get(fn)
//
// This creates the following composition when executing a func and handling its result:
//
//	Fallback(RetryPolicy(CircuitBreaker(func)))
func With[R any](policies ...Policy[R]) Executor[R] { _ = "STUB: not implemented"; return nil }

// WithAny creates and returns a new Executor that can be used to compose other policies with the result type R, for the
// final composed execution. The executor will handle failures according to the given policies.
func WithAny[R any](policy ResultAgnosticPolicy[any]) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) Compose(innerPolicy Policy[R]) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) ComposeAny(innerPolicy ResultAgnosticPolicy[any]) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (e *executor[R]) WithContext(ctx context.Context) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) OnDone(listener func(ExecutionDoneEvent[R])) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) OnSuccess(listener func(ExecutionDoneEvent[R])) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) OnFailure(listener func(ExecutionDoneEvent[R])) Executor[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) Run(fn func() error) error { _ = "STUB: not implemented"; return nil }

func (e *executor[R]) RunWithExecution(fn func(exec Execution[R]) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) Get(fn func() (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func (e *executor[R]) GetWithExecution(fn func(exec Execution[R]) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func (e *executor[R]) RunAsync(fn func() error) ExecutionResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) RunAsyncWithExecution(fn func(exec Execution[R]) error) ExecutionResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) GetAsync(fn func() (R, error)) ExecutionResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) GetAsyncWithExecution(fn func(exec Execution[R]) (R, error)) ExecutionResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// This type mirrors part of policy.Executor, which we don't import here to avoid a cycle.
type policyExecutor[R any] interface {
	Apply(innerFn func(Execution[R]) *common.PolicyResult[R]) func(Execution[R]) *common.PolicyResult[R]
}

func (e *executor[R]) executeSync(fn func(exec Execution[R]) (R, error), withExec bool) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func (e *executor[R]) executeAsync(fn func(exec Execution[R]) (R, error), withExec bool) ExecutionResult[R] {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor[R]) execute(fn func(exec Execution[R]) (R, error), outerExec *execution[R], withExec bool) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Only copy and provide an execution to the user fn if needed

// Compose policy executors from the innermost policy to the outermost

// Execute

// policyAnyWrapper adapts Policy[R] to Policy[any], allowing Policy[any] to be used in compositions with Policy[R].
type policyAnyWrapper[R any] struct {
	inner Policy[any]
}

func (p *policyAnyWrapper[R]) ToExecutor(_ R) any { _ = "STUB: not implemented"; return *new(any) }

type policyExecutorAnyWrapper[R any] struct {
	inner policyExecutor[any]
}

func (pe *policyExecutorAnyWrapper[R]) Apply(innerFn func(Execution[R]) *common.PolicyResult[R]) func(Execution[R]) *common.PolicyResult[R] {
	_ = "STUB: not implemented"
	return nil
}

// Adapt func(Execution[any]) *PolicyResult[any] to func(Execution[R]) *PolicyResult[R]
