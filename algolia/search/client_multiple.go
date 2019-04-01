package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
)

func (c *Client) MultipleBatch(operations []BatchOperationIndexed, opts ...interface{}) (res MultipleBatchRes, err error) {
	body := map[string][]BatchOperationIndexed{"requests": operations}
	path := c.path("/indexes/*/batch")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = c.waitTask
	return
}

func (c *Client) waitTask(index string, taskID int) error {
	return c.InitIndex(index).WaitTask(taskID)
}

func (c *Client) MultipleGetObjects(requests []IndexedGetObject, objects interface{}, opts ...interface{}) (err error) {
	if requests == nil || len(requests) == 0 {
		return
	}
	res := getObjectsRes{objects}
	body := map[string]interface{}{"requests": requests}
	path := c.path("/indexes/*/objects")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

func (c *Client) MultipleQueries(queries []IndexedQuery, strategy string, opts ...interface{}) (res MultipleQueriesRes, err error) {
	body := newMultipleQueriesReq(queries, strategy)
	path := c.path("/indexes/*/queries")
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}
