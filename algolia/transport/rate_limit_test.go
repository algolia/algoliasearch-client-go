//nolint:testpackage // Tests unexported parseRetryAfter, resolveMaxRateLimitRetries, and isRateLimited.
package transport

import (
	"net/http"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

func TestParseRetryAfter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		header http.Header
		want   time.Duration
	}{
		{name: "positive seconds", header: http.Header{"Retry-After": []string{"2"}}, want: 2 * time.Second},
		{name: "missing header", header: nil, want: time.Second},
		{name: "empty map", header: http.Header{}, want: time.Second},
		{name: "empty value", header: http.Header{"Retry-After": []string{""}}, want: time.Second},
		{name: "zero", header: http.Header{"Retry-After": []string{"0"}}, want: time.Second},
		{name: "junk", header: http.Header{"Retry-After": []string{"120abc"}}, want: time.Second},
		{name: "http-date", header: http.Header{"Retry-After": []string{"Wed, 21 Oct 2015 07:28:00 GMT"}}, want: time.Second},
		{name: "trimmed", header: http.Header{"Retry-After": []string{" 4 "}}, want: 4 * time.Second},
		{name: "leading zeros", header: http.Header{"Retry-After": []string{"007"}}, want: 7 * time.Second},
		{name: "too large for a Duration", header: http.Header{"Retry-After": []string{"10000000000"}}, want: maxRateLimitWait},
		{name: "does not fit in int64", header: http.Header{"Retry-After": []string{"99999999999999999999"}}, want: maxRateLimitWait},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := parseRetryAfter(tt.header); got != tt.want {
				t.Fatalf("parseRetryAfter() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestResolveMaxRateLimitRetries(t *testing.T) {
	t.Parallel()

	if got := resolveMaxRateLimitRetries(nil); got != DefaultMaxRateLimitRetries {
		t.Fatalf("nil = %d, want %d", got, DefaultMaxRateLimitRetries)
	}

	if got := resolveMaxRateLimitRetries(utils.ToPtr(0)); got != 0 {
		t.Fatalf("0 = %d, want 0", got)
	}

	if got := resolveMaxRateLimitRetries(utils.ToPtr(5)); got != 5 {
		t.Fatalf("5 = %d, want 5", got)
	}

	if got := resolveMaxRateLimitRetries(utils.ToPtr(-1)); got != 0 {
		t.Fatalf("negative = %d, want 0", got)
	}
}

func TestIsRateLimited(t *testing.T) {
	t.Parallel()

	if !isRateLimited(429) {
		t.Fatal("429 should be rate limited")
	}

	if isRateLimited(400) || isRateLimited(500) {
		t.Fatal("only 429 is rate limited")
	}
}
