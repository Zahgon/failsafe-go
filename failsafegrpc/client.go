package failsafegrpc

import (
	"google.golang.org/grpc"

	"github.com/failsafe-go/failsafe-go"
)

const (
	priorityMetadataKey = "x-failsafe-priority"
	levelMetadataKey    = "x-failsafe-level"
)

// NewUnaryClientInterceptor returns a grpc.UnaryClientInterceptor that wraps the invoker with the policies.
//
// R is the response type.
func NewUnaryClientInterceptor[R any](policies ...failsafe.Policy[R]) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// NewUnaryClientInterceptorWithExecutor returns a grpc.UnaryClientInterceptor that wraps the invoker with a failsafe.Executor.
//
// R is the response type.
func NewUnaryClientInterceptorWithExecutor[R any](executor failsafe.Executor[R]) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// Merge the request context with the Executor so it's available for policies

// Merge the latest execution context for each attempt

// NewUnaryClientInterceptorWithLevel propagates adaptivelimiter priority and level information from a client context to
// a server via metadata. If a level is present it's propagated, else a priority is propagated if present.
func NewUnaryClientInterceptorWithLevel() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// Lazily construct or copy metadata
