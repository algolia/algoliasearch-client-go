package search

import (
	"encoding/json"
	"fmt"
)

type AltCorrection2 AltCorrection1

func NewAltCorrection2(objectID, word string, corrections ...string) AltCorrection2 {
	return AltCorrection2{objectID, word, corrections}
}

func (s AltCorrection2) ObjectID() string  { return s.objectID }
func (s AltCorrection2) Type() SynonymType { return AltCorrection2Type }

func (s AltCorrection2) MarshalJSON() ([]byte, error) {
	return json.Marshal(altCorrection{s.ObjectID(), s.Type(), s.Word, s.Corrections})
}

func (s *AltCorrection2) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var synonym altCorrection
	err := json.Unmarshal(data, &synonym)
	if err != nil {
		return fmt.Errorf("cannot unmarshal AltCorrection2: %v", err)
	}

	if synonym.Type != s.Type() && synonym.Type != altCorrection2TypeLower {
		return fmt.Errorf("cannot deserialize synonym of type %s into AltCorretion2", synonym.Type)
	}

	s.objectID = synonym.ObjectID
	s.Word = synonym.Word
	s.Corrections = synonym.Corrections
	return nil
}
