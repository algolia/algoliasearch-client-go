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

type Requester interface {
	Request(req *http.Request) (*http.Response, error)
}

type defaultRequester struct {
	client *http.Client
}

func newDefaultRequester() *defaultRequester {
	return &defaultRequester{
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					KeepAlive: DefaultKeepAliveDuration,
					Timeout:   DefaultConnectTimeout,
				}).Dial,
				DisableKeepAlives:   false,
				MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
				Proxy:               http.ProxyFromEnvironment,
				TLSHandshakeTimeout: DefaultTLSHandshakeTimeout,
			},
		},
	}
}

func (r *defaultRequester) Request(req *http.Request) (*http.Response, error) {
	return r.client.Do(req)
}
