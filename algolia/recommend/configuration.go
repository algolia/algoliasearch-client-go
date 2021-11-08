package recommend

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Configuration contains all the different parameters one can change to
// instantiate a new client for the Recommend API.
type Configuration struct {
	AppID          string
	APIKey         string
	Hosts          []string
	Requester      transport.Requester
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Headers        map[string]string
	ExtraUserAgent string
	Compression    compression.Compression
}
