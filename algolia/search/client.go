package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

const (
	DefaultMaxBatchSize = 1000
)

type Client struct {
	appID        string
	maxBatchSize int
	transport    *transport.Transport
}

func NewClient(config Configuration) *Client {
	var (
		hosts        []*transport.StatefulHost
		maxBatchSize int
	)

	if config.Hosts == nil {
		hosts = defaultHosts(config.AppID)
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsReadWrite))
		}
	}

	if config.MaxBatchSize == 0 {
		maxBatchSize = DefaultMaxBatchSize
	} else {
		maxBatchSize = config.MaxBatchSize
	}

	return &Client{
		appID:        config.AppID,
		maxBatchSize: maxBatchSize,
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

func (c *Client) InitIndex(indexName string) *Index {
	return newIndex(c.appID, indexName, c.maxBatchSize, c.transport)
}

func (c *Client) path(format string, a ...interface{}) string {
	return "/1" + fmt.Sprintf(format, a...)
}

func (c *Client) ListIndexes(opts ...interface{}) (res ListIndexesRes, err error) {
	path := c.path("/indexes")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}
