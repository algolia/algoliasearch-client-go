package algoliasearch

type Key struct {
	ACL                    []string `json:"acl"`
	CreatedAt              int64    `json:"createdAt,omitempty"`
	Description            string   `json:"description,omitempty"`
	MaxHitsPerQuery        int64    `json:"maxHitsPerQuery,omitempty"`
	MaxQueriesPerIPPerHour int64    `json:"maxQueriesPerIPPerHour,omitempty"`
	QueryParamaters        string   `json:"queryParameters,omitempty"`
	Referers               []string `json:"referers,omitempty"`
	Validity               int64    `json:"validity,omitempty"`
	Value                  string   `json:"value,omitempty"`
}

type KeyRes struct {
	CreatedAt string `json:"createdAt"`
	Key       string `json:"key"`
}
