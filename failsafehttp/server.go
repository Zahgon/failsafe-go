package failsafehttp

import (
	"net/http"

	"github.com/failsafe-go/failsafe-go"
)

// NewHandler returns a new http.Handler that will perform failsafe request handling via the policies and
// innerHandler. The policies are composed around responses and will handle responses in reverse order.
func NewHandler(innerHandler http.Handler, policies ...failsafe.Policy[*http.Response]) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// NewHandlerWithExecutor returns a new http.Handler that will perform failsafe request handling via the executor and
// innerHandler. The policies are composed around responses and will handle responses in reverse order.
func NewHandlerWithExecutor(innerHandler http.Handler, executor failsafe.Executor[*http.Response]) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// NewHandlerWithLevel extracts adaptivelimiter priority and level information from an incoming request
// and adds a level to the handling context. If a level is present in the incoming request header, it's added to the
// context. If a level is not present but a priority is, and ensureLevel is true, then the priority will be converted
// to a level, else if a priority is present it will be passed through the context.
func NewHandlerWithLevel(innerHandler http.Handler, ensureLevel bool) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
