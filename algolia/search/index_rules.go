package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (i *Index) GetRule(objectID string, opts ...interface{}) (rule Rule, err error) {
	path := i.path("/rules/%s", objectID)
	err = i.transport.Request(&rule, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (i *Index) SaveRule(rule Rule, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/rules/%s", rule.ObjectID)
	err = i.transport.Request(&res, http.MethodPut, path, rule, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) SaveRules(rules []Rule, opts ...interface{}) (res UpdateTaskRes, err error) {
	if clearExistingRules := iopt.ExtractClearExistingRules(opts...); clearExistingRules != nil {
		opts = append(opts, opt.ExtraURLParams(
			map[string]string{"clearExistingRules": fmt.Sprintf("%t", clearExistingRules.Get())},
		))
	}
	path := i.path("/rules/batch")
	err = i.transport.Request(&res, http.MethodPost, path, rules, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) ClearRules(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/rules/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) DeleteRule(objectID string, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/rules/%s", objectID)
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) SearchRules(query string, opts ...interface{}) (res SearchRulesRes, err error) {
	body := newSearchRulesParams(query, opts...)
	path := i.path("/rules/search")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

func (i *Index) ReplaceAllRules(rules []Rule, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.ClearExistingRules(true))
	return i.SaveRules(rules, opts...)
}
