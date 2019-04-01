package search

import (
	"encoding/json"
	"fmt"
)

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
