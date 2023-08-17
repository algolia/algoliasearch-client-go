package search

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Index provides methods to interact with the Algolia Search API on a single
// index.
type Index struct {
	appID        string
	name         string
	maxBatchSize int
	client       *Client
	transport    *transport.Transport
}

func newIndex(client *Client, name string) *Index {
	return &Index{
		appID:        client.appID,
		client:       client,
		name:         name,
		maxBatchSize: client.maxBatchSize,
		transport:    client.transport,
	}
}

func (i *Index) path(format string, a ...interface{}) string {
	prefix := fmt.Sprintf("/1/indexes/%s", url.QueryEscape(i.name))
	suffix := fmt.Sprintf(format, a...)
	return prefix + suffix
}

// WaitTask blocks until the task identified by the given taskID is completed on
// Algolia engine.
func (i *Index) WaitTask(taskID int64, opts ...interface{}) error {
	return waitWithRetry(func() (bool, error) {
		res, err := i.GetStatus(taskID, opts...)
		if err != nil {
			return true, err
		}
		return res.Status == taskPublished, nil
	}, iopt.ExtractWaitConfiguration(opts...))
}

// WaitRecommendTask blocks until the task identified by the given Recommend-scope
// taskID is completed on Algolia engine.
func (i *Index) WaitRecommendTask(taskID int64, opts ...interface{}) error {
	return waitWithRetry(func() (bool, error) {
		res, err := i.GetRecommendStatus(taskID, opts...)
		if err != nil {
			return true, err
		}
		return res.Status == taskPublished, nil
	}, iopt.ExtractWaitConfiguration(opts...))
}

func (i *Index) operation(destination, op string, opts ...interface{}) (res UpdateTaskRes, err error) {
	var scopes []string
	if opt := iopt.ExtractScopes(opts...); opt != nil {
		scopes = opt.Get()
	}
	req := IndexOperation{
		Destination: destination,
		Operation:   op,
		Scopes:      scopes,
	}
	path := i.path("/operation")
	err = i.transport.Request(&res, http.MethodPost, path, req, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// GetAppID returns the Algolia application ID on where the current index
// leaves.
func (i *Index) GetAppID() string {
	return i.appID
}

// GetName returns the current index name.
func (i *Index) GetName() string {
	return i.name
}

// ClearObjects deletes all the records of the index.
func (i *Index) ClearObjects(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// Delete removes the entire index. After this call, new indexing calls can be
// sent with the same index instance.
func (i *Index) Delete(opts ...interface{}) (res DeleteTaskRes, err error) {
	path := i.path("")
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// GetStatus retrieves the task status according to the Algolia engine for the
// given task.
func (i *Index) GetStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error) {
	path := i.path("/task/%d", taskID)
	err = i.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// GetRecommendStatus retrieves the task status according to the Algolia engine
// for the given Recommend task.
func (i *Index) GetRecommendStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error) {
	// modelName is arbitrarily defined as related-products because this parameter does not matter anymore
	modelName := "related-products"
	path := i.path("/%s/task/%d", modelName, taskID)
	err = i.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// Exists returns whether an initialized index exists or not, along with a nil
// error. When encountering a network error, a non-nil error is returned along
// with false.
func (i *Index) Exists() (bool, error) {
	_, err := i.GetSettings()
	if err == nil {
		return true, nil
	}
	if _, ok := errs.IsAlgoliaErrWithCode(err, http.StatusNotFound); ok {
		return false, nil
	}
	return false, err
}
