package search

import (
	"encoding/json"
	"fmt"
)

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
