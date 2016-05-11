package algoliasearch

import "fmt"

type Synonyms struct {
	AltCorrectionSynonyms []AltCorrectionSynonym
	OneWaySynonyms        []OneWaySynonym
	PlaceholderSynonyms   []PlaceholderSynonym
	SimpleSynonyms        []SimpleSynonym
}

type Synonym interface{}

type AltCorrectionSynonym struct {
	Corrections       []string               `json:"corrections"`
	HighlightedResult map[string]interface{} `json:"_highlightedResult"`
	ObjectID          string                 `json:"objectID"`
	Type              string                 `json:"type"`
	Word              string                 `json:"word"`
}

type OneWaySynonym struct {
	HighlightedResult map[string]interface{} `json:"_highlightedResult"`
	Input             string                 `json:"input"`
	ObjectID          string                 `json:"objectID"`
	Synonyms          []string               `json:"synonyms"`
	Type              string                 `json:"type"`
}

type PlaceholderSynonym struct {
	HighlightedResult map[string]interface{} `json:"_highlightedResult"`
	ObjectID          string                 `json:"objectID"`
	Placeholder       string                 `json:"placeholder"`
	Replacements      []string               `json:"replacements"`
	Type              string                 `json:"type"`
}

type SimpleSynonym struct {
	HighlightedResult map[string]interface{} `json:"_highlightedResult"`
	ObjectID          string                 `json:"objectID"`
	Synonyms          []string               `json:"synonyms"`
	Type              string                 `json:"type"`
}

func generateSynonyms(rawHits interface{}) (synonyms Synonyms, err error) {
	hits, ok := rawHits.([]interface{})
	if !ok {
		err = fmt.Errorf("Cannot cast `hits` to `[]interface{}`")
		return
	}

	var s Synonym
	for _, raw := range hits {
		s, err = generateSynonym(raw)
		if err != nil {
			return
		}

		switch s := s.(type) {

		case AltCorrectionSynonym:
			synonyms.AltCorrectionSynonyms = append(synonyms.AltCorrectionSynonyms, s)

		case OneWaySynonym:
			synonyms.OneWaySynonyms = append(synonyms.OneWaySynonyms, s)

		case PlaceholderSynonym:
			synonyms.PlaceholderSynonyms = append(synonyms.PlaceholderSynonyms, s)

		case SimpleSynonym:
			synonyms.SimpleSynonyms = append(synonyms.SimpleSynonyms, s)

		default:
			err = fmt.Errorf("Synonym type `%T` unknown", s)

		}
	}

	return
}

func generateSynonym(rawHit interface{}) (s Synonym, err error) {
	hit, ok := rawHit.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("Cannot cast `hit` to `map[string]interface{}`")
		return
	}

	// Check that the `type` field exists
	t, ok := hit["type"]
	if !ok {
		err = invalidParameter("type")
		return
	}

	// Check that the `type` field is a string
	t, ok = t.(string)
	if !ok {
		err = invalidParameter("type")
		return
	}

	switch t {

	case "altCorrection1", "altCorrection2", "altCorrection1|altCorrection2":
		if s, ok = rawHit.(AltCorrectionSynonym); !ok {
			err = fmt.Errorf("Cannot cast synonym to `%s`", t)
		}

	case "oneWaySynonym":
		if s, ok = rawHit.(OneWaySynonym); ok {
			err = fmt.Errorf("Cannot cast synonym to `%s`", t)
		}

	case "placeholder":
		if s, ok = rawHit.(PlaceholderSynonym); ok {
			err = fmt.Errorf("Cannot cast synonym to `%s`", t)
		}

	case "synonym":
		if s, ok = rawHit.(SimpleSynonym); ok {
			err = fmt.Errorf("Cannot cast synonym to `%s`", t)
		}

	default:
		err = fmt.Errorf("Synonym type `%s` doesn't exist`", t)

	}

	return
}

func (s *Synonyms) NbHits() int64 {
	return int64(len(s.AltCorrectionSynonyms) + len(s.OneWaySynonyms) +
		len(s.SimpleSynonyms) + len(s.PlaceholderSynonyms))
}
