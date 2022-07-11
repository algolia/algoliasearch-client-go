package recommend

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Client provides methods to interact with the Algolia Recommend API.
type Client struct {
	transport *transport.Transport
}

// NewClient instantiates a new client able to interact with the Algolia
// Recommend API.
func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Algolia Recommend API.
func NewClientWithConfig(config Configuration) *Client {
	var (
		hosts []*transport.StatefulHost
	)

	if len(config.Hosts) == 0 {
		hosts = defaultHosts(config.AppID)
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
			config.Compression,
			config.DisableSSL,
		),
	}
}

func defaultHosts(appID string) (hosts []*transport.StatefulHost) {
	hosts = append(hosts, transport.NewStatefulHost(appID+"-dsn.algolia.net", call.IsRead))
	hosts = append(hosts, transport.NewStatefulHost(appID+".algolia.net", call.IsWrite))
	hosts = append(hosts, transport.Shuffle(
		[]*transport.StatefulHost{
			transport.NewStatefulHost(appID+"-1.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-2.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-3.algolianet.com", call.IsReadWrite),
		},
	)...)
	return
}
