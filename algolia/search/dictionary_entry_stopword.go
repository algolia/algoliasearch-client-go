package search

import (
	"encoding/json"
	"fmt"
)

// Stopword represents an entry in the Compounds dictionary.
type Stopword struct {
	objectID string
	language string
	Word     string
	State    string
}

// Stopword implements DictionaryEntry
func (s Stopword) ObjectID() string { return s.objectID }
func (s Stopword) Language() string { return s.language }

func NewStopword(objectID, language, word, state string) Stopword {
	return Stopword{objectID, language, word, state}
}

type stopword struct {
	ObjectID string  `json:"objectID"`
	Language string  `json:"language"`
	Word     string  `json:"word"`
	State    *string `json:"state,omitempty"`
}

func (s Stopword) MarshalJSON() ([]byte, error) {
	return json.Marshal(stopword{s.objectID, s.language, s.Word, &s.State})
}

func (s *Stopword) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var stopword stopword
	err := json.Unmarshal(data, &stopword)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Stopword dictionary entry: %v", err)
	}

	s.objectID = stopword.ObjectID
	s.language = stopword.Language
	s.Word = stopword.Word
	if stopword.State == nil {
		s.State = "enabled"
	} else {
		s.State = *stopword.State
	}
	return nil
}
