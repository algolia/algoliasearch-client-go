package algoliasearch


type Client struct {
  transport *Transport
}

func NewClient(appID, apiKey string) *Client {
  client := new(Client)
  client.transport = NewTransport(appID, apiKey)
  return client
}

func (c *Client) ListIndexes() (interface{}, error) {
     return c.transport.request("GET", "/1/indexes", nil)
}

func (c *Client) InitIndex(indexName string) *Index {
  return NewIndex(indexName, c)
}

func (c *Client) ListKeys() (interface{}, error) {
  return c.transport.request("GET", "/1/keys", nil)
}

func (c *Client) AddKey(acl, indexes []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
  body := make(map[string]interface{})
  body["acl"] = acl
  body["maxHitsPerQuery"] = maxHitsPerQuery
  body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
  body["validity"] = validity
  body["indexes"] = indexes
  return c.transport.request("POST", "/1/keys/", body)
}

func (c *Client) GetKey(key string) (interface{}, error) {
  return c.transport.request("GET", "/1/keys/" + key, nil)
}

func (c *Client) DeleteKey(key string) (interface{}, error) {
  return c.transport.request("DELETE", "/1/keys/" + key, nil)
}

func (c *Client) GetLogs(offset, length int, onlyErrors bool) (interface{}, error) {
  body := make(map[string]interface{})
  body["offset"] = offset
  body["length"] = length
  body["onlyErrors"] = onlyErrors
  return c.transport.request("GET", "/1/logs", body)
}
