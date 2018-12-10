package algoliasearch

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch/call"
	"github.com/stretchr/testify/require"
)

type FakeNetError struct{}

func (e *FakeNetError) Timeout() bool   { return false }
func (e *FakeNetError) Temporary() bool { return false }
func (e *FakeNetError) Error() string   { return "fake net.error" }

var fakeNetError net.Error = &FakeNetError{}

func TestRetryStrategy_Sequential(t *testing.T) {
	strategy := NewRetryStrategy("latency", []string{"example.com"})

	// Try sucessful read/write/analytics calls with provided hosts
	{

		hosts := strategy.GetTryableHosts(call.Read)
		expected := &tryableHost{"example.com", 5 * time.Second}
		require.ElementsMatch(t, []TryableHost{expected}, hosts)
		require.Equal(t, Success, strategy.Decide(expected, 200, nil))

		hosts = strategy.GetTryableHosts(call.Write)
		expected = &tryableHost{"example.com", 30 * time.Second}
		require.ElementsMatch(t, []TryableHost{expected}, hosts)
		require.Equal(t, Success, strategy.Decide(expected, 200, nil))

		hosts = strategy.GetTryableHosts(call.Analytics)
		expected = &tryableHost{"analytics.algolia.com", 30 * time.Second}
		require.ElementsMatch(t, []TryableHost{expected}, hosts)
		require.Equal(t, Success, strategy.Decide(expected, 200, nil))
	}

	strategy = NewRetryStrategy("latency", nil)

	// Try read calls until exhaustion
	{
		expected := []TryableHost{
			&tryableHost{"latency-dsn.algolia.net", 5 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 5 * time.Second},
		}
		hosts := strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		for _, h := range hosts {
			require.Equal(t, Success, strategy.Decide(h, 200, nil))
		}

		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, context.DeadlineExceeded))

		expected = []TryableHost{
			&tryableHost{"latency-dsn.algolia.net", 10 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 5 * time.Second},
		}
		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))

		expected = []TryableHost{
			hosts[1],
			hosts[2],
			hosts[3],
		}
		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))

		expected = []TryableHost{
			hosts[1],
			hosts[2],
		}
		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 300, nil))

		expected = []TryableHost{
			hosts[1],
		}
		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Failure, strategy.Decide(hosts[0], 200, errors.New("JSON decoding error")))

		hosts = strategy.GetTryableHosts(call.Read)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 300, nil))

		expected = []TryableHost{
			&tryableHost{"latency-dsn.algolia.net", 5 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 5 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 5 * time.Second},
		}
		hosts = strategy.GetTryableHosts(call.Read)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))
	}

	// Try write calls until exhaustion
	{
		expected := []TryableHost{
			&tryableHost{"latency.algolia.net", 30 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 30 * time.Second},
		}
		hosts := strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		for _, h := range hosts {
			require.Equal(t, Success, strategy.Decide(h, 200, nil))
		}

		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, context.DeadlineExceeded))

		expected = []TryableHost{
			&tryableHost{"latency.algolia.net", 60 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 30 * time.Second},
		}
		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))

		expected = []TryableHost{
			hosts[1],
			hosts[2],
			hosts[3],
		}
		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))

		expected = []TryableHost{
			hosts[1],
			hosts[2],
		}
		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 300, nil))

		expected = []TryableHost{
			hosts[1],
		}
		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Failure, strategy.Decide(hosts[0], 200, errors.New("JSON decoding error")))

		hosts = strategy.GetTryableHosts(call.Write)
		require.ElementsMatch(t, expected, hosts)
		require.Equal(t, Retry, strategy.Decide(hosts[0], 300, nil))

		expected = []TryableHost{
			&tryableHost{"latency.algolia.net", 30 * time.Second},
			&tryableHost{"latency-1.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-2.algolianet.com", 30 * time.Second},
			&tryableHost{"latency-3.algolianet.com", 30 * time.Second},
		}
		hosts = strategy.GetTryableHosts(call.Write)
		require.Equal(t, expected[0], hosts[0])
		require.Equal(t, Retry, strategy.Decide(hosts[0], 0, fakeNetError))
	}
}

