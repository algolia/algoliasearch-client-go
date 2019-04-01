package analytics

import (
	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Client struct {
	appID        string
	searchClient *search.Client
	transport    *transport.Transport
}

func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

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
		AppID:     config.AppID,
		APIKey:    config.APIKey,
		Requester: config.Requester,
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
		),
	}
}

func (c *Client) waitTaskSearchClient(index string, taskID int) error {
	return c.searchClient.InitIndex(index).WaitTask(taskID)
}
