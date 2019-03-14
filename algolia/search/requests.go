package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

type deleteByReq struct {
	AroundLatLng      *opt.AroundLatLngOption      `json:"aroundLatLng,omitempty"`
	AroundRadius      *opt.AroundRadiusOption      `json:"aroundRadius,omitempty"`
	FacetFilters      *opt.FacetFiltersOption      `json:"facetFilters,omitempty"`
	Filters           *opt.FiltersOption           `json:"filters,omitempty"`
	InsideBoundingBox *opt.InsideBoundingBoxOption `json:"insideBoundingBox,omitempty"`
	InsidePolygon     *opt.InsidePolygonOption     `json:"insidePolygon,omitempty"`
	NumericFilters    *opt.NumericFiltersOption    `json:"numericFilters,omitempty"`
}

type getObjectsReq struct {
	IndexName            string                          `json:"indexName"`
	ObjectID             string                          `json:"objectID"`
	AttributesToRetrieve *opt.AttributesToRetrieveOption `json:"attributesToRetrieve,omitempty"`
}

type batchReq struct {
	Requests []BatchOperation `json:"requests,omitempty"`
}

type searchReq struct {
	Params string `json:"params"`
}

type browseReq struct {
	Cursor string `json:"cursor"`
	Params string `json:"params"`
}

type searchRulesParams struct {
	Query       string                  `json:"query"`
	Anchoring   *opt.AnchoringOption    `json:"anchoring,omitempty"`
	Context     *opt.RuleContextsOption `json:"context,omitempty"`
	Page        *opt.PageOption         `json:"page,omitempty"`
	HitsPerPage *opt.HitsPerPageOption  `json:"hitsPerPage,omitempty"`
	Enabled     *opt.EnableRulesOption  `json:"enabled,omitempty"`
}

func newSearchRulesParams(query string, opts ...interface{}) searchRulesParams {
	return searchRulesParams{
		Query:       query,
		Anchoring:   iopt.ExtractAnchoring(opts...),
		Context:     iopt.ExtractRuleContexts(opts...),
		Page:        iopt.ExtractPage(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
		Enabled:     iopt.ExtractEnableRules(opts...),
	}
}

type searchSynonymsParams struct {
	Query       string                 `json:"query"`
	Page        *opt.PageOption        `json:"page,omitempty"`
	Type        *opt.TypeOption        `json:"type,omitempty"`
	HitsPerPage *opt.HitsPerPageOption `json:"hitsPerPage,omitempty"`
}

func newSearchSynonymsParams(query string, opts ...interface{}) searchSynonymsParams {
	return searchSynonymsParams{
		Query:       query,
		Page:        iopt.ExtractPage(opts...),
		Type:        iopt.ExtractType(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
	}
}

type IndexOperation struct {
	Destination string   `json:"destination"`
	Operation   string   `json:"operation"`
	Scopes      []string `json:"scope,omitempty"`
}
