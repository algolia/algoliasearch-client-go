package recommendation

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

func (c *Client) GetPersonalizationStrategy(opts ...interface{}) (strategy Strategy, err error) {
	path := c.pathPersonalization("")
	err = c.transport.Request(&strategy, http.MethodGet, path, nil, call.Read)
	return
}

func (c *Client) SetPersonalizationStrategy(strategy Strategy, opts ...interface{}) (res SetPersonalizationStrategyRes, err error) {
	path := c.pathPersonalization("")
	err = c.transport.Request(&strategy, http.MethodPost, path, strategy, call.Write)
	return
}

func (c *Client) pathPersonalization(format string, a ...interface{}) string {
	return "/1/strategies/personalization" + fmt.Sprintf(format, a...)
}
