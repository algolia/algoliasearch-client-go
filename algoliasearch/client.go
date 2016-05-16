package algoliasearch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"time"
)

// Client is a representation of an Algolia application. Once initialized it
// allows manipulations over the indexes of the application as well as
// network related parameters.
type client struct {
	transport *Transport
}

// NewClient creates a new Client from the provided `appID` and `apiKey`. The
// default hosts are used for the transport layer.
func NewClient(appID, apiKey string) Client {
	return &client{
		transport: NewTransport(appID, apiKey),
	}
}

// NewClientWithHosts creates a new Client from the provided `appID,` `apiKey`,
// and `hosts` used to connect to the Algolia servers.
func NewClientWithHosts(appID, apiKey string, hosts []string) Client {
	return &client{
		transport: NewTransportWithHosts(appID, apiKey, hosts),
	}
}

// SetExtraHeader allows to set custom headers while reaching out to
// Algolia servers.
func (c *client) SetExtraHeader(key, value string) {
	c.transport.setExtraHeader(key, value)
}

// SetTimeout specifies timeouts to use with the HTTP connection.
func (c *client) SetTimeout(connectTimeout, readTimeout int) {
	c.transport.setTimeout(time.Duration(connectTimeout)*time.Millisecond, time.Duration(readTimeout)*time.Millisecond)
}

// ListIndexes returns the list of all indexes belonging to this Algolia
// application.
func (c *client) ListIndexes() (indexes []IndexRes, err error) {
	var res listIndexesRes

	err = c.request(&res, "GET", "/1/indexes", nil, read)
	indexes = res.items
	return
}

// InitIndex returns an Index object targeting `indexName`.
func (c *client) InitIndex(name string) Index {
	return NewIndex(name, c)
}

// ListKeys returns all the API keys available for this Algolia application.
func (c *client) ListKeys() (keys []Key, err error) {
	var res listKeysRes

	err = c.request(&res, "GET", "/1/keys", nil, read)
	keys = res.Keys
	return
}

// MoveIndex renames the index named `source` as `destination`.
func (c *client) MoveIndex(source, destination string) (UpdateTaskRes, error) {
	index := c.InitIndex(source)
	return index.Move(destination)
}

// CopyIndex duplicates the index named `source` as `destination`.
func (c *client) CopyIndex(source, destination string) (UpdateTaskRes, error) {
	index := c.InitIndex(source)
	return index.Copy(destination)
}

// AddKey creates a new API key from the supplied `ACL` and the specified
// optional parameters.
func (c *client) AddKey(ACL []string, params Map) (res AddKeyRes, err error) {
	req := duplicateMap(params)
	req["acl"] = ACL

	if err = checkKey(req); err != nil {
		return
	}

	err = c.request(&res, "POST", "/1/keys/", req, read)
	return
}

// UpdateKey updates the API key named `key` with the supplied
// parameters.
func (c *client) UpdateKey(key string, params Map) (res UpdateKeyRes, err error) {
	if err = checkKey(params); err != nil {
		return
	}

	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "PUT", path, params, write)
	return
}

// GetKey returns the ACL and validity of the API key named `key`.
func (c *client) GetKey(key string) (res Key, err error) {
	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "GET", path, nil, read)
	return
}

// DeleteKey deletes the API key named `key`.
func (c *client) DeleteKey(key string) (res DeleteRes, err error) {
	path := "/1/keys/" + url.QueryEscape(key)
	err = c.request(&res, "DELETE", path, nil, write)
	return
}

// GetLogs retrieves the `length` latest logs, starting at `offset`. Logs can
// be filtered by type via `logType` being either "query", "build" or "error".
func (c *client) GetLogs(params Map) (logs []LogRes, err error) {
	var res getLogsRes

	if err = checkGetLogs(params); err != nil {
		return
	}

	err = c.request(&res, "GET", "/1/logs", params, write)
	logs = res.Logs
	return
}

// GenerateSecuredAPIKey generates a public API key intended to restrict access
// to certain records.
// This new key is built upon the existing key named `apiKey`. Tag filters
// or query parameters used to restrict access to certain records are specified
// via the `public` argument. A single `userToken` may be supplied, in order to
// use rate limited access.
func (c *client) GenerateSecuredAPIKey(apiKey string, params Map) (key string, err error) {
	if err = checkGenerateSecuredAPIKey(params); err != nil {
		return
	}

	message := encodeMap(params)

	h := hmac.New(sha256.New, []byte(apiKey))
	h.Write([]byte(message))
	securedKey := hex.EncodeToString(h.Sum(nil))

	key = base64.StdEncoding.EncodeToString([]byte(securedKey + message))
	return
}

// MultipleQueries performs all the queries specified in `queries` and
// aggregates the results. It accepts two additional arguments: the name of
// the field used to store the index name in the queries, and the strategy used
// to perform the multiple queries.
// The strategy can either be "none" or "stopIfEnoughMatches".
func (c *client) MultipleQueries(queries []IndexedQuery, strategy string) (res []MultipleQueryRes, err error) {
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
	}

	var m multipleQueriesRes
	err = c.request(&m, "POST", "/1/indexes/*/queries?strategy="+strategy, body, search)
	res = m.Results
	return
}

// Batch performs all queries in `queries`. Each query should contain the
// targeted index, as well as the type of operation wanted.
func (c *client) Batch(records []BatchOperationIndexed) (res MultipleBatchRes, err error) {
	// TODO: Use check functions of index.go

	request := map[string][]BatchOperationIndexed{
		"requests": records,
	}

	err = c.request(&res, "POST", "/1/indexes/*/batch", request, write)
	return
}

func (c *client) request(res interface{}, method, path string, body interface{}, typeCall int) error {
	r, err := c.transport.request(method, path, body, typeCall)
	if err != nil {
		return err
	}

	return json.Unmarshal(r, res)
}
