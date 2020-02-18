package recommendation

import "github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

type Strategy struct {
	EventsScoring         []EventsScoring                 `json:"eventsScoring"`
	FacetsScoring         []FacetsScoring                 `json:"facetsScoring"`
	PersonalizationImpact opt.PersonalizationImpactOption `json:"personalizationImpact"`
}

type EventsScoring struct {
	EventName string `json:"eventName"`
	EventType string `json:"eventType"`
	Score     int    `json:"score"`
}

type FacetsScoring struct {
	FacetName string `json:"facetName"`
	Score     int    `json:"score"`
}
