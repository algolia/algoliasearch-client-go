package algoliasearch

type FacetHit struct {
	Value       string `json:"value"`
	Highlighted string `json:"highlighted"`
	Count       int    `json:"count"`
}

type SearchFacetRes struct {
	FacetHits             []FacetHit `json:"facetHits"`
	ExhaustiveFacetsCount bool       `json:"exhaustiveFacetsCount"`
	ProcessingTimeMS      int        `json:"processingTimeMS"`
}
