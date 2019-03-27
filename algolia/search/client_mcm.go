package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (c *Client) ListClusters(opts ...interface{}) (res ListClustersRes, err error) {
	err = c.transport.Request(&res, http.MethodGet, "/1/clusters", nil, call.Read, opts...)
	return
}

func (c *Client) ListUserIDs(opts ...interface{}) (res ListUserIDsRes, err error) {
	if page := iopt.ExtractPage(opts...); page != nil {
		opts = opt.InsertExtraURLParam(opts, "page", page.Get())
	}
	if hitsPerPage := iopt.ExtractHitsPerPage(opts...); hitsPerPage != nil {
		opts = opt.InsertExtraURLParam(opts, "hitsPerPage", hitsPerPage.Get())
	}
	err = c.transport.Request(&res, http.MethodGet, "/1/clusters/mapping", nil, call.Read, opts...)
	return
}

func (c *Client) GetUserID(userID string, opts ...interface{}) (res UserID, err error) {
	path := fmt.Sprintf("/1/clusters/mapping/%s", userID)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) AssignUserID(userID, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error) {
	opts = opt.InsertExtraHeader(opts, "X-Algolia-User-ID", userID)
	body := map[string]string{"cluster": clusterName}
	err = c.transport.Request(&res, http.MethodPost, "/1/clusters/mapping", body, call.Write, opts...)
	return
}

func (c *Client) RemoveUserID(userID string, opts ...interface{}) (res RemoveUserIDRes, err error) {
	opts = opt.InsertExtraHeader(opts, "X-Algolia-User-ID", userID)
	err = c.transport.Request(&res, http.MethodDelete, "/1/clusters/mapping", nil, call.Write, opts...)
	return
}

func (c *Client) GetTopUserIDs(opts ...interface{}) (res TopUserIDs, err error) {
	err = c.transport.Request(&res, "GET", "/1/clusters/mapping/top", nil, call.Read, opts...)
	return
}

func (c *Client) SearchUserIDs(query string, opts ...interface{}) (res SearchUserIDRes, err error) {
	req := newSearchUserIDsReq(query, opts...)
	err = c.transport.Request(&res, http.MethodPost, "/1/clusters/mapping/search", req, call.Read, opts...)
	return
}
