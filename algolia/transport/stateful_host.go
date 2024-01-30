package transport

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
)

const (
	DefaultResetPeriod = 5 * time.Minute
)

type StatefulHost struct {
	scheme     string
	host       string
	isDown     bool
	retryCount int
	lastUpdate time.Time
	accept     func(k call.Kind) bool
}

func NewStatefulHost(scheme string, host string, accept func(k call.Kind) bool) StatefulHost {
	return StatefulHost{
		scheme:     scheme,
		host:       host,
		isDown:     false,
		retryCount: 0,
		lastUpdate: time.Now(),
		accept:     accept,
	}
}

func (h *StatefulHost) markUp() {
	h.lastUpdate = time.Now()
	h.isDown = false
	h.retryCount = 0
}

func (h *StatefulHost) markTimeout() {
	h.lastUpdate = time.Now()
	h.retryCount++
}

func (h *StatefulHost) markDown() {
	h.lastUpdate = time.Now()
	h.isDown = true
	h.retryCount = 0
}

func (h *StatefulHost) isExpired() bool {
	return h.isDown && time.Since(h.lastUpdate) > DefaultResetPeriod
}

func (h *StatefulHost) reset() {
	h.lastUpdate = time.Now()
	h.isDown = false
	h.retryCount = 0
}
