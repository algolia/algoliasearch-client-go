package recommendation

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/region"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Deprecated: use personalization.Configuration instead
type Configuration struct {
	AppID          string
	APIKey         string
	Hosts          []string
	Requester      transport.Requester
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Region         region.Region
	Headers        map[string]string
	ExtraUserAgent string
}
