package algoliasearch

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch/call"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

type Outcome int

const (
	DefaultReadTimeout      = 5 * time.Second
	DefaultWriteTimeout     = 30 * time.Second
	DefaultAnalyticsTimeout = 30 * time.Second
	DefaultInsightsTimeout  = 30 * time.Second

	Success Outcome = iota
	Failure
	Retry
)

type TryableHost interface {
	Host() string
	Timeout() time.Duration
}

type tryableHost struct {
	host    string
	timeout time.Duration
}

func (h *tryableHost) Host() string           { return h.host }
func (h *tryableHost) Timeout() time.Duration { return h.timeout }
func (h *tryableHost) String() string         { return fmt.Sprintf("tryableHost{%s,%s}", h.host, h.timeout) }

type RetryStrategy interface {
	// GetTryableHosts returns the slice of host to try to send the request to.
	GetTryableHosts(k call.Kind) []TryableHost

	// Decide returns an Outcome defining if the call have succeded, or failed
	// or to be retried.
	Decide(h TryableHost, code int, err error) Outcome

	// SetTimeouts updates the internal timeouts for read, write (i.e.
	// indexing) and analytics calls. Negative values are simply ignored,
	// leaving the original timeouts unchanged.
	SetTimeouts(read, write, analytics, insights time.Duration)
}

type retryStrategy struct {
	sync.RWMutex
	hosts            []*statefulHost
	readTimeout      time.Duration
	writeTimeout     time.Duration
	analyticsTimeout time.Duration
	insightsTimeout  time.Duration
}

type statefulHost struct {
	host       string
	isDown     bool
	retryCount int
	lastUpdate time.Time
	accept     func(k call.Kind) bool
}

func (h *statefulHost) String() string {
	return fmt.Sprintf(
		"statefulHost{host:%s, isDown: %t, retryCount:%d}",
		h.host,
		h.isDown,
		h.retryCount,
	)
}

func (h *statefulHost) reset() {
	h.isDown = false
	h.lastUpdate = time.Now()
	h.retryCount = 0
}

func NewRetryStrategy(appID string, providedHosts []string) *retryStrategy {
	var allHosts []*statefulHost
	now := time.Now()

	if providedHosts != nil && len(providedHosts) > 0 {
		for _, h := range providedHosts {
			allHosts = append(allHosts, &statefulHost{host: h, lastUpdate: now, accept: call.IsReadWrite})
		}
	} else {
		allHosts = append(allHosts, &statefulHost{host: appID + "-dsn.algolia.net", lastUpdate: now, accept: call.IsRead})
		allHosts = append(allHosts, &statefulHost{host: appID + ".algolia.net", lastUpdate: now, accept: call.IsWrite})
		allHosts = append(allHosts, shuffle(
			[]*statefulHost{
				&statefulHost{host: appID + "-1.algolianet.com", lastUpdate: now, accept: call.IsReadWrite},
				&statefulHost{host: appID + "-2.algolianet.com", lastUpdate: now, accept: call.IsReadWrite},
				&statefulHost{host: appID + "-3.algolianet.com", lastUpdate: now, accept: call.IsReadWrite},
			},
		)...)
	}
	allHosts = append(allHosts, &statefulHost{host: "analytics.algolia.com", lastUpdate: now, accept: call.IsAnalytics})
	allHosts = append(allHosts, &statefulHost{host: "insights.algolia.io", lastUpdate: now, accept: call.IsInsights})

	return &retryStrategy{
		hosts:            allHosts,
		readTimeout:      DefaultReadTimeout,
		writeTimeout:     DefaultWriteTimeout,
		analyticsTimeout: DefaultAnalyticsTimeout,
		insightsTimeout:  DefaultInsightsTimeout,
	}
}

