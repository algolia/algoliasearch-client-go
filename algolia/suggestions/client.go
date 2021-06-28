package suggestions

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Client provides methods to interact with the Algolia Query Suggestions API.
type Client struct {
	appID        string
	searchClient *search.Client
	transport    *transport.Transport
}

// NewClient instantiates a new client able to interact with the Algolia
// Query Suggestions API.
func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Algolia Query Suggestions API.
func NewClientWithConfig(config Configuration) *Client {
	var hosts []*transport.StatefulHost

	if config.Hosts == nil {
		hosts = defaultHosts(config.Region)
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsReadWrite))
		}
	}

	searchConfig := search.Configuration{
		AppID:          config.AppID,
		APIKey:         config.APIKey,
		Requester:      config.Requester,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		Headers:        config.Headers,
		ExtraUserAgent: config.ExtraUserAgent,
	}

	return &Client{
		appID:        config.AppID,
		searchClient: search.NewClientWithConfig(searchConfig),
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

func (c *Client) path(format string, a ...interface{}) string { //nolint:unparam
	return "/1" + fmt.Sprintf(format, a...)
}
