//nolint:testpackage // Tests unexported Transport.sleep and newTestTransport helpers.
package transport

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

func newTestTransport(t *testing.T, serverURL string, extraHosts int, maxRetries *int) *Transport {
	t.Helper()

	serverHost, err := url.Parse(serverURL)
	if err != nil {
		t.Fatal(err)
	}

	hosts := []StatefulHost{
		NewStatefulHost(serverHost.Scheme, serverHost.Host, func(call.Kind) bool { return true }),
	}
	for i := 0; i < extraHosts; i++ {
		hosts = append(hosts, NewStatefulHost(serverHost.Scheme, "never-called.example", func(call.Kind) bool { return true }))
	}

	tr := New(Configuration{
		Hosts:               hosts,
		MaxRateLimitRetries: maxRetries,
	})
	tr.sleep = func(context.Context, time.Duration) error { return nil }

	return tr
}

func newGetRequest(t *testing.T, path string) *http.Request {
	t.Helper()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	return req
}

func TestRequestWaitsRetryAfterOnSameHost(t *testing.T) {
	t.Parallel()

	var (
		calls atomic.Int32
		waits []time.Duration
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)

		if calls.Load() == 1 {
			w.Header().Set("Retry-After", "2")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"Too many requests"}`))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer srv.Close()

	tr := newTestTransport(t, srv.URL, 1, nil)
	tr.sleep = func(_ context.Context, d time.Duration) error {
		waits = append(waits, d)

		return nil
	}

	res, body, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", res.StatusCode)
	}

	if string(body) != `{"message":"ok"}` {
		t.Fatalf("body = %s", body)
	}

	if calls.Load() != 2 {
		t.Fatalf("calls = %d, want 2", calls.Load())
	}

	if len(waits) != 1 || waits[0] != 2*time.Second {
		t.Fatalf("waits = %v, want [2s]", waits)
	}
}

func TestRequestMissingRetryAfterWaitsOneSecond(t *testing.T) {
	t.Parallel()

	var (
		calls atomic.Int32
		waits []time.Duration
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)

		if calls.Load() == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"Too many requests"}`))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer srv.Close()

	tr := newTestTransport(t, srv.URL, 0, nil)
	tr.sleep = func(_ context.Context, d time.Duration) error {
		waits = append(waits, d)

		return nil
	}

	res, _, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if len(waits) != 1 || waits[0] != time.Second {
		t.Fatalf("waits = %v, want [1s]", waits)
	}
}

func TestRequestMaxRateLimitRetriesZeroFailsImmediately(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"Too many requests"}`))
	}))
	defer srv.Close()

	tr := newTestTransport(t, srv.URL, 0, utils.ToPtr(0))

	res, _, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("status = %d", res.StatusCode)
	}

	if calls.Load() != 1 {
		t.Fatalf("calls = %d, want 1", calls.Load())
	}
}

func TestRequestExhaustsMaxRateLimitRetries(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"Too many requests"}`))
	}))
	defer srv.Close()

	tr := newTestTransport(t, srv.URL, 0, nil)

	res, _, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("status = %d", res.StatusCode)
	}

	if calls.Load() != 4 {
		t.Fatalf("calls = %d, want 4", calls.Load())
	}
}

func TestRequestRateLimitBudgetIsSharedAcrossHosts(t *testing.T) {
	t.Parallel()

	// first host: 429 (budget 3 -> 2), then 500 fails over; second host: 429, 429 (-> 0), then a
	// 429 that must surface because the budget is per request, not per host
	var firstCalls, secondCalls, waits atomic.Int32

	first := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if firstCalls.Add(1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		_, _ = w.Write([]byte(`{"message":"error"}`))
	}))
	defer first.Close()

	second := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if secondCalls.Add(1) <= 3 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"Too many requests"}`))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer second.Close()

	firstHost, _ := url.Parse(first.URL)
	secondHost, _ := url.Parse(second.URL)
	tr := New(Configuration{
		Hosts: []StatefulHost{
			NewStatefulHost(firstHost.Scheme, firstHost.Host, func(call.Kind) bool { return true }),
			NewStatefulHost(secondHost.Scheme, secondHost.Host, func(call.Kind) bool { return true }),
		},
	})
	tr.sleep = func(context.Context, time.Duration) error {
		waits.Add(1)

		return nil
	}

	res, _, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429 once the shared budget is spent", res.StatusCode)
	}

	if firstCalls.Load() != 2 || secondCalls.Load() != 3 {
		t.Fatalf("calls = %d/%d, want 2/3", firstCalls.Load(), secondCalls.Load())
	}

	if waits.Load() != 3 {
		t.Fatalf("waits = %d, want 3", waits.Load())
	}
}

func TestRequestStillFailsOverOn5xx(t *testing.T) {
	t.Parallel()

	var firstCalls atomic.Int32

	first := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		firstCalls.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"error"}`))
	}))
	defer first.Close()

	second := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer second.Close()

	firstHost, _ := url.Parse(first.URL)
	secondHost, _ := url.Parse(second.URL)
	tr := New(Configuration{
		Hosts: []StatefulHost{
			NewStatefulHost(firstHost.Scheme, firstHost.Host, func(call.Kind) bool { return true }),
			NewStatefulHost(secondHost.Scheme, secondHost.Host, func(call.Kind) bool { return true }),
		},
	})
	tr.sleep = func(context.Context, time.Duration) error { return nil }

	res, _, err := tr.Request(context.Background(), newGetRequest(t, "/1/test"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", res.StatusCode)
	}

	if firstCalls.Load() != 1 {
		t.Fatalf("first host calls = %d", firstCalls.Load())
	}
}

func TestRequestStreamRetries429(t *testing.T) {
	t.Parallel()

	var (
		calls atomic.Int32
		waits []time.Duration
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)

		if calls.Load() == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"Too many requests"}`))

			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = io.WriteString(w, "data: ok\n\n")
	}))
	defer srv.Close()

	tr := newTestTransport(t, srv.URL, 0, nil)
	tr.sleep = func(_ context.Context, d time.Duration) error {
		waits = append(waits, d)

		return nil
	}

	res, err := tr.RequestStream(context.Background(), newGetRequest(t, "/1/events"), call.Read, RequestConfiguration{})
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if calls.Load() != 2 {
		t.Fatalf("calls = %d, want 2", calls.Load())
	}

	if len(waits) != 1 || waits[0] != time.Second {
		t.Fatalf("waits = %v", waits)
	}
}
