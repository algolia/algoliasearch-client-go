package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
)

func (c *Client) SetPersonalizationStrategy(strategy Strategy, opts ...interface{}) (res SetPersonalizationStrategyRes, err error) {
	path := c.path("/recommendation/personalization/strategy")
	err = c.transport.Request(&res, http.MethodPost, path, strategy, call.Write, opts...)
	return
}

func (c *Client) GetPersonalizationStrategy(opts ...interface{}) (res GetPersonalizationStrategyRes, err error) {
	path := c.path("/recommendation/personalization/strategy")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}
