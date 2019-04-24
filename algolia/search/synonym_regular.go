package search

import (
	"encoding/json"
	"fmt"
)

type RegularSynonym struct {
	objectID string
	Synonyms []string
}

func NewRegularSynonym(objectID string, synonyms ...string) RegularSynonym {
	return RegularSynonym{objectID, synonyms}
}

func (s RegularSynonym) ObjectID() string  { return s.objectID }
func (s RegularSynonym) Type() SynonymType { return RegularSynonymType }

type regularSynonym struct {
	ObjectID string      `json:"objectID"`
	Type     SynonymType `json:"type"`
	Synonyms []string    `json:"synonyms"`
}

func (s RegularSynonym) MarshalJSON() ([]byte, error) {
	return json.Marshal(regularSynonym{s.ObjectID(), s.Type(), s.Synonyms})
}

func (s *RegularSynonym) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var synonym regularSynonym
	err := json.Unmarshal(data, &synonym)
	if err != nil {
		return fmt.Errorf("cannot unmarshal RegularSynonym: %v", err)
	}

	if synonym.Type != s.Type() {
		return fmt.Errorf("cannot deserialize synonym of type %s into RegularSynonym", synonym.Type)
	}

	s.objectID = synonym.ObjectID
	s.Synonyms = synonym.Synonyms
	return nil
}
