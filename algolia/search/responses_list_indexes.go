package search

import (
	"encoding/json"
	"time"
)

type ListIndexesRes struct {
	Items   []IndexRes `json:"items"`
	NbPages int        `json:"nbPages"`
}

type IndexRes struct {
	CreatedAt            time.Time     `json:"createdAt"`
	DataSize             int64         `json:"dataSize"`
	Entries              int64         `json:"entries"`
	FileSize             int64         `json:"fileSize"`
	LastBuildTime        time.Duration `json:"-"`
	Name                 string        `json:"name"`
	NumberOfPendingTasks int64         `json:"numberOfPendingTasks"`
	PendingTask          bool          `json:"pendingTask"`
	UpdatedAt            time.Time     `json:"updatedAt"`
}

func (r *IndexRes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	type indexRes IndexRes
	var res struct {
		LastBuildTimeS int64 `json:"lastBuildTimeS"`
		indexRes
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	*r = IndexRes(res.indexRes)
	r.LastBuildTime = time.Duration(res.LastBuildTimeS) * time.Second
	return nil
}
