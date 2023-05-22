package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

const (
	// DefaultMaxBatchSize defines the default maximum batch size to be used to
	// automatically split record batches when using Index.SaveObjects.
	DefaultMaxBatchSize = 1000
)

// Client provides methods to interact with the Algolia Search API on multiple
// indices which belong to the same Algolia application.
type Client struct {
	appID        string
	maxBatchSize int
	transport    *transport.Transport
}

// NewClient instantiates a new client able to interact with the Algolia
// Search API on multiple indices which belong to the same Algolia application.
func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:       appID,
			APIKey:      apiKey,
			Compression: compression.None,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Algolia Search API on multiple indices which belong to the same Algolia
// application.
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

	if config.MaxBatchSize <= 0 {
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
			config.ExtraUserAgent,
			config.Compression,
		),
	}
}

// InitIndex instantiates a new index able to interact with the Algolia
// Search API on a single index.
func (c *Client) InitIndex(indexName string) *Index {
	return newIndex(c, indexName)
}

func (c *Client) path(format string, a ...interface{}) string {
	return "/1" + fmt.Sprintf(format, a...)
}

// ListIndices lists all the indices of the Algolia application in a single
// call.
func (c *Client) ListIndices(opts ...interface{}) (res ListIndicesRes, err error) {
	path := c.path("/indexes")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// GetLogs returns the most recent information logs of the Algolia application.
func (c *Client) GetLogs(opts ...interface{}) (res GetLogsRes, err error) {
	if offset := iopt.ExtractOffset(opts...); offset != nil {
		opts = opt.InsertExtraURLParam(opts, "offset", offset.Get())
	}
	if length := iopt.ExtractLength(opts...); length != nil {
		opts = opt.InsertExtraURLParam(opts, "length", length.Get())
	}
	if t := iopt.ExtractType(opts...).Get(); len(t) > 0 {
		opts = opt.InsertExtraURLParam(opts, "type", t[0])
	}
	if indexName := iopt.ExtractIndexName(opts...); indexName != nil {
		opts = opt.InsertExtraURLParam(opts, "indexName", indexName.Get())
	}
	path := c.path("/logs")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// CustomRequest is a low-level function which build a request from the given
// parameters and send it through the requester, making use of the underlying
// retry strategy.
func (c *Client) CustomRequest(
	res interface{},
	method string,
	path string,
	body interface{},
	k call.Kind,
	opts ...interface{},
) error {
	return c.transport.Request(&res, method, path, body, k, opts...)
}

// GetStatus retrieves the task status according to the Algolia engine for the
// given task.
func (c *Client) GetStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error) {
	path := c.path("/task/%d", taskID)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// GetRecommendStatus retrieves the task status according to the Algolia engine
// for the given Recommend task.
func (c *Client) GetRecommendStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error) {
	// modelName is arbitrarily defined as related-products because this parameter does not matter anymore
	modelName := "related-products"
	path := c.path("/%s/task/%d", modelName, taskID)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// WaitTask blocks until the task identified by the given taskID is completed on
// Algolia engine.
func (c *Client) WaitTask(taskID int64, opts ...interface{}) error {
	return waitWithRetry(func() (bool, error) {
		res, err := c.GetStatus(taskID, opts...)
		if err != nil {
			return true, err
		}
		return res.Status == taskPublished, nil
	}, iopt.ExtractWaitConfiguration(opts...))
}

// WaitRecommendTask blocks until the task identified by the given Recommend-scope
// taskID is completed on Algolia engine.
func (c *Client) WaitRecommendTask(taskID int64, opts ...interface{}) error {
	return waitWithRetry(func() (bool, error) {
		res, err := c.GetRecommendStatus(taskID, opts...)
		if err != nil {
			return true, err
		}
		return res.Status == taskPublished, nil
	}, iopt.ExtractWaitConfiguration(opts...))
}
