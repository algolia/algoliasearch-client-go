package search

import (
	"encoding/json"
	"fmt"
)

type SearchDictionariesRes struct {
	Hits    interface{} `json:"hits"`
	NbHits  int         `json:"nbHits"`
	Page    int         `json:"page"`
	NbPages int         `json:"nbPages"`
}

func (r SearchDictionariesRes) DictionaryEntries() ([]DictionaryEntry, error) {
	var (
		rawDictionaryEntries []rawDictionaryEntry
		dictionaryEntries    []DictionaryEntry
		err                  error
	)

	data, err := json.Marshal(r.Hits)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal dictionary entries: error while marshalling original dictionary entries: %v", err)
	}

	err = json.Unmarshal(data, &rawDictionaryEntries)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal dictionary entries: error while unmarshalling to intermediate type: %v", err)
	}

	for _, s := range rawDictionaryEntries {
		dictionaryEntries = append(dictionaryEntries, s.impl)
	}

	return dictionaryEntries, nil
}
