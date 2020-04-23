package search

import (
	"encoding/json"
	"strconv"
	"time"
)

type GetLogsRes struct {
	Logs []LogRes `json:"logs"`
}

type LogRes struct {
	Answer         string          `json:"answer"`
	AnswerCode     int             `json:"-"`
	Exhaustive     bool            `json:"exhaustive"`
	IP             string          `json:"ip"`
	Index          string          `json:"index"`
	InnerQueries   []InnerQueryRes `json:"inner_queries"`
	Method         string          `json:"method"`
	NbAPICalls     int             `json:"-"`
	ProcessingTime time.Duration   `json:"-"`
	QueryBody      string          `json:"query_body"`
	QueryHeaders   string          `json:"query_headers"`
	QueryNbHits    int             `json:"-"`
	SHA1           string          `json:"sha1"`
	Timestamp      time.Time       `json:"timestamp"`
	URL            string          `json:"url"`
}

type InnerQueryRes struct {
	IndexName string `json:"index_name"`
	QueryID   string `json:"query_id"`
	Offset    int    `json:"offset"`
	UserToken string `json:"user_token"`
}

func (res *LogRes) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	type Alias LogRes
	var tmp Alias
	_ = json.Unmarshal(data, &tmp)
	*res = LogRes(tmp)

	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		switch k {
		case "answer_code":
			if v, ok := v.(string); ok {
				res.AnswerCode, err = strconv.Atoi(v)
			}
		case "nb_api_calls":
			if v, ok := v.(string); ok {
				res.NbAPICalls, err = strconv.Atoi(v)
			}
		case "processing_time_ms":
			if v, ok := v.(string); ok {
				var i int
				i, err = strconv.Atoi(v)
				res.ProcessingTime = time.Duration(i) * time.Millisecond
			}
		case "query_nb_hits":
			if v, ok := v.(string); ok {
				res.QueryNbHits, err = strconv.Atoi(v)
			}
		}
		if err != nil {
			return err
		}
	}

	return nil
}
