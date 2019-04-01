package search

import "github.com/algolia/algoliasearch-client-go/algolia/transport"

type IndexedGetObject struct {
	IndexName            string `json:"indexName"`
	ObjectID             string `json:"objectID"`
	AttributesToRetrieve string `json:"attributesToRetrieve,omitempty"`
}

type multipleQueriesReq struct {
	Requests []indexedRequest `json:"requests"`
	Strategy string           `json:"strategy"`
}

type indexedRequest struct {
	IndexName string `json:"indexName"`
	Params    string `json:"params"`
}

func newMultipleQueriesReq(queries []IndexedQuery, strategy string) multipleQueriesReq {
	if strategy == "" {
		strategy = "none"
	}

	var requests []indexedRequest
	for _, q := range queries {
		requests = append(requests, indexedRequest{
			IndexName: q.IndexName,
			Params:    transport.URLEncode(q.QueryParams),
		})
	}

	return multipleQueriesReq{
		Requests: requests,
		Strategy: strategy,
	}
}

type IndexedQuery struct {
	IndexName string `json:"indexName"`
	QueryParams
}

func NewIndexedQuery(index string, opts ...interface{}) IndexedQuery {
	return IndexedQuery{
		IndexName:   index,
		QueryParams: newQueryParams(opts...),
	}
}
