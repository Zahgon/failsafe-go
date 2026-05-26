package failsafehttp

import (
	"io"
	"net/http"

	"github.com/failsafe-go/failsafe-go"
)

const (
	priorityHeaderKey = "X-Failsafe-Priority"
	levelHeaderKey    = "X-Failsafe-Level"
)

type roundTripper struct {
	next     http.RoundTripper
	executor failsafe.Executor[*http.Response]
}

// NewRoundTripper returns a new http.RoundTripper that will perform failsafe round trips via the policies and
// innerRoundTripper. If innerRoundTripper is nil, http.DefaultTransport will be used. The policies are composed around
// requests and will handle responses in reverse order.
func NewRoundTripper(innerRoundTripper http.RoundTripper, policies ...failsafe.Policy[*http.Response]) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// NewRoundTripperWithExecutor returns a new http.RoundTripper that will perform failsafe round trips via the executor and
// innerRoundTripper. If innerRoundTripper is nil, http.DefaultTransport will be used.
func NewRoundTripperWithExecutor(innerRoundTripper http.RoundTripper, executor failsafe.Executor[*http.Response]) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

type levelRoundTripper struct {
	next http.RoundTripper
}

// NewRoundTripperWithLevel propagates adaptivelimiter priority and level information from a client context to
// a server via HTTP headers. If a level is present it's propagated, else a priority is propagated if present.
func NewRoundTripperWithLevel(innerRoundTripper http.RoundTripper) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (p *levelRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *roundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Request struct {
	executor failsafe.Executor[*http.Response]
	request  *http.Request
	client   *http.Client
}

// NewRequest creates and returns a new Request that will perform failsafe round trips via the request, client, and
// policies. The policies are composed around requests and will handle responses in reverse order.
func NewRequest(request *http.Request, client *http.Client, policies ...failsafe.Policy[*http.Response]) *Request {
	_ = "STUB: not implemented"
	return nil
}

// NewRequestWithExecutor creates and returns a new Request that will perform failsafe round trips via the request,
// client, and executor.
func NewRequestWithExecutor(request *http.Request, client *http.Client, executor failsafe.Executor[*http.Response]) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) Do() (*http.Response, error) { _ = "STUB: not implemented"; return nil, nil }

func doRequest(request *http.Request, executor failsafe.Executor[*http.Response], reqFn func(r *http.Request) (*http.Response, error)) (resp *http.Response, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bodyWithCancel will handle context cancellation

// Merge the request context with the Executor so it's available for policies

// Calls cancelOuter if a bodyWithCancel is not returned

// Merge the latest execution context into the request for each attempt

// Calls cancelInner if a bodyWithCancel is not returned

// Get new body for each attempt

// Wrap the response body to cancel both contexts when the body is closed

// bodyWithCancel wraps a response body and calls the cancel functions when the body is closed.
type bodyWithCancel struct {
	io.ReadCloser
	cancelOuter func(error)
	cancelInner func(error)
	cancelFn    func()
}

func (b *bodyWithCancel) Close() error { _ = "STUB: not implemented"; return nil }

// bodyReader returns a function that can repeatedly read the untypedBody of an http.Request.
func bodyReader(untypedBody any) (func() (io.Reader, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Match bytes.Reader first to avoid seeking via ReadSeeker match
