package search

// Deprecated: use recommendation.Strategy instead
type Strategy struct {
	EventsScoring map[string]EventsScoring `json:"eventsScoring"`
	FacetsScoring map[string]FacetsScoring `json:"facetsScoring"`
}

// Deprecated: use recommendation.EventsScoring instead
type EventsScoring struct {
	Score int    `json:"score"`
	Type  string `json:"type"`
}

// Deprecated: use recommendation.FacetsScoring instead
type FacetsScoring struct {
	Score int `json:"score"`
}
