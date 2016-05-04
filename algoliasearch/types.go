package algoliasearch

import "strings"

type ListIndexesRes struct {
	Items []IndexRes `json:"items"`
}

type IndexRes struct {
	CreatedAt      string `json:"createdAt"`
	DataSize       int64  `json:"dataSize"`
	Entries        int64  `json:"entries"`
	LastBuildTimeS int64  `json:"lastBuildTimeS"`
	Name           string `json:"name"`
	PendingTask    bool   `json:"pendingTask"`
	UpdatedAt      string `json:"updatedAt"`
}

type ListKeysRes struct {
	Keys []KeyRes `json:"keys"`
}

type KeyRes struct {
	ACL      []string `json:"acl"`
	Validity int64    `json:"validity"`
	Value    string   `json:"value"`
}

type AddKeyRes struct {
	CreatedAt string `json:"createdAt"`
	Key       string `json:"key"`
}

type UpdateKeyRes struct {
	UpdatedAt string `json:"updatedAt"`
	Key       string `json:"key"`
}

type GetKeyRes struct {
	ACL      []string `json:"acl"`
	Validity int64    `json:"validity"`
	Value    string   `json:"value"`
}

type DeleteKeyRes struct {
	DeletedAt string `json:"deletedAt"`
}

type GetLogsRes struct {
	Logs []LogRes `json:"logs"`
}

type LogRes struct {
	Answer       string `json:"answer"`
	AnswerCode   int    `json:"answer_code"`
	IP           string `json:"ip"`
	Method       string `json:"method"`
	QueryBody    string `json:"query_body"`
	QueryHeaders string `json:"query_headers"`
	SHA1         string `json:"sha1"`
	Timestamp    string `json:"timestamp"`
	URL          string `json:"url"`
}

type CustomBatchRes struct {
	TaskID    string   `json:"taskID"`
	ObjectIDs []string `json:"objectIDs"`
}

type BatchRecord struct {
	IndexName string                 `json:"indexName"`
	Action    string                 `json:"action"`
	Body      map[string]interface{} `json:"body"`
}

type MultipleQueriesRes struct {
}

func transformQuery(query map[string]interface{}, ignore string) map[string]interface{} {
	norm := map[string]interface{}{}
	for k, v := range query {
		switch k {
		case ignore:
		// Ignore.
		case "analyticsTags", "optionalWords", "disableTypoToleranceOnAttributes",
			"attributesToRetrieve", "attributesToHighlight", "numericFilters",
			"tagFilters", "facets", "facetFilters":
			norm[k] = strings.Join(v.([]string), ",")

		case "attributesToSnippet":
		}
	}
}
