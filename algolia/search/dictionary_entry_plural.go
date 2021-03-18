package search

import (
	"encoding/json"
	"fmt"
)

// Plural represents an entry in the Plurals dictionary.
type Plural struct {
	objectID string
	language string
	Words    []string
}

// Plural implements DictionaryEntry
func (p Plural) ObjectID() string { return p.objectID }
func (p Plural) Language() string { return p.language }

func NewPlural(objectID, language string, words []string) Plural {
	return Plural{objectID, language, words}
}

type plural struct {
	ObjectID string   `json:"objectID"`
	Language string   `json:"language"`
	Words    []string `json:"words"`
}

func (p Plural) MarshalJSON() ([]byte, error) {
	return json.Marshal(plural{p.objectID, p.language, p.Words})
}

func (p *Plural) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var plural plural
	err := json.Unmarshal(data, &plural)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Plural dictionary entry: %v", err)
	}

	p.objectID = plural.ObjectID
	p.language = plural.Language
	p.Words = plural.Words
	return nil
}
