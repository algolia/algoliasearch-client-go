package search

import (
	"encoding/json"
	"fmt"
)

type Synonym interface {
	ObjectID() string
	Type() SynonymType
}

type SynonymType string

const (
	RegularSynonymType SynonymType = "synonym"
	OneWaySynonymType  SynonymType = "oneWaySynonym"
	AltCorrection1Type SynonymType = "altCorrection1"
	AltCorrection2Type SynonymType = "altCorrection2"
	PlaceholderType    SynonymType = "placeholder"
)

type rawSynonym struct{ impl Synonym }

func (s *rawSynonym) UnmarshalJSON(data []byte) error {
	var (
		rawSyn map[string]interface{}
		err    error
	)

	if err = json.Unmarshal(data, &rawSyn); err != nil {
		return fmt.Errorf("cannot unmarshal synonym: error when unmarshalling %v", err)
	}

	typeInterface, ok := rawSyn["type"]
	if !ok {
		return fmt.Errorf("cannot unmarshal synonym: `type` field is missing")
	}

	typeString, ok := typeInterface.(string)
	if !ok {
		return fmt.Errorf("cannot unmarshal synonym: `type` field is not a string")
	}

	switch SynonymType(typeString) {
	case RegularSynonymType:
		var syn RegularSynonym
		err = json.Unmarshal(data, &syn)
		s.impl = syn

	case OneWaySynonymType:
		var syn OneWaySynonym
		err = json.Unmarshal(data, &syn)
		s.impl = syn

	case AltCorrection1Type:
		var syn AltCorrection1
		err = json.Unmarshal(data, &syn)
		s.impl = syn

	case AltCorrection2Type:
		var syn AltCorrection2
		err = json.Unmarshal(data, &syn)
		s.impl = syn

	case PlaceholderType:
		var syn Placeholder
		err = json.Unmarshal(data, &syn)
		s.impl = syn

	default:
		return fmt.Errorf("cannot unmarshal synonym: unknown type %s", typeString)
	}

	return err
}
