package search

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type TaskStatusRes struct {
	Status      string `json:"status"`
	PendingTask bool   `json:"pendingTask"`
}

type UpdateTaskRes struct {
	TaskID    int       `json:"taskID"`
	UpdatedAt time.Time `json:"updatedAt"`
	wait      func(taskID int) error
}

func (r UpdateTaskRes) Wait() error {
	return r.wait(r.TaskID)
}

type SaveObjectRes struct {
	CreatedAt time.Time `json:"createdAt"`
	ObjectID  string    `json:"objectID"`
	TaskID    int       `json:"taskID"`
	wait      func(taskID int) error
}

func (r SaveObjectRes) Wait() error {
	return r.wait(r.TaskID)
}

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int      `json:"taskID"`
	wait      func(taskID int) error
}

func (r BatchRes) Wait() error {
	return r.wait(r.TaskID)
}

type MultipleBatchRes struct {
	Responses []BatchRes
}

func (r MultipleBatchRes) Wait() error {
	var wg sync.WaitGroup
	errs := make(chan error, len(r.Responses))

	for _, res := range r.Responses {
		wg.Add(1)
		go func(wg *sync.WaitGroup, res BatchRes) {
			errs <- res.Wait()
			wg.Done()
		}(&wg, res)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return fmt.Errorf("at least one batch could not complete: %s", err)
		}
	}

	return nil
}

type DeleteTaskRes struct {
	DeletedAt time.Time `json:"deletedAt"`
	TaskID    int       `json:"taskID"`
	wait      func(taskID int) error
}

func (r DeleteTaskRes) Wait() error {
	return r.wait(r.TaskID)
}

type getObjectsRes struct {
	Results interface{} `json:"results"`
}

type SearchRes struct {
	AppliedRules          interface{}              `json:"appliedRules"` // TODO to type correctly or add an unmarshalling method
	AroundLatLng          string                   `json:"aroundLatLng"`
	AutomaticRadius       string                   `json:"automaticRadius"`
	ExhaustiveFacetsCount bool                     `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits      bool                     `json:"exhaustiveNbHits"`
	Explain               interface{}              `json:"explain"`      // TODO to type correctly or add an unmarshalling method
	Facets                interface{}              `json:"facets"`       // TODO to type correctly or add an unmarshalling method
	FacetsStats           interface{}              `json:"facets_stats"` // TODO to type correctly or add an unmarshalling method
	Hits                  []map[string]interface{} `json:"Hits"`
	HitsPerPage           int                      `json:"hitsPerPage"`
	Index                 string                   `json:"index"`
	IndexUsed             string                   `json:"indexUsed"`
	Length                int                      `json:"length"`
	Message               string                   `json:"message"`
	NbHits                int                      `json:"nbHits"`
	NbPages               int                      `json:"nbPages"`
	Offset                int                      `json:"offset"`
	Page                  int                      `json:"page"`
	Params                string                   `json:"params"`
	ParsedQuery           string                   `json:"parsedQuery"`
	ProcessingTimeMS      int                      `json:"processingTimeMS"`
	Query                 string                   `json:"query"`
	QueryAfterRemoval     string                   `json:"queryAfterRemoval"`
	QueryID               string                   `json:"queryID"`
	ServerUsed            string                   `json:"serverUsed"`
	TimeoutCounts         bool                     `json:"timeoutCounts"`
	TimeoutHits           bool                     `json:"timeoutHits"`
	UserData              interface{}              `json:"userData"` // TODO to type correctly or add an unmarshalling method
	// TODO: add and type `abTestVariantID` field
	// TODO: add and type `message` field
	// TODO: add and type `warning` field
	// TODO: add and type `cursor` field
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

func (r SearchRes) UnmarshalHits(v interface{}) error {
	hitsPayload, err := json.Marshal(r.Hits)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Hits from search response: %v", err)
	}
	return json.Unmarshal(hitsPayload, &v)
}

type browseRes struct {
	Cursor  string `json:"cursor"`
	Warning string `json:"warning"`
	SearchRes
}

type SearchRulesRes struct {
	Hits    interface{} `json:"hits"`
	NbHits  int         `json:"nbHits"`
	Page    int         `json:"page"`
	NbPages int         `json:"nbPages"`
}

func (r SearchRulesRes) UnmarshalHits(v interface{}) error {
	hitsPayload, err := json.Marshal(r.Hits)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Hits from search rules response: %v", err)
	}
	return json.Unmarshal(hitsPayload, v)
}

func (r SearchRulesRes) Rules() (rules []Rule, err error) {
	err = r.UnmarshalHits(&rules)
	return
}

type SearchSynonymsRes struct {
	Hits   []map[string]interface{} `json:"hits"`
	NbHits int                      `json:"nbHits"`
}

func (r SearchSynonymsRes) Synonyms() ([]Synonym, error) {
	var (
		synonyms []Synonym
		err      error
	)

	for i, hit := range r.Hits {
		itf, ok := hit["type"]
		if !ok {
			return nil, fmt.Errorf("cannot unmarshal SearchSynonyms response's hits: missing `type` field at position %d", i)
		}

		t, ok := itf.(string)
		if !ok {
			return nil, fmt.Errorf("cannot unmarshal SearchSynonyms response's hits: expecting `type` field as string %d but got %#v", i, itf)
		}

		switch t {

		case string(RegularSynonymType):
			var syn RegularSynonym
			err = decodeSynonym(hit, &syn)
			synonyms = append(synonyms, syn)

		case string(OneWaySynonmType):
			var syn OneWaySynonym
			err = decodeSynonym(hit, &syn)
			synonyms = append(synonyms, syn)

		case string(AltCorrection1Type):
			var syn AltCorrection1
			err = decodeSynonym(hit, &syn)
			synonyms = append(synonyms, syn)

		case string(AltCorrection2Type):
			var syn AltCorrection2
			err = decodeSynonym(hit, &syn)
			synonyms = append(synonyms, syn)

		case string(PlaceholderType):
			var syn Placeholder
			err = decodeSynonym(hit, &syn)
			synonyms = append(synonyms, syn)

		default:
			err = fmt.Errorf("cannot unmarshal SearchSynonyms response's hits: unknown `type` field %q at position %d", t, i)

		}

		if err != nil {
			return nil, err
		}
	}

	return synonyms, nil
}

func decodeSynonym(hit map[string]interface{}, syn interface{}) error {
	data, err := json.Marshal(hit)
	if err != nil {
		return fmt.Errorf("cannot unmarshal synonym: failed marshalling: %v", err)
	}
	err = json.Unmarshal(data, &syn)
	if err != nil {
		return fmt.Errorf("cannot unmarshal synonym: failed unmarshalling: %v", err)
	}
	return nil
}
