package main


type Client struct {
  transport *Transport
}

func NewClient(appID, apiKey string) *Client {
  client := new(Client)
  client.transport = NewTransport(appID, apiKey)
  return client
}

func (c *Client) listIndexes() (interface{}, error) {
     return c.transport.request("GET", "/1/indexes", "")
}

func (c *Client) initIndex(indexName string) *Index {
  return NewIndex(indexName, c)
}

func (c *Client) listKeys() (interface{}, error) {
  return c.transport.request("GET", "/1/keys", "")
}

func (c *Client) addKey(acl, indexes []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
  body := make(map[string]interface{})
  body["acl"] = acl
  body["maxHitsPerQuery"] = maxHitsPerQuery
  body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
  body["validity"] = validity
  body["indexes"] = indexes
  return c.transport.request("POST", "/1/keys/", body)
}

func (c *Client) getKey(key string) (interface{}, error) {
  return c.transport.request("GET", "/1/keys/" + key, "")
}

func (c *Client) deleteKey(key string) (interface{}, error) {
  return c.transport.request("DELETE", "/1/keys/" + key, "")
}

func (c *Client) getLogs(offset, length int, onlyErrors bool) (interface{}, error) {
  body := make(map[string]interface{})
  body["offset"] = offset
  body["length"] = length
  body["onlyErrors"] = onlyErrors
  return c.transport.request("GET", "/1/logs", body)
}
