package algoliasearch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"reflect"
	"time"
	"strings"
	"encoding/base64"
)

type Client struct {
	Transport *Transport
}

func NewClient(appID, apiKey string) *Client {
	client := new(Client)
	client.Transport = NewTransport(appID, apiKey)
	return client
}

func NewClientWithHosts(appID, apiKey string, hosts []string) *Client {
	client := new(Client)
	client.Transport = NewTransportWithHosts(appID, apiKey, hosts)
	return client
}

func (c *Client) SetExtraHeader(key string, value string) {
	c.Transport.setExtraHeader(key, value)
}

func (c *Client) SetTimeout(connectTimeout int, readTimeout int) {
	c.Transport.setTimeout(time.Duration(connectTimeout)*time.Millisecond, time.Duration(readTimeout)*time.Millisecond)
}

func (c *Client) ListIndexes() (interface{}, error) {
	return c.Transport.request("GET", "/1/indexes", nil, read)
}

func (c *Client) InitIndex(indexName string) *Index {
	return NewIndex(indexName, c)
}

func (c *Client) ListKeys() (interface{}, error) {
	return c.Transport.request("GET", "/1/keys", nil, read)
}

func (c *Client) AddKey(acl, indexes []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
	body := make(map[string]interface{})
	body["acl"] = acl
	body["maxHitsPerQuery"] = maxHitsPerQuery
	body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
	body["validity"] = validity
	body["indexes"] = indexes
	return c.AddKeyWithParam(body)
}

func (c *Client) AddKeyWithParam(params interface{}) (interface{}, error) {
	return c.Transport.request("POST", "/1/keys/", params, read)
}

func (c *Client) UpdateKey(key string, acl, indexes []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
	body := make(map[string]interface{})
	body["acl"] = acl
	body["maxHitsPerQuery"] = maxHitsPerQuery
	body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
	body["validity"] = validity
	body["indexes"] = indexes
	return c.UpdateKeyWithParam(key, body)
}

func (c *Client) UpdateKeyWithParam(key string, params interface{}) (interface{}, error) {
	return c.Transport.request("PUT", "/1/keys/"+key, params, write)
}

func (c *Client) GetKey(key string) (interface{}, error) {
	return c.Transport.request("GET", "/1/keys/"+key, nil, read)
}

func (c *Client) DeleteKey(key string) (interface{}, error) {
	return c.Transport.request("DELETE", "/1/keys/"+key, nil, write)
}

func (c *Client) GetLogs(offset, length int, logType string) (interface{}, error) {
	body := make(map[string]interface{})
	body["offset"] = offset
	body["length"] = length
	body["type"] = logType
	return c.Transport.request("GET", "/1/logs", body, write)
}

func (c *Client) GenerateSecuredApiKey(apiKey string, public interface{}, userToken ...string) (string, error) {
	if len(userToken) > 1 {
		return "", errors.New("Too many parameters")
	}

	var userTokenStr string
	var message string
	if len(userToken) == 1 {
		userTokenStr = userToken[0]
	} else {
		userTokenStr = ""
	}
	if reflect.TypeOf(public).Name() != "string" { // QueryParameters
		if len(userTokenStr) != 0 {
			public.(map[string]interface{})["userToken"] = userTokenStr
		}
		message = c.Transport.EncodeParams(public)
	} else if strings.Contains(public.(string), "=") && len(userTokenStr) != 0 { // Url encoded query parameters
		message = public.(string) + "userToken=" + c.Transport.urlEncode(userTokenStr)
	} else { // TagFilters
		queryParameters := make(map[string]interface{})
		queryParameters["tagFilters"] = public
		if len(userTokenStr) != 0 {
			queryParameters["userToken"] = userTokenStr
		}
		message = c.Transport.EncodeParams(queryParameters)
	}

	key := []byte(apiKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	securedKey := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(securedKey + message)), nil
}

func (c *Client) EncodeParams(body interface{}) string {
	return c.Transport.EncodeParams(body)
}

func (c *Client) MultipleQueries(queries []interface{}, optionals ...string) (interface{}, error) {
	if len(optionals) > 2 {
		return "", errors.New("Too many parametters")
	}
	var nameKey string
	if len(optionals) >= 1 {
		nameKey = optionals[0]
	} else {
		nameKey = "indexName"
	}
	var strategy string = "none"
	if len(optionals) == 2 {
		strategy = optionals[1]
	}
	requests := make([]map[string]interface{}, len(queries))
	for i := range queries {
		requests[i] = make(map[string]interface{})
		requests[i]["indexName"] = queries[i].(map[string]interface{})[nameKey].(string)
		delete(queries[i].(map[string]interface{}), nameKey)
		requests[i]["params"] = c.Transport.EncodeParams(queries[i])
	}
	body := make(map[string]interface{})
	body["requests"] = requests
	return c.Transport.request("POST", "/1/indexes/*/queries?strategy="+strategy, body, search)
}

func (c *Client) CustomBatch(queries interface{}) (interface{}, error) {
	request := make(map[string]interface{})
	request["requests"] = queries
	return c.Transport.request("POST", "/1/indexes/*/batch", request, write)
}
