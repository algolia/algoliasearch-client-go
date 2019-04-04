package search

import "time"

type SetPersonalizationStrategyRes struct {
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetPersonalizationStrategyRes struct {
	TaskID int `json:"taskID"`
	Strategy
}
