package algoliasearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	appID     string
	transport *Transport
}

// NewClient instantiates a new `Client` from the provided `appID` and
// `apiKey`. Default hosts are used for the transport layer.
func NewClient(appID, apiKey string) Client {
	return &client{
		appID:     appID,
		transport: NewTransport(appID, apiKey),
	}
}

// NewClientWithHosts instantiates a new `Client` from the provided `appID` and
// `apiKey`. The transport layers' hosts are initialized with the given
// `hosts`.
func NewClientWithHosts(appID, apiKey string, hosts []string) Client {
	return &client{
		appID:     appID,
		transport: NewTransportWithHosts(appID, apiKey, hosts),
	}
}

func (c *client) GetAppID() string {
	return c.appID
}

func (c *client) SetExtraHeader(key, value string) {
	c.transport.setExtraHeader(key, value)
}

func (c *client) SetTimeout(connectTimeout, readTimeout int) {
	// As SetTimeout is about to be deprecated, we simply forward the
	// readTimeout but ignores the connectTimeout that is not longer
	// configurable.
	c.SetReadTimeout(time.Duration(readTimeout) * time.Second)
}
func (c *client) SetReadTimeout(t time.Duration)      { c.transport.setTimeouts(t, -1, -1, -1) }
func (c *client) SetWriteTimeout(t time.Duration)     { c.transport.setTimeouts(-1, t, -1, -1) }
func (c *client) SetAnalyticsTimeout(t time.Duration) { c.transport.setTimeouts(-1, -1, t, -1) }
func (c *client) SetInsightsTimeout(t time.Duration)  { c.transport.setTimeouts(-1, -1, -1, t) }

func (c *client) SetMaxIdleConnsPerHosts(maxIdleConnsPerHost int) {
	c.transport.setMaxIdleConnsPerHost(maxIdleConnsPerHost)
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

func (c *client) InitAnalytics() Analytics {
	return NewAnalytics(c)
}

func (c *client) InitInsights() Insights {
	return NewInsights(c)
}

func (c *client) ListKeys() (keys []Key, err error) {
	return c.ListAPIKeys()
}

func (c *client) ListKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error) {
	return c.ListAPIKeysWithRequestOptions(opts)
}

func (c *client) ListAPIKeys() (keys []Key, err error) {
	return c.ListAPIKeysWithRequestOptions(nil)
}

func (c *client) ListAPIKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error) {
	var res listAPIKeysRes
	err = c.request(&res, "GET", "/1/keys", nil, read, opts)
	keys = res.Keys
	return
}

func (c *client) MoveIndex(source, destination string) (UpdateTaskRes, error) {
	return c.MoveIndexWithRequestOptions(source, destination, nil)
}

func (c *client) MoveIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.operation(source, destination, "move", nil, opts)
}

func (c *client) CopyIndex(source, destination string) (UpdateTaskRes, error) {
	return c.CopyIndexWithRequestOptions(source, destination, nil)
}

func (c *client) CopyIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.ScopedCopyIndexWithRequestOptions(source, destination, nil, opts)
}

func (c *client) ScopedCopyIndex(source, destination string, scopes []string) (UpdateTaskRes, error) {
	return c.ScopedCopyIndexWithRequestOptions(source, destination, scopes, nil)
}

func (c *client) ScopedCopyIndexWithRequestOptions(source, destination string, scopes []string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.operation(source, destination, "copy", scopes, opts)
}

