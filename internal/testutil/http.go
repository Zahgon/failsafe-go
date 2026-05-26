package testutil

import (
	"net/http"
	"net/http/httptest"
	"time"
)

func MockResponse(statusCode int, body string) *httptest.Server {
	_ = "STUB: not implemented"
	return nil
}

func MockHandler(statusCode int, body string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func MockFlakyServer(failTimes int, responseCode int, retryAfterDelay time.Duration, finalResponse string) (server *httptest.Server, resetFailures func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MockDelayedResponse(statusCode int, body string, delay time.Duration) *httptest.Server {
	_ = "STUB: not implemented"
	return nil
}

func MockDelayedHandler(statusCode int, body string, delay time.Duration) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func MockDelayedResponseWithEarlyFlush(statusCode int, body string, delay time.Duration) *httptest.Server {
	_ = "STUB: not implemented"
	return nil
}

// Ensure data and error is sent to the client
