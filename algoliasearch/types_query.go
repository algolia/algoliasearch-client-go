package algoliasearch

type multipleQueriesRes struct {
	Results []MultipleQueryRes `json:"results"`
}

type MultipleQueryRes struct {
	Index string `json:"index"`
	QueryRes
}

type QueryRes struct {
	AroundLatLng          string `json:"aroundLatLng"`
	AutomaticRadius       string `json:"automaticRadius"`
	ExhaustiveFacetsCount bool   `json:"exhaustiveFacetsCount"`
	Facets                Map    `json:"facets"`
	ExhaustiveNbHits      bool   `json:"exhaustiveNbHits"`
	FacetsStats           Map    `json:"facets_stats"`
	Hits                  []Map  `json:"hits"`
	HitsPerPage           int    `json:"hitsPerPage"`
	Index                 string `json:"index"`
	Length                int    `json:"length"`
	Message               string `json:"message"`
	NbHits                int    `json:"nbHits"`
	NbPages               int    `json:"nbPages"`
	Offset                int    `json:"offset"`
	Page                  int    `json:"page"`
	Params                string `json:"params"`
	ParsedQuery           string `json:"parsedQuery"`
	ProcessingTimeMS      int    `json:"processingTimeMS"`
	Query                 string `json:"query"`
	QueryAfterRemoval     string `json:"queryAfterRemoval"`
	ServerUsed            string `json:"serverUsed"`
	TimeoutCounts         bool   `json:"timeoutCounts"`
	TimeoutHits           bool   `json:"timeoutHits"`
}

type IndexedQuery struct {
	IndexName string
	Params    Map
}
