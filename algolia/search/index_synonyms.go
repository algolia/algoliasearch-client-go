package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (i *Index) GetSynonym(objectID string, opts ...interface{}) (synonym Synonym, err error) {
	var syn rawSynonym
	path := i.path("/synonyms/%s", objectID)
	err = i.transport.Request(&syn, http.MethodGet, path, nil, call.Read, opts...)
	if err == nil {
		synonym = syn.impl
	}
	return
}

func (i *Index) SaveSynonym(synonym Synonym, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/synonyms/%s", synonym.ObjectID())
	err = i.transport.Request(&res, http.MethodPut, path, synonym, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) SaveSynonyms(synonyms []Synonym, opts ...interface{}) (res UpdateTaskRes, err error) {
	if replaceExistingSynonyms := iopt.ExtractReplaceExistingSynonyms(opts...); replaceExistingSynonyms != nil {
		opts = opt.InsertExtraURLParam(opts, "replaceExistingSynonyms", replaceExistingSynonyms.Get())
	}
	path := i.path("/synonyms/batch")
	err = i.transport.Request(&res, http.MethodPost, path, synonyms, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) ClearSynonyms(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/synonyms/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) DeleteSynonym(objectID string, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/synonyms/%s", objectID)
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) SearchSynonyms(query string, opts ...interface{}) (res SearchSynonymsRes, err error) {
	body := newSearchSynonymsParams(query, opts...)
	path := i.path("/synonyms/search")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

func (i *Index) ReplaceAllSynonyms(synonyms []Synonym, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.ReplaceExistingSynonyms(true))
	return i.SaveSynonyms(synonyms, opts...)
}

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
