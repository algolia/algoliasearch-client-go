package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

// MultipleBatch applies multiple indexing operations on potentially multiple
// indices in a single call.
func (c *Client) MultipleBatch(operations []BatchOperationIndexed, opts ...interface{}) (res MultipleBatchRes, err error) {
	body := map[string][]BatchOperationIndexed{"requests": operations}
	path := c.path("/indexes/*/batch")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = c.waitTask
	return
}

func (c *Client) waitTask(index string, taskID int64, opts ...interface{}) error {
	return c.InitIndex(index).WaitTask(taskID, opts...)
}

// MultipleGetObjects retrieves multiple objects from potentially multiple
// indices in a single call.
func (c *Client) MultipleGetObjects(requests []IndexedGetObject, objects interface{}, opts ...interface{}) (err error) {
	if len(requests) == 0 {
		return
	}
	res := getObjectsRes{objects}
	body := map[string]interface{}{"requests": requests}
	path := c.path("/indexes/*/objects")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// MultipleQueries performs multiple search queries on potentially multiple
// indices in a single call.
func (c *Client) MultipleQueries(queries []IndexedQuery, strategy string, opts ...interface{}) (res MultipleQueriesRes, err error) {
	body := newMultipleQueriesReq(queries, strategy)
	path := c.path("/indexes/*/queries")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}
