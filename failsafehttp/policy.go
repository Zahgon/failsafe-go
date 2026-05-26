package failsafehttp

import (
	"net/http"
	"regexp"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
)

var (
	unsupportedScheme     = regexp.MustCompile(`unsupported protocol scheme`)
	certNotTrusted        = regexp.MustCompile(`certificate is not trusted`)
	stoppedAfterRedirects = regexp.MustCompile(`stopped after \d+ redirects\z`)
)

// NewRetryPolicyBuilder returns a retrypolicy.Builder that will retry non-terminal HTTP errors and responses up
// to 2 times, by default. If a Retry-After header is present in the response, it will be used as a delay between
// retries. Additional handling and delay configuration can be added to the resulting builder.
func NewRetryPolicyBuilder() retrypolicy.Builder[*http.Response] {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors

// Do not retry unsupported protocol scheme error
// This will be a url.Error when using an http.Client, and an errorString when using a RoundTripper

// Do not retry when certain error messages are observed

// Do not retry on unknown authority errors

// Retry on all other url errors

// Handle response

// Retry on 429

// Retry on most 5xx responses

// DelayFunc delays according to an http.Response Retry-After header. This can be used as a delay in a RetryPolicy or a CircuitBreaker.
func DelayFunc(exec failsafe.ExecutionAttempt[*http.Response]) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
