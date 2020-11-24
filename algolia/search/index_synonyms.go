package search

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// GetSynonym retrieves the synonym identified by the given objectID.
func (i *Index) GetSynonym(objectID string, opts ...interface{}) (synonym Synonym, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		return
	}

	var syn rawSynonym
	path := i.path("/synonyms/%s", url.QueryEscape(objectID))
	err = i.transport.Request(&syn, http.MethodGet, path, nil, call.Read, opts...)
	if err == nil {
		synonym = syn.impl
	}
	return
}

// SaveSynonym saves the given synonym.
func (i *Index) SaveSynonym(synonym Synonym, opts ...interface{}) (res UpdateTaskRes, err error) {
	if synonym.ObjectID() == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/synonyms/%s", url.QueryEscape(synonym.ObjectID()))
	err = i.transport.Request(&res, http.MethodPut, path, synonym, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// SaveSynonym saves the given synonyms.
//
// Unlike SaveObjects, this method does not batch the given synonyms i.e. all synonyms
// are sent in a single call.
func (i *Index) SaveSynonyms(synonyms []Synonym, opts ...interface{}) (res UpdateTaskRes, err error) {

	if clearExistingSynonyms := iopt.ExtractClearExistingSynonyms(opts...); clearExistingSynonyms != nil {
		opts = opt.InsertExtraURLParam(opts, "replaceExistingSynonyms", clearExistingSynonyms.Get())
	} else if replaceExistingSynonyms := iopt.ExtractReplaceExistingSynonyms(opts...); replaceExistingSynonyms != nil {
		// if clearExistingSynonyms parameter is missing, attempt to parse legacy replaceExistingSynonyms option
		opts = opt.InsertExtraURLParam(opts, "replaceExistingSynonyms", replaceExistingSynonyms.Get())
	}

	path := i.path("/synonyms/batch")
	err = i.transport.Request(&res, http.MethodPost, path, synonyms, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// ClearSynonyms removes all the synonyms from the index.
func (i *Index) ClearSynonyms(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/synonyms/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// DeleteSynonym removes the synonym identified by the given objectID.
func (i *Index) DeleteSynonym(objectID string, opts ...interface{}) (res DeleteTaskRes, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/synonyms/%s", url.QueryEscape(objectID))
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// SearchSynonyms search for synonyms according to the given query string and any synonym
// parameter, as documented here:
// https://www.algolia.com/doc/api-reference/api-methods/search-synonyms/
func (i *Index) SearchSynonyms(query string, opts ...interface{}) (res SearchSynonymsRes, err error) {
	body := newSearchSynonymsParams(query, opts...)
	path := i.path("/synonyms/search")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// ReplaceAllSynonyms replaces any existing synonyms with the given ones.
func (i *Index) ReplaceAllSynonyms(synonyms []Synonym, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.ReplaceExistingSynonyms(true))
	return i.SaveSynonyms(synonyms, opts...)
}

// BrowseSynonyms returns an iterator which will retrieve synonyms one by one from the
// index.
func (i *Index) BrowseSynonyms(opts ...interface{}) (*SynonymIterator, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.HitsPerPage(1000))
	res, err := i.SearchSynonyms("", opts...)
	if err != nil {
		return nil, fmt.Errorf("cannot browse synonyms: search failed: %v", err)
	}
	synonyms, err := res.Synonyms()
	if err != nil {
		return nil, fmt.Errorf("cannot browse synonyms: cannot decode synonyms: %v", err)
	}
	return newSynonymIterator(synonyms), nil
}
