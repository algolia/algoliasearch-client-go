package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ListClusters list all the clusters managed by MCM.
func (c *Client) ListClusters(opts ...interface{}) (res ListClustersRes, err error) {
	err = c.transport.Request(&res, http.MethodGet, "/1/clusters", nil, call.Read, opts...)
	return
}

// ListUserIDs list all the userIDs managed by MCM.
//
// Pagination can be implemented using the opt.Page and opt.HitsPerPage option
// parameters.
func (c *Client) ListUserIDs(opts ...interface{}) (res ListUserIDsRes, err error) {
	opts = opt.InsertExtraURLParam(opts, "page", iopt.ExtractPage(opts...).Get())
	opts = opt.InsertExtraURLParam(opts, "hitsPerPage", iopt.ExtractHitsPerPage(opts...).Get())
	err = c.transport.Request(&res, http.MethodGet, "/1/clusters/mapping", nil, call.Read, opts...)
	return
}

// GetUserID retrieves the user, managed by MCM, according to the given userID.
func (c *Client) GetUserID(userID string, opts ...interface{}) (res UserID, err error) {
	path := fmt.Sprintf("/1/clusters/mapping/%s", userID)
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// AssignUserID assign the given userID to the given cluster.
func (c *Client) AssignUserID(userID, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error) {
	opts = opt.InsertExtraHeader(opts, "X-Algolia-User-ID", userID)
	body := map[string]string{"cluster": clusterName}
	err = c.transport.Request(&res, http.MethodPost, "/1/clusters/mapping", body, call.Write, opts...)
	return
}

func (c *Client) AssignUserIDs(userIDs []string, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error) {
	body := map[string]interface{}{"cluster": clusterName, "users": userIDs}
	err = c.transport.Request(&res, http.MethodPost, "/1/clusters/mapping/batch", body, call.Write, opts...)
	return
}

// RemoveUserID deletes the given userID managed by MCM.
func (c *Client) RemoveUserID(userID string, opts ...interface{}) (res RemoveUserIDRes, err error) {
	opts = opt.InsertExtraHeader(opts, "X-Algolia-User-ID", userID)
	err = c.transport.Request(&res, http.MethodDelete, "/1/clusters/mapping", nil, call.Write, opts...)
	return
}

// GetTopUserIDs retrieves the top list of userIDs managed by MCM.
func (c *Client) GetTopUserIDs(opts ...interface{}) (res TopUserIDs, err error) {
	err = c.transport.Request(&res, "GET", "/1/clusters/mapping/top", nil, call.Read, opts...)
	return
}

func (c *Client) SearchUserIDs(query string, opts ...interface{}) (res SearchUserIDRes, err error) {
	req := newSearchUserIDsReq(query, opts...)
	err = c.transport.Request(&res, http.MethodPost, "/1/clusters/mapping/search", req, call.Read, opts...)
	return
}

// HasPendingMappings answers with a boolean indicating whether or not at least
// one mapping is currently pending.
//
// If opt.RetrieveMappings(true) is passed,
// the method also retrieves the list of userIDs for which an operation is
// currently pending, grouped by cluster, if any.
func (c *Client) HasPendingMappings(opts ...interface{}) (res HasPendingMappingsRes, err error) {
	if retrieveMappings := iopt.ExtractRetrieveMappings(opts...); retrieveMappings != nil {
		opts = opt.InsertExtraURLParam(opts, "getClusters", retrieveMappings.Get())
	}
	err = c.transport.Request(&res, http.MethodGet, "/1/clusters/mapping/pending", nil, call.Read, opts...)
	return
}
