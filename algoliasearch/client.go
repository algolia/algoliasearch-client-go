package algoliasearch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
)

// Client is a representation of an Algolia application. Once initialized it
// allows manipulations over the indexes of the application as well as
// network related parameters.
type Client struct {
	transport *Transport
}

// NewClient creates a new Client from the provided `appID` and `apiKey`. The
// default hosts are used for the transport layer.
func NewClient(appID, apiKey string) Client {
	return Client{
		transport: NewTransport(appID, apiKey),
	}
}

// NewClientWithHosts creates a new Client from the provided `appID,` `apiKey`,
// and `hosts` used to connect to the Algolia servers.
func NewClientWithHosts(appID, apiKey string, hosts []string) Client {
	return Client{
		transport: NewTransportWithHosts(appID, apiKey, hosts),
	}
}

// SetExtraHeader allows to set custom headers while reaching out to
// Algolia servers.
func (c *Client) SetExtraHeader(key string, value string) {
	c.transport.setExtraHeader(key, value)
}

// SetTimeout specifies timeouts to use with the HTTP connection.
func (c *Client) SetTimeout(connectTimeout int, readTimeout int) {
	c.transport.setTimeout(time.Duration(connectTimeout)*time.Millisecond, time.Duration(readTimeout)*time.Millisecond)
}

// ListIndexes returns the list of all indexes belonging to this Algolia
// application.
func (c *Client) ListIndexes() (ListIndexesRes, error) {
	res, err := c.transport.request("GET", "/1/indexes", nil, read)
	if err != nil {
		return ListIndexesRes{}, err
	}

	indexes := ListIndexesRes{}
	err = json.Unmarshal(res, &indexes)

	return indexes, err
}

// InitIndex returns an Index object targeting `indexName`.
func (c *Client) InitIndex(indexName string) Index {
	return *NewIndex(indexName, c)
}

// ListKeys returns all the API keys available for this Algolia application.
func (c *Client) ListKeys() (ListKeysRes, error) {
	keys := ListKeysRes{}
	err := c.request(&keys, "GET", "/1/keys", nil, read)
	return keys, err
}

// MoveIndex renames the index named `source` as `destination`.
func (c *Client) MoveIndex(source string, destination string) (interface{}, error) {
	index := c.InitIndex(source)
	return index.Move(destination)
}

// CopyIndex duplicates the index named `source` as `destination`.
func (c *Client) CopyIndex(source string, destination string) (interface{}, error) {
	index := c.InitIndex(source)
	return index.Copy(destination)
}

// AddKey creates a new API key from the supplied `ACL` and the specified
// optional parameters.
func (c *Client) AddKey(ACL []string, params map[string]interface{}) (AddKeyRes, error) {
	req := make(map[string]interface{})
	for k, v := range params {
		req[k] = v
	}
	req["acl"] = ACL

	if err := checkKeyReq(req); err != nil {
		return AddKeyRes{}, err
	}

	add := AddKeyRes{}
	err := c.request(&add, "POST", "/1/keys/", req, read)
	return add, err
}

// UpdateKeyWithParam updates the API key named `key` with the supplied
// parameters.
func (c *Client) UpdateKeyWithParam(key string, params map[string]interface{}) (UpdateKeyRes, error) {
	if err := checkKeyReq(params); err != nil {
		return UpdateKeyRes{}, err
	}

	update := UpdateKeyRes{}
	err := c.request(&update, "PUT", "/1/keys/"+key, params, write)
	return update, err
}

// GetKey returns the ACL and validity of the API key named `key`.
func (c *Client) GetKey(key string) (GetKeyRes, error) {
	get := GetKeyRes{}
	err := c.request(&get, "GET", "/1/keys/"+key, nil, read)
	return get, err
}

// DeleteKey deletes the API key named `key`.
func (c *Client) DeleteKey(key string) (DeleteKeyRes, error) {
	delete := DeleteKeyRes{}
	err := c.request(&delete, "DELETE", "/1/keys/"+key, nil, write)
	return delete, err
}

// GetLogs retrieves the `length` latest logs, starting at `offset`. Logs can
// be filtered by type via `logType` being either "query", "build" or "error".
func (c *Client) GetLogs(params map[string]interface{}) (GetLogsRes, error) {
	if err := checkGetLogs(params); err != nil {
		return GetLogsRes{}, err
	}

	logs := GetLogsRes{}
	err := c.request(&logs, "GET", "/1/logs", params, write)
	return logs, err
}

// GenerateSecuredApiKey generates a public API key intended to restrict access
// to certain records.
// This new key is built upon the existing key named `apiKey`. Tag filters
// or query parameters used to restrict access to certain records are specified
// via the `public` argument. A single `userToken` may be supplied, in order to
// use rate limited access.
func (c *Client) GenerateSecuredApiKey(apiKey string, params map[string]interface{}) (string, error) {
	if err := checkGenerateSecuredAPIKey(params); err != nil {
		return "", nil
	}

	req := make(map[string]interface{})
	for k, v := range params {
		req[k] = v
	}
	req["tagFilters"] = strings.Join(params["tagFilters"].([]string), ",")
	message := c.transport.EncodeParams(req)

	h := hmac.New(sha256.New, []byte(apiKey))
	h.Write([]byte(message))
	securedKey := hex.EncodeToString(h.Sum(nil))

	return base64.StdEncoding.EncodeToString([]byte(securedKey + message)), nil
}

// MultipleQueries performs all the queries specified in `queries` and
// aggregates the results. It accepts two additional arguments: the name of
// the field used to store the index name in the queries, and the strategy used
// to perform the multiple queries.
// The strategy can either be "none" or "stopIfEnoughMatches".
func (c *Client) MultipleQueries(queries []map[string]interface{}, indexField, strategy string) (MultipleQueriesRes, error) {
	if indexField == "" {
		indexField = "indexName"
	}

	if strategy == "" {
		strategy = "none"
	} else if strategy != "none" && strategy != "stopIfEnoughMatches" {
		return MultipleQueriesRes{}, invalidValue(strategy, "strategy")
	}

	for _, query := range queries {
		if err := checkParams(query, checkQueryMap, indexField); err != nil {
			return MultipleQueriesRes{}, err
		}
	}

	requests := make([]map[string]string, len(queries))
	for i, q := range queries {
		requests[i] = map[string]string{
			"indexName": q[indexField].(string),
		}

		requests[i]["params"] = c.transport.EncodeParams(transformQuery(q, indexField))
	}

	body := map[string]interface{}{
		"requests": requests,
	}

	res := MultipleQueriesRes{}
	err := c.request(&res, "POST", "/1/indexes/*/queries?strategy="+strategy, body, search)
	return res, err
}

// Batch performs all queries in `queries`. Each query should contain the
// targeted index, as well as the type of operation wanted.
func (c *Client) Batch(records []BatchRecord) (CustomBatchRes, error) {
	for _, record := range records {
		if err := checkBatchAction(record.Action); err != nil {
			return CustomBatchRes{}, err
		}
	}

	request := map[string][]BatchRecord{
		"requests": records,
	}

	batch := CustomBatchRes{}
	err := c.request(&batch, "POST", "/1/indexes/*/batch", request, write)
	return batch, err
}

func (c *Client) request(res interface{}, method, path string, body interface{}, typeCall int) error {
	r, err := c.transport.request(method, path, body, typeCall)

	if err != nil {
		return err
	}

	return json.Unmarshal(r, res)
}
