package main


type Client struct {
  transport *Transport
}

func NewClient(appID, apiKey string) *Client {
  client := new(Client)
  client.transport = NewTransport(appID, apiKey)
  return client
}

func (c *Client) listIndexes() (string) {
     return c.transport.request("GET", "/1/indexes", "")
}

func (c *Client) initIndex(indexName string) *Index {
  return NewIndex(indexName, c)
}

func (c *Client) listKeys() string {
  return c.transport.request("GET", "/1/keys", "")
}

func (c *Client) getKey(key string) string {
  return c.transport.request("GET", "/1/keys/" + key, "")
}

func (c *Client) deleteKey(key string) string {
  return c.transport.request("DELETE", "/1/keys/" + key, "")
}
