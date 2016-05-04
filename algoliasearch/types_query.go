package algoliasearch

type multipleQueriesRes struct {
	results []MultipleQueryRes
}

type MultipleQueryRes struct {
	Index string `json:"index"`
	QueryRes
}

type QueryRes struct {
	Hits             []map[string]interface{} `json:"hits"`
	HitsPerPage      int64                    `json:"hitsPerPage"`
	NbHits           int64                    `json:"nbHits"`
	NbPages          int64                    `json:"nbPages"`
	Page             int64                    `json:"page"`
	Params           string                   `json:"params"`
	ParsedQuery      string                   `json:"parsedQuery"`
	ProcessingTimeMS int64                    `json:"processingTimeMS"`
	Query            string                   `json:"query"`
}
