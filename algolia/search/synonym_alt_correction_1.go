package search

import (
	"encoding/json"
	"fmt"
)

type AltCorrection1 struct {
	objectID    string
	Word        string
	Corrections []string
}

func NewAltCorrection1(objectID, word string, corrections ...string) AltCorrection1 {
	return AltCorrection1{objectID, word, corrections}
}

func (s AltCorrection1) ObjectID() string  { return s.objectID }
func (s AltCorrection1) Type() SynonymType { return AltCorrection1Type }

type altCorrection struct {
	ObjectID    string      `json:"objectID"`
	Type        SynonymType `json:"type"`
	Word        string      `json:"word"`
	Corrections []string    `json:"corrections"`
}

func (s AltCorrection1) MarshalJSON() ([]byte, error) {
	return json.Marshal(altCorrection{s.ObjectID(), s.Type(), s.Word, s.Corrections})
}

func (s *AltCorrection1) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var synonym altCorrection
	err := json.Unmarshal(data, &synonym)
	if err != nil {
		return fmt.Errorf("cannot unmarshal AltCorrection1: %v", err)
	}

	if synonym.Type != s.Type() && synonym.Type != altCorrection1TypeLower {
		return fmt.Errorf("cannot deserialize synonym of type %s into AltCorretion1", synonym.Type)
	}

	s.objectID = synonym.ObjectID
	s.Word = synonym.Word
	s.Corrections = synonym.Corrections
	return nil
}
