package algoliasearch

type SearchResult struct {
	Hits             []map[string]interface{}
	Page             int
	NbHits           int
	NbPages          int
	HitsPerPage      int
	ProcessingTimeMS int
	Facets           map[string]interface{}
	FacetStats       map[string]interface{} `json:"facet_stats"`
	Query            string
	Params           string
}
