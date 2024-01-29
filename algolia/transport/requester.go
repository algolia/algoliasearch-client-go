package transport

import (
	"net"
	"net/http"
	"time"
)

const (
	DefaultConnectTimeout      = 2 * time.Second
	DefaultKeepAliveDuration   = 5 * time.Minute
	DefaultMaxIdleConnsPerHost = 64
	DefaultTLSHandshakeTimeout = 2 * time.Second
)

// DefaultHTTPClient exposes the default *http.Client used by the different
// Client instances of the Algolia API client.
//
// Most users should not need to access this http.Client.
//
// This helper is only useful for users who wish to keep the default behavior of
// the Algolia API client (by wrapping it in a user-defined instance of
// transport.Requester) but would like to wrap it into a middleware layer or
// pass it to an HTTP interceptor.
func DefaultHTTPClient(connectTimeout *time.Duration) *http.Client {
	connectTimeoutValue := DefaultConnectTimeout
	if connectTimeout != nil {
		connectTimeoutValue = *connectTimeout
	}

	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				KeepAlive: DefaultKeepAliveDuration,
				Timeout:   connectTimeoutValue,
			}).DialContext,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
			Proxy:               http.ProxyFromEnvironment,
			TLSHandshakeTimeout: DefaultTLSHandshakeTimeout,
		},
	}
}

type Requester interface {
	Request(req *http.Request, timeout time.Duration, connectTimeout time.Duration) (*http.Response, error)
}

type defaultRequester struct {
	client *http.Client
}

func NewDefaultRequester(connectTimeout *time.Duration) *defaultRequester {
	return &defaultRequester{
		client: DefaultHTTPClient(connectTimeout),
	}
}

func (r *defaultRequester) Request(req *http.Request, _, _ time.Duration) (*http.Response, error) {
	return r.client.Do(req) //nolint:wrapcheck
}
