package search

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalLogRes(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected LogRes
	}{
		{`{}`, LogRes{}},
		{`{"answer": "something"}`, LogRes{Answer: "something"}},
		{`{"exhaustive": true}`, LogRes{Exhaustive: true}},
		{`{"answer_code": "200"}`, LogRes{AnswerCode: 200}},
		{`{"ip": "127.0.0.1"}`, LogRes{IP: "127.0.0.1"}},
		{`{"index": "something"}`, LogRes{Index: "something"}},
		{`{"method": "GET"}`, LogRes{Method: "GET"}},
		{`{"nb_api_calls": "42"}`, LogRes{NbAPICalls: 42}},
		{`{"processing_time_ms": "42"}`, LogRes{ProcessingTime: 42 * time.Millisecond}},
		{`{"query_body": "something"}`, LogRes{QueryBody: "something"}},
		{`{"query_headers": "something"}`, LogRes{QueryHeaders: "something"}},
		{`{"query_nb_hits": "42"}`, LogRes{QueryNbHits: 42}},
		{`{"sha1": "something"}`, LogRes{SHA1: "something"}},
		{`{"timestamp": "2017-12-29T18:15:57Z"}`, LogRes{Timestamp: time.Date(2017, 12, 29, 18, 15, 57, 0, time.UTC)}},
		{`{"url": "something"}`, LogRes{URL: "something"}},
	} {
		var got LogRes
		err := json.Unmarshal([]byte(c.payload), &got)
		assert.NoError(t, err, "should decode payload %#v without error", c.payload)
		assert.Equal(t, c.expected, got, "should decode payload %#v correctly", c.payload)
	}
}