func (s *retryStrategy) GetTryableHosts(k call.Kind) []TryableHost {
	s.resetExpiredHosts()
	s.displayState()

	s.Lock()
	defer s.Unlock()

	var baseTimeout time.Duration
	switch k {
	case call.Read:
		baseTimeout = s.readTimeout
	case call.Write:
		baseTimeout = s.writeTimeout
	case call.Analytics:
		baseTimeout = s.analyticsTimeout
	case call.Insights:
		baseTimeout = s.insightsTimeout
	default:
		return nil
	}

	var hosts []TryableHost
	for _, h := range s.hosts {
		if !h.isDown && h.accept(k) {
			hosts = append(hosts, &tryableHost{h.host, baseTimeout * time.Duration(h.retryCount+1)})
		}
	}
	if len(hosts) > 0 {
		return hosts
	}
	for _, h := range s.hosts {
		if h.accept(k) {
			h.reset()
			hosts = append(hosts, &tryableHost{h.host, baseTimeout})
		}
	}
	return hosts
}

func (s *retryStrategy) Decide(h TryableHost, code int, err error) Outcome {
	if err == nil && is2xx(code) {
		debug("* RETRY STRATEGY DECISION host=%s code=%d err=%v -> SUCCESS", h.Host(), code, err)
		s.markUp(h.Host())
		return Success
	}

	if isTimeoutError(err) {
		debug("* RETRY STRATEGY DECISION host=%s code=%d err=%v -> RETRY (TIMEOUT)", h.Host(), code, err)
		s.markTimeouted(h.Host())
		return Retry
	}

	if !(isZero(code) || is4xx(code) || is2xx(code)) || isNetworkError(err) {
		debug("* RETRY STRATEGY DECISION host=%s code=%d err=%v -> RETRY (DOWN)", h.Host(), code, err)
		s.markDown(h.Host())
		return Retry
	}

	debug("* RETRY STRATEGY DECISION host=%s code=%d err=%v -> FAILURE", h.Host(), code, err)
	return Failure
}

func (s *retryStrategy) SetTimeouts(read, write, analytics, insights time.Duration) {
	s.Lock()
	defer s.Unlock()

	if read > 0 {
		s.readTimeout = read
	}
	if write > 0 {
		s.writeTimeout = write
	}
	if analytics > 0 {
		s.analyticsTimeout = analytics
	}
	if insights > 0 {
		s.insightsTimeout = insights
	}
}

func (s *retryStrategy) markUp(host string)        { s.update(host, false, false) }
func (s *retryStrategy) markDown(host string)      { s.update(host, true, false) }
func (s *retryStrategy) markTimeouted(host string) { s.update(host, false, true) }
func (s *retryStrategy) update(host string, isDown, isTimeout bool) {
	s.Lock()
	defer s.Unlock()

	for _, h := range s.hosts {
		if h.host == host {
			h.isDown = isDown
			h.lastUpdate = time.Now()
			if isTimeout {
				h.retryCount++
			} else {
				h.retryCount = 0
			}
			return
		}
	}
}

func (s *retryStrategy) resetExpiredHosts() {
	s.Lock()
	defer s.Unlock()

	for _, h := range s.hosts {
		if h.isDown && time.Since(h.lastUpdate) > 5*time.Minute {
			h.reset()
		}
	}
}

func (s *retryStrategy) displayState() {
	s.RLock()
	defer s.RUnlock()

	for _, h := range s.hosts {
		debug("* RETRY STRATEGY STATE %s", h)
	}
}

func shuffle(hosts []*statefulHost) []*statefulHost {
	if hosts == nil {
		return nil
	}
	shuffled := make([]*statefulHost, len(hosts))
	for i, v := range rand.Perm(len(hosts)) {
		shuffled[i] = hosts[v]
	}
	return shuffled
}

func isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(net.Error)
	// We need to ensure that the error is a net.Error but not a
	// context.DeadlineExceeded error (which is actually a net.Error), because
	// we do not want to consider context.DeadlineExceeded as an error.
	return ok && !isTimeoutError(err)
}

func isTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), context.DeadlineExceeded.Error())
}

func isZero(code int) bool { return code == 0 }
func is2xx(code int) bool  { return 200 <= code && code < 300 }
func is4xx(code int) bool  { return 400 <= code && code < 500 }
