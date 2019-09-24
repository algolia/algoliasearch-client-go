package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// GetSettings retrieves the settings of the index.
func (i *Index) GetSettings(opts ...interface{}) (settings Settings, err error) {
	if advanced := iopt.ExtractAdvanced(opts...); advanced != nil {
		opts = opt.InsertExtraURLParam(opts, "advanced", advanced.Get())
	}
	opts = opt.InsertExtraURLParam(opts, "getVersion", "2")
	path := i.path("/settings")
	err = i.transport.Request(&settings, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// SetSettings applies all the non-nil Settings field to the settings
// configuration of the index.
func (i *Index) SetSettings(settings Settings, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/settings")
	err = i.transport.Request(&res, http.MethodPut, path, settings, call.Write, opts...)
	res.wait = i.WaitTask
	return
}
