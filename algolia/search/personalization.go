package search

type Strategy struct {
	EventsScoring map[string]EventsScoring `json:"eventsScoring"`
	FacetsScoring map[string]FacetsScoring `json:"facetsScoring"`
}

type EventsScoring struct {
	Score int    `json:"score"`
	Type  string `json:"type"`
}

type FacetsScoring struct {
	Score int `json:"score"`
}
