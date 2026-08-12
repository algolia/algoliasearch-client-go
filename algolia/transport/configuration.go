package transport

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/compression"
)

type Configuration struct {
	AppID  string
	ApiKey string //nolint:staticcheck

	Hosts                           []StatefulHost
	DefaultHeader                   map[string]string
	UserAgent                       string
	Requester                       Requester
	ReadTimeout                     time.Duration
	WriteTimeout                    time.Duration
	ConnectTimeout                  time.Duration
	Compression                     compression.Compression
	ExposeIntermediateNetworkErrors bool

	// RequestIDEnabled makes the transport send a Request-ID header, minted
	// once per call and reused across its retry attempts, so that Algolia
	// support can tie the attempts of one request together. When nil, the
	// generated per-client default applies (enabled on the APIs that support
	// it); a non-nil value always wins, for any client. A Request-ID supplied
	// through request options or DefaultHeader is never overwritten.
	RequestIDEnabled *bool
}

type RequestConfiguration struct {
	ReadTimeout    *time.Duration
	WriteTimeout   *time.Duration
	ConnectTimeout *time.Duration
}
