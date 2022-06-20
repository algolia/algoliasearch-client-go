package personalization

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/region"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Client provides methods to interact with the Algolia Personalization API.
type Client struct {
	transport *transport.Transport
}

// NewClient instantiates a new client able to interact with the Algolia
// Personalization API.
func NewClient(appID, apiKey string, region region.Region) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
			Region: region,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Personalization API.
func NewClientWithConfig(config Configuration) *Client {
	var hosts []*transport.StatefulHost

	if config.Hosts == nil {
		hosts = defaultHosts(config.Region)
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsReadWrite))
		}
	}

	return &Client{
		transport: transport.New(
			hosts,
			config.Requester,
			config.AppID,
			config.APIKey,
			config.ReadTimeout,
			config.WriteTimeout,
			config.Headers,
			config.ExtraUserAgent,
			compression.None,
		),
	}
}
