package insights

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/region"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Configuration struct {
	AppID        string
	APIKey       string
	Hosts        []string
	Requester    transport.Requester
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Region       region.Region
	Headers      map[string]string
}
