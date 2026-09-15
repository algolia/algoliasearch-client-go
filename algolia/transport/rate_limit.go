package transport

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultMaxRateLimitRetries is how many same-host 429 waits a call gets
	// when Configuration.MaxRateLimitRetries is nil.
	DefaultMaxRateLimitRetries = 3
	defaultRateLimitWait       = time.Second
	// maxRateLimitWait is the longest wait a time.Duration can hold; larger
	// Retry-After values would otherwise overflow and fire immediately.
	maxRateLimitWait = time.Duration(math.MaxInt64)
)

var retryAfterDigits = regexp.MustCompile(`^\d+$`)

func parseRetryAfter(header http.Header) time.Duration {
	// http.Header.Get canonicalises the key and tolerates a nil map, so no manual lookup is needed
	raw := strings.TrimSpace(header.Get("Retry-After"))
	if !retryAfterDigits.MatchString(raw) {
		return defaultRateLimitWait
	}

	seconds, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		// Digits only reach this point, so the value does not fit in an int64.
		return maxRateLimitWait
	}

	if seconds <= 0 {
		return defaultRateLimitWait
	}

	if seconds > int64(maxRateLimitWait/time.Second) {
		return maxRateLimitWait
	}

	return time.Duration(seconds) * time.Second
}

func isRateLimited(code int) bool {
	return code == http.StatusTooManyRequests
}

func resolveMaxRateLimitRetries(configured *int) int {
	if configured == nil {
		return DefaultMaxRateLimitRetries
	}

	if *configured < 0 {
		return 0
	}

	return *configured
}

func defaultSleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("rate limit wait cancelled: %w", ctx.Err())
	}
}
