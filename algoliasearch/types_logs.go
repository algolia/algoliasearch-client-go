package algoliasearch

type getLogsRes struct {
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
