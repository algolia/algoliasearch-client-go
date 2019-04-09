package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
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

func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

func NewClientWithConfig(config Configuration) *Client {
	var (
		hosts        []*transport.StatefulHost
		maxBatchSize int
	)

	if len(config.Hosts) == 0 {
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
	return newIndex(c, indexName)
}

func (c *Client) path(format string, a ...interface{}) string {
	return "/1" + fmt.Sprintf(format, a...)
}

func (c *Client) ListIndexes(opts ...interface{}) (res ListIndexesRes, err error) {
	path := c.path("/indexes")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) GetLogs(opts ...interface{}) (res GetLogsRes, err error) {
	if offset := iopt.ExtractOffset(opts...); offset != nil {
		opts = opt.InsertExtraURLParam(opts, "offset", offset.Get())
	}
	if length := iopt.ExtractLength(opts...); length != nil {
		opts = opt.InsertExtraURLParam(opts, "length", length.Get())
	}
	if t := iopt.ExtractType(opts...); t != nil {
		opts = opt.InsertExtraURLParam(opts, "type", t.Get())
	}
	if indexName := iopt.ExtractIndexName(opts...); indexName != nil {
		opts = opt.InsertExtraURLParam(opts, "indexName", indexName.Get())
	}
	path := c.path("/logs")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}
