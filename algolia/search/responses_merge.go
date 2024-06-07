package search

type Merge struct {
	NBSearchers        int              `json:"nbSearchers"`
	NBHitsLimit        int              `json:"nbHitslimit"`
	NBHitsMax          int              `json:"nbHitsMax"`
	NBLastHitToDisplay int              `json:"lastHitToDisplay"`
	NBHitsNumberingEnd int              `json:"nbHitsNumberingEnd"`
	NBHitsProcessed    int              `json:"nbHitsProcessed"`
	Personalization    *Personalization `json:"personalization"`
}

type Personalization struct {
	Enabled            bool                    `json:"enabled"`
	Impact             int                     `json:"impact"`
	NBPersoScanned     int                     `json:"nbPersoScanned"`
	NBPersoSelected    int                     `json:"nbPersoSelected"`
	NBPersoReranked    int                     `json:"nbPersoReranked"`
	NBPersoReturned    int                     `json:"nbPersoReturned"`
	NBPersoSkipped     int                     `json:"nbPersoSkipped"`
	Percentile         int                     `json:"percentile"`
	NBRelevanceBuckets int                     `json:"nbRelevanceBuckets"`
	Profile            *PersonalizationProfile `json:"profile"`
}

type PersonalizationProfile struct {
	TaskID int64 `json:"taskID"`
	Time   int64 `json:"time"`
	Facets map[string]int `json:"facets"`
}
