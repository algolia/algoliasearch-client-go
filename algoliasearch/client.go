package algoliasearch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	transport *Transport
}

// NewClient instantiates a new `Client` from the provided `appID` and
// `apiKey`. Default hosts are used for the transport layer.
func NewClient(appID, apiKey string) Client {
	return &client{
		transport: NewTransport(appID, apiKey),
	}
}

// NewClientWithHosts instantiates a new `Client` from the provided `appID` and
// `apiKey`. The transport layers' hosts are initialized with the given
// `hosts`.
func NewClientWithHosts(appID, apiKey string, hosts []string) Client {
	return &client{
		transport: NewTransportWithHosts(appID, apiKey, hosts),
	}
}

func (c *client) SetExtraHeader(key, value string) {
	c.transport.setExtraHeader(key, value)
}

func (c *client) SetTimeout(connectTimeout, readTimeout int) {
	c.transport.setTimeout(
		time.Duration(connectTimeout)*time.Millisecond,
		time.Duration(readTimeout)*time.Millisecond,
	)
}

func (c *client) SetHTTPClient(client *http.Client) {
	c.transport.httpClient = client
}

func (c *client) ListIndexes() (indexes []IndexRes, err error) {
	return c.ListIndexesWithRequestOptions(nil)
}

func (c *client) ListIndexesWithRequestOptions(opts *RequestOptions) (indexes []IndexRes, err error) {
	var res listIndexesRes

	err = c.request(&res, "GET", "/1/indexes", nil, read, opts)
	indexes = res.Items
	return
}

func (c *client) InitIndex(name string) Index {
	return NewIndex(name, c)
}

func (c *client) ListKeys() (keys []Key, err error) {
	return c.ListKeysWithRequestOptions(nil)
}

func (c *client) ListKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error) {
	var res listKeysRes
	err = c.request(&res, "GET", "/1/keys", nil, read, opts)
	keys = res.Keys
	return
}

func (c *client) MoveIndex(source, destination string) (UpdateTaskRes, error) {
	return c.MoveIndexWithRequestOptions(source, destination, nil)
}

func (c *client) MoveIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	index := c.InitIndex(source)
	return index.MoveWithRequestOptions(destination, opts)
}

func (c *client) CopyIndex(source, destination string) (UpdateTaskRes, error) {
	return c.CopyIndexWithRequestOptions(source, destination, nil)
}

func (c *client) CopyIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	index := c.InitIndex(source)
	return index.CopyWithRequestOptions(destination, opts)
}

func (c *client) DeleteIndex(name string) (res DeleteTaskRes, err error) {
	return c.DeleteIndexWithRequestOptions(name, nil)
}

func (c *client) DeleteIndexWithRequestOptions(name string, opts *RequestOptions) (res DeleteTaskRes, err error) {
	index := c.InitIndex(name)
	return index.DeleteWithRequestOptions(opts)
}

func (c *client) ClearIndex(name string) (res UpdateTaskRes, err error) {
	return c.ClearIndexWithRequestOptions(name, nil)
}

func (c *client) ClearIndexWithRequestOptions(name string, opts *RequestOptions) (res UpdateTaskRes, err error) {
	index := c.InitIndex(name)
	return index.ClearWithRequestOptions(opts)
}

func (c *client) AddUserKey(ACL []string, params Map) (AddKeyRes, error) {
	return c.AddAPIKey(ACL, params)
}

func (c *client) AddAPIKey(ACL []string, params Map) (res AddKeyRes, err error) {
	return c.AddAPIKeyWithRequestOptions(ACL, params, nil)
}

func (c *client) AddAPIKeyWithRequestOptions(ACL []string, params Map, opts *RequestOptions) (res AddKeyRes, err error) {
	req := duplicateMap(params)
	req["acl"] = ACL

	if err = checkKey(req); err != nil {
		return
	}

	err = c.request(&res, "POST", "/1/keys/", req, read, opts)
	return
}

func (c *client) UpdateUserKey(key string, params Map) (UpdateKeyRes, error) {
	return c.UpdateAPIKey(key, params)
}

func (c *client) UpdateAPIKey(key string, params Map) (res UpdateKeyRes, err error) {
	return c.UpdateAPIKeyWithRequestOptions(key, params, nil)
}

func (c *client) UpdateAPIKeyWithRequestOptions(key string, params Map, opts *RequestOptions) (res UpdateKeyRes, err error) {
	if err = checkKey(params); err != nil {
		return
	}

	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "PUT", path, params, write, opts)
	return
}

func (c *client) GetUserKey(key string) (Key, error) {
	return c.GetAPIKey(key)
}

func (c *client) GetAPIKey(key string) (res Key, err error) {
	return c.GetAPIKeyWithRequestOptions(key, nil)
}

func (c *client) GetAPIKeyWithRequestOptions(key string, opts *RequestOptions) (res Key, err error) {
	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "GET", path, nil, read, opts)
	return
}

func (c *client) DeleteUserKey(key string) (DeleteRes, error) {
	return c.DeleteAPIKey(key)
}

func (c *client) DeleteAPIKey(key string) (res DeleteRes, err error) {
	return c.DeleteAPIKeyWithRequestOptions(key, nil)
}

func (c *client) DeleteAPIKeyWithRequestOptions(key string, opts *RequestOptions) (res DeleteRes, err error) {
	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (c *client) GetLogs(params Map) (logs []LogRes, err error) {
	return c.GetLogsWithRequestOptions(params, nil)
}

func (c *client) GetLogsWithRequestOptions(params Map, opts *RequestOptions) (logs []LogRes, err error) {
	var res getLogsRes

	if err = checkGetLogs(params); err != nil {
		return
	}

	err = c.request(&res, "GET", "/1/logs", params, write, opts)
	logs = res.Logs
	return
}

func (c *client) MultipleQueries(queries []IndexedQuery, strategy string) (res []MultipleQueryRes, err error) {
	return c.MultipleQueriesWithRequestOptions(queries, strategy, nil)
}

func (c *client) MultipleQueriesWithRequestOptions(queries []IndexedQuery, strategy string, opts *RequestOptions) (res []MultipleQueryRes, err error) {
	if strategy == "" {
		strategy = "none"
	}

	for _, q := range queries {
		if err = checkQuery(q.Params); err != nil {
			return
		}
	}

	requests := make([]map[string]string, len(queries))
	for i, q := range queries {
		requests[i] = map[string]string{
			"indexName": q.IndexName,
			"params":    encodeMap(q.Params),
		}
	}

	body := Map{
		"requests": requests,
		"strategy": strategy,
	}

	var m multipleQueriesRes
	err = c.request(&m, "POST", "/1/indexes/*/queries", body, search, opts)
	res = m.Results
	return
}

func (c *client) Batch(operations []BatchOperationIndexed) (res MultipleBatchRes, err error) {
	return c.BatchWithRequestOptions(operations, nil)
}

func (c *client) BatchWithRequestOptions(operations []BatchOperationIndexed, opts *RequestOptions) (res MultipleBatchRes, err error) {
	// TODO: Use check functions of index.go

	request := map[string][]BatchOperationIndexed{
		"requests": operations,
	}

	err = c.request(&res, "POST", "/1/indexes/*/batch", request, write, opts)
	return
}

func (c *client) request(res interface{}, method, path string, body interface{}, typeCall int, opts *RequestOptions) error {
	r, err := c.transport.request(method, path, body, typeCall, opts)
	if err != nil {
		return err
	}

	return json.Unmarshal(r, res)
}
