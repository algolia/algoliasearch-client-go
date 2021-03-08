package search

import (
	"encoding/json"
	"time"
)

type ListIndicesRes struct {
	Items   []IndexRes `json:"items"`
	NbPages int        `json:"nbPages"`
}

type IndexRes struct {
	CreatedAt            time.Time     `json:"-"`
	DataSize             int64         `json:"dataSize"`
	Entries              int64         `json:"entries"`
	FileSize             int64         `json:"fileSize"`
	LastBuildTime        time.Duration `json:"-"`
	Name                 string        `json:"name"`
	NumberOfPendingTasks int64         `json:"numberOfPendingTasks"`
	PendingTask          bool          `json:"pendingTask"`
	UpdatedAt            time.Time     `json:"-"`
	Primary              string        `json:"primary"`
	Replicas             []string      `json:"replicas"`
}

func (r *IndexRes) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	type indexRes IndexRes
	var res struct {
		LastBuildTimeS int64  `json:"lastBuildTimeS"`
		CreatedAt      string `json:"createdAt"`
		UpdatedAt      string `json:"updatedAt"`
		indexRes
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*r = IndexRes(res.indexRes)
	r.LastBuildTime = time.Duration(res.LastBuildTimeS) * time.Second
	if res.CreatedAt != "" {
		r.CreatedAt, err = time.Parse(time.RFC3339, res.CreatedAt)
		if err != nil {
			return err
		}
	}
	if res.UpdatedAt != "" {
		r.UpdatedAt, err = time.Parse(time.RFC3339, res.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}
