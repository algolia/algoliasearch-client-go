package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type IndexOperation struct {
	Destination string   `json:"destination"`
	Operation   string   `json:"operation"`
	Scopes      []string `json:"scope,omitempty"`
}

type deleteByReq struct {
	AroundLatLng      *opt.AroundLatLngOption      `json:"aroundLatLng,omitempty"`
	AroundRadius      *opt.AroundRadiusOption      `json:"aroundRadius,omitempty"`
	FacetFilters      *opt.FacetFiltersOption      `json:"facetFilters,omitempty"`
	Filters           *opt.FiltersOption           `json:"filters,omitempty"`
	InsideBoundingBox *opt.InsideBoundingBoxOption `json:"insideBoundingBox,omitempty"`
	InsidePolygon     *opt.InsidePolygonOption     `json:"insidePolygon,omitempty"`
	NumericFilters    *opt.NumericFiltersOption    `json:"numericFilters,omitempty"`
}

func newDeleteByReq(opts ...interface{}) deleteByReq {
	return deleteByReq{
		AroundLatLng:      iopt.ExtractAroundLatLng(opts...),
		AroundRadius:      iopt.ExtractAroundRadius(opts...),
		FacetFilters:      iopt.ExtractFacetFilters(opts...),
		Filters:           iopt.ExtractFilters(opts...),
		InsideBoundingBox: iopt.ExtractInsideBoundingBox(opts...),
		InsidePolygon:     iopt.ExtractInsidePolygon(opts...),
		NumericFilters:    iopt.ExtractNumericFilters(opts...),
	}
}

type getObjectsReq struct {
	IndexName            string                          `json:"indexName"`
	ObjectID             string                          `json:"objectID"`
	AttributesToRetrieve *opt.AttributesToRetrieveOption `json:"attributesToRetrieve,omitempty"`
}

type batchReq struct {
	Requests []BatchOperation `json:"requests,omitempty"`
}

type PartialUpdateOperation struct {
	Operation string      `json:"_operation"`
	Value     interface{} `json:"value"`
}
