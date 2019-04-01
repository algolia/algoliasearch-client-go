package analytics

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (c *Client) AddABTest(abTest ABTest, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("")
	err = c.transport.Request(&res, http.MethodPost, path, abTest, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

func (c *Client) StopABTest(id int, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("/%d/stop", id)
	err = c.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

func (c *Client) DeleteABTest(id int, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("/%d", id)
	err = c.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

func (c *Client) GetABTest(id int, opts ...interface{}) (res ABTestResponse, err error) {
	path := c.pathABTesting("/%d", id)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) GetABTests(opts ...interface{}) (res GetABTestsRes, err error) {
	if offset := iopt.ExtractOffset(opts...); offset != nil {
		opts = opt.InsertExtraURLParam(opts, "offset", offset.Get())
	}
	if limit := iopt.ExtractLimit(opts...); limit != nil {
		opts = opt.InsertExtraURLParam(opts, "limit", limit.Get())
	}
	path := c.pathABTesting("")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) pathABTesting(format string, a ...interface{}) string {
	return "/2/abtests" + fmt.Sprintf(format, a...)
}
