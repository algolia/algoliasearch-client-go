package insights

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

// Client provides methods to interact with the Algolia Insights API.
type Client struct {
	appID     string
	transport *transport.Transport
}

// NewClient instantiates a new client able to interact with the Algolia
// Insights API.
func NewClient(appID, apiKey string) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
		},
	)
}

// NewClientWithConfig instantiates a new client able to interact with the
// Algolia Insights API.
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

// User instantiates a new client able to interact with the Algolia Insights API
// where all events will be sent with the userToken field preset according to
// the given userToken.
func (c *Client) User(userToken string) *UserClient {
	return &UserClient{
		UserToken: userToken,
		Client:    *c,
	}
}

// SendEvent sends a new raw event to the Algolia Insights API. This method
// should only be used to send a custom event to the API. To send more common
// and predefined events, please use methods from UserClient instead.
func (c *Client) SendEvent(event Event, opts ...interface{}) (res StatusMessageRes, err error) {
	return c.SendEvents([]Event{event}, opts...)
}

// SendEvents sends new raw events to the Algolia Insights API. This method
// should only be used to send custom events to the API. To send more common and
// predefined events, please use methods from UserClient instead.
func (c *Client) SendEvents(events []Event, opts ...interface{}) (res StatusMessageRes, err error) {
	body := newSendEventsReq(events, opts...)
	err = c.transport.Request(&res, http.MethodPost, "/1/events", body, call.Write, opts...)
	return
}
