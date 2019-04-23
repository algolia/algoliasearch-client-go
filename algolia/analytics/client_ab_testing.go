package analytics

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// AddABTest creates a new AB test.
func (c *Client) AddABTest(abTest ABTest, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("")
	err = c.transport.Request(&res, http.MethodPost, path, abTest, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

// StopABTest stops the AB test with the given ID, which means that this
// specific AB test will not be triggered anymore. However, the AB test and its
// associated results are still accessible.
func (c *Client) StopABTest(id int, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("/%d/stop", id)
	err = c.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

// DeleteABTest effectively deletes the AB test with the given ID, making it
// both disabled and unreachable afterwards i.e. AB tests results are not
// available anymore.
func (c *Client) DeleteABTest(id int, opts ...interface{}) (res ABTestTaskRes, err error) {
	path := c.pathABTesting("/%d", id)
	err = c.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = c.waitTaskSearchClient
	return
}

// GetABTest returns the AB test content and results for the given ID.
func (c *Client) GetABTest(id int, opts ...interface{}) (res ABTestResponse, err error) {
	path := c.pathABTesting("/%d", id)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// GetABTests returns a list of AB tests according to any opt.Offset or
// opt.Limit parameters.
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
