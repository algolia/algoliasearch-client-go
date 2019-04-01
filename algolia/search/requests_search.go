package search

type searchReq struct {
	Params string `json:"params"`
}

type searchParams struct {
	Query string `json:"query"`
	QueryParams
}

func newSearchParams(query string, opts ...interface{}) searchParams {
	return searchParams{
		Query:       query,
		QueryParams: newQueryParams(opts...),
	}
}

type searchForFacetValuesParams struct {
	FacetQuery string `json:"facetQuery"`
	QueryParams
}

func newSearchForFacetValuesParams(query string, opts ...interface{}) searchForFacetValuesParams {
	return searchForFacetValuesParams{
		FacetQuery:  query,
		QueryParams: newQueryParams(opts...),
	}
}
