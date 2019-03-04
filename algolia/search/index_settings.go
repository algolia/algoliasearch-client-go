package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (i *Index) GetSettings(opts ...interface{}) (settings Settings, err error) {
	opts = append(opts, opt.ExtraURLParams(map[string]string{"getVersion": "2"}))
	path := i.path("/settings")
	err = i.transport.Request(&settings, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (i *Index) SetSettings(settings Settings, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/settings")
	err = i.transport.Request(&res, http.MethodPut, path, settings, call.Write, opts...)
	res.wait = i.waitTask
	return
}
