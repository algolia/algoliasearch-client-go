package algoliasearch

type Key struct {
	ACL                    []string `json:"acl"`
	CreatedAt              int64    `json:"createdAt"`
	Description            string   `json:"description"`
	MaxHitsPerQuery        int64    `json:"maxHitsPerQuery"`
	MaxQueriesPerIPPerHour int64    `json:"maxQueriesPerIPPerHour"`
	QueryParamaters        string   `json:"queryParameters"`
	Referers               []string `json:"referers"`
	Validity               int64    `json:"validity"`
	Value                  string   `json:"value"`
}

type KeyRes struct {
	CreatedAt string `json:"createdAt"`
	Key       string `json:"key"`
}
