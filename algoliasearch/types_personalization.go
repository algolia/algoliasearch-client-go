package algoliasearch

type Strategy struct {
	EventsScoring map[string]ScoreType `json:"eventsScoring"`
	FacetsScoring map[string]Score     `json:"facetsScoring"`
}

type ScoreType struct {
	Score int    `json:"score"`
	Type  string `json:"type"`
}

type Score struct {
	Score int `json:"score"`
}

type SetStrategyRes struct {
	UpdatedAt string `json:"updatedAt"`
}
