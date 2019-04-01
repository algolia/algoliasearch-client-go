package search

import (
	"encoding/json"
	"fmt"
	"strconv"
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

func (r MultipleBatchRes) ObjectIDs() []string {
	var objectIDs []string
	for _, res := range r.Responses {
		objectIDs = append(objectIDs, res.ObjectIDs...)
	}
	return objectIDs
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
		rawSynonyms []rawSynonym
		synonyms    []Synonym
		err         error
	)

	data, err := json.Marshal(r.Hits)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal synonyms: error while marshalling original synonyms: %v", err)
	}

	err = json.Unmarshal(data, &rawSynonyms)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal synonyms: error while unmarshalling to intermediate type: %v", err)
	}

	for _, s := range rawSynonyms {
		synonyms = append(synonyms, s.impl)
	}

	return synonyms, nil
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

type ListIndexesRes struct {
	Items   []IndexRes `json:"items"`
	NbPages int        `json:"nbPages"`
}

type CreateKeyRes struct {
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	wait      func() error
}

func (r CreateKeyRes) Wait() error {
	return r.wait()
}

type RestoreKeyRes struct {
	CreatedAt time.Time `json:"createdAt"`
	wait      func() error
}

func (r RestoreKeyRes) Wait() error {
	return r.wait()
}

type UpdateKeyRes struct {
	Key       string    `json:"key"`
	UpdatedAt time.Time `json:"updatedAt"`
	wait      func() error
}

func (r UpdateKeyRes) Wait() error {
	return r.wait()
}

type DeleteKeyRes struct {
	DeletedAt time.Time `json:"deletedAt"`
	wait      func() error
}

func (r DeleteKeyRes) Wait() error {
	return r.wait()
}

type ListAPIKeysRes struct {
	Keys []Key `json:"keys"`
}

type GetLogsRes struct {
	Logs []LogRes `json:"logs"`
}

type LogRes struct {
	Answer         string
	AnswerCode     int
	IP             string
	Method         string
	NbAPICalls     int
	ProcessingTime time.Duration
	QueryBody      string
	QueryHeaders   string
	QueryNbHits    int
	SHA1           string
	Timestamp      time.Time
	URL            string
}

func (res *LogRes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var m map[string]string
	var i int

	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		switch k {
		case "answer":
			res.Answer = v
		case "answer_code":
			res.AnswerCode, err = strconv.Atoi(v)
		case "ip":
			res.IP = v
		case "method":
			res.Method = v
		case "nb_api_calls":
			res.NbAPICalls, err = strconv.Atoi(v)
		case "processing_time_ms":
			i, err = strconv.Atoi(v)
			res.ProcessingTime = time.Duration(i) * time.Millisecond
		case "query_body":
			res.QueryBody = v
		case "query_headers":
			res.QueryHeaders = v
		case "query_nb_hits":
			res.QueryNbHits, err = strconv.Atoi(v)
		case "sha1":
			res.SHA1 = v
		case "timestamp":
			res.Timestamp, err = time.Parse(time.RFC3339, v)
		case "url":
			res.URL = v
		}
		if err != nil {
			return err
		}
	}

	return nil
}
