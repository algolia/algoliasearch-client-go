package search

import (
	"encoding/json"
	"fmt"
)

type OneWaySynonym struct {
	objectID string
	Input    string
	Synonyms []string
}

func NewOneWaySynonym(objectID, input string, synonyms ...string) OneWaySynonym {
	return OneWaySynonym{objectID, input, synonyms}
}

func (s OneWaySynonym) ObjectID() string  { return s.objectID }
func (s OneWaySynonym) Type() SynonymType { return OneWaySynonymType }

type oneWaySynonym struct {
	ObjectID string      `json:"objectID"`
	Type     SynonymType `json:"type"`
	Input    string      `json:"input"`
	Synonyms []string    `json:"synonyms"`
}

func (s OneWaySynonym) MarshalJSON() ([]byte, error) {
	return json.Marshal(oneWaySynonym{s.ObjectID(), s.Type(), s.Input, s.Synonyms})
}

func (s *OneWaySynonym) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var synonym oneWaySynonym
	err := json.Unmarshal(data, &synonym)
	if err != nil {
		return fmt.Errorf("cannot unmarshal OneWaySynonym: %v", err)
	}

	if synonym.Type != s.Type() && synonym.Type != oneWaySynonymTypeLower {
		return fmt.Errorf("cannot deserialize synonym of type %s into OneWaySynonym", synonym.Type)
	}

	s.objectID = synonym.ObjectID
	s.Input = synonym.Input
	s.Synonyms = synonym.Synonyms
	return nil
}
