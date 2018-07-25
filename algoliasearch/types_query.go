package algoliasearch

type multipleQueriesRes struct {
	Results []MultipleQueryRes `json:"results"`
}

type MultipleQueryRes struct {
	Index     string `json:"index"`
	Processed bool   `json:"processed"`
	QueryRes
}

type QueryRes struct {
	AppliedRules          []Map  `json:"appliedRules"`
	AroundLatLng          string `json:"aroundLatLng"`
	AutomaticRadius       string `json:"automaticRadius"`
	ExhaustiveFacetsCount bool   `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits      bool   `json:"exhaustiveNbHits"`
	Explain               Map    `json:"explain"`
	Facets                Map    `json:"facets"`
	FacetsStats           Map    `json:"facets_stats"`
	Hits                  []Map  `json:"hits"`
	HitsPerPage           int    `json:"hitsPerPage"`
	Index                 string `json:"index"`
	IndexUsed             string `json:"indexUsed"`
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
	QueryID               string `json:"queryID"`
	ServerUsed            string `json:"serverUsed"`
	TimeoutCounts         bool   `json:"timeoutCounts"`
	TimeoutHits           bool   `json:"timeoutHits"`
	UserData              []Map  `json:"userData"`
}

type IndexedQuery struct {
	IndexName string
	Params    Map
}
