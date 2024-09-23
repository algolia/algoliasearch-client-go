package search

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryRes struct {
	AIReRanking                *AIReRanking                      `json:"aiReRanking"`
	AppliedRules               []AppliedRule                     `json:"appliedRules"`
	AppliedRelevancyStrictness int                               `json:"appliedRelevancyStrictness"`
	AroundLatLng               string                            `json:"aroundLatLng"`
	AutomaticRadius            string                            `json:"automaticRadius"`
	ExhaustiveFacetsCount      bool                              `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits           bool                              `json:"exhaustiveNbHits"`
	Explain                    map[string]map[string]interface{} `json:"explain"`
	Extensions                 map[string]map[string]interface{} `json:"extensions"`
	Facets                     map[string]map[string]int         `json:"facets"`
	FacetsStats                map[string]FacetStat              `json:"facets_stats"`
	Hits                       []map[string]interface{}          `json:"hits"`
	HitsPerPage                int                               `json:"hitsPerPage"`
	Index                      string                            `json:"index"`
	IndexUsed                  string                            `json:"indexUsed"`
	Length                     int                               `json:"length"`
	Merge                      *Merge                            `json:"merge"`
	Message                    string                            `json:"message"`
	NbHits                     int                               `json:"nbHits"`
	NbPages                    int                               `json:"nbPages"`
	NbSortedHits               int                               `json:"nbSortedHits"`
	Offset                     int                               `json:"offset"`
	Page                       int                               `json:"page"`
	Params                     string                            `json:"params"`
	ParsedQuery                string                            `json:"parsedQuery"`
	ProcessingTimeMS           int                               `json:"processingTimeMS"`
	Query                      string                            `json:"query"`
	QueryAfterRemoval          string                            `json:"queryAfterRemoval"`
	QueryID                    string                            `json:"queryID"`
	ServerUsed                 string                            `json:"serverUsed"`
	TimeoutCounts              bool                              `json:"timeoutCounts"`
	TimeoutHits                bool                              `json:"timeoutHits"`
	UserData                   []interface{}                     `json:"userData"`
	ABTestVariantID            int                               `json:"abTestVariantID"`
	ABTestID                   uint32                            `json:"abTestID"`
	RenderingContent           *RenderingContent                 `json:"renderingContent"`
	AutomaticInsights          *bool                             `json:"_automaticInsights,omitempty"`
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

func (r QueryRes) UnmarshalUserData(v interface{}) error {
	userDataPayload, err := json.Marshal(r.UserData)
	if err != nil {
		return fmt.Errorf("cannot unmarshal UserData from search response: %v", err)
	}
	return json.Unmarshal(userDataPayload, &v)
}

// GetObjectPosition returns the position (0-based) within the `Hits`
// result list of the record matching against the given objectID. If the
// objectID is not found, -1 is returned.
func (r QueryRes) GetObjectPosition(objectID string) int {
	for i, hit := range r.Hits {
		itf, ok := hit["objectID"]
		if !ok {
			continue
		}
		hitObjectID, ok := itf.(string)
		if ok && hitObjectID == objectID {
			return i
		}
	}
	return -1
}

// GetObjectIDPosition does the same as GetObjectPosition.
//
// Deprecated: use GetObjectPosition instead.
func (r QueryRes) GetObjectIDPosition(objectID string) int {
	return r.GetObjectPosition(objectID)
}

type ObjectWithPosition struct {
	Object   map[string]interface{}
	Position int
	Page     int
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
	if string(data) == jsonNull {
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
