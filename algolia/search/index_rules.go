package search

import (
	"net/http"
	"net/url"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// GetRule retrieves the rule identified by the given objectID.
func (i *Index) GetRule(objectID string, opts ...interface{}) (rule Rule, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		return
	}

	path := i.path("/rules/%s", url.QueryEscape(objectID))
	err = i.transport.Request(&rule, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// SaveRule saves the given rule.
func (i *Index) SaveRule(rule Rule, opts ...interface{}) (res UpdateTaskRes, err error) {
	if rule.ObjectID == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/rules/%s", url.QueryEscape(rule.ObjectID))
	err = i.transport.Request(&res, http.MethodPut, path, rule, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// SaveRule saves the given rules.
//
// Unlike SaveObjects, this method does not batch the given rules i.e. all rules
// are sent in a single call.
func (i *Index) SaveRules(rules []Rule, opts ...interface{}) (res UpdateTaskRes, err error) {
	if clearExistingRules := iopt.ExtractClearExistingRules(opts...); clearExistingRules != nil {
		opts = opt.InsertExtraURLParam(opts, "clearExistingRules", clearExistingRules.Get())
	}
	path := i.path("/rules/batch")
	err = i.transport.Request(&res, http.MethodPost, path, rules, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// ClearRules removes all the rules from the index.
func (i *Index) ClearRules(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/rules/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// DeleteRule removes the rule identified by the given objectID.
func (i *Index) DeleteRule(objectID string, opts ...interface{}) (res UpdateTaskRes, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/rules/%s", url.QueryEscape(objectID))
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// SearchRules search for rules according to the given query string and any rule
// parameter, as documented here:
// https://www.algolia.com/doc/api-reference/api-methods/search-rules/
func (i *Index) SearchRules(query string, opts ...interface{}) (res SearchRulesRes, err error) {
	body := newSearchRulesParams(query, opts...)
	path := i.path("/rules/search")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// ReplaceAllRules replaces any existing rules with the given ones.
func (i *Index) ReplaceAllRules(rules []Rule, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.ClearExistingRules(true))
	return i.SaveRules(rules, opts...)
}

// BrowseRules returns an iterator which will retrieve rules one by one from the
// index.
func (i *Index) BrowseRules(opts ...interface{}) (*RuleIterator, error) {
	browser := func(page int) (SearchRulesRes, error) {
		opts = opt.InsertOrReplaceOption(opts, opt.Page(page))
		return i.SearchRules("", opts...)
	}
	return newRuleIterator(browser)
}
