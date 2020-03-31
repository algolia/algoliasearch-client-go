package monitoring

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Client provides methods to interact with the Algolia Monitoring API.
type Client struct {
	transport *transport.Transport
}

func (c *Client) path(format string, a ...interface{}) string {
	return "/1" + fmt.Sprintf(format, a...)
}

// NewClient instantiates a new client able to interact with the Algolia
// Monitoring API.
func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Monitoring API.
func NewClientWithConfig(config Configuration) *Client {
	var hosts []*transport.StatefulHost

	if config.Hosts == nil {
		hosts = append(hosts, transport.NewStatefulHost("status.algolia.com", call.IsRead))
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsRead))
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
