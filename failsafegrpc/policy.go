package failsafegrpc

import (
	"github.com/failsafe-go/failsafe-go/retrypolicy"
)

// NewRetryPolicyBuilder returns a retrypolicy.Builder that will retry on gRPC status codes that are considered
// retryable (UNAVAILABLE, DEADLINE_EXCEEDED, RESOURCE_EXHAUSTED), up to 2 times by default, with no delay between
// attempts. Additional handling can be added by chaining the builder with more conditions.
//
// R is the execution result type.
func NewRetryPolicyBuilder[R any]() retrypolicy.Builder[R] { _ = "STUB: not implemented"; return nil }
