package algoliasearch

type multipleQueriesRes struct {
	Results []MultipleQueryRes `json:"results"`
}

type MultipleQueryRes struct {
	Index string `json:"index"`
	QueryRes
}

type QueryRes struct {
	Hits             []map[string]interface{} `json:"hits"`
	HitsPerPage      int                      `json:"hitsPerPage"`
	NbHits           int                      `json:"nbHits"`
	NbPages          int                      `json:"nbPages"`
	Page             int                      `json:"page"`
	Params           string                   `json:"params"`
	ParsedQuery      string                   `json:"parsedQuery,omitempty"`
	ProcessingTimeMS int                      `json:"processingTimeMS"`
	Query            string                   `json:"query"`
}