func TestRetryStrategy_Concurrent(t *testing.T) {
	var wg sync.WaitGroup
	strategy := NewRetryStrategy("latency", []string{"example.com"})

	var (
		readTimeout      = 1 * time.Second
		writeTimeout     = 2 * time.Second
		analyticsTimeout = 3 * time.Second
		insightsTimeout  = 4 * time.Second
	)
	strategy.SetTimeouts(readTimeout, writeTimeout, analyticsTimeout, insightsTimeout)

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			expected := []TryableHost{&tryableHost{"example.com", readTimeout}}
			hosts := strategy.GetTryableHosts(call.Read)
			require.ElementsMatch(t, expected, hosts)

			expected = []TryableHost{&tryableHost{"example.com", writeTimeout}}
			hosts = strategy.GetTryableHosts(call.Write)
			require.ElementsMatch(t, expected, hosts)

			expected = []TryableHost{&tryableHost{"analytics.algolia.com", analyticsTimeout}}
			hosts = strategy.GetTryableHosts(call.Analytics)
			require.ElementsMatch(t, expected, hosts)

			expected = []TryableHost{&tryableHost{"insights.algolia.io", insightsTimeout}}
			hosts = strategy.GetTryableHosts(call.Insights)
			require.ElementsMatch(t, expected, hosts)

		}(&wg)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			strategy.SetTimeouts(readTimeout, writeTimeout, analyticsTimeout, insightsTimeout)
		}(&wg)
	}

	wg.Wait()
}

func TestRetryStrategy_shuffle(t *testing.T) {
	var hosts []*statefulHost
	for i := 0; i < 1000; i++ {
		hosts = append(hosts, &statefulHost{host: fmt.Sprintf("%d", i)})
	}

	shuffled := shuffle(hosts)
	require.ElementsMatch(t, shuffled, hosts)

	shuffledElementsCount := 0
	for i := 0; i < len(hosts); i++ {
		if hosts[i].host != shuffled[i].host {
			shuffledElementsCount++
		}
	}
	require.NotEqual(t, 0, shuffledElementsCount)
}

func TestRetryStrategy_isNetworkError(t *testing.T) {
	for _, c := range []struct {
		err            error
		isNetworkError bool
	}{
		{nil, false},
		{errors.New("regular error"), false},
		{context.DeadlineExceeded, false},
		{fakeNetError, true},
	} {
		require.Equal(t, c.isNetworkError, isNetworkError(c.err))
	}
}

func TestRetryStrategy_isTimeoutError(t *testing.T) {
	for _, c := range []struct {
		err            error
		isTimeoutError bool
	}{
		{nil, false},
		{errors.New("regular error"), false},
		{fakeNetError, false},
		{context.DeadlineExceeded, true},
		{fmt.Errorf("custom error: %s", context.DeadlineExceeded), true},
	} {
		require.Equal(t, c.isTimeoutError, isTimeoutError(c.err))
	}
}

func TestRetryStrategy_is2xx(t *testing.T) {
	for i := -1000; i < 1000; i++ {
		if 200 <= i && i < 300 {
			require.True(t, is2xx(i))
		} else {
			require.False(t, is2xx(i))
		}
	}
}

func TestRetryStrategy_is4xx(t *testing.T) {
	for i := -1000; i < 1000; i++ {
		if 400 <= i && i < 500 {
			require.True(t, is4xx(i))
		} else {
			require.False(t, is4xx(i))
		}
	}
}

func TestRetryStrategy_isZero(t *testing.T) {
	for i := -1000; i < 1000; i++ {
		if i == 0 {
			require.True(t, isZero(i))
		} else {
			require.False(t, isZero(i))
		}
	}
}
