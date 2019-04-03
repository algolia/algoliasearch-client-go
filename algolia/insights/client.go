package insights

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Client struct {
	appID     string
	transport *transport.Transport
}

func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

func NewClientWithConfig(config Configuration) *Client {
	var hosts []*transport.StatefulHost

	if config.Hosts == nil {
		hosts = defaultHosts(config.Region)
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsReadWrite))
		}
	}

	return &Client{
		appID: config.AppID,
		transport: transport.New(
			hosts,
			config.Requester,
			config.AppID,
			config.APIKey,
			config.ReadTimeout,
			config.WriteTimeout,
			config.Headers,
		),
	}
}

func (c *Client) User(userToken string) *UserClient {
	return &UserClient{
		UserToken: userToken,
		Client:    *c,
	}
}

func (c *Client) SendEvent(event Event, opts ...interface{}) (res StatusMessageRes, err error) {
	return c.SendEvents([]Event{event}, opts...)
}

func (c *Client) SendEvents(events []Event, opts ...interface{}) (res StatusMessageRes, err error) {
	err = c.transport.Request(&res, http.MethodPost, "/1/events", events, call.Write, opts...)
	return
}
