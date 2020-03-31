package monitoring

import (
	"net/http"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

// GetCurrentApiStatus - This method gets the current status of all clusters/replicas.
func (c *Client) GetCurrentApiStatus() (status *StatusResp, err error) {
	err = c.transport.Request(&status, http.MethodGet, c.path("/status"), nil, call.Read)
	return
}

// GetCurrentApiStatus - This method gets the current status of all clusters/replicas.
func (c *Client) GetCurrentServerStatus(servers ...string) (status *StatusResp, err error) {
	err = c.transport.Request(&status, http.MethodGet, c.path("/status/%s", strings.Join(servers, ",")), nil, call.Read)
	return
}
