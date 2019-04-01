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
	Answer         string
	AnswerCode     int
	IP             string
	Method         string
	NbAPICalls     int
	ProcessingTime time.Duration
	QueryBody      string
	QueryHeaders   string
	QueryNbHits    int
	SHA1           string
	Timestamp      time.Time
	URL            string
}

func (res *LogRes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var m map[string]string
	var i int

	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		switch k {
		case "answer":
			res.Answer = v
		case "answer_code":
			res.AnswerCode, err = strconv.Atoi(v)
		case "ip":
			res.IP = v
		case "method":
			res.Method = v
		case "nb_api_calls":
			res.NbAPICalls, err = strconv.Atoi(v)
		case "processing_time_ms":
			i, err = strconv.Atoi(v)
			res.ProcessingTime = time.Duration(i) * time.Millisecond
		case "query_body":
			res.QueryBody = v
		case "query_headers":
			res.QueryHeaders = v
		case "query_nb_hits":
			res.QueryNbHits, err = strconv.Atoi(v)
		case "sha1":
			res.SHA1 = v
		case "timestamp":
			res.Timestamp, err = time.Parse(time.RFC3339, v)
		case "url":
			res.URL = v
		}
		if err != nil {
			return err
		}
	}

	return nil
}
