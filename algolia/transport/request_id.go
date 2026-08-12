package transport

import (
	"crypto/rand"
	"net/url"
	"strings"
	"sync/atomic"
	"time"
)

// RequestIDHeader is the name of the header carrying the Request-ID minted
// when Configuration.RequestIDEnabled resolves to true.
const RequestIDHeader = "Request-ID"

// RequestIDQueryParam is the name of the query parameter through which a
// caller can supply its own Request-ID. The server consults it only when the
// Request-ID header is absent, so minting a header would shadow it.
const RequestIDQueryParam = "x-algolia-request-id"

const requestIDAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

const requestIDLength = 11

// requestIDFallbackCounter keeps the fallback seeds of concurrent mints
// distinct when they land within the same clock tick.
var requestIDFallbackCounter atomic.Int64

// NewRequestID returns a fresh 11-character base62 identifier suitable for
// the Request-ID header. The modulo bias of the byte mapping is acceptable:
// the ID is a tracing breadcrumb, not a secret.
func NewRequestID() string {
	b := make([]byte, requestIDLength)

	_, err := rand.Read(b)
	if err != nil {
		// A request must not fail over its tracing metadata: when the entropy
		// source is unavailable, degrade to a time-seeded sequence.
		n := time.Now().UnixNano() + requestIDFallbackCounter.Add(1)
		for i := range b {
			n = n*6364136223846793005 + 1442695040888963407
			b[i] = byte((n >> 33) & 0xff)
		}
	}

	for i := range b {
		b[i] = requestIDAlphabet[int(b[i])%len(requestIDAlphabet)]
	}

	return string(b)
}

// HasRequestID reports whether any of the given header maps carries a
// Request-ID entry, whatever its casing. Unlike http.Header, plain header
// maps keep the caller's literal casing, so the lookup must not assume a
// canonical form.
func HasRequestID(headers ...map[string]string) bool {
	for _, m := range headers {
		for k := range m {
			if strings.EqualFold(k, RequestIDHeader) {
				return true
			}
		}
	}

	return false
}

// HasRequestIDQueryParam reports whether the given query parameters carry a
// Request-ID entry, whatever its casing. url.Values keeps the caller's
// literal key casing, so the lookup must not assume a canonical form.
func HasRequestIDQueryParam(queryParams url.Values) bool {
	for k := range queryParams {
		if strings.EqualFold(k, RequestIDQueryParam) {
			return true
		}
	}

	return false
}