func (c *client) operation(src, dst, op string, scopes []string, opts *RequestOptions) (res UpdateTaskRes, err error) {
	if err = checkScopes(scopes); err != nil {
		return
	}

	o := IndexOperation{
		Destination: dst,
		Operation:   op,
		Scopes:      scopes,
	}

	path := "/1/indexes/" + url.QueryEscape(src) + "/operation"
	err = c.request(&res, "POST", path, o, write, opts)
	return
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

	err = c.request(&res, "POST", "/1/keys/", req, write, opts)
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

func (c *client) ListClusters() (res []Cluster, err error) {
	return c.ListClustersWithRequestOptions(nil)
}

func (c *client) ListClustersWithRequestOptions(opts *RequestOptions) (res []Cluster, err error) {
	var rawRes map[string][]Cluster
	var ok bool
	err = c.request(&rawRes, "GET", "/1/clusters", nil, read, opts)
	if res, ok = rawRes["clusters"]; !ok {
		res = nil
		err = errors.New("missing field `clusters` in JSON response")
	}
	return
}

func (c *client) ListUserIDs(page int, hitsPerPage int) (res ListUserIDsRes, err error) {
	return c.ListUserIDsWithRequestOptions(page, hitsPerPage, nil)
}

func (c *client) ListUserIDsWithRequestOptions(page int, hitsPerPage int, opts *RequestOptions) (res ListUserIDsRes, err error) {
	params := Map{
		"page":        page,
		"hitsPerPage": hitsPerPage,
	}
	err = c.request(&res, "GET", "/1/clusters/mapping?"+encodeMap(params), nil, read, opts)
	return
}

func (c *client) GetUserID(userID string) (res UserID, err error) {
	return c.GetUserIDWithRequestOptions(userID, nil)
}

func (c *client) GetUserIDWithRequestOptions(userID string, opts *RequestOptions) (res UserID, err error) {
	err = c.request(&res, "GET", "/1/clusters/mapping/"+url.QueryEscape(userID), nil, read, opts)
	return
}

func (c *client) AssignUserID(userID string, clusterName string) (res AssignUserIDRes, err error) {
	return c.AssignUserIDWithRequestOptions(userID, clusterName, nil)
}

func (c *client) AssignUserIDWithRequestOptions(userID string, clusterName string, opts *RequestOptions) (res AssignUserIDRes, err error) {
	if opts == nil {
		opts = &RequestOptions{
			ExtraHeaders: make(map[string]string),
		}
	}
	opts.ExtraHeaders["X-Algolia-User-ID"] = userID
	body := map[string]string{"cluster": clusterName}
	err = c.request(&res, "POST", "/1/clusters/mapping", body, write, opts)
	return
}

func (c *client) RemoveUserID(userID string) (res RemoveUserIDRes, err error) {
	return c.RemoveUserIDWithRequestOptions(userID, nil)
}

func (c *client) RemoveUserIDWithRequestOptions(userID string, opts *RequestOptions) (res RemoveUserIDRes, err error) {
	if opts == nil {
		opts = &RequestOptions{
			ExtraHeaders: make(map[string]string),
		}
	}
	opts.ExtraHeaders["X-Algolia-User-ID"] = userID

	err = c.request(&res, "DELETE", "/1/clusters/mapping", nil, write, opts)
	return
}

func (c *client) GetTopUserIDs() (res TopUserIDs, err error) {
	return c.GetTopUserIDsWithRequestOptions(nil)
}

func (c *client) GetTopUserIDsWithRequestOptions(opts *RequestOptions) (res TopUserIDs, err error) {
	err = c.request(&res, "GET", "/1/clusters/mapping/top", nil, read, opts)
	return
}

func (c *client) SearchUserIDs(query string, params Map) (res SearchUserIDRes, err error) {
	return c.SearchUserIDsWithRequestOptions(query, params, nil)
}

func (c *client) SearchUserIDsWithRequestOptions(query string, params Map, opts *RequestOptions) (res SearchUserIDRes, err error) {
	params["query"] = query
	err = c.request(&res, "POST", "/1/clusters/mapping/search", params, read, opts)
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

func (c *client) WaitTask(indexName string, taskID int) error {
	return c.WaitTaskWithRequestOptions(indexName, taskID, nil)
}

func (c *client) WaitTaskWithRequestOptions(indexName string, taskID int, opts *RequestOptions) error {
	var maxDuration = time.Second

	for {
		res, err := c.GetStatusWithRequestOptions(indexName,
			taskID, opts)
		if err != nil {
			return err
		}

		if res.Status == "published" {
			return nil
		}

		sleepDuration := randDuration(maxDuration)
		time.Sleep(sleepDuration)

		// Increase the upper boundary used to generate the sleep duration
		if maxDuration < 10*time.Minute {
			maxDuration *= 2
			if maxDuration > 10*time.Minute {
				maxDuration = 10 * time.Minute
			}
		}
	}
}

func (c *client) GetStatus(indexName string, taskID int) (res TaskStatusRes, err error) {
	return c.GetStatusWithRequestOptions(indexName, taskID, nil)
}

func (c *client) GetStatusWithRequestOptions(indexName string, taskID int, opts *RequestOptions) (res TaskStatusRes, err error) {
	path := fmt.Sprintf("/1/indexes/%s/task/%d", url.QueryEscape(indexName), taskID)
	err = c.request(&res, "GET", path, nil, read, opts)
	return
}

func (c *client) CopySettings(source, destination string) (UpdateTaskRes, error) {
	return c.CopySettingsWithRequestOptions(source, destination, nil)
}

func (c *client) CopySettingsWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.ScopedCopyIndexWithRequestOptions(source, destination, []string{"settings"}, opts)
}

func (c *client) CopySynonyms(source, destination string) (UpdateTaskRes, error) {
	return c.CopySynonymsWithRequestOptions(source, destination, nil)
}

func (c *client) CopySynonymsWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.ScopedCopyIndexWithRequestOptions(source, destination, []string{"synonyms"}, opts)
}

func (c *client) CopyRules(source, destination string) (UpdateTaskRes, error) {
	return c.CopyRulesWithRequestOptions(source, destination, nil)
}

func (c *client) CopyRulesWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error) {
	return c.ScopedCopyIndexWithRequestOptions(source, destination, []string{"rules"}, opts)
}

func (c *client) SetPersonalizationStrategy(strategy Strategy) (SetStrategyRes, error) {
	return c.SetPersonalizationStrategyWithRequestOptions(strategy, nil)
}

func (c *client) SetPersonalizationStrategyWithRequestOptions(strategy Strategy, opts *RequestOptions) (res SetStrategyRes, err error) {
	path := "/1/recommendation/personalization/strategy"
	err = c.request(&res, "POST", path, strategy, write, opts)
	return
}

func (c *client) GetPersonalizationStrategy() (Strategy, error) {
	return c.GetPersonalizationStrategyWithRequestOptions(nil)
}

func (c *client) GetPersonalizationStrategyWithRequestOptions(opts *RequestOptions) (strategy Strategy, err error) {
	path := "/1/recommendation/personalization/strategy"
	err = c.request(&strategy, "GET", path, nil, read, opts)
	return
}

func (c *client) request(res interface{}, method, path string, body interface{}, typeCall int, opts *RequestOptions) error {
	r, err := c.transport.request(method, path, body, typeCall, opts)
	if err != nil {
		return err
	}

	return json.Unmarshal(r, res)
}
