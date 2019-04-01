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
	CreatedAt            time.Time
	DataSize             int64
	Entries              int64
	FileSize             int64
	LastBuildTime        time.Duration
	Name                 string
	NumberOfPendingTasks int64
	PendingTask          bool
	UpdatedAt            time.Time
}

func (r *IndexRes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	type indexRes struct {
		IndexRes
		LastBuildTimeS int `json:"lastBuildTimeS"`
	}

	var res struct {
		CreatedAt            time.Time `json:"createdAt"`
		DataSize             int64     `json:"dataSize"`
		Entries              int64     `json:"entries"`
		FileSize             int64     `json:"fileSize"`
		LastBuildTimeS       int64     `json:"lastBuildTimeS"`
		Name                 string    `json:"name"`
		NumberOfPendingTasks int64     `json:"numberOfPendingTasks"`
		PendingTask          bool      `json:"pendingTask"`
		UpdatedAt            time.Time `json:"updatedAt"`
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	r.CreatedAt = res.CreatedAt
	r.DataSize = res.DataSize
	r.Entries = res.Entries
	r.FileSize = res.FileSize
	r.LastBuildTime = time.Duration(res.LastBuildTimeS) * time.Second
	r.Name = res.Name
	r.NumberOfPendingTasks = res.NumberOfPendingTasks
	r.PendingTask = res.PendingTask
	r.UpdatedAt = res.UpdatedAt
	return nil
}
