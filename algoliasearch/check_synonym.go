package algoliasearch

import "fmt"

func checkSynonym(s Synonym) error {
	switch s := s.(type) {
	case AltCorrectionSynonym,
		OneWaySynonym,
		PlaceholderSynonym,
		SimpleSynonym:
		// Type is valid. Do nothing.
	default:
		return fmt.Errorf("Invalid synonym type `%T`", s)
	}

	return nil
}

func checkSynonyms(synonyms []Synonym) error {
	for _, s := range synonyms {
		if err := checkSynonym(s); err != nil {
			return err
		}
	}

	return nil
}
