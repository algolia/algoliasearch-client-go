package search

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryRes struct {
	AppliedRules          []AppliedRule                     `json:"appliedRules"`
	AroundLatLng          string                            `json:"aroundLatLng"`
	AutomaticRadius       string                            `json:"automaticRadius"`
	ExhaustiveFacetsCount bool                              `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits      bool                              `json:"exhaustiveNbHits"`
	Explain               map[string]map[string]interface{} `json:"explain"`
	Facets                map[string]map[string]int         `json:"facets"`
	FacetsStats           map[string]FacetStat              `json:"facets_stats"`
	Hits                  []map[string]interface{}          `json:"Hits"`
	HitsPerPage           int                               `json:"hitsPerPage"`
	Index                 string                            `json:"index"`
	IndexUsed             string                            `json:"indexUsed"`
	Length                int                               `json:"length"`
	Message               string                            `json:"message"`
	NbHits                int                               `json:"nbHits"`
	NbPages               int                               `json:"nbPages"`
	Offset                int                               `json:"offset"`
	Page                  int                               `json:"page"`
	Params                string                            `json:"params"`
	ParsedQuery           string                            `json:"parsedQuery"`
	ProcessingTimeMS      int                               `json:"processingTimeMS"`
	Query                 string                            `json:"query"`
	QueryAfterRemoval     string                            `json:"queryAfterRemoval"`
	QueryID               string                            `json:"queryID"`
	ServerUsed            string                            `json:"serverUsed"`
	TimeoutCounts         bool                              `json:"timeoutCounts"`
	TimeoutHits           bool                              `json:"timeoutHits"`
	UserData              map[string]interface{}            `json:"userData"`
	ABTestVariantID       string                            `json:"abTestVariantID"`
}

type AppliedRule struct {
	ObjectID string `json:"objectID"`
}

type FacetStat struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	Sum float64 `json:"sum"`
}

type HighlightResult map[string]HighlightedResult

type HighlightedResult struct {
	FullyHighlighted bool     `json:"fullyHighlighted"`
	MatchedWords     []string `json:"matchedWords"`
	MatchLevel       string   `json:"matchLevel"`
	Value            string   `json:"value"`
}

type RankingInfo struct {
	Filters           int `json:"filters"`
	FirstMatchedWord  int `json:"firstMatchedWord"`
	GeoDistance       int `json:"geoDistance"`
	GeoPrecision      int `json:"geoPrecision"`
	NbExactWords      int `json:"nbExactWords"`
	NbTypos           int `json:"nbTypos"`
	ProximityDistance int `json:"proximityDistance"`
	UserScore         int `json:"userScore"`
	Words             int `json:"words"`
}

func (r QueryRes) UnmarshalHits(v interface{}) error {
	hitsPayload, err := json.Marshal(r.Hits)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Hits from search response: %v", err)
	}
	return json.Unmarshal(hitsPayload, &v)
}

type FacetHit struct {
	Value       string `json:"value"`
	Highlighted string `json:"highlighted"`
	Count       int    `json:"count"`
}

type SearchForFacetValuesRes struct {
	FacetHits             []FacetHit
	ExhaustiveFacetsCount bool
	ProcessingTime        time.Duration
}

func (r *SearchForFacetValuesRes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var res struct {
		FacetHits             []FacetHit `json:"facetHits"`
		ExhaustiveFacetsCount bool       `json:"exhaustiveFacetsCount"`
		ProcessingTimeMS      int        `json:"processingTimeMS"`
	}

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	r.FacetHits = res.FacetHits
	r.ExhaustiveFacetsCount = res.ExhaustiveFacetsCount
	r.ProcessingTime = time.Duration(res.ProcessingTimeMS) * time.Millisecond
	return nil
}
