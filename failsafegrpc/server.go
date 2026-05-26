package failsafegrpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/tap"

	"github.com/failsafe-go/failsafe-go"
)

// NewServerInHandle returns a tap.ServerInHandle that wraps the handler with the policies. This can be used to limit
// server side load with policies such as AdaptiveLimiter, AdaptiveThrottler, CircuitBreaker, Bulkhead, RateLimiter, and
// Cache, and should be prefered over NewUnaryServerInterceptor since it does not waste resources for requests that are
// rejected. Since a tap.ServerInHandle is meant to be non-blocking, be sure that any policies you're using do not have a
// wait time configured.
func NewServerInHandle[R any](policies ...failsafe.Policy[R]) tap.ServerInHandle {
	_ = "STUB: not implemented"
	return *new(tap.ServerInHandle)
}

// NewServerInHandleWithExecutor returns a tap.ServerInHandle that wraps the handler with a failsafe.Executor. This can
// be used to limit server side load with policies such as AdaptiveLimiter, AdaptiveThrottler, CircuitBreaker, Bulkhead,
// RateLimiter, and Cache, and should be prefered over NewUnaryServerInterceptorWithExecutor since it does not waste
// resources for requests that are rejected. Since a tap.ServerInHandle is meant to be non-blocking, be sure that any
// policies you're using do not have a wait time configured.
func NewServerInHandleWithExecutor[R any](executor failsafe.Executor[R]) tap.ServerInHandle {
	_ = "STUB: not implemented"
	return *new(tap.ServerInHandle)
}

// Pass a hint to executors that they should perform a check-only execution

// The execution is a noop since it's meant to be used with load limiting policies

// NewUnaryServerInterceptor returns a grpc.UnaryServerInterceptor that wraps the handler the policies. This can be used
// to limit server side load where the content of the request might influence whether it's rejected or not, such as with
// a CircuitBreaker. For load limiting that does not require inspecting requests, prefer NewServerInHandle.
// R is the response type.
func NewUnaryServerInterceptor[R any](policies ...failsafe.Policy[R]) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// NewUnaryServerInterceptorWithExecutor returns a grpc.UnaryServerInterceptor that wraps the handler with a failsafe.Executor. This can
// be used to limit server side load where the content of the request might influence whether it's rejected or not, such
// as with a CircuitBreaker. For load limiting that does not require inspecting requests, prefer NewServerInHandleWithExecutor.
// R is the response type.
func NewUnaryServerInterceptorWithExecutor[R any](executor failsafe.Executor[R]) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// NewUnaryServerInterceptorWithLevel extracts adaptivelimiter priority and level information from an incoming request
// and adds a level to the handling context. If a level is present in the incoming request metadata, it's added to the
// context. If a level is not present but a priority is, and ensureLevel is true, then the priority will be converted
// to a level, else if a priority is present it will be passed through the context.
func NewUnaryServerInterceptorWithLevel(ensureLevel bool) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}
